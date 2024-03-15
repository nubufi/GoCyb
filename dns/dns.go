package dns

import (
	"github.com/miekg/dns"
)

func LookupA(domain, server string) ([]string, error) {
	ipList, err := lookUpRecord(domain, server, dns.TypeA)

	return ipList, err
}

func LookupCname(domain, server string) ([]string, error) {
	cnames, err := lookUpRecord(domain, server, dns.TypeCNAME)

	return cnames, err
}

type result struct {
	Hostname string
	Ip       string
}

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

func SubDomainScanner(domain, server string, workers int, wordList []string) []result {
	var results []result
	var fqdns = make(chan string, workers)
	var gather = make(chan []result)
	var tracker = make(chan empty)

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

	return results
}
