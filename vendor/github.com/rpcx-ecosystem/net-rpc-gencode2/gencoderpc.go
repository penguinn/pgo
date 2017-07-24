package gencodec2

import (
	"io"
	"net"
	"time"

	"github.com/smallnest/rpcx/core"
)

// Dial connects to a Gencode-RPC server at the specified network address.
func Dial(network, address string) (*core.Client, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return NewClient(conn), err
}

// DialTimeout connects to a Gencode-RPC server at the specified network address with timeout.
func DialTimeout(network, address string, timeout time.Duration) (*core.Client, error) {
	conn, err := net.DialTimeout(network, address, timeout)
	if err != nil {
		return nil, err
	}
	return NewClient(conn), err
}

// NewClient returns a new core.Client to handle requests to the set of
// services at the other end of the connection.
func NewClient(conn io.ReadWriteCloser) *core.Client {
	return core.NewClientWithCodec(NewGencodeClientCodec(conn))
}

// ServeConn runs the Gencode-RPC server on a single connection. ServeConn
// blocks, serving the connection until the client hangs up. The caller
// typically invokes ServeConn in a go statement.
func ServeConn(conn io.ReadWriteCloser) {
	core.ServeCodec(NewGencodeServerCodec(conn))
}
