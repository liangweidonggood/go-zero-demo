// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package common

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type BaseAppReq struct {
	Beid                 int64    `protobuf:"varint,1,opt,name=beid,proto3" json:"beid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseAppReq) Reset()         { *m = BaseAppReq{} }
func (m *BaseAppReq) String() string { return proto.CompactTextString(m) }
func (*BaseAppReq) ProtoMessage()    {}
func (*BaseAppReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *BaseAppReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseAppReq.Unmarshal(m, b)
}
func (m *BaseAppReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseAppReq.Marshal(b, m, deterministic)
}
func (m *BaseAppReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseAppReq.Merge(m, src)
}
func (m *BaseAppReq) XXX_Size() int {
	return xxx_messageInfo_BaseAppReq.Size(m)
}
func (m *BaseAppReq) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseAppReq.DiscardUnknown(m)
}

var xxx_messageInfo_BaseAppReq proto.InternalMessageInfo

func (m *BaseAppReq) GetBeid() int64 {
	if m != nil {
		return m.Beid
	}
	return 0
}

type BaseAppResp struct {
	Beid                 int64    `protobuf:"varint,1,opt,name=beid,proto3" json:"beid,omitempty"`
	Logo                 string   `protobuf:"bytes,2,opt,name=logo,proto3" json:"logo,omitempty"`
	Sname                string   `protobuf:"bytes,3,opt,name=sname,proto3" json:"sname,omitempty"`
	Isclose              int64    `protobuf:"varint,4,opt,name=isclose,proto3" json:"isclose,omitempty"`
	Fullwebsite          string   `protobuf:"bytes,5,opt,name=fullwebsite,proto3" json:"fullwebsite,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseAppResp) Reset()         { *m = BaseAppResp{} }
func (m *BaseAppResp) String() string { return proto.CompactTextString(m) }
func (*BaseAppResp) ProtoMessage()    {}
func (*BaseAppResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *BaseAppResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseAppResp.Unmarshal(m, b)
}
func (m *BaseAppResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseAppResp.Marshal(b, m, deterministic)
}
func (m *BaseAppResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseAppResp.Merge(m, src)
}
func (m *BaseAppResp) XXX_Size() int {
	return xxx_messageInfo_BaseAppResp.Size(m)
}
func (m *BaseAppResp) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseAppResp.DiscardUnknown(m)
}

var xxx_messageInfo_BaseAppResp proto.InternalMessageInfo

func (m *BaseAppResp) GetBeid() int64 {
	if m != nil {
		return m.Beid
	}
	return 0
}

func (m *BaseAppResp) GetLogo() string {
	if m != nil {
		return m.Logo
	}
	return ""
}

func (m *BaseAppResp) GetSname() string {
	if m != nil {
		return m.Sname
	}
	return ""
}

func (m *BaseAppResp) GetIsclose() int64 {
	if m != nil {
		return m.Isclose
	}
	return 0
}

func (m *BaseAppResp) GetFullwebsite() string {
	if m != nil {
		return m.Fullwebsite
	}
	return ""
}

//请求的api
type AppConfigReq struct {
	Beid                 int64    `protobuf:"varint,1,opt,name=beid,proto3" json:"beid,omitempty"`
	Ptyid                int64    `protobuf:"varint,2,opt,name=ptyid,proto3" json:"ptyid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppConfigReq) Reset()         { *m = AppConfigReq{} }
func (m *AppConfigReq) String() string { return proto.CompactTextString(m) }
func (*AppConfigReq) ProtoMessage()    {}
func (*AppConfigReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *AppConfigReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppConfigReq.Unmarshal(m, b)
}
func (m *AppConfigReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppConfigReq.Marshal(b, m, deterministic)
}
func (m *AppConfigReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppConfigReq.Merge(m, src)
}
func (m *AppConfigReq) XXX_Size() int {
	return xxx_messageInfo_AppConfigReq.Size(m)
}
func (m *AppConfigReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AppConfigReq.DiscardUnknown(m)
}

var xxx_messageInfo_AppConfigReq proto.InternalMessageInfo

func (m *AppConfigReq) GetBeid() int64 {
	if m != nil {
		return m.Beid
	}
	return 0
}

func (m *AppConfigReq) GetPtyid() int64 {
	if m != nil {
		return m.Ptyid
	}
	return 0
}

//返回的值
type AppConfigResp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Beid                 int64    `protobuf:"varint,2,opt,name=beid,proto3" json:"beid,omitempty"`
	Ptyid                int64    `protobuf:"varint,3,opt,name=ptyid,proto3" json:"ptyid,omitempty"`
	Appid                string   `protobuf:"bytes,4,opt,name=appid,proto3" json:"appid,omitempty"`
	Appsecret            string   `protobuf:"bytes,5,opt,name=appsecret,proto3" json:"appsecret,omitempty"`
	Title                string   `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AppConfigResp) Reset()         { *m = AppConfigResp{} }
