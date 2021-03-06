// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lachesis.proto

package wire

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type ToServer struct {
	// Types that are valid to be assigned to Event:
	//	*ToServer_Tx_
	//	*ToServer_Answer_
	Event                isToServer_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ToServer) Reset()         { *m = ToServer{} }
func (m *ToServer) String() string { return proto.CompactTextString(m) }
func (*ToServer) ProtoMessage()    {}
func (*ToServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9112d127c29c618, []int{0}
}

func (m *ToServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToServer.Unmarshal(m, b)
}
func (m *ToServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToServer.Marshal(b, m, deterministic)
}
func (m *ToServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToServer.Merge(m, src)
}
func (m *ToServer) XXX_Size() int {
	return xxx_messageInfo_ToServer.Size(m)
}
func (m *ToServer) XXX_DiscardUnknown() {
	xxx_messageInfo_ToServer.DiscardUnknown(m)
}

var xxx_messageInfo_ToServer proto.InternalMessageInfo

type isToServer_Event interface {
	isToServer_Event()
}

type ToServer_Tx_ struct {
	Tx *ToServer_Tx `protobuf:"bytes,1,opt,name=tx,proto3,oneof"`
}

type ToServer_Answer_ struct {
	Answer *ToServer_Answer `protobuf:"bytes,2,opt,name=answer,proto3,oneof"`
}

func (*ToServer_Tx_) isToServer_Event() {}

func (*ToServer_Answer_) isToServer_Event() {}

func (m *ToServer) GetEvent() isToServer_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ToServer) GetTx() *ToServer_Tx {
	if x, ok := m.GetEvent().(*ToServer_Tx_); ok {
		return x.Tx
	}
	return nil
}

func (m *ToServer) GetAnswer() *ToServer_Answer {
	if x, ok := m.GetEvent().(*ToServer_Answer_); ok {
		return x.Answer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ToServer) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ToServer_Tx_)(nil),
		(*ToServer_Answer_)(nil),
	}
}

type ToServer_Tx struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToServer_Tx) Reset()         { *m = ToServer_Tx{} }
func (m *ToServer_Tx) String() string { return proto.CompactTextString(m) }
func (*ToServer_Tx) ProtoMessage()    {}
func (*ToServer_Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9112d127c29c618, []int{0, 0}
}

func (m *ToServer_Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToServer_Tx.Unmarshal(m, b)
}
func (m *ToServer_Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToServer_Tx.Marshal(b, m, deterministic)
}
func (m *ToServer_Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToServer_Tx.Merge(m, src)
}
func (m *ToServer_Tx) XXX_Size() int {
	return xxx_messageInfo_ToServer_Tx.Size(m)
}
func (m *ToServer_Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_ToServer_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_ToServer_Tx proto.InternalMessageInfo

func (m *ToServer_Tx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ToServer_Answer struct {
	Uid []byte `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*ToServer_Answer_Data
	//	*ToServer_Answer_Error
	Payload              isToServer_Answer_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ToServer_Answer) Reset()         { *m = ToServer_Answer{} }
func (m *ToServer_Answer) String() string { return proto.CompactTextString(m) }
func (*ToServer_Answer) ProtoMessage()    {}
func (*ToServer_Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9112d127c29c618, []int{0, 1}
}

func (m *ToServer_Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToServer_Answer.Unmarshal(m, b)
}
func (m *ToServer_Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToServer_Answer.Marshal(b, m, deterministic)
}
func (m *ToServer_Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToServer_Answer.Merge(m, src)
}
func (m *ToServer_Answer) XXX_Size() int {
	return xxx_messageInfo_ToServer_Answer.Size(m)
}
func (m *ToServer_Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_ToServer_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_ToServer_Answer proto.InternalMessageInfo

func (m *ToServer_Answer) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

type isToServer_Answer_Payload interface {
	isToServer_Answer_Payload()
}

type ToServer_Answer_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type ToServer_Answer_Error struct {
	Error string `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*ToServer_Answer_Data) isToServer_Answer_Payload() {}

func (*ToServer_Answer_Error) isToServer_Answer_Payload() {}

func (m *ToServer_Answer) GetPayload() isToServer_Answer_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ToServer_Answer) GetData() []byte {
	if x, ok := m.GetPayload().(*ToServer_Answer_Data); ok {
		return x.Data
	}
	return nil
}

func (m *ToServer_Answer) GetError() string {
	if x, ok := m.GetPayload().(*ToServer_Answer_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ToServer_Answer) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ToServer_Answer_Data)(nil),
		(*ToServer_Answer_Error)(nil),
	}
}

