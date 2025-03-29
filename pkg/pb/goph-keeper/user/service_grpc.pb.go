// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0
// source: service.proto

package gophkeeperuser

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
	UsersService_RegisterUser_FullMethodName = "/goph_keeper.UsersService/RegisterUser"
	UsersService_Authenticate_FullMethodName = "/goph_keeper.UsersService/Authenticate"
)

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для работы с пользователями: регистрация, аутентификация
type UsersServiceClient interface {
	// Регистрация нового пользователя
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	// Аутентификация пользователя
	Authenticate(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error)
}

type usersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersServiceClient(cc grpc.ClientConnInterface) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, UsersService_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *usersServiceClient) Authenticate(ctx context.Context, in *AuthenticationRequest, opts ...grpc.CallOption) (*AuthenticationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthenticationResponse)
	err := c.cc.Invoke(ctx, UsersService_Authenticate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
// All implementations must embed UnimplementedUsersServiceServer
// for forward compatibility.
//
// Сервис для работы с пользователями: регистрация, аутентификация
type UsersServiceServer interface {
	// Регистрация нового пользователя
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	// Аутентификация пользователя
	Authenticate(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error)
	mustEmbedUnimplementedUsersServiceServer()
}

// UnimplementedUsersServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUsersServiceServer struct{}

func (UnimplementedUsersServiceServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUsersServiceServer) Authenticate(context.Context, *AuthenticationRequest) (*AuthenticationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedUsersServiceServer) mustEmbedUnimplementedUsersServiceServer() {}
func (UnimplementedUsersServiceServer) testEmbeddedByValue()                      {}

// UnsafeUsersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UsersServiceServer will
// result in compilation errors.
type UnsafeUsersServiceServer interface {
	mustEmbedUnimplementedUsersServiceServer()
}

func RegisterUsersServiceServer(s grpc.ServiceRegistrar, srv UsersServiceServer) {
	// If the following call pancis, it indicates UnimplementedUsersServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UsersService_ServiceDesc, srv)
}

func _UsersService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UsersService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UsersService_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).Authenticate(ctx, req.(*AuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UsersService_ServiceDesc is the grpc.ServiceDesc for UsersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UsersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goph_keeper.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UsersService_RegisterUser_Handler,
		},
		{
			MethodName: "Authenticate",
			Handler:    _UsersService_Authenticate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
