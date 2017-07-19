// Package msgpackrpc implements a MessagePack-RPC ClientCodec and ServerCodec
// for the rpc package, using the same API as the Go standard library
// for jsoncore.
package msgpackrpc2

import (
	"io"
	"net"

	"github.com/smallnest/rpcx/core"
)

// Dial connects to a MessagePack-RPC server at the specified network address.
func Dial(network, address string) (*core.Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewClient(conn), err
}

// NewClient returns a new core.Client to handle requests to the set of
// services at the other end of the connection.
func NewClient(conn io.ReadWriteCloser) *core.Client {
	return core.NewClientWithCodec(NewClientCodec(conn))
}

// NewClientCodec returns a new core.ClientCodec using MessagePack-RPC on conn.
func NewClientCodec(conn io.ReadWriteCloser) core.ClientCodec {
	return NewCodec(true, true, conn)
}

// NewServerCodec returns a new core.ServerCodec using MessagePack-RPC on conn.
func NewServerCodec(conn io.ReadWriteCloser) core.ServerCodec {
	return NewCodec(true, true, conn)
}

// ServeConn runs the MessagePack-RPC server on a single connection. ServeConn
// blocks, serving the connection until the client hangs up. The caller
// typically invokes ServeConn in a go statement.
func ServeConn(conn io.ReadWriteCloser) {
	core.ServeCodec(NewServerCodec(conn))
}
