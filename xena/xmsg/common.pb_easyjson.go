// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package xmsg

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi(in *jlexer.Lexer, out *PriceRange) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "enabled":
			out.Enabled = bool(in.Bool())
		case "distance":
			out.Distance = string(in.String())
		case "movingBoundary":
			out.MovingBoundary = string(in.String())
		case "movingTime":
			out.MovingTime = int64(in.Int64())
		case "lowIndex":
			out.LowIndex = string(in.String())
		case "highIndex":
			out.HighIndex = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi(out *jwriter.Writer, in PriceRange) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Enabled {
		const prefix string = ",\"enabled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Enabled))
	}
	if in.Distance != "" {
		const prefix string = ",\"distance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Distance))
	}
	if in.MovingBoundary != "" {
		const prefix string = ",\"movingBoundary\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MovingBoundary))
	}
	if in.MovingTime != 0 {
		const prefix string = ",\"movingTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.MovingTime))
	}
	if in.LowIndex != "" {
		const prefix string = ",\"lowIndex\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LowIndex))
	}
	if in.HighIndex != "" {
		const prefix string = ",\"highIndex\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HighIndex))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PriceRange) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PriceRange) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi(l, v)
}
func easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi1(in *jlexer.Lexer, out *PriceLimits) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "enabled":
			out.Enabled = bool(in.Bool())
		case "distance":
			out.Distance = string(in.String())
		case "lowIndex":
			out.LowIndex = string(in.String())
		case "highIndex":
			out.HighIndex = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi1(out *jwriter.Writer, in PriceLimits) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Enabled {
		const prefix string = ",\"enabled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Enabled))
	}
	if in.Distance != "" {
		const prefix string = ",\"distance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Distance))
	}
	if in.LowIndex != "" {
		const prefix string = ",\"lowIndex\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LowIndex))
	}
	if in.HighIndex != "" {
		const prefix string = ",\"highIndex\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.HighIndex))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PriceLimits) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PriceLimits) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi1(l, v)
}
func easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi2(in *jlexer.Lexer, out *MsgTypeHeader) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "MsgType":
			out.MsgType = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi2(out *jwriter.Writer, in MsgTypeHeader) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MsgType != "" {
		const prefix string = ",\"MsgType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MsgType))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MsgTypeHeader) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MsgTypeHeader) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi2(l, v)
}
func easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi3(in *jlexer.Lexer, out *MarginRate) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "maxVolume":
			out.MaxVolume = string(in.String())
		case "initialRate":
			out.InitialRate = string(in.String())
		case "maintenanceRate":
			out.MaintenanceRate = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi3(out *jwriter.Writer, in MarginRate) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MaxVolume != "" {
		const prefix string = ",\"maxVolume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MaxVolume))
	}
	if in.InitialRate != "" {
		const prefix string = ",\"initialRate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.InitialRate))
	}
	if in.MaintenanceRate != "" {
		const prefix string = ",\"maintenanceRate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MaintenanceRate))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v MarginRate) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *MarginRate) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi3(l, v)
}
func easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi4(in *jlexer.Lexer, out *Margin) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "netting":
			out.Netting = string(in.String())
		case "rates":
			if in.IsNull() {
				in.Skip()
				out.Rates = nil
			} else {
				in.Delim('[')
				if out.Rates == nil {
					if !in.IsDelim(']') {
						out.Rates = make([]*MarginRate, 0, 8)
					} else {
						out.Rates = []*MarginRate{}
					}
				} else {
					out.Rates = (out.Rates)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *MarginRate
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(MarginRate)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Rates = append(out.Rates, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "rateMultipliers":
			if in.IsNull() {
				in.Skip()
			} else {
				in.Delim('{')
				if !in.IsDelim('}') {
					out.RateMultipliers = make(map[string]string)
				} else {
					out.RateMultipliers = nil
				}
				for !in.IsDelim('}') {
					key := string(in.String())
					in.WantColon()
					var v2 string
					v2 = string(in.String())
					(out.RateMultipliers)[key] = v2
					in.WantComma()
				}
				in.Delim('}')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi4(out *jwriter.Writer, in Margin) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Netting != "" {
		const prefix string = ",\"netting\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Netting))
	}
	if len(in.Rates) != 0 {
		const prefix string = ",\"rates\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v3, v4 := range in.Rates {
				if v3 > 0 {
					out.RawByte(',')
				}
				if v4 == nil {
					out.RawString("null")
				} else {
					(*v4).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.RateMultipliers) != 0 {
		const prefix string = ",\"rateMultipliers\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('{')
			v5First := true
			for v5Name, v5Value := range in.RateMultipliers {
				if v5First {
					v5First = false
				} else {
					out.RawByte(',')
				}
				out.String(string(v5Name))
				out.RawByte(':')
				out.String(string(v5Value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Margin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi4(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Margin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi4(l, v)
}
func easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi5(in *jlexer.Lexer, out *Instrument) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "type":
			out.Type = string(in.String())
		case "symbol":
			out.Symbol = string(in.String())
		case "baseCurrency":
			out.BaseCurrencyName = string(in.String())
		case "quoteCurrency":
			out.QuoteCurrencyName = string(in.String())
		case "settlCurrency":
			out.SettlCurrencyName = string(in.String())
		case "tickSize":
			out.TickSize = int32(in.Int32())
		case "minOrderQuantity":
			out.MinOrderQty = string(in.String())
		case "orderQtyStep":
			out.OrderQtyStep = string(in.String())
		case "limitOrderMaxDistance":
			out.LimitOrderMaxDistance = string(in.String())
		case "priceInputMask":
			out.PriceInputMask = string(in.String())
		case "indexes":
			if in.IsNull() {
				in.Skip()
				out.Indexes = nil
			} else {
				in.Delim('[')
				if out.Indexes == nil {
					if !in.IsDelim(']') {
						out.Indexes = make([]string, 0, 4)
					} else {
						out.Indexes = []string{}
					}
				} else {
					out.Indexes = (out.Indexes)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.Indexes = append(out.Indexes, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "enabled":
			out.Enabled = bool(in.Bool())
		case "liquidationMaxDistance":
			out.LiquidationMaxDistance = string(in.String())
		case "contractValue":
			out.ContractValue = string(in.String())
		case "contractCurrency":
			out.ContractCurrency = string(in.String())
		case "lotSize":
			out.LotSize = string(in.String())
		case "tickValue":
			out.TickValue = string(in.String())
		case "maxOrderQty":
			out.MaxOrderQty = string(in.String())
		case "maxPosVolume":
			out.MaxPosVolume = string(in.String())
		case "mark":
			out.Mark = string(in.String())
		case "floatingPL":
			out.FloatingPL = string(in.String())
		case "addUvmToFreeMargin":
			out.AddUvmToFreeMargin = string(in.String())
		case "minLeverage":
			out.MinLeverage = string(in.String())
		case "maxLeverage":
			out.MaxLeverage = string(in.String())
		case "margin":
			if in.IsNull() {
				in.Skip()
				out.Margin = nil
			} else {
				if out.Margin == nil {
					out.Margin = new(Margin)
				}
				(*out.Margin).UnmarshalEasyJSON(in)
			}
		case "clearing":
			if in.IsNull() {
				in.Skip()
				out.Clearing = nil
			} else {
				if out.Clearing == nil {
					out.Clearing = new(DerivativeOperation)
				}
				(*out.Clearing).UnmarshalEasyJSON(in)
			}
		case "interest":
			if in.IsNull() {
				in.Skip()
				out.Interest = nil
			} else {
				if out.Interest == nil {
					out.Interest = new(DerivativeOperation)
				}
				(*out.Interest).UnmarshalEasyJSON(in)
			}
		case "premium":
			if in.IsNull() {
				in.Skip()
				out.Premium = nil
			} else {
				if out.Premium == nil {
					out.Premium = new(DerivativeOperation)
				}
				(*out.Premium).UnmarshalEasyJSON(in)
			}
		case "riskAdjustment":
			if in.IsNull() {
				in.Skip()
				out.RiskAdjustment = nil
			} else {
				if out.RiskAdjustment == nil {
					out.RiskAdjustment = new(DerivativeOperation)
				}
				(*out.RiskAdjustment).UnmarshalEasyJSON(in)
			}
		case "pricePrecision":
			out.PricePrecision = int32(in.Int32())
		case "priceRange":
			if in.IsNull() {
				in.Skip()
				out.PriceRange = nil
			} else {
				if out.PriceRange == nil {
					out.PriceRange = new(PriceRange)
				}
				(*out.PriceRange).UnmarshalEasyJSON(in)
			}
		case "priceLimits":
			if in.IsNull() {
				in.Skip()
				out.PriceLimits = nil
			} else {
				if out.PriceLimits == nil {
					out.PriceLimits = new(PriceLimits)
				}
				(*out.PriceLimits).UnmarshalEasyJSON(in)
			}
		case "inverse":
			out.Inverse = bool(in.Bool())
		case "tradingStartDate":
			out.TradingStartDate = string(in.String())
		case "expiryDate":
			out.ExpiryDate = string(in.String())
		case "basis":
			out.Basis = int32(in.Int32())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi5(out *jwriter.Writer, in Instrument) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ID))
	}
	if in.Type != "" {
		const prefix string = ",\"type\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Type))
	}
	if in.Symbol != "" {
		const prefix string = ",\"symbol\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Symbol))
	}
	if in.BaseCurrencyName != "" {
		const prefix string = ",\"baseCurrency\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.BaseCurrencyName))
	}
	if in.QuoteCurrencyName != "" {
		const prefix string = ",\"quoteCurrency\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.QuoteCurrencyName))
	}
	if in.SettlCurrencyName != "" {
		const prefix string = ",\"settlCurrency\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.SettlCurrencyName))
	}
	if in.TickSize != 0 {
		const prefix string = ",\"tickSize\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.TickSize))
	}
	if in.MinOrderQty != "" {
		const prefix string = ",\"minOrderQuantity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MinOrderQty))
	}
	if in.OrderQtyStep != "" {
		const prefix string = ",\"orderQtyStep\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OrderQtyStep))
	}
	if in.LimitOrderMaxDistance != "" {
		const prefix string = ",\"limitOrderMaxDistance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LimitOrderMaxDistance))
	}
	if in.PriceInputMask != "" {
		const prefix string = ",\"priceInputMask\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PriceInputMask))
	}
	if len(in.Indexes) != 0 {
		const prefix string = ",\"indexes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v7, v8 := range in.Indexes {
				if v7 > 0 {
					out.RawByte(',')
				}
				out.String(string(v8))
			}
			out.RawByte(']')
		}
	}
	if in.Enabled {
		const prefix string = ",\"enabled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Enabled))
	}
	if in.LiquidationMaxDistance != "" {
		const prefix string = ",\"liquidationMaxDistance\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LiquidationMaxDistance))
	}
	if in.ContractValue != "" {
		const prefix string = ",\"contractValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ContractValue))
	}
	if in.ContractCurrency != "" {
		const prefix string = ",\"contractCurrency\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ContractCurrency))
	}
	if in.LotSize != "" {
		const prefix string = ",\"lotSize\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.LotSize))
	}
	if in.TickValue != "" {
		const prefix string = ",\"tickValue\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TickValue))
	}
	if in.MaxOrderQty != "" {
		const prefix string = ",\"maxOrderQty\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MaxOrderQty))
	}
	if in.MaxPosVolume != "" {
		const prefix string = ",\"maxPosVolume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MaxPosVolume))
	}
	if in.Mark != "" {
		const prefix string = ",\"mark\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Mark))
	}
	if in.FloatingPL != "" {
		const prefix string = ",\"floatingPL\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.FloatingPL))
	}
	if in.AddUvmToFreeMargin != "" {
		const prefix string = ",\"addUvmToFreeMargin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AddUvmToFreeMargin))
	}
	if in.MinLeverage != "" {
		const prefix string = ",\"minLeverage\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MinLeverage))
	}
	if in.MaxLeverage != "" {
		const prefix string = ",\"maxLeverage\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MaxLeverage))
	}
	if in.Margin != nil {
		const prefix string = ",\"margin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Margin).MarshalEasyJSON(out)
	}
	if in.Clearing != nil {
		const prefix string = ",\"clearing\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Clearing).MarshalEasyJSON(out)
	}
	if in.Interest != nil {
		const prefix string = ",\"interest\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Interest).MarshalEasyJSON(out)
	}
	if in.Premium != nil {
		const prefix string = ",\"premium\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.Premium).MarshalEasyJSON(out)
	}
	if in.RiskAdjustment != nil {
		const prefix string = ",\"riskAdjustment\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.RiskAdjustment).MarshalEasyJSON(out)
	}
	if in.PricePrecision != 0 {
		const prefix string = ",\"pricePrecision\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.PricePrecision))
	}
	if in.PriceRange != nil {
		const prefix string = ",\"priceRange\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.PriceRange).MarshalEasyJSON(out)
	}
	if in.PriceLimits != nil {
		const prefix string = ",\"priceLimits\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		(*in.PriceLimits).MarshalEasyJSON(out)
	}
	if in.Inverse {
		const prefix string = ",\"inverse\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Inverse))
	}
	if in.TradingStartDate != "" {
		const prefix string = ",\"tradingStartDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TradingStartDate))
	}
	if in.ExpiryDate != "" {
		const prefix string = ",\"expiryDate\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ExpiryDate))
	}
	if in.Basis != 0 {
		const prefix string = ",\"basis\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.Basis))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Instrument) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi5(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Instrument) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi5(l, v)
}
func easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi6(in *jlexer.Lexer, out *Heartbeat) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "msgType":
			out.MsgType = string(in.String())
		case "testReqId":
			out.TestReqId = string(in.String())
		case "transactTime":
			out.TransactTime = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi6(out *jwriter.Writer, in Heartbeat) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MsgType != "" {
		const prefix string = ",\"msgType\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.MsgType))
	}
	if in.TestReqId != "" {
		const prefix string = ",\"testReqId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TestReqId))
	}
	if in.TransactTime != 0 {
		const prefix string = ",\"transactTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.TransactTime))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Heartbeat) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi6(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Heartbeat) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi6(l, v)
}
func easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi7(in *jlexer.Lexer, out *DerivativeOperation) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "maxVolenabledume":
			out.Enabled = bool(in.Bool())
		case "index":
			out.Index = string(in.String())
		case "schedule":
			out.Schedule = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi7(out *jwriter.Writer, in DerivativeOperation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Enabled {
		const prefix string = ",\"maxVolenabledume\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Bool(bool(in.Enabled))
	}
	if in.Index != "" {
		const prefix string = ",\"index\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Index))
	}
	if in.Schedule != 0 {
		const prefix string = ",\"schedule\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Schedule))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DerivativeOperation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89b94fcfEncodeGitXenaIoXenaInterfacesPublicApi7(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DerivativeOperation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89b94fcfDecodeGitXenaIoXenaInterfacesPublicApi7(l, v)
}
