// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push.proto

package push

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

type MessagesRequest struct {
	ChatroomID           string   `protobuf:"bytes,1,opt,name=chatroomID,proto3" json:"chatroomID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagesRequest) Reset()         { *m = MessagesRequest{} }
func (m *MessagesRequest) String() string { return proto.CompactTextString(m) }
func (*MessagesRequest) ProtoMessage()    {}
func (*MessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{0}
}

func (m *MessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagesRequest.Unmarshal(m, b)
}
func (m *MessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagesRequest.Marshal(b, m, deterministic)
}
func (m *MessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagesRequest.Merge(m, src)
}
func (m *MessagesRequest) XXX_Size() int {
	return xxx_messageInfo_MessagesRequest.Size(m)
}
func (m *MessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessagesRequest proto.InternalMessageInfo

func (m *MessagesRequest) GetChatroomID() string {
	if m != nil {
		return m.ChatroomID
	}
	return ""
}

type MessagesReply struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessagesReply) Reset()         { *m = MessagesReply{} }
func (m *MessagesReply) String() string { return proto.CompactTextString(m) }
func (*MessagesReply) ProtoMessage()    {}
func (*MessagesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{1}
}

func (m *MessagesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessagesReply.Unmarshal(m, b)
}
func (m *MessagesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessagesReply.Marshal(b, m, deterministic)
}
func (m *MessagesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessagesReply.Merge(m, src)
}
func (m *MessagesReply) XXX_Size() int {
	return xxx_messageInfo_MessagesReply.Size(m)
}
func (m *MessagesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MessagesReply.DiscardUnknown(m)
}

var xxx_messageInfo_MessagesReply proto.InternalMessageInfo

func (m *MessagesReply) GetMessage() *Message {
	if m != nil {
		return m.Message
	}
	return nil
}

type Message struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ChatroomID           string   `protobuf:"bytes,2,opt,name=chatroomID,proto3" json:"chatroomID,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreateTime           int64    `protobuf:"varint,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1e4bfd2e9d102bb, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Message) GetChatroomID() string {
	if m != nil {
		return m.ChatroomID
	}
	return ""
}

func (m *Message) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*MessagesRequest)(nil), "push.MessagesRequest")
	proto.RegisterType((*MessagesReply)(nil), "push.MessagesReply")
	proto.RegisterType((*Message)(nil), "push.Message")
}

func init() { proto.RegisterFile("push.proto", fileDescriptor_d1e4bfd2e9d102bb) }

var fileDescriptor_d1e4bfd2e9d102bb = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x28, 0x2d, 0xce,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x0c, 0xb9, 0xf8, 0x7d, 0x53,
	0x8b, 0x8b, 0x13, 0xd3, 0x53, 0x8b, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xe4, 0xb8,
	0xb8, 0x92, 0x33, 0x12, 0x4b, 0x8a, 0xf2, 0xf3, 0x73, 0x3d, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0x90, 0x44, 0x94, 0x2c, 0xb8, 0x78, 0x11, 0x5a, 0x0a, 0x72, 0x2a, 0x85, 0xd4, 0xb9,
	0xd8, 0x73, 0x21, 0x02, 0x60, 0xd5, 0xdc, 0x46, 0xbc, 0x7a, 0x60, 0x7b, 0xa0, 0xaa, 0x82, 0x60,
	0xb2, 0x4a, 0xfd, 0x8c, 0x5c, 0xec, 0x50, 0x41, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0xa8, 0xe9,
	0x4c, 0x99, 0x29, 0x68, 0xb6, 0x32, 0xa1, 0xdb, 0x2a, 0x24, 0xc5, 0xc5, 0x51, 0x5a, 0x9c, 0x5a,
	0x94, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x0c, 0x96, 0x85, 0xf3, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0xf3,
	0xf3, 0x4a, 0x52, 0xf3, 0x4a, 0x24, 0x58, 0xc0, 0x52, 0x30, 0x2e, 0xd8, 0xd4, 0xa2, 0xd4, 0xc4,
	0x92, 0xd4, 0x90, 0xcc, 0xdc, 0x54, 0x09, 0x56, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x24, 0x11, 0x23,
	0x77, 0x2e, 0xee, 0x80, 0xd2, 0xe2, 0x8c, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x0b,
	0x2e, 0x0e, 0x98, 0xd7, 0x84, 0x44, 0x51, 0x3c, 0x01, 0x0b, 0x1d, 0x29, 0x61, 0x74, 0xe1, 0x82,
	0x9c, 0x4a, 0x03, 0xc6, 0x24, 0x36, 0x70, 0xa0, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xff,
	0x36, 0xc6, 0x16, 0x62, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PushServiceClient is the client API for PushService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PushServiceClient interface {
	Messages(ctx context.Context, in *MessagesRequest, opts ...grpc.CallOption) (PushService_MessagesClient, error)
}

type pushServiceClient struct {
	cc *grpc.ClientConn
}

func NewPushServiceClient(cc *grpc.ClientConn) PushServiceClient {
	return &pushServiceClient{cc}
}

func (c *pushServiceClient) Messages(ctx context.Context, in *MessagesRequest, opts ...grpc.CallOption) (PushService_MessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PushService_serviceDesc.Streams[0], "/push.PushService/Messages", opts...)
	if err != nil {
		return nil, err
	}
	x := &pushServiceMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PushService_MessagesClient interface {
	Recv() (*MessagesReply, error)
	grpc.ClientStream
}

type pushServiceMessagesClient struct {
	grpc.ClientStream
}

func (x *pushServiceMessagesClient) Recv() (*MessagesReply, error) {
	m := new(MessagesReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PushServiceServer is the server API for PushService service.
type PushServiceServer interface {
	Messages(*MessagesRequest, PushService_MessagesServer) error
}

func RegisterPushServiceServer(s *grpc.Server, srv PushServiceServer) {
	s.RegisterService(&_PushService_serviceDesc, srv)
}

func _PushService_Messages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(MessagesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PushServiceServer).Messages(m, &pushServiceMessagesServer{stream})
}

type PushService_MessagesServer interface {
	Send(*MessagesReply) error
	grpc.ServerStream
}

type pushServiceMessagesServer struct {
	grpc.ServerStream
}

func (x *pushServiceMessagesServer) Send(m *MessagesReply) error {
	return x.ServerStream.SendMsg(m)
}

var _PushService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "push.PushService",
	HandlerType: (*PushServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Messages",
			Handler:       _PushService_Messages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "push.proto",
}
