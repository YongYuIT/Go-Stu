// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.2
// source: HelloServerStream.proto

package HelloServerStreamPkg

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ClassInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GradeName string `protobuf:"bytes,1,opt,name=gradeName,proto3" json:"gradeName,omitempty"` //在消息定义中，每个字段都有唯一的一个数字标识符。这些标识符是用来在消息的二进制格式中识别各个字段的，一旦开始使用就不能够再改变。
	ClassName string `protobuf:"bytes,2,opt,name=className,proto3" json:"className,omitempty"`
}

func (x *ClassInfo) Reset() {
	*x = ClassInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HelloServerStream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassInfo) ProtoMessage() {}

func (x *ClassInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HelloServerStream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassInfo.ProtoReflect.Descriptor instead.
func (*ClassInfo) Descriptor() ([]byte, []int) {
	return file_HelloServerStream_proto_rawDescGZIP(), []int{0}
}

func (x *ClassInfo) GetGradeName() string {
	if x != nil {
		return x.GradeName
	}
	return ""
}

func (x *ClassInfo) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

type StuInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StuName    string `protobuf:"bytes,1,opt,name=stuName,proto3" json:"stuName,omitempty"`
	StuGender  string `protobuf:"bytes,2,opt,name=stuGender,proto3" json:"stuGender,omitempty"`
	StuHomeAdd string `protobuf:"bytes,3,opt,name=stuHomeAdd,proto3" json:"stuHomeAdd,omitempty"`
}

func (x *StuInfo) Reset() {
	*x = StuInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HelloServerStream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StuInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StuInfo) ProtoMessage() {}

