// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import "github.com/stretchr/testify/require"

func equalHeader(require *require.Assertions, want, have Header) {
	require.Equal(want.ChainID(), have.ChainID())
	require.Equal(want.ParentID(), have.ParentID())
	require.Equal(want.BodyID(), have.BodyID())
}
