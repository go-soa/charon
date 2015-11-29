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
	GetUserPermissionsRequest
	GetUserPermissionsResponse
*/
package charon

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import mnemosyne "github.com/piotrkowalczuk/mnemosyne"
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

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type LoginResponse struct {
	Token *mnemosyne.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LoginResponse) Reset()                    { *m = LoginResponse{} }
func (m *LoginResponse) String() string            { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()               {}
func (*LoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginResponse) GetToken() *mnemosyne.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type LogoutRequest struct {
	Token *mnemosyne.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *LogoutRequest) Reset()                    { *m = LogoutRequest{} }
func (m *LogoutRequest) String() string            { return proto.CompactTextString(m) }
func (*LogoutRequest) ProtoMessage()               {}
func (*LogoutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *LogoutRequest) GetToken() *mnemosyne.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type LogoutResponse struct {
}

func (m *LogoutResponse) Reset()                    { *m = LogoutResponse{} }
func (m *LogoutResponse) String() string            { return proto.CompactTextString(m) }
func (*LogoutResponse) ProtoMessage()               {}
func (*LogoutResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type IsAuthenticatedRequest struct {
	Token *mnemosyne.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *IsAuthenticatedRequest) Reset()                    { *m = IsAuthenticatedRequest{} }
func (m *IsAuthenticatedRequest) String() string            { return proto.CompactTextString(m) }
func (*IsAuthenticatedRequest) ProtoMessage()               {}
func (*IsAuthenticatedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *IsAuthenticatedRequest) GetToken() *mnemosyne.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type IsAuthenticatedResponse struct {
	IsAuthenticated bool `protobuf:"varint,1,opt,name=is_authenticated" json:"is_authenticated,omitempty"`
}

func (m *IsAuthenticatedResponse) Reset()                    { *m = IsAuthenticatedResponse{} }
func (m *IsAuthenticatedResponse) String() string            { return proto.CompactTextString(m) }
func (*IsAuthenticatedResponse) ProtoMessage()               {}
func (*IsAuthenticatedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type IsGrantedRequest struct {
	Token      *mnemosyne.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	UserId     int64            `protobuf:"varint,2,opt,name=user_id" json:"user_id,omitempty"`
	Permission string           `protobuf:"bytes,3,opt,name=permission" json:"permission,omitempty"`
}

func (m *IsGrantedRequest) Reset()                    { *m = IsGrantedRequest{} }
func (m *IsGrantedRequest) String() string            { return proto.CompactTextString(m) }
func (*IsGrantedRequest) ProtoMessage()               {}
func (*IsGrantedRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *IsGrantedRequest) GetToken() *mnemosyne.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type IsGrantedResponse struct {
	IsGranted bool `protobuf:"varint,1,opt,name=is_granted" json:"is_granted,omitempty"`
}

func (m *IsGrantedResponse) Reset()                    { *m = IsGrantedResponse{} }
func (m *IsGrantedResponse) String() string            { return proto.CompactTextString(m) }
func (*IsGrantedResponse) ProtoMessage()               {}
func (*IsGrantedResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type BelongsToRequest struct {
}

func (m *BelongsToRequest) Reset()                    { *m = BelongsToRequest{} }
func (m *BelongsToRequest) String() string            { return proto.CompactTextString(m) }
func (*BelongsToRequest) ProtoMessage()               {}
func (*BelongsToRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type BelongsToResponse struct {
}

func (m *BelongsToResponse) Reset()                    { *m = BelongsToResponse{} }
func (m *BelongsToResponse) String() string            { return proto.CompactTextString(m) }
func (*BelongsToResponse) ProtoMessage()               {}
func (*BelongsToResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

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

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

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
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *CreateUserResponse) Reset()                    { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()               {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CreateUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserRequest struct {
	Token *mnemosyne.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Id    int64            `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *GetUserRequest) GetToken() *mnemosyne.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type GetUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *GetUserResponse) Reset()                    { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()               {}
func (*GetUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *GetUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUsersRequest struct {
	Offset *protot.NilInt64 `protobuf:"bytes,100,opt,name=offset" json:"offset,omitempty"`
	Limit  *protot.NilInt64 `protobuf:"bytes,101,opt,name=limit" json:"limit,omitempty"`
}

func (m *GetUsersRequest) Reset()                    { *m = GetUsersRequest{} }
func (m *GetUsersRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUsersRequest) ProtoMessage()               {}
func (*GetUsersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *GetUsersRequest) GetOffset() *protot.NilInt64 {
	if m != nil {
		return m.Offset
	}
	return nil
}

func (m *GetUsersRequest) GetLimit() *protot.NilInt64 {
	if m != nil {
		return m.Limit
	}
	return nil
}

type GetUsersResponse struct {
	Users []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *GetUsersResponse) Reset()                    { *m = GetUsersResponse{} }
func (m *GetUsersResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUsersResponse) ProtoMessage()               {}
func (*GetUsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *GetUsersResponse) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

type DeleteUserRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteUserRequest) Reset()                    { *m = DeleteUserRequest{} }
func (m *DeleteUserRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserRequest) ProtoMessage()               {}
func (*DeleteUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type DeleteUserResponse struct {
	Affected int64 `protobuf:"varint,1,opt,name=affected" json:"affected,omitempty"`
}

func (m *DeleteUserResponse) Reset()                    { *m = DeleteUserResponse{} }
func (m *DeleteUserResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserResponse) ProtoMessage()               {}
func (*DeleteUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

type ModifyUserRequest struct {
	Id             int64             `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username       *protot.NilString `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	PlainPassword  *protot.NilString `protobuf:"bytes,3,opt,name=plain_password" json:"plain_password,omitempty"`
	SecurePassword *protot.NilString `protobuf:"bytes,4,opt,name=secure_password" json:"secure_password,omitempty"`
	FirstName      *protot.NilString `protobuf:"bytes,5,opt,name=first_name" json:"first_name,omitempty"`
	LastName       *protot.NilString `protobuf:"bytes,6,opt,name=last_name" json:"last_name,omitempty"`
	IsSuperuser    *protot.NilBool   `protobuf:"bytes,7,opt,name=is_superuser" json:"is_superuser,omitempty"`
	IsActive       *protot.NilBool   `protobuf:"bytes,8,opt,name=is_active" json:"is_active,omitempty"`
	IsStaff        *protot.NilBool   `protobuf:"bytes,9,opt,name=is_staff" json:"is_staff,omitempty"`
	IsConfirmed    *protot.NilBool   `protobuf:"bytes,10,opt,name=is_confirmed" json:"is_confirmed,omitempty"`
}

func (m *ModifyUserRequest) Reset()                    { *m = ModifyUserRequest{} }
func (m *ModifyUserRequest) String() string            { return proto.CompactTextString(m) }
func (*ModifyUserRequest) ProtoMessage()               {}
func (*ModifyUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *ModifyUserRequest) GetUsername() *protot.NilString {
	if m != nil {
		return m.Username
	}
	return nil
}

func (m *ModifyUserRequest) GetPlainPassword() *protot.NilString {
	if m != nil {
		return m.PlainPassword
	}
	return nil
}

func (m *ModifyUserRequest) GetSecurePassword() *protot.NilString {
	if m != nil {
		return m.SecurePassword
	}
	return nil
}

func (m *ModifyUserRequest) GetFirstName() *protot.NilString {
	if m != nil {
		return m.FirstName
	}
	return nil
}

func (m *ModifyUserRequest) GetLastName() *protot.NilString {
	if m != nil {
		return m.LastName
	}
	return nil
}

func (m *ModifyUserRequest) GetIsSuperuser() *protot.NilBool {
	if m != nil {
		return m.IsSuperuser
	}
	return nil
}

func (m *ModifyUserRequest) GetIsActive() *protot.NilBool {
	if m != nil {
		return m.IsActive
	}
	return nil
}

func (m *ModifyUserRequest) GetIsStaff() *protot.NilBool {
	if m != nil {
		return m.IsStaff
	}
	return nil
}

func (m *ModifyUserRequest) GetIsConfirmed() *protot.NilBool {
	if m != nil {
		return m.IsConfirmed
	}
	return nil
}

type ModifyUserResponse struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *ModifyUserResponse) Reset()                    { *m = ModifyUserResponse{} }
func (m *ModifyUserResponse) String() string            { return proto.CompactTextString(m) }
func (*ModifyUserResponse) ProtoMessage()               {}
func (*ModifyUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *ModifyUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserPermissionsRequest struct {
	Token  *mnemosyne.Token `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	UserId int64            `protobuf:"varint,2,opt,name=user_id" json:"user_id,omitempty"`
}

func (m *GetUserPermissionsRequest) Reset()                    { *m = GetUserPermissionsRequest{} }
func (m *GetUserPermissionsRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserPermissionsRequest) ProtoMessage()               {}
func (*GetUserPermissionsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *GetUserPermissionsRequest) GetToken() *mnemosyne.Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type GetUserPermissionsResponse struct {
	Permissions []string `protobuf:"bytes,1,rep,name=permissions" json:"permissions,omitempty"`
}

func (m *GetUserPermissionsResponse) Reset()                    { *m = GetUserPermissionsResponse{} }
func (m *GetUserPermissionsResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserPermissionsResponse) ProtoMessage()               {}
func (*GetUserPermissionsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func init() {
	proto.RegisterType((*User)(nil), "charon.User")
	proto.RegisterType((*LoginRequest)(nil), "charon.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "charon.LoginResponse")
	proto.RegisterType((*LogoutRequest)(nil), "charon.LogoutRequest")
	proto.RegisterType((*LogoutResponse)(nil), "charon.LogoutResponse")
	proto.RegisterType((*IsAuthenticatedRequest)(nil), "charon.IsAuthenticatedRequest")
	proto.RegisterType((*IsAuthenticatedResponse)(nil), "charon.IsAuthenticatedResponse")
	proto.RegisterType((*IsGrantedRequest)(nil), "charon.IsGrantedRequest")
	proto.RegisterType((*IsGrantedResponse)(nil), "charon.IsGrantedResponse")
	proto.RegisterType((*BelongsToRequest)(nil), "charon.BelongsToRequest")
	proto.RegisterType((*BelongsToResponse)(nil), "charon.BelongsToResponse")
	proto.RegisterType((*CreateUserRequest)(nil), "charon.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "charon.CreateUserResponse")
	proto.RegisterType((*GetUserRequest)(nil), "charon.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "charon.GetUserResponse")
	proto.RegisterType((*GetUsersRequest)(nil), "charon.GetUsersRequest")
	proto.RegisterType((*GetUsersResponse)(nil), "charon.GetUsersResponse")
	proto.RegisterType((*DeleteUserRequest)(nil), "charon.DeleteUserRequest")
	proto.RegisterType((*DeleteUserResponse)(nil), "charon.DeleteUserResponse")
	proto.RegisterType((*ModifyUserRequest)(nil), "charon.ModifyUserRequest")
	proto.RegisterType((*ModifyUserResponse)(nil), "charon.ModifyUserResponse")
	proto.RegisterType((*GetUserPermissionsRequest)(nil), "charon.GetUserPermissionsRequest")
	proto.RegisterType((*GetUserPermissionsResponse)(nil), "charon.GetUserPermissionsResponse")
}

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
			MethodName: "GetUserPermissions",
			Handler:    _RPC_GetUserPermissions_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 914 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0xed, 0x8e, 0xdb, 0x44,
	0x14, 0x6d, 0xd6, 0x89, 0x63, 0xdf, 0x4d, 0x37, 0xce, 0xb4, 0xec, 0xba, 0x46, 0x62, 0xb7, 0x86,
	0x85, 0x82, 0xc4, 0x02, 0x29, 0xaa, 0x54, 0x09, 0x84, 0xd8, 0x82, 0xaa, 0xad, 0x28, 0xaa, 0x4a,
	0x90, 0xf8, 0x83, 0x22, 0x37, 0x9e, 0xa4, 0x23, 0x62, 0x4f, 0xf0, 0x4c, 0x40, 0xfb, 0x08, 0xfc,
	0xe5, 0x25, 0x78, 0x16, 0x1e, 0x0a, 0x89, 0xf1, 0x78, 0xec, 0x99, 0xf8, 0xa3, 0x1b, 0xf8, 0x65,
	0xf9, 0xcc, 0xb9, 0x67, 0xee, 0xdc, 0x39, 0x73, 0x2f, 0x8c, 0x16, 0xaf, 0xa3, 0x8c, 0xa6, 0x17,
	0x9b, 0x8c, 0x72, 0x8a, 0xec, 0xe2, 0x2f, 0x18, 0x27, 0x29, 0x4e, 0x28, 0xbb, 0x4e, 0x71, 0xb1,
	0x10, 0x8c, 0xe4, 0x87, 0x17, 0x7f, 0xe1, 0x5f, 0x07, 0xd0, 0xff, 0x91, 0xe1, 0x0c, 0x01, 0x1c,
	0x90, 0xd8, 0xef, 0x9d, 0xf5, 0x1e, 0x58, 0xc8, 0x03, 0x67, 0x2b, 0xb0, 0x34, 0x4a, 0xb0, 0x7f,
	0x20, 0x10, 0x17, 0x89, 0xe5, 0x25, 0xc9, 0x18, 0x9f, 0x4b, 0xcc, 0x92, 0xd8, 0x04, 0xdc, 0x75,
	0x54, 0x42, 0x7d, 0x09, 0xdd, 0x85, 0x11, 0x61, 0x73, 0xb6, 0xdd, 0xe0, 0x2c, 0x17, 0xf0, 0x07,
	0x02, 0x75, 0x72, 0xa2, 0x40, 0xa3, 0x05, 0x27, 0xbf, 0x61, 0xdf, 0x96, 0x90, 0xd8, 0x21, 0x27,
	0xf2, 0x68, 0xb9, 0xf4, 0x87, 0x12, 0x29, 0x42, 0x17, 0x34, 0x15, 0xfb, 0x24, 0x38, 0xf6, 0x1d,
	0x89, 0x9e, 0x03, 0x2c, 0x32, 0x1c, 0x71, 0x1c, 0xcf, 0x23, 0xee, 0xbb, 0x02, 0x3b, 0x9c, 0x4e,
	0x2e, 0xd4, 0x09, 0x66, 0x24, 0xc1, 0x42, 0x22, 0xd9, 0xe4, 0xe9, 0x95, 0xb4, 0x57, 0xd7, 0x3e,
	0xc8, 0x43, 0x88, 0xd0, 0xed, 0x26, 0x2e, 0x43, 0x0f, 0xdf, 0x10, 0x5a, 0xd2, 0x44, 0xe8, 0x48,
	0x86, 0xde, 0x81, 0x43, 0x71, 0x82, 0x84, 0x30, 0x46, 0x68, 0xca, 0xfc, 0xdb, 0x67, 0xd6, 0x03,
	0x37, 0x9c, 0xc2, 0xe8, 0x3b, 0xba, 0x22, 0xe9, 0x4b, 0xfc, 0xeb, 0x56, 0xc4, 0xee, 0x14, 0xa9,
	0x27, 0x4f, 0x2f, 0x90, 0x4d, 0xc4, 0xd8, 0xef, 0x34, 0x8b, 0x8b, 0xb2, 0x85, 0x9f, 0xc2, 0x6d,
	0x15, 0xc3, 0x36, 0x42, 0x09, 0xa3, 0x53, 0x18, 0x70, 0xfa, 0x0b, 0x4e, 0x65, 0xc4, 0xe1, 0xd4,
	0xbb, 0xd0, 0xb7, 0x33, 0xcb, 0x71, 0x15, 0x41, 0xb7, 0xbc, 0xdc, 0xe6, 0xc6, 0x08, 0x0f, 0x8e,
	0xca, 0x88, 0x62, 0x93, 0xf0, 0x31, 0x1c, 0x5f, 0xb1, 0xaf, 0xb7, 0xfc, 0x35, 0x4e, 0x39, 0x59,
	0xe4, 0x47, 0xdb, 0x5b, 0xec, 0x21, 0x9c, 0x34, 0x42, 0x55, 0xea, 0x3e, 0x78, 0xf9, 0x2d, 0x9a,
	0x6b, 0x52, 0xc6, 0x09, 0x7f, 0x02, 0xef, 0x8a, 0x3d, 0xcd, 0xa2, 0xf4, 0x3f, 0xec, 0x84, 0xc6,
	0x30, 0xcc, 0xcb, 0x37, 0x27, 0x45, 0xad, 0xac, 0xfc, 0x22, 0x74, 0xd1, 0x0b, 0x8b, 0x85, 0x1f,
	0xc0, 0xc4, 0x50, 0x56, 0x89, 0x08, 0xa2, 0x48, 0x64, 0x55, 0xa0, 0x2a, 0x05, 0x04, 0xde, 0x25,
	0x5e, 0xd3, 0x74, 0xc5, 0x66, 0x54, 0xa5, 0x10, 0xde, 0x81, 0x89, 0x81, 0xa9, 0xda, 0xfc, 0x79,
	0x00, 0x93, 0x27, 0xd2, 0x2a, 0xb9, 0xeb, 0xbb, 0xef, 0xf2, 0x18, 0x8e, 0x36, 0xeb, 0x88, 0xa4,
	0xf3, 0xdd, 0x1b, 0x45, 0x27, 0x30, 0x66, 0x78, 0xb1, 0xcd, 0xb0, 0x5e, 0xb0, 0x5a, 0x5e, 0x48,
	0xbf, 0xf9, 0x42, 0x06, 0x12, 0x3a, 0xaf, 0xbd, 0x10, 0x5b, 0x96, 0x67, 0x5c, 0xfa, 0xf2, 0x7b,
	0xb2, 0xbe, 0xa4, 0x74, 0x8d, 0x42, 0xf3, 0xc9, 0x0c, 0xdb, 0x39, 0xf7, 0x8d, 0x37, 0xe4, 0xb4,
	0x53, 0xce, 0x6b, 0x8f, 0xca, 0x6d, 0xa5, 0x09, 0xd3, 0x21, 0xb3, 0x26, 0xaa, 0xce, 0x01, 0xf4,
	0x65, 0x8a, 0xc5, 0x0d, 0x8e, 0x2e, 0x54, 0x7b, 0xc9, 0x39, 0xe1, 0x97, 0x70, 0xf4, 0x14, 0x73,
	0xb3, 0x84, 0x37, 0x5e, 0x78, 0xd1, 0x60, 0xe4, 0x5d, 0x87, 0x1f, 0xc3, 0xb8, 0x0a, 0xdf, 0x63,
	0xb7, 0xa8, 0xa2, 0xb3, 0x72, 0xbb, 0x33, 0xb0, 0xe9, 0x72, 0xc9, 0x30, 0xf7, 0x63, 0xb5, 0x9f,
	0x3e, 0xd3, 0x55, 0xca, 0x1f, 0x7d, 0x9e, 0x27, 0xb4, 0x26, 0x09, 0xe1, 0x3e, 0x6e, 0x27, 0x3c,
	0xeb, 0x3b, 0x3d, 0x2f, 0x7e, 0x36, 0x70, 0x96, 0xde, 0xdf, 0xbd, 0xf0, 0x13, 0xf0, 0xf4, 0x16,
	0x2a, 0xa5, 0xb7, 0x61, 0x90, 0xa7, 0xc4, 0x44, 0x4e, 0x56, 0x23, 0xa7, 0x53, 0x98, 0x7c, 0x83,
	0xd7, 0x78, 0xd7, 0x47, 0x46, 0x13, 0x0d, 0xdf, 0x07, 0x64, 0x12, 0x94, 0xa6, 0x70, 0x9a, 0xb8,
	0x2f, 0xbc, 0x28, 0xad, 0x6b, 0x85, 0xff, 0x08, 0x47, 0x3e, 0xa7, 0x31, 0x59, 0x5e, 0x77, 0x28,
	0xa1, 0x77, 0x6b, 0xed, 0xd8, 0xe8, 0x63, 0xe2, 0x30, 0x3f, 0xf0, 0x8c, 0xa4, 0x2b, 0xf4, 0x61,
	0xc3, 0xb0, 0x56, 0x17, 0xf5, 0xa3, 0xa6, 0x87, 0xfb, 0x5d, 0xdc, 0xf3, 0x1d, 0x5b, 0x0f, 0xba,
	0x68, 0xef, 0x99, 0x4e, 0xb7, 0xbb, 0xc5, 0x76, 0xcd, 0x3f, 0xdc, 0xc3, 0xfc, 0xce, 0xcd, 0xe6,
	0x77, 0xf7, 0x33, 0x3f, 0x74, 0x9a, 0xdf, 0x2c, 0xff, 0x1e, 0x76, 0x7c, 0x0e, 0xf7, 0x94, 0x57,
	0x5e, 0xe8, 0x29, 0xf1, 0xbf, 0x1b, 0x5f, 0xf8, 0x19, 0x04, 0x6d, 0x72, 0x2a, 0x91, 0xda, 0x2c,
	0xca, 0xad, 0xe8, 0x4e, 0xff, 0xb0, 0xc1, 0x7a, 0xf9, 0xe2, 0x09, 0x7a, 0x04, 0x03, 0x39, 0x5f,
	0xd0, 0xdd, 0x32, 0x41, 0x73, 0x44, 0x05, 0x6f, 0xd5, 0x50, 0xd5, 0x03, 0x6f, 0xa1, 0xc7, 0x60,
	0x17, 0x33, 0x03, 0x99, 0x14, 0x3d, 0x75, 0x82, 0xe3, 0x3a, 0x5c, 0x85, 0xce, 0x60, 0x5c, 0x9b,
	0x10, 0xe8, 0x9d, 0x92, 0xdc, 0x3e, 0x75, 0x82, 0xd3, 0xce, 0xf5, 0x4a, 0xf5, 0x12, 0xdc, 0xaa,
	0xd1, 0x23, 0x5f, 0xf3, 0x77, 0xa7, 0x4a, 0x70, 0xaf, 0x65, 0xc5, 0xd4, 0xa8, 0xfa, 0xbd, 0xd6,
	0xa8, 0x8f, 0x05, 0xad, 0xd1, 0x1c, 0x0e, 0xb7, 0xd0, 0xb7, 0x00, 0xba, 0x13, 0xa2, 0x8a, 0xda,
	0x98, 0x18, 0x41, 0xd0, 0xb6, 0x64, 0xca, 0x68, 0x4f, 0x69, 0x99, 0xc6, 0x33, 0xd7, 0x32, 0x4d,
	0x0b, 0x0a, 0x99, 0x2f, 0x60, 0xa8, 0x9c, 0x81, 0xaa, 0x0b, 0xd9, 0x6d, 0xbb, 0xc1, 0x49, 0x03,
	0xaf, 0xa2, 0xbf, 0x02, 0xa7, 0x6c, 0x69, 0xa8, 0x4e, 0x2b, 0xed, 0x1a, 0xf8, 0xcd, 0x05, 0xf3,
	0x14, 0xba, 0x83, 0xe9, 0x53, 0x34, 0xda, 0x9e, 0x3e, 0x45, 0xb3, 0xe1, 0x09, 0x99, 0x9f, 0x01,
	0x35, 0xfd, 0x8d, 0xee, 0xd7, 0x36, 0x6e, 0x3e, 0xa5, 0x20, 0x7c, 0x13, 0xa5, 0x94, 0x7f, 0x65,
	0xcb, 0xf7, 0xfc, 0xf0, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x9f, 0x88, 0x66, 0xff, 0x0a,
	0x00, 0x00,
}
