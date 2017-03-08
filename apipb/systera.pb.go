// Code generated by protoc-gen-go.
// source: systera.proto
// DO NOT EDIT!

/*
Package apipb is a generated protocol buffer package.

It is generated from these files:
	systera.proto

It has these top-level messages:
	StreamRequest
	ActionStreamResponse
	AnnounceRequest
	Empty
	InitPlayerProfileRequest
	InitPlayerProfileResponse
	FetchPlayerProfileRequest
	FetchPlayerProfileResponse
*/
package apipb

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

type StreamRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *StreamRequest) Reset()                    { *m = StreamRequest{} }
func (m *StreamRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamRequest) ProtoMessage()               {}
func (*StreamRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StreamRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ActionStreamResponse struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Cmd  string `protobuf:"bytes,2,opt,name=cmd" json:"cmd,omitempty"`
}

func (m *ActionStreamResponse) Reset()                    { *m = ActionStreamResponse{} }
func (m *ActionStreamResponse) String() string            { return proto.CompactTextString(m) }
func (*ActionStreamResponse) ProtoMessage()               {}
func (*ActionStreamResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ActionStreamResponse) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ActionStreamResponse) GetCmd() string {
	if m != nil {
		return m.Cmd
	}
	return ""
}

type AnnounceRequest struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *AnnounceRequest) Reset()                    { *m = AnnounceRequest{} }
func (m *AnnounceRequest) String() string            { return proto.CompactTextString(m) }
func (*AnnounceRequest) ProtoMessage()               {}
func (*AnnounceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AnnounceRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

//
// PLAYER PROFILE
type InitPlayerProfileRequest struct {
	PlayerUUID      string `protobuf:"bytes,1,opt,name=playerUUID" json:"playerUUID,omitempty"`
	PlayerName      string `protobuf:"bytes,2,opt,name=playerName" json:"playerName,omitempty"`
	PlayerIPAddress string `protobuf:"bytes,3,opt,name=playerIPAddress" json:"playerIPAddress,omitempty"`
}

func (m *InitPlayerProfileRequest) Reset()                    { *m = InitPlayerProfileRequest{} }
func (m *InitPlayerProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*InitPlayerProfileRequest) ProtoMessage()               {}
func (*InitPlayerProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InitPlayerProfileRequest) GetPlayerUUID() string {
	if m != nil {
		return m.PlayerUUID
	}
	return ""
}

func (m *InitPlayerProfileRequest) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *InitPlayerProfileRequest) GetPlayerIPAddress() string {
	if m != nil {
		return m.PlayerIPAddress
	}
	return ""
}

type InitPlayerProfileResponse struct {
	HasProfile bool `protobuf:"varint,1,opt,name=hasProfile" json:"hasProfile,omitempty"`
}

func (m *InitPlayerProfileResponse) Reset()                    { *m = InitPlayerProfileResponse{} }
func (m *InitPlayerProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*InitPlayerProfileResponse) ProtoMessage()               {}
func (*InitPlayerProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *InitPlayerProfileResponse) GetHasProfile() bool {
	if m != nil {
		return m.HasProfile
	}
	return false
}

type FetchPlayerProfileRequest struct {
	PlayerUUID string `protobuf:"bytes,1,opt,name=playerUUID" json:"playerUUID,omitempty"`
}

func (m *FetchPlayerProfileRequest) Reset()                    { *m = FetchPlayerProfileRequest{} }
func (m *FetchPlayerProfileRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchPlayerProfileRequest) ProtoMessage()               {}
func (*FetchPlayerProfileRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FetchPlayerProfileRequest) GetPlayerUUID() string {
	if m != nil {
		return m.PlayerUUID
	}
	return ""
}

