package common

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/sse"
)

type (
	HumaHandler[I, O any] func(context.Context, *I) (*O, huma.StatusError)
)

type Response[B any] struct {
	Body    *B
	Status  int
	Headers map[string]string
}
type Request[B any] struct {
	Body    *B
	RawBody []byte
}

func WithBody[B any](status int, body B) *Response[B] {
	return &Response[B]{
		Body:    &body,
		Status:  status,
		Headers: nil,
	}
}

type Middleware func(ctx huma.Context, next func(huma.Context))

type RouteCreationContext struct {
	HTTP *http.ServeMux
	API  huma.API
}

func AddHumaRoute[I, O any](rctx RouteCreationContext, humaHandle HumaHandler[I, O], operation huma.Operation) {
	// huma.Register expects the return type of the handle function to be
	// '(*O, error)' but we return 'huma.StatusError' for easier testing. This
	// wrapper type casts the 'huma.StatusError' to error
	convertedHandle := func(ctx context.Context, i *I) (*O, error) {
		output, err := humaHandle(ctx, i)
		return output, err
	}

	huma.Register(
		rctx.API,
		operation,
		convertedHandle,
	)
}

type (
	HumaSSEHandler[I any] func(context.Context, *I, sse.Sender)
)

func AddHumaRouteSSE[I any](rctx RouteCreationContext, eventMap map[string]any, humaHandle HumaSSEHandler[I], operation huma.Operation) {
	// clone to guarantee that there won't be side effects, since append can have them
	sse.Register(rctx.API, operation, eventMap, humaHandle)
}
