// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package coinflect

import (
	"github.com/coinflect/coinflectchain/snow"
	"github.com/coinflect/coinflectchain/snow/consensus/coinflect"
	"github.com/coinflect/coinflectchain/snow/engine/coinflect/vertex"
	"github.com/coinflect/coinflectchain/snow/engine/common"
	"github.com/coinflect/coinflectchain/snow/validators"
)

// Config wraps all the parameters needed for an coinflect engine
type Config struct {
	Ctx *snow.ConsensusContext
	common.AllGetsServer
	VM         vertex.DAGVM
	Manager    vertex.Manager
	Sender     common.Sender
	Validators validators.Set

	Params    coinflect.Parameters
	Consensus coinflect.Consensus
}
