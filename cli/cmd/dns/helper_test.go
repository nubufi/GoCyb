package dns

import (
	"testing"

	"github.com/miekg/dns"
)

func TestLookUpRecord(t *testing.T) {
	records, _ := lookUpRecord("facebook.com", "8.8.8.8:53", dns.TypeA)

	if len(records) == 0 {
		t.Errorf("No records found")
	}
}
