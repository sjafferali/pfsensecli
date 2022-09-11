package main

import (
	"context"

	"github.com/sjafferali/pfsense-api-goclient/pfsenseapi"
	"github.com/spf13/cobra"
)

// listInterfacesCmd represents the list-interfaces command
var listInterfacesCmd = &cobra.Command{
	Use:   "list-interfaces",
	Short: "List Interfaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		client := getClientForUser(pfsenseConfig.host, pfsenseConfig.username, pfsenseConfig.password)
		interfaces, err := client.Interface.ListInterfaces(ctx)
		if err != nil {
			return err
		}

		if jsonOutput {
			if err := printJson(interfaces); err != nil {
				return err
			}
			return nil
		}

		printInterfacesTable(interfaces)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listInterfacesCmd)
}

func printInterfacesTable(interfaces []*pfsenseapi.Interface) {
	data := make([][]string, 0)
	for _, i := range interfaces {
		data = append(
			data,
			[]string{
				i.Name,
				i.If,
				i.Descr,
				i.Ipaddr,
			},
		)
	}
	header := []string{"Name", "Interface", "Desc", "IP"}
	printTable(header, data)
}
