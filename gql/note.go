package gql

import (
	"context"
	"fmt"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/gql/model"
)

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input model.NoteInput) (*ent.Note, error) {
	panic(fmt.Errorf("not implemented: CreateNote - createNote"))
}

// UpdateNote is the resolver for the updateNote field.
func (r *mutationResolver) UpdateNote(ctx context.Context, id int, input model.NoteInput) (*ent.Note, error) {
	panic(fmt.Errorf("not implemented: UpdateNote - updateNote"))
}

// DeleteNote is the resolver for the deleteNote field.
func (r *mutationResolver) DeleteNote(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteNote - deleteNote"))
}

// NodeID is the resolver for the nodeId field.
func (r *noteResolver) NodeID(ctx context.Context, obj *ent.Note) (string, error) {
	panic(fmt.Errorf("not implemented: NodeID - nodeId"))
}

// BodyMarkdown is the resolver for the bodyMarkdown field.
func (r *noteResolver) BodyMarkdown(ctx context.Context, obj *ent.Note) (string, error) {
	panic(fmt.Errorf("not implemented: BodyMarkdown - bodyMarkdown"))
}

// BodyHTML is the resolver for the bodyHtml field.
func (r *noteResolver) BodyHTML(ctx context.Context, obj *ent.Note) (string, error) {
	panic(fmt.Errorf("not implemented: BodyHTML - bodyHtml"))
}

// Notes is the resolver for the notes field.
func (r *queryResolver) Notes(ctx context.Context) ([]*ent.Note, error) {
	return r.ent.Note.Query().All(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Note returns NoteResolver implementation.
func (r *Resolver) Note() NoteResolver { return &noteResolver{r} }

type mutationResolver struct{ *Resolver }
type noteResolver struct{ *Resolver }
