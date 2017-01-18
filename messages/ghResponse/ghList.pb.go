// Code generated by protoc-gen-go.
// source: messages/ghResponse/ghList.proto
// DO NOT EDIT!

/*
Package ghResponse is a generated protocol buffer package.

It is generated from these files:
	messages/ghResponse/ghList.proto

It has these top-level messages:
	List
*/
package ghResponse

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

type List struct {
	Token  string        `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	GhUser string        `protobuf:"bytes,2,opt,name=ghUser" json:"ghUser,omitempty"`
	List   []*ListGhList `protobuf:"bytes,3,rep,name=list" json:"list,omitempty"`
}

func (m *List) Reset()                    { *m = List{} }
func (m *List) String() string            { return proto.CompactTextString(m) }
func (*List) ProtoMessage()               {}
func (*List) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *List) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *List) GetGhUser() string {
	if m != nil {
		return m.GhUser
	}
	return ""
}

func (m *List) GetList() []*ListGhList {
	if m != nil {
		return m.List
	}
	return nil
}

type ListGhList struct {
	Name            string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	StargazersCount int32  `protobuf:"varint,2,opt,name=stargazers_count,json=stargazersCount" json:"stargazers_count,omitempty"`
}

func (m *ListGhList) Reset()                    { *m = ListGhList{} }
func (m *ListGhList) String() string            { return proto.CompactTextString(m) }
func (*ListGhList) ProtoMessage()               {}
func (*ListGhList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *ListGhList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListGhList) GetStargazersCount() int32 {
	if m != nil {
		return m.StargazersCount
	}
	return 0
}

func init() {
	proto.RegisterType((*List)(nil), "ghResponse.List")
	proto.RegisterType((*ListGhList)(nil), "ghResponse.List.ghList")
}

func init() { proto.RegisterFile("messages/ghResponse/ghList.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0xc8, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x4f, 0xcf, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x05,
	0x32, 0x7d, 0x32, 0x8b, 0x4b, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb8, 0x10, 0x12, 0x4a,
	0xab, 0x18, 0xb9, 0x58, 0x40, 0x52, 0x42, 0x22, 0x5c, 0xac, 0x25, 0xf9, 0xd9, 0xa9, 0x79, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x10, 0x8e, 0x90, 0x18, 0x17, 0x5b, 0x7a, 0x46, 0x68, 0x71,
	0x6a, 0x91, 0x04, 0x13, 0x58, 0x18, 0xca, 0x13, 0xd2, 0xe6, 0x62, 0xc9, 0x01, 0xea, 0x92, 0x60,
	0x56, 0x60, 0xd6, 0xe0, 0x36, 0x12, 0xd7, 0x43, 0x98, 0xa8, 0x07, 0xb6, 0x08, 0x62, 0x5f, 0x10,
	0x58, 0x91, 0x94, 0x3b, 0xc8, 0x10, 0xb0, 0x25, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x50,
	0x3b, 0xc0, 0x6c, 0x21, 0x4d, 0x2e, 0x81, 0xe2, 0x92, 0xc4, 0xa2, 0xf4, 0xc4, 0xaa, 0xd4, 0xa2,
	0xe2, 0xf8, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0xb0, 0x65, 0xac, 0x41, 0xfc, 0x08, 0x71, 0x67, 0x90,
	0x70, 0x12, 0x1b, 0xd8, 0xfd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x06, 0xbc, 0xdc,
	0xe3, 0x00, 0x00, 0x00,
}
