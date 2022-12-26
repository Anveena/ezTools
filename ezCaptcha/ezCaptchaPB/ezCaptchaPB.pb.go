// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.3
// source: ezCaptchaPB.proto

package ezCaptchaPB

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

type EZCaptchaRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Suc           bool   `protobuf:"varint,1,opt,name=suc,proto3" json:"suc,omitempty"`
	ErrDesc       string `protobuf:"bytes,2,opt,name=err_desc,json=errDesc,proto3" json:"err_desc,omitempty"`
	CorrectAnswer string `protobuf:"bytes,3,opt,name=correct_answer,json=correctAnswer,proto3" json:"correct_answer,omitempty"`
	PngBase64     string `protobuf:"bytes,4,opt,name=png_base64,json=pngBase64,proto3" json:"png_base64,omitempty"`
}

func (x *EZCaptchaRsp) Reset() {
	*x = EZCaptchaRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ezCaptchaPB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EZCaptchaRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EZCaptchaRsp) ProtoMessage() {}

func (x *EZCaptchaRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ezCaptchaPB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EZCaptchaRsp.ProtoReflect.Descriptor instead.
func (*EZCaptchaRsp) Descriptor() ([]byte, []int) {
	return file_ezCaptchaPB_proto_rawDescGZIP(), []int{0}
}

func (x *EZCaptchaRsp) GetSuc() bool {
	if x != nil {
		return x.Suc
	}
	return false
}

func (x *EZCaptchaRsp) GetErrDesc() string {
	if x != nil {
		return x.ErrDesc
	}
	return ""
}

func (x *EZCaptchaRsp) GetCorrectAnswer() string {
	if x != nil {
		return x.CorrectAnswer
	}
	return ""
}

func (x *EZCaptchaRsp) GetPngBase64() string {
	if x != nil {
		return x.PngBase64
	}
	return ""
}

type EZCaptchaEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EZCaptchaEmpty) Reset() {
	*x = EZCaptchaEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ezCaptchaPB_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EZCaptchaEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EZCaptchaEmpty) ProtoMessage() {}

func (x *EZCaptchaEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_ezCaptchaPB_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EZCaptchaEmpty.ProtoReflect.Descriptor instead.
func (*EZCaptchaEmpty) Descriptor() ([]byte, []int) {
	return file_ezCaptchaPB_proto_rawDescGZIP(), []int{1}
}

var File_ezCaptchaPB_proto protoreflect.FileDescriptor

var file_ezCaptchaPB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x7a, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x50, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x0c, 0x45, 0x5a, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68,
	0x61, 0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x03, 0x73, 0x75, 0x63, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x72, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65,
	0x63, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6e, 0x67, 0x5f,
	0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6e,
	0x67, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x22, 0x10, 0x0a, 0x0e, 0x45, 0x5a, 0x43, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x42, 0x0a, 0x10, 0x45, 0x5a, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x0f, 0x2e, 0x45, 0x5a,
	0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x45,
	0x5a, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x3b, 0x65, 0x7a, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x50, 0x42, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ezCaptchaPB_proto_rawDescOnce sync.Once
	file_ezCaptchaPB_proto_rawDescData = file_ezCaptchaPB_proto_rawDesc
)

func file_ezCaptchaPB_proto_rawDescGZIP() []byte {
	file_ezCaptchaPB_proto_rawDescOnce.Do(func() {
		file_ezCaptchaPB_proto_rawDescData = protoimpl.X.CompressGZIP(file_ezCaptchaPB_proto_rawDescData)
	})
	return file_ezCaptchaPB_proto_rawDescData
}

var file_ezCaptchaPB_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ezCaptchaPB_proto_goTypes = []any{
	(*EZCaptchaRsp)(nil),   // 0: EZCaptchaRsp
	(*EZCaptchaEmpty)(nil), // 1: EZCaptchaEmpty
}
var file_ezCaptchaPB_proto_depIdxs = []int32{
	1, // 0: EZCaptchaService.GetCaptcha:input_type -> EZCaptchaEmpty
	0, // 1: EZCaptchaService.GetCaptcha:output_type -> EZCaptchaRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ezCaptchaPB_proto_init() }
func file_ezCaptchaPB_proto_init() {
	if File_ezCaptchaPB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ezCaptchaPB_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EZCaptchaRsp); i {
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
		file_ezCaptchaPB_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EZCaptchaEmpty); i {
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
			RawDescriptor: file_ezCaptchaPB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ezCaptchaPB_proto_goTypes,
		DependencyIndexes: file_ezCaptchaPB_proto_depIdxs,
		MessageInfos:      file_ezCaptchaPB_proto_msgTypes,
	}.Build()
	File_ezCaptchaPB_proto = out.File
	file_ezCaptchaPB_proto_rawDesc = nil
	file_ezCaptchaPB_proto_goTypes = nil
	file_ezCaptchaPB_proto_depIdxs = nil
}