func (x *StuInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HelloServerStream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StuInfo.ProtoReflect.Descriptor instead.
func (*StuInfo) Descriptor() ([]byte, []int) {
	return file_HelloServerStream_proto_rawDescGZIP(), []int{1}
}

func (x *StuInfo) GetStuName() string {
	if x != nil {
		return x.StuName
	}
	return ""
}

func (x *StuInfo) GetStuGender() string {
	if x != nil {
		return x.StuGender
	}
	return ""
}

func (x *StuInfo) GetStuHomeAdd() string {
	if x != nil {
		return x.StuHomeAdd
	}
	return ""
}

var File_HelloServerStream_proto protoreflect.FileDescriptor

var file_HelloServerStream_proto_rawDesc = []byte{
	0x0a, 0x17, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x6b, 0x67, 0x22,
	0x47, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x09,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x67, 0x72, 0x61, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x61, 0x0a, 0x07, 0x53, 0x74, 0x75, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x75, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x74, 0x75, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x73,
	0x74, 0x75, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x73, 0x74, 0x75, 0x48, 0x6f, 0x6d, 0x65, 0x41, 0x64, 0x64, 0x32, 0x70, 0x0a, 0x14, 0x53,
	0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x53, 0x74, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x12, 0x67, 0x65, 0x74, 0x53, 0x74, 0x75, 0x73, 0x42, 0x79,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x6b, 0x67,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x1d, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x50, 0x6b,
	0x67, 0x2e, 0x53, 0x74, 0x75, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HelloServerStream_proto_rawDescOnce sync.Once
	file_HelloServerStream_proto_rawDescData = file_HelloServerStream_proto_rawDesc
)

func file_HelloServerStream_proto_rawDescGZIP() []byte {
	file_HelloServerStream_proto_rawDescOnce.Do(func() {
		file_HelloServerStream_proto_rawDescData = protoimpl.X.CompressGZIP(file_HelloServerStream_proto_rawDescData)
	})
	return file_HelloServerStream_proto_rawDescData
}

var file_HelloServerStream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_HelloServerStream_proto_goTypes = []interface{}{
	(*ClassInfo)(nil), // 0: HelloServerStreamPkg.ClassInfo
	(*StuInfo)(nil),   // 1: HelloServerStreamPkg.StuInfo
}
var file_HelloServerStream_proto_depIdxs = []int32{
	0, // 0: HelloServerStreamPkg.SchoolStuInfoService.getStusByClassInfo:input_type -> HelloServerStreamPkg.ClassInfo
	1, // 1: HelloServerStreamPkg.SchoolStuInfoService.getStusByClassInfo:output_type -> HelloServerStreamPkg.StuInfo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HelloServerStream_proto_init() }
func file_HelloServerStream_proto_init() {
	if File_HelloServerStream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HelloServerStream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassInfo); i {
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
		file_HelloServerStream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StuInfo); i {
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
			RawDescriptor: file_HelloServerStream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_HelloServerStream_proto_goTypes,
		DependencyIndexes: file_HelloServerStream_proto_depIdxs,
		MessageInfos:      file_HelloServerStream_proto_msgTypes,
	}.Build()
	File_HelloServerStream_proto = out.File
	file_HelloServerStream_proto_rawDesc = nil
	file_HelloServerStream_proto_goTypes = nil
	file_HelloServerStream_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SchoolStuInfoServiceClient is the client API for SchoolStuInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchoolStuInfoServiceClient interface {
	//根据班级id获取所有学生名单
	GetStusByClassInfo(ctx context.Context, in *ClassInfo, opts ...grpc.CallOption) (SchoolStuInfoService_GetStusByClassInfoClient, error)
}

type schoolStuInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSchoolStuInfoServiceClient(cc grpc.ClientConnInterface) SchoolStuInfoServiceClient {
	return &schoolStuInfoServiceClient{cc}
}

func (c *schoolStuInfoServiceClient) GetStusByClassInfo(ctx context.Context, in *ClassInfo, opts ...grpc.CallOption) (SchoolStuInfoService_GetStusByClassInfoClient, error) {
	stream, err := c.cc.NewStream(ctx, &_SchoolStuInfoService_serviceDesc.Streams[0], "/HelloServerStreamPkg.SchoolStuInfoService/getStusByClassInfo", opts...)
	if err != nil {
		return nil, err
	}
	x := &schoolStuInfoServiceGetStusByClassInfoClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SchoolStuInfoService_GetStusByClassInfoClient interface {
	Recv() (*StuInfo, error)
	grpc.ClientStream
}

type schoolStuInfoServiceGetStusByClassInfoClient struct {
	grpc.ClientStream
}

func (x *schoolStuInfoServiceGetStusByClassInfoClient) Recv() (*StuInfo, error) {
	m := new(StuInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SchoolStuInfoServiceServer is the server API for SchoolStuInfoService service.
type SchoolStuInfoServiceServer interface {
	//根据班级id获取所有学生名单
	GetStusByClassInfo(*ClassInfo, SchoolStuInfoService_GetStusByClassInfoServer) error
}

// UnimplementedSchoolStuInfoServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSchoolStuInfoServiceServer struct {
}

func (*UnimplementedSchoolStuInfoServiceServer) GetStusByClassInfo(*ClassInfo, SchoolStuInfoService_GetStusByClassInfoServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStusByClassInfo not implemented")
}

func RegisterSchoolStuInfoServiceServer(s *grpc.Server, srv SchoolStuInfoServiceServer) {
	s.RegisterService(&_SchoolStuInfoService_serviceDesc, srv)
}

func _SchoolStuInfoService_GetStusByClassInfo_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClassInfo)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SchoolStuInfoServiceServer).GetStusByClassInfo(m, &schoolStuInfoServiceGetStusByClassInfoServer{stream})
}

type SchoolStuInfoService_GetStusByClassInfoServer interface {
	Send(*StuInfo) error
	grpc.ServerStream
}

type schoolStuInfoServiceGetStusByClassInfoServer struct {
	grpc.ServerStream
}

func (x *schoolStuInfoServiceGetStusByClassInfoServer) Send(m *StuInfo) error {
	return x.ServerStream.SendMsg(m)
}

var _SchoolStuInfoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "HelloServerStreamPkg.SchoolStuInfoService",
	HandlerType: (*SchoolStuInfoServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getStusByClassInfo",
			Handler:       _SchoolStuInfoService_GetStusByClassInfo_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "HelloServerStream.proto",
}
