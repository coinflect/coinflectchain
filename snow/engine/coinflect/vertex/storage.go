// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"context"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/snow/consensus/coinflect"
)

// Storage defines the persistent storage that is required by the consensus
// engine.
type Storage interface {
	// Get a vertex by its hash from storage.
	GetVtx(ctx context.Context, vtxID ids.ID) (coinflect.Vertex, error)
	// Edge returns a list of accepted vertex IDs with no accepted children.
	Edge(ctx context.Context) (vtxIDs []ids.ID)
	// Returns "true" if accepted frontier ("Edge") is stop vertex.
	StopVertexAccepted(ctx context.Context) (bool, error)
}
