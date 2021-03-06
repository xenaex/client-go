// Code generated by protoc-gen-go. DO NOT EDIT.
// source: market.proto

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

type MDEntry struct {
	Symbol               string   `protobuf:"bytes,55,opt,name=Symbol,json=symbol,proto3" json:"symbol,omitempty"`
	MDUpdateAction       string   `protobuf:"bytes,279,opt,name=MDUpdateAction,json=mdUpdateAction,proto3" json:"mdUpdateAction,omitempty"`
	MDEntryType          string   `protobuf:"bytes,269,opt,name=MDEntryType,json=mdEntryType,proto3" json:"mdEntryType,omitempty"`
	MDEntryPx            string   `protobuf:"bytes,270,opt,name=MDEntryPx,json=mdEntryPx,proto3" json:"mdEntryPx,omitempty"`
	MDEntrySize          string   `protobuf:"bytes,271,opt,name=MDEntrySize,json=mdEntrySize,proto3" json:"mdEntrySize,omitempty"`
	NumberOfOrders       uint32   `protobuf:"varint,346,opt,name=NumberOfOrders,json=numberOfOrders,proto3" json:"numberOfOrders,omitempty"`
	TransactTime         int64    `protobuf:"varint,60,opt,name=TransactTime,json=transactTime,proto3" json:"transactTime,omitempty"`
	TradeId              string   `protobuf:"bytes,1003,opt,name=TradeId,json=tradeId,proto3" json:"tradeId,omitempty"`
	AggressorSide        string   `protobuf:"bytes,1501,opt,name=AggressorSide,json=aggressorSide,proto3" json:"aggressorSide,omitempty"`
	FirstPx              string   `protobuf:"bytes,1025,opt,name=FirstPx,json=firstPx,proto3" json:"firstPx,omitempty"`
	LastPx               string   `protobuf:"bytes,31,opt,name=LastPx,json=lastPx,proto3" json:"lastPx,omitempty"`
	HighPx               string   `protobuf:"bytes,332,opt,name=HighPx,json=highPx,proto3" json:"highPx,omitempty"`
	LowPx                string   `protobuf:"bytes,333,opt,name=LowPx,json=lowPx,proto3" json:"lowPx,omitempty"`
	BuyVolume            string   `protobuf:"bytes,330,opt,name=BuyVolume,json=buyVolume,proto3" json:"buyVolume,omitempty"`
	SellVolume           string   `protobuf:"bytes,331,opt,name=SellVolume,json=sellVolume,proto3" json:"sellVolume,omitempty"`
	Bid                  string   `protobuf:"bytes,1502,opt,name=Bid,json=bid,proto3" json:"bid,omitempty"`
	Ask                  string   `protobuf:"bytes,1503,opt,name=Ask,json=ask,proto3" json:"ask,omitempty"`
	LowRangePx           string   `protobuf:"bytes,35601,opt,name=LowRangePx,json=lowRangePx,proto3" json:"lowRangePx,omitempty"`
	HighRangePx          string   `protobuf:"bytes,35602,opt,name=HighRangePx,json=highRangePx,proto3" json:"highRangePx,omitempty"`
	LowLimitPx           string   `protobuf:"bytes,35603,opt,name=LowLimitPx,json=lowLimitPx,proto3" json:"lowLimitPx,omitempty"`
	HighLimitPx          string   `protobuf:"bytes,35604,opt,name=HighLimitPx,json=highLimitPx,proto3" json:"highLimitPx,omitempty"`
	ClearingPx           string   `protobuf:"bytes,35605,opt,name=ClearingPx,json=clearingPx,proto3" json:"clearingPx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MDEntry) Reset()         { *m = MDEntry{} }
func (m *MDEntry) String() string { return proto.CompactTextString(m) }
func (*MDEntry) ProtoMessage()    {}
func (*MDEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{0}
}

func (m *MDEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MDEntry.Unmarshal(m, b)
}
func (m *MDEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MDEntry.Marshal(b, m, deterministic)
}
func (m *MDEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MDEntry.Merge(m, src)
}
func (m *MDEntry) XXX_Size() int {
	return xxx_messageInfo_MDEntry.Size(m)
}
func (m *MDEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_MDEntry.DiscardUnknown(m)
}

var xxx_messageInfo_MDEntry proto.InternalMessageInfo

func (m *MDEntry) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MDEntry) GetMDUpdateAction() string {
	if m != nil {
		return m.MDUpdateAction
	}
	return ""
}

func (m *MDEntry) GetMDEntryType() string {
	if m != nil {
		return m.MDEntryType
	}
	return ""
}

func (m *MDEntry) GetMDEntryPx() string {
	if m != nil {
		return m.MDEntryPx
	}
	return ""
}

func (m *MDEntry) GetMDEntrySize() string {
	if m != nil {
		return m.MDEntrySize
	}
	return ""
}

func (m *MDEntry) GetNumberOfOrders() uint32 {
	if m != nil {
		return m.NumberOfOrders
	}
	return 0
}

func (m *MDEntry) GetTransactTime() int64 {
	if m != nil {
		return m.TransactTime
	}
	return 0
}

func (m *MDEntry) GetTradeId() string {
	if m != nil {
		return m.TradeId
	}
	return ""
}

func (m *MDEntry) GetAggressorSide() string {
	if m != nil {
		return m.AggressorSide
	}
	return ""
}

func (m *MDEntry) GetFirstPx() string {
	if m != nil {
		return m.FirstPx
	}
	return ""
}

func (m *MDEntry) GetLastPx() string {
	if m != nil {
		return m.LastPx
	}
	return ""
}

func (m *MDEntry) GetHighPx() string {
	if m != nil {
		return m.HighPx
	}
	return ""
}

func (m *MDEntry) GetLowPx() string {
	if m != nil {
		return m.LowPx
	}
	return ""
}

func (m *MDEntry) GetBuyVolume() string {
	if m != nil {
		return m.BuyVolume
	}
	return ""
}

func (m *MDEntry) GetSellVolume() string {
	if m != nil {
		return m.SellVolume
	}
	return ""
}

func (m *MDEntry) GetBid() string {
	if m != nil {
		return m.Bid
	}
	return ""
}

func (m *MDEntry) GetAsk() string {
	if m != nil {
		return m.Ask
	}
	return ""
}

func (m *MDEntry) GetLowRangePx() string {
	if m != nil {
		return m.LowRangePx
	}
	return ""
}

func (m *MDEntry) GetHighRangePx() string {
	if m != nil {
		return m.HighRangePx
	}
	return ""
}

func (m *MDEntry) GetLowLimitPx() string {
	if m != nil {
		return m.LowLimitPx
	}
	return ""
}

func (m *MDEntry) GetHighLimitPx() string {
	if m != nil {
		return m.HighLimitPx
	}
	return ""
}

func (m *MDEntry) GetClearingPx() string {
	if m != nil {
		return m.ClearingPx
	}
	return ""
}

type MarketDataRefresh struct {
	MsgType              string     `protobuf:"bytes,35,opt,name=MsgType,json=msgType,proto3" json:"msgType,omitempty"`
	MDStreamId           string     `protobuf:"bytes,1500,opt,name=MDStreamId,json=mdStreamId,proto3" json:"mdStreamId,omitempty"`
	LastUpdateTime       int64      `protobuf:"varint,779,opt,name=LastUpdateTime,json=lastUpdateTime,proto3" json:"lastUpdateTime,omitempty"`
	MDBookType           string     `protobuf:"bytes,1021,opt,name=MDBookType,json=mdBookType,proto3" json:"mdBookType,omitempty"`
	Symbol               string     `protobuf:"bytes,55,opt,name=Symbol,json=symbol,proto3" json:"symbol,omitempty"`
	LowRangePx           string     `protobuf:"bytes,35601,opt,name=LowRangePx,json=lowRangePx,proto3" json:"lowRangePx,omitempty"`
	HighRangePx          string     `protobuf:"bytes,35602,opt,name=HighRangePx,json=highRangePx,proto3" json:"highRangePx,omitempty"`
	LowLimitPx           string     `protobuf:"bytes,35603,opt,name=LowLimitPx,json=lowLimitPx,proto3" json:"lowLimitPx,omitempty"`
	HighLimitPx          string     `protobuf:"bytes,35604,opt,name=HighLimitPx,json=highLimitPx,proto3" json:"highLimitPx,omitempty"`
	ClearingPx           string     `protobuf:"bytes,35605,opt,name=ClearingPx,json=clearingPx,proto3" json:"clearingPx,omitempty"`
	BestBid              string     `protobuf:"bytes,1502,opt,name=BestBid,json=bestBid,proto3" json:"bestBid,omitempty"`
	BestAsk              string     `protobuf:"bytes,1503,opt,name=BestAsk,json=bestAsk,proto3" json:"bestAsk,omitempty"`
	MDEntry              []*MDEntry `protobuf:"bytes,268,rep,name=MDEntry,json=mdEntry,proto3" json:"mdEntry,omitempty"`
	Ratios               []*MDEntry `protobuf:"bytes,1504,rep,name=Ratios,json=ratios,proto3" json:"ratios,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MarketDataRefresh) Reset()         { *m = MarketDataRefresh{} }
func (m *MarketDataRefresh) String() string { return proto.CompactTextString(m) }
func (*MarketDataRefresh) ProtoMessage()    {}
func (*MarketDataRefresh) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{1}
}

