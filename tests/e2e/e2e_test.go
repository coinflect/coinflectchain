// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package e2e_test

import (
	"flag"
	"testing"

	ginkgo "github.com/onsi/ginkgo/v2"

	"github.com/onsi/gomega"

	"github.com/coinflect/coinflectchain/tests/e2e"

	// ensure test packages are scanned by ginkgo
	_ "github.com/coinflect/coinflectchain/tests/e2e/banff"
	_ "github.com/coinflect/coinflectchain/tests/e2e/p"
	_ "github.com/coinflect/coinflectchain/tests/e2e/ping"
	_ "github.com/coinflect/coinflectchain/tests/e2e/static-handlers"
	_ "github.com/coinflect/coinflectchain/tests/e2e/x/transfer"
	_ "github.com/coinflect/coinflectchain/tests/e2e/x/whitelist-vtx"
)

func TestE2E(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "e2e test suites")
}

var (
	// helpers to parse test flags
	logLevel string

	networkRunnerGRPCEp              string
	networkRunnerCoinflectChainExecPath string
	networkRunnerCoinflectChainLogLevel string

	uris string

	testKeysFile string
)

func init() {
	flag.StringVar(
		&logLevel,
		"log-level",
		"info",
		"log level",
	)

	flag.StringVar(
		&networkRunnerGRPCEp,
		"network-runner-grpc-endpoint",
		"",
		"[optional] gRPC server endpoint for network-runner (only required for local network-runner tests)",
	)
	flag.StringVar(
		&networkRunnerCoinflectChainExecPath,
		"network-runner-coinflectchain-path",
		"",
		"[optional] coinflectchain executable path (only required for local network-runner tests)",
	)
	flag.StringVar(
		&networkRunnerCoinflectChainLogLevel,
		"network-runner-coinflectchain-log-level",
		"INFO",
		"[optional] coinflectchain log-level (only required for local network-runner tests)",
	)

	// e.g., custom network HTTP RPC endpoints
	flag.StringVar(
		&uris,
		"uris",
		"",
		"HTTP RPC endpoint URIs for coinflect node (comma-separated, required to run against existing cluster)",
	)

	// file that contains a list of new-line separated secp256k1 private keys
	flag.StringVar(
		&testKeysFile,
		"test-keys-file",
		"",
		"file that contains a list of new-line separated hex-encoded secp256k1 private keys (assume test keys are pre-funded, for test networks)",
	)
}

var _ = ginkgo.BeforeSuite(func() {
	err := e2e.Env.ConfigCluster(
		logLevel,
		networkRunnerGRPCEp,
		networkRunnerCoinflectChainExecPath,
		networkRunnerCoinflectChainLogLevel,
		uris,
		testKeysFile,
	)
	gomega.Expect(err).Should(gomega.BeNil())

	// check cluster can be started
	err = e2e.Env.StartCluster()
	gomega.Expect(err).Should(gomega.BeNil())

	// load keys
	err = e2e.Env.LoadKeys()
	gomega.Expect(err).Should(gomega.BeNil())

	// take initial snapshot. cluster will be switched off
	err = e2e.Env.SnapInitialState()
	gomega.Expect(err).Should(gomega.BeNil())

	// restart cluster
	err = e2e.Env.RestoreInitialState(false /*switchOffNetworkFirst*/)
	gomega.Expect(err).Should(gomega.BeNil())
})

var _ = ginkgo.AfterSuite(func() {
	err := e2e.Env.ShutdownCluster()
	gomega.Expect(err).Should(gomega.BeNil())
})
