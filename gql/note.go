package gql

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/ent/note"
	"github.com/nikkomiu/gentql/gql/model"
	"github.com/yuin/goldmark"
)

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input model.NoteInput) (*ent.Note, error) {
	return r.ent.Note.Create().
		SetTitle(input.Title).
		SetBody(input.Body).
		Save(ctx)
}

// UpdateNote is the resolver for the updateNote field.
func (r *mutationResolver) UpdateNote(ctx context.Context, id int, input model.NoteInput) (*ent.Note, error) {
	return r.ent.Note.UpdateOneID(id).
		SetTitle(input.Title).
		SetBody(input.Body).
		Save(ctx)
}

// DeleteNote is the resolver for the deleteNote field.
func (r *mutationResolver) DeleteNote(ctx context.Context, id int) (bool, error) {
	err := r.ent.Note.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// NodeID is the resolver for the nodeId field.
func (r *noteResolver) NodeID(ctx context.Context, obj *ent.Note) (string, error) {
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%d", note.Table, obj.ID))), nil
}

// BodyMarkdown is the resolver for the bodyMarkdown field.
func (r *noteResolver) BodyMarkdown(ctx context.Context, obj *ent.Note) (string, error) {
	return obj.Body, nil
}

// BodyHTML is the resolver for the bodyHtml field.
func (r *noteResolver) BodyHTML(ctx context.Context, obj *ent.Note) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(obj.Body), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
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
