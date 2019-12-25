// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package api

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

// MsgType header for JSON deserialization
type MsgTypeHeader struct {
	MsgType              string   `protobuf:"bytes,35,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgTypeHeader) Reset()         { *m = MsgTypeHeader{} }
func (m *MsgTypeHeader) String() string { return proto.CompactTextString(m) }
func (*MsgTypeHeader) ProtoMessage()    {}
func (*MsgTypeHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *MsgTypeHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgTypeHeader.Unmarshal(m, b)
}
func (m *MsgTypeHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgTypeHeader.Marshal(b, m, deterministic)
}
func (m *MsgTypeHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTypeHeader.Merge(m, src)
}
func (m *MsgTypeHeader) XXX_Size() int {
	return xxx_messageInfo_MsgTypeHeader.Size(m)
}
func (m *MsgTypeHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTypeHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTypeHeader proto.InternalMessageInfo

func (m *MsgTypeHeader) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

type Heartbeat struct {
	MsgType              string   `protobuf:"bytes,35,opt,name=MsgType,proto3" json:"MsgType,omitempty"`
	TestReqId            string   `protobuf:"bytes,112,opt,name=TestReqId,proto3" json:"TestReqId,omitempty"`
	TransactTime         int64    `protobuf:"varint,60,opt,name=TransactTime,proto3" json:"TransactTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Heartbeat) Reset()         { *m = Heartbeat{} }
func (m *Heartbeat) String() string { return proto.CompactTextString(m) }
func (*Heartbeat) ProtoMessage()    {}
func (*Heartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *Heartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Heartbeat.Unmarshal(m, b)
}
func (m *Heartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Heartbeat.Marshal(b, m, deterministic)
}
func (m *Heartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Heartbeat.Merge(m, src)
}
func (m *Heartbeat) XXX_Size() int {
	return xxx_messageInfo_Heartbeat.Size(m)
}
func (m *Heartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_Heartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_Heartbeat proto.InternalMessageInfo

func (m *Heartbeat) GetMsgType() string {
	if m != nil {
		return m.MsgType
	}
	return ""
}

func (m *Heartbeat) GetTestReqId() string {
	if m != nil {
		return m.TestReqId
	}
	return ""
}

func (m *Heartbeat) GetTransactTime() int64 {
	if m != nil {
		return m.TransactTime
	}
	return 0
}

type Instrument struct {
	ID                    string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"ID,omitempty"`
	Type                  string   `protobuf:"bytes,2,opt,name=Type,json=type,proto3" json:"Type,omitempty"`
	Symbol                string   `protobuf:"bytes,3,opt,name=Symbol,json=symbol,proto3" json:"Symbol,omitempty"`
	BaseCurrencyName      string   `protobuf:"bytes,4,opt,name=BaseCurrencyName,json=baseCurrency,proto3" json:"BaseCurrencyName,omitempty"`
	QuoteCurrencyName     string   `protobuf:"bytes,5,opt,name=QuoteCurrencyName,json=quoteCurrency,proto3" json:"QuoteCurrencyName,omitempty"`
	SettlCurrencyName     string   `protobuf:"bytes,6,opt,name=SettlCurrencyName,json=settlCurrency,proto3" json:"SettlCurrencyName,omitempty"`
	TickSize              int32    `protobuf:"varint,7,opt,name=TickSize,json=tickSize,proto3" json:"TickSize,omitempty"`
	MinOrderQty           string   `protobuf:"bytes,8,opt,name=MinOrderQty,json=minOrderQuantity,proto3" json:"MinOrderQty,omitempty"`
	OrderQtyStep          string   `protobuf:"bytes,9,opt,name=OrderQtyStep,json=orderQtyStep,proto3" json:"OrderQtyStep,omitempty"`
	LimitOrderMaxDistance string   `protobuf:"bytes,10,opt,name=LimitOrderMaxDistance,json=limitOrderMaxDistance,proto3" json:"LimitOrderMaxDistance,omitempty"`
	PriceInputMask        string   `protobuf:"bytes,11,opt,name=PriceInputMask,json=priceInputMask,proto3" json:"PriceInputMask,omitempty"`
	Indexes               []string `protobuf:"bytes,12,rep,name=Indexes,json=indexes,proto3" json:"Indexes,omitempty"`
	Enabled               bool     `protobuf:"varint,13,opt,name=Enabled,json=enabled,proto3" json:"Enabled,omitempty"`
	// Fields from derivative
	LiquidationMaxDistance string               `protobuf:"bytes,14,opt,name=LiquidationMaxDistance,json=liquidationMaxDistance,proto3" json:"LiquidationMaxDistance,omitempty"`
	ContractValue          string               `protobuf:"bytes,15,opt,name=ContractValue,json=contractValue,proto3" json:"ContractValue,omitempty"`
	ContractCurrency       string               `protobuf:"bytes,16,opt,name=ContractCurrency,json=contractCurrency,proto3" json:"ContractCurrency,omitempty"`
	LotSize                string               `protobuf:"bytes,17,opt,name=LotSize,json=lotSize,proto3" json:"LotSize,omitempty"`
	TickValue              string               `protobuf:"bytes,18,opt,name=TickValue,json=tickValue,proto3" json:"TickValue,omitempty"`
	MaxOrderQty            string               `protobuf:"bytes,19,opt,name=MaxOrderQty,json=maxOrderQty,proto3" json:"MaxOrderQty,omitempty"`
	MaxPosVolume           string               `protobuf:"bytes,20,opt,name=MaxPosVolume,json=maxPosVolume,proto3" json:"MaxPosVolume,omitempty"`
	Mark                   string               `protobuf:"bytes,21,opt,name=Mark,json=mark,proto3" json:"Mark,omitempty"`
	FloatingPL             string               `protobuf:"bytes,22,opt,name=FloatingPL,json=floatingPL,proto3" json:"FloatingPL,omitempty"`
	AddUvmToFreeMargin     string               `protobuf:"bytes,23,opt,name=AddUvmToFreeMargin,json=addUvmToFreeMargin,proto3" json:"AddUvmToFreeMargin,omitempty"`
	MinLeverage            string               `protobuf:"bytes,24,opt,name=MinLeverage,json=minLeverage,proto3" json:"MinLeverage,omitempty"`
	MaxLeverage            string               `protobuf:"bytes,25,opt,name=MaxLeverage,json=maxLeverage,proto3" json:"MaxLeverage,omitempty"`
	Margin                 *Margin              `protobuf:"bytes,26,opt,name=Margin,json=margin,proto3" json:"Margin,omitempty"`
	Clearing               *DerivativeOperation `protobuf:"bytes,27,opt,name=Clearing,json=clearing,proto3" json:"Clearing,omitempty"`
	Interest               *DerivativeOperation `protobuf:"bytes,28,opt,name=Interest,json=interest,proto3" json:"Interest,omitempty"`
	Premium                *DerivativeOperation `protobuf:"bytes,29,opt,name=Premium,json=premium,proto3" json:"Premium,omitempty"`
	RiskAdjustment         *DerivativeOperation `protobuf:"bytes,30,opt,name=RiskAdjustment,json=riskAdjustment,proto3" json:"RiskAdjustment,omitempty"`
	PricePrecision         int32                `protobuf:"varint,31,opt,name=PricePrecision,json=pricePrecision,proto3" json:"PricePrecision,omitempty"`
	PriceRange             *PriceRange          `protobuf:"bytes,32,opt,name=PriceRange,json=priceRange,proto3" json:"PriceRange,omitempty"`
	PriceLimits            *PriceLimits         `protobuf:"bytes,33,opt,name=PriceLimits,json=priceLimits,proto3" json:"PriceLimits,omitempty"`
	Inverse                bool                 `protobuf:"varint,34,opt,name=Inverse,json=inverse,proto3" json:"Inverse,omitempty"`
	// Futures
	TradingStartDate string `protobuf:"bytes,35,opt,name=TradingStartDate,json=tradingStartDate,proto3" json:"TradingStartDate,omitempty"`
	ExpiryDate       string `protobuf:"bytes,36,opt,name=ExpiryDate,json=expiryDate,proto3" json:"ExpiryDate,omitempty"`
	// Fields from index
	Basis                int32    `protobuf:"varint,37,opt,name=Basis,json=basis,proto3" json:"Basis,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Instrument) Reset()         { *m = Instrument{} }
func (m *Instrument) String() string { return proto.CompactTextString(m) }
func (*Instrument) ProtoMessage()    {}
func (*Instrument) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *Instrument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instrument.Unmarshal(m, b)
}
func (m *Instrument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instrument.Marshal(b, m, deterministic)
}
func (m *Instrument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instrument.Merge(m, src)
}
func (m *Instrument) XXX_Size() int {
	return xxx_messageInfo_Instrument.Size(m)
}
func (m *Instrument) XXX_DiscardUnknown() {
	xxx_messageInfo_Instrument.DiscardUnknown(m)
}

var xxx_messageInfo_Instrument proto.InternalMessageInfo

func (m *Instrument) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Instrument) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Instrument) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Instrument) GetBaseCurrencyName() string {
	if m != nil {
		return m.BaseCurrencyName
	}
	return ""
}

