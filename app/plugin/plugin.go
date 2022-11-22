// Copyright (C) 2019-2022, Ava Labs, Inc. All rights reserved.
// Copyright (C) 2022, Coinflect, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package plugin

import (
	"context"

	"google.golang.org/grpc"

	"github.com/hashicorp/go-plugin"

	"github.com/coinflect/coinflectchain/app"

	pluginpb "github.com/coinflect/coinflectchain/proto/pb/plugin"
)

const Name = "nodeProcess"

var (
	Handshake = plugin.HandshakeConfig{
		ProtocolVersion:  2,
		MagicCookieKey:   "NODE_PROCESS_PLUGIN",
		MagicCookieValue: "dynamic",
	}

	// PluginMap is the map of plugins we can dispense.
	PluginMap = map[string]plugin.Plugin{
		Name: &appPlugin{},
	}

	_ plugin.Plugin     = (*appPlugin)(nil)
	_ plugin.GRPCPlugin = (*appPlugin)(nil)
)

type appPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	app app.App
}

// New will be called by the server side of the plugin to pass into the server
// side PluginMap for dispatching.
func New(app app.App) plugin.Plugin {
	return &appPlugin{app: app}
}

// GRPCServer registers a new GRPC server.
func (p *appPlugin) GRPCServer(_ *plugin.GRPCBroker, s *grpc.Server) error {
	pluginpb.RegisterNodeServer(s, NewServer(p.app))
	return nil
}

// GRPCClient returns a new GRPC client
func (*appPlugin) GRPCClient(_ context.Context, _ *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return NewClient(pluginpb.NewNodeClient(c)), nil
}
