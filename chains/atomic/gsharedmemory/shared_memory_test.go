// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gsharedmemory

import (
	"context"
	"io"
	"net"
	"testing"

	"github.com/stretchr/testify/require"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/coinflect/coinflectchain/chains/atomic"
	"github.com/coinflect/coinflectchain/database"
	"github.com/coinflect/coinflectchain/database/memdb"
	"github.com/coinflect/coinflectchain/database/prefixdb"
	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/utils/units"
	"github.com/coinflect/coinflectchain/vms/rpcchainvm/grpcutils"

	sharedmemorypb "github.com/coinflect/coinflectchain/proto/pb/sharedmemory"
)

const (
	bufSize = units.MiB
)

func TestInterface(t *testing.T) {
	require := require.New(t)

	chainID0 := ids.GenerateTestID()
	chainID1 := ids.GenerateTestID()

	for _, test := range atomic.SharedMemoryTests {
		baseDB := memdb.New()
		memoryDB := prefixdb.New([]byte{0}, baseDB)
		testDB := prefixdb.New([]byte{1}, baseDB)

		m := atomic.NewMemory(memoryDB)

		sm0, conn0 := wrapSharedMemory(t, m.NewSharedMemory(chainID0), baseDB)
		sm1, conn1 := wrapSharedMemory(t, m.NewSharedMemory(chainID1), baseDB)

		test(t, chainID0, chainID1, sm0, sm1, testDB)

		err := conn0.Close()
		require.NoError(err)

		err = conn1.Close()
		require.NoError(err)
	}
}

func wrapSharedMemory(t *testing.T, sm atomic.SharedMemory, db database.Database) (atomic.SharedMemory, io.Closer) {
	listener := bufconn.Listen(bufSize)
	serverCloser := grpcutils.ServerCloser{}

	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		server := grpcutils.NewDefaultServer(opts)
		sharedmemorypb.RegisterSharedMemoryServer(server, NewServer(sm, db))
		serverCloser.Add(server)
		return server
	}

	go grpcutils.Serve(listener, serverFunc)

	dialer := grpc.WithContextDialer(
		func(context.Context, string) (net.Conn, error) {
			return listener.Dial()
		},
	)

	dopts := grpcutils.DefaultDialOptions
	dopts = append(dopts, dialer)
	conn, err := grpcutils.Dial("", dopts...)
	if err != nil {
		t.Fatalf("Failed to dial: %s", err)
	}

	rpcsm := NewClient(sharedmemorypb.NewSharedMemoryClient(conn))
	return rpcsm, conn
}