func (m *Instrument) GetQuoteCurrencyName() string {
	if m != nil {
		return m.QuoteCurrencyName
	}
	return ""
}

func (m *Instrument) GetSettlCurrencyName() string {
	if m != nil {
		return m.SettlCurrencyName
	}
	return ""
}

func (m *Instrument) GetTickSize() int32 {
	if m != nil {
		return m.TickSize
	}
	return 0
}

func (m *Instrument) GetMinOrderQty() string {
	if m != nil {
		return m.MinOrderQty
	}
	return ""
}

func (m *Instrument) GetOrderQtyStep() string {
	if m != nil {
		return m.OrderQtyStep
	}
	return ""
}

func (m *Instrument) GetLimitOrderMaxDistance() string {
	if m != nil {
		return m.LimitOrderMaxDistance
	}
	return ""
}

func (m *Instrument) GetPriceInputMask() string {
	if m != nil {
		return m.PriceInputMask
	}
	return ""
}

func (m *Instrument) GetIndexes() []string {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func (m *Instrument) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Instrument) GetLiquidationMaxDistance() string {
	if m != nil {
		return m.LiquidationMaxDistance
	}
	return ""
}

func (m *Instrument) GetContractValue() string {
	if m != nil {
		return m.ContractValue
	}
	return ""
}

