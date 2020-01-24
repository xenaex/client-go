// Code generated by protoc-gen-go. DO NOT EDIT.
// source: positions.proto

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

type Payment struct {
	PaymentType          string   `protobuf:"bytes,40213,opt,name=PaymentType,json=paymentType,proto3" json:"paymentType,omitempty"`
	PaymentCurrency      string   `protobuf:"bytes,40216,opt,name=PaymentCurrency,json=paymentCurrency,proto3" json:"paymentCurrency,omitempty"`
	PaymentAmount        string   `protobuf:"bytes,40217,opt,name=PaymentAmount,json=paymentAmount,proto3" json:"paymentAmount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2827456090109, []int{0}
}

func (m *Payment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payment.Unmarshal(m, b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
}
func (m *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(m, src)
}
func (m *Payment) XXX_Size() int {
	return xxx_messageInfo_Payment.Size(m)
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetPaymentType() string {
	if m != nil {
		return m.PaymentType
	}
	return ""
}

func (m *Payment) GetPaymentCurrency() string {
	if m != nil {
		return m.PaymentCurrency
	}
	return ""
}

func (m *Payment) GetPaymentAmount() string {
	if m != nil {
		return m.PaymentAmount
	}
	return ""
}

type RelatedTrade struct {
	RelatedTradeId       string   `protobuf:"bytes,1856,opt,name=RelatedTradeId,json=relatedTradeId,proto3" json:"relatedTradeId,omitempty"`
	RelatedTradeType     string   `protobuf:"bytes,1857,opt,name=RelatedTradeType,json=relatedTradeType,proto3" json:"relatedTradeType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RelatedTrade) Reset()         { *m = RelatedTrade{} }
func (m *RelatedTrade) String() string { return proto.CompactTextString(m) }
func (*RelatedTrade) ProtoMessage()    {}
func (*RelatedTrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2827456090109, []int{1}
}

func (m *RelatedTrade) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RelatedTrade.Unmarshal(m, b)
}
func (m *RelatedTrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RelatedTrade.Marshal(b, m, deterministic)
}
func (m *RelatedTrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RelatedTrade.Merge(m, src)
}
func (m *RelatedTrade) XXX_Size() int {
	return xxx_messageInfo_RelatedTrade.Size(m)
}
func (m *RelatedTrade) XXX_DiscardUnknown() {
	xxx_messageInfo_RelatedTrade.DiscardUnknown(m)
}

var xxx_messageInfo_RelatedTrade proto.InternalMessageInfo

func (m *RelatedTrade) GetRelatedTradeId() string {
	if m != nil {
		return m.RelatedTradeId
	}
	return ""
}

func (m *RelatedTrade) GetRelatedTradeType() string {
	if m != nil {
		return m.RelatedTradeType
	}
	return ""
}

