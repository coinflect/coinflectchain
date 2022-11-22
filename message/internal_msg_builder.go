// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

//nolint:stylecheck // proto generates interfaces that fail linting
package message

import (
	"time"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/utils/timer/mockable"
	"github.com/coinflect/coinflectchain/version"
)

var (
	disconnected  = &Disconnected{}
	timeout       = &Timeout{}
	gossipRequest = &GossipRequest{}

	_ chainIDGetter       = (*GetStateSummaryFrontierFailed)(nil)
	_ requestIDGetter     = (*GetStateSummaryFrontierFailed)(nil)
	_ chainIDGetter       = (*GetAcceptedStateSummaryFailed)(nil)
	_ requestIDGetter     = (*GetAcceptedStateSummaryFailed)(nil)
	_ chainIDGetter       = (*GetAcceptedFrontierFailed)(nil)
	_ requestIDGetter     = (*GetAcceptedFrontierFailed)(nil)
	_ chainIDGetter       = (*GetAcceptedFailed)(nil)
	_ requestIDGetter     = (*GetAcceptedFailed)(nil)
	_ chainIDGetter       = (*GetAncestorsFailed)(nil)
	_ requestIDGetter     = (*GetAncestorsFailed)(nil)
	_ chainIDGetter       = (*GetFailed)(nil)
	_ requestIDGetter     = (*GetFailed)(nil)
	_ chainIDGetter       = (*QueryFailed)(nil)
	_ requestIDGetter     = (*QueryFailed)(nil)
	_ chainIDGetter       = (*AppRequestFailed)(nil)
	_ requestIDGetter     = (*AppRequestFailed)(nil)
	_ sourceChainIDGetter = (*CrossChainAppRequest)(nil)
	_ chainIDGetter       = (*CrossChainAppRequest)(nil)
	_ requestIDGetter     = (*CrossChainAppRequest)(nil)
	_ sourceChainIDGetter = (*CrossChainAppRequestFailed)(nil)
	_ chainIDGetter       = (*CrossChainAppRequestFailed)(nil)
	_ requestIDGetter     = (*CrossChainAppRequestFailed)(nil)
	_ sourceChainIDGetter = (*CrossChainAppResponse)(nil)
	_ chainIDGetter       = (*CrossChainAppResponse)(nil)
	_ requestIDGetter     = (*CrossChainAppResponse)(nil)
)

type GetStateSummaryFrontierFailed struct {
	ChainID   ids.ID
	RequestID uint32
}

func (m *GetStateSummaryFrontierFailed) GetChainId() []byte {
	return m.ChainID[:]
}

func (m *GetStateSummaryFrontierFailed) GetRequestId() uint32 {
	return m.RequestID
}

func InternalGetStateSummaryFrontierFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     GetStateSummaryFrontierFailedOp,
		message: &GetStateSummaryFrontierFailed{
			ChainID:   chainID,
			RequestID: requestID,
		},
		expiration: mockable.MaxTime,
	}
}

type GetAcceptedStateSummaryFailed struct {
	ChainID   ids.ID
	RequestID uint32
}

func (m *GetAcceptedStateSummaryFailed) GetChainId() []byte {
	return m.ChainID[:]
}

func (m *GetAcceptedStateSummaryFailed) GetRequestId() uint32 {
	return m.RequestID
}

func InternalGetAcceptedStateSummaryFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     GetAcceptedStateSummaryFailedOp,
		message: &GetAcceptedStateSummaryFailed{
			ChainID:   chainID,
			RequestID: requestID,
		},
		expiration: mockable.MaxTime,
	}
}

type GetAcceptedFrontierFailed struct {
	ChainID   ids.ID
	RequestID uint32
}

func (m *GetAcceptedFrontierFailed) GetChainId() []byte {
	return m.ChainID[:]
}

func (m *GetAcceptedFrontierFailed) GetRequestId() uint32 {
	return m.RequestID
}

func InternalGetAcceptedFrontierFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     GetAcceptedFrontierFailedOp,
		message: &GetAcceptedFrontierFailed{
			ChainID:   chainID,
			RequestID: requestID,
		},
		expiration: mockable.MaxTime,
	}
}

type GetAcceptedFailed struct {
	ChainID   ids.ID
	RequestID uint32
}

func (m *GetAcceptedFailed) GetChainId() []byte {
	return m.ChainID[:]
}

func (m *GetAcceptedFailed) GetRequestId() uint32 {
	return m.RequestID
}

func InternalGetAcceptedFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     GetAcceptedFailedOp,
		message: &GetAcceptedFailed{
			ChainID:   chainID,
			RequestID: requestID,
		},
		expiration: mockable.MaxTime,
	}
}

type GetAncestorsFailed struct {
	ChainID   ids.ID
	RequestID uint32
}

func (m *GetAncestorsFailed) GetChainId() []byte {
	return m.ChainID[:]
}

func (m *GetAncestorsFailed) GetRequestId() uint32 {
	return m.RequestID
}

func InternalGetAncestorsFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     GetAncestorsFailedOp,
		message: &GetAncestorsFailed{
			ChainID:   chainID,
			RequestID: requestID,
		},
		expiration: mockable.MaxTime,
	}
}

type GetFailed struct {
	ChainID   ids.ID
	RequestID uint32
}

func (m *GetFailed) GetChainId() []byte {
	return m.ChainID[:]
}

func (m *GetFailed) GetRequestId() uint32 {
	return m.RequestID
}

func InternalGetFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     GetFailedOp,
		message: &GetFailed{
			ChainID:   chainID,
			RequestID: requestID,
		},
		expiration: mockable.MaxTime,
	}
}

type QueryFailed struct {
	ChainID   ids.ID
	RequestID uint32
}

func (m *QueryFailed) GetChainId() []byte {
	return m.ChainID[:]
}

func (m *QueryFailed) GetRequestId() uint32 {
	return m.RequestID
}

func InternalQueryFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     QueryFailedOp,
		message: &QueryFailed{
			ChainID:   chainID,
			RequestID: requestID,
		},
		expiration: mockable.MaxTime,
	}
}

type AppRequestFailed struct {
	ChainID   ids.ID
	RequestID uint32
}

func (m *AppRequestFailed) GetChainId() []byte {
	return m.ChainID[:]
}

func (m *AppRequestFailed) GetRequestId() uint32 {
	return m.RequestID
}

func InternalAppRequestFailed(
	nodeID ids.NodeID,
	chainID ids.ID,
	requestID uint32,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     AppRequestFailedOp,
		message: &AppRequestFailed{
			ChainID:   chainID,
			RequestID: requestID,
		},
		expiration: mockable.MaxTime,
	}
}

type CrossChainAppRequest struct {
	SourceChainID      ids.ID
	DestinationChainID ids.ID
	RequestID          uint32
	Message            []byte
}

func (m *CrossChainAppRequest) GetSourceChainID() ids.ID {
	return m.SourceChainID
}

func (m *CrossChainAppRequest) GetChainId() []byte {
	return m.DestinationChainID[:]
}

func (m *CrossChainAppRequest) GetRequestId() uint32 {
	return m.RequestID
}

func InternalCrossChainAppRequest(
	nodeID ids.NodeID,
	sourceChainID ids.ID,
	destinationChainID ids.ID,
	requestID uint32,
	deadline time.Duration,
	msg []byte,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     CrossChainAppRequestOp,
		message: &CrossChainAppRequest{
			SourceChainID:      sourceChainID,
			DestinationChainID: destinationChainID,
			RequestID:          requestID,
			Message:            msg,
		},
		expiration: time.Now().Add(deadline),
	}
}

type CrossChainAppRequestFailed struct {
	SourceChainID      ids.ID
	DestinationChainID ids.ID
	RequestID          uint32
}

func (m *CrossChainAppRequestFailed) GetSourceChainID() ids.ID {
	return m.SourceChainID
}

func (m *CrossChainAppRequestFailed) GetChainId() []byte {
	return m.DestinationChainID[:]
}

func (m *CrossChainAppRequestFailed) GetRequestId() uint32 {
	return m.RequestID
}

func InternalCrossChainAppRequestFailed(
	nodeID ids.NodeID,
	sourceChainID ids.ID,
	destinationChainID ids.ID,
	requestID uint32,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     CrossChainAppRequestFailedOp,
		message: &CrossChainAppRequestFailed{
			SourceChainID:      sourceChainID,
			DestinationChainID: destinationChainID,
			RequestID:          requestID,
		},
		expiration: mockable.MaxTime,
	}
}

type CrossChainAppResponse struct {
	SourceChainID      ids.ID
	DestinationChainID ids.ID
	RequestID          uint32
	Message            []byte
}

func (m *CrossChainAppResponse) GetSourceChainID() ids.ID {
	return m.SourceChainID
}

func (m *CrossChainAppResponse) GetChainId() []byte {
	return m.DestinationChainID[:]
}

func (m *CrossChainAppResponse) GetRequestId() uint32 {
	return m.RequestID
}

func InternalCrossChainAppResponse(
	nodeID ids.NodeID,
	sourceChainID ids.ID,
	destinationChainID ids.ID,
	requestID uint32,
	msg []byte,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     CrossChainAppResponseOp,
		message: &CrossChainAppResponse{
			SourceChainID:      sourceChainID,
			DestinationChainID: destinationChainID,
			RequestID:          requestID,
			Message:            msg,
		},
		expiration: mockable.MaxTime,
	}
}

type Connected struct {
	NodeVersion *version.Application
}

func InternalConnected(nodeID ids.NodeID, nodeVersion *version.Application) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     ConnectedOp,
		message: &Connected{
			NodeVersion: nodeVersion,
		},
		expiration: mockable.MaxTime,
	}
}

type Disconnected struct{}

func InternalDisconnected(nodeID ids.NodeID) InboundMessage {
	return &inboundMessage{
		nodeID:     nodeID,
		op:         DisconnectedOp,
		message:    disconnected,
		expiration: mockable.MaxTime,
	}
}

type VMMessage struct {
	Notification uint32
}

func InternalVMMessage(
	nodeID ids.NodeID,
	notification uint32,
) InboundMessage {
	return &inboundMessage{
		nodeID: nodeID,
		op:     NotifyOp,
		message: &VMMessage{
			Notification: notification,
		},
		expiration: mockable.MaxTime,
	}
}

type GossipRequest struct{}

func InternalGossipRequest(
	nodeID ids.NodeID,
) InboundMessage {
	return &inboundMessage{
		nodeID:     nodeID,
		op:         GossipRequestOp,
		message:    gossipRequest,
		expiration: mockable.MaxTime,
	}
}

type Timeout struct{}

func InternalTimeout(nodeID ids.NodeID) InboundMessage {
	return &inboundMessage{
		nodeID:     nodeID,
		op:         TimeoutOp,
		message:    timeout,
		expiration: mockable.MaxTime,
	}
}
