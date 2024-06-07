// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: Party.proto

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

// PartyServiceClient is the client API for PartyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartyServiceClient interface {
	CreateParty(ctx context.Context, in *CreatePartyRequest, opts ...grpc.CallOption) (*VoidPartyResponse, error)
	GetPartyById(ctx context.Context, in *ByIdPartyRequest, opts ...grpc.CallOption) (*GetPartyResponse, error)
	GetAllParties(ctx context.Context, in *FilterPartyRequest, opts ...grpc.CallOption) (*GetAllPartyResponse, error)
	UpdateParty(ctx context.Context, in *GetPartyResponse, opts ...grpc.CallOption) (*VoidPartyResponse, error)
	DeleteParty(ctx context.Context, in *ByIdPartyRequest, opts ...grpc.CallOption) (*VoidPartyResponse, error)
}

type partyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartyServiceClient(cc grpc.ClientConnInterface) PartyServiceClient {
	return &partyServiceClient{cc}
}

func (c *partyServiceClient) CreateParty(ctx context.Context, in *CreatePartyRequest, opts ...grpc.CallOption) (*VoidPartyResponse, error) {
	out := new(VoidPartyResponse)
	err := c.cc.Invoke(ctx, "/voting.PartyService/CreateParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) GetPartyById(ctx context.Context, in *ByIdPartyRequest, opts ...grpc.CallOption) (*GetPartyResponse, error) {
	out := new(GetPartyResponse)
	err := c.cc.Invoke(ctx, "/voting.PartyService/GetPartyById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) GetAllParties(ctx context.Context, in *FilterPartyRequest, opts ...grpc.CallOption) (*GetAllPartyResponse, error) {
	out := new(GetAllPartyResponse)
	err := c.cc.Invoke(ctx, "/voting.PartyService/GetAllParties", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) UpdateParty(ctx context.Context, in *GetPartyResponse, opts ...grpc.CallOption) (*VoidPartyResponse, error) {
	out := new(VoidPartyResponse)
	err := c.cc.Invoke(ctx, "/voting.PartyService/UpdateParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partyServiceClient) DeleteParty(ctx context.Context, in *ByIdPartyRequest, opts ...grpc.CallOption) (*VoidPartyResponse, error) {
	out := new(VoidPartyResponse)
	err := c.cc.Invoke(ctx, "/voting.PartyService/DeleteParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartyServiceServer is the server API for PartyService service.
// All implementations must embed UnimplementedPartyServiceServer
// for forward compatibility
type PartyServiceServer interface {
	CreateParty(context.Context, *CreatePartyRequest) (*VoidPartyResponse, error)
	GetPartyById(context.Context, *ByIdPartyRequest) (*GetPartyResponse, error)
	GetAllParties(context.Context, *FilterPartyRequest) (*GetAllPartyResponse, error)
	UpdateParty(context.Context, *GetPartyResponse) (*VoidPartyResponse, error)
	DeleteParty(context.Context, *ByIdPartyRequest) (*VoidPartyResponse, error)
	mustEmbedUnimplementedPartyServiceServer()
}

// UnimplementedPartyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPartyServiceServer struct {
}

func (UnimplementedPartyServiceServer) CreateParty(context.Context, *CreatePartyRequest) (*VoidPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateParty not implemented")
}
func (UnimplementedPartyServiceServer) GetPartyById(context.Context, *ByIdPartyRequest) (*GetPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartyById not implemented")
}
func (UnimplementedPartyServiceServer) GetAllParties(context.Context, *FilterPartyRequest) (*GetAllPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllParties not implemented")
}
func (UnimplementedPartyServiceServer) UpdateParty(context.Context, *GetPartyResponse) (*VoidPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParty not implemented")
}
func (UnimplementedPartyServiceServer) DeleteParty(context.Context, *ByIdPartyRequest) (*VoidPartyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteParty not implemented")
}
func (UnimplementedPartyServiceServer) mustEmbedUnimplementedPartyServiceServer() {}

// UnsafePartyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartyServiceServer will
// result in compilation errors.
type UnsafePartyServiceServer interface {
	mustEmbedUnimplementedPartyServiceServer()
}

func RegisterPartyServiceServer(s grpc.ServiceRegistrar, srv PartyServiceServer) {
	s.RegisterService(&PartyService_ServiceDesc, srv)
}

func _PartyService_CreateParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).CreateParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting.PartyService/CreateParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).CreateParty(ctx, req.(*CreatePartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_GetPartyById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetPartyById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting.PartyService/GetPartyById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetPartyById(ctx, req.(*ByIdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_GetAllParties_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).GetAllParties(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting.PartyService/GetAllParties",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).GetAllParties(ctx, req.(*FilterPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_UpdateParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartyResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).UpdateParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting.PartyService/UpdateParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).UpdateParty(ctx, req.(*GetPartyResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartyService_DeleteParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByIdPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartyServiceServer).DeleteParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/voting.PartyService/DeleteParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartyServiceServer).DeleteParty(ctx, req.(*ByIdPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartyService_ServiceDesc is the grpc.ServiceDesc for PartyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "voting.PartyService",
	HandlerType: (*PartyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateParty",
			Handler:    _PartyService_CreateParty_Handler,
		},
		{
			MethodName: "GetPartyById",
			Handler:    _PartyService_GetPartyById_Handler,
		},
		{
			MethodName: "GetAllParties",
			Handler:    _PartyService_GetAllParties_Handler,
		},
		{
			MethodName: "UpdateParty",
			Handler:    _PartyService_UpdateParty_Handler,
		},
		{
			MethodName: "DeleteParty",
			Handler:    _PartyService_DeleteParty_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Party.proto",
}