func (m *MarketDataRefresh) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketDataRefresh.Unmarshal(m, b)
}
func (m *MarketDataRefresh) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketDataRefresh.Marshal(b, m, deterministic)
}
func (m *MarketDataRefresh) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketDataRefresh.Merge(m, src)
}
func (m *MarketDataRefresh) XXX_Size() int {
	return xxx_messageInfo_MarketDataRefresh.Size(m)
}
func (m *MarketDataRefresh) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketDataRefresh.DiscardUnknown(m)
}

var xxx_messageInfo_MarketDataRefresh proto.InternalMessageInfo

func (m *MarketDataRefresh) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MarketDataRefresh) GetMDStreamId() string {
	if m != nil {
		return m.MDStreamId
	}
	return ""
}

func (m *MarketDataRefresh) GetLastUpdateTime() int64 {
	if m != nil {
		return m.LastUpdateTime
	}
	return 0
}

func (m *MarketDataRefresh) GetMDBookType() string {
	if m != nil {
		return m.MDBookType
	}
	return ""
}

func (m *MarketDataRefresh) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *MarketDataRefresh) GetLowRangePx() string {
	if m != nil {
		return m.LowRangePx
	}
	return ""
}

func (m *MarketDataRefresh) GetHighRangePx() string {
	if m != nil {
		return m.HighRangePx
	}
	return ""
}

