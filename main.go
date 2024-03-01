package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nubufi/GoCyb/metasploit"
)

func main() {
	host := os.Getenv("MSFHOST")
	pass := os.Getenv("MSFPASS")
	user := "msf"

	msf, err := metasploit.New(host, user, pass)
	if err != nil {
		log.Panicln(err)
	}

	defer msf.Logout()

	sessions, err := msf.SessionList()
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("Sessions:")
	for _, session := range sessions {
		fmt.Printf("%5d %s\n", session.ID, session.Info)
	}
}
