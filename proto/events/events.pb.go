// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/events.proto

package events

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

type PublishRequest struct {
	Topic                string            `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload              []byte            `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp            int64             `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{0}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PublishRequest) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *PublishRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PublishRequest) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PublishResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{1}
}

func (m *PublishResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponse.Unmarshal(m, b)
}
func (m *PublishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponse.Marshal(b, m, deterministic)
}
func (m *PublishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponse.Merge(m, src)
}
func (m *PublishResponse) XXX_Size() int {
	return xxx_messageInfo_PublishResponse.Size(m)
}
func (m *PublishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponse proto.InternalMessageInfo

type SubscribeRequest struct {
	Queue       string `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
	Topic       string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	StartAtTime int64  `protobuf:"varint,3,opt,name=start_at_time,json=startAtTime,proto3" json:"start_at_time,omitempty"`
	AutoAck     bool   `protobuf:"varint,4,opt,name=auto_ack,json=autoAck,proto3" json:"auto_ack,omitempty"`
	// duration in nanoseconds
	AckWait              int64    `protobuf:"varint,5,opt,name=ack_wait,json=ackWait,proto3" json:"ack_wait,omitempty"`
	RetryLimit           int64    `protobuf:"varint,6,opt,name=retry_limit,json=retryLimit,proto3" json:"retry_limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{2}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetQueue() string {
	if m != nil {
		return m.Queue
	}
	return ""
}

func (m *SubscribeRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SubscribeRequest) GetStartAtTime() int64 {
	if m != nil {
		return m.StartAtTime
	}
	return 0
}

func (m *SubscribeRequest) GetAutoAck() bool {
	if m != nil {
		return m.AutoAck
	}
	return false
}

func (m *SubscribeRequest) GetAckWait() int64 {
	if m != nil {
		return m.AckWait
	}
	return 0
}

func (m *SubscribeRequest) GetRetryLimit() int64 {
	if m != nil {
		return m.RetryLimit
	}
	return 0
}

type Event struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string            `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Metadata             map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Payload              []byte            `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	Timestamp            int64             `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{3}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Event) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Event) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Event) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type ReadRequest struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Limit                uint64   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               uint64   `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{4}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *ReadRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReadRequest) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ReadResponse struct {
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{5}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type WriteRequest struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Ttl                  int64    `protobuf:"varint,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{6}
}

func (m *WriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteRequest.Unmarshal(m, b)
}
func (m *WriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteRequest.Marshal(b, m, deterministic)
}
func (m *WriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteRequest.Merge(m, src)
}
func (m *WriteRequest) XXX_Size() int {
	return xxx_messageInfo_WriteRequest.Size(m)
}
func (m *WriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WriteRequest proto.InternalMessageInfo

func (m *WriteRequest) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *WriteRequest) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

type WriteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WriteResponse) Reset()         { *m = WriteResponse{} }
func (m *WriteResponse) String() string { return proto.CompactTextString(m) }
func (*WriteResponse) ProtoMessage()    {}
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{7}
}

