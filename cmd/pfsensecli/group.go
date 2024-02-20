package main

import (
	"context"
	"strconv"

	"github.com/sjafferali/pfsense-api-goclient/pfsenseapi"
	"github.com/spf13/cobra"
)

var (
	groupCmd = &cobra.Command{
		Use:   "group",
		Short: "Commands associated with groups",
	}

	listGroupsCmd = &cobra.Command{
		Use:   "list",
		Short: "List Groups",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client := getClientForUser(pfsenseConfig.host, pfsenseConfig.username, pfsenseConfig.password)
			groups, err := client.User.ListGroups(ctx)
			if err != nil {
				return err
			}

			if jsonOutput {
				if err := printJson(groups); err != nil {
					return err
				}
				return nil
			}

			printGroupsTable(groups)
			return nil
		},
	}
)

func init() {
	groupCmd.AddCommand(listGroupsCmd)
	rootCmd.AddCommand(groupCmd)
}

func printGroupsTable(groups []*pfsenseapi.Group) {
	data := make([][]string, 0)
	for _, g := range groups {
		data = append(
			data,
			[]string{
				g.Name,
				strconv.Itoa(g.Gid),
				g.Description,
			},
		)
	}
	header := []string{"Name", "GID", "Description"}
	printTable(header, data)
}
