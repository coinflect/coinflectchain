// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/coinflect/coinflectchain/vms/components/cflt"
	"github.com/coinflect/coinflectchain/vms/platformvm/txs"
)

// UTXO adds messages to UTXOs
type UTXO struct {
	cflt.UTXO `serialize:"true"`
	Message   []byte `serialize:"true" json:"message"`
}

// Genesis represents a genesis state of the platform chain
type Genesis struct {
	UTXOs         []*UTXO   `serialize:"true"`
	Validators    []*txs.Tx `serialize:"true"`
	Chains        []*txs.Tx `serialize:"true"`
	Timestamp     uint64    `serialize:"true"`
	InitialSupply uint64    `serialize:"true"`
	Message       string    `serialize:"true"`
}

func Parse(genesisBytes []byte) (*Genesis, error) {
	gen := &Genesis{}
	if _, err := Codec.Unmarshal(genesisBytes, gen); err != nil {
		return nil, err
	}
	for _, tx := range gen.Validators {
		if err := tx.Sign(txs.GenesisCodec, nil); err != nil {
			return nil, err
		}
	}
	for _, tx := range gen.Chains {
		if err := tx.Sign(txs.GenesisCodec, nil); err != nil {
			return nil, err
		}
	}
	return gen, nil
}

// State represents the genesis state of the platform chain
type State struct {
	UTXOs         []*cflt.UTXO
	Validators    []*txs.Tx
	Chains        []*txs.Tx
	Timestamp     uint64
	InitialSupply uint64
}

func ParseState(genesisBytes []byte) (*State, error) {
	genesis, err := Parse(genesisBytes)
	if err != nil {
		return nil, err
	}

	utxos := make([]*cflt.UTXO, 0, len(genesis.UTXOs))
	for _, utxo := range genesis.UTXOs {
		utxos = append(utxos, &utxo.UTXO)
	}

	return &State{
		UTXOs:         utxos,
		Validators:    genesis.Validators,
		Chains:        genesis.Chains,
		Timestamp:     genesis.Timestamp,
		InitialSupply: genesis.InitialSupply,
	}, nil
}
