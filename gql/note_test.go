package gql_test

import (
	"context"
	"testing"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/gql"
)

func TestNoteNodeID(t *testing.T) {
	// Arrange
	expectedNodeID := "bm90ZXM6MTIz"
	resolver := gql.NewResolver(nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Act
	nodeID, err := resolver.Resolvers.Note().NodeID(ctx, &ent.Note{
		ID: 123,
	})

	// Assert
	if err != nil {
		t.Errorf("expected err to be nil, but got: %s", err)
	}

	if nodeID != expectedNodeID {
		t.Errorf("expected NodeID to be '%s', but got '%s'", nodeID, expectedNodeID)
	}
}

func TestNoteBodyMarkdown(t *testing.T) {
	// Arrange
	obj := &ent.Note{Body: "raw markdown content [blog](https://blog.miu.guru)"}
	resolver := gql.NewResolver(nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Act
	str, err := resolver.Resolvers.Note().BodyMarkdown(ctx, obj)

	// Assert
	if err != nil {
		t.Errorf("expected no error, but got: %s", err)
	}

	if str != obj.Body {
		t.Errorf("expected markdown to be: %s, but got: %s", obj.Body, str)
	}
}

func TestNoteBodyHTML(t *testing.T) {
	// Arrange
	obj := &ent.Note{Body: "raw markdown content [blog](https://blog.miu.guru)"}
	expected := "<p>raw markdown content <a href=\"https://blog.miu.guru\">blog</a></p>\n"
	resolver := gql.NewResolver(nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Act
	str, err := resolver.Resolvers.Note().BodyHTML(ctx, obj)

	// Assert
	if err != nil {
		t.Errorf("expected no error, but got: %s", err)
	}

	if str != expected {
		t.Errorf("expected markdown to be: %s, but got: %s", expected, str)
	}
}
