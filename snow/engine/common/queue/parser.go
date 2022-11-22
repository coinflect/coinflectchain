// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package queue

import "context"

// Parser allows parsing a job from bytes.
type Parser interface {
	Parse(context.Context, []byte) (Job, error)
}
