// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/service.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	api/service.proto

It has these top-level messages:
	AuthRequest
	AuthResponse
	GetUserRequest
	GetUserResponse
	User
	Payload
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type AuthRequest struct {
	Login string `protobuf:"bytes,1,opt,name=login" json:"login,omitempty"`
	Pass  string `protobuf:"bytes,2,opt,name=pass" json:"pass,omitempty"`
}

func (m *AuthRequest) Reset()                    { *m = AuthRequest{} }
func (m *AuthRequest) String() string            { return proto.CompactTextString(m) }
func (*AuthRequest) ProtoMessage()               {}
func (*AuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *AuthRequest) GetPass() string {
	if m != nil {
		return m.Pass
	}
	return ""
}

type AuthResponse struct {
}

func (m *AuthResponse) Reset()                    { *m = AuthResponse{} }
func (m *AuthResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()               {}
func (*AuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type GetUserRequest struct {
	Login string `protobuf:"bytes,1,opt,name=login" json:"login,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetUserRequest) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

type GetUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *GetUserResponse) Reset()                    { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()               {}
func (*GetUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type User struct {
	Login    string   `protobuf:"bytes,1,opt,name=login" json:"login,omitempty"`
	Pass     string   `protobuf:"bytes,2,opt,name=pass" json:"pass,omitempty"`
	Payload  *Payload `protobuf:"bytes,3,opt,name=payload" json:"payload,omitempty"`
	CreateAt string   `protobuf:"bytes,4,opt,name=create_at,json=createAt" json:"create_at,omitempty"`
	Active   bool     `protobuf:"varint,5,opt,name=active" json:"active,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *User) GetLogin() string {
	if m != nil {
		return m.Login
	}
	return ""
}

func (m *User) GetPass() string {
	if m != nil {
		return m.Pass
	}
	return ""
}

func (m *User) GetPayload() *Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *User) GetCreateAt() string {
	if m != nil {
		return m.CreateAt
	}
	return ""
}

func (m *User) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type Payload struct {
	CustomScript string `protobuf:"bytes,1,opt,name=custom_script,json=customScript" json:"custom_script,omitempty"`
	Url          string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Email        string `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
}

func (m *Payload) Reset()                    { *m = Payload{} }
func (m *Payload) String() string            { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()               {}
func (*Payload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Payload) GetCustomScript() string {
	if m != nil {
		return m.CustomScript
	}
	return ""
}

func (m *Payload) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Payload) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*AuthRequest)(nil), "auth.AuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "auth.AuthResponse")
	proto.RegisterType((*GetUserRequest)(nil), "auth.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "auth.GetUserResponse")
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*Payload)(nil), "auth.Payload")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthBasic service

type AuthBasicClient interface {
	Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type authBasicClient struct {
	cc *grpc.ClientConn
}

func NewAuthBasicClient(cc *grpc.ClientConn) AuthBasicClient {
	return &authBasicClient{cc}
}

func (c *authBasicClient) Auth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/auth.AuthBasic/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authBasicClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := grpc.Invoke(ctx, "/auth.AuthBasic/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthBasic service

type AuthBasicServer interface {
	Auth(context.Context, *AuthRequest) (*AuthResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
}

func RegisterAuthBasicServer(s *grpc.Server, srv AuthBasicServer) {
	s.RegisterService(&_AuthBasic_serviceDesc, srv)
}

func _AuthBasic_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthBasicServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthBasic/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthBasicServer).Auth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthBasic_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthBasicServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.AuthBasic/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthBasicServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthBasic_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthBasic",
	HandlerType: (*AuthBasicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _AuthBasic_Auth_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _AuthBasic_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service.proto",
}

func init() { proto.RegisterFile("api/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x25, 0x6d, 0xfa, 0x93, 0xdb, 0x9f, 0xef, 0xeb, 0xa5, 0xdf, 0x67, 0x6c, 0x55, 0x6a, 0x04,
	0x2d, 0x2e, 0x1a, 0x5a, 0x17, 0x82, 0xbb, 0xba, 0x71, 0xe3, 0x42, 0x22, 0x0a, 0xae, 0xca, 0x34,
	0x0e, 0xed, 0x40, 0x9a, 0x89, 0x99, 0x49, 0x41, 0xc4, 0x8d, 0x0f, 0xe0, 0xc6, 0x57, 0xf1, 0x4d,
	0x7c, 0x05, 0x1f, 0x44, 0x66, 0x26, 0x4a, 0xeb, 0x42, 0xdc, 0xcd, 0x39, 0xf7, 0xdc, 0xc3, 0xc9,
	0x3d, 0x81, 0x16, 0x49, 0x98, 0x2f, 0x68, 0xba, 0x64, 0x21, 0x1d, 0x24, 0x29, 0x97, 0x1c, 0x6d,
	0x92, 0xc9, 0x79, 0x67, 0x6b, 0xc6, 0xf9, 0x2c, 0xa2, 0xbe, 0x9a, 0x93, 0x38, 0xe6, 0x92, 0x48,
	0xc6, 0x63, 0x61, 0x34, 0xde, 0x31, 0xd4, 0xc6, 0x99, 0x9c, 0x07, 0xf4, 0x2e, 0xa3, 0x42, 0x62,
	0x1b, 0x4a, 0x11, 0x9f, 0xb1, 0xd8, 0xb5, 0x7a, 0x56, 0xdf, 0x09, 0x0c, 0x40, 0x04, 0x3b, 0x21,
	0x42, 0xb8, 0x05, 0x4d, 0xea, 0xb7, 0xd7, 0x84, 0xba, 0x59, 0x14, 0x09, 0x8f, 0x05, 0xf5, 0xf6,
	0xa1, 0x79, 0x46, 0xe5, 0x95, 0xa0, 0xe9, 0x8f, 0x5e, 0xde, 0x10, 0xfe, 0x7c, 0xe9, 0xcc, 0x2a,
	0xee, 0x80, 0x9d, 0x09, 0x9a, 0x6a, 0x5d, 0x6d, 0x04, 0x03, 0x15, 0x7b, 0xa0, 0x15, 0x9a, 0xf7,
	0x9e, 0x2d, 0xb0, 0x15, 0xfc, 0x7d, 0x3a, 0x3c, 0x80, 0x4a, 0x42, 0xee, 0x23, 0x4e, 0x6e, 0xdd,
	0xa2, 0x76, 0x6d, 0x18, 0xd7, 0x0b, 0x43, 0x06, 0x9f, 0x53, 0xec, 0x82, 0x13, 0xa6, 0x94, 0x48,
	0x3a, 0x21, 0xd2, 0xb5, 0xb5, 0x43, 0xd5, 0x10, 0x63, 0x89, 0xff, 0xa1, 0x4c, 0x42, 0xc9, 0x96,
	0xd4, 0x2d, 0xf5, 0xac, 0x7e, 0x35, 0xc8, 0x91, 0x77, 0x0d, 0x95, 0xdc, 0x08, 0xf7, 0xa0, 0x11,
	0x66, 0x42, 0xf2, 0xc5, 0x44, 0x84, 0x29, 0x4b, 0x64, 0x1e, 0xad, 0x6e, 0xc8, 0x4b, 0xcd, 0xe1,
	0x5f, 0x28, 0x66, 0x69, 0x94, 0x07, 0x54, 0x4f, 0xf5, 0x25, 0x74, 0x41, 0x58, 0xa4, 0xd3, 0x39,
	0x81, 0x01, 0xa3, 0x57, 0x0b, 0x1c, 0x75, 0xd4, 0x53, 0x22, 0x58, 0x88, 0xe7, 0x60, 0x2b, 0x80,
	0x2d, 0x13, 0x7d, 0xa5, 0xa6, 0x0e, 0xae, 0x52, 0x79, 0x01, 0xdb, 0x4f, 0x6f, 0xef, 0x2f, 0x85,
	0x0d, 0x0f, 0x75, 0xd3, 0xcb, 0xa1, 0xaf, 0x24, 0xfe, 0x54, 0x59, 0x9d, 0x58, 0x87, 0x78, 0x03,
	0x95, 0xfc, 0xee, 0xd8, 0x36, 0xdb, 0xeb, 0x75, 0x75, 0xfe, 0x7d, 0x63, 0x73, 0xdb, 0x5d, 0x6d,
	0xdb, 0xc5, 0xcd, 0x35, 0x5b, 0xd5, 0x8b, 0xff, 0xa0, 0xef, 0xff, 0x38, 0x2d, 0xeb, 0x5f, 0xe9,
	0xe8, 0x23, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xa5, 0xfb, 0x59, 0x83, 0x02, 0x00, 0x00,
}