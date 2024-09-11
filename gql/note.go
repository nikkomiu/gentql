package gql

import (
	"context"
	"fmt"

	"github.com/nikkomiu/gentql/gql/model"
)

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input model.NoteInput) (*model.Note, error) {
	panic(fmt.Errorf("not implemented: CreateNote - createNote"))
}

// UpdateNote is the resolver for the updateNote field.
func (r *mutationResolver) UpdateNote(ctx context.Context, id int, input model.NoteInput) (*model.Note, error) {
	panic(fmt.Errorf("not implemented: UpdateNote - updateNote"))
}

// DeleteNote is the resolver for the deleteNote field.
func (r *mutationResolver) DeleteNote(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteNote - deleteNote"))
}

// Notes is the resolver for the notes field.
func (r *queryResolver) Notes(ctx context.Context) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented: Notes - notes"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
