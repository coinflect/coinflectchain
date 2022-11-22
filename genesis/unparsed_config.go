// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"encoding/hex"
	"errors"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/utils/formatting/address"
)

var errInvalidETHAddress = errors.New("invalid eth address")

type UnparsedAllocation struct {
	ETHAddr        string         `json:"ethAddr"`
	CFLTAddr       string         `json:"cfltAddr"`
	InitialAmount  uint64         `json:"initialAmount"`
	UnlockSchedule []LockedAmount `json:"unlockSchedule"`
}

func (ua UnparsedAllocation) Parse() (Allocation, error) {
	a := Allocation{
		InitialAmount:  ua.InitialAmount,
		UnlockSchedule: ua.UnlockSchedule,
	}

	if len(ua.ETHAddr) < 2 {
		return a, errInvalidETHAddress
	}

	ethAddrBytes, err := hex.DecodeString(ua.ETHAddr[2:])
	if err != nil {
		return a, err
	}
	ethAddr, err := ids.ToShortID(ethAddrBytes)
	if err != nil {
		return a, err
	}
	a.ETHAddr = ethAddr

	_, _, cfltAddrBytes, err := address.Parse(ua.CFLTAddr)
	if err != nil {
		return a, err
	}
	cfltAddr, err := ids.ToShortID(cfltAddrBytes)
	if err != nil {
		return a, err
	}
	a.CFLTAddr = cfltAddr

	return a, nil
}

type UnparsedStaker struct {
	NodeID        ids.NodeID `json:"nodeID"`
	RewardAddress string     `json:"rewardAddress"`
	DelegationFee uint32     `json:"delegationFee"`
}

func (us UnparsedStaker) Parse() (Staker, error) {
	s := Staker{
		NodeID:        us.NodeID,
		DelegationFee: us.DelegationFee,
	}

	_, _, cfltAddrBytes, err := address.Parse(us.RewardAddress)
	if err != nil {
		return s, err
	}
	cfltAddr, err := ids.ToShortID(cfltAddrBytes)
	if err != nil {
		return s, err
	}
	s.RewardAddress = cfltAddr
	return s, nil
}

// UnparsedConfig contains the genesis addresses used to construct a genesis
type UnparsedConfig struct {
	NetworkID uint32 `json:"networkID"`

	Allocations []UnparsedAllocation `json:"allocations"`

	StartTime                  uint64           `json:"startTime"`
	InitialStakeDuration       uint64           `json:"initialStakeDuration"`
	InitialStakeDurationOffset uint64           `json:"initialStakeDurationOffset"`
	InitialStakedFunds         []string         `json:"initialStakedFunds"`
	InitialStakers             []UnparsedStaker `json:"initialStakers"`

	CChainGenesis string `json:"cChainGenesis"`

	Message string `json:"message"`
}

func (uc UnparsedConfig) Parse() (Config, error) {
	c := Config{
		NetworkID:                  uc.NetworkID,
		Allocations:                make([]Allocation, len(uc.Allocations)),
		StartTime:                  uc.StartTime,
		InitialStakeDuration:       uc.InitialStakeDuration,
		InitialStakeDurationOffset: uc.InitialStakeDurationOffset,
		InitialStakedFunds:         make([]ids.ShortID, len(uc.InitialStakedFunds)),
		InitialStakers:             make([]Staker, len(uc.InitialStakers)),
		CChainGenesis:              uc.CChainGenesis,
		Message:                    uc.Message,
	}
	for i, ua := range uc.Allocations {
		a, err := ua.Parse()
		if err != nil {
			return c, err
		}
		c.Allocations[i] = a
	}
	for i, isa := range uc.InitialStakedFunds {
		_, _, cfltAddrBytes, err := address.Parse(isa)
		if err != nil {
			return c, err
		}
		cfltAddr, err := ids.ToShortID(cfltAddrBytes)
		if err != nil {
			return c, err
		}
		c.InitialStakedFunds[i] = cfltAddr
	}
	for i, uis := range uc.InitialStakers {
		is, err := uis.Parse()
		if err != nil {
			return c, err
		}
		c.InitialStakers[i] = is
	}
	return c, nil
}
