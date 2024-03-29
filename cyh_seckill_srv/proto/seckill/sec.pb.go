// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: sec.proto

package cyh_seckill_srv

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

type SecKillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *SecKillRequest) Reset() {
	*x = SecKillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecKillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecKillRequest) ProtoMessage() {}

func (x *SecKillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecKillRequest.ProtoReflect.Descriptor instead.
func (*SecKillRequest) Descriptor() ([]byte, []int) {
	return file_sec_proto_rawDescGZIP(), []int{0}
}

func (x *SecKillRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SecKillRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SecKillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SecKillResponse) Reset() {
	*x = SecKillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecKillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecKillResponse) ProtoMessage() {}

func (x *SecKillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecKillResponse.ProtoReflect.Descriptor instead.
func (*SecKillResponse) Descriptor() ([]byte, []int) {
	return file_sec_proto_rawDescGZIP(), []int{1}
}

func (x *SecKillResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SecKillResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_sec_proto protoreflect.FileDescriptor

var file_sec_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x79, 0x68,
	0x5f, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x73, 0x72, 0x76, 0x22, 0x3c, 0x0a, 0x0e,
	0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x0f, 0x53, 0x65,
	0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x32, 0x5e, 0x0a, 0x07, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x12, 0x53,
	0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x12, 0x1f,
	0x2e, 0x63, 0x79, 0x68, 0x5f, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x73, 0x72, 0x76,
	0x2e, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x63, 0x79, 0x68, 0x5f, 0x73, 0x65, 0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x73, 0x72,
	0x76, 0x2e, 0x53, 0x65, 0x63, 0x4b, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x2e, 0x2f, 0x3b, 0x63, 0x79, 0x68, 0x5f, 0x73, 0x65,
	0x63, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x73, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_sec_proto_rawDescOnce sync.Once
	file_sec_proto_rawDescData = file_sec_proto_rawDesc
)

func file_sec_proto_rawDescGZIP() []byte {
	file_sec_proto_rawDescOnce.Do(func() {
		file_sec_proto_rawDescData = protoimpl.X.CompressGZIP(file_sec_proto_rawDescData)
	})
	return file_sec_proto_rawDescData
}

var file_sec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_sec_proto_goTypes = []interface{}{
	(*SecKillRequest)(nil),  // 0: cyh_seckill_srv.SecKillRequest
	(*SecKillResponse)(nil), // 1: cyh_seckill_srv.SecKillResponse
}
var file_sec_proto_depIdxs = []int32{
	0, // 0: cyh_seckill_srv.SecKill.FrontSecKill:input_type -> cyh_seckill_srv.SecKillRequest
	1, // 1: cyh_seckill_srv.SecKill.FrontSecKill:output_type -> cyh_seckill_srv.SecKillResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sec_proto_init() }
func file_sec_proto_init() {
	if File_sec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecKillRequest); i {
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
		file_sec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecKillResponse); i {
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
			RawDescriptor: file_sec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sec_proto_goTypes,
		DependencyIndexes: file_sec_proto_depIdxs,
		MessageInfos:      file_sec_proto_msgTypes,
	}.Build()
	File_sec_proto = out.File
	file_sec_proto_rawDesc = nil
	file_sec_proto_goTypes = nil
	file_sec_proto_depIdxs = nil
}
