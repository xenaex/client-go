// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package xmsg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Logon struct {
	MsgType string `protobuf:"bytes,35,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	// SessionStatus SessionStatus = 1409;
	HeartBtInt           int32    `protobuf:"varint,108,opt,name=HeartBtInt,proto3" json:"HeartBtInt,omitempty"`
	RejectText           string   `protobuf:"bytes,1328,opt,name=RejectText,proto3" json:"RejectText,omitempty"`
	Account              []uint64 `protobuf:"varint,1,rep,packed,name=Account,proto3" json:"Account,omitempty"`
	SendingTime          int64    `protobuf:"varint,52,opt,name=SendingTime,proto3" json:"SendingTime,omitempty"`
	CstmApplVerId        string   `protobuf:"bytes,1129,opt,name=CstmApplVerId,proto3" json:"CstmApplVerId,omitempty"`
	Username             string   `protobuf:"bytes,553,opt,name=Username,proto3" json:"Username,omitempty"`
	Password             string   `protobuf:"bytes,554,opt,name=Password,proto3" json:"Password,omitempty"`
	RawData              string   `protobuf:"bytes,96,opt,name=RawData,proto3" json:"RawData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Logon) Reset()         { *m = Logon{} }
func (m *Logon) String() string { return proto.CompactTextString(m) }
func (*Logon) ProtoMessage()    {}
func (*Logon) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{0}
}

func (m *Logon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Logon.Unmarshal(m, b)
}
func (m *Logon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Logon.Marshal(b, m, deterministic)
}
func (m *Logon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Logon.Merge(m, src)
}
func (m *Logon) XXX_Size() int {
	return xxx_messageInfo_Logon.Size(m)
}
func (m *Logon) XXX_DiscardUnknown() {
	xxx_messageInfo_Logon.DiscardUnknown(m)
}

var xxx_messageInfo_Logon proto.InternalMessageInfo

func (m *Logon) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *Logon) GetHeartBtInt() int32 {
	if m != nil {
		return m.HeartBtInt
	}
	return 0
}

func (m *Logon) GetRejectText() string {
	if m != nil {
		return m.RejectText
	}
	return ""
}

func (m *Logon) GetAccount() []uint64 {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *Logon) GetSendingTime() int64 {
	if m != nil {
		return m.SendingTime
	}
	return 0
}

func (m *Logon) GetCstmApplVerId() string {
	if m != nil {
		return m.CstmApplVerId
	}
	return ""
}

func (m *Logon) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *Logon) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Logon) GetRawData() string {
	if m != nil {
		return m.RawData
	}
	return ""
}

type AccountInfo struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=Id,json=id,proto3" json:"id,omitempty"`
	Kind                 string   `protobuf:"bytes,2,opt,name=Kind,json=kind,proto3" json:"kind,omitempty"`
	Currency             string   `protobuf:"bytes,3,opt,name=Currency,json=currency,proto3" json:"currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountInfo) Reset()         { *m = AccountInfo{} }
func (m *AccountInfo) String() string { return proto.CompactTextString(m) }
func (*AccountInfo) ProtoMessage()    {}
func (*AccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bbd6f3875b0e874, []int{1}
}

func (m *AccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountInfo.Unmarshal(m, b)
}
func (m *AccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountInfo.Marshal(b, m, deterministic)
}
func (m *AccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountInfo.Merge(m, src)
}
func (m *AccountInfo) XXX_Size() int {
	return xxx_messageInfo_AccountInfo.Size(m)
}
func (m *AccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AccountInfo proto.InternalMessageInfo

func (m *AccountInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *AccountInfo) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *AccountInfo) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func init() {
	proto.RegisterType((*Logon)(nil), "xmsg.Logon")
	proto.RegisterType((*AccountInfo)(nil), "xmsg.AccountInfo")
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_8bbd6f3875b0e874) }

var fileDescriptor_8bbd6f3875b0e874 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0xe5, 0xd4, 0x7d, 0xaf, 0x4c, 0x05, 0x0b, 0xaf, 0x46, 0x20, 0x81, 0x55, 0x84, 0x94,
	0x15, 0x1b, 0xb8, 0x40, 0x29, 0x0b, 0x22, 0xa8, 0x84, 0x4c, 0x61, 0x8d, 0x89, 0x4d, 0x08, 0x34,
	0x76, 0x64, 0x3b, 0x6a, 0x7b, 0x14, 0x8e, 0x00, 0x2b, 0xae, 0xc5, 0x2d, 0x50, 0xd2, 0x04, 0x85,
	0xe5, 0xf7, 0x8d, 0xe7, 0xf7, 0x6f, 0x03, 0xc8, 0x2a, 0xbc, 0x9c, 0x96, 0xce, 0x06, 0xcb, 0xe8,
	0xba, 0xf0, 0xd9, 0xe4, 0x3d, 0x82, 0xe1, 0x8d, 0xcd, 0xac, 0x61, 0x08, 0xff, 0xe7, 0x3e, 0x5b,
	0x6c, 0x4a, 0x8d, 0xc7, 0x9c, 0xc4, 0x3b, 0xa2, 0x43, 0x76, 0x08, 0x70, 0xa5, 0xa5, 0x0b, 0x17,
	0x21, 0x31, 0x01, 0x97, 0x9c, 0xc4, 0x43, 0xd1, 0x33, 0xec, 0x08, 0x40, 0xe8, 0x57, 0x9d, 0x86,
	0x85, 0x5e, 0x07, 0xfc, 0x82, 0x66, 0xbb, 0xa7, 0xea, 0xe8, 0x69, 0x9a, 0xda, 0xca, 0x04, 0x24,
	0x7c, 0x10, 0x53, 0xd1, 0x21, 0xe3, 0x30, 0xbe, 0xd3, 0x46, 0xe5, 0x26, 0x5b, 0xe4, 0x85, 0xc6,
	0x73, 0x4e, 0xe2, 0x81, 0xe8, 0x2b, 0x76, 0x02, 0xbb, 0x33, 0x1f, 0x8a, 0x69, 0x59, 0x2e, 0x1f,
	0xb4, 0x4b, 0x14, 0x7e, 0x8f, 0x9a, 0xfc, 0xbf, 0x96, 0x1d, 0xc0, 0xe8, 0xde, 0x6b, 0x67, 0x64,
	0xa1, 0xf1, 0x83, 0x36, 0x27, 0x7e, 0x45, 0x3d, 0xbc, 0x95, 0xde, 0xaf, 0xac, 0x53, 0xf8, 0xd9,
	0x0e, 0x3b, 0x51, 0x97, 0x13, 0x72, 0x75, 0x29, 0x83, 0xc4, 0xc7, 0xed, 0xbb, 0x5b, 0x9c, 0xcc,
	0x61, 0xdc, 0xf6, 0x4c, 0xcc, 0xb3, 0x65, 0x7b, 0x10, 0x25, 0x0a, 0x09, 0x27, 0x31, 0x15, 0x51,
	0xae, 0x18, 0x03, 0x7a, 0x9d, 0x1b, 0x85, 0x51, 0xb3, 0x45, 0xdf, 0x72, 0xa3, 0xd8, 0x3e, 0x8c,
	0x66, 0x95, 0x73, 0xda, 0xa4, 0x1b, 0x1c, 0x6c, 0x2f, 0x4a, 0x5b, 0x7e, 0xfa, 0xd7, 0xfc, 0xfb,
	0xd9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xd4, 0x3a, 0xb8, 0x85, 0x01, 0x00, 0x00,
}
