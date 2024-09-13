package cmd

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/gql"
	"github.com/nikkomiu/gentql/pkg/config"
	"github.com/nikkomiu/gentql/pkg/errors"
	"github.com/nikkomiu/gentql/pkg/sig"
)

func newAPICmd() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "Start the API services for gentql",
		RunE:  runAPI,
	}
}

func runAPI(cmd *cobra.Command, args []string) error {
	ctx, cfg := config.WithApp(cmd.Context())

	entClient, err := ent.Open(cfg.Database.Driver, cfg.Database.URL)
	if err != nil {
		return errors.NewExitCode(err, 3)
	}
	ctx = ent.NewContext(ctx, entClient)
	defer entClient.Close()

	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)

	srv := gql.NewServer(ctx)
	router.Handle("/graphql", srv)
	router.Handle("/graphiql", playground.Handler("GentQL", "/graphql"))

	fmt.Printf("starting server at %s\n", cfg.Server.DisplayAddr())
	return sig.ListenAndServe(ctx, cfg.Server.Addr(), router, cfg.Server.ShutdownTimeout)
}
