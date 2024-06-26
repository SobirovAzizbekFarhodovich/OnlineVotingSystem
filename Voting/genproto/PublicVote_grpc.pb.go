// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: PublicVote.proto

package genproto

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

// PublicVoteServiceClient is the client API for PublicVoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicVoteServiceClient interface {
	CreatePublicVote(ctx context.Context, in *CreatePublicVoteReq, opts ...grpc.CallOption) (*PublicVote_Void, error)
	GetAllPublicVote(ctx context.Context, in *PublicVote_Void, opts ...grpc.CallOption) (*GetPublicVoteResponse, error)
	GetByIdPublicVote(ctx context.Context, in *GetPublicVoteRequest, opts ...grpc.CallOption) (*PublicVote, error)
	UpdatePublicVote(ctx context.Context, in *UpdatePublicVoteRequest, opts ...grpc.CallOption) (*PublicVote_Void, error)
	DeletePublicVote(ctx context.Context, in *DeletePublicVoteRequest, opts ...grpc.CallOption) (*PublicVote_Void, error)
}

type publicVoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicVoteServiceClient(cc grpc.ClientConnInterface) PublicVoteServiceClient {
	return &publicVoteServiceClient{cc}
}

func (c *publicVoteServiceClient) CreatePublicVote(ctx context.Context, in *CreatePublicVoteReq, opts ...grpc.CallOption) (*PublicVote_Void, error) {
	out := new(PublicVote_Void)
	err := c.cc.Invoke(ctx, "/voting.PublicVoteService/CreatePublicVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicVoteServiceClient) GetAllPublicVote(ctx context.Context, in *PublicVote_Void, opts ...grpc.CallOption) (*GetPublicVoteResponse, error) {
	out := new(GetPublicVoteResponse)
	err := c.cc.Invoke(ctx, "/voting.PublicVoteService/GetAllPublicVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicVoteServiceClient) GetByIdPublicVote(ctx context.Context, in *GetPublicVoteRequest, opts ...grpc.CallOption) (*PublicVote, error) {
	out := new(PublicVote)
	err := c.cc.Invoke(ctx, "/voting.PublicVoteService/GetByIdPublicVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicVoteServiceClient) UpdatePublicVote(ctx context.Context, in *UpdatePublicVoteRequest, opts ...grpc.CallOption) (*PublicVote_Void, error) {
	out := new(PublicVote_Void)
	err := c.cc.Invoke(ctx, "/voting.PublicVoteService/UpdatePublicVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicVoteServiceClient) DeletePublicVote(ctx context.Context, in *DeletePublicVoteRequest, opts ...grpc.CallOption) (*PublicVote_Void, error) {
	out := new(PublicVote_Void)
	err := c.cc.Invoke(ctx, "/voting.PublicVoteService/DeletePublicVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicVoteServiceServer is the server API for PublicVoteService service.
// All implementations must embed UnimplementedPublicVoteServiceServer
// for forward compatibility
type PublicVoteServiceServer interface {
	CreatePublicVote(context.Context, *CreatePublicVoteReq) (*PublicVote_Void, error)
	GetAllPublicVote(context.Context, *PublicVote_Void) (*GetPublicVoteResponse, error)
	GetByIdPublicVote(context.Context, *GetPublicVoteRequest) (*PublicVote, error)
	UpdatePublicVote(context.Context, *UpdatePublicVoteRequest) (*PublicVote_Void, error)
	DeletePublicVote(context.Context, *DeletePublicVoteRequest) (*PublicVote_Void, error)
	mustEmbedUnimplementedPublicVoteServiceServer()
}

// UnimplementedPublicVoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublicVoteServiceServer struct {
}

func (UnimplementedPublicVoteServiceServer) CreatePublicVote(context.Context, *CreatePublicVoteReq) (*PublicVote_Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePublicVote not implemented")
}
func (UnimplementedPublicVoteServiceServer) GetAllPublicVote(context.Context, *PublicVote_Void) (*GetPublicVoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllPublicVote not implemented")
}
func (UnimplementedPublicVoteServiceServer) GetByIdPublicVote(context.Context, *GetPublicVoteRequest) (*PublicVote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdPublicVote not implemented")
}
func (UnimplementedPublicVoteServiceServer) UpdatePublicVote(context.Context, *UpdatePublicVoteRequest) (*PublicVote_Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePublicVote not implemented")
}
func (UnimplementedPublicVoteServiceServer) DeletePublicVote(context.Context, *DeletePublicVoteRequest) (*PublicVote_Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePublicVote not implemented")
}
func (UnimplementedPublicVoteServiceServer) mustEmbedUnimplementedPublicVoteServiceServer() {}

// UnsafePublicVoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicVoteServiceServer will
// result in compilation errors.
type UnsafePublicVoteServiceServer interface {
	mustEmbedUnimplementedPublicVoteServiceServer()
}

func RegisterPublicVoteServiceServer(s grpc.ServiceRegistrar, srv PublicVoteServiceServer) {
	s.RegisterService(&PublicVoteService_ServiceDesc, srv)
}

func _PublicVoteService_CreatePublicVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePublicVoteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).CreatePublicVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting.PublicVoteService/CreatePublicVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).CreatePublicVote(ctx, req.(*CreatePublicVoteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicVoteService_GetAllPublicVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublicVote_Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).GetAllPublicVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting.PublicVoteService/GetAllPublicVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).GetAllPublicVote(ctx, req.(*PublicVote_Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicVoteService_GetByIdPublicVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).GetByIdPublicVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting.PublicVoteService/GetByIdPublicVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).GetByIdPublicVote(ctx, req.(*GetPublicVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicVoteService_UpdatePublicVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePublicVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).UpdatePublicVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting.PublicVoteService/UpdatePublicVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).UpdatePublicVote(ctx, req.(*UpdatePublicVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicVoteService_DeletePublicVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePublicVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicVoteServiceServer).DeletePublicVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting.PublicVoteService/DeletePublicVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicVoteServiceServer).DeletePublicVote(ctx, req.(*DeletePublicVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicVoteService_ServiceDesc is the grpc.ServiceDesc for PublicVoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicVoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "voting.PublicVoteService",
	HandlerType: (*PublicVoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePublicVote",
			Handler:    _PublicVoteService_CreatePublicVote_Handler,
		},
		{
			MethodName: "GetAllPublicVote",
			Handler:    _PublicVoteService_GetAllPublicVote_Handler,
		},
		{
			MethodName: "GetByIdPublicVote",
			Handler:    _PublicVoteService_GetByIdPublicVote_Handler,
		},
		{
			MethodName: "UpdatePublicVote",
			Handler:    _PublicVoteService_UpdatePublicVote_Handler,
		},
		{
			MethodName: "DeletePublicVote",
			Handler:    _PublicVoteService_DeletePublicVote_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "PublicVote.proto",
}
