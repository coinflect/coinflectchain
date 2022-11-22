// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package units

// Denominations of value
const (
	NanoCflt  uint64 = 1
	MicroCflt uint64 = 1000 * NanoCflt
	Schmeckle uint64 = 49*MicroCflt + 463*NanoCflt
	MilliCflt uint64 = 1000 * MicroCflt
	Cflt      uint64 = 1000 * MilliCflt
	KiloCflt  uint64 = 1000 * Cflt
	MegaCflt  uint64 = 1000 * KiloCflt
)
