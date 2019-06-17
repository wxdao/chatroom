// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

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

type QueryMessagesInRangeRequest struct {
	ChatroomID           string   `protobuf:"bytes,1,opt,name=chatroomID,proto3" json:"chatroomID,omitempty"`
	StartID              string   `protobuf:"bytes,2,opt,name=startID,proto3" json:"startID,omitempty"`
	EndID                string   `protobuf:"bytes,3,opt,name=endID,proto3" json:"endID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryMessagesInRangeRequest) Reset()         { *m = QueryMessagesInRangeRequest{} }
func (m *QueryMessagesInRangeRequest) String() string { return proto.CompactTextString(m) }
func (*QueryMessagesInRangeRequest) ProtoMessage()    {}
func (*QueryMessagesInRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *QueryMessagesInRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryMessagesInRangeRequest.Unmarshal(m, b)
}
func (m *QueryMessagesInRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryMessagesInRangeRequest.Marshal(b, m, deterministic)
}
func (m *QueryMessagesInRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMessagesInRangeRequest.Merge(m, src)
}
func (m *QueryMessagesInRangeRequest) XXX_Size() int {
	return xxx_messageInfo_QueryMessagesInRangeRequest.Size(m)
}
func (m *QueryMessagesInRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMessagesInRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMessagesInRangeRequest proto.InternalMessageInfo

func (m *QueryMessagesInRangeRequest) GetChatroomID() string {
	if m != nil {
		return m.ChatroomID
	}
	return ""
}

func (m *QueryMessagesInRangeRequest) GetStartID() string {
	if m != nil {
		return m.StartID
	}
	return ""
}

func (m *QueryMessagesInRangeRequest) GetEndID() string {
	if m != nil {
		return m.EndID
	}
	return ""
}

type QueryMessagesInRangeReply struct {
	Messages             []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *QueryMessagesInRangeReply) Reset()         { *m = QueryMessagesInRangeReply{} }
func (m *QueryMessagesInRangeReply) String() string { return proto.CompactTextString(m) }
func (*QueryMessagesInRangeReply) ProtoMessage()    {}
func (*QueryMessagesInRangeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *QueryMessagesInRangeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryMessagesInRangeReply.Unmarshal(m, b)
}
func (m *QueryMessagesInRangeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryMessagesInRangeReply.Marshal(b, m, deterministic)
}
func (m *QueryMessagesInRangeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryMessagesInRangeReply.Merge(m, src)
}
func (m *QueryMessagesInRangeReply) XXX_Size() int {
	return xxx_messageInfo_QueryMessagesInRangeReply.Size(m)
}
func (m *QueryMessagesInRangeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryMessagesInRangeReply.DiscardUnknown(m)
}

var xxx_messageInfo_QueryMessagesInRangeReply proto.InternalMessageInfo

func (m *QueryMessagesInRangeReply) GetMessages() []*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type SendMessageRequest struct {
	ChatroomID           string   `protobuf:"bytes,1,opt,name=chatroomID,proto3" json:"chatroomID,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageRequest) Reset()         { *m = SendMessageRequest{} }
func (m *SendMessageRequest) String() string { return proto.CompactTextString(m) }
func (*SendMessageRequest) ProtoMessage()    {}
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *SendMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageRequest.Unmarshal(m, b)
}
func (m *SendMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageRequest.Marshal(b, m, deterministic)
}
func (m *SendMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageRequest.Merge(m, src)
}
func (m *SendMessageRequest) XXX_Size() int {
	return xxx_messageInfo_SendMessageRequest.Size(m)
}
func (m *SendMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageRequest proto.InternalMessageInfo

func (m *SendMessageRequest) GetChatroomID() string {
	if m != nil {
		return m.ChatroomID
	}
	return ""
}

func (m *SendMessageRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SendMessageRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type SendMessageReply struct {
	Message              *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMessageReply) Reset()         { *m = SendMessageReply{} }
func (m *SendMessageReply) String() string { return proto.CompactTextString(m) }
func (*SendMessageReply) ProtoMessage()    {}
func (*SendMessageReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *SendMessageReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMessageReply.Unmarshal(m, b)
}
func (m *SendMessageReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMessageReply.Marshal(b, m, deterministic)
}
func (m *SendMessageReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMessageReply.Merge(m, src)
}
func (m *SendMessageReply) XXX_Size() int {
	return xxx_messageInfo_SendMessageReply.Size(m)
}
func (m *SendMessageReply) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMessageReply.DiscardUnknown(m)
}

var xxx_messageInfo_SendMessageReply proto.InternalMessageInfo

func (m *SendMessageReply) GetMessage() *Message {
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
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
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
	proto.RegisterType((*QueryMessagesInRangeRequest)(nil), "pb.QueryMessagesInRangeRequest")
	proto.RegisterType((*QueryMessagesInRangeReply)(nil), "pb.QueryMessagesInRangeReply")
	proto.RegisterType((*SendMessageRequest)(nil), "pb.SendMessageRequest")
	proto.RegisterType((*SendMessageReply)(nil), "pb.SendMessageReply")
	proto.RegisterType((*Message)(nil), "pb.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0x49, 0xf2, 0xbe, 0xa6, 0xde, 0x60, 0x91, 0x4b, 0x90, 0x18, 0x51, 0x4b, 0x40, 0xec,
	0x2a, 0x8b, 0xba, 0x12, 0xb7, 0xd9, 0x64, 0xe1, 0xc2, 0x14, 0xdc, 0xe7, 0xe3, 0x52, 0x23, 0xcd,
	0x4c, 0x9c, 0x99, 0x08, 0xf9, 0x41, 0xfe, 0x4f, 0x49, 0x3a, 0x29, 0xfd, 0xc6, 0xe5, 0x39, 0xe7,
	0xc2, 0x79, 0xe6, 0xde, 0x81, 0x8b, 0x8a, 0xa4, 0x4c, 0x17, 0x14, 0xd6, 0x82, 0x2b, 0x8e, 0x66,
	0x9d, 0x05, 0x15, 0xdc, 0xbc, 0x35, 0x24, 0xda, 0xd7, 0x55, 0x22, 0x63, 0x96, 0xa4, 0x6c, 0x41,
	0x09, 0x7d, 0x35, 0x24, 0x15, 0xde, 0x01, 0xe4, 0x1f, 0xa9, 0x12, 0x9c, 0x57, 0x71, 0xe4, 0x19,
	0x13, 0x63, 0x7a, 0x9e, 0x6c, 0x38, 0xe8, 0x81, 0x2d, 0x55, 0x2a, 0x54, 0x1c, 0x79, 0x66, 0x1f,
	0x0e, 0x12, 0x5d, 0xf8, 0x4f, 0xac, 0x88, 0x23, 0xcf, 0xea, 0xfd, 0x95, 0x08, 0x22, 0xb8, 0x3e,
	0x5c, 0x57, 0x2f, 0x5b, 0x7c, 0x84, 0x91, 0x06, 0x94, 0x9e, 0x31, 0xb1, 0xa6, 0xce, 0xcc, 0x09,
	0xeb, 0x2c, 0xd4, 0xb3, 0xc9, 0x3a, 0x0c, 0x3e, 0x01, 0xe7, 0xc4, 0x8a, 0x21, 0xf8, 0x23, 0xab,
	0x0f, 0xa3, 0x46, 0x92, 0x60, 0x69, 0x45, 0x1a, 0x76, 0xad, 0xbb, 0x77, 0xe4, 0x9c, 0x29, 0x62,
	0x4a, 0xf3, 0x0e, 0x32, 0x78, 0x86, 0xcb, 0xad, 0xae, 0x0e, 0xf4, 0x01, 0x6c, 0xcd, 0xd2, 0xd7,
	0xec, 0x70, 0x0e, 0x59, 0xc0, 0xc1, 0xd6, 0x1e, 0x8e, 0xc1, 0x2c, 0x0b, 0xcd, 0x64, 0x96, 0xc5,
	0x0e, 0xab, 0x79, 0x92, 0xd5, 0x3a, 0xce, 0xfa, 0x6f, 0x8b, 0x75, 0xf6, 0x63, 0xc0, 0x58, 0x37,
	0xce, 0x49, 0x7c, 0x97, 0x39, 0xe1, 0x3b, 0xb8, 0x87, 0x16, 0x8e, 0xf7, 0x1d, 0xf1, 0x89, 0xcb,
	0xfb, 0xb7, 0xc7, 0x07, 0xba, 0x15, 0xbc, 0x80, 0xb3, 0xb1, 0x16, 0xbc, 0xea, 0xa6, 0xf7, 0x6f,
	0xe2, 0xbb, 0x7b, 0x7e, 0xbd, 0x6c, 0xb3, 0xb3, 0xfe, 0xff, 0x3d, 0xfd, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x5c, 0xfd, 0x12, 0x81, 0x90, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	QueryMessagesInRange(ctx context.Context, in *QueryMessagesInRangeRequest, opts ...grpc.CallOption) (*QueryMessagesInRangeReply, error)
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageReply, error)
}

type messageServiceClient struct {
	cc *grpc.ClientConn
}

func NewMessageServiceClient(cc *grpc.ClientConn) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) QueryMessagesInRange(ctx context.Context, in *QueryMessagesInRangeRequest, opts ...grpc.CallOption) (*QueryMessagesInRangeReply, error) {
	out := new(QueryMessagesInRangeReply)
	err := c.cc.Invoke(ctx, "/pb.MessageService/QueryMessagesInRange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageReply, error) {
	out := new(SendMessageReply)
	err := c.cc.Invoke(ctx, "/pb.MessageService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	QueryMessagesInRange(context.Context, *QueryMessagesInRangeRequest) (*QueryMessagesInRangeReply, error)
	SendMessage(context.Context, *SendMessageRequest) (*SendMessageReply, error)
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_QueryMessagesInRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMessagesInRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).QueryMessagesInRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageService/QueryMessagesInRange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).QueryMessagesInRange(ctx, req.(*QueryMessagesInRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.MessageService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryMessagesInRange",
			Handler:    _MessageService_QueryMessagesInRange_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _MessageService_SendMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}