type PositionReport struct {
	MsgType                      string          `protobuf:"bytes,35,opt,name=MsgType,json=msgType,proto3" json:"msgType,omitempty"`
	Account                      uint64          `protobuf:"varint,1,opt,name=Account,json=account,proto3" json:"account,omitempty"`
	PositionId                   uint64          `protobuf:"varint,2618,opt,name=PositionId,json=positionId,proto3" json:"positionId,omitempty"`
	TransactTime                 int64           `protobuf:"varint,60,opt,name=TransactTime,json=transactTime,proto3" json:"transactTime,omitempty"`
	Symbol                       string          `protobuf:"bytes,55,opt,name=Symbol,json=symbol,proto3" json:"symbol,omitempty"`
	PositionOpenTime             int64           `protobuf:"varint,1805,opt,name=PositionOpenTime,json=positionOpenTime,proto3" json:"positionOpenTime,omitempty"`
	AvgPx                        string          `protobuf:"bytes,6,opt,name=AvgPx,json=avgPx,proto3" json:"avgPx,omitempty"`
	Volume                       string          `protobuf:"bytes,53,opt,name=Volume,json=volume,proto3" json:"volume,omitempty"`
	Side                         string          `protobuf:"bytes,54,opt,name=Side,json=side,proto3" json:"side,omitempty"`
	SettlDate                    int64           `protobuf:"varint,64,opt,name=SettlDate,json=settlDate,proto3" json:"settlDate,omitempty"`
	SettlPrice                   string          `protobuf:"bytes,730,opt,name=SettlPrice,json=settlPrice,proto3" json:"settlPrice,omitempty"`
	PriorSettlPrice              string          `protobuf:"bytes,734,opt,name=PriorSettlPrice,json=priorSettlPrice,proto3" json:"priorSettlPrice,omitempty"`
	PreviousClearingBusinessDate int64           `protobuf:"varint,1084,opt,name=PreviousClearingBusinessDate,json=previousClearingBusinessDate,proto3" json:"previousClearingBusinessDate,omitempty"`
	ClearingBusinessDate         int64           `protobuf:"varint,715,opt,name=ClearingBusinessDate,json=clearingBusinessDate,proto3" json:"clearingBusinessDate,omitempty"`
	MarginAmounts                []*MarginAmount `protobuf:"bytes,1643,rep,name=MarginAmounts,json=marginAmounts,proto3" json:"marginAmounts,omitempty"`
	Payments                     []*Payment      `protobuf:"bytes,40212,rep,name=Payments,json=payments,proto3" json:"payments,omitempty"`
	RelatedTrades                []*RelatedTrade `protobuf:"bytes,1855,rep,name=RelatedTrades,json=relatedTrades,proto3" json:"relatedTrades,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}        `json:"-"`
	XXX_unrecognized             []byte          `json:"-"`
	XXX_sizecache                int32           `json:"-"`
}

func (m *PositionReport) Reset()         { *m = PositionReport{} }
func (m *PositionReport) String() string { return proto.CompactTextString(m) }
func (*PositionReport) ProtoMessage()    {}
func (*PositionReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2827456090109, []int{2}
}

func (m *PositionReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositionReport.Unmarshal(m, b)
}
func (m *PositionReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositionReport.Marshal(b, m, deterministic)
}
func (m *PositionReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionReport.Merge(m, src)
}
func (m *PositionReport) XXX_Size() int {
	return xxx_messageInfo_PositionReport.Size(m)
}
func (m *PositionReport) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionReport.DiscardUnknown(m)
}

var xxx_messageInfo_PositionReport proto.InternalMessageInfo

func (m *PositionReport) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *PositionReport) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *PositionReport) GetPositionId() uint64 {
	if m != nil {
		return m.PositionId
	}
	return 0
}

func (m *PositionReport) GetTransactTime() int64 {
	if m != nil {
		return m.TransactTime
	}
	return 0
}

func (m *PositionReport) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *PositionReport) GetPositionOpenTime() int64 {
	if m != nil {
		return m.PositionOpenTime
	}
	return 0
}

func (m *PositionReport) GetAvgPx() string {
	if m != nil {
		return m.AvgPx
	}
	return ""
}

func (m *PositionReport) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *PositionReport) GetSide() string {
	if m != nil {
		return m.Side
	}
	return ""
}

func (m *PositionReport) GetSettlDate() int64 {
	if m != nil {
		return m.SettlDate
	}
	return 0
}

func (m *PositionReport) GetSettlPrice() string {
	if m != nil {
		return m.SettlPrice
	}
	return ""
}

func (m *PositionReport) GetPriorSettlPrice() string {
	if m != nil {
		return m.PriorSettlPrice
	}
	return ""
}

func (m *PositionReport) GetPreviousClearingBusinessDate() int64 {
	if m != nil {
		return m.PreviousClearingBusinessDate
	}
	return 0
}

func (m *PositionReport) GetClearingBusinessDate() int64 {
	if m != nil {
		return m.ClearingBusinessDate
	}
	return 0
}

func (m *PositionReport) GetMarginAmounts() []*MarginAmount {
	if m != nil {
		return m.MarginAmounts
	}
	return nil
}

func (m *PositionReport) GetPayments() []*Payment {
	if m != nil {
		return m.Payments
	}
	return nil
}

func (m *PositionReport) GetRelatedTrades() []*RelatedTrade {
	if m != nil {
		return m.RelatedTrades
	}
	return nil
}

type MassPositionReport struct {
	MsgType         string            `protobuf:"bytes,35,opt,name=MsgType,json=msgType,proto3" json:"msgType,omitempty"`
	PosReqId        string            `protobuf:"bytes,710,opt,name=PosReqId,json=posReqId,proto3" json:"posReqId,omitempty"`
	Account         uint64            `protobuf:"varint,1,opt,name=Account,json=account,proto3" json:"account,omitempty"`
	TransactTime    int64             `protobuf:"varint,60,opt,name=TransactTime,json=transactTime,proto3" json:"transactTime,omitempty"`
	OpenPositions   []*PositionReport `protobuf:"bytes,727,rep,name=OpenPositions,json=openPositions,proto3" json:"openPositions,omitempty"`
	PositionHistory []*PositionReport `protobuf:"bytes,726,rep,name=PositionHistory,json=positionHistory,proto3" json:"positionHistory,omitempty"`
	// reject fields
	RejectReason         string   `protobuf:"bytes,380,opt,name=RejectReason,json=rejectReason,proto3" json:"rejectReason,omitempty"`
	Text                 string   `protobuf:"bytes,58,opt,name=Text,json=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MassPositionReport) Reset()         { *m = MassPositionReport{} }
func (m *MassPositionReport) String() string { return proto.CompactTextString(m) }
func (*MassPositionReport) ProtoMessage()    {}
func (*MassPositionReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2827456090109, []int{3}
}

func (m *MassPositionReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MassPositionReport.Unmarshal(m, b)
}
func (m *MassPositionReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MassPositionReport.Marshal(b, m, deterministic)
}
func (m *MassPositionReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MassPositionReport.Merge(m, src)
}
func (m *MassPositionReport) XXX_Size() int {
	return xxx_messageInfo_MassPositionReport.Size(m)
}
func (m *MassPositionReport) XXX_DiscardUnknown() {
	xxx_messageInfo_MassPositionReport.DiscardUnknown(m)
}

var xxx_messageInfo_MassPositionReport proto.InternalMessageInfo

func (m *MassPositionReport) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MassPositionReport) GetPosReqId() string {
	if m != nil {
		return m.PosReqId
	}
	return ""
}