func (m *MarketDataRefresh) GetLowLimitPx() string {
	if m != nil {
		return m.LowLimitPx
	}
	return ""
}

func (m *MarketDataRefresh) GetHighLimitPx() string {
	if m != nil {
		return m.HighLimitPx
	}
	return ""
}

func (m *MarketDataRefresh) GetClearingPx() string {
	if m != nil {
		return m.ClearingPx
	}
	return ""
}

func (m *MarketDataRefresh) GetBestBid() string {
	if m != nil {
		return m.BestBid
	}
	return ""
}

func (m *MarketDataRefresh) GetBestAsk() string {
	if m != nil {
		return m.BestAsk
	}
	return ""
}

func (m *MarketDataRefresh) GetMDEntry() []*MDEntry {
	if m != nil {
		return m.MDEntry
	}
	return nil
}

func (m *MarketDataRefresh) GetRatios() []*MDEntry {
	if m != nil {
		return m.Ratios
	}
	return nil
}

type MarketDataRequest struct {
	MsgType                 string   `protobuf:"bytes,35,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	MDStreamId              string   `protobuf:"bytes,1500,opt,name=MDStreamId,proto3" json:"MDStreamId,omitempty"`
	SubscriptionRequestType string   `protobuf:"bytes,263,opt,name=SubscriptionRequestType,proto3" json:"SubscriptionRequestType,omitempty"`
	ThrottleType            string   `protobuf:"bytes,1612,opt,name=ThrottleType,proto3" json:"ThrottleType,omitempty"`
	ThrottleTimeInterval    int64    `protobuf:"varint,1614,opt,name=ThrottleTimeInterval,proto3" json:"ThrottleTimeInterval,omitempty"`
	ThrottleTimeUnit        string   `protobuf:"bytes,1615,opt,name=ThrottleTimeUnit,proto3" json:"ThrottleTimeUnit,omitempty"`
	AggregatedBook          int64    `protobuf:"varint,266,opt,name=AggregatedBook,proto3" json:"AggregatedBook,omitempty"`
	MarketDepth             int64    `protobuf:"varint,264,opt,name=MarketDepth,proto3" json:"MarketDepth,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *MarketDataRequest) Reset()         { *m = MarketDataRequest{} }
func (m *MarketDataRequest) String() string { return proto.CompactTextString(m) }
func (*MarketDataRequest) ProtoMessage()    {}
func (*MarketDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{2}
}

func (m *MarketDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketDataRequest.Unmarshal(m, b)
}
func (m *MarketDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketDataRequest.Marshal(b, m, deterministic)
}
func (m *MarketDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketDataRequest.Merge(m, src)
}
func (m *MarketDataRequest) XXX_Size() int {
	return xxx_messageInfo_MarketDataRequest.Size(m)
}
func (m *MarketDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MarketDataRequest proto.InternalMessageInfo

func (m *MarketDataRequest) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MarketDataRequest) GetMDStreamId() string {
	if m != nil {
		return m.MDStreamId
	}
	return ""
}

func (m *MarketDataRequest) GetSubscriptionRequestType() string {
	if m != nil {
		return m.SubscriptionRequestType
	}
	return ""
}

func (m *MarketDataRequest) GetThrottleType() string {
	if m != nil {
		return m.ThrottleType
	}
	return ""
}

func (m *MarketDataRequest) GetThrottleTimeInterval() int64 {
	if m != nil {
		return m.ThrottleTimeInterval
	}
	return 0
}

func (m *MarketDataRequest) GetThrottleTimeUnit() string {
	if m != nil {
		return m.ThrottleTimeUnit
	}
	return ""
}

func (m *MarketDataRequest) GetAggregatedBook() int64 {
	if m != nil {
		return m.AggregatedBook
	}
	return 0
}

func (m *MarketDataRequest) GetMarketDepth() int64 {
	if m != nil {
		return m.MarketDepth
	}
	return 0
}

type MarketDataRequestReject struct {
	MsgType              string   `protobuf:"bytes,35,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	MDStreamId           string   `protobuf:"bytes,1500,opt,name=MDStreamId,proto3" json:"MDStreamId,omitempty"`
	RejectText           string   `protobuf:"bytes,1328,opt,name=RejectText,proto3" json:"RejectText,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarketDataRequestReject) Reset()         { *m = MarketDataRequestReject{} }
func (m *MarketDataRequestReject) String() string { return proto.CompactTextString(m) }
func (*MarketDataRequestReject) ProtoMessage()    {}
func (*MarketDataRequestReject) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{3}
}

func (m *MarketDataRequestReject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarketDataRequestReject.Unmarshal(m, b)
}
func (m *MarketDataRequestReject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarketDataRequestReject.Marshal(b, m, deterministic)
}
func (m *MarketDataRequestReject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarketDataRequestReject.Merge(m, src)
}
func (m *MarketDataRequestReject) XXX_Size() int {
	return xxx_messageInfo_MarketDataRequestReject.Size(m)
}
func (m *MarketDataRequestReject) XXX_DiscardUnknown() {
	xxx_messageInfo_MarketDataRequestReject.DiscardUnknown(m)
}

var xxx_messageInfo_MarketDataRequestReject proto.InternalMessageInfo

func (m *MarketDataRequestReject) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *MarketDataRequestReject) GetMDStreamId() string {
	if m != nil {
		return m.MDStreamId
	}
	return ""
}

