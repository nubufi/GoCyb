package adds

import (
	"fmt"
	"net"
	"time"
)

const (
	Open     = "Open"
	Closed   = "Closed"
	Filtered = "Filtered"
)

type ServiceVersion struct {
	Port     int
	Protocol string
	Service  string
	Response string
}

func ScanPort(ip string, port int) string {
	address := fmt.Sprintf("%s:%d", ip, port)

	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			return Filtered
		}
		return Closed
	}
	conn.Close()
	return Open
}

func DetectService(port int, services map[int]string) ServiceVersion {
	if svc, ok := services[port]; ok {
		return ServiceVersion{Port: port, Protocol: "tcp", Service: svc, Response: "Service Detected"}
	}
	return ServiceVersion{Port: port, Protocol: "Unknown", Service: "Unknown", Response: "Service Not Detected"}
}

func Worker(ip string, ports, results chan int, openPorts chan ServiceVersion, done chan bool, services map[int]string) {
	for port := range ports {
		state := ScanPort(ip, port)
		service := DetectService(port, services)
		results <- port
		if state == Open {
			openPorts <- service
		}
		fmt.Printf("Port %d: %s, Service: %s, Response: %s\n", port, state, service.Service, service.Response)
	}
	done <- true
}
