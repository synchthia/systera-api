// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: systera.proto

package systerapb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SysteraClient is the client API for Systera service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysteraClient interface {
	Announce(ctx context.Context, in *AnnounceRequest, opts ...grpc.CallOption) (*Empty, error)
	Dispatch(ctx context.Context, in *DispatchRequest, opts ...grpc.CallOption) (*Empty, error)
	// Chat
	Chat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*Empty, error)
	AddChatIgnore(ctx context.Context, in *AddChatIgnoreRequest, opts ...grpc.CallOption) (*Empty, error)
	RemoveChatIgnore(ctx context.Context, in *RemoveChatIgnoreRequest, opts ...grpc.CallOption) (*Empty, error)
	InitPlayerProfile(ctx context.Context, in *InitPlayerProfileRequest, opts ...grpc.CallOption) (*InitPlayerProfileResponse, error)
	FetchPlayerProfile(ctx context.Context, in *FetchPlayerProfileRequest, opts ...grpc.CallOption) (*FetchPlayerProfileResponse, error)
	FetchPlayerProfileByName(ctx context.Context, in *FetchPlayerProfileByNameRequest, opts ...grpc.CallOption) (*FetchPlayerProfileResponse, error)
	SetPlayerGroups(ctx context.Context, in *SetPlayerGroupsRequest, opts ...grpc.CallOption) (*Empty, error)
	SetPlayerServer(ctx context.Context, in *SetPlayerServerRequest, opts ...grpc.CallOption) (*Empty, error)
	RemovePlayerServer(ctx context.Context, in *RemovePlayerServerRequest, opts ...grpc.CallOption) (*Empty, error)
	SetPlayerSettings(ctx context.Context, in *SetPlayerSettingsRequest, opts ...grpc.CallOption) (*Empty, error)
	AltLookup(ctx context.Context, in *AltLookupRequest, opts ...grpc.CallOption) (*AltLookupResponse, error)
	GetPlayerPunish(ctx context.Context, in *GetPlayerPunishRequest, opts ...grpc.CallOption) (*GetPlayerPunishResponse, error)
	SetPlayerPunish(ctx context.Context, in *SetPlayerPunishRequest, opts ...grpc.CallOption) (*SetPlayerPunishResponse, error)
	UnBan(ctx context.Context, in *UnBanRequest, opts ...grpc.CallOption) (*UnBanResponse, error)
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
	FetchGroups(ctx context.Context, in *FetchGroupsRequest, opts ...grpc.CallOption) (*FetchGroupsResponse, error)
	CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Empty, error)
	RemoveGroup(ctx context.Context, in *RemoveGroupRequest, opts ...grpc.CallOption) (*Empty, error)
	AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*Empty, error)
	RemovePermission(ctx context.Context, in *RemovePermissionRequest, opts ...grpc.CallOption) (*Empty, error)
}

type systeraClient struct {
	cc grpc.ClientConnInterface
}

func NewSysteraClient(cc grpc.ClientConnInterface) SysteraClient {
	return &systeraClient{cc}
}