func (m *Instrument) GetContractCurrency() string {
	if m != nil {
		return m.ContractCurrency
	}
	return ""
}

func (m *Instrument) GetLotSize() string {
	if m != nil {
		return m.LotSize
	}
	return ""
}

func (m *Instrument) GetTickValue() string {
	if m != nil {
		return m.TickValue
	}
	return ""
}

func (m *Instrument) GetMaxOrderQty() string {
	if m != nil {
		return m.MaxOrderQty
	}
	return ""
}

func (m *Instrument) GetMaxPosVolume() string {
	if m != nil {
		return m.MaxPosVolume
	}
	return ""
}

func (m *Instrument) GetMark() string {
	if m != nil {
		return m.Mark
	}
	return ""
}

func (m *Instrument) GetFloatingPL() string {
	if m != nil {
		return m.FloatingPL
	}
	return ""
}

func (m *Instrument) GetAddUvmToFreeMargin() string {
	if m != nil {
		return m.AddUvmToFreeMargin
	}
	return ""
}

func (m *Instrument) GetMinLeverage() string {
	if m != nil {
		return m.MinLeverage
	}
	return ""
}

func (m *Instrument) GetMaxLeverage() string {
	if m != nil {
		return m.MaxLeverage
	}
	return ""
}

func (m *Instrument) GetMargin() *Margin {
	if m != nil {
		return m.Margin
	}
	return nil
}

func (m *Instrument) GetClearing() *DerivativeOperation {
	if m != nil {
		return m.Clearing
	}
	return nil
}

func (m *Instrument) GetInterest() *DerivativeOperation {
	if m != nil {
		return m.Interest
	}
	return nil
}

func (m *Instrument) GetPremium() *DerivativeOperation {
	if m != nil {
		return m.Premium
	}
	return nil
}

func (m *Instrument) GetRiskAdjustment() *DerivativeOperation {
	if m != nil {
		return m.RiskAdjustment
	}
	return nil
}

func (m *Instrument) GetPricePrecision() int32 {
	if m != nil {
		return m.PricePrecision
	}
	return 0
}

func (m *Instrument) GetPriceRange() *PriceRange {
	if m != nil {
		return m.PriceRange
	}
	return nil
}

func (m *Instrument) GetPriceLimits() *PriceLimits {
	if m != nil {
		return m.PriceLimits
	}
	return nil
}

func (m *Instrument) GetInverse() bool {
	if m != nil {
		return m.Inverse
	}
	return false
}

func (m *Instrument) GetTradingStartDate() string {
	if m != nil {
		return m.TradingStartDate
	}
	return ""
}

func (m *Instrument) GetExpiryDate() string {
	if m != nil {
		return m.ExpiryDate
	}
	return ""
}

func (m *Instrument) GetBasis() int32 {
	if m != nil {
		return m.Basis
	}
	return 0
}

