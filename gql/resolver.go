package gql

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/nikkomiu/gentql/ent"
)

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	ent *ent.Client
}

func NewResolver(entClient *ent.Client) Config {
	return Config{
		Resolvers: &Resolver{
			ent: entClient,
		},
	}
}

func NewServer(ctx context.Context) *handler.Server {
	resolver := NewResolver(ent.FromContext(ctx))
	return handler.NewDefaultServer(NewExecutableSchema(resolver))
}
