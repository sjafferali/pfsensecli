package main

import (
	"context"
	"strconv"

	"github.com/sjafferali/pfsense-api-goclient/pfsenseapi"
	"github.com/spf13/cobra"
)

// listDHCPReservationsCmd represents the listDHCPReservations command
var listDHCPReservationsCmd = &cobra.Command{
	Use:   "list-dhcp-reservations",
	Short: "List DHCP reservations for the interface specified",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client := getClientForUser(pfsenseConfig.host, pfsenseConfig.username, pfsenseConfig.password)
		reservations, err := client.DHCP.ListStaticMappings(ctx, args[0])
		if err != nil {
			return err
		}
		if jsonOutput {
			if err := printJson(reservations); err != nil {
				return err
			}
			return nil
		}
		printReservationsTable(reservations)
		return nil
	},
}

// listDHCPReservationsCmd represents the listDHCPReservations command
var createDHCPReservationsCmd = &cobra.Command{
	Use:   "create-dhcp-reservations",
	Short: "List DHCP reservations for the interface specified",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client := getClientForUser(pfsenseConfig.host, pfsenseConfig.username, pfsenseConfig.password)
		reservations, err := client.DHCP.ListStaticMappings(ctx, args[0])
		if err != nil {
			return err
		}
		if jsonOutput {
			if err := printJson(reservations); err != nil {
				return err
			}
			return nil
		}
		printReservationsTable(reservations)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listDHCPReservationsCmd)
}

func printReservationsTable(reservations []pfsenseapi.DHCPStaticMapping) {
	data := make([][]string, 0)
	for _, i := range reservations {
		data = append(
			data,
			[]string{
				strconv.Itoa(i.ID),
				i.Mac,
				i.IPaddr,
				i.Hostname,
				i.Descr,
			},
		)
	}
	header := []string{"ID", "Mac", "IPAddr", "Hostname", "Descr"}
	printTable(header, data)
}
