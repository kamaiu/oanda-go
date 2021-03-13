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

func easyjsonB0fe1dfcDecodeGithubComKamaiuOandaGoModel(in *jlexer.Lexer, out *TradeSummary) {
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
		case "id":
			out.Id = TradeID(in.String())
		case "instrument":
			out.Instrument = InstrumentName(in.String())
		case "price":
			out.Price = PriceValue(in.String())
		case "openTime":
			out.OpenTime = DateTime(in.String())
		case "state":
			out.State = TradeState(in.String())
		case "initialUnits":
			out.InitialUnits = DecimalNumber(in.String())
		case "initialMarginRequired":
			out.InitialMarginRequired = AccountUnits(in.String())
		case "currentUnits":
			out.CurrentUnits = DecimalNumber(in.String())
		case "realizedPL":
			out.RealizedPL = AccountUnits(in.String())
		case "unrealizedPL":
			out.UnrealizedPL = AccountUnits(in.String())
		case "marginUsed":
			out.MarginUsed = AccountUnits(in.String())
		case "averageClosePrice":
			out.AverageClosePrice = PriceValue(in.String())
		case "closingTransactionIDs":
			if in.IsNull() {
				in.Skip()
				out.ClosingTransactionIDs = nil
			} else {
				in.Delim('[')
				if out.ClosingTransactionIDs == nil {
					if !in.IsDelim(']') {
						out.ClosingTransactionIDs = make([]TransactionID, 0, 4)
					} else {
						out.ClosingTransactionIDs = []TransactionID{}
					}
				} else {
					out.ClosingTransactionIDs = (out.ClosingTransactionIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 TransactionID
					v1 = TransactionID(in.String())
					out.ClosingTransactionIDs = append(out.ClosingTransactionIDs, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "financing":
			out.Financing = AccountUnits(in.String())
		case "dividendAdjustment":
			out.DividendAdjustment = AccountUnits(in.String())
		case "closeTime":
			out.CloseTime = DateTime(in.String())
		case "clientExtensions":
			if in.IsNull() {
				in.Skip()
				out.ClientExtensions = nil
			} else {
				if out.ClientExtensions == nil {
					out.ClientExtensions = new(ClientExtensions)
				}
				(*out.ClientExtensions).UnmarshalEasyJSON(in)
			}
		case "takeProfitOrderID":
			out.TakeProfitOrderID = OrderID(in.String())
		case "stopLossOrderID":
			out.StopLossOrderID = OrderID(in.String())
		case "guaranteedStopLossOrderID":
			out.GuaranteedStopLossOrderID = OrderID(in.String())
		case "trailingStopLossOrderID":
			out.TrailingStopLossOrderID = OrderID(in.String())
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
func easyjsonB0fe1dfcEncodeGithubComKamaiuOandaGoModel(out *jwriter.Writer, in TradeSummary) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"instrument\":"
		out.RawString(prefix)
		out.String(string(in.Instrument))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.String(string(in.Price))
	}
	{
		const prefix string = ",\"openTime\":"
		out.RawString(prefix)
		out.String(string(in.OpenTime))
	}
	{
		const prefix string = ",\"state\":"
		out.RawString(prefix)
		out.String(string(in.State))
	}
	{
		const prefix string = ",\"initialUnits\":"
		out.RawString(prefix)
		out.String(string(in.InitialUnits))
	}
	{
		const prefix string = ",\"initialMarginRequired\":"
		out.RawString(prefix)
		out.String(string(in.InitialMarginRequired))
	}
	{
		const prefix string = ",\"currentUnits\":"
		out.RawString(prefix)
		out.String(string(in.CurrentUnits))
	}
	{
		const prefix string = ",\"realizedPL\":"
		out.RawString(prefix)
		out.String(string(in.RealizedPL))
	}
	{
		const prefix string = ",\"unrealizedPL\":"
		out.RawString(prefix)
		out.String(string(in.UnrealizedPL))
	}
	{
		const prefix string = ",\"marginUsed\":"
		out.RawString(prefix)
		out.String(string(in.MarginUsed))
	}
	{
		const prefix string = ",\"averageClosePrice\":"
		out.RawString(prefix)
		out.String(string(in.AverageClosePrice))
	}
	{
		const prefix string = ",\"closingTransactionIDs\":"
		out.RawString(prefix)
		if in.ClosingTransactionIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.ClosingTransactionIDs {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"financing\":"
		out.RawString(prefix)
		out.String(string(in.Financing))
	}
	{
		const prefix string = ",\"dividendAdjustment\":"
		out.RawString(prefix)
		out.String(string(in.DividendAdjustment))
	}
	{
		const prefix string = ",\"closeTime\":"
		out.RawString(prefix)
		out.String(string(in.CloseTime))
	}
	{
		const prefix string = ",\"clientExtensions\":"
		out.RawString(prefix)
		if in.ClientExtensions == nil {
			out.RawString("null")
		} else {
			(*in.ClientExtensions).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"takeProfitOrderID\":"
		out.RawString(prefix)
		out.String(string(in.TakeProfitOrderID))
	}
	{
		const prefix string = ",\"stopLossOrderID\":"
		out.RawString(prefix)
		out.String(string(in.StopLossOrderID))
	}
	{
		const prefix string = ",\"guaranteedStopLossOrderID\":"
		out.RawString(prefix)
		out.String(string(in.GuaranteedStopLossOrderID))
	}
	{
		const prefix string = ",\"trailingStopLossOrderID\":"
		out.RawString(prefix)
		out.String(string(in.TrailingStopLossOrderID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v TradeSummary) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB0fe1dfcEncodeGithubComKamaiuOandaGoModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v TradeSummary) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0fe1dfcEncodeGithubComKamaiuOandaGoModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *TradeSummary) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB0fe1dfcDecodeGithubComKamaiuOandaGoModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *TradeSummary) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0fe1dfcDecodeGithubComKamaiuOandaGoModel(l, v)
}
func easyjsonB0fe1dfcDecodeGithubComKamaiuOandaGoModel1(in *jlexer.Lexer, out *Trade) {
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
		case "id":
			out.Id = TradeID(in.String())
		case "instrument":
			out.Instrument = InstrumentName(in.String())
		case "price":
			out.Price = PriceValue(in.String())
		case "openTime":
			out.OpenTime = DateTime(in.String())
		case "state":
			out.State = TradeState(in.String())
		case "initialUnits":
			out.InitialUnits = DecimalNumber(in.String())
		case "initialMarginRequired":
			out.InitialMarginRequired = AccountUnits(in.String())
		case "currentUnits":
			out.CurrentUnits = DecimalNumber(in.String())
		case "realizedPL":
			out.RealizedPL = AccountUnits(in.String())
		case "unrealizedPL":
			out.UnrealizedPL = AccountUnits(in.String())
		case "marginUsed":
			out.MarginUsed = AccountUnits(in.String())
		case "averageClosePrice":
			out.AverageClosePrice = PriceValue(in.String())
		case "closingTransactionIDs":
			if in.IsNull() {
				in.Skip()
				out.ClosingTransactionIDs = nil
			} else {
				in.Delim('[')
				if out.ClosingTransactionIDs == nil {
					if !in.IsDelim(']') {
						out.ClosingTransactionIDs = make([]TransactionID, 0, 4)
					} else {
						out.ClosingTransactionIDs = []TransactionID{}
					}
				} else {
					out.ClosingTransactionIDs = (out.ClosingTransactionIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v4 TransactionID
					v4 = TransactionID(in.String())
					out.ClosingTransactionIDs = append(out.ClosingTransactionIDs, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "financing":
			out.Financing = AccountUnits(in.String())
		case "dividendAdjustment":
			out.DividendAdjustment = AccountUnits(in.String())
		case "closeTime":
			out.CloseTime = DateTime(in.String())
		case "clientExtensions":
			if in.IsNull() {
				in.Skip()
				out.ClientExtensions = nil
			} else {
				if out.ClientExtensions == nil {
					out.ClientExtensions = new(ClientExtensions)
				}
				(*out.ClientExtensions).UnmarshalEasyJSON(in)
			}
		case "takeProfitOrder":
			if in.IsNull() {
				in.Skip()
				out.TakeProfitOrder = nil
			} else {
				if out.TakeProfitOrder == nil {
					out.TakeProfitOrder = new(TakeProfitOrder)
				}
				(*out.TakeProfitOrder).UnmarshalEasyJSON(in)
			}
		case "stopLossOrder":
			if in.IsNull() {
				in.Skip()
				out.StopLossOrder = nil
			} else {
				if out.StopLossOrder == nil {
					out.StopLossOrder = new(StopLossOrder)
				}
				(*out.StopLossOrder).UnmarshalEasyJSON(in)
			}
		case "trailingStopLossOrder":
			if in.IsNull() {
				in.Skip()
				out.TrailingStopLossOrder = nil
			} else {
				if out.TrailingStopLossOrder == nil {
					out.TrailingStopLossOrder = new(TrailingStopLossOrder)
				}
				(*out.TrailingStopLossOrder).UnmarshalEasyJSON(in)
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
func easyjsonB0fe1dfcEncodeGithubComKamaiuOandaGoModel1(out *jwriter.Writer, in Trade) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"instrument\":"
		out.RawString(prefix)
		out.String(string(in.Instrument))
	}
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix)
		out.String(string(in.Price))
	}
	{
		const prefix string = ",\"openTime\":"
		out.RawString(prefix)
		out.String(string(in.OpenTime))
	}
	{
		const prefix string = ",\"state\":"
		out.RawString(prefix)
		out.String(string(in.State))
	}
	{
		const prefix string = ",\"initialUnits\":"
		out.RawString(prefix)
		out.String(string(in.InitialUnits))
	}
	{
		const prefix string = ",\"initialMarginRequired\":"
		out.RawString(prefix)
		out.String(string(in.InitialMarginRequired))
	}
	{
		const prefix string = ",\"currentUnits\":"
		out.RawString(prefix)
		out.String(string(in.CurrentUnits))
	}
	{
		const prefix string = ",\"realizedPL\":"
		out.RawString(prefix)
		out.String(string(in.RealizedPL))
	}
	{
		const prefix string = ",\"unrealizedPL\":"
		out.RawString(prefix)
		out.String(string(in.UnrealizedPL))
	}
	{
		const prefix string = ",\"marginUsed\":"
		out.RawString(prefix)
		out.String(string(in.MarginUsed))
	}
	{
		const prefix string = ",\"averageClosePrice\":"
		out.RawString(prefix)
		out.String(string(in.AverageClosePrice))
	}
	{
		const prefix string = ",\"closingTransactionIDs\":"
		out.RawString(prefix)
		if in.ClosingTransactionIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.ClosingTransactionIDs {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"financing\":"
		out.RawString(prefix)
		out.String(string(in.Financing))
	}
	{
		const prefix string = ",\"dividendAdjustment\":"
		out.RawString(prefix)
		out.String(string(in.DividendAdjustment))
	}
	{
		const prefix string = ",\"closeTime\":"
		out.RawString(prefix)
		out.String(string(in.CloseTime))
	}
	{
		const prefix string = ",\"clientExtensions\":"
		out.RawString(prefix)
		if in.ClientExtensions == nil {
			out.RawString("null")
		} else {
			(*in.ClientExtensions).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"takeProfitOrder\":"
		out.RawString(prefix)
		if in.TakeProfitOrder == nil {
			out.RawString("null")
		} else {
			(*in.TakeProfitOrder).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"stopLossOrder\":"
		out.RawString(prefix)
		if in.StopLossOrder == nil {
			out.RawString("null")
		} else {
			(*in.StopLossOrder).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"trailingStopLossOrder\":"
		out.RawString(prefix)
		if in.TrailingStopLossOrder == nil {
			out.RawString("null")
		} else {
			(*in.TrailingStopLossOrder).MarshalEasyJSON(out)
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Trade) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB0fe1dfcEncodeGithubComKamaiuOandaGoModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Trade) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0fe1dfcEncodeGithubComKamaiuOandaGoModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Trade) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB0fe1dfcDecodeGithubComKamaiuOandaGoModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Trade) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0fe1dfcDecodeGithubComKamaiuOandaGoModel1(l, v)
}
func easyjsonB0fe1dfcDecodeGithubComKamaiuOandaGoModel2(in *jlexer.Lexer, out *CalculatedTradeState) {
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
		case "id":
			out.ID = TradeID(in.String())
		case "unrealizedPL":
			out.UnrealizedPL = AccountUnits(in.String())
		case "marginUsed":
			out.MarginUsed = AccountUnits(in.String())
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
func easyjsonB0fe1dfcEncodeGithubComKamaiuOandaGoModel2(out *jwriter.Writer, in CalculatedTradeState) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"unrealizedPL\":"
		out.RawString(prefix)
		out.String(string(in.UnrealizedPL))
	}
	{
		const prefix string = ",\"marginUsed\":"
		out.RawString(prefix)
		out.String(string(in.MarginUsed))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CalculatedTradeState) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB0fe1dfcEncodeGithubComKamaiuOandaGoModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CalculatedTradeState) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB0fe1dfcEncodeGithubComKamaiuOandaGoModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CalculatedTradeState) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB0fe1dfcDecodeGithubComKamaiuOandaGoModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CalculatedTradeState) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB0fe1dfcDecodeGithubComKamaiuOandaGoModel2(l, v)
}