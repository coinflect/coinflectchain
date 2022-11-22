// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package coinflect

import (
	"context"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/snow/consensus/coinflect"
	"github.com/coinflect/coinflectchain/snow/engine/common"
)

// Engine describes the events that can occur on a consensus instance
type Engine interface {
	common.Engine

	// GetVtx returns a vertex by its ID.
	// Returns an error if unknown.
	GetVtx(ctx context.Context, vtxID ids.ID) (coinflect.Vertex, error)
}
