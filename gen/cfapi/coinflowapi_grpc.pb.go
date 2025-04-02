// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: proto/coinflowapi.proto

package cfapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Coinflow_PostTransaction_FullMethodName = "/coinflow_api.Coinflow/PostTransaction"
	Coinflow_GetTransaction_FullMethodName  = "/coinflow_api.Coinflow/GetTransaction"
)

// CoinflowClient is the client API for Coinflow service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoinflowClient interface {
	PostTransaction(ctx context.Context, in *PostTransactionRequest, opts ...grpc.CallOption) (*PostTransactionResponse, error)
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
}

type coinflowClient struct {
	cc grpc.ClientConnInterface
}

func NewCoinflowClient(cc grpc.ClientConnInterface) CoinflowClient {
	return &coinflowClient{cc}
}

func (c *coinflowClient) PostTransaction(ctx context.Context, in *PostTransactionRequest, opts ...grpc.CallOption) (*PostTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostTransactionResponse)
	err := c.cc.Invoke(ctx, Coinflow_PostTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinflowClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, Coinflow_GetTransaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoinflowServer is the server API for Coinflow service.
// All implementations must embed UnimplementedCoinflowServer
// for forward compatibility.
type CoinflowServer interface {
	PostTransaction(context.Context, *PostTransactionRequest) (*PostTransactionResponse, error)
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	mustEmbedUnimplementedCoinflowServer()
}

// UnimplementedCoinflowServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCoinflowServer struct{}

func (UnimplementedCoinflowServer) PostTransaction(context.Context, *PostTransactionRequest) (*PostTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostTransaction not implemented")
}
func (UnimplementedCoinflowServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedCoinflowServer) mustEmbedUnimplementedCoinflowServer() {}
func (UnimplementedCoinflowServer) testEmbeddedByValue()                  {}

// UnsafeCoinflowServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoinflowServer will
// result in compilation errors.
type UnsafeCoinflowServer interface {
	mustEmbedUnimplementedCoinflowServer()
}

func RegisterCoinflowServer(s grpc.ServiceRegistrar, srv CoinflowServer) {
	// If the following call pancis, it indicates UnimplementedCoinflowServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Coinflow_ServiceDesc, srv)
}

func _Coinflow_PostTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinflowServer).PostTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coinflow_PostTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinflowServer).PostTransaction(ctx, req.(*PostTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Coinflow_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinflowServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Coinflow_GetTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinflowServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Coinflow_ServiceDesc is the grpc.ServiceDesc for Coinflow service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Coinflow_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coinflow_api.Coinflow",
	HandlerType: (*CoinflowServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostTransaction",
			Handler:    _Coinflow_PostTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Coinflow_GetTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/coinflowapi.proto",
}
