package proxy

import (
	"golang.org/x/net/context"

	//tsdb "github.com/grafana/grafana/pkg/tsdb"

	proto "github.com/grafana/grafana/pkg/tsdb/models"
	plugin "github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var PluginMap = map[string]plugin.Plugin{
	"tsdb_mock": &TsdbPluginImpl{},
}

type TsdbPlugin interface {
	Query(ctx context.Context, req *proto.TsdbQuery) (*proto.Response, error)
}

type TsdbPluginImpl struct { //LOL IMPL LOL
	plugin.NetRPCUnsupportedPlugin
	Plugin TsdbPlugin
}

func (p *TsdbPluginImpl) GRPCServer(s *grpc.Server) error {
	proto.RegisterTsdbPluginServer(s, &GRPCServer{p.Plugin})
	return nil
}

func (p *TsdbPluginImpl) GRPCClient(c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{proto.NewTsdbPluginClient(c)}, nil
}