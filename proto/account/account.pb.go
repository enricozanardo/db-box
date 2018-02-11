// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/account.proto

/*
Package account is a generated protocol buffer package.

It is generated from these files:
	account/account.proto

It has these top-level messages:
	Empty
	Account
	Credentials
	Token
	Response
	Status
	UpdateStatus
	Accounts
	Email
*/
package account

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

type Account_Status int32

const (
	Account_NOTSET    Account_Status = 0
	Account_ENABLED   Account_Status = 1
	Account_DISABLED  Account_Status = 2
	Account_SUSPENDED Account_Status = 3
	Account_REVOKED   Account_Status = 4
)

var Account_Status_name = map[int32]string{
	0: "NOTSET",
	1: "ENABLED",
	2: "DISABLED",
	3: "SUSPENDED",
	4: "REVOKED",
}
var Account_Status_value = map[string]int32{
	"NOTSET":    0,
	"ENABLED":   1,
	"DISABLED":  2,
	"SUSPENDED": 3,
	"REVOKED":   4,
}

func (x Account_Status) String() string {
	return proto.EnumName(Account_Status_name, int32(x))
}
func (Account_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Status_Status int32

const (
	Status_NOTSET    Status_Status = 0
	Status_ENABLED   Status_Status = 1
	Status_DISABLED  Status_Status = 2
	Status_SUSPENDED Status_Status = 3
	Status_REVOKED   Status_Status = 4
)

var Status_Status_name = map[int32]string{
	0: "NOTSET",
	1: "ENABLED",
	2: "DISABLED",
	3: "SUSPENDED",
	4: "REVOKED",
}
var Status_Status_value = map[string]int32{
	"NOTSET":    0,
	"ENABLED":   1,
	"DISABLED":  2,
	"SUSPENDED": 3,
	"REVOKED":   4,
}

func (x Status_Status) String() string {
	return proto.EnumName(Status_Status_name, int32(x))
}
func (Status_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Account struct {
	Uuid       string         `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Username   string         `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Password   string         `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	Token      *Token         `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
	Status     Account_Status `protobuf:"varint,5,opt,name=status,enum=account.Account_Status" json:"status,omitempty"`
	Type       string         `protobuf:"bytes,6,opt,name=type" json:"type,omitempty"`
	Created    string         `protobuf:"bytes,7,opt,name=created" json:"created,omitempty"`
	Expiration string         `protobuf:"bytes,8,opt,name=expiration" json:"expiration,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Account) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Account) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Account) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Account) GetStatus() Account_Status {
	if m != nil {
		return m.Status
	}
	return Account_NOTSET
}

func (m *Account) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Account) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Account) GetExpiration() string {
	if m != nil {
		return m.Expiration
	}
	return ""
}

type Credentials struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Token    *Token `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
}

func (m *Credentials) Reset()                    { *m = Credentials{} }
func (m *Credentials) String() string            { return proto.CompactTextString(m) }
func (*Credentials) ProtoMessage()               {}
func (*Credentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Credentials) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Credentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Credentials) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type Token struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type Response struct {
	Code  int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Token *Token `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Response) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

type Status struct {
	Status Status_Status `protobuf:"varint,1,opt,name=status,enum=account.Status_Status" json:"status,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Status) GetStatus() Status_Status {
	if m != nil {
		return m.Status
	}
	return Status_NOTSET
}

type UpdateStatus struct {
	Token  *Token         `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Status Account_Status `protobuf:"varint,2,opt,name=status,enum=account.Account_Status" json:"status,omitempty"`
}

func (m *UpdateStatus) Reset()                    { *m = UpdateStatus{} }
func (m *UpdateStatus) String() string            { return proto.CompactTextString(m) }
func (*UpdateStatus) ProtoMessage()               {}
func (*UpdateStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateStatus) GetToken() *Token {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *UpdateStatus) GetStatus() Account_Status {
	if m != nil {
		return m.Status
	}
	return Account_NOTSET
}

type Accounts struct {
	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
}

func (m *Accounts) Reset()                    { *m = Accounts{} }
func (m *Accounts) String() string            { return proto.CompactTextString(m) }
func (*Accounts) ProtoMessage()               {}
func (*Accounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Accounts) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type Email struct {
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Email) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "account.Empty")
	proto.RegisterType((*Account)(nil), "account.Account")
	proto.RegisterType((*Credentials)(nil), "account.Credentials")
	proto.RegisterType((*Token)(nil), "account.Token")
	proto.RegisterType((*Response)(nil), "account.Response")
	proto.RegisterType((*Status)(nil), "account.Status")
	proto.RegisterType((*UpdateStatus)(nil), "account.UpdateStatus")
	proto.RegisterType((*Accounts)(nil), "account.Accounts")
	proto.RegisterType((*Email)(nil), "account.Email")
	proto.RegisterEnum("account.Account_Status", Account_Status_name, Account_Status_value)
	proto.RegisterEnum("account.Status_Status", Status_Status_name, Status_Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AccountService service

type AccountServiceClient interface {
	// Create a new Account
	CreateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Response, error)
	// Get Account given the Credentials
	GetAccountByCredentials(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Account, error)
	// Get an Account given the Token
	GetAccountByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Account, error)
	// Update an Account given the updated Account
	UpdateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Response, error)
	// Delete an Account given the Token
	DeleteAccount(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Response, error)
	// Check if an email address is already used
	CheckEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Response, error)
	// Get the Status of an account given the Token
	GetAccountStatus(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Status, error)
	// Set the Status of an account given the Updated Status
	SetAccountStatus(ctx context.Context, in *UpdateStatus, opts ...grpc.CallOption) (*Response, error)
	// Get all the accounts based on a specific Status
	GetAccountsByStatus(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Accounts, error)
	// Get the account collection
	GetAccounts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Accounts, error)
}

type accountServiceClient struct {
	cc *grpc.ClientConn
}

func NewAccountServiceClient(cc *grpc.ClientConn) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) CreateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/account.AccountService/CreateAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountByCredentials(ctx context.Context, in *Credentials, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := grpc.Invoke(ctx, "/account.AccountService/GetAccountByCredentials", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountByToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := grpc.Invoke(ctx, "/account.AccountService/GetAccountByToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) UpdateAccount(ctx context.Context, in *Account, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/account.AccountService/UpdateAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) DeleteAccount(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/account.AccountService/DeleteAccount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CheckEmail(ctx context.Context, in *Email, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/account.AccountService/CheckEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountStatus(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/account.AccountService/GetAccountStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) SetAccountStatus(ctx context.Context, in *UpdateStatus, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/account.AccountService/SetAccountStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccountsByStatus(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Accounts, error) {
	out := new(Accounts)
	err := grpc.Invoke(ctx, "/account.AccountService/GetAccountsByStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetAccounts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Accounts, error) {
	out := new(Accounts)
	err := grpc.Invoke(ctx, "/account.AccountService/GetAccounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountService service

type AccountServiceServer interface {
	// Create a new Account
	CreateAccount(context.Context, *Account) (*Response, error)
	// Get Account given the Credentials
	GetAccountByCredentials(context.Context, *Credentials) (*Account, error)
	// Get an Account given the Token
	GetAccountByToken(context.Context, *Token) (*Account, error)
	// Update an Account given the updated Account
	UpdateAccount(context.Context, *Account) (*Response, error)
	// Delete an Account given the Token
	DeleteAccount(context.Context, *Token) (*Response, error)
	// Check if an email address is already used
	CheckEmail(context.Context, *Email) (*Response, error)
	// Get the Status of an account given the Token
	GetAccountStatus(context.Context, *Token) (*Status, error)
	// Set the Status of an account given the Updated Status
	SetAccountStatus(context.Context, *UpdateStatus) (*Response, error)
	// Get all the accounts based on a specific Status
	GetAccountsByStatus(context.Context, *Status) (*Accounts, error)
	// Get the account collection
	GetAccounts(context.Context, *Empty) (*Accounts, error)
}

func RegisterAccountServiceServer(s *grpc.Server, srv AccountServiceServer) {
	s.RegisterService(&_AccountService_serviceDesc, srv)
}

func _AccountService_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CreateAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountByCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Credentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountByCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/GetAccountByCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountByCredentials(ctx, req.(*Credentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/GetAccountByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountByToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_UpdateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Account)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).UpdateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/UpdateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).UpdateAccount(ctx, req.(*Account))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/DeleteAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).DeleteAccount(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CheckEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Email)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CheckEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/CheckEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CheckEmail(ctx, req.(*Email))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/GetAccountStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountStatus(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_SetAccountStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStatus)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).SetAccountStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/SetAccountStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).SetAccountStatus(ctx, req.(*UpdateStatus))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccountsByStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Status)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccountsByStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/GetAccountsByStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccountsByStatus(ctx, req.(*Status))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.AccountService/GetAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetAccounts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AccountService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "account.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _AccountService_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccountByCredentials",
			Handler:    _AccountService_GetAccountByCredentials_Handler,
		},
		{
			MethodName: "GetAccountByToken",
			Handler:    _AccountService_GetAccountByToken_Handler,
		},
		{
			MethodName: "UpdateAccount",
			Handler:    _AccountService_UpdateAccount_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _AccountService_DeleteAccount_Handler,
		},
		{
			MethodName: "CheckEmail",
			Handler:    _AccountService_CheckEmail_Handler,
		},
		{
			MethodName: "GetAccountStatus",
			Handler:    _AccountService_GetAccountStatus_Handler,
		},
		{
			MethodName: "SetAccountStatus",
			Handler:    _AccountService_SetAccountStatus_Handler,
		},
		{
			MethodName: "GetAccountsByStatus",
			Handler:    _AccountService_GetAccountsByStatus_Handler,
		},
		{
			MethodName: "GetAccounts",
			Handler:    _AccountService_GetAccounts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/account.proto",
}

func init() { proto.RegisterFile("account/account.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 574 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x5f, 0x6f, 0x93, 0x50,
	0x14, 0x2f, 0xf4, 0x0f, 0xec, 0x74, 0xad, 0xec, 0xba, 0x39, 0xd2, 0x44, 0xd3, 0x90, 0x3d, 0xf4,
	0xc1, 0x74, 0xb1, 0xea, 0x34, 0x31, 0x31, 0x59, 0x0b, 0x31, 0x66, 0xa6, 0x33, 0xd0, 0xf9, 0x8e,
	0x70, 0x12, 0x49, 0x5b, 0x20, 0x70, 0xab, 0xf6, 0x03, 0xf8, 0xd9, 0x7c, 0xf4, 0x2b, 0x19, 0xee,
	0xbd, 0x50, 0xc6, 0xb0, 0x99, 0xc9, 0x9e, 0x7a, 0xcf, 0x39, 0xbf, 0xdf, 0xf9, 0xf3, 0x3b, 0xa7,
	0xc0, 0x89, 0xeb, 0x79, 0xd1, 0x26, 0xa4, 0xe7, 0xe2, 0x77, 0x1c, 0x27, 0x11, 0x8d, 0x88, 0x22,
	0x4c, 0x43, 0x81, 0xb6, 0xb5, 0x8e, 0xe9, 0xd6, 0xf8, 0x2d, 0x83, 0x72, 0xc9, 0x9d, 0x84, 0x40,
	0x6b, 0xb3, 0x09, 0x7c, 0x5d, 0x1a, 0x4a, 0xa3, 0x03, 0x9b, 0xbd, 0xc9, 0x00, 0xd4, 0x4d, 0x8a,
	0x49, 0xe8, 0xae, 0x51, 0x97, 0x99, 0xbf, 0xb0, 0xb3, 0x58, 0xec, 0xa6, 0xe9, 0x8f, 0x28, 0xf1,
	0xf5, 0x26, 0x8f, 0xe5, 0x36, 0x39, 0x83, 0x36, 0x8d, 0x96, 0x18, 0xea, 0xad, 0xa1, 0x34, 0xea,
	0x4e, 0xfa, 0xe3, 0xbc, 0x91, 0x45, 0xe6, 0xb5, 0x79, 0x90, 0x9c, 0x43, 0x27, 0xa5, 0x2e, 0xdd,
	0xa4, 0x7a, 0x7b, 0x28, 0x8d, 0xfa, 0x93, 0xd3, 0x02, 0x26, 0x7a, 0x1a, 0x3b, 0x2c, 0x6c, 0x0b,
	0x58, 0xd6, 0x22, 0xdd, 0xc6, 0xa8, 0x77, 0x78, 0x8b, 0xd9, 0x9b, 0xe8, 0xa0, 0x78, 0x09, 0xba,
	0x14, 0x7d, 0x5d, 0x61, 0xee, 0xdc, 0x24, 0xcf, 0x00, 0xf0, 0x67, 0x1c, 0x24, 0x2e, 0x0d, 0xa2,
	0x50, 0x57, 0x59, 0xb0, 0xe4, 0x31, 0xae, 0xa0, 0xc3, 0xf3, 0x13, 0x80, 0xce, 0xfc, 0x7a, 0xe1,
	0x58, 0x0b, 0xad, 0x41, 0xba, 0xa0, 0x58, 0xf3, 0xcb, 0xe9, 0x27, 0xcb, 0xd4, 0x24, 0x72, 0x08,
	0xaa, 0xf9, 0xd1, 0xe1, 0x96, 0x4c, 0x7a, 0x70, 0xe0, 0xdc, 0x38, 0x9f, 0xad, 0xb9, 0x69, 0x99,
	0x5a, 0x33, 0x43, 0xda, 0xd6, 0x97, 0xeb, 0x2b, 0xcb, 0xd4, 0x5a, 0xc6, 0x12, 0xba, 0xb3, 0x04,
	0x7d, 0x0c, 0x69, 0xe0, 0xae, 0xd2, 0x5b, 0xc2, 0x49, 0x7b, 0x84, 0x93, 0xff, 0x25, 0x5c, 0x73,
	0x8f, 0x70, 0xc6, 0x53, 0x68, 0x33, 0x9b, 0x1c, 0xe7, 0x70, 0x5e, 0x43, 0x84, 0x4d, 0x50, 0x6d,
	0x4c, 0xe3, 0x28, 0x4c, 0x31, 0x93, 0xcc, 0x8b, 0x7c, 0xde, 0x44, 0xdb, 0x66, 0xef, 0x5d, 0x11,
	0x79, 0x5f, 0x91, 0x5f, 0x52, 0xa1, 0xcf, 0xb8, 0x58, 0x94, 0xc4, 0x16, 0xf5, 0xa4, 0x60, 0x70,
	0x40, 0x65, 0x4f, 0x0f, 0xab, 0x2c, 0xc2, 0xe1, 0x4d, 0xec, 0xbb, 0x14, 0x45, 0xca, 0xb3, 0xf2,
	0xcc, 0xf7, 0xb8, 0x2d, 0xf9, 0x5e, 0xb7, 0x65, 0xbc, 0x05, 0x55, 0x44, 0x52, 0xf2, 0x1c, 0x54,
	0x81, 0xce, 0x26, 0x6e, 0x8e, 0xba, 0x13, 0xad, 0x4a, 0xb7, 0x0b, 0x44, 0xb6, 0x0d, 0x6b, 0xed,
	0x06, 0xab, 0x6c, 0x1b, 0x98, 0x3d, 0xf2, 0x6d, 0x30, 0x63, 0xf2, 0xa7, 0x05, 0x7d, 0x41, 0x72,
	0x30, 0xf9, 0x1e, 0x78, 0x48, 0x2e, 0xa0, 0x37, 0x63, 0x47, 0x9a, 0xff, 0xf7, 0xee, 0xa4, 0x1f,
	0x1c, 0x15, 0x9e, 0x7c, 0x95, 0x46, 0x83, 0xcc, 0xe0, 0xf4, 0x03, 0x52, 0x01, 0x99, 0x6e, 0xcb,
	0x07, 0x77, 0x5c, 0xe0, 0x4b, 0xde, 0xc1, 0x9d, 0xbc, 0x46, 0x83, 0xbc, 0x81, 0xa3, 0x72, 0x12,
	0x7e, 0x48, 0x15, 0x15, 0x6b, 0x89, 0x17, 0xd0, 0xe3, 0x8b, 0xf8, 0xcf, 0xae, 0x5f, 0x41, 0xcf,
	0xc4, 0x15, 0xee, 0x78, 0xd5, 0x62, 0xb5, 0xac, 0x17, 0x00, 0xb3, 0x6f, 0xe8, 0x2d, 0xb9, 0xb4,
	0x3b, 0x0a, 0xb3, 0xeb, 0x29, 0xaf, 0x41, 0xdb, 0x4d, 0x26, 0xae, 0xa5, 0x5a, 0xeb, 0x51, 0xe5,
	0x74, 0x8d, 0x06, 0x79, 0x0f, 0x9a, 0x53, 0xa5, 0x9d, 0x14, 0xb0, 0xf2, 0xed, 0xd5, 0x97, 0x7d,
	0x07, 0x8f, 0x77, 0x65, 0xd3, 0xe9, 0x56, 0xa4, 0xa8, 0x56, 0x2a, 0x91, 0x73, 0xac, 0xd1, 0x20,
	0x13, 0xe8, 0x96, 0xc8, 0xb7, 0xe6, 0x8c, 0xe9, 0xb6, 0x96, 0xf3, 0xb5, 0xc3, 0x3e, 0xe7, 0x2f,
	0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x12, 0x0f, 0x79, 0x69, 0xe7, 0x05, 0x00, 0x00,
}
