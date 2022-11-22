// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package blocks

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/vms/components/cflt"
	"github.com/coinflect/coinflectchain/vms/components/verify"
	"github.com/coinflect/coinflectchain/vms/platformvm/txs"
)

func TestNewApricotAtomicBlock(t *testing.T) {
	require := require.New(t)

	parentID := ids.GenerateTestID()
	height := uint64(1337)
	tx := &txs.Tx{
		Unsigned: &txs.ImportTx{
			BaseTx: txs.BaseTx{
				BaseTx: cflt.BaseTx{
					Ins:  []*cflt.TransferableInput{},
					Outs: []*cflt.TransferableOutput{},
				},
			},
			ImportedInputs: []*cflt.TransferableInput{},
		},
		Creds: []verify.Verifiable{},
	}
	require.NoError(tx.Sign(txs.Codec, nil))

	blk, err := NewApricotAtomicBlock(
		parentID,
		height,
		tx,
	)
	require.NoError(err)

	// Make sure the block and tx are initialized
	require.NotNil(blk.Bytes())
	require.NotNil(blk.Tx.Bytes())
	require.NotEqual(ids.Empty, blk.Tx.ID())
	require.Equal(tx.Bytes(), blk.Tx.Bytes())
	require.Equal(parentID, blk.Parent())
	require.Equal(height, blk.Height())
}