func (m *MassPositionReport) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *MassPositionReport) GetTransactTime() int64 {
	if m != nil {
		return m.TransactTime
	}
	return 0
}

func (m *MassPositionReport) GetOpenPositions() []*PositionReport {
	if m != nil {
		return m.OpenPositions
	}
	return nil
}

func (m *MassPositionReport) GetPositionHistory() []*PositionReport {
	if m != nil {
		return m.PositionHistory
	}
	return nil
}

func (m *MassPositionReport) GetRejectReason() string {
	if m != nil {
		return m.RejectReason
	}
	return ""
}

func (m *MassPositionReport) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type PositionsRequest struct {
	MsgType              string   `protobuf:"bytes,35,opt,name=MsgType,json=msgType,proto3" json:"msgType,omitempty"`
	PosReqId             string   `protobuf:"bytes,710,opt,name=PosReqId,json=posReqId,proto3" json:"posReqId,omitempty"`
	Account              uint64   `protobuf:"varint,1,opt,name=Account,json=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PositionsRequest) Reset()         { *m = PositionsRequest{} }
func (m *PositionsRequest) String() string { return proto.CompactTextString(m) }
func (*PositionsRequest) ProtoMessage()    {}
func (*PositionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2827456090109, []int{4}
}

func (m *PositionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositionsRequest.Unmarshal(m, b)
}
func (m *PositionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositionsRequest.Marshal(b, m, deterministic)
}
func (m *PositionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionsRequest.Merge(m, src)
}
func (m *PositionsRequest) XXX_Size() int {
	return xxx_messageInfo_PositionsRequest.Size(m)
}
func (m *PositionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PositionsRequest proto.InternalMessageInfo

func (m *PositionsRequest) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *PositionsRequest) GetPosReqId() string {
	if m != nil {
		return m.PosReqId
	}
	return ""
}

func (m *PositionsRequest) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

type PositionMaintenanceRequest struct {
	MsgType              string   `protobuf:"bytes,35,opt,name=MsgType,json=msgType,proto3" json:"msgType,omitempty"`
	PosReqId             string   `protobuf:"bytes,710,opt,name=PosReqId,json=posReqId,proto3" json:"posReqId,omitempty"`
	PosTransType         string   `protobuf:"bytes,709,opt,name=PosTransType,json=posTransType,proto3" json:"posTransType,omitempty"`
	PosMaintAction       string   `protobuf:"bytes,712,opt,name=PosMaintAction,json=posMaintAction,proto3" json:"posMaintAction,omitempty"`
	Account              uint64   `protobuf:"varint,1,opt,name=Account,json=account,proto3" json:"account,omitempty"`
	Symbol               string   `protobuf:"bytes,55,opt,name=Symbol,json=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PositionMaintenanceRequest) Reset()         { *m = PositionMaintenanceRequest{} }
func (m *PositionMaintenanceRequest) String() string { return proto.CompactTextString(m) }
func (*PositionMaintenanceRequest) ProtoMessage()    {}
func (*PositionMaintenanceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2827456090109, []int{5}
}

func (m *PositionMaintenanceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositionMaintenanceRequest.Unmarshal(m, b)
}
func (m *PositionMaintenanceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositionMaintenanceRequest.Marshal(b, m, deterministic)
}
func (m *PositionMaintenanceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionMaintenanceRequest.Merge(m, src)
}
func (m *PositionMaintenanceRequest) XXX_Size() int {
	return xxx_messageInfo_PositionMaintenanceRequest.Size(m)
}
func (m *PositionMaintenanceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionMaintenanceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PositionMaintenanceRequest proto.InternalMessageInfo

func (m *PositionMaintenanceRequest) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *PositionMaintenanceRequest) GetPosReqId() string {
	if m != nil {
		return m.PosReqId
	}
	return ""
}

func (m *PositionMaintenanceRequest) GetPosTransType() string {
	if m != nil {
		return m.PosTransType
	}
	return ""
}

func (m *PositionMaintenanceRequest) GetPosMaintAction() string {
	if m != nil {
		return m.PosMaintAction
	}
	return ""
}

func (m *PositionMaintenanceRequest) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *PositionMaintenanceRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type PositionMaintenanceReport struct {
	MsgType              string   `protobuf:"bytes,35,opt,name=MsgType,json=msgType,proto3" json:"msgType,omitempty"`
	PosReqId             string   `protobuf:"bytes,710,opt,name=PosReqId,json=posReqId,proto3" json:"posReqId,omitempty"`
	PosTransType         string   `protobuf:"bytes,709,opt,name=PosTransType,json=posTransType,proto3" json:"posTransType,omitempty"`
	PosMaintAction       string   `protobuf:"bytes,712,opt,name=PosMaintAction,json=posMaintAction,proto3" json:"posMaintAction,omitempty"`
	Account              uint64   `protobuf:"varint,1,opt,name=Account,json=account,proto3" json:"account,omitempty"`
	Symbol               string   `protobuf:"bytes,55,opt,name=Symbol,json=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PositionMaintenanceReport) Reset()         { *m = PositionMaintenanceReport{} }
func (m *PositionMaintenanceReport) String() string { return proto.CompactTextString(m) }
func (*PositionMaintenanceReport) ProtoMessage()    {}
func (*PositionMaintenanceReport) Descriptor() ([]byte, []int) {
	return fileDescriptor_86e2827456090109, []int{6}
}

func (m *PositionMaintenanceReport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PositionMaintenanceReport.Unmarshal(m, b)
}
func (m *PositionMaintenanceReport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PositionMaintenanceReport.Marshal(b, m, deterministic)
}
func (m *PositionMaintenanceReport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PositionMaintenanceReport.Merge(m, src)
}
func (m *PositionMaintenanceReport) XXX_Size() int {
	return xxx_messageInfo_PositionMaintenanceReport.Size(m)
}
func (m *PositionMaintenanceReport) XXX_DiscardUnknown() {
	xxx_messageInfo_PositionMaintenanceReport.DiscardUnknown(m)
}

var xxx_messageInfo_PositionMaintenanceReport proto.InternalMessageInfo

func (m *PositionMaintenanceReport) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *PositionMaintenanceReport) GetPosReqId() string {
	if m != nil {
		return m.PosReqId
	}
	return ""
}

func (m *PositionMaintenanceReport) GetPosTransType() string {
	if m != nil {
		return m.PosTransType
	}
	return ""
}

func (m *PositionMaintenanceReport) GetPosMaintAction() string {
	if m != nil {
		return m.PosMaintAction
	}
	return ""
}

func (m *PositionMaintenanceReport) GetAccount() uint64 {
	if m != nil {
		return m.Account
	}
	return 0
}

func (m *PositionMaintenanceReport) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func init() {
	proto.RegisterType((*Payment)(nil), "xmsg.Payment")
	proto.RegisterType((*RelatedTrade)(nil), "xmsg.RelatedTrade")
	proto.RegisterType((*PositionReport)(nil), "xmsg.PositionReport")
	proto.RegisterType((*MassPositionReport)(nil), "xmsg.MassPositionReport")
	proto.RegisterType((*PositionsRequest)(nil), "xmsg.PositionsRequest")
	proto.RegisterType((*PositionMaintenanceRequest)(nil), "xmsg.PositionMaintenanceRequest")
	proto.RegisterType((*PositionMaintenanceReport)(nil), "xmsg.PositionMaintenanceReport")
}

func init() { proto.RegisterFile("positions.proto", fileDescriptor_86e2827456090109) }

var fileDescriptor_86e2827456090109 = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x55, 0x4f, 0x6b, 0x13, 0x4f,
	0x18, 0x66, 0xdb, 0x6c, 0x92, 0xbe, 0xdd, 0xdd, 0x94, 0x21, 0xfc, 0x98, 0x5f, 0x2d, 0x18, 0xb6,
	0xa0, 0xd1, 0x42, 0x0f, 0x16, 0x15, 0xff, 0x80, 0xc6, 0x7a, 0xb0, 0x87, 0xe0, 0xb2, 0x0d, 0xde,
	0xa7, 0xbb, 0x43, 0x58, 0xc9, 0xee, 0x8c, 0x33, 0x93, 0x92, 0x1c, 0xfd, 0x00, 0xde, 0x14, 0xf5,
	0xab, 0x14, 0xc1, 0x7a, 0x50, 0x04, 0x6f, 0x0a, 0x0a, 0x1e, 0xfc, 0x02, 0x7e, 0x05, 0x0f, 0xb2,
	0xb3, 0x9b, 0x76, 0x52, 0x62, 0x28, 0x54, 0xf0, 0x96, 0xf7, 0x79, 0x9f, 0xe7, 0x7d, 0xe7, 0x9d,
	0x7d, 0xe6, 0x0d, 0x34, 0x38, 0x93, 0x89, 0x4a, 0x58, 0x26, 0x37, 0xb9, 0x60, 0x8a, 0xa1, 0xca,
	0x28, 0x95, 0xfd, 0x55, 0x27, 0x25, 0xa2, 0x9f, 0x64, 0x05, 0xe6, 0x3f, 0xb5, 0xa0, 0x16, 0x90,
	0x71, 0x4a, 0x33, 0x85, 0x7c, 0x58, 0x2e, 0x7f, 0xf6, 0xc6, 0x9c, 0xe2, 0x17, 0x07, 0x0b, 0x2d,
	0xab, 0xbd, 0x14, 0x2e, 0xf3, 0x63, 0x10, 0x5d, 0x86, 0x46, 0xc9, 0xd9, 0x1e, 0x0a, 0x41, 0xb3,
	0x68, 0x8c, 0x5f, 0x95, 0xbc, 0x06, 0x9f, 0x4e, 0xa0, 0x0b, 0xe0, 0x96, 0xdc, 0x4e, 0xca, 0x86,
	0x99, 0xc2, 0xaf, 0x4b, 0xa6, 0xcb, 0x4d, 0xd8, 0x8f, 0xc1, 0x09, 0xe9, 0x80, 0x28, 0x1a, 0xf7,
	0x04, 0x89, 0x29, 0xba, 0x08, 0x9e, 0x19, 0xef, 0xc4, 0xf8, 0xd0, 0xd3, 0x3a, 0x4f, 0x4c, 0xc1,
	0x68, 0x03, 0x56, 0x4c, 0xa2, 0x3e, 0xf5, 0xbb, 0x82, 0xba, 0x22, 0x4e, 0x24, 0xfc, 0x97, 0x36,
	0x78, 0x41, 0x79, 0x23, 0x21, 0xe5, 0x4c, 0x28, 0x84, 0xa1, 0xd6, 0x95, 0x7d, 0x2d, 0x5b, 0xd7,
	0xaa, 0x5a, 0x5a, 0x84, 0x79, 0xa6, 0x13, 0x45, 0xfa, 0xd0, 0x56, 0xcb, 0x6a, 0x57, 0xc2, 0x1a,
	0x29, 0x42, 0x74, 0x1e, 0x60, 0x52, 0x65, 0x27, 0xc6, 0x07, 0x4d, 0x9d, 0x05, 0x7e, 0x04, 0x21,
	0x1f, 0x9c, 0x9e, 0x20, 0x99, 0x24, 0x91, 0xea, 0x25, 0x29, 0xc5, 0xb7, 0x5b, 0x56, 0x7b, 0x31,
	0x74, 0x94, 0x81, 0xa1, 0xff, 0xa0, 0xba, 0x3b, 0x4e, 0xf7, 0xd8, 0x00, 0x5f, 0xd7, 0x7d, 0xab,
	0x52, 0x47, 0xf9, 0x40, 0x93, 0xe2, 0x0f, 0x39, 0xcd, 0xb4, 0xfe, 0x99, 0xa7, 0x0b, 0xac, 0xf0,
	0x13, 0x09, 0xd4, 0x04, 0xbb, 0xb3, 0xdf, 0x0f, 0x46, 0xb8, 0xaa, 0x6b, 0xd8, 0x24, 0x0f, 0xf2,
	0xd2, 0x8f, 0xd8, 0x60, 0x98, 0x52, 0x7c, 0xb5, 0x28, 0xbd, 0xaf, 0x23, 0x84, 0xa0, 0xb2, 0x9b,
	0xc4, 0x14, 0x5f, 0xd3, 0x68, 0x45, 0x26, 0x31, 0x45, 0x6b, 0xb0, 0xb4, 0x4b, 0x95, 0x1a, 0xdc,
	0x27, 0x8a, 0xe2, 0xbb, 0xba, 0xcd, 0x92, 0x9c, 0x00, 0xf9, 0xa4, 0x3a, 0x1b, 0x88, 0x24, 0xa2,
	0xf8, 0xbb, 0xad, 0x85, 0x20, 0x8f, 0x20, 0x74, 0x09, 0x1a, 0x81, 0x48, 0x98, 0x30, 0x58, 0x3f,
	0xec, 0xd2, 0x0a, 0xd3, 0x38, 0xda, 0x86, 0xb5, 0x40, 0xd0, 0xfd, 0x84, 0x0d, 0xe5, 0xf6, 0x80,
	0x12, 0x91, 0x64, 0xfd, 0x7b, 0x43, 0x99, 0x64, 0x54, 0x4a, 0xdd, 0xfc, 0x4d, 0x5d, 0x77, 0x5f,
	0xe3, 0x73, 0x48, 0x68, 0x0b, 0x9a, 0x33, 0xc5, 0x9f, 0x6c, 0x2d, 0x6e, 0x46, 0xb3, 0x44, 0x37,
	0xc0, 0xed, 0x6a, 0xc3, 0x17, 0x66, 0x93, 0xf8, 0xa7, 0xd3, 0x5a, 0x6c, 0x2f, 0x5f, 0x41, 0x9b,
	0xf9, 0x6b, 0xd8, 0x34, 0x73, 0xa1, 0x9b, 0x9a, 0x4c, 0xb4, 0x01, 0xf5, 0xd2, 0xbf, 0x12, 0x3f,
	0x3f, 0x58, 0xd0, 0x32, 0xb7, 0x90, 0x95, 0x78, 0x58, 0x2f, 0x9d, 0x2c, 0xf3, 0x3e, 0xa6, 0x17,
	0x25, 0x7e, 0xeb, 0x99, 0x7d, 0xcc, 0x5c, 0xe8, 0x9a, 0xe6, 0x94, 0xfe, 0xe1, 0x02, 0xa0, 0x2e,
	0x91, 0xf2, 0xd4, 0xee, 0x3c, 0x07, 0xf5, 0x80, 0xc9, 0x90, 0x3e, 0xd9, 0x89, 0xf1, 0x87, 0xe2,
	0xc6, 0xeb, 0xbc, 0x04, 0xe6, 0x58, 0xf7, 0x34, 0xce, 0xbc, 0x05, 0x6e, 0x6e, 0xb0, 0xc9, 0x51,
	0x24, 0xfe, 0x66, 0xeb, 0x31, 0x9a, 0xe5, 0xdc, 0x53, 0x47, 0x0c, 0x5d, 0x66, 0x72, 0xd1, 0x1d,
	0x68, 0x4c, 0x82, 0x07, 0x89, 0x54, 0x4c, 0x8c, 0xf1, 0xd7, 0x79, 0xf2, 0xa3, 0x0d, 0x55, 0xb2,
	0xd1, 0x7a, 0xbe, 0x09, 0x1e, 0xd3, 0x48, 0x85, 0x94, 0x48, 0x96, 0xe1, 0x5f, 0xc5, 0xbe, 0x70,
	0x84, 0x01, 0xe6, 0x4e, 0xee, 0xd1, 0x91, 0xc2, 0x37, 0x0b, 0x27, 0x2b, 0x3a, 0x52, 0x7e, 0x74,
	0xfc, 0x70, 0xf2, 0x6b, 0x18, 0x52, 0xf9, 0xf7, 0xef, 0xcf, 0xff, 0x62, 0xc1, 0xea, 0xa4, 0x4b,
	0x97, 0x24, 0x99, 0xa2, 0x19, 0xc9, 0x22, 0x7a, 0xc6, 0x7e, 0xeb, 0xe0, 0x04, 0x4c, 0xea, 0x0f,
	0xa3, 0xb5, 0xef, 0x0b, 0x82, 0xc3, 0x0d, 0x30, 0x5f, 0x89, 0x01, 0x93, 0xba, 0x69, 0x27, 0xca,
	0xfb, 0xe3, 0x8f, 0x05, 0xcd, 0xe3, 0x53, 0xf0, 0x9c, 0xaf, 0xff, 0x87, 0x9d, 0xe3, 0x7f, 0xb6,
	0xe0, 0xff, 0x99, 0x53, 0x9d, 0xc5, 0x84, 0xff, 0x78, 0xa8, 0xbd, 0xaa, 0xfe, 0x77, 0xdb, 0xfa,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xc7, 0x8b, 0x38, 0x04, 0x07, 0x00, 0x00,
}
