// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package coinflect

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/coinflect/coinflectchain/database/memdb"
	"github.com/coinflect/coinflectchain/snow/consensus/coinflect"
	"github.com/coinflect/coinflectchain/snow/consensus/snowball"
	"github.com/coinflect/coinflectchain/snow/engine/coinflect/bootstrap"
	"github.com/coinflect/coinflectchain/snow/engine/coinflect/vertex"
	"github.com/coinflect/coinflectchain/snow/engine/common"
	"github.com/coinflect/coinflectchain/snow/engine/common/queue"
)

func DefaultConfig() (common.Config, bootstrap.Config, Config) {
	vtxBlocked, _ := queue.NewWithMissing(memdb.New(), "", prometheus.NewRegistry())
	txBlocked, _ := queue.New(memdb.New(), "", prometheus.NewRegistry())

	commonCfg := common.DefaultConfigTest()

	bootstrapConfig := bootstrap.Config{
		Config:     commonCfg,
		VtxBlocked: vtxBlocked,
		TxBlocked:  txBlocked,
		Manager:    &vertex.TestManager{},
		VM:         &vertex.TestVM{},
	}

	engineConfig := Config{
		Ctx:        bootstrapConfig.Ctx,
		VM:         bootstrapConfig.VM,
		Manager:    bootstrapConfig.Manager,
		Sender:     bootstrapConfig.Sender,
		Validators: bootstrapConfig.Validators,
		Params: coinflect.Parameters{
			Parameters: snowball.Parameters{
				K:                       1,
				Alpha:                   1,
				BetaVirtuous:            1,
				BetaRogue:               2,
				ConcurrentRepolls:       1,
				OptimalProcessing:       100,
				MaxOutstandingItems:     1,
				MaxItemProcessingTime:   1,
				MixedQueryNumPushVdr:    1,
				MixedQueryNumPushNonVdr: 1,
			},
			Parents:   2,
			BatchSize: 1,
		},
		Consensus: &coinflect.Topological{},
	}

	return commonCfg, bootstrapConfig, engineConfig
}
