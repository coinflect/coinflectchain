// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"github.com/coinflect/coinflectchain/ids"
)

// Subnet describes the standard interface of a subnet description
type Subnet interface {
	// Returns true iff the subnet is done bootstrapping
	IsBootstrapped() bool

	// Bootstrapped marks the named chain as being bootstrapped
	Bootstrapped(chainID ids.ID)
}