func (c *systeraClient) Announce(ctx context.Context, in *AnnounceRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/Announce", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) Dispatch(ctx context.Context, in *DispatchRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) Chat(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/Chat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) AddChatIgnore(ctx context.Context, in *AddChatIgnoreRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/AddChatIgnore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) RemoveChatIgnore(ctx context.Context, in *RemoveChatIgnoreRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/RemoveChatIgnore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) InitPlayerProfile(ctx context.Context, in *InitPlayerProfileRequest, opts ...grpc.CallOption) (*InitPlayerProfileResponse, error) {
	out := new(InitPlayerProfileResponse)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/InitPlayerProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) FetchPlayerProfile(ctx context.Context, in *FetchPlayerProfileRequest, opts ...grpc.CallOption) (*FetchPlayerProfileResponse, error) {
	out := new(FetchPlayerProfileResponse)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/FetchPlayerProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) FetchPlayerProfileByName(ctx context.Context, in *FetchPlayerProfileByNameRequest, opts ...grpc.CallOption) (*FetchPlayerProfileResponse, error) {
	out := new(FetchPlayerProfileResponse)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/FetchPlayerProfileByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) SetPlayerGroups(ctx context.Context, in *SetPlayerGroupsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/SetPlayerGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) SetPlayerServer(ctx context.Context, in *SetPlayerServerRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/SetPlayerServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) RemovePlayerServer(ctx context.Context, in *RemovePlayerServerRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/RemovePlayerServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) SetPlayerSettings(ctx context.Context, in *SetPlayerSettingsRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/SetPlayerSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) AltLookup(ctx context.Context, in *AltLookupRequest, opts ...grpc.CallOption) (*AltLookupResponse, error) {
	out := new(AltLookupResponse)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/AltLookup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) GetPlayerPunish(ctx context.Context, in *GetPlayerPunishRequest, opts ...grpc.CallOption) (*GetPlayerPunishResponse, error) {
	out := new(GetPlayerPunishResponse)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/GetPlayerPunish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) SetPlayerPunish(ctx context.Context, in *SetPlayerPunishRequest, opts ...grpc.CallOption) (*SetPlayerPunishResponse, error) {
	out := new(SetPlayerPunishResponse)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/SetPlayerPunish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) UnBan(ctx context.Context, in *UnBanRequest, opts ...grpc.CallOption) (*UnBanResponse, error) {
	out := new(UnBanResponse)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/UnBan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) FetchGroups(ctx context.Context, in *FetchGroupsRequest, opts ...grpc.CallOption) (*FetchGroupsResponse, error) {
	out := new(FetchGroupsResponse)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/FetchGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) CreateGroup(ctx context.Context, in *CreateGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) RemoveGroup(ctx context.Context, in *RemoveGroupRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/RemoveGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) AddPermission(ctx context.Context, in *AddPermissionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/AddPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systeraClient) RemovePermission(ctx context.Context, in *RemovePermissionRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/systerapb.Systera/RemovePermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysteraServer is the server API for Systera service.
// All implementations should embed UnimplementedSysteraServer
// for forward compatibility
type SysteraServer interface {
	Announce(context.Context, *AnnounceRequest) (*Empty, error)
	Dispatch(context.Context, *DispatchRequest) (*Empty, error)
	// Chat
	Chat(context.Context, *ChatRequest) (*Empty, error)
	AddChatIgnore(context.Context, *AddChatIgnoreRequest) (*Empty, error)
	RemoveChatIgnore(context.Context, *RemoveChatIgnoreRequest) (*Empty, error)
	InitPlayerProfile(context.Context, *InitPlayerProfileRequest) (*InitPlayerProfileResponse, error)
	FetchPlayerProfile(context.Context, *FetchPlayerProfileRequest) (*FetchPlayerProfileResponse, error)
	FetchPlayerProfileByName(context.Context, *FetchPlayerProfileByNameRequest) (*FetchPlayerProfileResponse, error)
	SetPlayerGroups(context.Context, *SetPlayerGroupsRequest) (*Empty, error)
	SetPlayerServer(context.Context, *SetPlayerServerRequest) (*Empty, error)
	RemovePlayerServer(context.Context, *RemovePlayerServerRequest) (*Empty, error)
	SetPlayerSettings(context.Context, *SetPlayerSettingsRequest) (*Empty, error)
	AltLookup(context.Context, *AltLookupRequest) (*AltLookupResponse, error)
	GetPlayerPunish(context.Context, *GetPlayerPunishRequest) (*GetPlayerPunishResponse, error)
	SetPlayerPunish(context.Context, *SetPlayerPunishRequest) (*SetPlayerPunishResponse, error)
	UnBan(context.Context, *UnBanRequest) (*UnBanResponse, error)
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
	FetchGroups(context.Context, *FetchGroupsRequest) (*FetchGroupsResponse, error)
	CreateGroup(context.Context, *CreateGroupRequest) (*Empty, error)
	RemoveGroup(context.Context, *RemoveGroupRequest) (*Empty, error)
	AddPermission(context.Context, *AddPermissionRequest) (*Empty, error)
	RemovePermission(context.Context, *RemovePermissionRequest) (*Empty, error)
}

// UnimplementedSysteraServer should be embedded to have forward compatible implementations.
type UnimplementedSysteraServer struct {
}

func (UnimplementedSysteraServer) Announce(context.Context, *AnnounceRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Announce not implemented")
}
func (UnimplementedSysteraServer) Dispatch(context.Context, *DispatchRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispatch not implemented")
}
func (UnimplementedSysteraServer) Chat(context.Context, *ChatRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedSysteraServer) AddChatIgnore(context.Context, *AddChatIgnoreRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddChatIgnore not implemented")
}
func (UnimplementedSysteraServer) RemoveChatIgnore(context.Context, *RemoveChatIgnoreRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveChatIgnore not implemented")
}
func (UnimplementedSysteraServer) InitPlayerProfile(context.Context, *InitPlayerProfileRequest) (*InitPlayerProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitPlayerProfile not implemented")
}
func (UnimplementedSysteraServer) FetchPlayerProfile(context.Context, *FetchPlayerProfileRequest) (*FetchPlayerProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPlayerProfile not implemented")
}
func (UnimplementedSysteraServer) FetchPlayerProfileByName(context.Context, *FetchPlayerProfileByNameRequest) (*FetchPlayerProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchPlayerProfileByName not implemented")
}
func (UnimplementedSysteraServer) SetPlayerGroups(context.Context, *SetPlayerGroupsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPlayerGroups not implemented")
}
func (UnimplementedSysteraServer) SetPlayerServer(context.Context, *SetPlayerServerRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPlayerServer not implemented")
}
func (UnimplementedSysteraServer) RemovePlayerServer(context.Context, *RemovePlayerServerRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePlayerServer not implemented")
}
func (UnimplementedSysteraServer) SetPlayerSettings(context.Context, *SetPlayerSettingsRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPlayerSettings not implemented")
}
func (UnimplementedSysteraServer) AltLookup(context.Context, *AltLookupRequest) (*AltLookupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AltLookup not implemented")
}
func (UnimplementedSysteraServer) GetPlayerPunish(context.Context, *GetPlayerPunishRequest) (*GetPlayerPunishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayerPunish not implemented")
}
func (UnimplementedSysteraServer) SetPlayerPunish(context.Context, *SetPlayerPunishRequest) (*SetPlayerPunishResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPlayerPunish not implemented")
}
func (UnimplementedSysteraServer) UnBan(context.Context, *UnBanRequest) (*UnBanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnBan not implemented")
}
func (UnimplementedSysteraServer) Report(context.Context, *ReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}
func (UnimplementedSysteraServer) FetchGroups(context.Context, *FetchGroupsRequest) (*FetchGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchGroups not implemented")
}
func (UnimplementedSysteraServer) CreateGroup(context.Context, *CreateGroupRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedSysteraServer) RemoveGroup(context.Context, *RemoveGroupRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGroup not implemented")
}
func (UnimplementedSysteraServer) AddPermission(context.Context, *AddPermissionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPermission not implemented")
}
func (UnimplementedSysteraServer) RemovePermission(context.Context, *RemovePermissionRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemovePermission not implemented")
}

// UnsafeSysteraServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysteraServer will
// result in compilation errors.
type UnsafeSysteraServer interface {
	mustEmbedUnimplementedSysteraServer()
}

func RegisterSysteraServer(s grpc.ServiceRegistrar, srv SysteraServer) {
	s.RegisterService(&Systera_ServiceDesc, srv)
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
		FullMethod: "/systerapb.Systera/Announce",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).Announce(ctx, req.(*AnnounceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DispatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).Dispatch(ctx, req.(*DispatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_Chat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).Chat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/Chat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).Chat(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_AddChatIgnore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddChatIgnoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).AddChatIgnore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/AddChatIgnore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).AddChatIgnore(ctx, req.(*AddChatIgnoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_RemoveChatIgnore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveChatIgnoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).RemoveChatIgnore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/RemoveChatIgnore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).RemoveChatIgnore(ctx, req.(*RemoveChatIgnoreRequest))
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
		FullMethod: "/systerapb.Systera/InitPlayerProfile",
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
		FullMethod: "/systerapb.Systera/FetchPlayerProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).FetchPlayerProfile(ctx, req.(*FetchPlayerProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_FetchPlayerProfileByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchPlayerProfileByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).FetchPlayerProfileByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/FetchPlayerProfileByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).FetchPlayerProfileByName(ctx, req.(*FetchPlayerProfileByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_SetPlayerGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPlayerGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).SetPlayerGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/SetPlayerGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).SetPlayerGroups(ctx, req.(*SetPlayerGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_SetPlayerServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPlayerServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).SetPlayerServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/SetPlayerServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).SetPlayerServer(ctx, req.(*SetPlayerServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_RemovePlayerServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePlayerServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).RemovePlayerServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/RemovePlayerServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).RemovePlayerServer(ctx, req.(*RemovePlayerServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_SetPlayerSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPlayerSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).SetPlayerSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/SetPlayerSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).SetPlayerSettings(ctx, req.(*SetPlayerSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_AltLookup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AltLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).AltLookup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/AltLookup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).AltLookup(ctx, req.(*AltLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_GetPlayerPunish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayerPunishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).GetPlayerPunish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/GetPlayerPunish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).GetPlayerPunish(ctx, req.(*GetPlayerPunishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_SetPlayerPunish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPlayerPunishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).SetPlayerPunish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/SetPlayerPunish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).SetPlayerPunish(ctx, req.(*SetPlayerPunishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_UnBan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnBanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).UnBan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/UnBan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).UnBan(ctx, req.(*UnBanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_FetchGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).FetchGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/FetchGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).FetchGroups(ctx, req.(*FetchGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).CreateGroup(ctx, req.(*CreateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_RemoveGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).RemoveGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/RemoveGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).RemoveGroup(ctx, req.(*RemoveGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_AddPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).AddPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/AddPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).AddPermission(ctx, req.(*AddPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Systera_RemovePermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemovePermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysteraServer).RemovePermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/systerapb.Systera/RemovePermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysteraServer).RemovePermission(ctx, req.(*RemovePermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Systera_ServiceDesc is the grpc.ServiceDesc for Systera service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Systera_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "systerapb.Systera",
	HandlerType: (*SysteraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Announce",
			Handler:    _Systera_Announce_Handler,
		},
		{
			MethodName: "Dispatch",
			Handler:    _Systera_Dispatch_Handler,
		},
		{
			MethodName: "Chat",
			Handler:    _Systera_Chat_Handler,
		},
		{
			MethodName: "AddChatIgnore",
			Handler:    _Systera_AddChatIgnore_Handler,
		},
		{
			MethodName: "RemoveChatIgnore",
			Handler:    _Systera_RemoveChatIgnore_Handler,
		},
		{
			MethodName: "InitPlayerProfile",
			Handler:    _Systera_InitPlayerProfile_Handler,
		},
		{
			MethodName: "FetchPlayerProfile",
			Handler:    _Systera_FetchPlayerProfile_Handler,
		},
		{
			MethodName: "FetchPlayerProfileByName",
			Handler:    _Systera_FetchPlayerProfileByName_Handler,
		},
		{
			MethodName: "SetPlayerGroups",
			Handler:    _Systera_SetPlayerGroups_Handler,
		},
		{
			MethodName: "SetPlayerServer",
			Handler:    _Systera_SetPlayerServer_Handler,
		},
		{
			MethodName: "RemovePlayerServer",
			Handler:    _Systera_RemovePlayerServer_Handler,
		},
		{
			MethodName: "SetPlayerSettings",
			Handler:    _Systera_SetPlayerSettings_Handler,
		},
		{
			MethodName: "AltLookup",
			Handler:    _Systera_AltLookup_Handler,
		},
		{
			MethodName: "GetPlayerPunish",
			Handler:    _Systera_GetPlayerPunish_Handler,
		},
		{
			MethodName: "SetPlayerPunish",
			Handler:    _Systera_SetPlayerPunish_Handler,
		},
		{
			MethodName: "UnBan",
			Handler:    _Systera_UnBan_Handler,
		},
		{
			MethodName: "Report",
			Handler:    _Systera_Report_Handler,
		},
		{
			MethodName: "FetchGroups",
			Handler:    _Systera_FetchGroups_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _Systera_CreateGroup_Handler,
		},
		{
			MethodName: "RemoveGroup",
			Handler:    _Systera_RemoveGroup_Handler,
		},
		{
			MethodName: "AddPermission",
			Handler:    _Systera_AddPermission_Handler,
		},
		{
			MethodName: "RemovePermission",
			Handler:    _Systera_RemovePermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "systera.proto",
}