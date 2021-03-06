// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

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
	MsgType              string   `protobuf:"bytes,35,opt,name=MsgType,json=msgType,proto3" json:"msgType,omitempty"`
	TestReqId            string   `protobuf:"bytes,112,opt,name=TestReqId,json=testReqId,proto3" json:"testReqId,omitempty"`
	TransactTime         int64    `protobuf:"varint,60,opt,name=TransactTime,json=transactTime,proto3" json:"transactTime,omitempty"`
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
	ID                    string   `protobuf:"bytes,1,opt,name=ID,json=id,proto3" json:"id,omitempty"`
	Type                  string   `protobuf:"bytes,2,opt,name=Type,json=type,proto3" json:"type,omitempty"`
	Symbol                string   `protobuf:"bytes,3,opt,name=Symbol,json=symbol,proto3" json:"symbol,omitempty"`
	BaseCurrencyName      string   `protobuf:"bytes,4,opt,name=BaseCurrencyName,json=baseCurrency,proto3" json:"baseCurrency,omitempty"`
	QuoteCurrencyName     string   `protobuf:"bytes,5,opt,name=QuoteCurrencyName,json=quoteCurrency,proto3" json:"quoteCurrency,omitempty"`
	SettlCurrencyName     string   `protobuf:"bytes,6,opt,name=SettlCurrencyName,json=settlCurrency,proto3" json:"settlCurrency,omitempty"`
	TickSize              int32    `protobuf:"varint,7,opt,name=TickSize,json=tickSize,proto3" json:"tickSize,omitempty"`
	MinOrderQty           string   `protobuf:"bytes,8,opt,name=MinOrderQty,json=minOrderQuantity,proto3" json:"minOrderQuantity,omitempty"`
	OrderQtyStep          string   `protobuf:"bytes,9,opt,name=OrderQtyStep,json=orderQtyStep,proto3" json:"orderQtyStep,omitempty"`
	LimitOrderMaxDistance string   `protobuf:"bytes,10,opt,name=LimitOrderMaxDistance,json=limitOrderMaxDistance,proto3" json:"limitOrderMaxDistance,omitempty"`
	PriceInputMask        string   `protobuf:"bytes,11,opt,name=PriceInputMask,json=priceInputMask,proto3" json:"priceInputMask,omitempty"`
	Indexes               []string `protobuf:"bytes,12,rep,name=Indexes,json=indexes,proto3" json:"indexes,omitempty"`
	Enabled               bool     `protobuf:"varint,13,opt,name=Enabled,json=enabled,proto3" json:"enabled,omitempty"`
	// Fields from derivative
	LiquidationMaxDistance string               `protobuf:"bytes,14,opt,name=LiquidationMaxDistance,json=liquidationMaxDistance,proto3" json:"liquidationMaxDistance,omitempty"`
	ContractValue          string               `protobuf:"bytes,15,opt,name=ContractValue,json=contractValue,proto3" json:"contractValue,omitempty"`
	ContractCurrency       string               `protobuf:"bytes,16,opt,name=ContractCurrency,json=contractCurrency,proto3" json:"contractCurrency,omitempty"`
	LotSize                string               `protobuf:"bytes,17,opt,name=LotSize,json=lotSize,proto3" json:"lotSize,omitempty"`
	TickValue              string               `protobuf:"bytes,18,opt,name=TickValue,json=tickValue,proto3" json:"tickValue,omitempty"`
	MaxOrderQty            string               `protobuf:"bytes,19,opt,name=MaxOrderQty,json=maxOrderQty,proto3" json:"maxOrderQty,omitempty"`
	MaxPosVolume           string               `protobuf:"bytes,20,opt,name=MaxPosVolume,json=maxPosVolume,proto3" json:"maxPosVolume,omitempty"`
	Mark                   string               `protobuf:"bytes,21,opt,name=Mark,json=mark,proto3" json:"mark,omitempty"`
	FloatingPL             string               `protobuf:"bytes,22,opt,name=FloatingPL,json=floatingPL,proto3" json:"floatingPL,omitempty"`
	AddUvmToFreeMargin     string               `protobuf:"bytes,23,opt,name=AddUvmToFreeMargin,json=addUvmToFreeMargin,proto3" json:"addUvmToFreeMargin,omitempty"`
	MinLeverage            string               `protobuf:"bytes,24,opt,name=MinLeverage,json=minLeverage,proto3" json:"minLeverage,omitempty"`
	MaxLeverage            string               `protobuf:"bytes,25,opt,name=MaxLeverage,json=maxLeverage,proto3" json:"maxLeverage,omitempty"`
	Margin                 *Margin              `protobuf:"bytes,26,opt,name=Margin,json=margin,proto3" json:"margin,omitempty"`
	Clearing               *DerivativeOperation `protobuf:"bytes,27,opt,name=Clearing,json=clearing,proto3" json:"clearing,omitempty"`
	Interest               *DerivativeOperation `protobuf:"bytes,28,opt,name=Interest,json=interest,proto3" json:"interest,omitempty"`
	Premium                *DerivativeOperation `protobuf:"bytes,29,opt,name=Premium,json=premium,proto3" json:"premium,omitempty"`
	RiskAdjustment         *DerivativeOperation `protobuf:"bytes,30,opt,name=RiskAdjustment,json=riskAdjustment,proto3" json:"riskAdjustment,omitempty"`
	PricePrecision         int32                `protobuf:"varint,31,opt,name=PricePrecision,json=pricePrecision,proto3" json:"pricePrecision,omitempty"`
	PriceRange             *PriceRange          `protobuf:"bytes,32,opt,name=PriceRange,json=priceRange,proto3" json:"priceRange,omitempty"`
	PriceLimits            *PriceLimits         `protobuf:"bytes,33,opt,name=PriceLimits,json=priceLimits,proto3" json:"priceLimits,omitempty"`
	Inverse                bool                 `protobuf:"varint,34,opt,name=Inverse,json=inverse,proto3" json:"inverse,omitempty"`
	// Futures
	TradingStartDate string `protobuf:"bytes,35,opt,name=TradingStartDate,json=tradingStartDate,proto3" json:"tradingStartDate,omitempty"`
	ExpiryDate       string `protobuf:"bytes,36,opt,name=ExpiryDate,json=expiryDate,proto3" json:"expiryDate,omitempty"`
	// Fields from index
	Basis                int32    `protobuf:"varint,37,opt,name=Basis,json=basis,proto3" json:"basis,omitempty"`
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
	Netting              string            `protobuf:"bytes,1,opt,name=Netting,json=netting,proto3" json:"netting,omitempty"`
	Rates                []*MarginRate     `protobuf:"bytes,2,rep,name=Rates,json=rates,proto3" json:"rates,omitempty"`
	RateMultipliers      map[string]string `protobuf:"bytes,3,rep,name=RateMultipliers,json=rateMultipliers,proto3" json:"rateMultipliers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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
	MaxVolume            string   `protobuf:"bytes,1,opt,name=MaxVolume,json=maxVolume,proto3" json:"maxVolume,omitempty"`
	InitialRate          string   `protobuf:"bytes,2,opt,name=InitialRate,json=initialRate,proto3" json:"initialRate,omitempty"`
	MaintenanceRate      string   `protobuf:"bytes,3,opt,name=MaintenanceRate,json=maintenanceRate,proto3" json:"maintenanceRate,omitempty"`
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
	Enabled              bool     `protobuf:"varint,1,opt,name=Enabled,json=maxVolenabledume,proto3" json:"maxVolenabledume,omitempty"`
	Index                string   `protobuf:"bytes,2,opt,name=Index,json=index,proto3" json:"index,omitempty"`
	Schedule             int64    `protobuf:"varint,3,opt,name=Schedule,json=schedule,proto3" json:"schedule,omitempty"`
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
	Enabled              bool     `protobuf:"varint,1,opt,name=Enabled,json=enabled,proto3" json:"enabled,omitempty"`
	Distance             string   `protobuf:"bytes,2,opt,name=Distance,json=distance,proto3" json:"distance,omitempty"`
	MovingBoundary       string   `protobuf:"bytes,3,opt,name=MovingBoundary,json=movingBoundary,proto3" json:"movingBoundary,omitempty"`
	MovingTime           int64    `protobuf:"varint,4,opt,name=MovingTime,json=movingTime,proto3" json:"movingTime,omitempty"`
	LowIndex             string   `protobuf:"bytes,5,opt,name=LowIndex,json=lowIndex,proto3" json:"lowIndex,omitempty"`
	HighIndex            string   `protobuf:"bytes,6,opt,name=HighIndex,json=highIndex,proto3" json:"highIndex,omitempty"`
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
	Enabled              bool     `protobuf:"varint,1,opt,name=Enabled,json=enabled,proto3" json:"enabled,omitempty"`
	Distance             string   `protobuf:"bytes,2,opt,name=Distance,json=distance,proto3" json:"distance,omitempty"`
	LowIndex             string   `protobuf:"bytes,3,opt,name=LowIndex,json=lowIndex,proto3" json:"lowIndex,omitempty"`
	HighIndex            string   `protobuf:"bytes,4,opt,name=HighIndex,json=highIndex,proto3" json:"highIndex,omitempty"`
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
	proto.RegisterType((*MsgTypeHeader)(nil), "xmsg.MsgTypeHeader")
	proto.RegisterType((*Heartbeat)(nil), "xmsg.Heartbeat")
	proto.RegisterType((*Instrument)(nil), "xmsg.Instrument")
	proto.RegisterType((*Margin)(nil), "xmsg.Margin")
	proto.RegisterMapType((map[string]string)(nil), "xmsg.Margin.RateMultipliersEntry")
	proto.RegisterType((*MarginRate)(nil), "xmsg.MarginRate")
	proto.RegisterType((*DerivativeOperation)(nil), "xmsg.DerivativeOperation")
	proto.RegisterType((*PriceRange)(nil), "xmsg.PriceRange")
	proto.RegisterType((*PriceLimits)(nil), "xmsg.PriceLimits")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 1072 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x96, 0xe3, 0xff, 0x63, 0xc7, 0x71, 0xa7, 0x69, 0x98, 0x86, 0x52, 0x5c, 0x93, 0x46, 0x86,
	0x8b, 0x08, 0x11, 0x40, 0x08, 0x71, 0x93, 0xbf, 0x2a, 0x16, 0xde, 0x36, 0x5d, 0x9b, 0xde, 0x8f,
	0x77, 0x0f, 0xce, 0xd4, 0xbb, 0xb3, 0x9b, 0xd9, 0x59, 0x63, 0x23, 0x71, 0xc7, 0xa3, 0xf1, 0x16,
	0xbc, 0x0c, 0x9a, 0x99, 0x5d, 0x7b, 0x1d, 0x22, 0x22, 0x71, 0xe7, 0xef, 0x3b, 0xdf, 0xd9, 0x73,
	0x66, 0xce, 0xcf, 0x18, 0xda, 0x5e, 0x14, 0x86, 0x91, 0x38, 0x89, 0x65, 0xa4, 0x22, 0x52, 0x59,
	0x86, 0xc9, 0xac, 0xff, 0x25, 0xec, 0x3a, 0xc9, 0x6c, 0xb2, 0x8a, 0xf1, 0x1a, 0x99, 0x8f, 0x92,
	0x50, 0xa8, 0x67, 0x04, 0xfd, 0xa2, 0x57, 0x1a, 0x34, 0xdd, 0x1c, 0xf6, 0x67, 0xd0, 0xbc, 0x46,
	0x26, 0xd5, 0x14, 0x99, 0x7a, 0x40, 0x16, 0x5a, 0x48, 0x5e, 0x40, 0x73, 0x82, 0x89, 0x72, 0xf1,
	0x6e, 0xe8, 0xd3, 0xd8, 0xd8, 0x9a, 0x2a, 0x27, 0x48, 0x1f, 0xda, 0x13, 0xc9, 0x44, 0xc2, 0x3c,
	0x35, 0xe1, 0x21, 0xd2, 0x9f, 0x7a, 0xa5, 0x41, 0xd9, 0x6d, 0xab, 0x02, 0xd7, 0xff, 0xb3, 0x05,
	0x30, 0x14, 0x89, 0x92, 0x69, 0x88, 0x42, 0x91, 0x0e, 0xec, 0x0c, 0x2f, 0x69, 0xc9, 0x7c, 0x69,
	0x87, 0xfb, 0x84, 0x40, 0xc5, 0xc4, 0xdd, 0x31, 0x4c, 0x45, 0xe9, 0xa0, 0x07, 0x50, 0x1b, 0xaf,
	0xc2, 0x69, 0x14, 0xd0, 0xb2, 0x61, 0x6b, 0x89, 0x41, 0xe4, 0x18, 0xba, 0xe7, 0x2c, 0xc1, 0x8b,
	0x54, 0x4a, 0x14, 0xde, 0xea, 0x2d, 0x0b, 0x91, 0x56, 0x8c, 0xa2, 0x3d, 0x2d, 0xf0, 0x64, 0x00,
	0x4f, 0xde, 0xa7, 0x91, 0xda, 0x16, 0x56, 0x8d, 0x70, 0xf7, 0xae, 0x68, 0xd0, 0xca, 0x31, 0x2a,
	0x15, 0x6c, 0x29, 0x6b, 0x56, 0x99, 0x14, 0x0d, 0xe4, 0x10, 0x1a, 0x13, 0xee, 0xcd, 0xc7, 0xfc,
	0x77, 0xa4, 0xf5, 0x5e, 0x69, 0x50, 0x75, 0x1b, 0x2a, 0xc3, 0xe4, 0x35, 0xb4, 0x1c, 0x2e, 0xde,
	0x49, 0x1f, 0xe5, 0x7b, 0xb5, 0xa2, 0x0d, 0xe3, 0xdf, 0x0d, 0x73, 0x2a, 0x65, 0x42, 0x71, 0xb5,
	0xd2, 0xb7, 0x95, 0x6b, 0xc6, 0x0a, 0x63, 0xda, 0xb4, 0xa9, 0x47, 0x05, 0x8e, 0x7c, 0x0b, 0xcf,
	0x46, 0x3c, 0xe4, 0xca, 0x08, 0x1d, 0xb6, 0xbc, 0xe4, 0x89, 0x62, 0xc2, 0x43, 0x0a, 0x46, 0xfc,
	0x2c, 0x78, 0xc8, 0x48, 0x8e, 0xa1, 0x73, 0x23, 0xb9, 0x87, 0x43, 0x11, 0xa7, 0xca, 0x61, 0xc9,
	0x9c, 0xb6, 0x8c, 0xbc, 0x13, 0x6f, 0xb1, 0xba, 0xce, 0x43, 0xe1, 0xe3, 0x12, 0x13, 0xda, 0xee,
	0x95, 0x75, 0x9d, 0xb9, 0x85, 0xda, 0x72, 0x25, 0xd8, 0x34, 0x40, 0x9f, 0xee, 0xf6, 0x4a, 0x83,
	0x86, 0x5b, 0x47, 0x0b, 0xc9, 0xf7, 0x70, 0x30, 0xe2, 0x77, 0x29, 0xf7, 0x99, 0xe2, 0x91, 0x28,
	0xa6, 0xd4, 0x31, 0x31, 0x0e, 0x82, 0x07, 0xad, 0xe4, 0x08, 0x76, 0x2f, 0x22, 0xa1, 0x24, 0xf3,
	0xd4, 0x07, 0x16, 0xa4, 0x48, 0xf7, 0xec, 0xb5, 0x7a, 0x45, 0x92, 0x7c, 0x05, 0xdd, 0x5c, 0x95,
	0x5f, 0x35, 0xed, 0xda, 0xfb, 0xf3, 0xee, 0xf1, 0x3a, 0xc7, 0x51, 0xa4, 0x4c, 0x05, 0x9e, 0xd8,
	0x2e, 0x0d, 0x2c, 0x34, 0x5d, 0xca, 0xbd, 0xb9, 0x8d, 0x43, 0xb2, 0x2e, 0xcd, 0x09, 0xd2, 0x83,
	0x96, 0xc3, 0x96, 0xeb, 0xf2, 0x3c, 0x35, 0xf6, 0x56, 0xb8, 0xa1, 0x74, 0x65, 0x1c, 0xb6, 0xbc,
	0x89, 0x92, 0x0f, 0x51, 0x90, 0x86, 0x48, 0xf7, 0x6d, 0x65, 0xc2, 0x02, 0xa7, 0x1b, 0xd5, 0x61,
	0x72, 0x4e, 0x9f, 0xd9, 0x46, 0x0d, 0x99, 0x9c, 0x93, 0x97, 0x00, 0x6f, 0x82, 0x88, 0x29, 0x2e,
	0x66, 0x37, 0x23, 0x7a, 0x60, 0x2c, 0xf0, 0xeb, 0x9a, 0x21, 0x27, 0x40, 0xce, 0x7c, 0xff, 0x97,
	0x45, 0x38, 0x89, 0xde, 0x48, 0x44, 0x87, 0xc9, 0x19, 0x17, 0xf4, 0x13, 0xa3, 0x23, 0xec, 0x5f,
	0x16, 0x93, 0x29, 0x17, 0x23, 0x5c, 0xa0, 0x64, 0x33, 0xa4, 0x34, 0xcb, 0x74, 0x43, 0x65, 0x67,
	0x59, 0x2b, 0x9e, 0xaf, 0xcf, 0xb2, 0x56, 0x1c, 0x41, 0x2d, 0x8b, 0x73, 0xd8, 0x2b, 0x0d, 0x5a,
	0xdf, 0xb4, 0x4f, 0xf4, 0x6a, 0x38, 0xb1, 0x9c, 0x5b, 0x0b, 0x6d, 0xa4, 0xef, 0xa0, 0x71, 0x11,
	0x20, 0x93, 0x5c, 0xcc, 0xe8, 0xa7, 0x46, 0xf7, 0xdc, 0xea, 0x2e, 0x51, 0xf2, 0x05, 0x53, 0x7c,
	0x81, 0xef, 0x62, 0x94, 0xa6, 0xaa, 0x6e, 0xc3, 0xcb, 0xa4, 0xda, 0x6d, 0x28, 0x14, 0x4a, 0x4c,
	0x14, 0x7d, 0xf1, 0xa8, 0x1b, 0xcf, 0xa4, 0xe4, 0x14, 0xea, 0x37, 0x12, 0x43, 0x9e, 0x86, 0xf4,
	0xb3, 0xc7, 0xbc, 0xea, 0xb1, 0x55, 0x92, 0x33, 0xe8, 0xb8, 0x3c, 0x99, 0x9f, 0xf9, 0x1f, 0xd3,
	0x44, 0xe9, 0xdd, 0x41, 0x5f, 0x3e, 0xe6, 0xdb, 0x91, 0x5b, 0x0e, 0xeb, 0xb9, 0xb8, 0x91, 0xe8,
	0xf1, 0x84, 0x47, 0x82, 0x7e, 0x6e, 0x46, 0xd7, 0xce, 0xc5, 0x9a, 0x25, 0x5f, 0x03, 0x18, 0x9d,
	0xcb, 0xc4, 0x0c, 0x69, 0xcf, 0x84, 0xe9, 0xda, 0x30, 0x1b, 0xde, 0x85, 0x78, 0xfd, 0x9b, 0x9c,
	0x42, 0xcb, 0x58, 0xcc, 0xb0, 0x26, 0xf4, 0x95, 0x71, 0x79, 0x52, 0x70, 0xb1, 0x06, 0xb7, 0x15,
	0x6f, 0x80, 0x1d, 0xbf, 0x05, 0xca, 0x04, 0x69, 0xdf, 0x0e, 0x19, 0xb7, 0x50, 0x8f, 0xc1, 0x44,
	0x32, 0x9f, 0x8b, 0xd9, 0x58, 0x31, 0xa9, 0x2e, 0x99, 0xca, 0x37, 0x71, 0x57, 0xdd, 0xe3, 0x75,
	0xd3, 0x5d, 0x2d, 0x63, 0x2e, 0x57, 0x46, 0x75, 0x64, 0x9b, 0x0e, 0xd7, 0x0c, 0xd9, 0x87, 0xea,
	0x39, 0x4b, 0x78, 0x42, 0x5f, 0x9b, 0xb3, 0x56, 0xa7, 0x1a, 0xf4, 0xff, 0x2e, 0xe5, 0x7d, 0xa1,
	0xd3, 0x78, 0x8b, 0x4a, 0xb7, 0x68, 0xb6, 0x87, 0xeb, 0xc2, 0x42, 0x72, 0x0c, 0x55, 0x97, 0x29,
	0x4c, 0xe8, 0x4e, 0xaf, 0xbc, 0xb9, 0x82, 0xac, 0x75, 0x98, 0x42, 0xb7, 0x2a, 0xb5, 0x99, 0xfc,
	0x0c, 0x7b, 0x1a, 0x3a, 0x69, 0xa0, 0x78, 0x1c, 0x70, 0x94, 0x09, 0x2d, 0x1b, 0x8f, 0x57, 0x45,
	0x8f, 0x93, 0x7b, 0x9a, 0x2b, 0xa1, 0xe4, 0xca, 0xdd, 0x93, 0xdb, 0xec, 0xe1, 0x39, 0xec, 0x3f,
	0x24, 0x24, 0x5d, 0x28, 0xcf, 0x71, 0x95, 0xa5, 0xa8, 0x7f, 0xea, 0x93, 0x2d, 0xcc, 0x88, 0xdb,
	0xc7, 0xc2, 0x82, 0x1f, 0x77, 0x7e, 0x28, 0xf5, 0x17, 0x00, 0x9b, 0x2c, 0xf5, 0x3a, 0x70, 0xd8,
	0x32, 0x9b, 0x65, 0xeb, 0xdf, 0x0c, 0x73, 0x42, 0x8f, 0xd0, 0x50, 0x70, 0xc5, 0x59, 0xa0, 0xc5,
	0xd9, 0xb7, 0x5a, 0x7c, 0x43, 0x91, 0x01, 0xec, 0x39, 0x4c, 0x37, 0xaf, 0xd0, 0x9b, 0xcc, 0xa8,
	0xec, 0x43, 0xb4, 0x17, 0x6e, 0xd3, 0xfd, 0x8f, 0xf0, 0xf4, 0x81, 0x3e, 0x24, 0xaf, 0x36, 0xdb,
	0xb4, 0x64, 0x0a, 0xdd, 0xb5, 0xe1, 0xb3, 0x9d, 0xaa, 0xb3, 0xd8, 0x87, 0xaa, 0x59, 0xc5, 0xf9,
	0x59, 0xcc, 0x22, 0xd6, 0xaf, 0xcc, 0xd8, 0xbb, 0x45, 0x3f, 0x0d, 0x6c, 0xc8, 0xb2, 0xdb, 0x48,
	0x32, 0xdc, 0xff, 0xab, 0x54, 0xec, 0xd2, 0xe2, 0xc6, 0x2e, 0x6d, 0x6f, 0xec, 0x43, 0x68, 0xac,
	0x77, 0xb4, 0xfd, 0x7a, 0xc3, 0x2f, 0xbc, 0x14, 0x4e, 0xb4, 0xe0, 0x62, 0x76, 0x1e, 0xa5, 0xc2,
	0x67, 0x72, 0x95, 0x9d, 0xac, 0x13, 0x6e, 0xb1, 0xba, 0xc9, 0xac, 0xce, 0xbc, 0xeb, 0x15, 0x93,
	0x0a, 0x84, 0x6b, 0x46, 0xc7, 0x18, 0x45, 0xbf, 0xd9, 0x13, 0xd8, 0x97, 0xb5, 0x11, 0x64, 0x58,
	0x5f, 0xff, 0x35, 0x9f, 0xdd, 0x5a, 0xa3, 0x7d, 0x4c, 0x9b, 0xb7, 0x39, 0xd1, 0xff, 0x63, 0x6b,
	0x72, 0xfe, 0xe7, 0x31, 0x8a, 0xe1, 0xcb, 0xff, 0x15, 0xbe, 0x72, 0x2f, 0xfc, 0xb4, 0x66, 0xfe,
	0x2f, 0x9d, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xfa, 0xd9, 0xb7, 0x3f, 0x09, 0x00, 0x00,
}
