// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

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

func easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel(in *jlexer.Lexer, out *PositionsResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
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
					var v1 *Position
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Position)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Positions = append(out.Positions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "lastTransactionID":
			out.LastTransactionID = TransactionID(in.String())
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
func easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel(out *jwriter.Writer, in PositionsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"positions\":"
		out.RawString(prefix[1:])
		if in.Positions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Positions {
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
	{
		const prefix string = ",\"lastTransactionID\":"
		out.RawString(prefix)
		out.String(string(in.LastTransactionID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PositionsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PositionsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PositionsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PositionsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel(l, v)
}
func easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel1(in *jlexer.Lexer, out *PositionResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "position":
			if in.IsNull() {
				in.Skip()
				out.Position = nil
			} else {
				if out.Position == nil {
					out.Position = new(Position)
				}
				(*out.Position).UnmarshalEasyJSON(in)
			}
		case "lastTransactionID":
			out.LastTransactionID = TransactionID(in.String())
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
func easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel1(out *jwriter.Writer, in PositionResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"position\":"
		out.RawString(prefix[1:])
		if in.Position == nil {
			out.RawString("null")
		} else {
			(*in.Position).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"lastTransactionID\":"
		out.RawString(prefix)
		out.String(string(in.LastTransactionID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PositionResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PositionResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PositionResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PositionResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel1(l, v)
}
func easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel2(in *jlexer.Lexer, out *PositionCloseResponse) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "longOrderCreateTransaction":
			if in.IsNull() {
				in.Skip()
				out.LongOrderCreateTransaction = nil
			} else {
				if out.LongOrderCreateTransaction == nil {
					out.LongOrderCreateTransaction = new(MarketOrderTransaction)
				}
				(*out.LongOrderCreateTransaction).UnmarshalEasyJSON(in)
			}
		case "longOrderFillTransaction":
			if in.IsNull() {
				in.Skip()
				out.LongOrderFillTransaction = nil
			} else {
				if out.LongOrderFillTransaction == nil {
					out.LongOrderFillTransaction = new(OrderFillTransaction)
				}
				(*out.LongOrderFillTransaction).UnmarshalEasyJSON(in)
			}
		case "longOrderCancelTransaction":
			if in.IsNull() {
				in.Skip()
				out.LongOrderCancelTransaction = nil
			} else {
				if out.LongOrderCancelTransaction == nil {
					out.LongOrderCancelTransaction = new(OrderCancelTransaction)
				}
				(*out.LongOrderCancelTransaction).UnmarshalEasyJSON(in)
			}
		case "shortOrderCreateTransaction":
			if in.IsNull() {
				in.Skip()
				out.ShortOrderCreateTransaction = nil
			} else {
				if out.ShortOrderCreateTransaction == nil {
					out.ShortOrderCreateTransaction = new(MarketOrderTransaction)
				}
				(*out.ShortOrderCreateTransaction).UnmarshalEasyJSON(in)
			}
		case "shortOrderFillTransaction":
			if in.IsNull() {
				in.Skip()
				out.ShortOrderFillTransaction = nil
			} else {
				if out.ShortOrderFillTransaction == nil {
					out.ShortOrderFillTransaction = new(OrderFillTransaction)
				}
				(*out.ShortOrderFillTransaction).UnmarshalEasyJSON(in)
			}
		case "shortOrderCancelTransaction":
			if in.IsNull() {
				in.Skip()
				out.ShortOrderCancelTransaction = nil
			} else {
				if out.ShortOrderCancelTransaction == nil {
					out.ShortOrderCancelTransaction = new(OrderCancelTransaction)
				}
				(*out.ShortOrderCancelTransaction).UnmarshalEasyJSON(in)
			}
		case "relatedTransactionIDs":
			if in.IsNull() {
				in.Skip()
				out.RelatedTransactionIDs = nil
			} else {
				in.Delim('[')
				if out.RelatedTransactionIDs == nil {
					if !in.IsDelim(']') {
						out.RelatedTransactionIDs = make([]TransactionID, 0, 4)
					} else {
						out.RelatedTransactionIDs = []TransactionID{}
					}
				} else {
					out.RelatedTransactionIDs = (out.RelatedTransactionIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v4 TransactionID
					v4 = TransactionID(in.String())
					out.RelatedTransactionIDs = append(out.RelatedTransactionIDs, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "lastTransactionID":
			out.LastTransactionID = TransactionID(in.String())
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
func easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel2(out *jwriter.Writer, in PositionCloseResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"longOrderCreateTransaction\":"
		out.RawString(prefix[1:])
		if in.LongOrderCreateTransaction == nil {
			out.RawString("null")
		} else {
			(*in.LongOrderCreateTransaction).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"longOrderFillTransaction\":"
		out.RawString(prefix)
		if in.LongOrderFillTransaction == nil {
			out.RawString("null")
		} else {
			(*in.LongOrderFillTransaction).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"longOrderCancelTransaction\":"
		out.RawString(prefix)
		if in.LongOrderCancelTransaction == nil {
			out.RawString("null")
		} else {
			(*in.LongOrderCancelTransaction).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"shortOrderCreateTransaction\":"
		out.RawString(prefix)
		if in.ShortOrderCreateTransaction == nil {
			out.RawString("null")
		} else {
			(*in.ShortOrderCreateTransaction).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"shortOrderFillTransaction\":"
		out.RawString(prefix)
		if in.ShortOrderFillTransaction == nil {
			out.RawString("null")
		} else {
			(*in.ShortOrderFillTransaction).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"shortOrderCancelTransaction\":"
		out.RawString(prefix)
		if in.ShortOrderCancelTransaction == nil {
			out.RawString("null")
		} else {
			(*in.ShortOrderCancelTransaction).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"relatedTransactionIDs\":"
		out.RawString(prefix)
		if in.RelatedTransactionIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.RelatedTransactionIDs {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"lastTransactionID\":"
		out.RawString(prefix)
		out.String(string(in.LastTransactionID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PositionCloseResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PositionCloseResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PositionCloseResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PositionCloseResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel2(l, v)
}
func easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel3(in *jlexer.Lexer, out *PositionCloseRequest) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "longUnits":
			out.LongUnits = string(in.String())
		case "longClientExtensions":
			if in.IsNull() {
				in.Skip()
				out.LongClientExtensions = nil
			} else {
				if out.LongClientExtensions == nil {
					out.LongClientExtensions = new(ClientExtensions)
				}
				(*out.LongClientExtensions).UnmarshalEasyJSON(in)
			}
		case "shortUnits":
			out.ShortUnits = string(in.String())
		case "shortClientExtensions":
			if in.IsNull() {
				in.Skip()
				out.ShortClientExtensions = nil
			} else {
				if out.ShortClientExtensions == nil {
					out.ShortClientExtensions = new(ClientExtensions)
				}
				(*out.ShortClientExtensions).UnmarshalEasyJSON(in)
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
func easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel3(out *jwriter.Writer, in PositionCloseRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"longUnits\":"
		out.RawString(prefix[1:])
		out.String(string(in.LongUnits))
	}
	{
		const prefix string = ",\"longClientExtensions\":"
		out.RawString(prefix)
		if in.LongClientExtensions == nil {
			out.RawString("null")
		} else {
			(*in.LongClientExtensions).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"shortUnits\":"
		out.RawString(prefix)
		out.String(string(in.ShortUnits))
	}
	{
		const prefix string = ",\"shortClientExtensions\":"
		out.RawString(prefix)
		if in.ShortClientExtensions == nil {
			out.RawString("null")
		} else {
			(*in.ShortClientExtensions).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PositionCloseRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PositionCloseRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PositionCloseRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PositionCloseRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel3(l, v)
}
func easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel4(in *jlexer.Lexer, out *PositionCloseError) {
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
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "longOrderRejectTransaction":
			if in.IsNull() {
				in.Skip()
				out.LongOrderRejectTransaction = nil
			} else {
				if out.LongOrderRejectTransaction == nil {
					out.LongOrderRejectTransaction = new(MarketOrderRejectTransaction)
				}
				(*out.LongOrderRejectTransaction).UnmarshalEasyJSON(in)
			}
		case "shortOrderRejectTransaction":
			if in.IsNull() {
				in.Skip()
				out.ShortOrderRejectTransaction = nil
			} else {
				if out.ShortOrderRejectTransaction == nil {
					out.ShortOrderRejectTransaction = new(MarketOrderRejectTransaction)
				}
				(*out.ShortOrderRejectTransaction).UnmarshalEasyJSON(in)
			}
		case "relatedTransactionIDs":
			if in.IsNull() {
				in.Skip()
				out.RelatedTransactionIDs = nil
			} else {
				in.Delim('[')
				if out.RelatedTransactionIDs == nil {
					if !in.IsDelim(']') {
						out.RelatedTransactionIDs = make([]TransactionID, 0, 4)
					} else {
						out.RelatedTransactionIDs = []TransactionID{}
					}
				} else {
					out.RelatedTransactionIDs = (out.RelatedTransactionIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v7 TransactionID
					v7 = TransactionID(in.String())
					out.RelatedTransactionIDs = append(out.RelatedTransactionIDs, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "lastTransactionID":
			out.LastTransactionID = TransactionID(in.String())
		case "errorCode":
			out.ErrorCode = string(in.String())
		case "errorMessage":
			out.ErrorMessage = string(in.String())
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
func easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel4(out *jwriter.Writer, in PositionCloseError) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"longOrderRejectTransaction\":"
		out.RawString(prefix[1:])
		if in.LongOrderRejectTransaction == nil {
			out.RawString("null")
		} else {
			(*in.LongOrderRejectTransaction).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"shortOrderRejectTransaction\":"
		out.RawString(prefix)
		if in.ShortOrderRejectTransaction == nil {
			out.RawString("null")
		} else {
			(*in.ShortOrderRejectTransaction).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"relatedTransactionIDs\":"
		out.RawString(prefix)
		if in.RelatedTransactionIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.RelatedTransactionIDs {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"lastTransactionID\":"
		out.RawString(prefix)
		out.String(string(in.LastTransactionID))
	}
	{
		const prefix string = ",\"errorCode\":"
		out.RawString(prefix)
		out.String(string(in.ErrorCode))
	}
	{
		const prefix string = ",\"errorMessage\":"
		out.RawString(prefix)
		out.String(string(in.ErrorMessage))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PositionCloseError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PositionCloseError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonDe089c11EncodeGithubComKamaiuOandaGoModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PositionCloseError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PositionCloseError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonDe089c11DecodeGithubComKamaiuOandaGoModel4(l, v)
}
