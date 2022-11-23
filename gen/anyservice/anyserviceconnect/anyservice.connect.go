// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: anyservice/anyservice.proto

package anyserviceconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	anypb "google.golang.org/protobuf/types/known/anypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// AnyServiceName is the fully-qualified name of the AnyService service.
	AnyServiceName = "anyservice.AnyService"
)

// AnyServiceClient is a client for the anyservice.AnyService service.
type AnyServiceClient interface {
	Call(context.Context, *connect_go.Request[anypb.Any]) (*connect_go.Response[anypb.Any], error)
}

// NewAnyServiceClient constructs a client for the anyservice.AnyService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAnyServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AnyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &anyServiceClient{
		call: connect_go.NewClient[anypb.Any, anypb.Any](
			httpClient,
			baseURL+"/anyservice.AnyService/Call",
			opts...,
		),
	}
}

// anyServiceClient implements AnyServiceClient.
type anyServiceClient struct {
	call *connect_go.Client[anypb.Any, anypb.Any]
}

// Call calls anyservice.AnyService.Call.
func (c *anyServiceClient) Call(ctx context.Context, req *connect_go.Request[anypb.Any]) (*connect_go.Response[anypb.Any], error) {
	return c.call.CallUnary(ctx, req)
}

// AnyServiceHandler is an implementation of the anyservice.AnyService service.
type AnyServiceHandler interface {
	Call(context.Context, *connect_go.Request[anypb.Any]) (*connect_go.Response[anypb.Any], error)
}

// NewAnyServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAnyServiceHandler(svc AnyServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/anyservice.AnyService/Call", connect_go.NewUnaryHandler(
		"/anyservice.AnyService/Call",
		svc.Call,
		opts...,
	))
	return "/anyservice.AnyService/", mux
}

// UnimplementedAnyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAnyServiceHandler struct{}

func (UnimplementedAnyServiceHandler) Call(context.Context, *connect_go.Request[anypb.Any]) (*connect_go.Response[anypb.Any], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("anyservice.AnyService.Call is not implemented"))
}
