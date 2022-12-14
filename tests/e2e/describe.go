// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package e2e

import (
	ginkgo "github.com/onsi/ginkgo/v2"
)

// DescribeXChain annotates the tests for X-Chain.
// Can run with any type of cluster (e.g., local, addismya, mainnet).
func DescribeXChain(text string, body func()) bool {
	return ginkgo.Describe("[X-Chain] "+text, body)
}

// DescribePChain annotates the tests for P-Chain.
// Can run with any type of cluster (e.g., local, addismya, mainnet).
func DescribePChain(text string, body func()) bool {
	return ginkgo.Describe("[P-Chain] "+text, body)
}
