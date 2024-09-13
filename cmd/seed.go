package cmd

import (
	"fmt"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/pkg/config"
	"github.com/spf13/cobra"
)

func newSeedCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "seed",
		Short: "Seed the database with initial values",
		RunE:  runSeed,
	}
}

func runSeed(cmd *cobra.Command, args []string) (err error) {
	ctx, cfg := config.WithApp(cmd.Context())

	entClient, err := ent.Open(cfg.Database.Driver, cfg.Database.URL)
	if err != nil {
		return
	}
	defer entClient.Close()

	notes := []*ent.NoteCreate{
		entClient.Note.Create().
			SetTitle("My First Note").
			SetBody("## My First Note Section\n\nSome content for the note. With a [link](https://blog.miu.guru) to a cool site!"),
		entClient.Note.Create().
			SetTitle("My Second Note").
			SetBody("## My Other Note\n\nMore random note content...\n\n- with\n- a\n- list\n\nAll this formatting and no where to go."),
	}

	fmt.Println("Seeding notes...")
	err = entClient.Note.CreateBulk(notes...).Exec(ctx)

	// create additional seeds here

	return
}
