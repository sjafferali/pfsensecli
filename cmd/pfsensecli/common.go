package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/sjafferali/pfsense-api-goclient/pfsenseapi"
)

func getClientForUser(host, user, password string) *pfsenseapi.Client {
	config := pfsenseapi.Config{
		Host:             host,
		LocalAuthEnabled: true,
		User:             user,
		Password:         password,
		SkipTLS:          true,
		Timeout:          10 * time.Second,
	}
	return pfsenseapi.NewClient(config)
}

func printJson(data interface{}) error {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	fmt.Println(string(val))
	return nil
}

func printTable(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: false, Top: false, Right: false, Bottom: false})
	table.SetCenterSeparator("")
	table.SetHeaderLine(false)
	table.SetColumnSeparator("")
	table.SetAutoWrapText(false)
	table.AppendBulk(data)
	table.Render()
}

func sliceFilter[T1 any](in []T1, fn func(T1) bool) []T1 {
	out := make([]T1, 0, len(in))

	for _, v := range in {
		if fn(v) {
			out = append(out, v)
		}
	}

	return out
}
