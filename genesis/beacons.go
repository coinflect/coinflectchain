// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"github.com/coinflect/coinflectchain/utils/constants"
	"github.com/coinflect/coinflectchain/utils/sampler"
)

// getIPs returns the beacon IPs for each network
func getIPs(networkID uint32) []string {
	switch networkID {
	case constants.MainnetID:
		return []string{
			"54.94.43.49:9651",
			"52.79.47.77:9651",
			"18.229.206.191:9651",
			"3.34.221.73:9651",
		}
	case constants.AddismyaID:
		return []string{
			"3.218.101.121:9651",
			"107.21.104.83:9651",
			"18.208.118.222:9651",
			"44.209.136.202:9651",
		}
	default:
		return nil
	}
}

// getNodeIDs returns the beacon node IDs for each network
func getNodeIDs(networkID uint32) []string {
	switch networkID {
	case constants.MainnetID:
		return []string{
			"NodeID-8TBmUoLt5nWTg3SrPjPo2yvzAj8ejaU7Y",
			"NodeID-DVfmNENxkR1r8pMVfNqr8ZCcVgWCFyPGr",
			"NodeID-HnMcbE1is18q9oXYcyXsYCMXiS7S3KqFM",
			"NodeID-LLryLUacrtGKXKJ7vMQzmZNeoAP5YVbXg",
		}
	case constants.AddismyaID:
		return []string{
			"NodeID-8TBmUoLt5nWTg3SrPjPo2yvzAj8ejaU7Y",
			"NodeID-DVfmNENxkR1r8pMVfNqr8ZCcVgWCFyPGr",
			"NodeID-HnMcbE1is18q9oXYcyXsYCMXiS7S3KqFM",
			"NodeID-LLryLUacrtGKXKJ7vMQzmZNeoAP5YVbXg",
		}
	default:
		return nil
	}
}

// SampleBeacons returns the some beacons this node should connect to
func SampleBeacons(networkID uint32, count int) ([]string, []string) {
	ips := getIPs(networkID)
	ids := getNodeIDs(networkID)

	if numIPs := len(ips); numIPs < count {
		count = numIPs
	}

	sampledIPs := make([]string, 0, count)
	sampledIDs := make([]string, 0, count)

	s := sampler.NewUniform()
	_ = s.Initialize(uint64(len(ips)))
	indices, _ := s.Sample(count)
	for _, index := range indices {
		sampledIPs = append(sampledIPs, ips[int(index)])
		sampledIDs = append(sampledIDs, ids[int(index)])
	}

	return sampledIPs, sampledIDs
}
