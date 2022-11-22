// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package server

import (
	"net/http"
)

type middlewareHandler struct {
	before, after func()
	handler       http.Handler
}

func (mh middlewareHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if mh.before != nil {
		mh.before()
	}
	if mh.after != nil {
		defer mh.after()
	}
	mh.handler.ServeHTTP(writer, request)
}