type ToClient struct {
	// Types that are valid to be assigned to Event:
	//	*ToClient_Block_
	//	*ToClient_Query_
	//	*ToClient_Restore_
	Event                isToClient_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ToClient) Reset()         { *m = ToClient{} }
func (m *ToClient) String() string { return proto.CompactTextString(m) }
func (*ToClient) ProtoMessage()    {}
func (*ToClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9112d127c29c618, []int{1}
}

func (m *ToClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient.Unmarshal(m, b)
}
func (m *ToClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient.Marshal(b, m, deterministic)
}
func (m *ToClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient.Merge(m, src)
}
func (m *ToClient) XXX_Size() int {
	return xxx_messageInfo_ToClient.Size(m)
}
func (m *ToClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient proto.InternalMessageInfo

type isToClient_Event interface {
	isToClient_Event()
}

type ToClient_Block_ struct {
	Block *ToClient_Block `protobuf:"bytes,1,opt,name=block,proto3,oneof"`
}

type ToClient_Query_ struct {
	Query *ToClient_Query `protobuf:"bytes,2,opt,name=query,proto3,oneof"`
}

type ToClient_Restore_ struct {
	Restore *ToClient_Restore `protobuf:"bytes,3,opt,name=restore,proto3,oneof"`
}

func (*ToClient_Block_) isToClient_Event() {}

func (*ToClient_Query_) isToClient_Event() {}

func (*ToClient_Restore_) isToClient_Event() {}

func (m *ToClient) GetEvent() isToClient_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ToClient) GetBlock() *ToClient_Block {
	if x, ok := m.GetEvent().(*ToClient_Block_); ok {
		return x.Block
	}
	return nil
}

func (m *ToClient) GetQuery() *ToClient_Query {
	if x, ok := m.GetEvent().(*ToClient_Query_); ok {
		return x.Query
	}
	return nil
}

func (m *ToClient) GetRestore() *ToClient_Restore {
	if x, ok := m.GetEvent().(*ToClient_Restore_); ok {
		return x.Restore
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ToClient) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ToClient_Block_)(nil),
		(*ToClient_Query_)(nil),
		(*ToClient_Restore_)(nil),
	}
}

type ToClient_Block struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToClient_Block) Reset()         { *m = ToClient_Block{} }
func (m *ToClient_Block) String() string { return proto.CompactTextString(m) }
func (*ToClient_Block) ProtoMessage()    {}
func (*ToClient_Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9112d127c29c618, []int{1, 0}
}

func (m *ToClient_Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient_Block.Unmarshal(m, b)
}
func (m *ToClient_Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient_Block.Marshal(b, m, deterministic)
}
func (m *ToClient_Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient_Block.Merge(m, src)
}
func (m *ToClient_Block) XXX_Size() int {
	return xxx_messageInfo_ToClient_Block.Size(m)
}
func (m *ToClient_Block) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient_Block.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient_Block proto.InternalMessageInfo

func (m *ToClient_Block) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *ToClient_Block) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ToClient_Query struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Index                int64    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToClient_Query) Reset()         { *m = ToClient_Query{} }
func (m *ToClient_Query) String() string { return proto.CompactTextString(m) }
func (*ToClient_Query) ProtoMessage()    {}
func (*ToClient_Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9112d127c29c618, []int{1, 1}
}

func (m *ToClient_Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient_Query.Unmarshal(m, b)
}
func (m *ToClient_Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient_Query.Marshal(b, m, deterministic)
}
func (m *ToClient_Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient_Query.Merge(m, src)
}
func (m *ToClient_Query) XXX_Size() int {
	return xxx_messageInfo_ToClient_Query.Size(m)
}
func (m *ToClient_Query) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient_Query.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient_Query proto.InternalMessageInfo

