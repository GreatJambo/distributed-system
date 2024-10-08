// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: pet.proto

package pet

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
	PetService_RegisterNewPet_FullMethodName = "/pet.PetService/RegisterNewPet"
	PetService_SearchPet_FullMethodName      = "/pet.PetService/SearchPet"
)

// PetServiceClient is the client API for PetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetServiceClient interface {
	RegisterNewPet(ctx context.Context, in *RegisterNewPetRequest, opts ...grpc.CallOption) (*RegisterNewPetReply, error)
	SearchPet(ctx context.Context, in *SearchPetRequest, opts ...grpc.CallOption) (*SearchPetReply, error)
}

type petServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetServiceClient(cc grpc.ClientConnInterface) PetServiceClient {
	return &petServiceClient{cc}
}

func (c *petServiceClient) RegisterNewPet(ctx context.Context, in *RegisterNewPetRequest, opts ...grpc.CallOption) (*RegisterNewPetReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterNewPetReply)
	err := c.cc.Invoke(ctx, PetService_RegisterNewPet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) SearchPet(ctx context.Context, in *SearchPetRequest, opts ...grpc.CallOption) (*SearchPetReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchPetReply)
	err := c.cc.Invoke(ctx, PetService_SearchPet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetServiceServer is the server API for PetService service.
// All implementations must embed UnimplementedPetServiceServer
// for forward compatibility.
type PetServiceServer interface {
	RegisterNewPet(context.Context, *RegisterNewPetRequest) (*RegisterNewPetReply, error)
	SearchPet(context.Context, *SearchPetRequest) (*SearchPetReply, error)
	mustEmbedUnimplementedPetServiceServer()
}

// UnimplementedPetServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPetServiceServer struct{}

func (UnimplementedPetServiceServer) RegisterNewPet(context.Context, *RegisterNewPetRequest) (*RegisterNewPetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterNewPet not implemented")
}
func (UnimplementedPetServiceServer) SearchPet(context.Context, *SearchPetRequest) (*SearchPetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPet not implemented")
}
func (UnimplementedPetServiceServer) mustEmbedUnimplementedPetServiceServer() {}
func (UnimplementedPetServiceServer) testEmbeddedByValue()                    {}

// UnsafePetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetServiceServer will
// result in compilation errors.
type UnsafePetServiceServer interface {
	mustEmbedUnimplementedPetServiceServer()
}

func RegisterPetServiceServer(s grpc.ServiceRegistrar, srv PetServiceServer) {
	// If the following call pancis, it indicates UnimplementedPetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PetService_ServiceDesc, srv)
}

func _PetService_RegisterNewPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterNewPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).RegisterNewPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetService_RegisterNewPet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).RegisterNewPet(ctx, req.(*RegisterNewPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_SearchPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).SearchPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetService_SearchPet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).SearchPet(ctx, req.(*SearchPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PetService_ServiceDesc is the grpc.ServiceDesc for PetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pet.PetService",
	HandlerType: (*PetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterNewPet",
			Handler:    _PetService_RegisterNewPet_Handler,
		},
		{
			MethodName: "SearchPet",
			Handler:    _PetService_SearchPet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pet.proto",
}
