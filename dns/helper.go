package dns

import (
	"errors"

	"github.com/miekg/dns"
)

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
