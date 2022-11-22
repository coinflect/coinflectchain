// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package peer

import (
	"time"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/message"
	"github.com/coinflect/coinflectchain/network/throttling"
	"github.com/coinflect/coinflectchain/snow/networking/router"
	"github.com/coinflect/coinflectchain/snow/networking/tracker"
	"github.com/coinflect/coinflectchain/snow/validators"
	"github.com/coinflect/coinflectchain/utils/logging"
	"github.com/coinflect/coinflectchain/utils/timer/mockable"
	"github.com/coinflect/coinflectchain/version"
)

type Config struct {
	// Size, in bytes, of the buffer this peer reads messages into
	ReadBufferSize int
	// Size, in bytes, of the buffer this peer writes messages into
	WriteBufferSize int
	Clock           mockable.Clock
	Metrics         *Metrics
	MessageCreator  message.Creator

	Log                  logging.Logger
	InboundMsgThrottler  throttling.InboundMsgThrottler
	Network              Network
	Router               router.InboundHandler
	VersionCompatibility version.Compatibility
	MySubnets            ids.Set
	Beacons              validators.Set
	NetworkID            uint32
	PingFrequency        time.Duration
	PongTimeout          time.Duration
	MaxClockDifference   time.Duration

	// Unix time of the last message sent and received respectively
	// Must only be accessed atomically
	LastSent, LastReceived int64

	// Tracks CPU/disk usage caused by each peer.
	ResourceTracker tracker.ResourceTracker
}
