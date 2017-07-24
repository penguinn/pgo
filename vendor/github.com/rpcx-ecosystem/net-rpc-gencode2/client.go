package gencodec2

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/smallnest/rpcx/core"
)

type clientCodec struct {
	mu   sync.Mutex // exclusive writer lock
	enc  *gencodeEncoder
	dec  *gencodeDecoder
	req  RequestHeader
	resp ResponseHeader

	w *bufio.Writer
	c io.Closer
}

// NewGencodeClientCodec returns a new core.Client.
//
// A ClientCodec implements writing of RPC requests and reading of RPC
// responses for the client side of an RPC session. The client calls
// WriteRequest to write a request to the connection and calls
// ReadResponseHeader and ReadResponseBody in pairs to read responses. The
// client calls Close when finished with the connection.
func NewGencodeClientCodec(rwc io.ReadWriteCloser) core.ClientCodec {
	w := bufio.NewWriter(rwc)
	r := bufio.NewReader(rwc)
	return &clientCodec{
		enc: newGencodeEncoder(w),
		dec: newGencodeDecoder(r),
		w:   w,
		c:   rwc,
	}
}

func (c *clientCodec) WriteRequest(ctx context.Context, req *core.Request, body interface{}) error {
	c.mu.Lock()
	c.req.ServiceMethod = req.ServiceMethod
	c.req.Seq = req.Seq

	err := c.enc.Encode(&c.req)
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

func (c *clientCodec) ReadResponseHeader(resp *core.Response) error {
	c.resp.Error = ""
	if err := c.dec.Decode(&c.resp); err != nil {
		return err
	}

	resp.ServiceMethod = c.resp.ServiceMethod
	resp.Seq = c.resp.Seq
	resp.Error = c.resp.Error
	return nil
}

func (c *clientCodec) ReadResponseBody(body interface{}) (err error) {
	if pb, ok := body.(genCodeMessage); ok {
		return c.dec.Decode(pb)
	}
	return fmt.Errorf("%T does not implement genCodeMessage", body)
}

func (c *clientCodec) Close() error { return c.c.Close() }
