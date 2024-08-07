package adds

import (
	"fmt"
	"net"
	"time"
)

func ScanPortUDP(ip string, port int) string {
	address := fmt.Sprintf("%s:%d", ip, port)
	timeout := 5 * time.Second
	conn, err := net.DialTimeout("udp", address, timeout)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			return Filtered
		}
		return Closed
	}
	conn.Close()
	return Open
}

func WorkerUDP(ip string, ports, results chan int, openPorts chan ServiceVersion, done chan bool, services map[int]string) {
	for port := range ports {
		state := ScanPortUDP(ip, port)
		service := DetectService(port, services)
		results <- port
		if state == Open {
			openPorts <- service
		}
		fmt.Printf("Port %d: %s, Service: %s, Response: %s\n", port, state, service.Service, service.Response)
	}
	done <- true
}
