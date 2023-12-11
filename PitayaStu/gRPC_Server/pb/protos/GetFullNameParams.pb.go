// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: proto_src/GetFullNameParams.proto

package protos

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

type GetFullNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetFullNameRequest) Reset() {
	*x = GetFullNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_src_GetFullNameParams_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullNameRequest) ProtoMessage() {}

func (x *GetFullNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_src_GetFullNameParams_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullNameRequest.ProtoReflect.Descriptor instead.
func (*GetFullNameRequest) Descriptor() ([]byte, []int) {
	return file_proto_src_GetFullNameParams_proto_rawDescGZIP(), []int{0}
}

func (x *GetFullNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetFullNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullName string `protobuf:"bytes,1,opt,name=fullName,proto3" json:"fullName,omitempty"`
}

func (x *GetFullNameResponse) Reset() {
	*x = GetFullNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_src_GetFullNameParams_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFullNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFullNameResponse) ProtoMessage() {}

func (x *GetFullNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_src_GetFullNameParams_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFullNameResponse.ProtoReflect.Descriptor instead.
func (*GetFullNameResponse) Descriptor() ([]byte, []int) {
	return file_proto_src_GetFullNameParams_proto_rawDescGZIP(), []int{1}
}

func (x *GetFullNameResponse) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

var File_proto_src_GetFullNameParams_proto protoreflect.FileDescriptor

var file_proto_src_GetFullNameParams_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x73, 0x72, 0x63, 0x2f, 0x47, 0x65, 0x74, 0x46,
	0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_src_GetFullNameParams_proto_rawDescOnce sync.Once
	file_proto_src_GetFullNameParams_proto_rawDescData = file_proto_src_GetFullNameParams_proto_rawDesc
)

func file_proto_src_GetFullNameParams_proto_rawDescGZIP() []byte {
	file_proto_src_GetFullNameParams_proto_rawDescOnce.Do(func() {
		file_proto_src_GetFullNameParams_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_src_GetFullNameParams_proto_rawDescData)
	})
	return file_proto_src_GetFullNameParams_proto_rawDescData
}

var file_proto_src_GetFullNameParams_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_src_GetFullNameParams_proto_goTypes = []interface{}{
	(*GetFullNameRequest)(nil),  // 0: GetFullNameRequest
	(*GetFullNameResponse)(nil), // 1: GetFullNameResponse
}
var file_proto_src_GetFullNameParams_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_src_GetFullNameParams_proto_init() }
func file_proto_src_GetFullNameParams_proto_init() {
	if File_proto_src_GetFullNameParams_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_src_GetFullNameParams_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullNameRequest); i {
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
		file_proto_src_GetFullNameParams_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFullNameResponse); i {
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
			RawDescriptor: file_proto_src_GetFullNameParams_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_src_GetFullNameParams_proto_goTypes,
		DependencyIndexes: file_proto_src_GetFullNameParams_proto_depIdxs,
		MessageInfos:      file_proto_src_GetFullNameParams_proto_msgTypes,
	}.Build()
	File_proto_src_GetFullNameParams_proto = out.File
	file_proto_src_GetFullNameParams_proto_rawDesc = nil
	file_proto_src_GetFullNameParams_proto_goTypes = nil
	file_proto_src_GetFullNameParams_proto_depIdxs = nil
}
