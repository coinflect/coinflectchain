// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/coinflect/coinflectchain/vms/platformvm/blocks"
)

const Version = blocks.Version

var Codec = blocks.GenesisCodec