func (m *WriteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WriteResponse.Unmarshal(m, b)
}
func (m *WriteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WriteResponse.Marshal(b, m, deterministic)
}
func (m *WriteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WriteResponse.Merge(m, src)
}
func (m *WriteResponse) XXX_Size() int {
	return xxx_messageInfo_WriteResponse.Size(m)
}
func (m *WriteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WriteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WriteResponse proto.InternalMessageInfo

type AckRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Success              bool     `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AckRequest) Reset()         { *m = AckRequest{} }
func (m *AckRequest) String() string { return proto.CompactTextString(m) }
func (*AckRequest) ProtoMessage()    {}
func (*AckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ec31f2d2a3db598, []int{8}
}

func (m *AckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AckRequest.Unmarshal(m, b)
}
func (m *AckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AckRequest.Marshal(b, m, deterministic)
}
func (m *AckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AckRequest.Merge(m, src)
}
func (m *AckRequest) XXX_Size() int {
	return xxx_messageInfo_AckRequest.Size(m)
}
func (m *AckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AckRequest proto.InternalMessageInfo

func (m *AckRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AckRequest) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*PublishRequest)(nil), "events.PublishRequest")
	proto.RegisterMapType((map[string]string)(nil), "events.PublishRequest.MetadataEntry")
	proto.RegisterType((*PublishResponse)(nil), "events.PublishResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "events.SubscribeRequest")
	proto.RegisterType((*Event)(nil), "events.Event")
	proto.RegisterMapType((map[string]string)(nil), "events.Event.MetadataEntry")
	proto.RegisterType((*ReadRequest)(nil), "events.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "events.ReadResponse")
	proto.RegisterType((*WriteRequest)(nil), "events.WriteRequest")
	proto.RegisterType((*WriteResponse)(nil), "events.WriteResponse")
	proto.RegisterType((*AckRequest)(nil), "events.AckRequest")
}

func init() { proto.RegisterFile("events/events.proto", fileDescriptor_8ec31f2d2a3db598) }

var fileDescriptor_8ec31f2d2a3db598 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0xed, 0x38, 0x1f, 0x93, 0xa4, 0x2d, 0xdb, 0x50, 0x4c, 0x40, 0x22, 0x32, 0x20, 0xe5,
	0x80, 0x12, 0x48, 0xa1, 0xa0, 0xf6, 0x42, 0x90, 0x72, 0x03, 0x09, 0x36, 0x48, 0x95, 0xb8, 0x44,
	0x1b, 0x67, 0x4b, 0x57, 0x8e, 0x63, 0xd7, 0x3b, 0x0e, 0x8a, 0xc4, 0x1f, 0xe3, 0xb7, 0x20, 0xfe,
	0x0b, 0xf2, 0xee, 0x3a, 0x4d, 0x52, 0xa8, 0x38, 0x70, 0x49, 0xfc, 0x66, 0x76, 0xc6, 0x6f, 0xde,
	0x9b, 0x35, 0x1c, 0xf2, 0x25, 0x5f, 0xa0, 0xec, 0xeb, 0xbf, 0x5e, 0x92, 0xc6, 0x18, 0x93, 0xb2,
	0x46, 0xfe, 0x2f, 0x0b, 0xf6, 0x3e, 0x66, 0xd3, 0xb9, 0x90, 0x97, 0x94, 0x5f, 0x65, 0x5c, 0x22,
	0x69, 0x81, 0x8b, 0x71, 0x22, 0x02, 0xcf, 0xea, 0x58, 0xdd, 0x1a, 0xd5, 0x80, 0xbc, 0x85, 0x6a,
	0xc4, 0x91, 0xcd, 0x18, 0x32, 0xcf, 0xee, 0x38, 0xdd, 0xfa, 0xe0, 0x49, 0xcf, 0x74, 0xdc, 0xae,
	0xef, 0x7d, 0x30, 0xc7, 0x46, 0x0b, 0x4c, 0x57, 0x74, 0x5d, 0x45, 0x3c, 0xa8, 0x24, 0x6c, 0x35,
	0x8f, 0xd9, 0xcc, 0x73, 0x3a, 0x56, 0xb7, 0x41, 0x0b, 0x48, 0x1e, 0x42, 0x0d, 0x45, 0xc4, 0x25,
	0xb2, 0x28, 0xf1, 0x4a, 0x1d, 0xab, 0xeb, 0xd0, 0xeb, 0x40, 0xfb, 0x0c, 0x9a, 0x5b, 0x2d, 0xc9,
	0x01, 0x38, 0x21, 0x5f, 0x19, 0x7a, 0xf9, 0x63, 0x4e, 0x79, 0xc9, 0xe6, 0x19, 0xf7, 0x6c, 0x4d,
	0x59, 0x81, 0x53, 0xfb, 0x8d, 0xe5, 0xdf, 0x81, 0xfd, 0x35, 0x3d, 0x99, 0xc4, 0x0b, 0xc9, 0xfd,
	0x1f, 0x16, 0x1c, 0x8c, 0xb3, 0xa9, 0x0c, 0x52, 0x31, 0xe5, 0x1b, 0x43, 0x5f, 0x65, 0x3c, 0xe3,
	0xc5, 0xd0, 0x0a, 0x5c, 0x4b, 0x61, 0x6f, 0x4a, 0xe1, 0x43, 0x53, 0x22, 0x4b, 0x71, 0xc2, 0x70,
	0x92, 0xd3, 0x54, 0xe3, 0x38, 0xb4, 0xae, 0x82, 0x43, 0xfc, 0x2c, 0x22, 0x4e, 0xee, 0x43, 0x95,
	0x65, 0x18, 0x4f, 0x58, 0x10, 0xaa, 0x89, 0xaa, 0xb4, 0x92, 0xe3, 0x61, 0x10, 0xaa, 0x54, 0x10,
	0x4e, 0xbe, 0x31, 0x81, 0x9e, 0xab, 0x2a, 0x2b, 0x2c, 0x08, 0xcf, 0x99, 0x40, 0xf2, 0x08, 0xea,
	0x29, 0xc7, 0x74, 0x35, 0x99, 0x8b, 0x48, 0xa0, 0x57, 0x56, 0x59, 0x50, 0xa1, 0xf7, 0x79, 0xc4,
	0xff, 0x69, 0x81, 0x3b, 0xca, 0x55, 0x27, 0x7b, 0x60, 0x8b, 0x99, 0x61, 0x6b, 0x8b, 0xd9, 0x5f,
	0xa8, 0xbe, 0xde, 0x70, 0xcd, 0x51, 0xae, 0x3d, 0x28, 0x5c, 0x53, 0x6d, 0xfe, 0xc5, 0xac, 0xd2,
	0x2d, 0x66, 0xb9, 0xff, 0xd5, 0xac, 0x4f, 0x50, 0xa7, 0x9c, 0xcd, 0x6e, 0x5f, 0xc4, 0x16, 0xb8,
	0x5a, 0x9d, 0xbc, 0xbc, 0x44, 0x35, 0x20, 0x47, 0x50, 0x8e, 0x2f, 0x2e, 0x24, 0x47, 0x65, 0x46,
	0x89, 0x1a, 0xe4, 0xbf, 0x82, 0x86, 0x6e, 0xa9, 0xcd, 0x27, 0x4f, 0xc1, 0x6c, 0xbe, 0x67, 0x29,
	0x39, 0x9a, 0x5b, 0x72, 0xd0, 0xe2, 0x5a, 0x8c, 0xa0, 0x71, 0x9e, 0x0a, 0x5c, 0xaf, 0xc7, 0x63,
	0x70, 0x55, 0x46, 0x51, 0xb9, 0x51, 0xa5, 0x73, 0xf9, 0xa8, 0x88, 0x73, 0xc5, 0xcb, 0xa1, 0xf9,
	0xa3, 0xbf, 0x0f, 0x4d, 0xd3, 0xc6, 0xec, 0xde, 0x09, 0xc0, 0x30, 0x08, 0x8b, 0xae, 0xbb, 0x1e,
	0x7a, 0x50, 0x91, 0x59, 0x10, 0x70, 0x29, 0x55, 0x93, 0x2a, 0x2d, 0xe0, 0xe0, 0x3b, 0x94, 0xc7,
	0x98, 0x72, 0x16, 0x91, 0x53, 0xa8, 0x98, 0x85, 0x26, 0x47, 0x7f, 0xbe, 0x80, 0xed, 0x7b, 0x37,
	0xe2, 0x66, 0xf8, 0x13, 0xa8, 0xad, 0x17, 0x9f, 0x78, 0xc5, 0xa9, 0xdd, 0xbb, 0xd0, 0xde, 0x9e,
	0xee, 0xb9, 0x35, 0x48, 0xc0, 0x1d, 0x63, 0x9c, 0x72, 0xf2, 0x02, 0x4a, 0xb9, 0x9a, 0xe4, 0xb0,
	0x38, 0xb1, 0x61, 0x57, 0xbb, 0xb5, 0x1d, 0x34, 0xef, 0x7c, 0x09, 0xae, 0x92, 0x80, 0xac, 0xd3,
	0x9b, 0xc2, 0xb6, 0xef, 0xee, 0x44, 0x75, 0xd5, 0xbb, 0xde, 0x97, 0x67, 0x5f, 0x05, 0x5e, 0x66,
	0xd3, 0x5e, 0x10, 0x47, 0xfd, 0x48, 0x04, 0x69, 0x6c, 0x7e, 0x97, 0xc7, 0x7d, 0xf5, 0x05, 0xd3,
	0x9f, 0xb3, 0x33, 0x5d, 0x3d, 0x2d, 0xab, 0xd8, 0xf1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21,
	0x39, 0xe5, 0xe4, 0xec, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Stream_SubscribeClient, error)
}

type streamClient struct {
	cc *grpc.ClientConn
}

func NewStreamClient(cc *grpc.ClientConn) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/events.Stream/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (Stream_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Stream_serviceDesc.Streams[0], "/events.Stream/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stream_SubscribeClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type streamSubscribeClient struct {
	grpc.ClientStream
}

func (x *streamSubscribeClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamServer is the server API for Stream service.
type StreamServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	Subscribe(*SubscribeRequest, Stream_SubscribeServer) error
}

// UnimplementedStreamServer can be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (*UnimplementedStreamServer) Publish(ctx context.Context, req *PublishRequest) (*PublishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (*UnimplementedStreamServer) Subscribe(req *SubscribeRequest, srv Stream_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterStreamServer(s *grpc.Server, srv StreamServer) {
	s.RegisterService(&_Stream_serviceDesc, srv)
}

func _Stream_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.Stream/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stream_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServer).Subscribe(m, &streamSubscribeServer{stream})
}

type Stream_SubscribeServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type streamSubscribeServer struct {
	grpc.ServerStream
}

func (x *streamSubscribeServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

var _Stream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _Stream_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Stream_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "events/events.proto",
}

// StoreClient is the client API for Store service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
}

type storeClient struct {
	cc *grpc.ClientConn
}

func NewStoreClient(cc *grpc.ClientConn) StoreClient {
	return &storeClient{cc}
}

func (c *storeClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/events.Store/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/events.Store/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreServer is the server API for Store service.
type StoreServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
}

// UnimplementedStoreServer can be embedded to have forward compatible implementations.
type UnimplementedStoreServer struct {
}

func (*UnimplementedStoreServer) Read(ctx context.Context, req *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (*UnimplementedStoreServer) Write(ctx context.Context, req *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}

func RegisterStoreServer(s *grpc.Server, srv StoreServer) {
	s.RegisterService(&_Store_serviceDesc, srv)
}

func _Store_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.Store/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Store_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/events.Store/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Store_serviceDesc = grpc.ServiceDesc{
	ServiceName: "events.Store",
	HandlerType: (*StoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Store_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Store_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "events/events.proto",
}
