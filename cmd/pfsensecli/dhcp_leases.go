package main

import (
	"context"

	"github.com/sjafferali/pfsense-api-goclient/pfsenseapi"
	"github.com/spf13/cobra"
)

// listDhcpLeasesCmd represents the listDhcpLeases command
var (
	dhcpLeasesCmd = &cobra.Command{
		Use:   "dhcp-leases",
		Short: "Commands associated with DHCP leases",
	}

	listDhcpLeasesCmd = &cobra.Command{
		Use:   "list",
		Short: "List DHCP Leases",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client := getClientForUser(pfsenseConfig.host, pfsenseConfig.username, pfsenseConfig.password)
			leases, err := client.DHCP.ListLeases(ctx)
			if err != nil {
				return err
			}
			withExpired, _ := cmd.Flags().GetBool("withexpired")
			if !withExpired {
				leases = sliceFilter(leases, func(lease *pfsenseapi.DHCPLease) bool {
					return lease.State != "expired"
				})
			}

			if jsonOutput {
				if err := printJson(leases); err != nil {
					return err
				}
				return nil
			}

			printLeasesTable(leases)
			return nil
		},
	}
)

func init() {
	listDhcpLeasesCmd.Flags().Bool("withexpired", false, "show expired leases")
	dhcpLeasesCmd.AddCommand(listDhcpLeasesCmd)
	rootCmd.AddCommand(dhcpLeasesCmd)
}

func printLeasesTable(leases []*pfsenseapi.DHCPLease) {
	data := make([][]string, 0)
	for _, i := range leases {
		data = append(
			data,
			[]string{
				i.Ip,
				i.Mac,
				i.If,
				i.Type,
				i.State,
				i.Online,
				i.Hostname,
				i.Descr,
			},
		)
	}
	header := []string{"IP", "Mac", "Iface", "Type", "State", "Online", "Hostname", "Desc"}
	printTable(header, data)
}
