// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/protocol/service.proto

package service

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

type UploadStatusCode int32

const (
	UploadStatusCode_Unknown UploadStatusCode = 0
	UploadStatusCode_Ok      UploadStatusCode = 1
	UploadStatusCode_Failed  UploadStatusCode = 2
)

var UploadStatusCode_name = map[int32]string{
	0: "Unknown",
	1: "Ok",
	2: "Failed",
}
var UploadStatusCode_value = map[string]int32{
	"Unknown": 0,
	"Ok":      1,
	"Failed":  2,
}

func (x UploadStatusCode) String() string {
	return proto.EnumName(UploadStatusCode_name, int32(x))
}
func (UploadStatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_service_ac3958e7e0d4cd5a, []int{0}
}

type GetMyCatMessage struct {
	TargetCat            string   `protobuf:"bytes,1,opt,name=target_cat,json=targetCat,proto3" json:"target_cat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMyCatMessage) Reset()         { *m = GetMyCatMessage{} }
func (m *GetMyCatMessage) String() string { return proto.CompactTextString(m) }
func (*GetMyCatMessage) ProtoMessage()    {}
func (*GetMyCatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ac3958e7e0d4cd5a, []int{0}
}
func (m *GetMyCatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMyCatMessage.Unmarshal(m, b)
}
func (m *GetMyCatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMyCatMessage.Marshal(b, m, deterministic)
}
func (dst *GetMyCatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMyCatMessage.Merge(dst, src)
}
func (m *GetMyCatMessage) XXX_Size() int {
	return xxx_messageInfo_GetMyCatMessage.Size(m)
}
func (m *GetMyCatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMyCatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_GetMyCatMessage proto.InternalMessageInfo

func (m *GetMyCatMessage) GetTargetCat() string {
	if m != nil {
		return m.TargetCat
	}
	return ""
}

type MyCatResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyCatResponse) Reset()         { *m = MyCatResponse{} }
func (m *MyCatResponse) String() string { return proto.CompactTextString(m) }
func (*MyCatResponse) ProtoMessage()    {}
func (*MyCatResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ac3958e7e0d4cd5a, []int{1}
}
func (m *MyCatResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyCatResponse.Unmarshal(m, b)
}
func (m *MyCatResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyCatResponse.Marshal(b, m, deterministic)
}
func (dst *MyCatResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyCatResponse.Merge(dst, src)
}
func (m *MyCatResponse) XXX_Size() int {
	return xxx_messageInfo_MyCatResponse.Size(m)
}
func (m *MyCatResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MyCatResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MyCatResponse proto.InternalMessageInfo

func (m *MyCatResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MyCatResponse) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

type UploadFileMessage struct {
	FileName             string   `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadFileMessage) Reset()         { *m = UploadFileMessage{} }
func (m *UploadFileMessage) String() string { return proto.CompactTextString(m) }
func (*UploadFileMessage) ProtoMessage()    {}
func (*UploadFileMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ac3958e7e0d4cd5a, []int{2}
}
func (m *UploadFileMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileMessage.Unmarshal(m, b)
}
func (m *UploadFileMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileMessage.Marshal(b, m, deterministic)
}
func (dst *UploadFileMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileMessage.Merge(dst, src)
}
func (m *UploadFileMessage) XXX_Size() int {
	return xxx_messageInfo_UploadFileMessage.Size(m)
}
func (m *UploadFileMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileMessage proto.InternalMessageInfo

func (m *UploadFileMessage) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *UploadFileMessage) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type UploadFileResponce struct {
	Code                 UploadStatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=UploadStatusCode" json:"code,omitempty"`
	Message              string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *UploadFileResponce) Reset()         { *m = UploadFileResponce{} }
