package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:  "api",
	RunE: runAPI,

	SilenceUsage: true,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func runAPI(cmd *cobra.Command, args []string) error {
	fmt.Println("hello api")
	return nil
}
