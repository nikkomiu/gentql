package gql_test

import (
	"context"
	"encoding/base64"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/ent/enttest"
	"github.com/nikkomiu/gentql/gql"
)

func ContextT(t *testing.T) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	return ctx
}

func EntT(t *testing.T) *ent.Client {
	entClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	t.Cleanup(func() { entClient.Close() })
	return entClient
}

func TestPing(t *testing.T) {
	// Arrange
	expected := "pong"
	resolver := gql.NewResolver(nil)
	ctx := ContextT(t)

	// Act
	res, err := resolver.Resolvers.Query().Ping(ctx)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
	if err != nil {
		t.Errorf("expected error to be nil but got: %s", err)
	}

	if res != expected {
		t.Errorf("expected %s but got %s", expected, res)
	}
}

func TestNode(t *testing.T) {
	ctx := ContextT(t)
	entClient := EntT(t)
	resolver := gql.NewResolver(entClient)

	note, err := entClient.Note.Create().SetTitle("Test Note 1").SetBody("Test Note Body 1").Save(ctx)
	assert.NoError(t, err)
	if err != nil {
		t.Errorf("expected note to be created, but got err: %s", err)
	}

	noteNodeID, err := resolver.Resolvers.Note().NodeID(ctx, note)
	if err != nil {
		t.Errorf("expected note to resolve node id, but got err: %s", err)
	}

	notFoundNoteNodeID, err := resolver.Resolvers.Note().NodeID(ctx, &ent.Note{ID: 0})
	if err != nil {
		t.Errorf("expected note to resolve node id, but got err: %s", err)
	}

	tt := []struct {
		name   string
		nodeID string

		expectedNode bool
		expectedErr  bool
	}{
		{
			name:   "note",
			nodeID: noteNodeID,

			expectedNode: true,
		},

		{
			name:   "not found",
			nodeID: notFoundNoteNodeID,

			expectedErr: true,
		},
		{
			name:   "invalid base64",
			nodeID: "bad string",

			expectedErr: true,
		},
		{
			name:   "not enough parts",
			nodeID: base64.RawURLEncoding.EncodeToString([]byte("notes")),

			expectedErr: true,
		},
		{
			name:   "too many parts",
			nodeID: base64.RawURLEncoding.EncodeToString([]byte("notes:1:3")),

			expectedErr: true,
		},
		{
			name:   "bad id value",
			nodeID: base64.RawURLEncoding.EncodeToString([]byte("notes:numless")),

			expectedErr: true,
		},
		{
			name:   "bad table name",
			nodeID: base64.RawURLEncoding.EncodeToString([]byte("not_my_table:11")),

			expectedErr: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := ContextT(t)

			node, err := resolver.Resolvers.Query().Node(ctx, tc.nodeID)

			assert.Equal(t, tc.expectedErr, err != nil, "expected error to be %v but got: %s", tc.expectedErr, err)
			if tc.expectedNode {
				assert.NotNil(t, node, "expected node to be not nil")
			} else {
				assert.Nil(t, node, "expected node to be nil")
			}
		})
	}
}
