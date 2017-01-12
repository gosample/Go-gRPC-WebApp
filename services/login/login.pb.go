// Code generated by protoc-gen-go.
// source: services/login/login.proto
// DO NOT EDIT!

/*
Package services is a generated protocol buffer package.

It is generated from these files:
	services/login/login.proto

It has these top-level messages:
*/
package services

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google/api"
import user "stars-app/messages/user"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for LoginService service

type LoginServiceClient interface {
	Login(ctx context.Context, in *user.User, opts ...grpc.CallOption) (*user.User, error)
}

type loginServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoginServiceClient(cc *grpc.ClientConn) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *user.User, opts ...grpc.CallOption) (*user.User, error) {
	out := new(user.User)
	err := grpc.Invoke(ctx, "/services.LoginService/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for LoginService service

type LoginServiceServer interface {
	Login(context.Context, *user.User) (*user.User, error)
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(user.User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*user.User))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/login/login.proto",
}

func init() { proto.RegisterFile("services/login/login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0xcf, 0xc9, 0x4f, 0xcf, 0xcc, 0x83, 0x90, 0x7a, 0x05, 0x45, 0xf9,
	0x25, 0xf9, 0x42, 0x1c, 0x30, 0x39, 0x29, 0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4,
	0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x3a,
	0x29, 0xa5, 0xe2, 0x92, 0xc4, 0xa2, 0x62, 0xdd, 0xc4, 0x82, 0x02, 0xfd, 0xdc, 0xd4, 0xe2, 0xe2,
	0xc4, 0xf4, 0xd4, 0x62, 0xfd, 0xd2, 0xe2, 0xd4, 0x22, 0x30, 0x01, 0x51, 0x63, 0xe4, 0xca, 0xc5,
	0xe3, 0x03, 0x32, 0x3a, 0x18, 0x62, 0xa4, 0x90, 0x29, 0x17, 0x2b, 0x98, 0x2f, 0xc4, 0xa5, 0x07,
	0x56, 0x15, 0x5a, 0x9c, 0x5a, 0x24, 0x85, 0xc4, 0x56, 0x12, 0x69, 0xba, 0xfc, 0x64, 0x32, 0x13,
	0x9f, 0x12, 0xa7, 0x7e, 0x99, 0x21, 0xc4, 0x55, 0x56, 0x8c, 0x5a, 0x49, 0x6c, 0x60, 0xd3, 0x8c,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2d, 0xde, 0x3e, 0x32, 0xb7, 0x00, 0x00, 0x00,
}