func (m *AppConfigResp) String() string { return proto.CompactTextString(m) }
func (*AppConfigResp) ProtoMessage()    {}
func (*AppConfigResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *AppConfigResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AppConfigResp.Unmarshal(m, b)
}
func (m *AppConfigResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AppConfigResp.Marshal(b, m, deterministic)
}
func (m *AppConfigResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AppConfigResp.Merge(m, src)
}
func (m *AppConfigResp) XXX_Size() int {
	return xxx_messageInfo_AppConfigResp.Size(m)
}
func (m *AppConfigResp) XXX_DiscardUnknown() {
	xxx_messageInfo_AppConfigResp.DiscardUnknown(m)
}

var xxx_messageInfo_AppConfigResp proto.InternalMessageInfo

func (m *AppConfigResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AppConfigResp) GetBeid() int64 {
	if m != nil {
		return m.Beid
	}
	return 0
}

func (m *AppConfigResp) GetPtyid() int64 {
	if m != nil {
		return m.Ptyid
	}
	return 0
}

func (m *AppConfigResp) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *AppConfigResp) GetAppsecret() string {
	if m != nil {
		return m.Appsecret
	}
	return ""
}

func (m *AppConfigResp) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseAppReq)(nil), "common.BaseAppReq")
	proto.RegisterType((*BaseAppResp)(nil), "common.BaseAppResp")
	proto.RegisterType((*AppConfigReq)(nil), "common.AppConfigReq")
	proto.RegisterType((*AppConfigResp)(nil), "common.AppConfigResp")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x25, 0x49, 0x1b, 0xc9, 0x34, 0x7a, 0x18, 0x2b, 0x2c, 0xc5, 0x43, 0xc8, 0xa9, 0xa7, 0x1e,
	0x14, 0x41, 0xf0, 0x54, 0x7b, 0xe8, 0x3d, 0xff, 0x20, 0x1f, 0xd3, 0xb2, 0x90, 0x64, 0xc7, 0xee,
	0x8a, 0x08, 0xde, 0xfd, 0x01, 0xfe, 0x61, 0xc9, 0x6e, 0x9a, 0x06, 0x8c, 0xb7, 0x79, 0x8f, 0x79,
	0x6f, 0xdf, 0xf0, 0x16, 0xe2, 0x52, 0x35, 0x8d, 0x6a, 0x37, 0x7c, 0x52, 0x46, 0x61, 0xe8, 0x50,
	0x9a, 0x00, 0xbc, 0xe6, 0x9a, 0xb6, 0xcc, 0x19, 0xbd, 0x21, 0xc2, 0xac, 0x20, 0x59, 0x09, 0x2f,
	0xf1, 0xd6, 0x41, 0x66, 0xe7, 0xf4, 0xdb, 0x83, 0xc5, 0xb0, 0xa2, 0x79, 0x6a, 0xa7, 0xe3, 0x6a,
	0x75, 0x54, 0xc2, 0x4f, 0xbc, 0x75, 0x94, 0xd9, 0x19, 0x97, 0x30, 0xd7, 0x6d, 0xde, 0x90, 0x08,
	0x2c, 0xe9, 0x00, 0x0a, 0xb8, 0x92, 0xba, 0xac, 0x95, 0x26, 0x31, 0xb3, 0x06, 0x67, 0x88, 0x09,
	0x2c, 0x0e, 0xef, 0x75, 0xfd, 0x41, 0x85, 0x96, 0x86, 0xc4, 0xdc, 0xaa, 0xc6, 0x54, 0xfa, 0x0c,
	0xf1, 0x96, 0x79, 0xa7, 0xda, 0x83, 0x3c, 0xfe, 0x93, 0xb6, 0x7b, 0x95, 0xcd, 0xa7, 0xac, 0x6c,
	0x94, 0x20, 0x73, 0x20, 0xfd, 0xf1, 0xe0, 0x7a, 0x24, 0xd5, 0x8c, 0x37, 0xe0, 0x0f, 0x4a, 0xdf,
	0x5d, 0x60, 0xbd, 0xfc, 0x29, 0xaf, 0x60, 0xe4, 0xd5, 0xb1, 0x39, 0xb3, 0xac, 0x6c, 0xfe, 0x28,
	0x73, 0x00, 0xef, 0x21, 0xca, 0x99, 0x35, 0x95, 0x27, 0x32, 0x7d, 0xf6, 0x0b, 0xd1, 0x69, 0x8c,
	0x34, 0x35, 0x89, 0xd0, 0x69, 0x2c, 0x78, 0xf8, 0x82, 0x70, 0x67, 0x5b, 0xc0, 0x17, 0x88, 0xf7,
	0x64, 0x86, 0x84, 0xb8, 0xdc, 0xf4, 0x65, 0x8d, 0xef, 0x5d, 0xdd, 0x4d, 0xb0, 0x9a, 0xf1, 0x09,
	0x60, 0x4f, 0xa6, 0xaf, 0x08, 0xf1, 0xbc, 0x74, 0xa9, 0x75, 0x75, 0xfb, 0x87, 0xd3, 0x5c, 0x84,
	0xf6, 0x23, 0x3c, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xed, 0xd7, 0xc7, 0x41, 0x18, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CommonClient is the client API for Common service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CommonClient interface {
	GetAppConfig(ctx context.Context, in *AppConfigReq, opts ...grpc.CallOption) (*AppConfigResp, error)
	GetBaseApp(ctx context.Context, in *BaseAppReq, opts ...grpc.CallOption) (*BaseAppResp, error)
}

