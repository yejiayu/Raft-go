// Code generated by protoc-gen-go. DO NOT EDIT.
// source: raft.proto

/*
Package internal is a generated protocol buffer package.

It is generated from these files:
	raft.proto

It has these top-level messages:
	VoteRequest
	VoteResponse
	AppendRequest
	AppendResponse
*/
package internal

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

type VoteRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *VoteRequest) Reset()                    { *m = VoteRequest{} }
func (m *VoteRequest) String() string            { return proto.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()               {}
func (*VoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VoteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type VoteResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *VoteResponse) Reset()                    { *m = VoteResponse{} }
func (m *VoteResponse) String() string            { return proto.CompactTextString(m) }
func (*VoteResponse) ProtoMessage()               {}
func (*VoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VoteResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AppendRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *AppendRequest) Reset()                    { *m = AppendRequest{} }
func (m *AppendRequest) String() string            { return proto.CompactTextString(m) }
func (*AppendRequest) ProtoMessage()               {}
func (*AppendRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AppendRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AppendResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *AppendResponse) Reset()                    { *m = AppendResponse{} }
func (m *AppendResponse) String() string            { return proto.CompactTextString(m) }
func (*AppendResponse) ProtoMessage()               {}
func (*AppendResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AppendResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*VoteRequest)(nil), "internal.VoteRequest")
	proto.RegisterType((*VoteResponse)(nil), "internal.VoteResponse")
	proto.RegisterType((*AppendRequest)(nil), "internal.AppendRequest")
	proto.RegisterType((*AppendResponse)(nil), "internal.AppendResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Raft service

type RaftClient interface {
	RequestVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	AppendEntries(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendResponse, error)
}

type raftClient struct {
	cc *grpc.ClientConn
}

func NewRaftClient(cc *grpc.ClientConn) RaftClient {
	return &raftClient{cc}
}

func (c *raftClient) RequestVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	out := new(VoteResponse)
	err := grpc.Invoke(ctx, "/internal.Raft/RequestVote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftClient) AppendEntries(ctx context.Context, in *AppendRequest, opts ...grpc.CallOption) (*AppendResponse, error) {
	out := new(AppendResponse)
	err := grpc.Invoke(ctx, "/internal.Raft/AppendEntries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Raft service

type RaftServer interface {
	RequestVote(context.Context, *VoteRequest) (*VoteResponse, error)
	AppendEntries(context.Context, *AppendRequest) (*AppendResponse, error)
}

func RegisterRaftServer(s *grpc.Server, srv RaftServer) {
	s.RegisterService(&_Raft_serviceDesc, srv)
}

func _Raft_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Raft/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).RequestVote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Raft_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Raft/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftServer).AppendEntries(ctx, req.(*AppendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Raft_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal.Raft",
	HandlerType: (*RaftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestVote",
			Handler:    _Raft_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _Raft_AppendEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft.proto",
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4a, 0x4c, 0x2b,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0xcc, 0x2b, 0x49, 0x2d, 0xca, 0x4b, 0xcc,
	0x51, 0x52, 0xe4, 0xe2, 0x0e, 0xcb, 0x2f, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x95, 0x34, 0xb8, 0x78, 0x20, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x24, 0xb8, 0xd8,
	0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3, 0x61, 0xca, 0x60, 0x5c, 0x25, 0x65, 0x2e, 0x5e, 0xc7, 0x82,
	0x82, 0xd4, 0xbc, 0x14, 0x7c, 0xc6, 0x69, 0x71, 0xf1, 0xc1, 0x14, 0x11, 0x32, 0xd0, 0xa8, 0x83,
	0x91, 0x8b, 0x25, 0x28, 0x31, 0xad, 0x44, 0xc8, 0x86, 0x8b, 0x1b, 0x6a, 0x26, 0xc8, 0x29, 0x42,
	0xa2, 0x7a, 0x30, 0x0f, 0xe8, 0x21, 0xb9, 0x5e, 0x4a, 0x0c, 0x5d, 0x18, 0x6a, 0x81, 0x13, 0xcc,
	0x5d, 0xae, 0x79, 0x25, 0x45, 0x99, 0xa9, 0xc5, 0x42, 0xe2, 0x08, 0x85, 0x28, 0x0e, 0x96, 0x92,
	0xc0, 0x94, 0x80, 0x98, 0x91, 0xc4, 0x06, 0x0e, 0x39, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x41, 0xf9, 0xf5, 0xd0, 0x47, 0x01, 0x00, 0x00,
}
