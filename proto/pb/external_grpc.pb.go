// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ExternalClient is the client API for External service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type externalClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalClient(cc grpc.ClientConnInterface) ExternalClient {
	return &externalClient{cc}
}

func (c *externalClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/eugeneuskov.chat.external.External/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalServer is the server API for External service.
// All implementations should embed UnimplementedExternalServer
// for forward compatibility
type ExternalServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*emptypb.Empty, error)
}

// UnimplementedExternalServer should be embedded to have forward compatible implementations.
type UnimplementedExternalServer struct {
}

func (UnimplementedExternalServer) CreateUser(context.Context, *CreateUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}

// UnsafeExternalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalServer will
// result in compilation errors.
type UnsafeExternalServer interface {
	mustEmbedUnimplementedExternalServer()
}

func RegisterExternalServer(s grpc.ServiceRegistrar, srv ExternalServer) {
	s.RegisterService(&External_ServiceDesc, srv)
}

func _External_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eugeneuskov.chat.external.External/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// External_ServiceDesc is the grpc.ServiceDesc for External service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var External_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eugeneuskov.chat.external.External",
	HandlerType: (*ExternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _External_CreateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/external.proto",
}
