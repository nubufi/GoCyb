package scanner

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	err := scan(80, "127.0.0.1")
	if err != nil {
		t.Error(err)
	}
}

func TestPortScanner(t *testing.T) {
	openPorts := PortScanner("127.0.0.1", 100)
	for _, port := range openPorts {
		fmt.Printf("%d open\n", port)
	}
}
