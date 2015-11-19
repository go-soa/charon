// Code generated by protoc-gen-go.
// source: charon.proto
// DO NOT EDIT!

/*
Package charon is a generated protocol buffer package.

It is generated from these files:
	charon.proto

It has these top-level messages:
	User
	LoginRequest
	LoginResponse
	LogoutRequest
	LogoutResponse
	IsAuthenticatedRequest
	IsAuthenticatedResponse
	IsGrantedRequest
	IsGrantedResponse
	BelongsToRequest
	BelongsToResponse
	CreateUserRequest
	CreateUserResponse
	GetUserRequest
	GetUserResponse
	GetUsersRequest
	GetUsersResponse
	DeleteUserRequest
	DeleteUserResponse
	ModifyUserRequest
	ModifyUserResponse
	ModifyUserPasswordRequest
	ModifyUserPasswordResponse
	GetUserPermissionsRequest
	GetUserPermissionsResponse
*/
package charon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mnemosyne1 "github.com/piotrkowalczuk/mnemosyne"
import protot "github.com/piotrkowalczuk/protot"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	Id          int64             `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username    string            `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	FirstName   string            `protobuf:"bytes,3,opt,name=first_name" json:"first_name,omitempty"`
	LastName    string            `protobuf:"bytes,4,opt,name=last_name" json:"last_name,omitempty"`
	IsSuperuser bool              `protobuf:"varint,5,opt,name=is_superuser" json:"is_superuser,omitempty"`
	IsActive    bool              `protobuf:"varint,6,opt,name=is_active" json:"is_active,omitempty"`
	IsStaff     bool              `protobuf:"varint,7,opt,name=is_staff" json:"is_staff,omitempty"`
	IsConfirmed bool              `protobuf:"varint,8,opt,name=is_confirmed" json:"is_confirmed,omitempty"`
	CreatedAt   *protot.Timestamp `protobuf:"bytes,9,opt,name=created_at" json:"created_at,omitempty"`
	CreatedBy   int64             `protobuf:"varint,10,opt,name=created_by" json:"created_by,omitempty"`
	UpdatedAt   *protot.Timestamp `protobuf:"bytes,11,opt,name=updated_at" json:"updated_at,omitempty"`
	UpdatedBy   int64             `protobuf:"varint,12,opt,name=updated_by" json:"updated_by,omitempty"`
	Permissions []string          `protobuf:"bytes,13,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}

func (m *User) GetCreatedAt() *protot.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *User) GetUpdatedAt() *protot.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}

type LoginResponse struct {
	Token *mnemosyne1.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}

func (m *LoginResponse) GetToken() *mnemosyne1.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type LogoutRequest struct {
	Token *mnemosyne1.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LogoutRequest) Reset()         { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()    {}

func (m *LogoutRequest) GetToken() *mnemosyne1.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type LogoutResponse struct {
}

func (m *LogoutResponse) Reset()         { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()    {}

type IsAuthenticatedRequest struct {
	Token *mnemosyne1.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *IsAuthenticatedRequest) Reset()         { *m = IsAuthenticatedRequest{} }
func (m *IsAuthenticatedRequest) String() string { return proto.CompactTextString(m) }
func (*IsAuthenticatedRequest) ProtoMessage()    {}

func (m *IsAuthenticatedRequest) GetToken() *mnemosyne1.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type IsAuthenticatedResponse struct {
	IsAuthenticated bool `protobuf:"varint,1,opt,name=is_authenticated" json:"is_authenticated,omitempty"`
}

func (m *IsAuthenticatedResponse) Reset()         { *m = IsAuthenticatedResponse{} }
func (m *IsAuthenticatedResponse) String() string { return proto.CompactTextString(m) }
func (*IsAuthenticatedResponse) ProtoMessage()    {}

type IsGrantedRequest struct {
	Token      *mnemosyne1.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	UserId     int64             `protobuf:"varint,2,opt,name=user_id" json:"user_id,omitempty"`
	Permission string            `protobuf:"bytes,3,opt,name=permission" json:"permission,omitempty"`
}

func (m *IsGrantedRequest) Reset()         { *m = IsGrantedRequest{} }
func (m *IsGrantedRequest) String() string { return proto.CompactTextString(m) }
func (*IsGrantedRequest) ProtoMessage()    {}

func (m *IsGrantedRequest) GetToken() *mnemosyne1.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type IsGrantedResponse struct {
	IsGranted bool `protobuf:"varint,1,opt,name=is_granted" json:"is_granted,omitempty"`
}

func (m *IsGrantedResponse) Reset()         { *m = IsGrantedResponse{} }
func (m *IsGrantedResponse) String() string { return proto.CompactTextString(m) }
func (*IsGrantedResponse) ProtoMessage()    {}

type BelongsToRequest struct {
}

func (m *BelongsToRequest) Reset()         { *m = BelongsToRequest{} }
func (m *BelongsToRequest) String() string { return proto.CompactTextString(m) }
func (*BelongsToRequest) ProtoMessage()    {}

type BelongsToResponse struct {
}

func (m *BelongsToResponse) Reset()         { *m = BelongsToResponse{} }
func (m *BelongsToResponse) String() string { return proto.CompactTextString(m) }
func (*BelongsToResponse) ProtoMessage()    {}

type CreateUserRequest struct {
	Username       string          `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	PlainPassword  string          `protobuf:"bytes,2,opt,name=plain_password" json:"plain_password,omitempty"`
	SecurePassword string          `protobuf:"bytes,3,opt,name=secure_password" json:"secure_password,omitempty"`
	FirstName      string          `protobuf:"bytes,4,opt,name=first_name" json:"first_name,omitempty"`
	LastName       string          `protobuf:"bytes,5,opt,name=last_name" json:"last_name,omitempty"`
	IsSuperuser    *protot.NilBool `protobuf:"bytes,6,opt,name=is_superuser" json:"is_superuser,omitempty"`
	IsActive       *protot.NilBool `protobuf:"bytes,7,opt,name=is_active" json:"is_active,omitempty"`
	IsStaff        *protot.NilBool `protobuf:"bytes,8,opt,name=is_staff" json:"is_staff,omitempty"`
	IsConfirmed    *protot.NilBool `protobuf:"bytes,9,opt,name=is_confirmed" json:"is_confirmed,omitempty"`
}

