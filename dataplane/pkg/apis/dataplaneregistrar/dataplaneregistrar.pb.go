// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dataplaneregistrar.proto

package dataplaneregistrar

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

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

// DataplaneRegistrationRequest is sent by the dataplane to NSM
// to advertise itself and inform NSM about the location of the dataplane socket
// and its initially supported parameters.
type DataplaneRegistrationRequest struct {
	DataplaneName        string   `protobuf:"bytes,1,opt,name=dataplane_name,json=dataplaneName,proto3" json:"dataplane_name,omitempty"`
	DataplaneSocket      string   `protobuf:"bytes,2,opt,name=dataplane_socket,json=dataplaneSocket,proto3" json:"dataplane_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneRegistrationRequest) Reset()         { *m = DataplaneRegistrationRequest{} }
func (m *DataplaneRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationRequest) ProtoMessage()    {}
func (*DataplaneRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_84909ec76baae1e8, []int{0}
}
func (m *DataplaneRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneRegistrationRequest.Unmarshal(m, b)
}
func (m *DataplaneRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneRegistrationRequest.Marshal(b, m, deterministic)
}
func (dst *DataplaneRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneRegistrationRequest.Merge(dst, src)
}
func (m *DataplaneRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_DataplaneRegistrationRequest.Size(m)
}
func (m *DataplaneRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneRegistrationRequest proto.InternalMessageInfo

func (m *DataplaneRegistrationRequest) GetDataplaneName() string {
	if m != nil {
		return m.DataplaneName
	}
	return ""
}

func (m *DataplaneRegistrationRequest) GetDataplaneSocket() string {
	if m != nil {
		return m.DataplaneSocket
	}
	return ""
}

type DataplaneRegistrationReply struct {
	Registered           bool     `protobuf:"varint,1,opt,name=registered,proto3" json:"registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneRegistrationReply) Reset()         { *m = DataplaneRegistrationReply{} }
func (m *DataplaneRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*DataplaneRegistrationReply) ProtoMessage()    {}
func (*DataplaneRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_84909ec76baae1e8, []int{1}
}
func (m *DataplaneRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneRegistrationReply.Unmarshal(m, b)
}
func (m *DataplaneRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneRegistrationReply.Marshal(b, m, deterministic)
}
func (dst *DataplaneRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneRegistrationReply.Merge(dst, src)
}
func (m *DataplaneRegistrationReply) XXX_Size() int {
	return xxx_messageInfo_DataplaneRegistrationReply.Size(m)
}
func (m *DataplaneRegistrationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneRegistrationReply.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneRegistrationReply proto.InternalMessageInfo

func (m *DataplaneRegistrationReply) GetRegistered() bool {
	if m != nil {
		return m.Registered
	}
	return false
}

// DataplaneUnRegistrationRequest is sent by the dataplane to NSM
// to remove itself from the list of available dataplanes.
type DataplaneUnRegistrationRequest struct {
	DataplaneName        string   `protobuf:"bytes,1,opt,name=dataplane_name,json=dataplaneName,proto3" json:"dataplane_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneUnRegistrationRequest) Reset()         { *m = DataplaneUnRegistrationRequest{} }
func (m *DataplaneUnRegistrationRequest) String() string { return proto.CompactTextString(m) }
func (*DataplaneUnRegistrationRequest) ProtoMessage()    {}
func (*DataplaneUnRegistrationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_84909ec76baae1e8, []int{2}
}
func (m *DataplaneUnRegistrationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Unmarshal(m, b)
}
func (m *DataplaneUnRegistrationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Marshal(b, m, deterministic)
}
func (dst *DataplaneUnRegistrationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneUnRegistrationRequest.Merge(dst, src)
}
func (m *DataplaneUnRegistrationRequest) XXX_Size() int {
	return xxx_messageInfo_DataplaneUnRegistrationRequest.Size(m)
}
func (m *DataplaneUnRegistrationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneUnRegistrationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneUnRegistrationRequest proto.InternalMessageInfo

func (m *DataplaneUnRegistrationRequest) GetDataplaneName() string {
	if m != nil {
		return m.DataplaneName
	}
	return ""
}

type DataplaneUnRegistrationReply struct {
	UnRegistered         bool     `protobuf:"varint,1,opt,name=un_registered,json=unRegistered,proto3" json:"un_registered,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataplaneUnRegistrationReply) Reset()         { *m = DataplaneUnRegistrationReply{} }
func (m *DataplaneUnRegistrationReply) String() string { return proto.CompactTextString(m) }
func (*DataplaneUnRegistrationReply) ProtoMessage()    {}
func (*DataplaneUnRegistrationReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_dataplaneregistrar_84909ec76baae1e8, []int{3}
}
func (m *DataplaneUnRegistrationReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Unmarshal(m, b)
}
func (m *DataplaneUnRegistrationReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Marshal(b, m, deterministic)
}
func (dst *DataplaneUnRegistrationReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataplaneUnRegistrationReply.Merge(dst, src)
}
func (m *DataplaneUnRegistrationReply) XXX_Size() int {
	return xxx_messageInfo_DataplaneUnRegistrationReply.Size(m)
}
func (m *DataplaneUnRegistrationReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DataplaneUnRegistrationReply.DiscardUnknown(m)
}

var xxx_messageInfo_DataplaneUnRegistrationReply proto.InternalMessageInfo

func (m *DataplaneUnRegistrationReply) GetUnRegistered() bool {
	if m != nil {
		return m.UnRegistered
	}
	return false
}

func init() {
	proto.RegisterType((*DataplaneRegistrationRequest)(nil), "dataplaneregistrar.DataplaneRegistrationRequest")
	proto.RegisterType((*DataplaneRegistrationReply)(nil), "dataplaneregistrar.DataplaneRegistrationReply")
	proto.RegisterType((*DataplaneUnRegistrationRequest)(nil), "dataplaneregistrar.DataplaneUnRegistrationRequest")
	proto.RegisterType((*DataplaneUnRegistrationReply)(nil), "dataplaneregistrar.DataplaneUnRegistrationReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataplaneRegistrationClient is the client API for DataplaneRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneRegistrationClient interface {
	RequestDataplaneRegistration(ctx context.Context, in *DataplaneRegistrationRequest, opts ...grpc.CallOption) (*DataplaneRegistrationReply, error)
	// RequestLiveness is a stream initiated by NSM to inform the dataplane that NSM is still alive and
	// no re-registration is required. Detection a failure on this "channel" will mean
	// that NSM is gone and the dataplane needs to start re-registration logic.
	RequestLiveness(ctx context.Context, opts ...grpc.CallOption) (DataplaneRegistration_RequestLivenessClient, error)
}

type dataplaneRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneRegistrationClient(cc *grpc.ClientConn) DataplaneRegistrationClient {
	return &dataplaneRegistrationClient{cc}
}

func (c *dataplaneRegistrationClient) RequestDataplaneRegistration(ctx context.Context, in *DataplaneRegistrationRequest, opts ...grpc.CallOption) (*DataplaneRegistrationReply, error) {
	out := new(DataplaneRegistrationReply)
	err := c.cc.Invoke(ctx, "/dataplaneregistrar.DataplaneRegistration/RequestDataplaneRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataplaneRegistrationClient) RequestLiveness(ctx context.Context, opts ...grpc.CallOption) (DataplaneRegistration_RequestLivenessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_DataplaneRegistration_serviceDesc.Streams[0], "/dataplaneregistrar.DataplaneRegistration/RequestLiveness", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataplaneRegistrationRequestLivenessClient{stream}
	return x, nil
}

type DataplaneRegistration_RequestLivenessClient interface {
	Send(*empty.Empty) error
	Recv() (*empty.Empty, error)
	grpc.ClientStream
}

type dataplaneRegistrationRequestLivenessClient struct {
	grpc.ClientStream
}

func (x *dataplaneRegistrationRequestLivenessClient) Send(m *empty.Empty) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataplaneRegistrationRequestLivenessClient) Recv() (*empty.Empty, error) {
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataplaneRegistrationServer is the server API for DataplaneRegistration service.
type DataplaneRegistrationServer interface {
	RequestDataplaneRegistration(context.Context, *DataplaneRegistrationRequest) (*DataplaneRegistrationReply, error)
	// RequestLiveness is a stream initiated by NSM to inform the dataplane that NSM is still alive and
	// no re-registration is required. Detection a failure on this "channel" will mean
	// that NSM is gone and the dataplane needs to start re-registration logic.
	RequestLiveness(DataplaneRegistration_RequestLivenessServer) error
}

func RegisterDataplaneRegistrationServer(s *grpc.Server, srv DataplaneRegistrationServer) {
	s.RegisterService(&_DataplaneRegistration_serviceDesc, srv)
}

func _DataplaneRegistration_RequestDataplaneRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataplaneRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneRegistrationServer).RequestDataplaneRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplaneregistrar.DataplaneRegistration/RequestDataplaneRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneRegistrationServer).RequestDataplaneRegistration(ctx, req.(*DataplaneRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataplaneRegistration_RequestLiveness_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataplaneRegistrationServer).RequestLiveness(&dataplaneRegistrationRequestLivenessServer{stream})
}

type DataplaneRegistration_RequestLivenessServer interface {
	Send(*empty.Empty) error
	Recv() (*empty.Empty, error)
	grpc.ServerStream
}

type dataplaneRegistrationRequestLivenessServer struct {
	grpc.ServerStream
}

func (x *dataplaneRegistrationRequestLivenessServer) Send(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataplaneRegistrationRequestLivenessServer) Recv() (*empty.Empty, error) {
	m := new(empty.Empty)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DataplaneRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplaneregistrar.DataplaneRegistration",
	HandlerType: (*DataplaneRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDataplaneRegistration",
			Handler:    _DataplaneRegistration_RequestDataplaneRegistration_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestLiveness",
			Handler:       _DataplaneRegistration_RequestLiveness_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "dataplaneregistrar.proto",
}

// DataplaneUnRegistrationClient is the client API for DataplaneUnRegistration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataplaneUnRegistrationClient interface {
	RequestDataplaneUnRegistration(ctx context.Context, in *DataplaneUnRegistrationRequest, opts ...grpc.CallOption) (*DataplaneUnRegistrationReply, error)
}

type dataplaneUnRegistrationClient struct {
	cc *grpc.ClientConn
}

func NewDataplaneUnRegistrationClient(cc *grpc.ClientConn) DataplaneUnRegistrationClient {
	return &dataplaneUnRegistrationClient{cc}
}

func (c *dataplaneUnRegistrationClient) RequestDataplaneUnRegistration(ctx context.Context, in *DataplaneUnRegistrationRequest, opts ...grpc.CallOption) (*DataplaneUnRegistrationReply, error) {
	out := new(DataplaneUnRegistrationReply)
	err := c.cc.Invoke(ctx, "/dataplaneregistrar.DataplaneUnRegistration/RequestDataplaneUnRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataplaneUnRegistrationServer is the server API for DataplaneUnRegistration service.
type DataplaneUnRegistrationServer interface {
	RequestDataplaneUnRegistration(context.Context, *DataplaneUnRegistrationRequest) (*DataplaneUnRegistrationReply, error)
}

func RegisterDataplaneUnRegistrationServer(s *grpc.Server, srv DataplaneUnRegistrationServer) {
	s.RegisterService(&_DataplaneUnRegistration_serviceDesc, srv)
}

func _DataplaneUnRegistration_RequestDataplaneUnRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataplaneUnRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataplaneUnRegistrationServer).RequestDataplaneUnRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dataplaneregistrar.DataplaneUnRegistration/RequestDataplaneUnRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataplaneUnRegistrationServer).RequestDataplaneUnRegistration(ctx, req.(*DataplaneUnRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataplaneUnRegistration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dataplaneregistrar.DataplaneUnRegistration",
	HandlerType: (*DataplaneUnRegistrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestDataplaneUnRegistration",
			Handler:    _DataplaneUnRegistration_RequestDataplaneUnRegistration_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataplaneregistrar.proto",
}

func init() {
	proto.RegisterFile("dataplaneregistrar.proto", fileDescriptor_dataplaneregistrar_84909ec76baae1e8)
}

var fileDescriptor_dataplaneregistrar_84909ec76baae1e8 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x14, 0x64, 0x3d, 0x88, 0x3e, 0xac, 0x95, 0x05, 0xb5, 0x84, 0x52, 0x24, 0x22, 0xd4, 0xcb, 0xa6,
	0xb4, 0x57, 0x6f, 0x5a, 0xbc, 0x88, 0x87, 0x88, 0xe7, 0xb2, 0x6d, 0x9f, 0x6b, 0x30, 0xd9, 0x5d,
	0xb3, 0x1b, 0x21, 0x37, 0x4f, 0xfe, 0x10, 0xff, 0x99, 0xff, 0x44, 0xf2, 0xd1, 0x58, 0xd2, 0xa4,
	0x50, 0x2f, 0x39, 0x4c, 0x66, 0xe6, 0xcd, 0xec, 0x7b, 0xd0, 0x5b, 0x72, 0xcb, 0x75, 0xc8, 0x25,
	0xc6, 0x28, 0x02, 0x63, 0x63, 0x1e, 0x33, 0x1d, 0x2b, 0xab, 0x28, 0xdd, 0xfc, 0xe3, 0x4c, 0x44,
	0x60, 0x5f, 0x93, 0x39, 0x5b, 0xa8, 0xc8, 0x13, 0x2a, 0xe4, 0x52, 0x78, 0x39, 0x79, 0x9e, 0xbc,
	0x78, 0xda, 0xa6, 0x1a, 0x8d, 0x87, 0x91, 0xb6, 0x69, 0xf1, 0x2d, 0x8c, 0x5c, 0x0d, 0xfd, 0xbb,
	0x95, 0x95, 0x5f, 0x5a, 0xd9, 0x40, 0x49, 0x1f, 0xdf, 0x13, 0x34, 0x96, 0x5e, 0xc1, 0x71, 0x35,
	0x6a, 0x26, 0x79, 0x84, 0x3d, 0x72, 0x41, 0x86, 0x87, 0x7e, 0xa7, 0x42, 0x1f, 0x79, 0x84, 0xf4,
	0x1a, 0x4e, 0xfe, 0x68, 0x46, 0x2d, 0xde, 0xd0, 0xf6, 0xf6, 0x72, 0x62, 0xb7, 0xc2, 0x9f, 0x72,
	0xd8, 0xbd, 0x01, 0xa7, 0x65, 0xa2, 0x0e, 0x53, 0x3a, 0x00, 0x28, 0x1a, 0x61, 0x8c, 0xcb, 0x7c,
	0xd6, 0x81, 0xbf, 0x86, 0xb8, 0xf7, 0x30, 0xa8, 0xd4, 0xcf, 0xf2, 0xff, 0x89, 0xdd, 0xdb, 0xb5,
	0xe2, 0x75, 0xa3, 0x2c, 0xc8, 0x25, 0x74, 0x12, 0x39, 0xdb, 0xc8, 0x72, 0x94, 0x94, 0xdc, 0x0c,
	0x1b, 0xff, 0x10, 0x38, 0x6d, 0x2c, 0x43, 0x3f, 0x09, 0xf4, 0xcb, 0x44, 0xcd, 0x84, 0x11, 0x6b,
	0x58, 0xee, 0xb6, 0x55, 0x38, 0x6c, 0x07, 0x45, 0xd6, 0x60, 0x0a, 0xdd, 0x52, 0xfa, 0x10, 0x7c,
	0xa0, 0x44, 0x63, 0xe8, 0x19, 0x13, 0x4a, 0x89, 0x10, 0xd9, 0xea, 0x30, 0xd8, 0x34, 0xbb, 0x05,
	0xa7, 0x05, 0x1f, 0x92, 0x11, 0x19, 0x7f, 0x13, 0x38, 0x6f, 0x79, 0x29, 0xfa, 0x45, 0x60, 0x50,
	0x6f, 0x59, 0xa3, 0x8c, 0xb7, 0xa6, 0x6e, 0x5c, 0xa1, 0x33, 0xda, 0x49, 0xa3, 0xc3, 0x74, 0xbe,
	0x9f, 0x07, 0x9f, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x77, 0x0b, 0x09, 0x32, 0x03, 0x00,
	0x00,
}
