package main

import (
	"context"
	"strconv"

	"github.com/sjafferali/pfsense-api-goclient/pfsenseapi"
	"github.com/spf13/cobra"
)

// listDhcpLeasesCmd represents the listDhcpLeases command
var listDhcpLeasesCmd = &cobra.Command{
	Use:   "list-dhcp-leases",
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
			leases = sliceFilter(leases, func(lease pfsenseapi.DHCPLease) bool {
				if lease.State == "expired" {
					return false
				}
				return true
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

func init() {
	rootCmd.AddCommand(listDhcpLeasesCmd)
	listDhcpLeasesCmd.Flags().Bool("withexpired", false, "show expired leases")
}

func printLeasesTable(leases []pfsenseapi.DHCPLease) {
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
				strconv.FormatBool(i.Online),
				i.Hostname,
				i.Descr,
			},
		)
	}
	header := []string{"IP", "Mac", "Iface", "Type", "State", "Online", "Hostname", "Desc"}
	printTable(header, data)
}
