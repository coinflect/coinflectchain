// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package atomic

import (
	"github.com/coinflect/coinflectchain/codec"
	"github.com/coinflect/coinflectchain/codec/linearcodec"
)

const codecVersion = 0

// codecManager is used to marshal and unmarshal dbElements and chain IDs.
var codecManager codec.Manager

func init() {
	linearCodec := linearcodec.NewDefault()
	codecManager = codec.NewDefaultManager()
	if err := codecManager.RegisterCodec(codecVersion, linearCodec); err != nil {
		panic(err)
	}
}
