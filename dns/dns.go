package dns

import (
	"github.com/miekg/dns"
)

func DnsLookUp(domain string) ([]string, error) {
	var IpList []string
	var msg dns.Msg

	fqdn := dns.Fqdn(domain)

	msg.SetQuestion(fqdn, dns.TypeA)

	in, err := dns.Exchange(&msg, "8.8.8.8:53")
	if err != nil {
		return IpList, err
	}

	if len(in.Answer) < 1 {
		return IpList, nil
	}

	for _, answer := range in.Answer {
		if a, ok := answer.(*dns.A); ok {
			IpList = append(IpList, a.A.String())
		}
	}

	return IpList, nil
}
