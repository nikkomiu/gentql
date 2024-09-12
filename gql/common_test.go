package gql_test

import (
	"context"
	"testing"

	"github.com/nikkomiu/gentql/gql"
)

func TestPing(t *testing.T) {
	// Arrange
	expected := "pong"
	resolver := gql.NewResolver(nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Act
	res, err := resolver.Resolvers.Query().Ping(ctx)

	// Assert
	if err != nil {
		t.Errorf("expected error to be nil but got: %s", err)
	}

	if res != expected {
		t.Errorf("expected %s but got %s", expected, res)
	}
}