func (m *CreateUserRequest) Reset()         { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()    {}

func (m *CreateUserRequest) GetIsSuperuser() *protot.NilBool {
	if m != nil {
		return m.IsSuperuser
	}
	return nil
}

func (m *CreateUserRequest) GetIsActive() *protot.NilBool {
	if m != nil {
		return m.IsActive
	}
	return nil
}

func (m *CreateUserRequest) GetIsStaff() *protot.NilBool {
	if m != nil {
		return m.IsStaff
	}
	return nil
}

func (m *CreateUserRequest) GetIsConfirmed() *protot.NilBool {
	if m != nil {
		return m.IsConfirmed
	}
	return nil
}

type CreateUserResponse struct {
	Id        int64             `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	CreatedAt *protot.Timestamp `protobuf:"bytes,2,opt,name=created_at" json:"created_at,omitempty"`
}

func (m *CreateUserResponse) Reset()         { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()    {}

func (m *CreateUserResponse) GetCreatedAt() *protot.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type GetUserRequest struct {
	Token *mnemosyne1.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Id    int64             `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}

func (m *GetUserRequest) GetToken() *mnemosyne1.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type GetUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUsersRequest struct {
}

func (m *GetUsersRequest) Reset()         { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()    {}

type GetUsersResponse struct {
	User []*User `protobuf:"bytes,1,rep,name=user" json:"user,omitempty"`
}

func (m *GetUsersResponse) Reset()         { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()    {}

func (m *GetUsersResponse) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

type DeleteUserRequest struct {
}

func (m *DeleteUserRequest) Reset()         { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()    {}

type DeleteUserResponse struct {
}

func (m *DeleteUserResponse) Reset()         { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()    {}

type ModifyUserRequest struct {
}

func (m *ModifyUserRequest) Reset()         { *m = ModifyUserRequest{} }
func (m *ModifyUserRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyUserRequest) ProtoMessage()    {}

type ModifyUserResponse struct {
}

func (m *ModifyUserResponse) Reset()         { *m = ModifyUserResponse{} }
func (m *ModifyUserResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyUserResponse) ProtoMessage()    {}

type ModifyUserPasswordRequest struct {
	Id             int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	PlainPassword  string `protobuf:"bytes,2,opt,name=plain_password" json:"plain_password,omitempty"`
	SecurePassword string `protobuf:"bytes,3,opt,name=secure_password" json:"secure_password,omitempty"`
	Regenerate     bool   `protobuf:"varint,4,opt,name=regenerate" json:"regenerate,omitempty"`
	IsConfirmed    bool   `protobuf:"varint,5,opt,name=is_confirmed" json:"is_confirmed,omitempty"`
}

func (m *ModifyUserPasswordRequest) Reset()         { *m = ModifyUserPasswordRequest{} }
func (m *ModifyUserPasswordRequest) String() string { return proto.CompactTextString(m) }
func (*ModifyUserPasswordRequest) ProtoMessage()    {}

type ModifyUserPasswordResponse struct {
	Generated string `protobuf:"bytes,1,opt,name=generated" json:"generated,omitempty"`
}

func (m *ModifyUserPasswordResponse) Reset()         { *m = ModifyUserPasswordResponse{} }
func (m *ModifyUserPasswordResponse) String() string { return proto.CompactTextString(m) }
func (*ModifyUserPasswordResponse) ProtoMessage()    {}

type GetUserPermissionsRequest struct {
	Token  *mnemosyne1.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	UserId int64             `protobuf:"varint,2,opt,name=user_id" json:"user_id,omitempty"`
}

func (m *GetUserPermissionsRequest) Reset()         { *m = GetUserPermissionsRequest{} }
func (m *GetUserPermissionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserPermissionsRequest) ProtoMessage()    {}

func (m *GetUserPermissionsRequest) GetToken() *mnemosyne1.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type GetUserPermissionsResponse struct {
	Permissions []string `protobuf:"bytes,1,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *GetUserPermissionsResponse) Reset()         { *m = GetUserPermissionsResponse{} }
func (m *GetUserPermissionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserPermissionsResponse) ProtoMessage()    {}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for RPC service

type RPCClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error)
	IsAuthenticated(ctx context.Context, in *IsAuthenticatedRequest, opts ...grpc.CallOption) (*IsAuthenticatedResponse, error)
	IsGranted(ctx context.Context, in *IsGrantedRequest, opts ...grpc.CallOption) (*IsGrantedResponse, error)
	BelongsTo(ctx context.Context, in *BelongsToRequest, opts ...grpc.CallOption) (*BelongsToResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	ModifyUser(ctx context.Context, in *ModifyUserRequest, opts ...grpc.CallOption) (*ModifyUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error)
	DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
	ModifyUserPassword(ctx context.Context, in *ModifyUserPasswordRequest, opts ...grpc.CallOption) (*ModifyUserPasswordResponse, error)
	GetUserPermissions(ctx context.Context, in *GetUserPermissionsRequest, opts ...grpc.CallOption) (*GetUserPermissionsResponse, error)
}

type rPCClient struct {
	cc *grpc.ClientConn
}

func NewRPCClient(cc *grpc.ClientConn) RPCClient {
	return &rPCClient{cc}
}

func (c *rPCClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) Logout(ctx context.Context, in *LogoutRequest, opts ...grpc.CallOption) (*LogoutResponse, error) {
	out := new(LogoutResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/Logout", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) IsAuthenticated(ctx context.Context, in *IsAuthenticatedRequest, opts ...grpc.CallOption) (*IsAuthenticatedResponse, error) {
	out := new(IsAuthenticatedResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/IsAuthenticated", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) IsGranted(ctx context.Context, in *IsGrantedRequest, opts ...grpc.CallOption) (*IsGrantedResponse, error) {
	out := new(IsGrantedResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/IsGranted", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) BelongsTo(ctx context.Context, in *BelongsToRequest, opts ...grpc.CallOption) (*BelongsToResponse, error) {
	out := new(BelongsToResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/BelongsTo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) ModifyUser(ctx context.Context, in *ModifyUserRequest, opts ...grpc.CallOption) (*ModifyUserResponse, error) {
	out := new(ModifyUserResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/ModifyUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) GetUsers(ctx context.Context, in *GetUsersRequest, opts ...grpc.CallOption) (*GetUsersResponse, error) {
	out := new(GetUsersResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/GetUsers", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) DeleteUser(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	out := new(DeleteUserResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/DeleteUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) ModifyUserPassword(ctx context.Context, in *ModifyUserPasswordRequest, opts ...grpc.CallOption) (*ModifyUserPasswordResponse, error) {
	out := new(ModifyUserPasswordResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/ModifyUserPassword", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rPCClient) GetUserPermissions(ctx context.Context, in *GetUserPermissionsRequest, opts ...grpc.CallOption) (*GetUserPermissionsResponse, error) {
	out := new(GetUserPermissionsResponse)
	err := grpc.Invoke(ctx, "/charon.RPC/GetUserPermissions", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RPC service

type RPCServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Logout(context.Context, *LogoutRequest) (*LogoutResponse, error)
	IsAuthenticated(context.Context, *IsAuthenticatedRequest) (*IsAuthenticatedResponse, error)
	IsGranted(context.Context, *IsGrantedRequest) (*IsGrantedResponse, error)
	BelongsTo(context.Context, *BelongsToRequest) (*BelongsToResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ModifyUser(context.Context, *ModifyUserRequest) (*ModifyUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	GetUsers(context.Context, *GetUsersRequest) (*GetUsersResponse, error)
	DeleteUser(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	ModifyUserPassword(context.Context, *ModifyUserPasswordRequest) (*ModifyUserPasswordResponse, error)
	GetUserPermissions(context.Context, *GetUserPermissionsRequest) (*GetUserPermissionsResponse, error)
}

func RegisterRPCServer(s *grpc.Server, srv RPCServer) {
	s.RegisterService(&_RPC_serviceDesc, srv)
}

func _RPC_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).Login(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(LogoutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).Logout(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_IsAuthenticated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IsAuthenticatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).IsAuthenticated(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_IsGranted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(IsGrantedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).IsGranted(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_BelongsTo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(BelongsToRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).BelongsTo(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).CreateUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_ModifyUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ModifyUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).ModifyUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).GetUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).GetUsers(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).DeleteUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_ModifyUserPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ModifyUserPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).ModifyUserPassword(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _RPC_GetUserPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(GetUserPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(RPCServer).GetUserPermissions(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _RPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "charon.RPC",
	HandlerType: (*RPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _RPC_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _RPC_Logout_Handler,
		},
		{
			MethodName: "IsAuthenticated",
			Handler:    _RPC_IsAuthenticated_Handler,
		},
		{
			MethodName: "IsGranted",
			Handler:    _RPC_IsGranted_Handler,
		},
		{
			MethodName: "BelongsTo",
			Handler:    _RPC_BelongsTo_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _RPC_CreateUser_Handler,
		},
		{
			MethodName: "ModifyUser",
			Handler:    _RPC_ModifyUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _RPC_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _RPC_GetUsers_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _RPC_DeleteUser_Handler,
		},
		{
			MethodName: "ModifyUserPassword",
			Handler:    _RPC_ModifyUserPassword_Handler,
		},
		{
			MethodName: "GetUserPermissions",
			Handler:    _RPC_GetUserPermissions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
