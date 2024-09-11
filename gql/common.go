package gql

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/nikkomiu/gentql/ent"
	"github.com/nikkomiu/gentql/ent/note"
)

// Hello is the resolver for the hello field.
func (r *queryResolver) Hello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, nodeID string) (ent.Noder, error) {
	rawNodeID, err := base64.RawURLEncoding.DecodeString(nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse node id: base64 decode error")
	}

	splitNodeID := strings.Split(string(rawNodeID), ":")
	if len(splitNodeID) != 2 {
		return nil, fmt.Errorf("failed to parse node id: wrong number of parts")
	}

	switch splitNodeID[0] {
	// add other cases here (custom table names, non-ent types, etc.)

	case note.Table:
		id, err := strconv.Atoi(splitNodeID[1])
		if err != nil {
			return nil, err
		}
		return r.ent.Noder(ctx, id, ent.WithFixedNodeType(splitNodeID[0]))

	default:
		return nil, fmt.Errorf("failed parse node id type")
	}
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
