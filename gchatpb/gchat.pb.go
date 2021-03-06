// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: gchatpb/gchat.proto

package gchatpb

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

type ChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size    int32  `protobuf:"varint,1,opt,name=Size,proto3" json:"Size,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gchatpb_gchat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_gchatpb_gchat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_gchatpb_gchat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatMessage) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ChatMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_gchatpb_gchat_proto protoreflect.FileDescriptor

var file_gchatpb_gchat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2f, 0x67, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x22, 0x3b,
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x4b, 0x0a, 0x05, 0x47,
	0x43, 0x68, 0x61, 0x74, 0x12, 0x42, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x41,
	0x6e, 0x64, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x2e, 0x67, 0x63, 0x68, 0x61, 0x74, 0x70, 0x62,
	0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x67,
	0x63, 0x68, 0x61, 0x74, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x67, 0x63,
	0x68, 0x61, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gchatpb_gchat_proto_rawDescOnce sync.Once
	file_gchatpb_gchat_proto_rawDescData = file_gchatpb_gchat_proto_rawDesc
)

func file_gchatpb_gchat_proto_rawDescGZIP() []byte {
	file_gchatpb_gchat_proto_rawDescOnce.Do(func() {
		file_gchatpb_gchat_proto_rawDescData = protoimpl.X.CompressGZIP(file_gchatpb_gchat_proto_rawDescData)
	})
	return file_gchatpb_gchat_proto_rawDescData
}

var file_gchatpb_gchat_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gchatpb_gchat_proto_goTypes = []interface{}{
	(*ChatMessage)(nil), // 0: gchatpb.ChatMessage
}
var file_gchatpb_gchat_proto_depIdxs = []int32{
	0, // 0: gchatpb.GChat.ReceiveAndSend:input_type -> gchatpb.ChatMessage
	0, // 1: gchatpb.GChat.ReceiveAndSend:output_type -> gchatpb.ChatMessage
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gchatpb_gchat_proto_init() }
func file_gchatpb_gchat_proto_init() {
	if File_gchatpb_gchat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gchatpb_gchat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatMessage); i {
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
			RawDescriptor: file_gchatpb_gchat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gchatpb_gchat_proto_goTypes,
		DependencyIndexes: file_gchatpb_gchat_proto_depIdxs,
		MessageInfos:      file_gchatpb_gchat_proto_msgTypes,
	}.Build()
	File_gchatpb_gchat_proto = out.File
	file_gchatpb_gchat_proto_rawDesc = nil
	file_gchatpb_gchat_proto_goTypes = nil
	file_gchatpb_gchat_proto_depIdxs = nil
}
