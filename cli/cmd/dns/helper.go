package dns

import (
	"errors"

	"github.com/miekg/dns"
)

// lookUpRecord performs a DNS lookup for the specified domain, using the given server and record type.
// It returns a slice of strings containing the IP addresses associated with the domain, and an error if any.
//
// Parameters:
//
// - domain: The domain to lookup.
//
// - server: The DNS server to use for the lookup.
//
// - recordType: The type of DNS record to lookup.
//
// Returns:
//
// - []string: A list of IP addresses associated with the domain.
//
// - error: Any error encountered during the lookup.
//
// Example:
//
// - ips, err := lookUpRecord("example.com", "8.8.8.8:53", dns.TypeA)
func lookUpRecord(domain, server string, recordType uint16) ([]string, error) {
	var records []string
	var msg dns.Msg

	fqdn := dns.Fqdn(domain)

	msg.SetQuestion(fqdn, recordType)

	in, err := dns.Exchange(&msg, server)
	if err != nil {
		return records, err
	}

	if len(in.Answer) < 1 {
		return records, errors.New("no records found")
	}

	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			records = append(records, a.A.String())
		}
	}

	return records, nil
}

type empty struct{}

func worker(tracker chan empty, fqdns chan string, gather chan []result, server string) {
	for fqdn := range fqdns {
		results := lookupSubDomain(fqdn, server)
		if len(results) > 0 {
			gather <- results
		}
	}
	var e empty
	tracker <- e
}
