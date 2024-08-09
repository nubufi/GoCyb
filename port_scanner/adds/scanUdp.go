package adds

import (
	"fmt"
	"net"
	"time"
)

// ScanPortUDP scans a UDP port on a given IP address to determine its state.
//
// Parameters:
// - ip: The IP address to scan.
// - port: The port number to scan.
//
// Returns:
// - string: The state of the port, which can be "Open", "Closed", or "Filtered".
//
// The function attempts to establish a UDP connection to the specified IP address and port.
// If the connection is successful, the port is considered "Open".
// If the connection fails due to a timeout or other network error, the port is considered "Filtered".
// If any other error occurs, the port is considered "Closed".
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

// WorkerUDP is a worker function that scans UDP ports and sends the results to channels.
//
// Parameters:
// - ip: The IP address to scan.
// - ports: A channel that provides port numbers to scan.
// - results: A channel to send the scanned port numbers.
// - openPorts: A channel to send details of open ports and their detected services.
// - done: A channel to signal when the worker has finished processing all ports.
// - services: A map of known services with port numbers as keys and service names as values.
//
// The function retrieves port numbers from the `ports` channel, scans them using `ScanPortUDP`,
// and detects services using `DetectService`. It sends the results to the `results` channel and
// details of open ports to the `openPorts` channel. When all ports have been processed, the function
// sends a signal through the `done` channel.
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
