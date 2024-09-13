package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/pkg/config"
	"github.com/nikkomiu/gentql/pkg/errors"
)

func newMigrateCmd() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Migrate the database between versions",
		RunE:  runMigrate,
	}

	migrateCmd.Flags().BoolP("dry", "d", false, "Write the schema output to stdout instead of updating the database")

	return migrateCmd
}

func runMigrate(cmd *cobra.Command, args []string) error {
	dryRun, err := cmd.Flags().GetBool("dry")
	if err != nil {
		return err
	}

	ctx, cfg := config.WithApp(cmd.Context())

	entClient, err := ent.Open(cfg.Database.Driver, cfg.Database.URL)
	if err != nil {
		return errors.NewExitCode(err, 3)
	}
	defer entClient.Close()

	if dryRun {
		return entClient.Schema.WriteTo(ctx, os.Stdout)
	} else {
		return entClient.Schema.Create(ctx)
	}
}