type Margin struct {
	Netting              string            `protobuf:"bytes,1,opt,name=Netting,json=netting,proto3" json:"Netting,omitempty"`
	Rates                []*MarginRate     `protobuf:"bytes,2,rep,name=Rates,json=rates,proto3" json:"Rates,omitempty"`
	RateMultipliers      map[string]string `protobuf:"bytes,3,rep,name=RateMultipliers,json=rateMultipliers,proto3" json:"RateMultipliers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Margin) Reset()         { *m = Margin{} }
func (m *Margin) String() string { return proto.CompactTextString(m) }
func (*Margin) ProtoMessage()    {}
func (*Margin) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *Margin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Margin.Unmarshal(m, b)
}
func (m *Margin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Margin.Marshal(b, m, deterministic)
}
func (m *Margin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Margin.Merge(m, src)
}
func (m *Margin) XXX_Size() int {
	return xxx_messageInfo_Margin.Size(m)
}
func (m *Margin) XXX_DiscardUnknown() {
	xxx_messageInfo_Margin.DiscardUnknown(m)
}

var xxx_messageInfo_Margin proto.InternalMessageInfo

func (m *Margin) GetNetting() string {
	if m != nil {
		return m.Netting
	}
	return ""
}

func (m *Margin) GetRates() []*MarginRate {
	if m != nil {
		return m.Rates
	}
	return nil
}

func (m *Margin) GetRateMultipliers() map[string]string {
	if m != nil {
		return m.RateMultipliers
	}
	return nil
}

type MarginRate struct {
	MaxVolume            string   `protobuf:"bytes,1,opt,name=MaxVolume,json=maxVolume,proto3" json:"MaxVolume,omitempty"`
	InitialRate          string   `protobuf:"bytes,2,opt,name=InitialRate,json=initialRate,proto3" json:"InitialRate,omitempty"`
	MaintenanceRate      string   `protobuf:"bytes,3,opt,name=MaintenanceRate,json=maintenanceRate,proto3" json:"MaintenanceRate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MarginRate) Reset()         { *m = MarginRate{} }
func (m *MarginRate) String() string { return proto.CompactTextString(m) }
func (*MarginRate) ProtoMessage()    {}
func (*MarginRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{4}
}

func (m *MarginRate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarginRate.Unmarshal(m, b)
}
func (m *MarginRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarginRate.Marshal(b, m, deterministic)
}
func (m *MarginRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarginRate.Merge(m, src)
}
func (m *MarginRate) XXX_Size() int {
	return xxx_messageInfo_MarginRate.Size(m)
}
func (m *MarginRate) XXX_DiscardUnknown() {
	xxx_messageInfo_MarginRate.DiscardUnknown(m)
}

var xxx_messageInfo_MarginRate proto.InternalMessageInfo

func (m *MarginRate) GetMaxVolume() string {
	if m != nil {
		return m.MaxVolume
	}
	return ""
}

func (m *MarginRate) GetInitialRate() string {
	if m != nil {
		return m.InitialRate
	}
	return ""
}

func (m *MarginRate) GetMaintenanceRate() string {
	if m != nil {
		return m.MaintenanceRate
	}
	return ""
}

type DerivativeOperation struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=Enabled,json=maxVolenabledume,proto3" json:"Enabled,omitempty"`
	Index                string   `protobuf:"bytes,2,opt,name=Index,json=index,proto3" json:"Index,omitempty"`
	Schedule             int64    `protobuf:"varint,3,opt,name=Schedule,json=schedule,proto3" json:"Schedule,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DerivativeOperation) Reset()         { *m = DerivativeOperation{} }
func (m *DerivativeOperation) String() string { return proto.CompactTextString(m) }
func (*DerivativeOperation) ProtoMessage()    {}
func (*DerivativeOperation) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{5}
}

func (m *DerivativeOperation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DerivativeOperation.Unmarshal(m, b)
}
func (m *DerivativeOperation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DerivativeOperation.Marshal(b, m, deterministic)
}
func (m *DerivativeOperation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DerivativeOperation.Merge(m, src)
}
func (m *DerivativeOperation) XXX_Size() int {
	return xxx_messageInfo_DerivativeOperation.Size(m)
}
func (m *DerivativeOperation) XXX_DiscardUnknown() {
	xxx_messageInfo_DerivativeOperation.DiscardUnknown(m)
}

var xxx_messageInfo_DerivativeOperation proto.InternalMessageInfo

func (m *DerivativeOperation) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *DerivativeOperation) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *DerivativeOperation) GetSchedule() int64 {
	if m != nil {
		return m.Schedule
	}
	return 0
}

type PriceRange struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=Enabled,json=enabled,proto3" json:"Enabled,omitempty"`
	Distance             string   `protobuf:"bytes,2,opt,name=Distance,json=distance,proto3" json:"Distance,omitempty"`
	MovingBoundary       string   `protobuf:"bytes,3,opt,name=MovingBoundary,json=movingBoundary,proto3" json:"MovingBoundary,omitempty"`
	MovingTime           int64    `protobuf:"varint,4,opt,name=MovingTime,json=movingTime,proto3" json:"MovingTime,omitempty"`
	LowIndex             string   `protobuf:"bytes,5,opt,name=LowIndex,json=lowIndex,proto3" json:"LowIndex,omitempty"`
	HighIndex            string   `protobuf:"bytes,6,opt,name=HighIndex,json=highIndex,proto3" json:"HighIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PriceRange) Reset()         { *m = PriceRange{} }
func (m *PriceRange) String() string { return proto.CompactTextString(m) }
func (*PriceRange) ProtoMessage()    {}
func (*PriceRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{6}
}

func (m *PriceRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceRange.Unmarshal(m, b)
}
func (m *PriceRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceRange.Marshal(b, m, deterministic)
}
func (m *PriceRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceRange.Merge(m, src)
}
func (m *PriceRange) XXX_Size() int {
	return xxx_messageInfo_PriceRange.Size(m)
}
func (m *PriceRange) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceRange.DiscardUnknown(m)
}

var xxx_messageInfo_PriceRange proto.InternalMessageInfo

func (m *PriceRange) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *PriceRange) GetDistance() string {
	if m != nil {
		return m.Distance
	}
	return ""
}

func (m *PriceRange) GetMovingBoundary() string {
	if m != nil {
		return m.MovingBoundary
	}
	return ""
}

func (m *PriceRange) GetMovingTime() int64 {
	if m != nil {
		return m.MovingTime
	}
	return 0
}

func (m *PriceRange) GetLowIndex() string {
	if m != nil {
		return m.LowIndex
	}
	return ""
}

func (m *PriceRange) GetHighIndex() string {
	if m != nil {
		return m.HighIndex
	}
	return ""
}

type PriceLimits struct {
	Enabled              bool     `protobuf:"varint,1,opt,name=Enabled,json=enabled,proto3" json:"Enabled,omitempty"`
	Distance             string   `protobuf:"bytes,2,opt,name=Distance,json=distance,proto3" json:"Distance,omitempty"`
	LowIndex             string   `protobuf:"bytes,3,opt,name=LowIndex,json=lowIndex,proto3" json:"LowIndex,omitempty"`
	HighIndex            string   `protobuf:"bytes,4,opt,name=HighIndex,json=highIndex,proto3" json:"HighIndex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PriceLimits) Reset()         { *m = PriceLimits{} }
func (m *PriceLimits) String() string { return proto.CompactTextString(m) }
func (*PriceLimits) ProtoMessage()    {}
func (*PriceLimits) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{7}
}

