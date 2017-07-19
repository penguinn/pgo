package prpcx

import "github.com/smallnest/rpcx"

var server *rpcx.Server

func InitRpcx(name string, class interface{}) {
	server = rpcx.NewServer()
	server.RegisterName(name, class)
}

func Run(addr string) error{
	err := server.Serve("tcp", addr)
	return err
}


