package gql

import (
	"context"
	"fmt"
)

// Hello is the resolver for the hello field.
func (r *queryResolver) Hello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
