// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package txs

import (
	"bytes"
	"errors"
	"testing"

	"github.com/coinflect/coinflectchain/codec"
	"github.com/coinflect/coinflectchain/codec/linearcodec"
	"github.com/coinflect/coinflectchain/ids"
	"github.com/coinflect/coinflectchain/vms/components/cflt"
	"github.com/coinflect/coinflectchain/vms/components/verify"
	"github.com/coinflect/coinflectchain/vms/secp256k1fx"
)

func TestInitialStateVerifySerialization(t *testing.T) {
	c := linearcodec.NewDefault()
	if err := c.RegisterType(&secp256k1fx.TransferOutput{}); err != nil {
		t.Fatal(err)
	}
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(CodecVersion, c); err != nil {
		t.Fatal(err)
	}

	expected := []byte{
		// Codec version:
		0x00, 0x00,
		// fxID:
		0x00, 0x00, 0x00, 0x00,
		// num outputs:
		0x00, 0x00, 0x00, 0x01,
		// output:
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x30, 0x39, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xd4, 0x31, 0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02, 0x51, 0x02, 0x5c, 0x61,
		0xfb, 0xcf, 0xc0, 0x78, 0xf6, 0x93, 0x34, 0xf8,
		0x34, 0xbe, 0x6d, 0xd2, 0x6d, 0x55, 0xa9, 0x55,
		0xc3, 0x34, 0x41, 0x28, 0xe0, 0x60, 0x12, 0x8e,
		0xde, 0x35, 0x23, 0xa2, 0x4a, 0x46, 0x1c, 0x89,
		0x43, 0xab, 0x08, 0x59,
	}

	is := &InitialState{
		FxIndex: 0,
		Outs: []verify.State{
			&secp256k1fx.TransferOutput{
				Amt: 12345,
				OutputOwners: secp256k1fx.OutputOwners{
					Locktime:  54321,
					Threshold: 1,
					Addrs: []ids.ShortID{
						{
							0x51, 0x02, 0x5c, 0x61, 0xfb, 0xcf, 0xc0, 0x78,
							0xf6, 0x93, 0x34, 0xf8, 0x34, 0xbe, 0x6d, 0xd2,
							0x6d, 0x55, 0xa9, 0x55,
						},
						{
							0xc3, 0x34, 0x41, 0x28, 0xe0, 0x60, 0x12, 0x8e,
							0xde, 0x35, 0x23, 0xa2, 0x4a, 0x46, 0x1c, 0x89,
							0x43, 0xab, 0x08, 0x59,
						},
					},
				},
			},
		},
	}

	isBytes, err := m.Marshal(CodecVersion, is)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(isBytes, expected) {
		t.Fatalf("Expected:\n0x%x\nResult:\n0x%x",
			expected,
			isBytes,
		)
	}
}

func TestInitialStateVerifyNil(t *testing.T) {
	c := linearcodec.NewDefault()
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(CodecVersion, c); err != nil {
		t.Fatal(err)
	}
	numFxs := 1

	is := (*InitialState)(nil)
	if err := is.Verify(m, numFxs); err == nil {
		t.Fatalf("Should have erred due to nil initial state")
	}
}

func TestInitialStateVerifyUnknownFxID(t *testing.T) {
	c := linearcodec.NewDefault()
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(CodecVersion, c); err != nil {
		t.Fatal(err)
	}
	numFxs := 1

	is := InitialState{
		FxIndex: 1,
	}
	if err := is.Verify(m, numFxs); err == nil {
		t.Fatalf("Should have erred due to unknown FxIndex")
	}
}

func TestInitialStateVerifyNilOutput(t *testing.T) {
	c := linearcodec.NewDefault()
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(CodecVersion, c); err != nil {
		t.Fatal(err)
	}
	numFxs := 1

	is := InitialState{
		FxIndex: 0,
		Outs:    []verify.State{nil},
	}
	if err := is.Verify(m, numFxs); err == nil {
		t.Fatalf("Should have erred due to a nil output")
	}
}

func TestInitialStateVerifyInvalidOutput(t *testing.T) {
	c := linearcodec.NewDefault()
	if err := c.RegisterType(&cflt.TestVerifiable{}); err != nil {
		t.Fatal(err)
	}
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(CodecVersion, c); err != nil {
		t.Fatal(err)
	}
	numFxs := 1

	is := InitialState{
		FxIndex: 0,
		Outs:    []verify.State{&cflt.TestVerifiable{Err: errors.New("")}},
	}
	if err := is.Verify(m, numFxs); err == nil {
		t.Fatalf("Should have erred due to an invalid output")
	}
}

func TestInitialStateVerifyUnsortedOutputs(t *testing.T) {
	c := linearcodec.NewDefault()
	if err := c.RegisterType(&cflt.TestTransferable{}); err != nil {
		t.Fatal(err)
	}
	m := codec.NewDefaultManager()
	if err := m.RegisterCodec(CodecVersion, c); err != nil {
		t.Fatal(err)
	}
	numFxs := 1

	is := InitialState{
		FxIndex: 0,
		Outs: []verify.State{
			&cflt.TestTransferable{Val: 1},
			&cflt.TestTransferable{Val: 0},
		},
	}
	if err := is.Verify(m, numFxs); err == nil {
		t.Fatalf("Should have erred due to unsorted outputs")
	}

	is.Sort(m)

	if err := is.Verify(m, numFxs); err != nil {
		t.Fatal(err)
	}
}
