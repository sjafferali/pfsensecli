package main

import (
	"context"

	"github.com/sjafferali/pfsense-api-goclient/pfsenseapi"
	"github.com/spf13/cobra"
)

var (
	interfaceCmd = &cobra.Command{
		Use:   "interface",
		Short: "Commands associated with interfaces",
	}

	listInterfacesCmd = &cobra.Command{
		Use:   "list",
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
)

func init() {
	interfaceCmd.AddCommand(listInterfacesCmd)
	rootCmd.AddCommand(interfaceCmd)
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
