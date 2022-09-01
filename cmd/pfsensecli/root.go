package main

import (
	"os"

	"github.com/sjafferali/pfsensecli/internal/version"
	"github.com/spf13/cobra"
)

var (
	pfsenseConfig = struct {
		username string
		password string
		host     string
	}{}

	jsonOutput bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "pfsensecli",
	Short:   "A CLI interface to the pfsense API",
	Version: version.FullVersionString(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(
		&pfsenseConfig.username,
		"username",
		"u",
		os.Getenv("PFSENSE_USERNAME"),
		"username to use for authentication",
	)
	rootCmd.PersistentFlags().StringVarP(
		&pfsenseConfig.password,
		"password",
		"p",
		os.Getenv("PFSENSE_PASSWORD"),
		"password to use for authentication",
	)
	rootCmd.PersistentFlags().StringVar(
		&pfsenseConfig.host,
		"host",
		os.Getenv("PFSENSE_HOST"),
		"pfsense host to connect to",
	)
	rootCmd.PersistentFlags().Bool(
		"json",
		false,
		"output json",
	)
}
