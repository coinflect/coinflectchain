// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"
	"fmt"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/snow"
	"github.com/coinflect/coinflectchain/vms/components/cflt"
	"github.com/coinflect/coinflectchain/vms/secp256k1fx"
)

var (
	_ UnsignedTx = (*ImportTx)(nil)

	errNoImportInputs = errors.New("tx has no imported inputs")
)

// ImportTx is an unsigned importTx
type ImportTx struct {
	BaseTx `serialize:"true"`

	// Which chain to consume the funds from
	SourceChain ids.ID `serialize:"true" json:"sourceChain"`

	// Inputs that consume UTXOs produced on the chain
	ImportedInputs []*cflt.TransferableInput `serialize:"true" json:"importedInputs"`
}

// InitCtx sets the FxID fields in the inputs and outputs of this
// [ImportTx]. Also sets the [ctx] to the given [vm.ctx] so that
// the addresses can be json marshalled into human readable format
func (tx *ImportTx) InitCtx(ctx *snow.Context) {
	tx.BaseTx.InitCtx(ctx)
	for _, in := range tx.ImportedInputs {
		in.FxID = secp256k1fx.ID
	}
}

// InputUTXOs returns the UTXOIDs of the imported funds
func (tx *ImportTx) InputUTXOs() ids.Set {
	set := ids.NewSet(len(tx.ImportedInputs))
	for _, in := range tx.ImportedInputs {
		set.Add(in.InputID())
	}
	return set
}

func (tx *ImportTx) InputIDs() ids.Set {
	inputs := tx.BaseTx.InputIDs()
	atomicInputs := tx.InputUTXOs()
	inputs.Union(atomicInputs)
	return inputs
}

// SyntacticVerify this transaction is well-formed
func (tx *ImportTx) SyntacticVerify(ctx *snow.Context) error {
	switch {
	case tx == nil:
		return ErrNilTx
	case tx.SyntacticallyVerified: // already passed syntactic verification
		return nil
	case len(tx.ImportedInputs) == 0:
		return errNoImportInputs
	}

	if err := tx.BaseTx.SyntacticVerify(ctx); err != nil {
		return err
	}

	for _, in := range tx.ImportedInputs {
		if err := in.Verify(); err != nil {
			return fmt.Errorf("input failed verification: %w", err)
		}
	}
	if !cflt.IsSortedAndUniqueTransferableInputs(tx.ImportedInputs) {
		return errInputsNotSortedUnique
	}

	tx.SyntacticallyVerified = true
	return nil
}

func (tx *ImportTx) Visit(visitor Visitor) error {
	return visitor.ImportTx(tx)
}