func (m *UploadFileResponce) String() string { return proto.CompactTextString(m) }
func (*UploadFileResponce) ProtoMessage()    {}
func (*UploadFileResponce) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_ac3958e7e0d4cd5a, []int{3}
}
func (m *UploadFileResponce) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadFileResponce.Unmarshal(m, b)
}
func (m *UploadFileResponce) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadFileResponce.Marshal(b, m, deterministic)
}
func (dst *UploadFileResponce) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadFileResponce.Merge(dst, src)
}
func (m *UploadFileResponce) XXX_Size() int {
	return xxx_messageInfo_UploadFileResponce.Size(m)
}
func (m *UploadFileResponce) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadFileResponce.DiscardUnknown(m)
}

var xxx_messageInfo_UploadFileResponce proto.InternalMessageInfo

func (m *UploadFileResponce) GetCode() UploadStatusCode {
	if m != nil {
		return m.Code
	}
	return UploadStatusCode_Unknown
}

func (m *UploadFileResponce) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMyCatMessage)(nil), "GetMyCatMessage")
	proto.RegisterType((*MyCatResponse)(nil), "MyCatResponse")
	proto.RegisterType((*UploadFileMessage)(nil), "UploadFileMessage")
	proto.RegisterType((*UploadFileResponce)(nil), "UploadFileResponce")
	proto.RegisterEnum("UploadStatusCode", UploadStatusCode_name, UploadStatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CatClient is the client API for Cat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CatClient interface {
	GetMyCat(ctx context.Context, in *GetMyCatMessage, opts ...grpc.CallOption) (*MyCatResponse, error)
}

type catClient struct {
	cc *grpc.ClientConn
}

func NewCatClient(cc *grpc.ClientConn) CatClient {
	return &catClient{cc}
}

func (c *catClient) GetMyCat(ctx context.Context, in *GetMyCatMessage, opts ...grpc.CallOption) (*MyCatResponse, error) {
	out := new(MyCatResponse)
	err := c.cc.Invoke(ctx, "/Cat/GetMyCat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CatServer is the server API for Cat service.
type CatServer interface {
	GetMyCat(context.Context, *GetMyCatMessage) (*MyCatResponse, error)
}

func RegisterCatServer(s *grpc.Server, srv CatServer) {
	s.RegisterService(&_Cat_serviceDesc, srv)
}

func _Cat_GetMyCat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyCatMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CatServer).GetMyCat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cat/GetMyCat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CatServer).GetMyCat(ctx, req.(*GetMyCatMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Cat",
	HandlerType: (*CatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMyCat",
			Handler:    _Cat_GetMyCat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/protocol/service.proto",
}

// UploaderClient is the client API for Uploader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UploaderClient interface {
	Upload(ctx context.Context, in *UploadFileMessage, opts ...grpc.CallOption) (*UploadFileResponce, error)
}

type uploaderClient struct {
	cc *grpc.ClientConn
}

func NewUploaderClient(cc *grpc.ClientConn) UploaderClient {
	return &uploaderClient{cc}
}

func (c *uploaderClient) Upload(ctx context.Context, in *UploadFileMessage, opts ...grpc.CallOption) (*UploadFileResponce, error) {
	out := new(UploadFileResponce)
	err := c.cc.Invoke(ctx, "/Uploader/upload", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploaderServer is the server API for Uploader service.
type UploaderServer interface {
	Upload(context.Context, *UploadFileMessage) (*UploadFileResponce, error)
}

func RegisterUploaderServer(s *grpc.Server, srv UploaderServer) {
	s.RegisterService(&_Uploader_serviceDesc, srv)
}

func _Uploader_Upload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploaderServer).Upload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Uploader/Upload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploaderServer).Upload(ctx, req.(*UploadFileMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Uploader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Uploader",
	HandlerType: (*UploaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "upload",
			Handler:    _Uploader_Upload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/protocol/service.proto",
}

func init() { proto.RegisterFile("src/protocol/service.proto", fileDescriptor_service_ac3958e7e0d4cd5a) }

var fileDescriptor_service_ac3958e7e0d4cd5a = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xd1, 0x4b, 0x02, 0x41,
	0x10, 0xc6, 0xd5, 0xe4, 0xd4, 0xa9, 0xec, 0x9c, 0x5e, 0xc4, 0x08, 0xe2, 0x20, 0x88, 0x1e, 0xd6,
	0x50, 0xa2, 0xc7, 0x1e, 0x0e, 0x0c, 0x02, 0x0b, 0x2e, 0x7c, 0x96, 0x6d, 0x6f, 0x92, 0xc3, 0x73,
	0xf7, 0xb8, 0x1d, 0x8b, 0xfe, 0xfb, 0xb8, 0x5d, 0x8f, 0x4a, 0xdf, 0x66, 0xbe, 0x9d, 0xfd, 0x7d,
	0xb3, 0xdf, 0xc2, 0xc8, 0x96, 0x6a, 0x5c, 0x94, 0x86, 0x8d, 0x32, 0xf9, 0xd8, 0x52, 0xf9, 0x99,
	0x29, 0x12, 0x4e, 0x88, 0xee, 0xe0, 0xec, 0x89, 0x78, 0xfe, 0x1d, 0x4b, 0x9e, 0x93, 0xb5, 0x72,
	0x45, 0x78, 0x09, 0xc0, 0xb2, 0x5c, 0x11, 0x2f, 0x95, 0xe4, 0x61, 0xf3, 0xaa, 0x79, 0xd3, 0x4b,
	0x7a, 0x5e, 0x89, 0x25, 0x47, 0x0f, 0x70, 0xea, 0xc6, 0x13, 0xb2, 0x85, 0xd1, 0x96, 0x10, 0xa1,
	0xad, 0xe5, 0x86, 0x76, 0x93, 0xae, 0xae, 0xb4, 0x75, 0xa6, 0xd3, 0x61, 0xcb, 0x6b, 0x55, 0x1d,
	0x3d, 0xc3, 0x60, 0x51, 0xe4, 0x46, 0xa6, 0xb3, 0x2c, 0xa7, 0xda, 0xec, 0x02, 0x7a, 0x1f, 0x59,
	0x4e, 0xcb, 0x3f, 0x84, 0x6e, 0x25, 0xbc, 0x54, 0x94, 0x21, 0x74, 0x94, 0xd1, 0x4c, 0x9a, 0x1d,
	0xe8, 0x24, 0xa9, 0xdb, 0x68, 0x01, 0xf8, 0xcb, 0xf2, 0x9b, 0x28, 0xc2, 0x6b, 0x68, 0x2b, 0x93,
	0x7a, 0x4e, 0x7f, 0x32, 0x10, 0x7e, 0xe4, 0x8d, 0x25, 0x6f, 0x6d, 0x6c, 0x52, 0x4a, 0xdc, 0x71,
	0x85, 0xdd, 0x78, 0xfb, 0xdd, 0x7e, 0x75, 0x7b, 0x3b, 0x85, 0x70, 0xff, 0x0e, 0x1e, 0x43, 0x67,
	0xa1, 0xd7, 0xda, 0x7c, 0xe9, 0xb0, 0x81, 0x01, 0xb4, 0x5e, 0xd7, 0x61, 0x13, 0x01, 0x82, 0x99,
	0xcc, 0x72, 0x4a, 0xc3, 0xd6, 0xe4, 0x1e, 0x8e, 0x62, 0xc9, 0x28, 0xa0, 0x5b, 0x27, 0x89, 0xa1,
	0xd8, 0x0b, 0x75, 0xd4, 0x17, 0xff, 0x42, 0x8b, 0x1a, 0x93, 0x47, 0xe8, 0x7a, 0x2f, 0x2a, 0x71,
	0x0a, 0xc1, 0xd6, 0xd5, 0x88, 0xe2, 0x20, 0xa3, 0xd1, 0xb9, 0x38, 0x7c, 0x6b, 0xd4, 0x78, 0x0f,
	0xdc, 0x0f, 0x4e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x91, 0x04, 0xec, 0xce, 0xdf, 0x01, 0x00,
	0x00,
}