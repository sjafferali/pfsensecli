package main

import (
	"context"

	"github.com/sjafferali/pfsense-api-goclient/pfsenseapi"
	"github.com/spf13/cobra"
)

var (
	gatewayCmd = &cobra.Command{
		Use:   "gateway",
		Short: "Commands associated with gateways",
	}

	listGatewaysCmd = &cobra.Command{
		Use:   "list",
		Short: "List Gateways",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client := getClientForUser(pfsenseConfig.host, pfsenseConfig.username, pfsenseConfig.password)
			gateways, err := client.Routing.ListGateways(ctx)
			if err != nil {
				return err
			}

			if jsonOutput {
				if err := printJson(gateways); err != nil {
					return err
				}
				return nil
			}

			printGatewaysTable(gateways)
			return nil
		},
	}
)

func init() {
	gatewayCmd.AddCommand(listGatewaysCmd)
	rootCmd.AddCommand(gatewayCmd)
}

func printGatewaysTable(gateways []*pfsenseapi.Gateway) {
	data := make([][]string, 0)
	for _, g := range gateways {
		data = append(
			data,
			[]string{
				g.Name,
				g.Interface,
				g.Gateway,
				g.Monitor,
				g.Descr,
			},
		)
	}
	header := []string{"Name", "Interface", "Gateway", "Monitor", "Description"}
	printTable(header, data)
}
