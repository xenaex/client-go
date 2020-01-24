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

func easyjson186008b4DecodeGithubComXenaexClientGoXenaXmsg(in *jlexer.Lexer, out *Position) {
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
		case "currency":
			out.Currency = string(in.String())
		case "avgPx":
			out.AvgPx = string(in.String())
		case "positionId":
			out.PositionId = uint64(in.Uint64())
		case "positionOpenTime":
			out.PositionOpenTime = int64(in.Int64())
		case "realizedPL":
			out.RealizedPL = string(in.String())
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
func easyjson186008b4EncodeGithubComXenaexClientGoXenaXmsg(out *jwriter.Writer, in Position) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Currency != "" {
		const prefix string = ",\"currency\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Currency))
	}
	if in.AvgPx != "" {
		const prefix string = ",\"avgPx\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AvgPx))
	}
	if in.PositionId != 0 {
		const prefix string = ",\"positionId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.PositionId))
	}
	if in.PositionOpenTime != 0 {
		const prefix string = ",\"positionOpenTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PositionOpenTime))
	}
	if in.RealizedPL != "" {
		const prefix string = ",\"realizedPL\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RealizedPL))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Position) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson186008b4EncodeGithubComXenaexClientGoXenaXmsg(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Position) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson186008b4DecodeGithubComXenaexClientGoXenaXmsg(l, v)
}
func easyjson186008b4DecodeGithubComXenaexClientGoXenaXmsg1(in *jlexer.Lexer, out *BalanceSnapshotRefresh) {
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
		case "account":
			out.Account = uint64(in.Uint64())
		case "accountStatusRequestId":
			out.AccountStatusRequestId = string(in.String())
		case "lastUpdateTime":
			out.LastUpdateTime = int64(in.Int64())
		case "balances":
			if in.IsNull() {
				in.Skip()
				out.Balances = nil
			} else {
				in.Delim('[')
				if out.Balances == nil {
					if !in.IsDelim(']') {
						out.Balances = make([]*Balance, 0, 8)
					} else {
						out.Balances = []*Balance{}
					}
				} else {
					out.Balances = (out.Balances)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Balance
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Balance)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Balances = append(out.Balances, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "rejectReason":
			out.RejectReason = string(in.String())
		case "text":
			out.Text = string(in.String())
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
func easyjson186008b4EncodeGithubComXenaexClientGoXenaXmsg1(out *jwriter.Writer, in BalanceSnapshotRefresh) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MsgType != "" {
		const prefix string = ",\"msgType\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.MsgType))
	}
	if in.Account != 0 {
		const prefix string = ",\"account\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Account))
	}
	if in.AccountStatusRequestId != "" {
		const prefix string = ",\"accountStatusRequestId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AccountStatusRequestId))
	}
	if in.LastUpdateTime != 0 {
		const prefix string = ",\"lastUpdateTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LastUpdateTime))
	}
	if len(in.Balances) != 0 {
		const prefix string = ",\"balances\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Balances {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.RejectReason != "" {
		const prefix string = ",\"rejectReason\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RejectReason))
	}
	if in.Text != "" {
		const prefix string = ",\"text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Text))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BalanceSnapshotRefresh) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson186008b4EncodeGithubComXenaexClientGoXenaXmsg1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BalanceSnapshotRefresh) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson186008b4DecodeGithubComXenaexClientGoXenaXmsg1(l, v)
}
func easyjson186008b4DecodeGithubComXenaexClientGoXenaXmsg2(in *jlexer.Lexer, out *BalanceIncrementalRefresh) {
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
		case "account":
			out.Account = uint64(in.Uint64())
		case "accountStatusRequestId":
			out.AccountStatusRequestId = string(in.String())
		case "lastUpdateTime":
			out.LastUpdateTime = int64(in.Int64())
		case "balances":
			if in.IsNull() {
				in.Skip()
				out.Balances = nil
			} else {
				in.Delim('[')
				if out.Balances == nil {
					if !in.IsDelim(']') {
						out.Balances = make([]*Balance, 0, 8)
					} else {
						out.Balances = []*Balance{}
					}
				} else {
					out.Balances = (out.Balances)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *Balance
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(Balance)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Balances = append(out.Balances, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "rejectReason":
			out.RejectReason = string(in.String())
		case "text":
			out.Text = string(in.String())
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
func easyjson186008b4EncodeGithubComXenaexClientGoXenaXmsg2(out *jwriter.Writer, in BalanceIncrementalRefresh) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MsgType != "" {
		const prefix string = ",\"msgType\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.MsgType))
	}
	if in.Account != 0 {
		const prefix string = ",\"account\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Account))
	}
	if in.AccountStatusRequestId != "" {
		const prefix string = ",\"accountStatusRequestId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AccountStatusRequestId))
	}
	if in.LastUpdateTime != 0 {
		const prefix string = ",\"lastUpdateTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LastUpdateTime))
	}
	if len(in.Balances) != 0 {
		const prefix string = ",\"balances\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v5, v6 := range in.Balances {
				if v5 > 0 {
					out.RawByte(',')
				}
				if v6 == nil {
					out.RawString("null")
				} else {
					(*v6).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.RejectReason != "" {
		const prefix string = ",\"rejectReason\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RejectReason))
	}
	if in.Text != "" {
		const prefix string = ",\"text\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Text))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BalanceIncrementalRefresh) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson186008b4EncodeGithubComXenaexClientGoXenaXmsg2(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BalanceIncrementalRefresh) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson186008b4DecodeGithubComXenaexClientGoXenaXmsg2(l, v)
}
func easyjson186008b4DecodeGithubComXenaexClientGoXenaXmsg3(in *jlexer.Lexer, out *Balance) {
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
		case "account":
			out.Account = uint64(in.Uint64())
		case "currency":
			out.Currency = string(in.String())
		case "lastUpdateTime":
			out.LastUpdateTime = int64(in.Int64())
		case "available":
			out.Available = string(in.String())
		case "onHold":
			out.OnHold = string(in.String())
		case "settled":
			out.Settled = string(in.String())
		case "equity":
			out.Equity = string(in.String())
		case "bonus":
			out.Bonus = string(in.String())
		case "positions":
			if in.IsNull() {
				in.Skip()
				out.Positions = nil
			} else {
				in.Delim('[')
				if out.Positions == nil {
					if !in.IsDelim(']') {
						out.Positions = make([]*Position, 0, 8)
					} else {
						out.Positions = []*Position{}
					}
				} else {
					out.Positions = (out.Positions)[:0]
				}
				for !in.IsDelim(']') {
					var v7 *Position
					if in.IsNull() {
						in.Skip()
						v7 = nil
					} else {
						if v7 == nil {
							v7 = new(Position)
						}
						(*v7).UnmarshalEasyJSON(in)
					}
					out.Positions = append(out.Positions, v7)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson186008b4EncodeGithubComXenaexClientGoXenaXmsg3(out *jwriter.Writer, in Balance) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Account != 0 {
		const prefix string = ",\"account\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Account))
	}
	if in.Currency != "" {
		const prefix string = ",\"currency\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Currency))
	}
	if in.LastUpdateTime != 0 {
		const prefix string = ",\"lastUpdateTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.LastUpdateTime))
	}
	if in.Available != "" {
		const prefix string = ",\"available\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Available))
	}
	if in.OnHold != "" {
		const prefix string = ",\"onHold\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OnHold))
	}
	if in.Settled != "" {
		const prefix string = ",\"settled\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Settled))
	}
	if in.Equity != "" {
		const prefix string = ",\"equity\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Equity))
	}
	if in.Bonus != "" {
		const prefix string = ",\"bonus\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Bonus))
	}
	if len(in.Positions) != 0 {
		const prefix string = ",\"positions\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.Positions {
				if v8 > 0 {
					out.RawByte(',')
				}
				if v9 == nil {
					out.RawString("null")
				} else {
					(*v9).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Balance) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson186008b4EncodeGithubComXenaexClientGoXenaXmsg3(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Balance) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson186008b4DecodeGithubComXenaexClientGoXenaXmsg3(l, v)
}
func easyjson186008b4DecodeGithubComXenaexClientGoXenaXmsg4(in *jlexer.Lexer, out *AccountStatusReportRequest) {
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
		case "account":
			out.Account = uint64(in.Uint64())
		case "accountStatusRequestId":
			out.AccountStatusRequestId = string(in.String())
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
func easyjson186008b4EncodeGithubComXenaexClientGoXenaXmsg4(out *jwriter.Writer, in AccountStatusReportRequest) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MsgType != "" {
		const prefix string = ",\"msgType\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.MsgType))
	}
	if in.Account != 0 {
		const prefix string = ",\"account\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Uint64(uint64(in.Account))
	}
	if in.AccountStatusRequestId != "" {
		const prefix string = ",\"accountStatusRequestId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.AccountStatusRequestId))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AccountStatusReportRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson186008b4EncodeGithubComXenaexClientGoXenaXmsg4(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AccountStatusReportRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson186008b4DecodeGithubComXenaexClientGoXenaXmsg4(l, v)
}
