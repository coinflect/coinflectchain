// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package state

import (
	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/vms/components/cflt"
)

type UTXOGetter interface {
	GetUTXO(utxoID ids.ID) (*cflt.UTXO, error)
}

type UTXOAdder interface {
	AddUTXO(utxo *cflt.UTXO)
}

type UTXODeleter interface {
	DeleteUTXO(utxoID ids.ID)
}
