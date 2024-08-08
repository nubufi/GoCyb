package portscanner

import (
	"fmt"

	"github.com/spf13/cobra"
)

func PortScannerCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "portscan [hostname]",
		Short: "Scans open ports on a given hostname",
		Long:  "Scans open ports on a given hostname",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			hostname := args[0]
			openPorts := ScanPort(hostname, 100)
			if len(openPorts) == 0 {
				fmt.Printf("No open ports found on %s\n", hostname)
			} else {
				fmt.Printf("Open ports on %s: %v\n", hostname, openPorts)
			}
		},
	}
}
