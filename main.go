package main

import (
	"fmt"

	"github.com/nubufi/GoCyb/dns"
)

func main() {
	IpList, _ := dns.DnsLookUp("test.soilprime.com")

	fmt.Println(IpList)
}
