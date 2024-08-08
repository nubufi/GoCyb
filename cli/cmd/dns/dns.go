package dns

import (
	"errors"

	"github.com/miekg/dns"
)

// LookupA performs a DNS lookup for the A record of the specified domain using the specified server.
// It returns a list of IP addresses associated with the domain and any error encountered during the lookup.
//
// Parameters:
//
//   - domain: The domain to lookup.
//
//   - server: The DNS server to use for the lookup.
//
// Returns:
//
// - []string: A list of IP addresses associated with the domain.
//
// - error: Any error encountered during the lookup.
//
// Example:
//
// - ips, err := LookupA("example.com", "8.8.8.8:53")
func LookupA(domain, server string) ([]string, error) {
	ipList, err := lookUpRecord(domain, server, dns.TypeA)

	return ipList, err
}

// LookupCname performs a DNS lookup for the CNAME records of the specified domain using the given server.
// It returns a slice of strings containing the CNAME records and an error if any.
//
// Parameters:
//
// - domain: The domain to lookup.
//
// - server: The DNS server to use for the lookup.
//
// Returns:
//
// - []string: A list of CNAME records associated with the domain.
//
// - error: Any error encountered during the lookup.
//
// Example:
//
// - cnames, err := LookupCname("test.example.com", "8.8.8.8:53")
func LookupCname(domain, server string) ([]string, error) {
	cnames, err := lookUpRecord(domain, server, dns.TypeCNAME)

	return cnames, err
}

type result struct {
	Hostname string
	Ip       string
}

// lookupSubDomain performs a DNS lookup for a given domain and server.
// It returns a slice of results, where each result contains the canonical FQDN (Fully Qualified Domain Name) and its corresponding IP address.
//
// Parameters:
//
// - domain: The domain to lookup.
//
// - server: The DNS server to use for the lookup.
//
// Returns:
//
// - []result: A list of results, where each result contains the canonical FQDN and its corresponding IP address.
//
// Example:
//
// - results := lookupSubDomain("test.example.com", "8.8.8.8:53")
func lookupSubDomain(domain, server string) []result {
	var results []result
	var cfqdn = domain

	for {
		cnames, err := LookupCname(cfqdn, server)
		if err == nil && len(cnames) > 0 {
			cfqdn = cnames[0]
			continue
		}
		ips, err := LookupA(cfqdn, server)
		if err != nil {
			break
		}
		for _, ip := range ips {
			results = append(results, result{cfqdn, ip})
		}
		break
	}
	return results
}

// SubDomainScanner scans for subdomains of a given domain using a specified DNS server.
// It utilizes multiple workers to concurrently query the DNS server for each subdomain in the wordList.
//
// Parameters:
//
// - domain: The domain to scan for subdomains.
//
// - server: The DNS server to use for the lookup.
//
// - workers: The number of workers to use for the scan.
//
// - wordList: A list of subdomains to scan.
//
// Returns:
//
// - []result: A list of results, where each result contains the canonical FQDN and its corresponding IP address.
//
// Example:
//
// - results := SubDomainScanner("example.com", "8.8.8.8:53", 200, wordList)
func SubDomainScanner(domain, server string, workers int, wordList []string) ([]result, error) {
	var results []result
	var fqdns = make(chan string, workers)
	var gather = make(chan []result)
	var tracker = make(chan empty)

	if workers < 2 {
		return results, errors.New("workers must be greater than 1")
	}

	for i := 0; i < workers; i++ {
		go worker(tracker, fqdns, gather, server)
	}

	for _, word := range wordList {
		fqdns <- word + "." + domain
	}

	go func() {
		for r := range gather {
			results = append(results, r...)
		}
		var e empty
		tracker <- e
	}()

	close(fqdns)
	for i := 0; i < workers; i++ {
		<-tracker
	}
	close(gather)

	if len(results) == 0 {
		return results, errors.New("no results found")
	}

	return results, nil
}
