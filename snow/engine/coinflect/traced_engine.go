// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package coinflect

import (
	"context"

	"go.opentelemetry.io/otel/attribute"

	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/snow/consensus/coinflect"
	"github.com/coinflect/coinflectchain/snow/engine/common"
	"github.com/coinflect/coinflectchain/trace"
)

var _ Engine = (*tracedEngine)(nil)

type tracedEngine struct {
	common.Engine
	engine Engine
	tracer trace.Tracer
}

func TraceEngine(engine Engine, tracer trace.Tracer) Engine {
	return &tracedEngine{
		Engine: common.TraceEngine(engine, tracer),
		engine: engine,
		tracer: tracer,
	}
}

func (e *tracedEngine) GetVtx(ctx context.Context, vtxID ids.ID) (coinflect.Vertex, error) {
	ctx, span := e.tracer.Start(ctx, "tracedEngine.GetVtx", oteltrace.WithAttributes(
		attribute.Stringer("vtxID", vtxID),
	))
	defer span.End()

	return e.engine.GetVtx(ctx, vtxID)
}
