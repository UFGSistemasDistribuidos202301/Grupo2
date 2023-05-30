// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: rpc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Aposentadoria_PodeAposentar_FullMethodName = "/aposentadoria.Aposentadoria/PodeAposentar"
)

// AposentadoriaClient is the client API for Aposentadoria service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AposentadoriaClient interface {
	PodeAposentar(ctx context.Context, in *PodeAposentarRequest, opts ...grpc.CallOption) (*PodeAposentarReply, error)
}

type aposentadoriaClient struct {
	cc grpc.ClientConnInterface
}

func NewAposentadoriaClient(cc grpc.ClientConnInterface) AposentadoriaClient {
	return &aposentadoriaClient{cc}
}

func (c *aposentadoriaClient) PodeAposentar(ctx context.Context, in *PodeAposentarRequest, opts ...grpc.CallOption) (*PodeAposentarReply, error) {
	out := new(PodeAposentarReply)
	err := c.cc.Invoke(ctx, Aposentadoria_PodeAposentar_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AposentadoriaServer is the server API for Aposentadoria service.
// All implementations must embed UnimplementedAposentadoriaServer
// for forward compatibility
type AposentadoriaServer interface {
	PodeAposentar(context.Context, *PodeAposentarRequest) (*PodeAposentarReply, error)
	mustEmbedUnimplementedAposentadoriaServer()
}

// UnimplementedAposentadoriaServer must be embedded to have forward compatible implementations.
type UnimplementedAposentadoriaServer struct {
}

func (UnimplementedAposentadoriaServer) PodeAposentar(context.Context, *PodeAposentarRequest) (*PodeAposentarReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PodeAposentar not implemented")
}
func (UnimplementedAposentadoriaServer) mustEmbedUnimplementedAposentadoriaServer() {}

// UnsafeAposentadoriaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AposentadoriaServer will
// result in compilation errors.
type UnsafeAposentadoriaServer interface {
	mustEmbedUnimplementedAposentadoriaServer()
}

func RegisterAposentadoriaServer(s grpc.ServiceRegistrar, srv AposentadoriaServer) {
	s.RegisterService(&Aposentadoria_ServiceDesc, srv)
}

func _Aposentadoria_PodeAposentar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodeAposentarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AposentadoriaServer).PodeAposentar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Aposentadoria_PodeAposentar_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AposentadoriaServer).PodeAposentar(ctx, req.(*PodeAposentarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Aposentadoria_ServiceDesc is the grpc.ServiceDesc for Aposentadoria service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Aposentadoria_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "aposentadoria.Aposentadoria",
	HandlerType: (*AposentadoriaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PodeAposentar",
			Handler:    _Aposentadoria_PodeAposentar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}