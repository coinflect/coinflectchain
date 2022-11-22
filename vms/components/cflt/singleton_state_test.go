// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cflt

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/coinflect/coinflectchain/database/memdb"
)

func TestSingletonState(t *testing.T) {
	require := require.New(t)

	db := memdb.New()
	s := NewSingletonState(db)

	isInitialized, err := s.IsInitialized()
	require.NoError(err)
	require.False(isInitialized)

	err = s.SetInitialized()
	require.NoError(err)

	isInitialized, err = s.IsInitialized()
	require.NoError(err)
	require.True(isInitialized)
}
