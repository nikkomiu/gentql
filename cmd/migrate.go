package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/pkg/config"
	"github.com/nikkomiu/gentql/pkg/errors"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database between versions",
	RunE:  runMigrate,
}

func init() {
	migrateCmd.Flags().BoolP("dry", "d", false, "Write the schema output to stdout instead of updating the database")

	rootCmd.AddCommand(migrateCmd)
}

func runMigrate(cmd *cobra.Command, args []string) error {
	dryRun, err := cmd.Flags().GetBool("dry")
	if err != nil {
		return err
	}

	cfg := config.GetApp()

	entClient, err := ent.Open(cfg.Database.Driver, cfg.Database.URL)
	if err != nil {
		return errors.NewExitCode(err, 3)
	}
	defer entClient.Close()

	if dryRun {
		return entClient.Schema.WriteTo(cmd.Context(), os.Stdout)
	} else {
		return entClient.Schema.Create(cmd.Context())
	}
}
