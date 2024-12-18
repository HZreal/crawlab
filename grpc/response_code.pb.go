// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.2
// source: entity/response_code.proto

package grpc

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

type ResponseCode int32

const (
	ResponseCode_OK    ResponseCode = 0
	ResponseCode_ERROR ResponseCode = 1
)

// Enum value maps for ResponseCode.
var (
	ResponseCode_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	ResponseCode_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x ResponseCode) Enum() *ResponseCode {
	p := new(ResponseCode)
	*p = x
	return p
}

func (x ResponseCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseCode) Descriptor() protoreflect.EnumDescriptor {
	return file_entity_response_code_proto_enumTypes[0].Descriptor()
}

func (ResponseCode) Type() protoreflect.EnumType {
	return &file_entity_response_code_proto_enumTypes[0]
}

func (x ResponseCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseCode.Descriptor instead.
func (ResponseCode) EnumDescriptor() ([]byte, []int) {
	return file_entity_response_code_proto_rawDescGZIP(), []int{0}
}

var File_entity_response_code_proto protoreflect.FileDescriptor

var file_entity_response_code_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x21, 0x0a, 0x0c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02,
	0x4f, 0x4b, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x42,
	0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_entity_response_code_proto_rawDescOnce sync.Once
	file_entity_response_code_proto_rawDescData = file_entity_response_code_proto_rawDesc
)

func file_entity_response_code_proto_rawDescGZIP() []byte {
	file_entity_response_code_proto_rawDescOnce.Do(func() {
		file_entity_response_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_response_code_proto_rawDescData)
	})
	return file_entity_response_code_proto_rawDescData
}

var file_entity_response_code_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_entity_response_code_proto_goTypes = []any{
	(ResponseCode)(0), // 0: ResponseCode
}
var file_entity_response_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_entity_response_code_proto_init() }
func file_entity_response_code_proto_init() {
	if File_entity_response_code_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_entity_response_code_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_response_code_proto_goTypes,
		DependencyIndexes: file_entity_response_code_proto_depIdxs,
		EnumInfos:         file_entity_response_code_proto_enumTypes,
	}.Build()
	File_entity_response_code_proto = out.File
	file_entity_response_code_proto_rawDesc = nil
	file_entity_response_code_proto_goTypes = nil
	file_entity_response_code_proto_depIdxs = nil
}
