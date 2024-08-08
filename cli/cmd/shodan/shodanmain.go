package shodan

import (
	"fmt"
	"greeter/cmd/shodanapi"

	"github.com/spf13/cobra"
)

// Shodan API anahtarı için bir global değişken
var apiKey string

func ShodanScanCommand() *cobra.Command {
	// shodanCmd represents the shodan command
	shodanCmd := &cobra.Command{
		Use:   "shodan",
		Short: "Shodan search",
		Long: `***Allows you to search based on a specific query using Shodan API.***
	Example:
	  shodan --apikey YOUR_API_KEY "apache"
	  shodan --apikey YOUR_API_KEY "port:22"
	  shodan --apikey YOUR_API_KEY "country:GR"        2 letter country code de,us
	  shodan --apikey YOUR_API_KEY "org:Google"
	  shodan --apikey YOUR_API_KEY "os:Windows 10"`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				fmt.Println("Enter search request.")
				return
			}
			query := args[0]

			client := shodanapi.New(apiKey)
			result, err := client.HostSearch(query)
			if err != nil {
				fmt.Printf("Unexpected error while searching: %v\n", err)
				return
			}

			for _, host := range result.Matches {
				fmt.Printf("IP: %s, Port: %d, ISP: %s, Country: %s\n",
					host.IPString, host.Port, host.ISP, host.Location.CountryName)
			}
		},
	}

	// API anahtarını CLI üzerinden almak için flag ekliyoruz
	shodanCmd.Flags().StringVarP(&apiKey, "apikey", "a", "", "Shodan API anahtarı")
	shodanCmd.MarkFlagRequired("apikey")
	return shodanCmd
}
