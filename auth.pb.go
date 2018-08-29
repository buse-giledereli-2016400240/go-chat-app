// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package go_micro_srv_auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_725845f6d2d5fbd0, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Errors               []*Error `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_725845f6d2d5fbd0, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_725845f6d2d5fbd0, []int{2}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_725845f6d2d5fbd0, []int{3}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.auth.User")
	proto.RegisterType((*Response)(nil), "go.micro.srv.auth.Response")
	proto.RegisterType((*Token)(nil), "go.micro.srv.auth.Token")
	proto.RegisterType((*Error)(nil), "go.micro.srv.auth.Error")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_725845f6d2d5fbd0) }

var fileDescriptor_auth_725845f6d2d5fbd0 = []byte{
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0xed, 0x47, 0x52, 0xea, 0x04, 0x0f, 0x0e, 0x82, 0xa1, 0x5e, 0x4a, 0x4e, 0x05, 0x61, 0x91,
	0x7a, 0x14, 0x45, 0x91, 0xfe, 0x81, 0xf5, 0xe3, 0xbe, 0xa6, 0x83, 0x5d, 0xb4, 0xd9, 0x30, 0xbb,
	0x8d, 0x3f, 0xd4, 0x3f, 0x24, 0xbb, 0x9b, 0x4a, 0xc1, 0x44, 0xbc, 0xcd, 0x9b, 0xf7, 0x66, 0xdf,
	0xcc, 0x63, 0x01, 0xd4, 0xce, 0x6d, 0x44, 0xcd, 0xc6, 0x19, 0x3c, 0x79, 0x33, 0x62, 0xab, 0x4b,
	0x36, 0xc2, 0x72, 0x23, 0x3c, 0x51, 0xdc, 0x42, 0xf2, 0x6c, 0x89, 0x71, 0x06, 0xd3, 0x9d, 0x25,
	0xae, 0xd4, 0x96, 0xf2, 0xe1, 0x7c, 0xb8, 0x38, 0x92, 0x3f, 0xd8, 0x73, 0xb5, 0xb2, 0xf6, 0xd3,
	0xf0, 0x3a, 0x1f, 0x45, 0x6e, 0x8f, 0x0b, 0x0d, 0x53, 0x49, 0xb6, 0x36, 0x95, 0x25, 0xbc, 0x80,
	0xc4, 0xcf, 0x84, 0xf9, 0x6c, 0x79, 0x26, 0x7e, 0xb9, 0x09, 0x6f, 0x25, 0x83, 0x08, 0x2f, 0x61,
	0x42, 0xcc, 0x86, 0x6d, 0x3e, 0x9a, 0x8f, 0x17, 0xd9, 0x32, 0xef, 0x90, 0xaf, 0xbc, 0x40, 0xb6,
	0xba, 0x82, 0x20, 0x7d, 0x32, 0xef, 0x54, 0xe1, 0x29, 0xa4, 0xce, 0x17, 0xed, 0xa2, 0x11, 0xf8,
	0x6e, 0xa3, 0x3e, 0x74, 0x5c, 0x71, 0x2a, 0x23, 0x38, 0xb0, 0x19, 0xff, 0xd3, 0xe6, 0x06, 0xd2,
	0xd0, 0x40, 0x84, 0xa4, 0x34, 0xeb, 0x18, 0x47, 0x2a, 0x43, 0x8d, 0x73, 0xc8, 0xd6, 0x64, 0x4b,
	0xd6, 0xb5, 0xd3, 0xa6, 0x6a, 0xd3, 0x38, 0x6c, 0x2d, 0xbf, 0x86, 0x90, 0xf9, 0x33, 0x1f, 0x89,
	0x1b, 0x5d, 0x12, 0xde, 0xc1, 0xe4, 0x81, 0x49, 0x39, 0xc2, 0xbe, 0x40, 0x66, 0xe7, 0x1d, 0xc4,
	0x3e, 0xd4, 0x62, 0x80, 0xd7, 0x90, 0xdc, 0xef, 0xdc, 0xa6, 0x7f, 0xbe, 0xeb, 0xa6, 0x90, 0x54,
	0x31, 0xc0, 0x15, 0x1c, 0xbf, 0xf8, 0x20, 0x94, 0xa3, 0x18, 0x5e, 0xaf, 0xf8, 0xaf, 0x67, 0x5e,
	0x27, 0xe1, 0x03, 0x5d, 0x7d, 0x07, 0x00, 0x00, 0xff, 0xff, 0x61, 0xa8, 0x81, 0x9d, 0x4e, 0x02,
	0x00, 0x00,
}
