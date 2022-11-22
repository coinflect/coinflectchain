// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package router

import (
	"context"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/message"
	"github.com/coinflect/coinflectchain/version"
)

var _ InboundHandler = InboundHandlerFunc(nil)

// InboundHandler handles inbound messages
type InboundHandler interface {
	HandleInbound(context.Context, message.InboundMessage)
}

// The ExternalRouterFunc type is an adapter to allow the use of ordinary
// functions as ExternalRouters. If f is a function with the appropriate
// signature, ExternalRouterFunc(f) is an ExternalRouter that calls f.
type InboundHandlerFunc func(context.Context, message.InboundMessage)

func (f InboundHandlerFunc) HandleInbound(ctx context.Context, msg message.InboundMessage) {
	f(ctx, msg)
}

// ExternalHandler handles messages from external parties
type ExternalHandler interface {
	InboundHandler

	Connected(nodeID ids.NodeID, nodeVersion *version.Application, subnetID ids.ID)
	Disconnected(nodeID ids.NodeID)
}
