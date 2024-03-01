package forwarder

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handler(src net.Conn, domain string, port int) {
	address := fmt.Sprintf("%s:%d", domain, port)
	dst, err := net.Dial("tcp", address)

	if err != nil {
		log.Fatalln(err)
	}

	defer dst.Close()

	go func() {
		if _, err := io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}
	}()

	if _, err := io.Copy(src, dst); err != nil {
		log.Fatalln(err)
	}
}

func PortForwarder(domain string, targetPort, hostPort int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", hostPort))
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handler(conn, domain, targetPort)
	}
}
