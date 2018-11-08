// Code generated by protoc-gen-go. DO NOT EDIT.
// source: raft.proto

package raftdemo

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

type AppendEntriesRequest struct {
	Term                 uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Leader               []byte   `protobuf:"bytes,2,opt,name=leader,proto3" json:"leader,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendEntriesRequest) Reset()         { *m = AppendEntriesRequest{} }
func (m *AppendEntriesRequest) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesRequest) ProtoMessage()    {}
func (*AppendEntriesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_f30d339453f41f8d, []int{0}
}
func (m *AppendEntriesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntriesRequest.Unmarshal(m, b)
}
func (m *AppendEntriesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntriesRequest.Marshal(b, m, deterministic)
}
func (dst *AppendEntriesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesRequest.Merge(dst, src)
}
func (m *AppendEntriesRequest) XXX_Size() int {
	return xxx_messageInfo_AppendEntriesRequest.Size(m)
}
func (m *AppendEntriesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesRequest proto.InternalMessageInfo

func (m *AppendEntriesRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesRequest) GetLeader() []byte {
	if m != nil {
		return m.Leader
	}
	return nil
}

type AppendEntriesResponse struct {
	Term                 uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppendEntriesResponse) Reset()         { *m = AppendEntriesResponse{} }
func (m *AppendEntriesResponse) String() string { return proto.CompactTextString(m) }
func (*AppendEntriesResponse) ProtoMessage()    {}
func (*AppendEntriesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_f30d339453f41f8d, []int{1}
}
func (m *AppendEntriesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppendEntriesResponse.Unmarshal(m, b)
}
func (m *AppendEntriesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppendEntriesResponse.Marshal(b, m, deterministic)
}
func (dst *AppendEntriesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppendEntriesResponse.Merge(dst, src)
}
func (m *AppendEntriesResponse) XXX_Size() int {
	return xxx_messageInfo_AppendEntriesResponse.Size(m)
}
func (m *AppendEntriesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AppendEntriesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AppendEntriesResponse proto.InternalMessageInfo

func (m *AppendEntriesResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *AppendEntriesResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RequestVoteRequest struct {
	Term                 uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Candidate            []byte   `protobuf:"bytes,2,opt,name=candidate,proto3" json:"candidate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVoteRequest) Reset()         { *m = RequestVoteRequest{} }
func (m *RequestVoteRequest) String() string { return proto.CompactTextString(m) }
func (*RequestVoteRequest) ProtoMessage()    {}
func (*RequestVoteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_f30d339453f41f8d, []int{2}
}
func (m *RequestVoteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVoteRequest.Unmarshal(m, b)
}
func (m *RequestVoteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVoteRequest.Marshal(b, m, deterministic)
}
func (dst *RequestVoteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVoteRequest.Merge(dst, src)
}
func (m *RequestVoteRequest) XXX_Size() int {
	return xxx_messageInfo_RequestVoteRequest.Size(m)
}
func (m *RequestVoteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVoteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVoteRequest proto.InternalMessageInfo

func (m *RequestVoteRequest) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteRequest) GetCandidate() []byte {
	if m != nil {
		return m.Candidate
	}
	return nil
}

type RequestVoteResponse struct {
	Term                 uint64   `protobuf:"varint,1,opt,name=term,proto3" json:"term,omitempty"`
	Granted              bool     `protobuf:"varint,2,opt,name=granted,proto3" json:"granted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestVoteResponse) Reset()         { *m = RequestVoteResponse{} }
func (m *RequestVoteResponse) String() string { return proto.CompactTextString(m) }
func (*RequestVoteResponse) ProtoMessage()    {}
func (*RequestVoteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_raft_f30d339453f41f8d, []int{3}
}
func (m *RequestVoteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestVoteResponse.Unmarshal(m, b)
}
func (m *RequestVoteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestVoteResponse.Marshal(b, m, deterministic)
}
func (dst *RequestVoteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestVoteResponse.Merge(dst, src)
}
func (m *RequestVoteResponse) XXX_Size() int {
	return xxx_messageInfo_RequestVoteResponse.Size(m)
}
func (m *RequestVoteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestVoteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RequestVoteResponse proto.InternalMessageInfo

func (m *RequestVoteResponse) GetTerm() uint64 {
	if m != nil {
		return m.Term
	}
	return 0
}

func (m *RequestVoteResponse) GetGranted() bool {
	if m != nil {
		return m.Granted
	}
	return false
}

func init() {
	proto.RegisterType((*AppendEntriesRequest)(nil), "raftdemo.AppendEntriesRequest")
	proto.RegisterType((*AppendEntriesResponse)(nil), "raftdemo.AppendEntriesResponse")
	proto.RegisterType((*RequestVoteRequest)(nil), "raftdemo.RequestVoteRequest")
	proto.RegisterType((*RequestVoteResponse)(nil), "raftdemo.RequestVoteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LeaderElectionClient is the client API for LeaderElection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LeaderElectionClient interface {
	RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error)
	AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error)
}

type leaderElectionClient struct {
	cc *grpc.ClientConn
}

func NewLeaderElectionClient(cc *grpc.ClientConn) LeaderElectionClient {
	return &leaderElectionClient{cc}
}

func (c *leaderElectionClient) RequestVote(ctx context.Context, in *RequestVoteRequest, opts ...grpc.CallOption) (*RequestVoteResponse, error) {
	out := new(RequestVoteResponse)
	err := c.cc.Invoke(ctx, "/raftdemo.LeaderElection/RequestVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaderElectionClient) AppendEntries(ctx context.Context, in *AppendEntriesRequest, opts ...grpc.CallOption) (*AppendEntriesResponse, error) {
	out := new(AppendEntriesResponse)
	err := c.cc.Invoke(ctx, "/raftdemo.LeaderElection/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaderElectionServer is the server API for LeaderElection service.
type LeaderElectionServer interface {
	RequestVote(context.Context, *RequestVoteRequest) (*RequestVoteResponse, error)
	AppendEntries(context.Context, *AppendEntriesRequest) (*AppendEntriesResponse, error)
}

func RegisterLeaderElectionServer(s *grpc.Server, srv LeaderElectionServer) {
	s.RegisterService(&_LeaderElection_serviceDesc, srv)
}

func _LeaderElection_RequestVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderElectionServer).RequestVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftdemo.LeaderElection/RequestVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderElectionServer).RequestVote(ctx, req.(*RequestVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaderElection_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppendEntriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaderElectionServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raftdemo.LeaderElection/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaderElectionServer).AppendEntries(ctx, req.(*AppendEntriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LeaderElection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "raftdemo.LeaderElection",
	HandlerType: (*LeaderElectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestVote",
			Handler:    _LeaderElection_RequestVote_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _LeaderElection_AppendEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raft.proto",
}

func init() { proto.RegisterFile("raft.proto", fileDescriptor_raft_f30d339453f41f8d) }

var fileDescriptor_raft_f30d339453f41f8d = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4e, 0x03, 0x21,
	0x10, 0xc7, 0x5d, 0xd3, 0xd4, 0x3a, 0x7e, 0x1c, 0xc6, 0x8f, 0x6c, 0x9a, 0xaa, 0x0d, 0xa7, 0x9e,
	0xf6, 0xa0, 0x4f, 0xa0, 0x66, 0x3d, 0xf5, 0xc4, 0xc1, 0x3b, 0xc2, 0x68, 0x36, 0x69, 0x01, 0x61,
	0xfa, 0x62, 0x3e, 0xa1, 0x11, 0x21, 0xb5, 0xb6, 0xdd, 0xdb, 0xfc, 0x99, 0xf0, 0xe3, 0xc7, 0x0c,
	0x40, 0x50, 0xef, 0xdc, 0xf8, 0xe0, 0xd8, 0xe1, 0xe8, 0xa7, 0x36, 0xb4, 0x74, 0xe2, 0x09, 0x2e,
	0x1f, 0xbd, 0x27, 0x6b, 0x5a, 0xcb, 0xa1, 0xa3, 0x28, 0xe9, 0x73, 0x45, 0x91, 0x11, 0x61, 0xc0,
	0x14, 0x96, 0x75, 0x35, 0xad, 0x66, 0x03, 0x99, 0x6a, 0xbc, 0x86, 0xe1, 0x82, 0x94, 0xa1, 0x50,
	0x1f, 0x4e, 0xab, 0xd9, 0xa9, 0xcc, 0x49, 0xb4, 0x70, 0xf5, 0x8f, 0x11, 0xbd, 0xb3, 0x91, 0x76,
	0x42, 0x6a, 0x38, 0x8a, 0x2b, 0xad, 0x29, 0xc6, 0x44, 0x19, 0xc9, 0x12, 0xc5, 0x0b, 0x60, 0x7e,
	0xfd, 0xd5, 0x31, 0xf5, 0x89, 0x4c, 0xe0, 0x58, 0x2b, 0x6b, 0x3a, 0xa3, 0x98, 0xb2, 0xcb, 0xfa,
	0x40, 0x3c, 0xc3, 0xc5, 0x06, 0xa7, 0x5f, 0xe6, 0x23, 0x28, 0xcb, 0x64, 0x8a, 0x4c, 0x8e, 0xf7,
	0x5f, 0x15, 0x9c, 0xcf, 0xd3, 0xf7, 0xda, 0x05, 0x69, 0xee, 0x9c, 0xc5, 0x39, 0x9c, 0xfc, 0xe1,
	0xe2, 0xa4, 0x29, 0x43, 0x6c, 0xb6, 0xb5, 0xc7, 0x37, 0x7b, 0xba, 0xbf, 0x32, 0xe2, 0x00, 0x25,
	0x9c, 0x6d, 0x0c, 0x0d, 0x6f, 0xd7, 0x37, 0x76, 0x6d, 0x64, 0x7c, 0xb7, 0xb7, 0x5f, 0x98, 0x6f,
	0xc3, 0xb4, 0xdd, 0x87, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x46, 0x43, 0xa2, 0xeb, 0x01,
	0x00, 0x00,
}