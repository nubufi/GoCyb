package main

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
	Response string
}

var services = map[int]ServiceVersion{

	80:   {Port: 80, Protocol: "HTTP", Response: "HTTP Response Headers"},
	443:  {Port: 443, Protocol: "HTTP", Response: "HTTP Response Headers"},
	21:   {Port: 21, Protocol: "FTP", Response: "220 Response Code"},
	22:   {Port: 22, Protocol: "SSH", Response: "SSH Protocol Response"},
	25:   {Port: 25, Protocol: "SMTP", Response: "220 Response Code"},
	110:  {Port: 110, Protocol: "POP3", Response: "+OK Response Code"},
	143:  {Port: 143, Protocol: "IMAP", Response: "* OK Response Code"},
	53:   {Port: 53, Protocol: "DNS", Response: "DNS Query Responses"},
	3306: {Port: 3306, Protocol: "MySQL", Response: "Protocol Headers"},
	5432: {Port: 5432, Protocol: "PostgreSQL", Response: "PostgreSQL Responses"},
	1433: {Port: 1433, Protocol: "MSSQL", Response: "SQL Server Responses"},
	161:  {Port: 161, Protocol: "SNMP", Response: "SNMP Query Responses"},
	23:   {Port: 23, Protocol: "Telnet", Response: "Telnet Responses"},
	3389: {Port: 3389, Protocol: "RDP", Response: "RDP Responses"},
	387:  {Port: 387, Protocol: "SCTP", Response: "INIT and COOKIE ECHO Chunks"},
	2049: {Port: 2049, Protocol: "NFS", Response: "NFS Responses"},
	389:  {Port: 389, Protocol: "LDAP", Response: "LDAP Responses"},
	137:  {Port: 137, Protocol: "NetBIOS", Response: "NetBIOS Responses"},
	138:  {Port: 138, Protocol: "NetBIOS", Response: "NetBIOS Responses"},
	139:  {Port: 139, Protocol: "NetBIOS", Response: "NetBIOS Responses"},
	873:  {Port: 873, Protocol: "Rsync", Response: "Rsync Responses"},
	1812: {Port: 1812, Protocol: "RADIUS", Response: "RADIUS Responses"},
	1813: {Port: 1813, Protocol: "RADIUS", Response: "RADIUS Responses"},
	69:   {Port: 69, Protocol: "TFTP", Response: "TFTP Responses"},
	445:  {Port: 445, Protocol: "SMB", Response: "SMB Responses"},
	5060: {Port: 5060, Protocol: "SIP", Response: "SIP Responses"},
	6660: {Port: 6660, Protocol: "IRC", Response: "IRC Responses"},
	6661: {Port: 6661, Protocol: "IRC", Response: "IRC Responses"},
	6662: {Port: 6662, Protocol: "IRC", Response: "IRC Responses"},
	6663: {Port: 6663, Protocol: "IRC", Response: "IRC Responses"},
	6664: {Port: 6664, Protocol: "IRC", Response: "IRC Responses"},
	6665: {Port: 6665, Protocol: "IRC", Response: "IRC Responses"},
	6666: {Port: 6666, Protocol: "IRC", Response: "IRC Responses"},
	6667: {Port: 6667, Protocol: "IRC", Response: "IRC Responses"},
	6668: {Port: 6668, Protocol: "IRC", Response: "IRC Responses"},
	6669: {Port: 6669, Protocol: "IRC", Response: "IRC Responses"},
	1883: {Port: 1883, Protocol: "MQTT", Response: "MQTT Responses"},
	5222: {Port: 5222, Protocol: "XMPP", Response: "XMPP Responses"},
	9042: {Port: 9042, Protocol: "Cassandra", Response: "Cassandra Responses"},
	9200: {Port: 9200, Protocol: "Elasticsearch", Response: "Elasticsearch Responses"},
}

func scanPort(ip string, port int) string {
	address := fmt.Sprintf("%s:%d", ip, port)

	timeout := 5 * time.Second
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		if _, ok := err.(*net.OpError); ok {
			return Filtered
		}
		return Filtered
	}
	conn.Close()
	return Open
}

func detectService(port int) ServiceVersion {
	if svc, ok := services[port]; ok {
		return svc
	}
	return ServiceVersion{Port: port, Protocol: "Unknown", Response: "Unknown"}
}

func worker(ip string, ports, results chan int, openPorts chan int, done chan bool) {
	for port := range ports {
		state := scanPort(ip, port)
		service := detectService(port)
		results <- port
		if state == Open {
			openPorts <- port
		}
		fmt.Printf("Port %d: %s, Service: %s, Response: %s\n", port, state, service.Protocol, service.Response)
	}
	done <- true
}

func main() {
	var domain string
	fmt.Print("Enter domain: ")
	fmt.Scanln(&domain)

	ips, err := net.LookupHost(domain)
	if err != nil {
		fmt.Printf("Error resolving domain: %s\n", err)
		return
	}

	var ports []int
	for port := 1; port <= 65535; port++ {
		ports = append(ports, port)
	}

	portChannel := make(chan int, len(ports))
	resultChannel := make(chan int, len(ports))
	openPorts := make(chan int, len(ports))
	done := make(chan bool, 100)

	numWorkers := 100
	for i := 0; i < numWorkers; i++ {
		go worker("", portChannel, resultChannel, openPorts, done)
	}

	for _, ip := range ips {
		fmt.Printf("Scanning IP: %s\n", ip)

		for _, port := range ports {
			portChannel <- port
		}
		close(portChannel)

		for range ports {
			<-resultChannel
		}

		for i := 0; i < numWorkers; i++ {
			<-done
		}
		close(openPorts)

		fmt.Println("Open Ports:")
		for port := range openPorts {
			fmt.Printf("Port %d is Open\n", port)
		}
	}

	close(resultChannel)
}
