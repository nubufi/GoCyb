package portscanner

import (
	"fmt"
	"net"
	"time"
)

func ScanPort(hostname string, pport int) []int {
	var openPorts []int
	for port := 1; port <= pport; port++ {
		address := fmt.Sprintf("%s:%d", hostname, port)
		conn, err := net.DialTimeout("tcp", address, 1*time.Second)
		if err == nil {
			openPorts = append(openPorts, port)
			conn.Close()
		}
	}
	return openPorts
}
