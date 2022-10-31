// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/template.proto

//https://github.com/PatrickMatthiesen/DSYS-gRPC-template#implementation

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//Amount is a type containing a string and int. They are intialized as the first and second parameter value.
type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName string `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Value      int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_proto_template_proto_rawDescGZIP(), []int{0}
}

func (x *Amount) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *Amount) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewValue int64 `protobuf:"varint,1,opt,name=newValue,proto3" json:"newValue,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_proto_template_proto_rawDescGZIP(), []int{1}
}

func (x *Ack) GetNewValue() int64 {
	if x != nil {
		return x.NewValue
	}
	return 0
}

type Greeding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientName string `protobuf:"bytes,1,opt,name=clientName,proto3" json:"clientName,omitempty"`
	Message    string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Greeding) Reset() {
	*x = Greeding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Greeding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Greeding) ProtoMessage() {}

func (x *Greeding) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Greeding.ProtoReflect.Descriptor instead.
func (*Greeding) Descriptor() ([]byte, []int) {
	return file_proto_template_proto_rawDescGZIP(), []int{2}
}

func (x *Greeding) GetClientName() string {
	if x != nil {
		return x.ClientName
	}
	return ""
}

func (x *Greeding) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Farewell struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Farewell) Reset() {
	*x = Farewell{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Farewell) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Farewell) ProtoMessage() {}

func (x *Farewell) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Farewell.ProtoReflect.Descriptor instead.
func (*Farewell) Descriptor() ([]byte, []int) {
	return file_proto_template_proto_rawDescGZIP(), []int{3}
}

func (x *Farewell) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type AskForTimeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId int64 `protobuf:"varint,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
}

func (x *AskForTimeMessage) Reset() {
	*x = AskForTimeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskForTimeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskForTimeMessage) ProtoMessage() {}

func (x *AskForTimeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskForTimeMessage.ProtoReflect.Descriptor instead.
func (*AskForTimeMessage) Descriptor() ([]byte, []int) {
	return file_proto_template_proto_rawDescGZIP(), []int{4}
}

func (x *AskForTimeMessage) GetClientId() int64 {
	if x != nil {
		return x.ClientId
	}
	return 0
}

type TimeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,1,opt,name=serverName,proto3" json:"serverName,omitempty"`
	Time       string `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *TimeMessage) Reset() {
	*x = TimeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeMessage) ProtoMessage() {}

func (x *TimeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeMessage.ProtoReflect.Descriptor instead.
func (*TimeMessage) Descriptor() ([]byte, []int) {
	return file_proto_template_proto_rawDescGZIP(), []int{5}
}

func (x *TimeMessage) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *TimeMessage) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	Time    int64  `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_proto_template_proto_rawDescGZIP(), []int{6}
}

func (x *ChatRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatRequest) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,4,opt,name=message,proto3" json:"message,omitempty"`
	Time    int64  `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_template_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_template_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_proto_template_proto_rawDescGZIP(), []int{7}
}

func (x *ChatResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatResponse) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_proto_template_proto protoreflect.FileDescriptor

var file_proto_template_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a,
	0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a,
	0x03, 0x41, 0x63, 0x6b, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x44, 0x0a, 0x08, 0x47, 0x72, 0x65, 0x65, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x46, 0x61, 0x72, 0x65, 0x77, 0x65,
	0x6c, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x11,
	0x41, 0x73, 0x6b, 0x46, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x41, 0x0a,
	0x0b, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x22, 0x3b, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3c, 0x0a,
	0x0c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x32, 0x7d, 0x0a, 0x08, 0x54,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x41, 0x73, 0x6b, 0x46, 0x6f,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x73,
	0x6b, 0x46, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x61, 0x74, 0x72, 0x69, 0x63, 0x6b,
	0x4d, 0x61, 0x74, 0x74, 0x68, 0x69, 0x65, 0x73, 0x65, 0x6e, 0x2f, 0x44, 0x53, 0x59, 0x53, 0x2d,
	0x67, 0x52, 0x50, 0x43, 0x2d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_template_proto_rawDescOnce sync.Once
	file_proto_template_proto_rawDescData = file_proto_template_proto_rawDesc
)

func file_proto_template_proto_rawDescGZIP() []byte {
	file_proto_template_proto_rawDescOnce.Do(func() {
		file_proto_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_template_proto_rawDescData)
	})
	return file_proto_template_proto_rawDescData
}

var file_proto_template_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_template_proto_goTypes = []interface{}{
	(*Amount)(nil),            // 0: proto.Amount
	(*Ack)(nil),               // 1: proto.Ack
	(*Greeding)(nil),          // 2: proto.Greeding
	(*Farewell)(nil),          // 3: proto.Farewell
	(*AskForTimeMessage)(nil), // 4: proto.AskForTimeMessage
	(*TimeMessage)(nil),       // 5: proto.TimeMessage
	(*ChatRequest)(nil),       // 6: proto.ChatRequest
	(*ChatResponse)(nil),      // 7: proto.ChatResponse
}
var file_proto_template_proto_depIdxs = []int32{
	4, // 0: proto.Template.AskForTime:input_type -> proto.AskForTimeMessage
	6, // 1: proto.Template.chat:input_type -> proto.ChatRequest
	5, // 2: proto.Template.AskForTime:output_type -> proto.TimeMessage
	7, // 3: proto.Template.chat:output_type -> proto.ChatResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_template_proto_init() }
func file_proto_template_proto_init() {
	if File_proto_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Greeding); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Farewell); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskForTimeMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_template_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_template_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_template_proto_goTypes,
		DependencyIndexes: file_proto_template_proto_depIdxs,
		MessageInfos:      file_proto_template_proto_msgTypes,
	}.Build()
	File_proto_template_proto = out.File
	file_proto_template_proto_rawDesc = nil
	file_proto_template_proto_goTypes = nil
	file_proto_template_proto_depIdxs = nil
}
