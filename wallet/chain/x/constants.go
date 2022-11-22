// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package x

import (
	"github.com/coinflect/coinflectchain/vms/avm/fxs"
	"github.com/coinflect/coinflectchain/vms/avm/txs"
	"github.com/coinflect/coinflectchain/vms/nftfx"
	"github.com/coinflect/coinflectchain/vms/propertyfx"
	"github.com/coinflect/coinflectchain/vms/secp256k1fx"
)

const (
	SECP256K1FxIndex = 0
	NFTFxIndex       = 1
	PropertyFxIndex  = 2
)

// Parser to support serialization and deserialization
var Parser txs.Parser

func init() {
	var err error
	Parser, err = txs.NewParser([]fxs.Fx{
		&secp256k1fx.Fx{},
		&nftfx.Fx{},
		&propertyfx.Fx{},
	})
	if err != nil {
		panic(err)
	}
}
