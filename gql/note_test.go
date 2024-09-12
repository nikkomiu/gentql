package gql_test

import (
	"fmt"
	"testing"
	"time"

	"entgo.io/contrib/entgql"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/gql"
	"github.com/nikkomiu/gentql/gql/model"
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

func TestNoteCreate(t *testing.T) {
	entClient := EntT(t)
	resolver := gql.NewResolver(entClient)

	tt := []struct {
		name string

		input model.NoteInput

		expectedErr  bool
		expectedNote *ent.Note
	}{
		{
			name: "default",

			input: model.NoteInput{
				Title: "Test Note",
				Body:  "Test Note Body",
			},

			expectedNote: &ent.Note{
				Title: "Test Note",
				Body:  "Test Note Body",
			},
		},
		{
			name: "empty title",

			input: model.NoteInput{
				Body: "Test Note Body",
			},

			expectedErr: true,
		},
		{
			name: "title too short",

			input: model.NoteInput{
				Title: "T",
				Body:  "Test Note Body",
			},

			expectedErr: true,
		},
		{
			name: "empty body",

			input: model.NoteInput{
				Title: "Test Note",
			},

			expectedNote: &ent.Note{
				Title: "Test Note",
				Body:  "",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ctx := ContextT(t)
			preCreateTime := time.Now()

			// Act
			note, err := resolver.Resolvers.Mutation().CreateNote(ctx, tc.input)

			// Assert
			assert.Equal(t, tc.expectedErr, err != nil, "expected error to be %v, got %v", tc.expectedErr, err)
			if tc.expectedNote != nil {
				assert.NotEmpty(t, note.ID)
				assert.Equal(t, tc.expectedNote.Title, note.Title)
				assert.Equal(t, tc.expectedNote.Body, note.Body)
				assert.True(t, note.CreatedAt.After(preCreateTime))
				assert.True(t, note.UpdatedAt.After(preCreateTime))
			}
		})
	}
}

func TestNoteUpdate(t *testing.T) {
	entClient := EntT(t)
	resolver := gql.NewResolver(entClient)
	note := entClient.Note.Create().SetTitle("Test Note").SetBody("Test Note Body").SaveX(ContextT(t))

	tt := []struct {
		name string

		id    int
		input model.NoteInput

		expectedErr  bool
		expectedNote *ent.Note
	}{
		{
			name: "default",

			id: note.ID,
			input: model.NoteInput{
				Title: "Test Note",
				Body:  "Test Note Body",
			},

			expectedNote: &ent.Note{
				Title: "Test Note",
				Body:  "Test Note Body",
			},
		},
		{
			name: "empty title",

			id: note.ID,
			input: model.NoteInput{
				Body: "Test Note Body",
			},

			expectedErr: true,
		},
		{
			name: "title too short",

			id: note.ID,
			input: model.NoteInput{
				Title: "T",
				Body:  "Test Note Body",
			},

			expectedErr: true,
		},
		{
			name: "empty body",

			id: note.ID,
			input: model.NoteInput{
				Title: "Test Note",
			},

			expectedNote: &ent.Note{
				Title: "Test Note",
				Body:  "",
			},
		},
		{
			name: "no change",

			id: note.ID,
			input: model.NoteInput{
				Title: note.Title,
				Body:  note.Body,
			},

			expectedNote: note,
		},
		{
			name: "not found",

			id: 999,
			input: model.NoteInput{
				Title: "Test Note",
				Body:  "Test Note Body",
			},

			expectedErr: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ctx := ContextT(t)
			preCreateTime := time.Now()

			// Act
			note, err := resolver.Resolvers.Mutation().UpdateNote(ctx, tc.id, tc.input)

			// Assert
			assert.Equal(t, tc.expectedErr, err != nil, "expected error to be %v, got %v", tc.expectedErr, err)
			if tc.expectedNote != nil {
				assert.NotEmpty(t, note.ID)
				assert.Equal(t, tc.expectedNote.Title, note.Title)
				assert.Equal(t, tc.expectedNote.Body, note.Body)
				assert.True(t, note.CreatedAt.Before(preCreateTime))
				assert.True(t, note.UpdatedAt.After(preCreateTime))
			}
		})
	}
}

func TestNoteDelete(t *testing.T) {
	entClient := EntT(t)
	resolver := gql.NewResolver(entClient)
	note := entClient.Note.Create().SetTitle("Test Note").SetBody("Test Note Body").SaveX(ContextT(t))

	tt := []struct {
		name string

		id int

		expectedErr bool
		expectedRes bool
	}{
		{
			name: "default",

			id: note.ID,

			expectedRes: true,
		},
		{
			name: "not found",

			id: 999,

			expectedRes: false,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			ctx := ContextT(t)

			// Act
			res, err := resolver.Resolvers.Mutation().DeleteNote(ctx, tc.id)

			// Assert
			assert.Equal(t, tc.expectedErr, err != nil, "expected error to be %v, got %v", tc.expectedErr, err)
			assert.Equal(t, tc.expectedRes, res)
		})
	}
}
