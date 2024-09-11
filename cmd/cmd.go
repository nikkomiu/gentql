package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "gentql",

	SilenceUsage: true,
}

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}
