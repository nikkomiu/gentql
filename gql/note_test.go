package gql_test

import (
	"fmt"
	"testing"

	"entgo.io/contrib/entgql"

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

func TestNoteList(t *testing.T) {
	ctx := ContextT(t)
	entClient := EntT(t)
	resolver := gql.NewResolver(entClient)

	totalCount := 10
	for i := 0; i < totalCount; i++ {
		_, err := entClient.Note.Create().
			SetTitle(fmt.Sprintf("Test Note %d", i)).
			SetBody("Test Note Body").
			Save(ctx)
		assert.NoError(t, err)
	}

	three := 3

	tt := []struct {
		name string

		after   *entgql.Cursor[int]
		first   *int
		before  *entgql.Cursor[int]
		last    *int
		orderBy *ent.NoteOrder
		where   *ent.NoteWhereInput

		expectedErr bool
		expectedLen int
	}{
		{
			name: "default",

			expectedLen: totalCount,
		},
		{
			name: "first 3",

			first: &three,

			expectedLen: 3,
		},
		{
			name: "last 3",

			last: &three,

			expectedLen: 3,
		},
		{
			name: "order by title asc",

			orderBy: &ent.NoteOrder{Field: ent.NoteOrderFieldTitle, Direction: entgql.OrderDirectionAsc},

			expectedLen: totalCount,
		},
		{
			name: "order by created at desc",

			orderBy: &ent.NoteOrder{Field: ent.NoteOrderFieldCreatedAt, Direction: entgql.OrderDirectionDesc},

			expectedLen: totalCount,
		},
		{
			name: "order by updated at asc",

			orderBy: &ent.NoteOrder{Field: ent.NoteOrderFieldUpdatedAt, Direction: entgql.OrderDirectionAsc},

			expectedLen: totalCount,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ctx := ContextT(t)

			// Act
			notes, err := resolver.Resolvers.Query().Notes(ctx, tc.after, tc.first, tc.before, tc.last, tc.orderBy, tc.where)

			// Assert
			assert.NoError(t, err)
			assert.Len(t, notes.Edges, tc.expectedLen)
			assert.Equal(t, totalCount, notes.TotalCount)
		})
	}
}
