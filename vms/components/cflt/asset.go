// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package cflt

import (
	"errors"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/vms/components/verify"
)

var (
	errNilAssetID   = errors.New("nil asset ID is not valid")
	errEmptyAssetID = errors.New("empty asset ID is not valid")

	_ verify.Verifiable = (*Asset)(nil)
)

type Asset struct {
	ID ids.ID `serialize:"true" json:"assetID"`
}

// AssetID returns the ID of the contained asset
func (asset *Asset) AssetID() ids.ID {
	return asset.ID
}

func (asset *Asset) Verify() error {
	switch {
	case asset == nil:
		return errNilAssetID
	case asset.ID == ids.Empty:
		return errEmptyAssetID
	default:
		return nil
	}
}
