// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/coinflect/coinflectchain/snow"
	"github.com/coinflect/coinflectchain/snow/consensus/snowball"
	"github.com/coinflect/coinflectchain/snow/consensus/snowman"
	"github.com/coinflect/coinflectchain/snow/engine/common"
	"github.com/coinflect/coinflectchain/snow/engine/snowman/block"
	"github.com/coinflect/coinflectchain/snow/validators"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	common.AllGetsServer

	Ctx        *snow.ConsensusContext
	VM         block.ChainVM
	Sender     common.Sender
	Validators validators.Set
	Params     snowball.Parameters
	Consensus  snowman.Consensus
}
