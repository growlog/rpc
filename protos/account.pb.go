// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/account.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type LoginRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b24b10312e6348, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b24b10312e6348, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *LoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyAccessTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyAccessTokenRequest) Reset()         { *m = VerifyAccessTokenRequest{} }
func (m *VerifyAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*VerifyAccessTokenRequest) ProtoMessage()    {}
func (*VerifyAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b24b10312e6348, []int{2}
}

func (m *VerifyAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyAccessTokenRequest.Unmarshal(m, b)
}
func (m *VerifyAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyAccessTokenRequest.Marshal(b, m, deterministic)
}
func (m *VerifyAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyAccessTokenRequest.Merge(m, src)
}
func (m *VerifyAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyAccessTokenRequest.Size(m)
}
func (m *VerifyAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyAccessTokenRequest proto.InternalMessageInfo

func (m *VerifyAccessTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type VerifyAccessTokenResponse struct {
	Status               bool                 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserId               int64                `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ThingId              int64                `protobuf:"varint,3,opt,name=thing_id,json=thingId,proto3" json:"thing_id,omitempty"`
	ExpiresAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *VerifyAccessTokenResponse) Reset()         { *m = VerifyAccessTokenResponse{} }
func (m *VerifyAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*VerifyAccessTokenResponse) ProtoMessage()    {}
func (*VerifyAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b24b10312e6348, []int{3}
}

func (m *VerifyAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyAccessTokenResponse.Unmarshal(m, b)
}
func (m *VerifyAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyAccessTokenResponse.Marshal(b, m, deterministic)
}
func (m *VerifyAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyAccessTokenResponse.Merge(m, src)
}
func (m *VerifyAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyAccessTokenResponse.Size(m)
}
func (m *VerifyAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyAccessTokenResponse proto.InternalMessageInfo

func (m *VerifyAccessTokenResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *VerifyAccessTokenResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *VerifyAccessTokenResponse) GetThingId() int64 {
	if m != nil {
		return m.ThingId
	}
	return 0
}

func (m *VerifyAccessTokenResponse) GetExpiresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

type RefreshAccessTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RefreshAccessTokenRequest) Reset()         { *m = RefreshAccessTokenRequest{} }
func (m *RefreshAccessTokenRequest) String() string { return proto.CompactTextString(m) }
func (*RefreshAccessTokenRequest) ProtoMessage()    {}
func (*RefreshAccessTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b24b10312e6348, []int{4}
}

func (m *RefreshAccessTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshAccessTokenRequest.Unmarshal(m, b)
}
func (m *RefreshAccessTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshAccessTokenRequest.Marshal(b, m, deterministic)
}
func (m *RefreshAccessTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshAccessTokenRequest.Merge(m, src)
}
func (m *RefreshAccessTokenRequest) XXX_Size() int {
	return xxx_messageInfo_RefreshAccessTokenRequest.Size(m)
}
func (m *RefreshAccessTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshAccessTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshAccessTokenRequest proto.InternalMessageInfo

func (m *RefreshAccessTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RefreshAccessTokenResponse struct {
	Status               bool                 `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	UserId               int64                `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ThingId              int64                `protobuf:"varint,3,opt,name=thing_id,json=thingId,proto3" json:"thing_id,omitempty"`
	ExpiresAt            *timestamp.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RefreshAccessTokenResponse) Reset()         { *m = RefreshAccessTokenResponse{} }
func (m *RefreshAccessTokenResponse) String() string { return proto.CompactTextString(m) }
func (*RefreshAccessTokenResponse) ProtoMessage()    {}
func (*RefreshAccessTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b24b10312e6348, []int{5}
}

func (m *RefreshAccessTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RefreshAccessTokenResponse.Unmarshal(m, b)
}
func (m *RefreshAccessTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RefreshAccessTokenResponse.Marshal(b, m, deterministic)
}
func (m *RefreshAccessTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RefreshAccessTokenResponse.Merge(m, src)
}
func (m *RefreshAccessTokenResponse) XXX_Size() int {
	return xxx_messageInfo_RefreshAccessTokenResponse.Size(m)
}
func (m *RefreshAccessTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RefreshAccessTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RefreshAccessTokenResponse proto.InternalMessageInfo

func (m *RefreshAccessTokenResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *RefreshAccessTokenResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RefreshAccessTokenResponse) GetThingId() int64 {
	if m != nil {
		return m.ThingId
	}
	return 0
}

func (m *RefreshAccessTokenResponse) GetExpiresAt() *timestamp.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "protos.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "protos.LoginResponse")
	proto.RegisterType((*VerifyAccessTokenRequest)(nil), "protos.VerifyAccessTokenRequest")
	proto.RegisterType((*VerifyAccessTokenResponse)(nil), "protos.VerifyAccessTokenResponse")
	proto.RegisterType((*RefreshAccessTokenRequest)(nil), "protos.RefreshAccessTokenRequest")
	proto.RegisterType((*RefreshAccessTokenResponse)(nil), "protos.RefreshAccessTokenResponse")
}

func init() { proto.RegisterFile("protos/account.proto", fileDescriptor_82b24b10312e6348) }

var fileDescriptor_82b24b10312e6348 = []byte{
	// 353 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0xb5, 0x22, 0x5f, 0xa3, 0x1e, 0xdc, 0xa0, 0x96, 0xbd, 0x88, 0x3d, 0x71, 0x2a, 0x8a, 0x89,
	0x89, 0x07, 0x13, 0x39, 0x92, 0x78, 0x6a, 0x88, 0x07, 0x13, 0x43, 0x4a, 0x19, 0xca, 0x46, 0xe8,
	0xd6, 0xce, 0x36, 0xea, 0xdf, 0xd1, 0xbf, 0xe9, 0xc1, 0x74, 0xb7, 0x25, 0x44, 0x3e, 0xf4, 0xe8,
	0xf1, 0xcd, 0xbc, 0xf7, 0xda, 0x79, 0x33, 0x0b, 0x8d, 0x38, 0x91, 0x4a, 0x52, 0xc7, 0x0f, 0x02,
	0x99, 0x46, 0xca, 0xd5, 0x90, 0x55, 0x4c, 0x95, 0x9f, 0x85, 0x52, 0x86, 0x33, 0xec, 0x68, 0x38,
	0x4a, 0x27, 0x1d, 0x25, 0xe6, 0x48, 0xca, 0x9f, 0xc7, 0x86, 0xe8, 0xdc, 0xc1, 0xc1, 0xbd, 0x0c,
	0x45, 0xe4, 0xe1, 0x4b, 0x8a, 0xa4, 0x58, 0x03, 0xca, 0x38, 0xf7, 0xc5, 0xcc, 0xb6, 0x5a, 0x56,
	0xbb, 0xee, 0x19, 0xc0, 0x38, 0xd4, 0x62, 0x9f, 0xe8, 0x55, 0x26, 0x63, 0x7b, 0x57, 0x37, 0x16,
	0xd8, 0xb9, 0x85, 0xc3, 0xdc, 0x81, 0x62, 0x19, 0x11, 0xb2, 0x13, 0xa8, 0x90, 0xf2, 0x55, 0x4a,
	0xda, 0xa3, 0xe6, 0xe5, 0x28, 0xb3, 0x56, 0xf2, 0x19, 0xa3, 0xdc, 0xc1, 0x00, 0xe7, 0x02, 0xec,
	0x07, 0x4c, 0xc4, 0xe4, 0xbd, 0x17, 0x04, 0x48, 0x34, 0xc8, 0x8a, 0x4b, 0x3f, 0x63, 0x14, 0xd6,
	0xb2, 0xe2, 0xc3, 0x82, 0xe6, 0x1a, 0xc9, 0x2f, 0x5f, 0x3f, 0x85, 0x6a, 0x4a, 0x98, 0x0c, 0x85,
	0x99, 0xa0, 0xe4, 0x55, 0x32, 0xd8, 0x1f, 0xb3, 0x26, 0xd4, 0xd4, 0x54, 0x44, 0x61, 0xd6, 0x29,
	0xe9, 0x4e, 0x55, 0xe3, 0xfe, 0x98, 0xdd, 0x00, 0xe0, 0x5b, 0x2c, 0x12, 0xa4, 0xa1, 0xaf, 0xec,
	0xbd, 0x96, 0xd5, 0xde, 0xef, 0x72, 0xd7, 0x44, 0xea, 0x16, 0x91, 0xba, 0x83, 0x22, 0x52, 0xaf,
	0x9e, 0xb3, 0x7b, 0xca, 0xb9, 0x84, 0xa6, 0x87, 0x93, 0x04, 0x69, 0xfa, 0xe7, 0xb9, 0x3e, 0x2d,
	0xe0, 0xeb, 0x34, 0xff, 0x6a, 0xb0, 0xee, 0x97, 0x05, 0xd5, 0x9e, 0xb9, 0x35, 0x76, 0x0d, 0x65,
	0xbd, 0x7a, 0xd6, 0x30, 0x22, 0x72, 0x97, 0x6f, 0x89, 0x1f, 0xff, 0xa8, 0x9a, 0x41, 0x9c, 0x1d,
	0xf6, 0x08, 0x47, 0x2b, 0x0b, 0x64, 0xad, 0x82, 0xbd, 0xe9, 0x1c, 0xf8, 0xf9, 0x16, 0xc6, 0xc2,
	0xfb, 0x09, 0xd8, 0x6a, 0x88, 0x6c, 0x21, 0xdd, 0xb8, 0x14, 0xee, 0x6c, 0xa3, 0x14, 0xf6, 0x23,
	0xf3, 0xb0, 0xae, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xe0, 0xcf, 0x6a, 0x77, 0x03, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	VerifyAccessToken(ctx context.Context, in *VerifyAccessTokenRequest, opts ...grpc.CallOption) (*VerifyAccessTokenResponse, error)
	RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/protos.Account/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) VerifyAccessToken(ctx context.Context, in *VerifyAccessTokenRequest, opts ...grpc.CallOption) (*VerifyAccessTokenResponse, error) {
	out := new(VerifyAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/protos.Account/VerifyAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) RefreshAccessToken(ctx context.Context, in *RefreshAccessTokenRequest, opts ...grpc.CallOption) (*RefreshAccessTokenResponse, error) {
	out := new(RefreshAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/protos.Account/RefreshAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	VerifyAccessToken(context.Context, *VerifyAccessTokenRequest) (*VerifyAccessTokenResponse, error)
	RefreshAccessToken(context.Context, *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (*UnimplementedAccountServer) VerifyAccessToken(ctx context.Context, req *VerifyAccessTokenRequest) (*VerifyAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyAccessToken not implemented")
}
func (*UnimplementedAccountServer) RefreshAccessToken(ctx context.Context, req *RefreshAccessTokenRequest) (*RefreshAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAccessToken not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Account/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_VerifyAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).VerifyAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Account/VerifyAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).VerifyAccessToken(ctx, req.(*VerifyAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_RefreshAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).RefreshAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.Account/RefreshAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).RefreshAccessToken(ctx, req.(*RefreshAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Account_Login_Handler,
		},
		{
			MethodName: "VerifyAccessToken",
			Handler:    _Account_VerifyAccessToken_Handler,
		},
		{
			MethodName: "RefreshAccessToken",
			Handler:    _Account_RefreshAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/account.proto",
}
