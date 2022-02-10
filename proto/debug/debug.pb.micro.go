// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: debug/debug.proto

package debug

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro-community/micro/v3/service/api"
	client "github.com/micro-community/micro/v3/service/client"
	server "github.com/micro-community/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Debug service

func NewDebugEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Debug service

type DebugService interface {
	Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (*LogResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...client.CallOption) (*HealthResponse, error)
	Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error)
	Trace(ctx context.Context, in *TraceRequest, opts ...client.CallOption) (*TraceResponse, error)
}

type debugService struct {
	c    client.Client
	name string
}

func NewDebugService(name string, c client.Client) DebugService {
	return &debugService{
		c:    c,
		name: name,
	}
}

func (c *debugService) Log(ctx context.Context, in *LogRequest, opts ...client.CallOption) (*LogResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Log", in)
	out := new(LogResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) Health(ctx context.Context, in *HealthRequest, opts ...client.CallOption) (*HealthResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Health", in)
	out := new(HealthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) Stats(ctx context.Context, in *StatsRequest, opts ...client.CallOption) (*StatsResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Stats", in)
	out := new(StatsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *debugService) Trace(ctx context.Context, in *TraceRequest, opts ...client.CallOption) (*TraceResponse, error) {
	req := c.c.NewRequest(c.name, "Debug.Trace", in)
	out := new(TraceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Debug service

type DebugHandler interface {
	Log(context.Context, *LogRequest, *LogResponse) error
	Health(context.Context, *HealthRequest, *HealthResponse) error
	Stats(context.Context, *StatsRequest, *StatsResponse) error
	Trace(context.Context, *TraceRequest, *TraceResponse) error
}

func RegisterDebugHandler(s server.Server, hdlr DebugHandler, opts ...server.HandlerOption) error {
	type debug interface {
		Log(ctx context.Context, in *LogRequest, out *LogResponse) error
		Health(ctx context.Context, in *HealthRequest, out *HealthResponse) error
		Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error
		Trace(ctx context.Context, in *TraceRequest, out *TraceResponse) error
	}
	type Debug struct {
		debug
	}
	h := &debugHandler{hdlr}
	return s.Handle(s.NewHandler(&Debug{h}, opts...))
}

type debugHandler struct {
	DebugHandler
}

func (h *debugHandler) Log(ctx context.Context, in *LogRequest, out *LogResponse) error {
	return h.DebugHandler.Log(ctx, in, out)
}

func (h *debugHandler) Health(ctx context.Context, in *HealthRequest, out *HealthResponse) error {
	return h.DebugHandler.Health(ctx, in, out)
}

func (h *debugHandler) Stats(ctx context.Context, in *StatsRequest, out *StatsResponse) error {
	return h.DebugHandler.Stats(ctx, in, out)
}

func (h *debugHandler) Trace(ctx context.Context, in *TraceRequest, out *TraceResponse) error {
	return h.DebugHandler.Trace(ctx, in, out)
}