func (m *PriceLimits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PriceLimits.Unmarshal(m, b)
}
func (m *PriceLimits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PriceLimits.Marshal(b, m, deterministic)
}
func (m *PriceLimits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PriceLimits.Merge(m, src)
}
func (m *PriceLimits) XXX_Size() int {
	return xxx_messageInfo_PriceLimits.Size(m)
}
func (m *PriceLimits) XXX_DiscardUnknown() {
	xxx_messageInfo_PriceLimits.DiscardUnknown(m)
}

var xxx_messageInfo_PriceLimits proto.InternalMessageInfo

func (m *PriceLimits) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *PriceLimits) GetDistance() string {
	if m != nil {
		return m.Distance
	}
	return ""
}

func (m *PriceLimits) GetLowIndex() string {
	if m != nil {
		return m.LowIndex
	}
	return ""
}

func (m *PriceLimits) GetHighIndex() string {
	if m != nil {
		return m.HighIndex
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgTypeHeader)(nil), "api.MsgTypeHeader")
	proto.RegisterType((*Heartbeat)(nil), "api.Heartbeat")
	proto.RegisterType((*Instrument)(nil), "api.Instrument")
	proto.RegisterType((*Margin)(nil), "api.Margin")
	proto.RegisterMapType((map[string]string)(nil), "api.Margin.RateMultipliersEntry")
	proto.RegisterType((*MarginRate)(nil), "api.MarginRate")
	proto.RegisterType((*DerivativeOperation)(nil), "api.DerivativeOperation")
	proto.RegisterType((*PriceRange)(nil), "api.PriceRange")
	proto.RegisterType((*PriceLimits)(nil), "api.PriceLimits")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 1067 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x6d, 0x6f, 0x1b, 0x45,
	0x10, 0x96, 0xe3, 0xf8, 0x6d, 0x9c, 0xd8, 0xee, 0x36, 0x0d, 0x4b, 0x28, 0xc5, 0x75, 0x5f, 0x64,
	0xf8, 0x10, 0xa4, 0x50, 0x21, 0x84, 0xf8, 0x40, 0xf3, 0x52, 0xc5, 0xc8, 0xd7, 0xa6, 0x67, 0xd3,
	0xef, 0xeb, 0xbb, 0xc1, 0xd9, 0xfa, 0x6e, 0xef, 0xb2, 0xb7, 0x67, 0x6c, 0x24, 0xfe, 0x1d, 0x7f,
	0x82, 0x1f, 0x83, 0x84, 0x76, 0xf7, 0xee, 0x7c, 0x0e, 0x81, 0x48, 0x7c, 0xf3, 0x3c, 0xf3, 0xcc,
	0xcd, 0xec, 0xce, 0x33, 0xb3, 0x86, 0x3d, 0x2f, 0x0a, 0xc3, 0x48, 0x1c, 0xc7, 0x32, 0x52, 0x11,
	0xa9, 0xb2, 0x98, 0x0f, 0xbe, 0x84, 0x7d, 0x27, 0x99, 0x4f, 0xd7, 0x31, 0x5e, 0x22, 0xf3, 0x51,
	0x12, 0x0a, 0x8d, 0x0c, 0xa0, 0xcf, 0xfa, 0x95, 0x61, 0xcb, 0xcd, 0xcd, 0xc1, 0x1c, 0x5a, 0x97,
	0xc8, 0xa4, 0x9a, 0x21, 0x53, 0xff, 0x4e, 0x23, 0x8f, 0xa1, 0x35, 0xc5, 0x44, 0xb9, 0x78, 0x33,
	0xf2, 0x69, 0x6c, 0x7c, 0x1b, 0x80, 0x0c, 0x60, 0x6f, 0x2a, 0x99, 0x48, 0x98, 0xa7, 0xa6, 0x3c,
	0x44, 0xfa, 0x43, 0xbf, 0x32, 0xac, 0xba, 0x5b, 0xd8, 0xe0, 0x2f, 0x00, 0x18, 0x89, 0x44, 0xc9,
	0x34, 0x44, 0xa1, 0x48, 0x07, 0x76, 0x46, 0xe7, 0xb4, 0x62, 0xbe, 0xb4, 0xc3, 0x7d, 0x42, 0x60,
	0xd7, 0xe4, 0xdd, 0x31, 0xc8, 0xae, 0xd2, 0x49, 0x0f, 0xa1, 0x3e, 0x59, 0x87, 0xb3, 0x28, 0xa0,
	0x55, 0x83, 0xd6, 0x13, 0x63, 0x91, 0x97, 0xd0, 0x3b, 0x65, 0x09, 0x9e, 0xa5, 0x52, 0xa2, 0xf0,
	0xd6, 0x6f, 0x59, 0x88, 0x74, 0xd7, 0x30, 0xf6, 0x66, 0x25, 0x9c, 0x0c, 0xe1, 0xc1, 0xfb, 0x34,
	0x52, 0xdb, 0xc4, 0x9a, 0x21, 0xee, 0xdf, 0x94, 0x1d, 0x9a, 0x39, 0x41, 0xa5, 0x82, 0x2d, 0x66,
	0xdd, 0x32, 0x93, 0xb2, 0x83, 0x1c, 0x41, 0x73, 0xca, 0xbd, 0xc5, 0x84, 0xff, 0x86, 0xb4, 0xd1,
	0xaf, 0x0c, 0x6b, 0x6e, 0x53, 0x65, 0x36, 0x79, 0x01, 0x6d, 0x87, 0x8b, 0x77, 0xd2, 0x47, 0xf9,
	0x5e, 0xad, 0x69, 0xd3, 0xc4, 0xf7, 0xc2, 0x1c, 0x4a, 0x99, 0x50, 0x5c, 0xad, 0xf5, 0x6d, 0xe5,
	0x9c, 0x89, 0xc2, 0x98, 0xb6, 0x6c, 0xe9, 0x51, 0x09, 0x23, 0xaf, 0xe0, 0xd1, 0x98, 0x87, 0x5c,
	0x19, 0xa2, 0xc3, 0x56, 0xe7, 0x3c, 0x51, 0x4c, 0x78, 0x48, 0xc1, 0x90, 0x1f, 0x05, 0x77, 0x39,
	0xc9, 0x4b, 0xe8, 0x5c, 0x49, 0xee, 0xe1, 0x48, 0xc4, 0xa9, 0x72, 0x58, 0xb2, 0xa0, 0x6d, 0x43,
	0xef, 0xc4, 0x5b, 0xa8, 0xee, 0xf3, 0x48, 0xf8, 0xb8, 0xc2, 0x84, 0xee, 0xf5, 0xab, 0xba, 0xcf,
	0xdc, 0x9a, 0xda, 0x73, 0x21, 0xd8, 0x2c, 0x40, 0x9f, 0xee, 0xf7, 0x2b, 0xc3, 0xa6, 0xdb, 0x40,
	0x6b, 0x92, 0x6f, 0xe1, 0x70, 0xcc, 0x6f, 0x52, 0xee, 0x33, 0xc5, 0x23, 0x51, 0x2e, 0xa9, 0x63,
	0x72, 0x1c, 0x06, 0x77, 0x7a, 0xc9, 0x73, 0xd8, 0x3f, 0x8b, 0x84, 0x92, 0xcc, 0x53, 0x1f, 0x58,
	0x90, 0x22, 0xed, 0xda, 0x6b, 0xf5, 0xca, 0x20, 0xf9, 0x0a, 0x7a, 0x39, 0x2b, 0xbf, 0x6a, 0xda,
	0xb3, 0xf7, 0xe7, 0xdd, 0xc2, 0x75, 0x8d, 0xe3, 0x48, 0x99, 0x0e, 0x3c, 0xb0, 0x2a, 0x0d, 0xac,
	0x69, 0x54, 0xca, 0xbd, 0x85, 0xcd, 0x43, 0xac, 0x4a, 0x55, 0x0e, 0x90, 0x3e, 0xb4, 0x1d, 0xb6,
	0x2a, 0xda, 0xf3, 0xd0, 0xf8, 0xdb, 0xe1, 0x06, 0xd2, 0x9d, 0x71, 0xd8, 0xea, 0x2a, 0x4a, 0x3e,
	0x44, 0x41, 0x1a, 0x22, 0x3d, 0xb0, 0x9d, 0x09, 0x4b, 0x98, 0x16, 0xaa, 0xc3, 0xe4, 0x82, 0x3e,
	0xb2, 0x42, 0x0d, 0x99, 0x5c, 0x90, 0x27, 0x00, 0x6f, 0x82, 0x88, 0x29, 0x2e, 0xe6, 0x57, 0x63,
	0x7a, 0x68, 0x3c, 0xf0, 0x4b, 0x81, 0x90, 0x63, 0x20, 0xaf, 0x7d, 0xff, 0xe7, 0x65, 0x38, 0x8d,
	0xde, 0x48, 0x44, 0x87, 0xc9, 0x39, 0x17, 0xf4, 0x13, 0xc3, 0x23, 0xec, 0x1f, 0x1e, 0x53, 0x29,
	0x17, 0x63, 0x5c, 0xa2, 0x64, 0x73, 0xa4, 0x34, 0xab, 0x74, 0x03, 0x65, 0x67, 0x29, 0x18, 0x9f,
	0x16, 0x67, 0x29, 0x18, 0xcf, 0xa0, 0x9e, 0xe5, 0x39, 0xea, 0x57, 0x86, 0xed, 0x93, 0xf6, 0x31,
	0x8b, 0xf9, 0xb1, 0x85, 0xdc, 0x7a, 0x68, 0x13, 0xbd, 0x82, 0xe6, 0x59, 0x80, 0x4c, 0x72, 0x31,
	0xa7, 0x9f, 0x19, 0x1a, 0x35, 0xb4, 0x73, 0x94, 0x7c, 0xc9, 0x14, 0x5f, 0xe2, 0xbb, 0x18, 0xa5,
	0xe9, 0xa9, 0xdb, 0xf4, 0x32, 0xa6, 0x8e, 0x1a, 0x09, 0x85, 0x12, 0x13, 0x45, 0x1f, 0xdf, 0x17,
	0xc5, 0x33, 0x26, 0x39, 0x81, 0xc6, 0x95, 0xc4, 0x90, 0xa7, 0x21, 0xfd, 0xfc, 0x9e, 0xa0, 0x46,
	0x6c, 0x89, 0xe4, 0x47, 0xe8, 0xb8, 0x3c, 0x59, 0xbc, 0xf6, 0x3f, 0xa6, 0x89, 0xd2, 0x7b, 0x83,
	0x3e, 0xb9, 0x27, 0xb4, 0x23, 0xb7, 0xf8, 0xc5, 0x48, 0x5c, 0x49, 0xf4, 0x78, 0xc2, 0x23, 0x41,
	0xbf, 0x30, 0x53, 0x6b, 0x47, 0xa2, 0x40, 0xc9, 0xd7, 0x00, 0x86, 0xe7, 0x32, 0x31, 0x47, 0xda,
	0x37, 0x59, 0xba, 0x26, 0xcb, 0x06, 0x76, 0x21, 0x2e, 0x7e, 0x93, 0x13, 0x68, 0x1b, 0x8f, 0x19,
	0xd3, 0x84, 0x3e, 0x35, 0x11, 0xbd, 0x4d, 0x84, 0xc5, 0xdd, 0x76, 0xbc, 0x31, 0xec, 0xdc, 0x2d,
	0x51, 0x26, 0x48, 0x07, 0x76, 0xba, 0xb8, 0x35, 0xb5, 0xfe, 0xa7, 0x92, 0xf9, 0x5c, 0xcc, 0x27,
	0x8a, 0x49, 0x75, 0xce, 0x54, 0xbe, 0x82, 0x7b, 0xea, 0x16, 0xae, 0xd5, 0x76, 0xb1, 0x8a, 0xb9,
	0x5c, 0x1b, 0xd6, 0x73, 0xab, 0x36, 0x2c, 0x10, 0x72, 0x00, 0xb5, 0x53, 0x96, 0xf0, 0x84, 0xbe,
	0x30, 0x27, 0xad, 0xcd, 0xb4, 0x31, 0xf8, 0xb3, 0x92, 0x0b, 0x42, 0x97, 0xf1, 0x16, 0x95, 0xd6,
	0x66, 0xb6, 0x80, 0x1b, 0xc2, 0x9a, 0xe4, 0x05, 0xd4, 0x5c, 0xa6, 0x30, 0xa1, 0x3b, 0xfd, 0x6a,
	0x71, 0x01, 0x99, 0x66, 0x98, 0x42, 0xb7, 0x26, 0xb5, 0x97, 0xfc, 0x04, 0x5d, 0x6d, 0x3a, 0x69,
	0xa0, 0x78, 0x1c, 0x70, 0x94, 0x09, 0xad, 0x9a, 0x80, 0x7e, 0x29, 0xe0, 0xf8, 0x16, 0xe5, 0x42,
	0x28, 0xb9, 0x76, 0xbb, 0x72, 0x1b, 0x3d, 0x3a, 0x85, 0x83, 0xbb, 0x88, 0xa4, 0x07, 0xd5, 0x05,
	0xae, 0xb3, 0x02, 0xf5, 0x4f, 0x7d, 0xae, 0xa5, 0x99, 0x6c, 0xfb, 0x46, 0x58, 0xe3, 0xfb, 0x9d,
	0xef, 0x2a, 0x83, 0x25, 0xc0, 0xa6, 0x48, 0xbd, 0x05, 0x1c, 0xb6, 0xca, 0x46, 0xd8, 0xc6, 0xb7,
	0xc2, 0x1c, 0xd0, 0x93, 0x33, 0x12, 0x5c, 0x71, 0x16, 0x68, 0x72, 0xf6, 0xad, 0x36, 0xdf, 0x40,
	0x64, 0x08, 0x5d, 0x87, 0x69, 0xd9, 0x0a, 0xbd, 0xc0, 0x0c, 0xcb, 0xbe, 0x3f, 0xdd, 0x70, 0x1b,
	0x1e, 0x7c, 0x84, 0x87, 0x77, 0x68, 0x90, 0x3c, 0xdd, 0x2c, 0xd1, 0x8a, 0x69, 0x73, 0xcf, 0xa6,
	0xcf, 0x56, 0xa9, 0xae, 0xe2, 0x00, 0x6a, 0x66, 0x03, 0xe7, 0x67, 0x31, 0xfb, 0x57, 0x3f, 0x2e,
	0x13, 0xef, 0x1a, 0xfd, 0x34, 0xb0, 0x29, 0xab, 0x6e, 0x33, 0xc9, 0xec, 0xc1, 0x1f, 0x95, 0xb2,
	0x42, 0xcb, 0x8b, 0xba, 0xb2, 0xbd, 0xa8, 0x8f, 0xa0, 0x59, 0xac, 0x66, 0xfb, 0xf5, 0xa6, 0x5f,
	0x7a, 0x20, 0x9c, 0x68, 0xc9, 0xc5, 0xfc, 0x34, 0x4a, 0x85, 0xcf, 0xe4, 0x3a, 0x3b, 0x59, 0x27,
	0xdc, 0x42, 0xb5, 0xc4, 0x2c, 0xcf, 0x3c, 0xe7, 0xbb, 0xa6, 0x14, 0x08, 0x0b, 0x44, 0xe7, 0x18,
	0x47, 0xbf, 0xda, 0x13, 0xd8, 0x07, 0xb5, 0x19, 0x64, 0xb6, 0xbe, 0xfe, 0x4b, 0x3e, 0xbf, 0xb6,
	0x4e, 0xfb, 0x86, 0xb6, 0xae, 0x73, 0x60, 0xf0, 0xfb, 0xd6, 0xd8, 0xfc, 0xcf, 0x63, 0x94, 0xd3,
	0x57, 0xff, 0x2b, 0xfd, 0xee, 0xad, 0xf4, 0xb3, 0xba, 0xf9, 0x97, 0xf4, 0xcd, 0xdf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0xf7, 0xfa, 0x0d, 0x35, 0x09, 0x00, 0x00,
}
