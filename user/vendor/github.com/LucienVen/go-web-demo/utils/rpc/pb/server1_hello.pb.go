// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: server1_hello.proto

package pb

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

type SayHelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,1,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"` // 服务名称
}

func (x *SayHelloReq) Reset() {
	*x = SayHelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server1_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloReq) ProtoMessage() {}

func (x *SayHelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_server1_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloReq.ProtoReflect.Descriptor instead.
func (*SayHelloReq) Descriptor() ([]byte, []int) {
	return file_server1_hello_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloReq) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

type SayHelloRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Greet string `protobuf:"bytes,1,opt,name=greet,proto3" json:"greet,omitempty"` // 问候语
}

func (x *SayHelloRes) Reset() {
	*x = SayHelloRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server1_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRes) ProtoMessage() {}

func (x *SayHelloRes) ProtoReflect() protoreflect.Message {
	mi := &file_server1_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRes.ProtoReflect.Descriptor instead.
func (*SayHelloRes) Descriptor() ([]byte, []int) {
	return file_server1_hello_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloRes) GetGreet() string {
	if x != nil {
		return x.Greet
	}
	return ""
}

var File_server1_hello_proto protoreflect.FileDescriptor

var file_server1_hello_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x31, 0x5f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x31, 0x22, 0x2e,
	0x0a, 0x0b, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x23,
	0x0a, 0x0b, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x72,
	0x65, 0x65, 0x74, 0x32, 0x41, 0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x38, 0x0a, 0x08,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x14,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x75, 0x63, 0x69, 0x65, 0x6e, 0x56, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2d, 0x77, 0x65, 0x62, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server1_hello_proto_rawDescOnce sync.Once
	file_server1_hello_proto_rawDescData = file_server1_hello_proto_rawDesc
)

func file_server1_hello_proto_rawDescGZIP() []byte {
	file_server1_hello_proto_rawDescOnce.Do(func() {
		file_server1_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_server1_hello_proto_rawDescData)
	})
	return file_server1_hello_proto_rawDescData
}

var file_server1_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server1_hello_proto_goTypes = []interface{}{
	(*SayHelloReq)(nil), // 0: user.SayHelloReq
	(*SayHelloRes)(nil), // 1: user.SayHelloRes
}
var file_server1_hello_proto_depIdxs = []int32{
	0, // 0: user.Hello.SayHello:input_type -> user.SayHelloReq
	1, // 1: user.Hello.SayHello:output_type -> user.SayHelloRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server1_hello_proto_init() }
func file_server1_hello_proto_init() {
	if File_server1_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server1_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloReq); i {
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
		file_server1_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRes); i {
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
			RawDescriptor: file_server1_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server1_hello_proto_goTypes,
		DependencyIndexes: file_server1_hello_proto_depIdxs,
		MessageInfos:      file_server1_hello_proto_msgTypes,
	}.Build()
	File_server1_hello_proto = out.File
	file_server1_hello_proto_rawDesc = nil
	file_server1_hello_proto_goTypes = nil
	file_server1_hello_proto_depIdxs = nil
}