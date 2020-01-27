// Code generated by protoc-gen-go. DO NOT EDIT.
// source: services/user/proto/user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Request struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f65f38375cad887, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Request) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Result struct {
	Correct              bool     `protobuf:"varint,1,opt,name=correct,proto3" json:"correct,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f65f38375cad887, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCorrect() bool {
	if m != nil {
		return m.Correct
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "user.Request")
	proto.RegisterType((*Result)(nil), "user.Result")
}

func init() { proto.RegisterFile("services/user/proto/user.proto", fileDescriptor_5f65f38375cad887) }

var fileDescriptor_5f65f38375cad887 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2f, 0x2d, 0x4e, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x07,
	0x33, 0xf5, 0xc0, 0x4c, 0x21, 0x16, 0x10, 0x5b, 0xc9, 0x91, 0x8b, 0x3d, 0x28, 0xb5, 0xb0, 0x34,
	0xb5, 0xb8, 0x44, 0x48, 0x8a, 0x8b, 0x03, 0x24, 0x94, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0,
	0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x83, 0xe4, 0x0a, 0x12, 0x8b, 0x8b, 0xcb, 0xf3, 0x8b, 0x52, 0x24,
	0x98, 0x20, 0x72, 0x30, 0xbe, 0x92, 0x12, 0x17, 0x5b, 0x50, 0x6a, 0x71, 0x69, 0x4e, 0x89, 0x90,
	0x04, 0x17, 0x7b, 0x72, 0x7e, 0x51, 0x51, 0x6a, 0x72, 0x09, 0xd8, 0x00, 0x8e, 0x20, 0x18, 0xd7,
	0xc8, 0x80, 0x8b, 0x25, 0xb4, 0x38, 0xb5, 0x48, 0x48, 0x83, 0x8b, 0xd3, 0x39, 0x23, 0x35, 0x39,
	0x1b, 0xcc, 0xe1, 0xd5, 0x03, 0x3b, 0x07, 0x6a, 0xbf, 0x14, 0x0f, 0x8c, 0x0b, 0x32, 0x2b, 0x89,
	0x0d, 0xec, 0x4a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x55, 0x12, 0xe5, 0xc7, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	// CheckUser returns whether the username and password are correct
	CheckUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) CheckUser(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/user.User/CheckUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	// CheckUser returns whether the username and password are correct
	CheckUser(context.Context, *Request) (*Result, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) CheckUser(ctx context.Context, req *Request) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_CheckUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CheckUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/CheckUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CheckUser(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUser",
			Handler:    _User_CheckUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/user/proto/user.proto",
}
