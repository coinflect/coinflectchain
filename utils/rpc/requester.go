// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpc

import (
	"context"
	"net/url"
)

var _ EndpointRequester = (*coinflectEndpointRequester)(nil)

type EndpointRequester interface {
	SendRequest(ctx context.Context, method string, params interface{}, reply interface{}, options ...Option) error
}

type coinflectEndpointRequester struct {
	uri string
}

func NewEndpointRequester(uri string) EndpointRequester {
	return &coinflectEndpointRequester{
		uri: uri,
	}
}

func (e *coinflectEndpointRequester) SendRequest(
	ctx context.Context,
	method string,
	params interface{},
	reply interface{},
	options ...Option,
) error {
	uri, err := url.Parse(e.uri)
	if err != nil {
		return err
	}

	return SendJSONRequest(
		ctx,
		uri,
		method,
		params,
		reply,
		options...,
	)
}