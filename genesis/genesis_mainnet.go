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
	//go:embed genesis_mainnet.json
	mainnetGenesisConfigJSON []byte

	// MainnetParams are the params used for mainnet
	MainnetParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                         units.MilliCflt,
			CreateAssetTxFee:              10 * units.MilliCflt,
			CreateSubnetTxFee:             1 * units.Cflt,
			TransformSubnetTxFee:          10 * units.Cflt,
			CreateBlockchainTxFee:         1 * units.Cflt,
			AddPrimaryNetworkValidatorFee: 0,
			AddPrimaryNetworkDelegatorFee: 0,
			AddSubnetValidatorFee:         units.MilliCflt,
			AddSubnetDelegatorFee:         units.MilliCflt,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 2 * units.KiloCflt,
			MaxValidatorStake: 3 * units.MegaCflt,
			MinDelegatorStake: 25 * units.Cflt,
			MinDelegationFee:  20000, // 2%
			MinStakeDuration:  2 * 7 * 24 * time.Hour,
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
