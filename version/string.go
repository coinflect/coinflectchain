// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package version

import (
	"fmt"
)

var (
	// String is displayed when CLI arg --version is used
	String string

	// GitCommit is set in the build script at compile time
	GitCommit string
)

func init() {
	format := "%s [database=%s, rpcchainvm=%d"
	args := []interface{}{
		CurrentApp,
		CurrentDatabase,
		RPCChainVMProtocol,
	}
	if GitCommit != "" {
		format += ", commit=%s"
		args = append(args, GitCommit)
	}
	format += "]\n"
	String = fmt.Sprintf(format, args...)
}
