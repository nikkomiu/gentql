package gql_test

import (
	"testing"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/gql"
	"github.com/stretchr/testify/assert"
)

func TestNoteNodeID(t *testing.T) {
	// Arrange
	expectedNodeID := "bm90ZXM6MTIz"
	resolver := gql.NewResolver(nil)
	ctx := ContextT(t)

	// Act
	nodeID, err := resolver.Resolvers.Note().NodeID(ctx, &ent.Note{
		ID: 123,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedNodeID, nodeID)
}

func TestNoteBodyMarkdown(t *testing.T) {
	// Arrange
	obj := &ent.Note{Body: "raw markdown content [blog](https://blog.miu.guru)"}
	resolver := gql.NewResolver(nil)
	ctx := ContextT(t)

	// Act
	str, err := resolver.Resolvers.Note().BodyMarkdown(ctx, obj)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, obj.Body, str)
}

func TestNoteBodyHTML(t *testing.T) {
	// Arrange
	obj := &ent.Note{Body: "raw markdown content [blog](https://blog.miu.guru)"}
	expected := "<p>raw markdown content <a href=\"https://blog.miu.guru\">blog</a></p>\n"
	resolver := gql.NewResolver(nil)
	ctx := ContextT(t)

	// Act
	str, err := resolver.Resolvers.Note().BodyHTML(ctx, obj)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, str)
}
