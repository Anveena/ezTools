// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: ezlogPB.proto

package ezLogPB

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EZLogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// https://developers.google.com/protocol-buffers/docs/proto3#scalar 大老师没有一种能对应go的int的.操
	Level    int32                  `protobuf:"varint,1,opt,name=level,proto3" json:"level,omitempty"`
	FileLine int32                  `protobuf:"varint,2,opt,name=file_line,json=fileLine,proto3" json:"file_line,omitempty"`
	Time     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	FileName string                 `protobuf:"bytes,4,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	AppName  string                 `protobuf:"bytes,5,opt,name=app_name,json=appName,proto3" json:"app_name,omitempty"`
	Tag      string                 `protobuf:"bytes,6,opt,name=tag,proto3" json:"tag,omitempty"`
	Content  string                 `protobuf:"bytes,7,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *EZLogReq) Reset() {
	*x = EZLogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ezlogPB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EZLogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EZLogReq) ProtoMessage() {}

func (x *EZLogReq) ProtoReflect() protoreflect.Message {
	mi := &file_ezlogPB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EZLogReq.ProtoReflect.Descriptor instead.
func (*EZLogReq) Descriptor() ([]byte, []int) {
	return file_ezlogPB_proto_rawDescGZIP(), []int{0}
}

func (x *EZLogReq) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *EZLogReq) GetFileLine() int32 {
	if x != nil {
		return x.FileLine
	}
	return 0
}

func (x *EZLogReq) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *EZLogReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *EZLogReq) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *EZLogReq) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

func (x *EZLogReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type EZLogEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EZLogEmpty) Reset() {
	*x = EZLogEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ezlogPB_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EZLogEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EZLogEmpty) ProtoMessage() {}

func (x *EZLogEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_ezlogPB_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EZLogEmpty.ProtoReflect.Descriptor instead.
func (*EZLogEmpty) Descriptor() ([]byte, []int) {
	return file_ezlogPB_proto_rawDescGZIP(), []int{1}
}

var File_ezlogPB_proto protoreflect.FileDescriptor

var file_ezlogPB_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x7a, 0x6c, 0x6f, 0x67, 0x50, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xd1, 0x01, 0x0a, 0x08, 0x45, 0x5a, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6c, 0x69, 0x6e, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4c, 0x69, 0x6e, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x5a, 0x4c, 0x6f, 0x67, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x32, 0x2e, 0x0a, 0x09, 0x45, 0x7a, 0x4c, 0x6f, 0x67, 0x47, 0x72, 0x70, 0x63, 0x12,
	0x21, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x09, 0x2e, 0x45, 0x5a, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x0b, 0x2e, 0x45, 0x5a, 0x4c, 0x6f, 0x67, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x28, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x3b, 0x65, 0x7a, 0x4c, 0x6f, 0x67, 0x50, 0x42, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ezlogPB_proto_rawDescOnce sync.Once
	file_ezlogPB_proto_rawDescData = file_ezlogPB_proto_rawDesc
)

func file_ezlogPB_proto_rawDescGZIP() []byte {
	file_ezlogPB_proto_rawDescOnce.Do(func() {
		file_ezlogPB_proto_rawDescData = protoimpl.X.CompressGZIP(file_ezlogPB_proto_rawDescData)
	})
	return file_ezlogPB_proto_rawDescData
}

var file_ezlogPB_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ezlogPB_proto_goTypes = []interface{}{
	(*EZLogReq)(nil),              // 0: EZLogReq
	(*EZLogEmpty)(nil),            // 1: EZLogEmpty
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_ezlogPB_proto_depIdxs = []int32{
	2, // 0: EZLogReq.time:type_name -> google.protobuf.Timestamp
	0, // 1: EzLogGrpc.Log:input_type -> EZLogReq
	1, // 2: EzLogGrpc.Log:output_type -> EZLogEmpty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ezlogPB_proto_init() }
func file_ezlogPB_proto_init() {
	if File_ezlogPB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ezlogPB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EZLogReq); i {
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
		file_ezlogPB_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EZLogEmpty); i {
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
			RawDescriptor: file_ezlogPB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ezlogPB_proto_goTypes,
		DependencyIndexes: file_ezlogPB_proto_depIdxs,
		MessageInfos:      file_ezlogPB_proto_msgTypes,
	}.Build()
	File_ezlogPB_proto = out.File
	file_ezlogPB_proto_rawDesc = nil
	file_ezlogPB_proto_goTypes = nil
	file_ezlogPB_proto_depIdxs = nil
}
