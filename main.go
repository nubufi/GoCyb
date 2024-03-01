package main

import (
	"fmt"

	"github.com/nubufi/GoCyb/shodan"
)

func main() {
	s := shodan.New(shodan.API_KEY)

	info, err := s.APIInfo()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Query Credits: %d\nScan Credits: %d\n\n",
		info.QueryCredits,
		info.ScanCredits)

	hostSearch, err := s.HostSearch("tomcat")
	if err != nil {
		panic(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
