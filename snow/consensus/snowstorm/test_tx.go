// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowstorm

import (
	"context"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/snow/choices"
)

var _ Tx = (*TestTx)(nil)

// TestTx is a useful test tx
type TestTx struct {
	choices.TestDecidable

	DependenciesV    []Tx
	DependenciesErrV error
	InputIDsV        []ids.ID
	HasWhitelistV    bool
	WhitelistV       ids.Set
	WhitelistErrV    error
	VerifyV          error
	BytesV           []byte
}

func (t *TestTx) Dependencies() ([]Tx, error) {
	return t.DependenciesV, t.DependenciesErrV
}

func (t *TestTx) InputIDs() []ids.ID {
	return t.InputIDsV
}

func (t *TestTx) HasWhitelist() bool {
	return t.HasWhitelistV
}

func (t *TestTx) Whitelist(context.Context) (ids.Set, error) {
	return t.WhitelistV, t.WhitelistErrV
}

func (t *TestTx) Verify(context.Context) error {
	return t.VerifyV
}

func (t *TestTx) Bytes() []byte {
	return t.BytesV
}
