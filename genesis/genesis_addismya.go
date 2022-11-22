// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	_ "embed"

	"github.com/coinflect/coinflectchain/utils/units"
	"github.com/coinflect/coinflectchain/vms/platformvm/reward"
)

var (
	//go:embed genesis_addismya.json
	addismyaGenesisConfigJSON []byte

	// AddismyaParams are the params used for the addismya testnet
	AddismyaParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                         units.MilliCflt,
			CreateAssetTxFee:              10 * units.MilliCflt,
			CreateSubnetTxFee:             100 * units.MilliCflt,
			TransformSubnetTxFee:          1 * units.Cflt,
			CreateBlockchainTxFee:         100 * units.MilliCflt,
			AddPrimaryNetworkValidatorFee: 0,
			AddPrimaryNetworkDelegatorFee: 0,
			AddSubnetValidatorFee:         units.MilliCflt,
			AddSubnetDelegatorFee:         units.MilliCflt,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 1 * units.Cflt,
			MaxValidatorStake: 3 * units.MegaCflt,
			MinDelegatorStake: 1 * units.Cflt,
			MinDelegationFee:  20000, // 2%
			MinStakeDuration:  24 * time.Hour,
			MaxStakeDuration:  365 * 24 * time.Hour,
			RewardConfig: reward.Config{
				MaxConsumptionRate: .12 * reward.PercentDenominator,
				MinConsumptionRate: .10 * reward.PercentDenominator,
				MintingPeriod:      365 * 24 * time.Hour,
				SupplyCap:          720 * units.MegaCflt,
			},
		},
	}
)
