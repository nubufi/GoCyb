package root

import (
	"fmt"
	"greeter/cmd/dns_subdomain"
	"greeter/cmd/greet"
	"greeter/cmd/portscanner"
	"greeter/cmd/shodan"

	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "greeter",
		Short: "Greeter basic CLI",
		Long:  "Greeter is a friendly command line application",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Greeter!")
		},
	}

	cmd.AddCommand(greet.GreetCommand())
	cmd.AddCommand(portscanner.PortScannerCommand())
	cmd.AddCommand(dns_subdomain.SubScannerCommand())
	cmd.AddCommand(shodan.ShodanScanCommand())

	return cmd
}
