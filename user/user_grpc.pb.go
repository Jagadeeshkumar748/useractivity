// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: user/user.proto

package user

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

// UserActivityClient is the client API for UserActivity service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserActivityClient interface {
	AddUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error)
	AddActivity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error)
	QueryActivity(ctx context.Context, in *QueryActivityRequest, opts ...grpc.CallOption) (*QueryActivityResponse, error)
}

type userActivityClient struct {
	cc grpc.ClientConnInterface
}

func NewUserActivityClient(cc grpc.ClientConnInterface) UserActivityClient {
	return &userActivityClient{cc}
}

func (c *userActivityClient) AddUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/user.UserActivity/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userActivityClient) QueryUser(ctx context.Context, in *QueryUserRequest, opts ...grpc.CallOption) (*QueryUserResponse, error) {
	out := new(QueryUserResponse)
	err := c.cc.Invoke(ctx, "/user.UserActivity/QueryUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userActivityClient) AddActivity(ctx context.Context, in *ActivityRequest, opts ...grpc.CallOption) (*ActivityResponse, error) {
	out := new(ActivityResponse)
	err := c.cc.Invoke(ctx, "/user.UserActivity/AddActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userActivityClient) QueryActivity(ctx context.Context, in *QueryActivityRequest, opts ...grpc.CallOption) (*QueryActivityResponse, error) {
	out := new(QueryActivityResponse)
	err := c.cc.Invoke(ctx, "/user.UserActivity/QueryActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserActivityServer is the server API for UserActivity service.
// All implementations must embed UnimplementedUserActivityServer
// for forward compatibility
type UserActivityServer interface {
	AddUser(context.Context, *UserRequest) (*UserResponse, error)
	QueryUser(context.Context, *QueryUserRequest) (*QueryUserResponse, error)
	AddActivity(context.Context, *ActivityRequest) (*ActivityResponse, error)
	QueryActivity(context.Context, *QueryActivityRequest) (*QueryActivityResponse, error)
	mustEmbedUnimplementedUserActivityServer()
}

// UnimplementedUserActivityServer must be embedded to have forward compatible implementations.
type UnimplementedUserActivityServer struct {
}

func (UnimplementedUserActivityServer) AddUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedUserActivityServer) QueryUser(context.Context, *QueryUserRequest) (*QueryUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUser not implemented")
}
func (UnimplementedUserActivityServer) AddActivity(context.Context, *ActivityRequest) (*ActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddActivity not implemented")
}
func (UnimplementedUserActivityServer) QueryActivity(context.Context, *QueryActivityRequest) (*QueryActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryActivity not implemented")
}
func (UnimplementedUserActivityServer) mustEmbedUnimplementedUserActivityServer() {}

// UnsafeUserActivityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserActivityServer will
// result in compilation errors.
type UnsafeUserActivityServer interface {
	mustEmbedUnimplementedUserActivityServer()
}

func RegisterUserActivityServer(s grpc.ServiceRegistrar, srv UserActivityServer) {
	s.RegisterService(&UserActivity_ServiceDesc, srv)
}

func _UserActivity_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserActivityServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserActivity/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserActivityServer).AddUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserActivity_QueryUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserActivityServer).QueryUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserActivity/QueryUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserActivityServer).QueryUser(ctx, req.(*QueryUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserActivity_AddActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserActivityServer).AddActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserActivity/AddActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserActivityServer).AddActivity(ctx, req.(*ActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserActivity_QueryActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserActivityServer).QueryActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserActivity/QueryActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserActivityServer).QueryActivity(ctx, req.(*QueryActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserActivity_ServiceDesc is the grpc.ServiceDesc for UserActivity service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserActivity_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserActivity",
	HandlerType: (*UserActivityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _UserActivity_AddUser_Handler,
		},
		{
			MethodName: "QueryUser",
			Handler:    _UserActivity_QueryUser_Handler,
		},
		{
			MethodName: "AddActivity",
			Handler:    _UserActivity_AddActivity_Handler,
		},
		{
			MethodName: "QueryActivity",
			Handler:    _UserActivity_QueryActivity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}
