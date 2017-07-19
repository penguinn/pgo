package msgpackrpc2

import (
	"context"
	"errors"
	"sync/atomic"

	"github.com/hashicorp/go-multierror"
	"github.com/smallnest/rpcx/core"
)

var (
	// nextCallSeq is used to assign a unique sequence number
	// to each call made with CallWithCodec
	nextCallSeq uint64
)

// CallWithCodec is used to perform the same actions as core.Client.Call but
// in a much cheaper way. It assumes the underlying connection is not being
// shared with multiple concurrent RPCs. The request/response must be syncronous.
func CallWithCodec(ctx context.Context, cc core.ClientCodec, method string, args interface{}, resp interface{}) error {
	request := core.Request{
		Seq:           atomic.AddUint64(&nextCallSeq, 1),
		ServiceMethod: method,
	}
	if err := cc.WriteRequest(ctx, &request, args); err != nil {
		return err
	}
	var response core.Response
	if err := cc.ReadResponseHeader(&response); err != nil {
		return err
	}
	if response.Error != "" {
		err := errors.New(response.Error)
		if readErr := cc.ReadResponseBody(nil); readErr != nil {
			err = multierror.Append(err, readErr)
		}
		return core.ServerError(err.Error())
	}
	if err := cc.ReadResponseBody(resp); err != nil {
		return err
	}
	return nil
}
