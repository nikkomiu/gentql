package cmd

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gentql",
	Short: "GentQL backend application services.",

	SilenceUsage: true,
}

func Execute(ctx context.Context) error {
	return rootCmd.ExecuteContext(ctx)
}
