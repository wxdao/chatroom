// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push.proto

package pb

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

func init() {
	proto.RegisterType((*MessagesRequest)(nil), "pb.MessagesRequest")
	proto.RegisterType((*MessagesReply)(nil), "pb.MessagesReply")
	proto.RegisterType((*Message)(nil), "pb.Message")
}

func init() { proto.RegisterFile("push.proto", fileDescriptor_d1e4bfd2e9d102bb) }

var fileDescriptor_d1e4bfd2e9d102bb = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x28, 0x2d, 0xce,
	0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x32, 0xe4, 0xe2, 0xf7, 0x4d,
	0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe3,
	0xe2, 0x4a, 0xce, 0x48, 0x2c, 0x29, 0xca, 0xcf, 0xcf, 0xf5, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x42, 0x12, 0x51, 0x32, 0xe3, 0xe2, 0x45, 0x68, 0x29, 0xc8, 0xa9, 0x14, 0x52, 0xe5,
	0x62, 0xcf, 0x85, 0x08, 0x80, 0x55, 0x73, 0x1b, 0x71, 0xeb, 0x15, 0x24, 0xe9, 0x41, 0xd5, 0x04,
	0xc1, 0xe4, 0x94, 0xf2, 0xb9, 0xd8, 0xa1, 0x62, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x50, 0xa3,
	0x99, 0x32, 0x53, 0xd0, 0xac, 0x64, 0x42, 0xb7, 0x52, 0x48, 0x8a, 0x8b, 0xa3, 0xb4, 0x38, 0xb5,
	0x28, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x19, 0x2c, 0x0b, 0xe7, 0x0b, 0x49, 0x70, 0xb1, 0x27, 0xe7,
	0xe7, 0x95, 0xa4, 0xe6, 0x95, 0x48, 0xb0, 0x80, 0xa5, 0x60, 0x5c, 0x23, 0x67, 0x2e, 0xee, 0x80,
	0xd2, 0xe2, 0x8c, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x13, 0x2e, 0x0e, 0x98, 0xbb,
	0x85, 0x84, 0x91, 0x5c, 0x08, 0xf3, 0xb8, 0x94, 0x20, 0xaa, 0x60, 0x41, 0x4e, 0xa5, 0x01, 0x63,
	0x12, 0x1b, 0x38, 0xac, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa6, 0x10, 0xee, 0xed, 0x39,
	0x01, 0x00, 0x00,
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
	stream, err := c.cc.NewStream(ctx, &_PushService_serviceDesc.Streams[0], "/pb.PushService/Messages", opts...)
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
	ServiceName: "pb.PushService",
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