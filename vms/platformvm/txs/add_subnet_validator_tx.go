// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"errors"
	"time"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/snow"
	"github.com/coinflect/coinflectchain/utils/constants"
	"github.com/coinflect/coinflectchain/utils/crypto/bls"
	"github.com/coinflect/coinflectchain/vms/components/verify"
	"github.com/coinflect/coinflectchain/vms/platformvm/validator"
)

var (
	_ StakerTx = (*AddSubnetValidatorTx)(nil)

	errAddPrimaryNetworkValidator = errors.New("can't add primary network validator with AddSubnetValidatorTx")
)

// AddSubnetValidatorTx is an unsigned addSubnetValidatorTx
type AddSubnetValidatorTx struct {
	// Metadata, inputs and outputs
	BaseTx `serialize:"true"`
	// The validator
	Validator validator.SubnetValidator `serialize:"true" json:"validator"`
	// Auth that will be allowing this validator into the network
	SubnetAuth verify.Verifiable `serialize:"true" json:"subnetAuthorization"`
}

func (tx *AddSubnetValidatorTx) SubnetID() ids.ID {
	return tx.Validator.Subnet
}

func (tx *AddSubnetValidatorTx) NodeID() ids.NodeID {
	return tx.Validator.NodeID
}

func (*AddSubnetValidatorTx) PublicKey() (*bls.PublicKey, bool, error) {
	return nil, false, nil
}

func (tx *AddSubnetValidatorTx) StartTime() time.Time {
	return tx.Validator.StartTime()
}

func (tx *AddSubnetValidatorTx) EndTime() time.Time {
	return tx.Validator.EndTime()
}

func (tx *AddSubnetValidatorTx) Weight() uint64 {
	return tx.Validator.Wght
}

func (*AddSubnetValidatorTx) PendingPriority() Priority {
	return SubnetPermissionedValidatorPendingPriority
}

func (*AddSubnetValidatorTx) CurrentPriority() Priority {
	return SubnetPermissionedValidatorCurrentPriority
}

// SyntacticVerify returns nil iff [tx] is valid
func (tx *AddSubnetValidatorTx) SyntacticVerify(ctx *snow.Context) error {
	switch {
	case tx == nil:
		return ErrNilTx
	case tx.SyntacticallyVerified: // already passed syntactic verification
		return nil
	case tx.Validator.Subnet == constants.PrimaryNetworkID:
		return errAddPrimaryNetworkValidator
	}

	if err := tx.BaseTx.SyntacticVerify(ctx); err != nil {
		return err
	}
	if err := verify.All(&tx.Validator, tx.SubnetAuth); err != nil {
		return err
	}

	// cache that this is valid
	tx.SyntacticallyVerified = true
	return nil
}

func (tx *AddSubnetValidatorTx) Visit(visitor Visitor) error {
	return visitor.AddSubnetValidatorTx(tx)
}
