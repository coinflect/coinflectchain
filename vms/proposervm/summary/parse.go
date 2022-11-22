// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package summary

import (
	"fmt"

	"github.com/coinflect/coinflectchain/utils/hashing"
)

func Parse(bytes []byte) (StateSummary, error) {
	summary := stateSummary{
		id:    hashing.ComputeHash256Array(bytes),
		bytes: bytes,
	}
	version, err := c.Unmarshal(bytes, &summary)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal summary due to: %w", err)
	}
	if version != codecVersion {
		return nil, errWrongCodecVersion
	}
	return &summary, nil
}