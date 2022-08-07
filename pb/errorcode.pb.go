// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: errorcode.proto

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

type ErrorCode int32

const (
	ErrorCode_INVALID                    ErrorCode = 0    //非法操作
	ErrorCode_SUCCESS                    ErrorCode = 1    // 操作成功
	ErrorCode_FAILED                     ErrorCode = 2    // 操作失败
	ErrorCode_SERVER_ERROR               ErrorCode = 3    // 服务器端错误
	ErrorCode_USERNAME_IS_NIL            ErrorCode = 1003 // 用户名为空
	ErrorCode_PASSWORD_IS_NIL            ErrorCode = 1004 // 密码为空
	ErrorCode_USERNAME_IS_USED           ErrorCode = 1005 // 用户名已被使用
	ErrorCode_USERNAME_OR_PASSWORD_WRONG ErrorCode = 1006 // 用户名或密码错误
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0:    "INVALID",
		1:    "SUCCESS",
		2:    "FAILED",
		3:    "SERVER_ERROR",
		1003: "USERNAME_IS_NIL",
		1004: "PASSWORD_IS_NIL",
		1005: "USERNAME_IS_USED",
		1006: "USERNAME_OR_PASSWORD_WRONG",
	}
	ErrorCode_value = map[string]int32{
		"INVALID":                    0,
		"SUCCESS":                    1,
		"FAILED":                     2,
		"SERVER_ERROR":               3,
		"USERNAME_IS_NIL":            1003,
		"PASSWORD_IS_NIL":            1004,
		"USERNAME_IS_USED":           1005,
		"USERNAME_OR_PASSWORD_WRONG": 1006,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_errorcode_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_errorcode_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_errorcode_proto_rawDescGZIP(), []int{0}
}

var File_errorcode_proto protoreflect.FileDescriptor

var file_errorcode_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0xa7, 0x01, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x4e,
	0x41, 0x4d, 0x45, 0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x49, 0x4c, 0x10, 0xeb, 0x07, 0x12, 0x14, 0x0a,
	0x0f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x49, 0x53, 0x5f, 0x4e, 0x49, 0x4c,
	0x10, 0xec, 0x07, 0x12, 0x15, 0x0a, 0x10, 0x55, 0x53, 0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x5f,
	0x49, 0x53, 0x5f, 0x55, 0x53, 0x45, 0x44, 0x10, 0xed, 0x07, 0x12, 0x1f, 0x0a, 0x1a, 0x55, 0x53,
	0x45, 0x52, 0x4e, 0x41, 0x4d, 0x45, 0x5f, 0x4f, 0x52, 0x5f, 0x50, 0x41, 0x53, 0x53, 0x57, 0x4f,
	0x52, 0x44, 0x5f, 0x57, 0x52, 0x4f, 0x4e, 0x47, 0x10, 0xee, 0x07, 0x42, 0x07, 0x5a, 0x05, 0x2e,
	0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errorcode_proto_rawDescOnce sync.Once
	file_errorcode_proto_rawDescData = file_errorcode_proto_rawDesc
)

func file_errorcode_proto_rawDescGZIP() []byte {
	file_errorcode_proto_rawDescOnce.Do(func() {
		file_errorcode_proto_rawDescData = protoimpl.X.CompressGZIP(file_errorcode_proto_rawDescData)
	})
	return file_errorcode_proto_rawDescData
}

var file_errorcode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errorcode_proto_goTypes = []interface{}{
	(ErrorCode)(0), // 0: ErrorCode
}
var file_errorcode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errorcode_proto_init() }
func file_errorcode_proto_init() {
	if File_errorcode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errorcode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errorcode_proto_goTypes,
		DependencyIndexes: file_errorcode_proto_depIdxs,
		EnumInfos:         file_errorcode_proto_enumTypes,
	}.Build()
	File_errorcode_proto = out.File
	file_errorcode_proto_rawDesc = nil
	file_errorcode_proto_goTypes = nil
	file_errorcode_proto_depIdxs = nil
}
