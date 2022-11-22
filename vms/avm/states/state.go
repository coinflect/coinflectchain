// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package states

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/coinflect/coinflectchain/database"
	"github.com/coinflect/coinflectchain/database/prefixdb"
	"github.com/coinflect/coinflectchain/vms/avm/txs"
	"github.com/coinflect/coinflectchain/vms/components/cflt"
)

var (
	utxoPrefix      = []byte("utxo")
	statusPrefix    = []byte("status")
	singletonPrefix = []byte("singleton")
	txPrefix        = []byte("tx")

	_ State = (*state)(nil)
)

// State persistently maintains a set of UTXOs, transaction, statuses, and
// singletons.
type State interface {
	cflt.UTXOState
	cflt.StatusState
	cflt.SingletonState
	TxState
}

type state struct {
	cflt.UTXOState
	cflt.StatusState
	cflt.SingletonState
	TxState
}

func New(db database.Database, parser txs.Parser, metrics prometheus.Registerer) (State, error) {
	utxoDB := prefixdb.New(utxoPrefix, db)
	statusDB := prefixdb.New(statusPrefix, db)
	singletonDB := prefixdb.New(singletonPrefix, db)
	txDB := prefixdb.New(txPrefix, db)

	utxoState, err := cflt.NewMeteredUTXOState(utxoDB, parser.Codec(), metrics)
	if err != nil {
		return nil, err
	}

	statusState, err := cflt.NewMeteredStatusState(statusDB, metrics)
	if err != nil {
		return nil, err
	}

	txState, err := NewTxState(txDB, parser, metrics)
	return &state{
		UTXOState:      utxoState,
		StatusState:    statusState,
		SingletonState: cflt.NewSingletonState(singletonDB),
		TxState:        txState,
	}, err
}
