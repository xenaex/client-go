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

func easyjsonA8fbe0d0DecodeGithubComXenaexClientGoXenaXmsg(in *jlexer.Lexer, out *Logon) {
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
		case "HeartBtInt":
			out.HeartBtInt = int32(in.Int32())
		case "RejectText":
			out.RejectText = string(in.String())
		case "Account":
			if in.IsNull() {
				in.Skip()
				out.Account = nil
			} else {
				in.Delim('[')
				if out.Account == nil {
					if !in.IsDelim(']') {
						out.Account = make([]uint64, 0, 8)
					} else {
						out.Account = []uint64{}
					}
				} else {
					out.Account = (out.Account)[:0]
				}
				for !in.IsDelim(']') {
					var v1 uint64
					v1 = uint64(in.Uint64())
					out.Account = append(out.Account, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "SendingTime":
			out.SendingTime = int64(in.Int64())
		case "CstmApplVerId":
			out.CstmApplVerId = string(in.String())
		case "Username":
			out.Username = string(in.String())
		case "Password":
			out.Password = string(in.String())
		case "RawData":
			out.RawData = string(in.String())
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
func easyjsonA8fbe0d0EncodeGithubComXenaexClientGoXenaXmsg(out *jwriter.Writer, in Logon) {
	out.RawByte('{')
	first := true
	_ = first
	if in.MsgType != "" {
		const prefix string = ",\"MsgType\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.MsgType))
	}
	if in.HeartBtInt != 0 {
		const prefix string = ",\"HeartBtInt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int32(int32(in.HeartBtInt))
	}
	if in.RejectText != "" {
		const prefix string = ",\"RejectText\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RejectText))
	}
	if len(in.Account) != 0 {
		const prefix string = ",\"Account\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Account {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.Uint64(uint64(v3))
			}
			out.RawByte(']')
		}
	}
	if in.SendingTime != 0 {
		const prefix string = ",\"SendingTime\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.SendingTime))
	}
	if in.CstmApplVerId != "" {
		const prefix string = ",\"CstmApplVerId\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.CstmApplVerId))
	}
	if in.Username != "" {
		const prefix string = ",\"Username\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Username))
	}
	if in.Password != "" {
		const prefix string = ",\"Password\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Password))
	}
	if in.RawData != "" {
		const prefix string = ",\"RawData\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.RawData))
	}
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Logon) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8fbe0d0EncodeGithubComXenaexClientGoXenaXmsg(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Logon) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8fbe0d0DecodeGithubComXenaexClientGoXenaXmsg(l, v)
}
func easyjsonA8fbe0d0DecodeGithubComXenaexClientGoXenaXmsg1(in *jlexer.Lexer, out *AccountInfo) {
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
			out.Id = uint64(in.Uint64())
		case "kind":
			out.Kind = string(in.String())
		case "currency":
			out.Currency = string(in.String())
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
func easyjsonA8fbe0d0EncodeGithubComXenaexClientGoXenaXmsg1(out *jwriter.Writer, in AccountInfo) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Id != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Uint64(uint64(in.Id))
	}
	if in.Kind != "" {
		const prefix string = ",\"kind\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Kind))
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
	out.RawByte('}')
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AccountInfo) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonA8fbe0d0EncodeGithubComXenaexClientGoXenaXmsg1(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AccountInfo) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonA8fbe0d0DecodeGithubComXenaexClientGoXenaXmsg1(l, v)
}