func (m *ToClient_Query) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *ToClient_Query) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ToClient_Restore struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToClient_Restore) Reset()         { *m = ToClient_Restore{} }
func (m *ToClient_Restore) String() string { return proto.CompactTextString(m) }
func (*ToClient_Restore) ProtoMessage()    {}
func (*ToClient_Restore) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9112d127c29c618, []int{1, 2}
}

func (m *ToClient_Restore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient_Restore.Unmarshal(m, b)
}
func (m *ToClient_Restore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient_Restore.Marshal(b, m, deterministic)
}
func (m *ToClient_Restore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient_Restore.Merge(m, src)
}
func (m *ToClient_Restore) XXX_Size() int {
	return xxx_messageInfo_ToClient_Restore.Size(m)
}
func (m *ToClient_Restore) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient_Restore.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient_Restore proto.InternalMessageInfo

func (m *ToClient_Restore) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *ToClient_Restore) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*ToServer)(nil), "wire.ToServer")
	proto.RegisterType((*ToServer_Tx)(nil), "wire.ToServer.Tx")
	proto.RegisterType((*ToServer_Answer)(nil), "wire.ToServer.Answer")
	proto.RegisterType((*ToClient)(nil), "wire.ToClient")
	proto.RegisterType((*ToClient_Block)(nil), "wire.ToClient.Block")
	proto.RegisterType((*ToClient_Query)(nil), "wire.ToClient.Query")
	proto.RegisterType((*ToClient_Restore)(nil), "wire.ToClient.Restore")
}

func init() { proto.RegisterFile("lachesis.proto", fileDescriptor_b9112d127c29c618) }

var fileDescriptor_b9112d127c29c618 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0xdb, 0x42, 0xe9, 0xe5, 0x40, 0xc8, 0xf5, 0x04, 0x49, 0xd3, 0x15, 0xc1, 0x0d, 0x0b,
	0x6d, 0x4d, 0x7d, 0x00, 0x23, 0x6c, 0xba, 0x30, 0x26, 0x0e, 0xbc, 0x40, 0xa1, 0x27, 0xb1, 0xb1,
	0xe9, 0xe0, 0x74, 0x80, 0xf2, 0x5e, 0x3e, 0x8c, 0x8f, 0x63, 0x66, 0x3a, 0x18, 0x21, 0x2c, 0xdc,
	0xcd, 0xe4, 0xfb, 0x3a, 0xff, 0xfc, 0xa7, 0x03, 0x83, 0x22, 0x5d, 0xbf, 0x51, 0x95, 0x57, 0xe1,
	0x46, 0x70, 0xc9, 0xb1, 0xbd, 0xcf, 0x05, 0x4d, 0xbe, 0x6c, 0xf8, 0xb7, 0xe4, 0x0b, 0x12, 0x3b,
	0x12, 0x78, 0x03, 0x8e, 0xac, 0x7d, 0x7b, 0x6c, 0x4f, 0x7b, 0xf1, 0x55, 0xa8, 0x78, 0x78, 0x64,
	0xe1, 0xb2, 0x4e, 0x2c, 0xe6, 0xc8, 0x1a, 0x23, 0xe8, 0xa4, 0x65, 0xb5, 0x27, 0xe1, 0x3b, 0x5a,
	0xbc, 0x3e, 0x13, 0x9f, 0x34, 0x4c, 0x2c, 0x66, 0xb4, 0xc0, 0x07, 0x67, 0x59, 0x23, 0x42, 0x3b,
	0x4b, 0x65, 0xaa, 0x4f, 0xef, 0x33, 0xbd, 0x0e, 0x16, 0xd0, 0x69, 0x6c, 0xfc, 0x0f, 0xad, 0x6d,
	0x9e, 0x19, 0xa8, 0x96, 0x38, 0x34, 0xbe, 0x0a, 0xe9, 0x27, 0x56, 0xf3, 0x05, 0x8e, 0xc0, 0x25,
	0x21, 0xb8, 0xf0, 0x5b, 0x63, 0x7b, 0xda, 0x4d, 0x2c, 0xd6, 0x6c, 0x67, 0x5d, 0xf0, 0x36, 0xe9,
	0xa1, 0xe0, 0x69, 0x36, 0xf3, 0xc0, 0xa5, 0x1d, 0x95, 0x72, 0xf2, 0xe9, 0xa8, 0x6a, 0xf3, 0x22,
	0xa7, 0x52, 0xe2, 0x2d, 0xb8, 0xab, 0x82, 0xaf, 0xdf, 0x4d, 0xbb, 0xe1, 0xf1, 0xd2, 0x0d, 0x0e,
	0x67, 0x8a, 0xa9, 0xe3, 0xb4, 0xa4, 0xec, 0x8f, 0x2d, 0x89, 0x83, 0xa9, 0x78, 0x6e, 0xbf, 0x2a,
	0xa6, 0x6c, 0x2d, 0x61, 0x0c, 0x9e, 0xa0, 0x4a, 0x72, 0x41, 0xfa, 0x5a, 0xbd, 0x78, 0x74, 0xe6,
	0xb3, 0x86, 0x26, 0x16, 0x3b, 0x8a, 0xc1, 0x1d, 0xb8, 0x3a, 0xf3, 0x42, 0x73, 0xfc, 0xdd, 0xdc,
	0x4c, 0x2a, 0x02, 0x57, 0x87, 0x5e, 0x1c, 0x94, 0x9b, 0x97, 0x19, 0xd5, 0xda, 0x6f, 0xb1, 0x66,
	0x13, 0x44, 0xe0, 0x99, 0xd4, 0xbf, 0x25, 0xfc, 0x8c, 0x2d, 0x7e, 0x84, 0xfe, 0xb3, 0x79, 0x29,
	0x2f, 0x3c, 0x23, 0x8c, 0xc0, 0x9b, 0xf3, 0xb2, 0xa4, 0xb5, 0xc4, 0xc1, 0xe9, 0xaf, 0x0e, 0x06,
	0xa7, 0x3d, 0x27, 0xd6, 0xd4, 0xbe, 0xb7, 0x57, 0x1d, 0xfd, 0xbe, 0x1e, 0xbe, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x72, 0xd4, 0x22, 0x29, 0x71, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LachesisNodeClient is the client API for LachesisNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LachesisNodeClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (LachesisNode_ConnectClient, error)
}

