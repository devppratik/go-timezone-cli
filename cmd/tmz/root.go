/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package tmz

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tmz",
	Short: "A CLI Toolkit for Timezones",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
