package dns_subdomain

import (
	"fmt"
	"greeter/cmd/dns"

	"github.com/spf13/cobra"
)

func SubScannerCommand() *cobra.Command {
	var domain string
	var server string = "8.8.8.8:53"
	var workers int

	subCmd := &cobra.Command{
		Use:   "subscan -d example.com -s 8.8.8.8:53 -w 200",
		Short: "Scans sub domains on a given hostname",
		Long:  "Scans subdomains on a given hostname with specified DNS server, number of workers, and a comma-separated list of subdomains",
		Run: func(cmd *cobra.Command, args []string) {
			wordList := []string{
				"www", "mail", "ftp", "docs", "localhost", "webmail", "smtp", "pop", "ns1",
				"webdisk", "ns2", "cpanel", "whm", "autodiscover", "autoconfig", "m", "imap",
				"test", "ns", "blog", "pop3", "dev", "www2", "admin", "forum", "news", "vpn",
				"ns3", "mail2", "new", "mysql", "old", "lists", "support", "mobile", "mx",
				"static", "docs", "beta", "shop", "sql", "secure", "demo", "cp", "calendar",
				"wiki", "web", "media", "email", "images", "img", "www1", "intranet", "portal",
				"video", "sip", "dns2", "api", "cdn", "stats", "dns1", "ns4", "www3", "dns",
				"search", "staging", "server", "mx1", "chat", "wap", "my", "svn", "mail1", "sites",
				"proxy", "ads", "host", "crm", "cms", "backup", "mx2", "lyncdiscover", "info",
				"apps", "download", "remote", "db", "forums", "store", "relay", "files",
				"newsletter", "app", "live", "owa", "en", "start", "sms", "office", "exchange",
				"ipv4", "Footer",
			}

			result, err := dns.SubDomainScanner(domain, server, workers, wordList)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Results:")
				for _, res := range result {
					fmt.Printf("Subdomain: %s,\t IP: %s\n", res.Hostname, res.Ip)
				}
			}
		},
	}

	subCmd.PersistentFlags().StringVarP(&domain, "domain", "d", "", "Domain to scan for subdomains")
	subCmd.PersistentFlags().StringVarP(&server, "server", "s", "8.8.8.8:53", "DNS server to use")
	subCmd.PersistentFlags().IntVarP(&workers, "workers", "w", 20, "Number of workers for concurrent processing")

	subCmd.MarkFlagRequired("domain")

	return subCmd
}
