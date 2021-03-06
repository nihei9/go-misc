// Code generated by protoc-gen-go. DO NOT EDIT.
// source: simple.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	simple.proto

It has these top-level messages:
	Message
	Result
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Result_StatusCode int32

const (
	Result_OK    Result_StatusCode = 0
	Result_ERROR Result_StatusCode = 1
)

var Result_StatusCode_name = map[int32]string{
	0: "OK",
	1: "ERROR",
}
var Result_StatusCode_value = map[string]int32{
	"OK":    0,
	"ERROR": 1,
}

func (x Result_StatusCode) String() string {
	return proto.EnumName(Result_StatusCode_name, int32(x))
}
func (Result_StatusCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Message struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Data string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type Result struct {
	Status  Result_StatusCode `protobuf:"varint,1,opt,name=status,enum=pb.Result_StatusCode" json:"status,omitempty"`
	Message string            `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *Result) Reset()                    { *m = Result{} }
func (m *Result) String() string            { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()               {}
func (*Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Result) GetStatus() Result_StatusCode {
	if m != nil {
		return m.Status
	}
	return Result_OK
}

func (m *Result) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Message)(nil), "pb.Message")
	proto.RegisterType((*Result)(nil), "pb.Result")
	proto.RegisterEnum("pb.Result_StatusCode", Result_StatusCode_name, Result_StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for MessageProcessor service

type MessageProcessorClient interface {
	Do(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Result, error)
}

type messageProcessorClient struct {
	cc *grpc.ClientConn
}

func NewMessageProcessorClient(cc *grpc.ClientConn) MessageProcessorClient {
	return &messageProcessorClient{cc}
}

func (c *messageProcessorClient) Do(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := grpc.Invoke(ctx, "/pb.MessageProcessor/Do", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for MessageProcessor service

type MessageProcessorServer interface {
	Do(context.Context, *Message) (*Result, error)
}

func RegisterMessageProcessorServer(s *grpc.Server, srv MessageProcessorServer) {
	s.RegisterService(&_MessageProcessor_serviceDesc, srv)
}

func _MessageProcessor_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageProcessorServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageProcessor/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageProcessorServer).Do(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageProcessor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageProcessor",
	HandlerType: (*MessageProcessorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _MessageProcessor_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simple.proto",
}

func init() { proto.RegisterFile("simple.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x8f, 0xc1, 0x4a, 0xc0, 0x30,
	0x0c, 0x86, 0xb7, 0xa2, 0x1d, 0x8b, 0x32, 0x46, 0x40, 0x18, 0x5e, 0x26, 0x3d, 0x79, 0x59, 0x0f,
	0xdb, 0x23, 0xa8, 0x27, 0x91, 0x49, 0x7d, 0x82, 0xce, 0x16, 0x19, 0x6c, 0xb4, 0x34, 0xdd, 0xfb,
	0x8b, 0x5d, 0x65, 0xb7, 0x24, 0xdf, 0x9f, 0xfc, 0x7f, 0xe0, 0x9e, 0xd6, 0xdd, 0x6f, 0x56, 0xfa,
	0xe0, 0xa2, 0x43, 0xe6, 0x17, 0x31, 0x40, 0xf5, 0x61, 0x89, 0xf4, 0x8f, 0xc5, 0x06, 0xd8, 0x6a,
	0xba, 0xf2, 0xa9, 0x7c, 0xae, 0x15, 0x5b, 0x0d, 0x22, 0xdc, 0x18, 0x1d, 0x75, 0xc7, 0xd2, 0x24,
	0xd5, 0x22, 0x00, 0x57, 0x96, 0x8e, 0x2d, 0xe2, 0x00, 0x9c, 0xa2, 0x8e, 0x07, 0xa5, 0x8d, 0x66,
	0x7c, 0x90, 0x7e, 0x91, 0x27, 0x93, 0x5f, 0x09, 0xbc, 0x38, 0x63, 0x55, 0x16, 0x61, 0x07, 0xd5,
	0x7e, 0xfa, 0xe4, 0x7b, 0xff, 0xad, 0xe8, 0x01, 0x2e, 0x3d, 0x72, 0x60, 0xf3, 0x7b, 0x5b, 0x60,
	0x0d, 0xb7, 0x6f, 0x4a, 0xcd, 0xaa, 0x2d, 0xc7, 0x09, 0xda, 0x1c, 0xf1, 0x33, 0xb8, 0x6f, 0x4b,
	0xe4, 0x02, 0xf6, 0xc0, 0x5e, 0x1d, 0xde, 0xfd, 0x79, 0x66, 0xf6, 0x08, 0x57, 0x00, 0x51, 0x2c,
	0x3c, 0xbd, 0x38, 0xfd, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xfb, 0x70, 0xab, 0xf2, 0x00, 0x00,
	0x00,
}
