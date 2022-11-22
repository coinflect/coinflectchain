// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package executor

import (
	"github.com/coinflect/coinflectchain/snow"
	"github.com/coinflect/coinflectchain/snow/uptime"
	"github.com/coinflect/coinflectchain/utils"
	"github.com/coinflect/coinflectchain/utils/timer/mockable"
	"github.com/coinflect/coinflectchain/vms/platformvm/config"
	"github.com/coinflect/coinflectchain/vms/platformvm/fx"
	"github.com/coinflect/coinflectchain/vms/platformvm/reward"
	"github.com/coinflect/coinflectchain/vms/platformvm/utxo"
)

type Backend struct {
	Config       *config.Config
	Ctx          *snow.Context
	Clk          *mockable.Clock
	Fx           fx.Fx
	FlowChecker  utxo.Verifier
	Uptimes      uptime.Manager
	Rewards      reward.Calculator
	Bootstrapped *utils.AtomicBool
}
