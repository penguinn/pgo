package prpcx

import (
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/codec"
	"github.com/smallnest/rpcx/plugin"
)

var server *rpcx.Server

func InitRpcx(name string, class interface{}) {
	server = rpcx.NewServer()
	server.PluginContainer.Add(plugin.NewMetricsPlugin())
	server.ServerCodecFunc = codec.NewJSONRPC2ServerCodec
	server.RegisterName(name, class)
}

func Run(addr string) error{
	err := server.Serve("tcp", addr)
	return err
}


