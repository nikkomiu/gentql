package cmd

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"

	"github.com/nikkomiu/gentql/gql"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start the API services for gentql",
	RunE:  runAPI,
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func runAPI(cmd *cobra.Command, args []string) error {
	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)

	srv := gql.NewServer()
	router.Handle("/graphql", srv)
	router.Handle("/graphiql", playground.Handler("GentQL", "/graphql"))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return err
	}

	return nil
}
