package scanner

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func worker(domain string, ports, results chan int) {
	for port := range ports {
		err := scan(port, domain)
		if err != nil {
			results <- 0
			continue
		}
		results <- port
	}
}

func scan(port int, domain string) error {
	address := fmt.Sprintf("%s:%d", domain, port)
	d := net.Dialer{Timeout: 30 * time.Second}
	conn, err := d.Dial("tcp", address)
	if err != nil {
		return err
	}
	conn.Close()
	return nil
}

func PortScanner(domain string, numberOfPorts int) []int {
	ports := make(chan int, min(numberOfPorts, 100))
	results := make(chan int)

	var openPorts []int

	for i := 0; i < cap(ports); i++ {
		go worker(domain, ports, results)
	}

	go func() {
		for i := 1; i < numberOfPorts; i++ {
			ports <- i
		}
	}()

	for i := 0; i < numberOfPorts; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)

	return openPorts
}
