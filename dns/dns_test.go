package dns

import "testing"

func TestLooukupA(t *testing.T) {
	_, err := LookupA("facebook.com", "8.8.8.8:53")

	if err != nil {
		t.Errorf("Error encountered: %v", err)
	}
}

func TestLooukupCname(t *testing.T) {
	_, err := LookupCname("tr.facebook.com", "8.8.8.8:53")

	if err != nil {
		t.Errorf("Error encountered: %v", err)
	}
}

func TestLooukupSubdomain(t *testing.T) {
	results := lookupSubDomain("tr.facebook.com", "8.8.8.8:53")

	if len(results) == 0 {
		t.Errorf("No results returned")
	}
}

func TestSubdomainScanner(t *testing.T) {
	wordlist := []string{"tr", "azxcafca", "en"}

	results, _ := SubDomainScanner("facebook.com", "8.8.8.8:53", 1, wordlist)

	if len(results) == 0 {
		t.Errorf("No results returned")
	}
}
