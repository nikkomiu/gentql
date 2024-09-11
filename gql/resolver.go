package gql

import "github.com/99designs/gqlgen/graphql/handler"

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct{}

func NewResolver() Config {
	return Config{
		Resolvers: &Resolver{},
	}
}

func NewServer() *handler.Server {
	return handler.NewDefaultServer(NewExecutableSchema(NewResolver()))
}