type lachesisNodeClient struct {
	cc *grpc.ClientConn
}

func NewLachesisNodeClient(cc *grpc.ClientConn) LachesisNodeClient {
	return &lachesisNodeClient{cc}
}

func (c *lachesisNodeClient) Connect(ctx context.Context, opts ...grpc.CallOption) (LachesisNode_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LachesisNode_serviceDesc.Streams[0], "/wire.LachesisNode/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &lachesisNodeConnectClient{stream}
	return x, nil
}

type LachesisNode_ConnectClient interface {
	Send(*ToServer) error
	Recv() (*ToClient, error)
	grpc.ClientStream
}

type lachesisNodeConnectClient struct {
	grpc.ClientStream
}

func (x *lachesisNodeConnectClient) Send(m *ToServer) error {
	return x.ClientStream.SendMsg(m)
}

func (x *lachesisNodeConnectClient) Recv() (*ToClient, error) {
	m := new(ToClient)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LachesisNodeServer is the server API for LachesisNode service.
type LachesisNodeServer interface {
	Connect(LachesisNode_ConnectServer) error
}

func RegisterLachesisNodeServer(s *grpc.Server, srv LachesisNodeServer) {
	s.RegisterService(&_LachesisNode_serviceDesc, srv)
}

func _LachesisNode_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LachesisNodeServer).Connect(&lachesisNodeConnectServer{stream})
}

type LachesisNode_ConnectServer interface {
	Send(*ToClient) error
	Recv() (*ToServer, error)
	grpc.ServerStream
}

type lachesisNodeConnectServer struct {
	grpc.ServerStream
}

func (x *lachesisNodeConnectServer) Send(m *ToClient) error {
	return x.ServerStream.SendMsg(m)
}

func (x *lachesisNodeConnectServer) Recv() (*ToServer, error) {
	m := new(ToServer)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LachesisNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wire.LachesisNode",
	HandlerType: (*LachesisNodeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _LachesisNode_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "lachesis.proto",
}
