// Code generated by protoc-gen-go. DO NOT EDIT.
// source: update.proto

package update

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpateData struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Version              float64  `protobuf:"fixed64,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpateData) Reset()         { *m = UpateData{} }
func (m *UpateData) String() string { return proto.CompactTextString(m) }
func (*UpateData) ProtoMessage()    {}
func (*UpateData) Descriptor() ([]byte, []int) {
	return fileDescriptor_update_0cbfaa24d988444e, []int{0}
}
func (m *UpateData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpateData.Unmarshal(m, b)
}
func (m *UpateData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpateData.Marshal(b, m, deterministic)
}
func (dst *UpateData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpateData.Merge(dst, src)
}
func (m *UpateData) XXX_Size() int {
	return xxx_messageInfo_UpateData.Size(m)
}
func (m *UpateData) XXX_DiscardUnknown() {
	xxx_messageInfo_UpateData.DiscardUnknown(m)
}

var xxx_messageInfo_UpateData proto.InternalMessageInfo

func (m *UpateData) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UpateData) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *UpateData) GetVersion() float64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type UpateDatas struct {
	Datas                []*UpateData `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpateDatas) Reset()         { *m = UpateDatas{} }
func (m *UpateDatas) String() string { return proto.CompactTextString(m) }
func (*UpateDatas) ProtoMessage()    {}
func (*UpateDatas) Descriptor() ([]byte, []int) {
	return fileDescriptor_update_0cbfaa24d988444e, []int{1}
}
func (m *UpateDatas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpateDatas.Unmarshal(m, b)
}
func (m *UpateDatas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpateDatas.Marshal(b, m, deterministic)
}
func (dst *UpateDatas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpateDatas.Merge(dst, src)
}
func (m *UpateDatas) XXX_Size() int {
	return xxx_messageInfo_UpateDatas.Size(m)
}
func (m *UpateDatas) XXX_DiscardUnknown() {
	xxx_messageInfo_UpateDatas.DiscardUnknown(m)
}

var xxx_messageInfo_UpateDatas proto.InternalMessageInfo

func (m *UpateDatas) GetDatas() []*UpateData {
	if m != nil {
		return m.Datas
	}
	return nil
}

func init() {
	proto.RegisterType((*UpateData)(nil), "update.upateData")
	proto.RegisterType((*UpateDatas)(nil), "update.upateDatas")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UpdateClient is the client API for Update service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UpdateClient interface {
	DoUpdateAll(ctx context.Context, opts ...grpc.CallOption) (Update_DoUpdateAllClient, error)
	DoUpdate(ctx context.Context, in *UpateData, opts ...grpc.CallOption) (*UpateData, error)
}

type updateClient struct {
	cc *grpc.ClientConn
}

func NewUpdateClient(cc *grpc.ClientConn) UpdateClient {
	return &updateClient{cc}
}

func (c *updateClient) DoUpdateAll(ctx context.Context, opts ...grpc.CallOption) (Update_DoUpdateAllClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Update_serviceDesc.Streams[0], "/update.Update/DoUpdateAll", opts...)
	if err != nil {
		return nil, err
	}
	x := &updateDoUpdateAllClient{stream}
	return x, nil
}

type Update_DoUpdateAllClient interface {
	Send(*UpateDatas) error
	Recv() (*UpateDatas, error)
	grpc.ClientStream
}

type updateDoUpdateAllClient struct {
	grpc.ClientStream
}

func (x *updateDoUpdateAllClient) Send(m *UpateDatas) error {
	return x.ClientStream.SendMsg(m)
}

func (x *updateDoUpdateAllClient) Recv() (*UpateDatas, error) {
	m := new(UpateDatas)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *updateClient) DoUpdate(ctx context.Context, in *UpateData, opts ...grpc.CallOption) (*UpateData, error) {
	out := new(UpateData)
	err := c.cc.Invoke(ctx, "/update.Update/DoUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateServer is the server API for Update service.
type UpdateServer interface {
	DoUpdateAll(Update_DoUpdateAllServer) error
	DoUpdate(context.Context, *UpateData) (*UpateData, error)
}

func RegisterUpdateServer(s *grpc.Server, srv UpdateServer) {
	s.RegisterService(&_Update_serviceDesc, srv)
}

func _Update_DoUpdateAll_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UpdateServer).DoUpdateAll(&updateDoUpdateAllServer{stream})
}

type Update_DoUpdateAllServer interface {
	Send(*UpateDatas) error
	Recv() (*UpateDatas, error)
	grpc.ServerStream
}

type updateDoUpdateAllServer struct {
	grpc.ServerStream
}

func (x *updateDoUpdateAllServer) Send(m *UpateDatas) error {
	return x.ServerStream.SendMsg(m)
}

func (x *updateDoUpdateAllServer) Recv() (*UpateDatas, error) {
	m := new(UpateDatas)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Update_DoUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpateData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UpdateServer).DoUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/update.Update/DoUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UpdateServer).DoUpdate(ctx, req.(*UpateData))
	}
	return interceptor(ctx, in, info, handler)
}

var _Update_serviceDesc = grpc.ServiceDesc{
	ServiceName: "update.Update",
	HandlerType: (*UpdateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoUpdate",
			Handler:    _Update_DoUpdate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoUpdateAll",
			Handler:       _Update_DoUpdateAll_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "update.proto",
}

func init() { proto.RegisterFile("update.proto", fileDescriptor_update_0cbfaa24d988444e) }

var fileDescriptor_update_0cbfaa24d988444e = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2d, 0x48, 0x49,
	0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x7c, 0xb9, 0x38,
	0x4b, 0x0b, 0x12, 0x4b, 0x52, 0x5d, 0x12, 0x4b, 0x12, 0x85, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b,
	0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c,
	0xd2, 0x54, 0x09, 0x26, 0xb0, 0x18, 0x84, 0x23, 0x24, 0xc1, 0xc5, 0x5e, 0x96, 0x5a, 0x54, 0x9c,
	0x99, 0x9f, 0x27, 0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x18, 0x04, 0xe3, 0x2a, 0x99, 0x72, 0x71, 0xc1,
	0x8d, 0x2b, 0x16, 0x52, 0xe7, 0x62, 0x4d, 0x01, 0x31, 0x24, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d,
	0x04, 0xf5, 0xa0, 0x4e, 0x80, 0x2b, 0x09, 0x82, 0xc8, 0x1b, 0x55, 0x72, 0xb1, 0x85, 0x82, 0xa5,
	0x84, 0xac, 0xb9, 0xb8, 0x5d, 0xf2, 0x21, 0x6c, 0xc7, 0x9c, 0x1c, 0x21, 0x21, 0x0c, 0x2d, 0xc5,
	0x52, 0x58, 0xc4, 0x94, 0x18, 0x34, 0x18, 0x0d, 0x18, 0x85, 0x8c, 0xb8, 0x38, 0x60, 0x9a, 0x85,
	0x30, 0x2d, 0x93, 0xc2, 0x14, 0x52, 0x62, 0x48, 0x62, 0x03, 0x87, 0x87, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xa6, 0x28, 0x18, 0xf1, 0x1f, 0x01, 0x00, 0x00,
}
