// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cflt

const (
	codecVersion = 0
)

// Addressable is the interface a feature extension must provide to be able to
// be tracked as a part of the utxo set for a set of addresses
type Addressable interface {
	Addresses() [][]byte
}