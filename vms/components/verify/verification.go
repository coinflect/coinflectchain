// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package verify

import "github.com/coinflect/coinflectchain/snow"

// Verifiable can be verified
type Verifiable interface {
	Verify() error
}

// State that can be verified
type State interface {
	snow.ContextInitializable
	Verifiable
	VerifyState() error
}

// All returns nil if all the verifiables were verified with no errors
func All(verifiables ...Verifiable) error {
	for _, verifiable := range verifiables {
		if err := verifiable.Verify(); err != nil {
			return err
		}
	}
	return nil
}