type commonClient struct {
	cc *grpc.ClientConn
}

func NewCommonClient(cc *grpc.ClientConn) CommonClient {
	return &commonClient{cc}
}

func (c *commonClient) GetAppConfig(ctx context.Context, in *AppConfigReq, opts ...grpc.CallOption) (*AppConfigResp, error) {
	out := new(AppConfigResp)
	err := c.cc.Invoke(ctx, "/common.Common/GetAppConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commonClient) GetBaseApp(ctx context.Context, in *BaseAppReq, opts ...grpc.CallOption) (*BaseAppResp, error) {
	out := new(BaseAppResp)
	err := c.cc.Invoke(ctx, "/common.Common/GetBaseApp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommonServer is the server API for Common service.
type CommonServer interface {
	GetAppConfig(context.Context, *AppConfigReq) (*AppConfigResp, error)
	GetBaseApp(context.Context, *BaseAppReq) (*BaseAppResp, error)
}

// UnimplementedCommonServer can be embedded to have forward compatible implementations.
type UnimplementedCommonServer struct {
}

func (*UnimplementedCommonServer) GetAppConfig(ctx context.Context, req *AppConfigReq) (*AppConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAppConfig not implemented")
}
func (*UnimplementedCommonServer) GetBaseApp(ctx context.Context, req *BaseAppReq) (*BaseAppResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBaseApp not implemented")
}

func RegisterCommonServer(s *grpc.Server, srv CommonServer) {
	s.RegisterService(&_Common_serviceDesc, srv)
}

func _Common_GetAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).GetAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.Common/GetAppConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).GetAppConfig(ctx, req.(*AppConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Common_GetBaseApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BaseAppReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommonServer).GetBaseApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/common.Common/GetBaseApp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommonServer).GetBaseApp(ctx, req.(*BaseAppReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Common_serviceDesc = grpc.ServiceDesc{
	ServiceName: "common.Common",
	HandlerType: (*CommonServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAppConfig",
			Handler:    _Common_GetAppConfig_Handler,
		},
		{
			MethodName: "GetBaseApp",
			Handler:    _Common_GetBaseApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common.proto",
}