// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package runner

type Config struct {
	// If true, displays version and exits during startup
	DisplayVersionAndExit bool

	// Path to the build directory
	BuildDir string

	// If true, run as a plugin
	PluginMode bool
}
