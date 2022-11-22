// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nftfx

import (
	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/snow"
	"github.com/coinflect/coinflectchain/vms"
)

var (
	_ vms.Factory = (*Factory)(nil)

	// ID that this Fx uses when labeled
	ID = ids.ID{'n', 'f', 't', 'f', 'x'}
)

type Factory struct{}

func (*Factory) New(*snow.Context) (interface{}, error) {
	return &Fx{}, nil
}
