package cmd

import (
	"context"

	_ "github.com/lib/pq"

	"github.com/spf13/cobra"
)

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "gentql",
		Short: "GentQL backend application services.",

		SilenceUsage: true,
	}

	rootCmd.AddCommand(
		newAPICmd(),
		newMigrateCmd(),
		newSeedCmd(),
	)

	return rootCmd
}

func Execute(ctx context.Context, opts ...Option) error {
	rootCmd := newRootCmd()

	for _, opt := range opts {
		opt.CmdOpt(rootCmd)
	}

	return rootCmd.ExecuteContext(ctx)
}
