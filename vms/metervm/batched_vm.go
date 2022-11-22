// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package metervm

import (
	"context"
	"time"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/snow/consensus/snowman"
	"github.com/coinflect/coinflectchain/snow/engine/snowman/block"
)

func (vm *blockVM) GetAncestors(
	ctx context.Context,
	blkID ids.ID,
	maxBlocksNum int,
	maxBlocksSize int,
	maxBlocksRetrivalTime time.Duration,
) ([][]byte, error) {
	if vm.bVM == nil {
		return nil, block.ErrRemoteVMNotImplemented
	}

	start := vm.clock.Time()
	ancestors, err := vm.bVM.GetAncestors(
		ctx,
		blkID,
		maxBlocksNum,
		maxBlocksSize,
		maxBlocksRetrivalTime,
	)
	end := vm.clock.Time()
	vm.blockMetrics.getAncestors.Observe(float64(end.Sub(start)))
	return ancestors, err
}

func (vm *blockVM) BatchedParseBlock(ctx context.Context, blks [][]byte) ([]snowman.Block, error) {
	if vm.bVM == nil {
		return nil, block.ErrRemoteVMNotImplemented
	}

	start := vm.clock.Time()
	blocks, err := vm.bVM.BatchedParseBlock(ctx, blks)
	end := vm.clock.Time()
	vm.blockMetrics.batchedParseBlock.Observe(float64(end.Sub(start)))

	wrappedBlocks := make([]snowman.Block, len(blocks))
	for i, block := range blocks {
		wrappedBlocks[i] = &meterBlock{
			Block: block,
			vm:    vm,
		}
	}
	return wrappedBlocks, err
}
