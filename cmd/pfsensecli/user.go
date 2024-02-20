package main

import (
	"context"

	"github.com/sjafferali/pfsense-api-goclient/pfsenseapi"
	"github.com/spf13/cobra"
)

var (
	userCmd = &cobra.Command{
		Use:   "user",
		Short: "Commands associated with users",
	}

	listUsersCmd = &cobra.Command{
		Use:   "list",
		Short: "List Users",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			client := getClientForUser(pfsenseConfig.host, pfsenseConfig.username, pfsenseConfig.password)
			users, err := client.User.ListUsers(ctx)
			if err != nil {
				return err
			}

			if jsonOutput {
				if err := printJson(users); err != nil {
					return err
				}
				return nil
			}

			printUsersTable(users)
			return nil
		},
	}
)

func init() {
	userCmd.AddCommand(listUsersCmd)
	rootCmd.AddCommand(userCmd)
}

func printUsersTable(users []*pfsenseapi.User) {
	data := make([][]string, 0)
	for _, u := range users {
		data = append(
			data,
			[]string{
				u.Name,
				u.Uid,
				u.Descr,
			},
		)
	}
	header := []string{"Name", "UID", "Description"}
	printTable(header, data)
}
