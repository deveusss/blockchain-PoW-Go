// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cli.proto

/*
Package cli is a generated protocol buffer package.

It is generated from these files:
	cli.proto

It has these top-level messages:
	GetBalanceRequest
	GetBalanceResponse
	Balance
	SendRequest
	SendResponse
*/
package cli

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

type GetBalanceRequest struct {
}

func (m *GetBalanceRequest) Reset()                    { *m = GetBalanceRequest{} }
func (m *GetBalanceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetBalanceRequest) ProtoMessage()               {}
func (*GetBalanceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetBalanceResponse struct {
	Addresses []*Balance `protobuf:"bytes,1,rep,name=addresses" json:"addresses,omitempty"`
}

func (m *GetBalanceResponse) Reset()                    { *m = GetBalanceResponse{} }
func (m *GetBalanceResponse) String() string            { return proto.CompactTextString(m) }
func (*GetBalanceResponse) ProtoMessage()               {}
func (*GetBalanceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetBalanceResponse) GetAddresses() []*Balance {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type Balance struct {
	Address string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Balance int64  `protobuf:"varint,2,opt,name=balance" json:"balance,omitempty"`
}

func (m *Balance) Reset()                    { *m = Balance{} }
func (m *Balance) String() string            { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()               {}
func (*Balance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Balance) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Balance) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

type SendRequest struct {
	ToAddress string `protobuf:"bytes,1,opt,name=to_address,json=toAddress" json:"to_address,omitempty"`
	Amount    int64  `protobuf:"varint,2,opt,name=amount" json:"amount,omitempty"`
}

func (m *SendRequest) Reset()                    { *m = SendRequest{} }
func (m *SendRequest) String() string            { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()               {}
func (*SendRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SendRequest) GetToAddress() string {
	if m != nil {
		return m.ToAddress
	}
	return ""
}

func (m *SendRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type SendResponse struct {
	TxId string `protobuf:"bytes,1,opt,name=tx_id,json=txId" json:"tx_id,omitempty"`
}

func (m *SendResponse) Reset()                    { *m = SendResponse{} }
func (m *SendResponse) String() string            { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()               {}
func (*SendResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SendResponse) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBalanceRequest)(nil), "cli.GetBalanceRequest")
	proto.RegisterType((*GetBalanceResponse)(nil), "cli.GetBalanceResponse")
	proto.RegisterType((*Balance)(nil), "cli.Balance")
	proto.RegisterType((*SendRequest)(nil), "cli.SendRequest")
	proto.RegisterType((*SendResponse)(nil), "cli.SendResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CliService service

type CliServiceClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error)
	GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error)
}

type cliServiceClient struct {
	cc *grpc.ClientConn
}

func NewCliServiceClient(cc *grpc.ClientConn) CliServiceClient {
	return &cliServiceClient{cc}
}

func (c *cliServiceClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := grpc.Invoke(ctx, "/cli.CliService/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliServiceClient) GetBalance(ctx context.Context, in *GetBalanceRequest, opts ...grpc.CallOption) (*GetBalanceResponse, error) {
	out := new(GetBalanceResponse)
	err := grpc.Invoke(ctx, "/cli.CliService/GetBalance", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CliService service

type CliServiceServer interface {
	Send(context.Context, *SendRequest) (*SendResponse, error)
	GetBalance(context.Context, *GetBalanceRequest) (*GetBalanceResponse, error)
}

func RegisterCliServiceServer(s *grpc.Server, srv CliServiceServer) {
	s.RegisterService(&_CliService_serviceDesc, srv)
}

func _CliService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cli.CliService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliService_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServiceServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cli.CliService/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServiceServer).GetBalance(ctx, req.(*GetBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CliService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cli.CliService",
	HandlerType: (*CliServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _CliService_Send_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _CliService_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cli.proto",
}

func init() { proto.RegisterFile("cli.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x4d, 0x6b, 0xc2, 0x40,
	0x10, 0x86, 0x9b, 0xc6, 0x2a, 0x79, 0xf5, 0x50, 0x47, 0xb0, 0x41, 0x28, 0x84, 0xed, 0x25, 0x14,
	0xea, 0xc1, 0x9e, 0x4b, 0x3f, 0xa1, 0xf4, 0x1a, 0x7f, 0x80, 0xc4, 0xec, 0x1c, 0x16, 0xd2, 0xac,
	0xcd, 0xae, 0xc5, 0x43, 0x7f, 0x7c, 0x49, 0xb2, 0x8b, 0x11, 0x8f, 0xf3, 0xcc, 0xbb, 0x0f, 0x33,
	0xb3, 0x88, 0x8a, 0x52, 0x2d, 0x77, 0xb5, 0xb6, 0x9a, 0xc2, 0xa2, 0x54, 0x62, 0x86, 0xe9, 0x27,
	0xdb, 0xb7, 0xbc, 0xcc, 0xab, 0x82, 0x33, 0xfe, 0xd9, 0xb3, 0xb1, 0xe2, 0x05, 0xd4, 0x87, 0x66,
	0xa7, 0x2b, 0xc3, 0x74, 0x8f, 0x28, 0x97, 0xb2, 0x66, 0x63, 0xd8, 0xc4, 0x41, 0x12, 0xa6, 0xe3,
	0xd5, 0x64, 0xd9, 0xe8, 0x7c, 0xf0, 0xd8, 0x16, 0x4f, 0x18, 0x39, 0x4a, 0x31, 0x46, 0x8e, 0xc7,
	0x41, 0x12, 0xa4, 0x51, 0xe6, 0xcb, 0xa6, 0xb3, 0xed, 0x42, 0xf1, 0x65, 0x12, 0xa4, 0x61, 0xe6,
	0x4b, 0xf1, 0x81, 0xf1, 0x9a, 0x2b, 0xe9, 0xe6, 0xa1, 0x5b, 0xc0, 0xea, 0xcd, 0xa9, 0x25, 0xb2,
	0xfa, 0xd5, 0x79, 0xe6, 0x18, 0xe6, 0xdf, 0x7a, 0x5f, 0x59, 0xa7, 0x71, 0x95, 0xb8, 0xc3, 0xa4,
	0xb3, 0xb8, 0x05, 0x66, 0xb8, 0xb2, 0x87, 0x8d, 0x92, 0xce, 0x30, 0xb0, 0x87, 0x2f, 0xb9, 0xfa,
	0x03, 0xde, 0x4b, 0xb5, 0xe6, 0xfa, 0x57, 0x15, 0x4c, 0x0f, 0x18, 0x34, 0x4f, 0xe8, 0xba, 0x5d,
	0xac, 0x37, 0xc3, 0x62, 0xda, 0x23, 0x9d, 0x4f, 0x5c, 0xd0, 0x33, 0x70, 0x3c, 0x14, 0xcd, 0xdb,
	0xc8, 0xd9, 0x39, 0x17, 0x37, 0x67, 0xdc, 0x0b, 0xb6, 0xc3, 0xf6, 0x2b, 0x1e, 0xff, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x8a, 0xf1, 0x7a, 0xea, 0x97, 0x01, 0x00, 0x00,
}