func (m *MarketDataRequestReject) GetRejectText() string {
	if m != nil {
		return m.RejectText
	}
	return ""
}

type Bars struct {
	MDEntry              []*MDEntry `protobuf:"bytes,268,rep,name=MDEntry,json=mdEntry,proto3" json:"mdEntry,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Bars) Reset()         { *m = Bars{} }
func (m *Bars) String() string { return proto.CompactTextString(m) }
func (*Bars) ProtoMessage()    {}
func (*Bars) Descriptor() ([]byte, []int) {
	return fileDescriptor_3f90997f23a2c3f8, []int{4}
}

func (m *Bars) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Bars.Unmarshal(m, b)
}
func (m *Bars) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Bars.Marshal(b, m, deterministic)
}
func (m *Bars) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bars.Merge(m, src)
}
func (m *Bars) XXX_Size() int {
	return xxx_messageInfo_Bars.Size(m)
}
func (m *Bars) XXX_DiscardUnknown() {
	xxx_messageInfo_Bars.DiscardUnknown(m)
}

var xxx_messageInfo_Bars proto.InternalMessageInfo

func (m *Bars) GetMDEntry() []*MDEntry {
	if m != nil {
		return m.MDEntry
	}
	return nil
}

func init() {
	proto.RegisterType((*MDEntry)(nil), "xmsg.MDEntry")
	proto.RegisterType((*MarketDataRefresh)(nil), "xmsg.MarketDataRefresh")
	proto.RegisterType((*MarketDataRequest)(nil), "xmsg.MarketDataRequest")
	proto.RegisterType((*MarketDataRequestReject)(nil), "xmsg.MarketDataRequestReject")
	proto.RegisterType((*Bars)(nil), "xmsg.Bars")
}

func init() { proto.RegisterFile("market.proto", fileDescriptor_3f90997f23a2c3f8) }

var fileDescriptor_3f90997f23a2c3f8 = []byte{
	// 740 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x95, 0xcb, 0x6e, 0x13, 0x3f,
	0x14, 0xc6, 0xd5, 0xa4, 0x9d, 0xf9, 0xf7, 0xe4, 0xa2, 0x7f, 0x2d, 0xa0, 0x66, 0x81, 0x1a, 0x52,
	0x01, 0x91, 0x90, 0x2a, 0x44, 0x17, 0x08, 0x89, 0x4d, 0x42, 0x40, 0x54, 0x4a, 0x68, 0x34, 0x49,
	0xd9, 0x3b, 0x19, 0x77, 0x32, 0x64, 0x2e, 0xc1, 0x76, 0xe8, 0x84, 0x1d, 0x2b, 0x10, 0x57, 0x71,
	0x7f, 0x0d, 0x36, 0xbc, 0x04, 0x14, 0xd8, 0xa3, 0x72, 0x59, 0xb1, 0xe1, 0x19, 0x90, 0x90, 0xed,
	0x99, 0x74, 0xa2, 0xb6, 0x08, 0x89, 0x15, 0xcb, 0xf3, 0x3b, 0x9f, 0xbe, 0xd8, 0xc7, 0xdf, 0x9c,
	0x40, 0xde, 0x27, 0x6c, 0x40, 0xc5, 0xca, 0x90, 0x85, 0x22, 0x44, 0xb3, 0x91, 0xcf, 0x9d, 0xf2,
	0x9b, 0x39, 0x30, 0x9b, 0xf5, 0x4b, 0x81, 0x60, 0x63, 0x74, 0x04, 0x8c, 0xf6, 0xd8, 0xef, 0x86,
	0x1e, 0x3e, 0x57, 0x9a, 0xa9, 0xcc, 0x5b, 0x06, 0x57, 0x15, 0x3a, 0x05, 0xc5, 0x66, 0x7d, 0x63,
	0x68, 0x13, 0x41, 0xab, 0x3d, 0xe1, 0x86, 0x01, 0x7e, 0x95, 0x51, 0x82, 0xa2, 0x6f, 0xa7, 0x31,
	0x3a, 0x0e, 0xb9, 0xd8, 0xab, 0x33, 0x1e, 0x52, 0xfc, 0x50, 0xab, 0x72, 0xbe, 0x3d, 0x61, 0xe8,
	0x18, 0xcc, 0xc7, 0x92, 0x56, 0x84, 0x1f, 0x69, 0xc1, 0x7c, 0x2c, 0x68, 0x45, 0x29, 0x87, 0xb6,
	0x7b, 0x8b, 0xe2, 0xc7, 0xd3, 0x0e, 0x92, 0xc9, 0xd3, 0x5c, 0x1d, 0xf9, 0x5d, 0xca, 0xd6, 0x37,
	0xd7, 0x99, 0x4d, 0x19, 0xc7, 0x9f, 0xa4, 0xaa, 0x60, 0x15, 0x83, 0x29, 0x8c, 0xca, 0x90, 0xef,
	0x30, 0x12, 0x70, 0xd2, 0x13, 0x1d, 0xd7, 0xa7, 0xf8, 0x42, 0x69, 0xa6, 0x92, 0xb5, 0xf2, 0x22,
	0xc5, 0xd0, 0x51, 0x30, 0x3b, 0x8c, 0xd8, 0x74, 0xcd, 0xc6, 0x3f, 0x4c, 0xf5, 0x5b, 0xa6, 0xd0,
	0x35, 0x3a, 0x01, 0x85, 0xaa, 0xe3, 0x30, 0xca, 0x79, 0xc8, 0xda, 0xae, 0x4d, 0xf1, 0xe7, 0x9c,
	0x12, 0x14, 0x48, 0x9a, 0x4a, 0x87, 0xcb, 0x2e, 0xe3, 0xa2, 0x15, 0xe1, 0xdb, 0xff, 0x69, 0x87,
	0x4d, 0x5d, 0xcb, 0x79, 0x36, 0x88, 0xea, 0x2c, 0xe9, 0x79, 0x7a, 0xaa, 0x42, 0x8b, 0x60, 0x5c,
	0x71, 0x9d, 0x7e, 0x2b, 0xc2, 0xdb, 0xfa, 0x7e, 0x46, 0x5f, 0x95, 0xe8, 0x30, 0xcc, 0x35, 0xc2,
	0xad, 0x56, 0x84, 0xdf, 0x6b, 0x3e, 0xe7, 0xc9, 0x4a, 0xce, 0xac, 0x36, 0x1a, 0x5f, 0x0b, 0xbd,
	0x91, 0x4f, 0xf1, 0xdb, 0x78, 0x66, 0xdd, 0x84, 0xa0, 0x25, 0x80, 0x36, 0xf5, 0xbc, 0xb8, 0xff,
	0x4e, 0xf7, 0x81, 0x4f, 0x10, 0x5a, 0x80, 0x6c, 0xcd, 0xb5, 0xf1, 0x17, 0x7d, 0xfe, 0x6c, 0xd7,
	0xb5, 0x25, 0xaa, 0xf2, 0x01, 0xfe, 0x1a, 0x23, 0xc2, 0x07, 0xa8, 0x04, 0xd0, 0x08, 0xb7, 0x2c,
	0x12, 0x38, 0xb4, 0x15, 0xe1, 0x27, 0x2f, 0x63, 0x1f, 0x6f, 0xc2, 0x50, 0x19, 0x72, 0xf2, 0xdc,
	0x89, 0xe4, 0x69, 0x2c, 0xc9, 0xf5, 0x77, 0x61, 0xec, 0xd2, 0x70, 0x7d, 0x57, 0xde, 0xfb, 0x59,
	0xca, 0x25, 0x66, 0x89, 0x4b, 0x22, 0x79, 0x9e, 0x76, 0x49, 0x34, 0x25, 0x80, 0x8b, 0x1e, 0x25,
	0xcc, 0x0d, 0x9c, 0x56, 0x84, 0x5f, 0x24, 0x2e, 0xbd, 0x09, 0x2b, 0x7f, 0xcf, 0xc2, 0x42, 0x53,
	0xc5, 0xb9, 0x4e, 0x04, 0xb1, 0xe8, 0x26, 0xa3, 0xbc, 0x8f, 0x30, 0x98, 0x4d, 0xee, 0xa8, 0xf0,
	0x2d, 0xeb, 0xb7, 0xf0, 0x75, 0x29, 0x87, 0xd4, 0xac, 0xb7, 0x05, 0xa3, 0xc4, 0x5f, 0xb3, 0xf1,
	0x8e, 0xbe, 0x37, 0xf8, 0x76, 0x82, 0x64, 0xac, 0xe4, 0x63, 0xe9, 0x3c, 0xab, 0xbc, 0xdc, 0x37,
	0x54, 0x60, 0x8a, 0xde, 0x14, 0xd6, 0x4e, 0xb5, 0x30, 0x1c, 0xa8, 0x9f, 0xf9, 0x69, 0x26, 0x4e,
	0x09, 0x3a, 0xf0, 0x33, 0xfa, 0xa7, 0x06, 0x2c, 0x73, 0x5d, 0xa3, 0x5c, 0xa4, 0x82, 0x63, 0x76,
	0x75, 0x9d, 0xb4, 0x52, 0x01, 0x52, 0xad, 0x2a, 0x1f, 0xa0, 0xca, 0x64, 0x9b, 0xe0, 0x07, 0x99,
	0x52, 0xb6, 0x92, 0x3b, 0x5b, 0x58, 0x91, 0x7b, 0x66, 0x25, 0xa6, 0x96, 0x19, 0x7f, 0xca, 0xe8,
	0x24, 0x18, 0x16, 0x11, 0x6e, 0xc8, 0xf1, 0xb7, 0xdc, 0x7e, 0x42, 0x83, 0xa9, 0x6e, 0x79, 0x27,
	0x33, 0xfd, 0xd0, 0x37, 0x46, 0x94, 0x8b, 0x7d, 0x1e, 0xba, 0xf9, 0xdb, 0x87, 0xde, 0x45, 0xe8,
	0x3c, 0x2c, 0xb6, 0x47, 0x5d, 0xde, 0x63, 0xee, 0x50, 0x2e, 0xad, 0xd8, 0x51, 0x59, 0xdd, 0xd1,
	0x63, 0x38, 0xa8, 0x8f, 0x96, 0x21, 0xdf, 0xe9, 0xb3, 0x50, 0x08, 0x8f, 0x2a, 0xfd, 0x76, 0x5e,
	0xe9, 0xa7, 0x20, 0x5a, 0x85, 0x43, 0x93, 0xda, 0xf5, 0xe9, 0x5a, 0x20, 0x28, 0xbb, 0x49, 0x3c,
	0xfc, 0x21, 0xaf, 0xe2, 0xb4, 0x6f, 0x13, 0x9d, 0x86, 0xff, 0xd3, 0x7c, 0x23, 0x70, 0x05, 0xfe,
	0xa8, 0xdd, 0xf7, 0x34, 0x64, 0x54, 0xd5, 0x66, 0x72, 0x88, 0xa0, 0x2a, 0x76, 0xf8, 0x5e, 0x46,
	0x47, 0x75, 0x1a, 0xab, 0x6d, 0xaa, 0x47, 0x47, 0x87, 0xa2, 0x8f, 0xef, 0x6a, 0x55, 0x9a, 0x95,
	0x47, 0xb0, 0xb8, 0x67, 0xba, 0x16, 0xbd, 0x4e, 0x7b, 0x7f, 0x35, 0xe3, 0x25, 0x00, 0x6d, 0xd2,
	0xa1, 0x91, 0xc0, 0xaf, 0x41, 0x0b, 0x76, 0x51, 0xf9, 0x0c, 0xcc, 0xd6, 0x08, 0xe3, 0x7f, 0x9e,
	0x97, 0xae, 0xa1, 0xfe, 0xb5, 0x56, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x25, 0xeb, 0x41,
	0xc5, 0x06, 0x00, 0x00,
}