type FetchPlayerProfileResponse struct {
	Settings map[string]bool `protobuf:"bytes,1,rep,name=settings" json:"settings,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *FetchPlayerProfileResponse) Reset()                    { *m = FetchPlayerProfileResponse{} }
func (m *FetchPlayerProfileResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchPlayerProfileResponse) ProtoMessage()               {}
func (*FetchPlayerProfileResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FetchPlayerProfileResponse) GetSettings() map[string]bool {
	if m != nil {
		return m.Settings
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamRequest)(nil), "apipb.StreamRequest")
	proto.RegisterType((*ActionStreamResponse)(nil), "apipb.ActionStreamResponse")
	proto.RegisterType((*AnnounceRequest)(nil), "apipb.AnnounceRequest")
	proto.RegisterType((*Empty)(nil), "apipb.Empty")
	proto.RegisterType((*InitPlayerProfileRequest)(nil), "apipb.InitPlayerProfileRequest")
	proto.RegisterType((*InitPlayerProfileResponse)(nil), "apipb.InitPlayerProfileResponse")
	proto.RegisterType((*FetchPlayerProfileRequest)(nil), "apipb.FetchPlayerProfileRequest")
	proto.RegisterType((*FetchPlayerProfileResponse)(nil), "apipb.FetchPlayerProfileResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Systera service

type SysteraClient interface {
	ActionStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Systera_ActionStreamClient, error)
	Announce(ctx context.Context, in *AnnounceRequest, opts ...grpc.CallOption) (*Empty, error)
	InitPlayerProfile(ctx context.Context, in *InitPlayerProfileRequest, opts ...grpc.CallOption) (*InitPlayerProfileResponse, error)
	FetchPlayerProfile(ctx context.Context, in *FetchPlayerProfileRequest, opts ...grpc.CallOption) (*FetchPlayerProfileResponse, error)
}

type systeraClient struct {
	cc *grpc.ClientConn
}

func NewSysteraClient(cc *grpc.ClientConn) SysteraClient {
	return &systeraClient{cc}
}

func (c *systeraClient) ActionStream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (Systera_ActionStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Systera_serviceDesc.Streams[0], c.cc, "/apipb.Systera/ActionStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &systeraActionStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Systera_ActionStreamClient interface {
	Recv() (*ActionStreamResponse, error)
	grpc.ClientStream
}

type systeraActionStreamClient struct {
	grpc.ClientStream
}

func (x *systeraActionStreamClient) Recv() (*ActionStreamResponse, error) {
	m := new(ActionStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *systeraClient) Announce(ctx context.Context, in *AnnounceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/apipb.Systera/Announce", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) InitPlayerProfile(ctx context.Context, in *InitPlayerProfileRequest, opts ...grpc.CallOption) (*InitPlayerProfileResponse, error) {
	out := new(InitPlayerProfileResponse)
	err := grpc.Invoke(ctx, "/apipb.Systera/InitPlayerProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) FetchPlayerProfile(ctx context.Context, in *FetchPlayerProfileRequest, opts ...grpc.CallOption) (*FetchPlayerProfileResponse, error) {
	out := new(FetchPlayerProfileResponse)
	err := grpc.Invoke(ctx, "/apipb.Systera/FetchPlayerProfile", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Systera service

type SysteraServer interface {
	ActionStream(*StreamRequest, Systera_ActionStreamServer) error
	Announce(context.Context, *AnnounceRequest) (*Empty, error)
	InitPlayerProfile(context.Context, *InitPlayerProfileRequest) (*InitPlayerProfileResponse, error)
	FetchPlayerProfile(context.Context, *FetchPlayerProfileRequest) (*FetchPlayerProfileResponse, error)
}

func RegisterSysteraServer(s *grpc.Server, srv SysteraServer) {
	s.RegisterService(&_Systera_serviceDesc, srv)
}

func _Systera_ActionStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SysteraServer).ActionStream(m, &systeraActionStreamServer{stream})
}

type Systera_ActionStreamServer interface {
	Send(*ActionStreamResponse) error
	grpc.ServerStream
}

type systeraActionStreamServer struct {
	grpc.ServerStream
}

func (x *systeraActionStreamServer) Send(m *ActionStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Systera_Announce_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnounceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).Announce(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.Systera/Announce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).Announce(ctx, req.(*AnnounceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_InitPlayerProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitPlayerProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).InitPlayerProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.Systera/InitPlayerProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).InitPlayerProfile(ctx, req.(*InitPlayerProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_FetchPlayerProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchPlayerProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).FetchPlayerProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apipb.Systera/FetchPlayerProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).FetchPlayerProfile(ctx, req.(*FetchPlayerProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Systera_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apipb.Systera",
	HandlerType: (*SysteraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Announce",
			Handler:    _Systera_Announce_Handler,
		},
		{
			MethodName: "InitPlayerProfile",
			Handler:    _Systera_InitPlayerProfile_Handler,
		},
		{
			MethodName: "FetchPlayerProfile",
			Handler:    _Systera_FetchPlayerProfile_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ActionStream",
			Handler:       _Systera_ActionStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "systera.proto",
}

func init() { proto.RegisterFile("systera.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 442 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x6d, 0x52, 0x6b, 0xe3, 0xb5, 0xa1, 0x3a, 0x14, 0x49, 0x23, 0xd4, 0x38, 0xbe, 0x14, 0x85,
	0x28, 0xf5, 0x45, 0xac, 0x2f, 0x2d, 0x56, 0x28, 0x82, 0x84, 0x94, 0x82, 0xe0, 0xd3, 0x34, 0x1d,
	0x9b, 0x60, 0x33, 0x89, 0x99, 0xe9, 0x42, 0xde, 0xf7, 0xb7, 0xec, 0x9f, 0xd8, 0x3f, 0xb7, 0x4c,
	0x32, 0xe9, 0xa6, 0xdb, 0x76, 0x17, 0xf6, 0xed, 0xce, 0x99, 0x73, 0x66, 0xee, 0xc7, 0xb9, 0x60,
	0xf2, 0x9c, 0x0b, 0x9a, 0x11, 0x37, 0xcd, 0x12, 0x91, 0xa0, 0x16, 0x49, 0xa3, 0x74, 0x85, 0xdf,
	0x81, 0xb9, 0x10, 0x19, 0x25, 0xb1, 0x4f, 0xff, 0xef, 0x28, 0x17, 0x08, 0xc1, 0x13, 0x46, 0x62,
	0x6a, 0x69, 0x8e, 0x36, 0x7c, 0xe6, 0x17, 0x31, 0xfe, 0x06, 0xbd, 0x49, 0x20, 0xa2, 0x84, 0x55,
	0x54, 0x9e, 0x26, 0x8c, 0x53, 0xc9, 0x15, 0x79, 0xba, 0xe7, 0xca, 0x18, 0xbd, 0x80, 0x66, 0x10,
	0xaf, 0x2d, 0xbd, 0x80, 0x64, 0x88, 0x3f, 0x40, 0x77, 0xc2, 0x58, 0xb2, 0x63, 0x01, 0xad, 0x3e,
	0xb1, 0xa0, 0x1d, 0x53, 0xce, 0xc9, 0xa6, 0xd2, 0x56, 0x47, 0xdc, 0x86, 0xd6, 0x2c, 0x4e, 0x45,
	0x8e, 0x2f, 0x35, 0xb0, 0xe6, 0x2c, 0x12, 0xde, 0x96, 0xe4, 0x34, 0xf3, 0xb2, 0xe4, 0x6f, 0xb4,
	0xdd, 0xeb, 0x07, 0x00, 0x69, 0x81, 0x2f, 0x97, 0xf3, 0xef, 0xea, 0x89, 0x1a, 0x72, 0x7b, 0xff,
	0x4b, 0x96, 0xa2, 0xd7, 0xef, 0x25, 0x82, 0x86, 0xd0, 0x2d, 0x4f, 0x73, 0x6f, 0xb2, 0x5e, 0x67,
	0x94, 0x73, 0xab, 0x59, 0x90, 0xee, 0xc2, 0x78, 0x0c, 0xfd, 0x13, 0x59, 0xa8, 0xfa, 0x07, 0x00,
	0x21, 0xe1, 0x0a, 0x2d, 0xd2, 0x30, 0xfc, 0x1a, 0x22, 0xc5, 0x3f, 0xa8, 0x08, 0xc2, 0xc7, 0xd4,
	0x80, 0xaf, 0x34, 0xb0, 0x4f, 0xa9, 0xd5, 0xdf, 0x3f, 0xc1, 0xe0, 0x54, 0x88, 0x88, 0x6d, 0xb8,
	0xa5, 0x39, 0xcd, 0xe1, 0xf3, 0xd1, 0x47, 0xb7, 0x18, 0xa9, 0x7b, 0x5e, 0xe4, 0x2e, 0x94, 0x62,
	0xc6, 0x44, 0x96, 0xfb, 0xfb, 0x07, 0xec, 0x31, 0x98, 0x07, 0x57, 0x72, 0x8a, 0xff, 0x68, 0xae,
	0xb2, 0x92, 0x21, 0xea, 0x41, 0xeb, 0x82, 0x6c, 0x77, 0x65, 0x37, 0x0d, 0xbf, 0x3c, 0x7c, 0xd5,
	0xbf, 0x68, 0xa3, 0x6b, 0x1d, 0xda, 0x8b, 0xd2, 0x5b, 0x68, 0x06, 0x9d, 0xba, 0x53, 0x50, 0x4f,
	0xe5, 0x74, 0xe0, 0x31, 0xfb, 0xb5, 0x42, 0x4f, 0x99, 0x0a, 0x37, 0x3e, 0x69, 0x68, 0x04, 0x46,
	0x65, 0x19, 0xf4, 0xaa, 0x22, 0x1f, 0x7a, 0xc8, 0xee, 0x28, 0xbc, 0xb4, 0x4b, 0x03, 0xfd, 0x86,
	0x97, 0x47, 0x93, 0x42, 0x6f, 0x14, 0xe9, 0x9c, 0x93, 0x6c, 0xe7, 0x3c, 0xa1, 0xca, 0x07, 0xfd,
	0x01, 0x74, 0xdc, 0x53, 0xe4, 0xdc, 0xd3, 0xee, 0xf2, 0xed, 0xb7, 0x0f, 0x0e, 0x04, 0x37, 0xa6,
	0xef, 0xa1, 0xcf, 0xa8, 0x70, 0x79, 0xce, 0x82, 0x50, 0x84, 0x11, 0x91, 0x1a, 0x57, 0xad, 0xea,
	0xd4, 0x54, 0x7d, 0xf5, 0xe4, 0xca, 0xf2, 0xd5, 0xd3, 0x62, 0x75, 0x3f, 0xdf, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x14, 0x45, 0x22, 0x97, 0xcb, 0x03, 0x00, 0x00,
}
