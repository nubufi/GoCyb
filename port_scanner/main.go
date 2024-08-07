package main

import (
	"fmt"
	"net"
	"os"
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
	OpenPortsUDP  chan adds.ServiceVersion
	Done          chan bool
	Services      map[int]string
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
		OpenPortsUDP:  make(chan adds.ServiceVersion, len(ports)),
		Done:          make(chan bool, numWorkers*2),
		Services:      make(map[int]string),
	}, nil
}

func (t *Target) Scan() {
	for i := 0; i < t.NumWorkers; i++ {
		go adds.WorkerTCP("", t.PortChannel, t.ResultChannel, t.OpenPorts, t.Done, t.Services)
		go adds.WorkerUDP("", t.PortChannel, t.ResultChannel, t.OpenPortsUDP, t.Done, t.Services)
	}

	for _, ip := range t.IPs {
		fmt.Printf("Scanning IP: %s\n", ip)

		for _, port := range t.Ports {
			fmt.Printf("Enqueueing port %d\n", port)
			t.PortChannel <- port
		}
		close(t.PortChannel)

		for range t.Ports {
			<-t.ResultChannel
		}

		for i := 0; i < t.NumWorkers*2; i++ {
			<-t.Done
		}
		close(t.OpenPorts)
		close(t.OpenPortsUDP)

		file, err := os.Create("output.txt")
		if err != nil {
			fmt.Printf("Error creating file: %s\n", err)
			return
		}
		defer file.Close()

		_, err = file.WriteString("Open TCP Ports with Services:\n")
		if err != nil {
			fmt.Printf("Error writing to file: %s\n", err)
			return
		}
		for service := range t.OpenPorts {
			_, err = file.WriteString(fmt.Sprintf("Port %d (TCP) is Open, Service: %s\n", service.Port, service.Service))
			if err != nil {
				fmt.Printf("Error writing to file: %s\n", err)
				return
			}
		}

		_, err = file.WriteString("Open UDP Ports with Services:\n")
		if err != nil {
			fmt.Printf("Error writing to file: %s\n", err)
			return
		}
		for service := range t.OpenPortsUDP {
			_, err = file.WriteString(fmt.Sprintf("Port %d (UDP) is Open, Service: %s\n", service.Port, service.Service))
			if err != nil {
				fmt.Printf("Error writing to file: %s\n", err)
				return
			}
		}
	}
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
