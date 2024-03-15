package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/nubufi/GoCyb/dns"
)

var wordList = []string{
	"ajax", "tr", "buy", "news", "ra", "smtp", "en", "asdasdasd", "test",
}

func main() {
	results := dns.SubDomainScanner("soilprime.com", "8.8.8.8:53", 5, wordList)

	w := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	for _, r := range results {
		fmt.Fprintf(w, "%s\t%s\n", r.Hostname, r.Ip)
	}
	w.Flush()
}
