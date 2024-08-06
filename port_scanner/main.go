package main

import (
	"fmt"
	"net"
	"port/adds"
)

type Target struct {
	Domain        string
	IPs           []string
	Ports         []int
	NumWorkers    int
	PortChannel   chan int
	ResultChannel chan int
	OpenPorts     chan adds.ServiceVersion
	Done          chan bool
}

func NewTarget(domain string, numWorkers int) (*Target, error) {
	ips, err := net.LookupHost(domain)
	if err != nil {
		return nil, err
	}

	ports := make([]int, 0, 65535)
	for port := 1; port <= 65535; port++ {
		ports = append(ports, port)
	}

	return &Target{
		Domain:        domain,
		IPs:           ips,
		Ports:         ports,
		NumWorkers:    numWorkers,
		PortChannel:   make(chan int, len(ports)),
		ResultChannel: make(chan int, len(ports)),
		OpenPorts:     make(chan adds.ServiceVersion, len(ports)),
		Done:          make(chan bool, numWorkers),
	}, nil
}

func (t *Target) Scan() {
	for i := 0; i < t.NumWorkers; i++ {
		go adds.Worker("", t.PortChannel, t.ResultChannel, t.OpenPorts, t.Done, adds.Servicess)
	}

	for _, ip := range t.IPs {
		fmt.Printf("Scanning IP: %s\n", ip)

		for _, port := range t.Ports {
			t.PortChannel <- port
		}
		close(t.PortChannel)

		for range t.Ports {
			<-t.ResultChannel
		}

		for i := 0; i < t.NumWorkers; i++ {
			<-t.Done
		}
		close(t.OpenPorts)

		fmt.Println("Open Ports with Services:")
		for service := range t.OpenPorts {
			fmt.Printf("Port %d is Open, Service: %s\n", service.Port, service.Service)
		}
	}

	close(t.ResultChannel)
}

func main() {
	var domain string
	fmt.Print("Enter domain: ")
	fmt.Scanln(&domain)

	target, err := NewTarget(domain, 100)
	if err != nil {
		fmt.Printf("Error resolving domain: %s\n", err)
		return
	}

	target.Scan()
}
