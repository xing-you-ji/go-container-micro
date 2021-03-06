// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

/*
Package go_micro_service_user is a generated protocol buffer package.

It is generated from these files:
	proto/user/user.proto

It has these top-level messages:
	UserLoginRequest
	UserLoginResponse
	UserRegisterRequest
	UserRegisterResponse
	UserInfoRequest
	UserInfoResponse
*/
package go_micro_service_user

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

type UserLoginRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	Pwd      string `protobuf:"bytes,2,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *UserLoginRequest) Reset()                    { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()               {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UserLoginRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserLoginRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserLoginResponse struct {
	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess" json:"is_success,omitempty"`
}

func (m *UserLoginResponse) Reset()                    { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string            { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()               {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UserLoginResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

type UserRegisterRequest struct {
	UserName  string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	Pwd       string `protobuf:"bytes,3,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *UserRegisterRequest) Reset()                    { *m = UserRegisterRequest{} }
func (m *UserRegisterRequest) String() string            { return proto.CompactTextString(m) }
func (*UserRegisterRequest) ProtoMessage()               {}
func (*UserRegisterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UserRegisterRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserRegisterRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UserRegisterRequest) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UserRegisterResponse struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *UserRegisterResponse) Reset()                    { *m = UserRegisterResponse{} }
func (m *UserRegisterResponse) String() string            { return proto.CompactTextString(m) }
func (*UserRegisterResponse) ProtoMessage()               {}
func (*UserRegisterResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UserRegisterResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type UserInfoRequest struct {
	UserName string `protobuf:"bytes,1,opt,name=user_name,json=userName" json:"user_name,omitempty"`
}

func (m *UserInfoRequest) Reset()                    { *m = UserInfoRequest{} }
func (m *UserInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*UserInfoRequest) ProtoMessage()               {}
func (*UserInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserInfoRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type UserInfoResponse struct {
	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=user_name,json=userName" json:"user_name,omitempty"`
	FirstName string `protobuf:"bytes,3,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
}

func (m *UserInfoResponse) Reset()                    { *m = UserInfoResponse{} }
func (m *UserInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*UserInfoResponse) ProtoMessage()               {}
func (*UserInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserInfoResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfoResponse) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserInfoResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func init() {
	proto.RegisterType((*UserLoginRequest)(nil), "UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "UserLoginResponse")
	proto.RegisterType((*UserRegisterRequest)(nil), "UserRegisterRequest")
	proto.RegisterType((*UserRegisterResponse)(nil), "UserRegisterResponse")
	proto.RegisterType((*UserInfoRequest)(nil), "UserInfoRequest")
	proto.RegisterType((*UserInfoResponse)(nil), "UserInfoResponse")
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x3f, 0x4f, 0x83, 0x40,
	0x14, 0x6f, 0x8b, 0xb6, 0xe5, 0x39, 0xd8, 0x9e, 0x6d, 0x24, 0x98, 0x26, 0xe6, 0x26, 0xa7, 0x6b,
	0x53, 0x5d, 0x8c, 0x93, 0x2e, 0xa6, 0x89, 0x71, 0xc0, 0xb8, 0xb8, 0x10, 0x84, 0x57, 0x72, 0x03,
	0x1c, 0xde, 0x03, 0xfd, 0x4c, 0x7e, 0x4b, 0x73, 0x47, 0x89, 0x2d, 0x36, 0xa6, 0x0b, 0xf0, 0x1e,
	0xef, 0xfd, 0xfe, 0xdd, 0xc1, 0xb4, 0xd0, 0xaa, 0x54, 0xf3, 0x8a, 0x50, 0xdb, 0x87, 0xb0, 0x35,
	0xbf, 0x87, 0xd1, 0x2b, 0xa1, 0x7e, 0x52, 0xa9, 0xcc, 0x03, 0xfc, 0xa8, 0x90, 0x4a, 0x76, 0x01,
	0xae, 0x99, 0x08, 0xf3, 0x28, 0x43, 0xaf, 0x7b, 0xd9, 0xbd, 0x72, 0x83, 0xa1, 0x69, 0x3c, 0x47,
	0x19, 0xb2, 0x11, 0x38, 0xc5, 0x57, 0xe2, 0xf5, 0x6c, 0xdb, 0x7c, 0xf2, 0x25, 0x8c, 0xb7, 0x20,
	0xa8, 0x50, 0x39, 0x21, 0x9b, 0x01, 0x48, 0x0a, 0xa9, 0x8a, 0x63, 0x24, 0xb2, 0x20, 0xc3, 0xc0,
	0x95, 0xf4, 0x52, 0x37, 0x78, 0x0c, 0x67, 0x66, 0x27, 0xc0, 0x54, 0x52, 0x69, 0xde, 0x07, 0x30,
	0xcf, 0x00, 0xd6, 0x52, 0x53, 0x59, 0xff, 0xad, 0x05, 0xb8, 0xb6, 0xb3, 0x2d, 0xcc, 0xf9, 0x15,
	0xb6, 0x80, 0xc9, 0x2e, 0xc9, 0x46, 0x9b, 0x07, 0x83, 0x0c, 0x89, 0xa2, 0xb4, 0xe1, 0x68, 0x4a,
	0x2e, 0xe0, 0xd4, 0x6c, 0xac, 0xf2, 0xb5, 0x3a, 0x44, 0x12, 0x4f, 0xeb, 0xf4, 0xea, 0xf9, 0x0d,
	0xfa, 0x39, 0x0c, 0xec, 0x82, 0x4c, 0xec, 0xb8, 0x13, 0xf4, 0x4d, 0xb9, 0x4a, 0x76, 0x91, 0x7a,
	0xff, 0x9a, 0x73, 0x5a, 0xe6, 0x96, 0xdf, 0x5d, 0x38, 0x32, 0x4c, 0x6c, 0x01, 0xc7, 0x36, 0x68,
	0x36, 0x16, 0xed, 0x73, 0xf3, 0x99, 0xf8, 0x73, 0x0e, 0xbc, 0xc3, 0x6e, 0x61, 0xd8, 0x24, 0xc0,
	0x26, 0x62, 0x4f, 0xea, 0xfe, 0x54, 0xec, 0x8b, 0x89, 0x77, 0xd8, 0x0d, 0x9c, 0x3c, 0x62, 0xd9,
	0x38, 0x64, 0x23, 0xd1, 0x0a, 0xc7, 0x1f, 0x8b, 0xb6, 0x7d, 0xde, 0x79, 0xf0, 0xdf, 0x3c, 0x31,
	0xbf, 0x4b, 0x55, 0x98, 0xc9, 0x58, 0xab, 0x90, 0x50, 0x7f, 0xca, 0x18, 0x43, 0x63, 0xf5, 0xbd,
	0x6f, 0x6f, 0xdd, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x35, 0xef, 0x84, 0x48, 0x8e, 0x02,
	0x00, 0x00,
}
