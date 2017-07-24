package gencodec2

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/smallnest/rpcx/core"
)

type serverCodec struct {
	conn io.ReadWriteCloser
	enc  *gencodeEncoder
	dec  *gencodeDecoder
	mu   sync.Mutex // exclusive writer lock
	resp ResponseHeader
	req  RequestHeader
	c    io.Closer
	w    *bufio.Writer
}

// NewGencodeServerCodec returns a new core.ServerCodec.
//
// A ServerCodec implements reading of RPC requests and writing of RPC
// responses for the server side of an RPC session. The server calls
// ReadRequestHeader and ReadRequestBody in pairs to read requests from the
// connection, and it calls WriteResponse to write a response back. The
// server calls Close when finished with the connection.
func NewGencodeServerCodec(rwc io.ReadWriteCloser) core.ServerCodec {
	w := bufio.NewWriter(rwc)
	r := bufio.NewReader(rwc)
	return &serverCodec{
		conn: rwc,
		enc:  newGencodeEncoder(w),
		dec:  newGencodeDecoder(r),
		w:    w,
		c:    rwc,
	}
}

func (c *serverCodec) WriteResponse(ctx context.Context, resp *core.Response, body interface{}) error {
	c.mu.Lock()
	c.resp.ServiceMethod = resp.ServiceMethod
	c.resp.Seq = resp.Seq
	c.resp.Error = resp.Error

	err := c.enc.Encode(&c.resp)
	if err != nil {
		c.mu.Unlock()
		return err
	}
	if err = c.enc.Encode(body); err != nil {
		c.mu.Unlock()
		return err
	}

	err = c.w.Flush()
	c.mu.Unlock()
	return err
}

func (c *serverCodec) ReadRequestHeader(ctx context.Context, req *core.Request) error {
	if err := c.dec.Decode(&c.req); err != nil {
		return err
	}

	req.ServiceMethod = c.req.ServiceMethod
	req.Seq = c.req.Seq
	return nil
}

func (c *serverCodec) ReadRequestBody(ctx context.Context, body interface{}) error {
	if pb, ok := body.(genCodeMessage); ok {
		return c.dec.Decode(pb)
	}
	return fmt.Errorf("%T does not implement genCodeMessage", body)
}

func (c *serverCodec) Close() error { return c.c.Close() }
