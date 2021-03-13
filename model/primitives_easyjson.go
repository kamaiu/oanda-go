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

func easyjson821530a6DecodeGithubComKamaiuOandaGoModel(in *jlexer.Lexer, out *Tag) {
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
		case "type":
			out.Type = string(in.String())
		case "name":
			out.Name = string(in.String())
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
func easyjson821530a6EncodeGithubComKamaiuOandaGoModel(out *jwriter.Writer, in Tag) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Tag) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Tag) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Tag) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Tag) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel(l, v)
}
func easyjson821530a6DecodeGithubComKamaiuOandaGoModel1(in *jlexer.Lexer, out *InstrumentFinancing) {
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
		case "longRate":
			out.LongRate = DecimalNumber(in.String())
		case "shortRate":
			out.ShortRate = DecimalNumber(in.String())
		case "financingDaysOfWeek":
			if in.IsNull() {
				in.Skip()
				out.FinancingDaysOfWeek = nil
			} else {
				in.Delim('[')
				if out.FinancingDaysOfWeek == nil {
					if !in.IsDelim(']') {
						out.FinancingDaysOfWeek = make([]FinancingDayOfWeek, 0, 2)
					} else {
						out.FinancingDaysOfWeek = []FinancingDayOfWeek{}
					}
				} else {
					out.FinancingDaysOfWeek = (out.FinancingDaysOfWeek)[:0]
				}
				for !in.IsDelim(']') {
					var v1 FinancingDayOfWeek
					(v1).UnmarshalEasyJSON(in)
					out.FinancingDaysOfWeek = append(out.FinancingDaysOfWeek, v1)
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
func easyjson821530a6EncodeGithubComKamaiuOandaGoModel1(out *jwriter.Writer, in InstrumentFinancing) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"longRate\":"
		out.RawString(prefix[1:])
		out.String(string(in.LongRate))
	}
	{
		const prefix string = ",\"shortRate\":"
		out.RawString(prefix)
		out.String(string(in.ShortRate))
	}
	{
		const prefix string = ",\"financingDaysOfWeek\":"
		out.RawString(prefix)
		if in.FinancingDaysOfWeek == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.FinancingDaysOfWeek {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InstrumentFinancing) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InstrumentFinancing) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InstrumentFinancing) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InstrumentFinancing) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel1(l, v)
}
func easyjson821530a6DecodeGithubComKamaiuOandaGoModel2(in *jlexer.Lexer, out *InstrumentCommission) {
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
		case "commission":
			out.Commission = DecimalNumber(in.String())
		case "unitsTraded":
			out.UnitsTraded = DecimalNumber(in.String())
		case "minimumCommission":
			out.MinimumCommission = DecimalNumber(in.String())
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
func easyjson821530a6EncodeGithubComKamaiuOandaGoModel2(out *jwriter.Writer, in InstrumentCommission) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"commission\":"
		out.RawString(prefix[1:])
		out.String(string(in.Commission))
	}
	{
		const prefix string = ",\"unitsTraded\":"
		out.RawString(prefix)
		out.String(string(in.UnitsTraded))
	}
	{
		const prefix string = ",\"minimumCommission\":"
		out.RawString(prefix)
		out.String(string(in.MinimumCommission))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v InstrumentCommission) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v InstrumentCommission) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *InstrumentCommission) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *InstrumentCommission) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel2(l, v)
}
func easyjson821530a6DecodeGithubComKamaiuOandaGoModel3(in *jlexer.Lexer, out *Instrument) {
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
		case "name":
			out.Name = InstrumentName(in.String())
		case "type":
			out.Type = InstrumentType(in.String())
		case "displayName":
			out.DisplayName = string(in.String())
		case "pipLocation":
			out.PipLocation = int64(in.Int64())
		case "displayPrecision":
			out.DisplayPrecision = int64(in.Int64())
		case "tradeUnitsPrecision":
			out.TradeUnitsPrecision = int64(in.Int64())
		case "minimumTradeSize":
			out.MinimumTradeSize = DecimalNumber(in.String())
		case "maximumTrailingStopDistance":
			out.MaximumTrailingStopDistance = DecimalNumber(in.String())
		case "minimumGuaranteedStopLossDistance":
			out.MinimumGuaranteedStopLossDistance = DecimalNumber(in.String())
		case "minimumTrailingStopDistance":
			out.MinimumTrailingStopDistance = DecimalNumber(in.String())
		case "maximumPositionSize":
			out.MaximumPositionSize = DecimalNumber(in.String())
		case "maximumOrderUnits":
			out.MaximumOrderUnits = DecimalNumber(in.String())
		case "marginRate":
			out.MarginRate = DecimalNumber(in.String())
		case "commission":
			out.Commission = DecimalNumber(in.String())
		case "guaranteedStopLossOrderMode":
			out.GuaranteedStopLossOrderMode = DecimalNumber(in.String())
		case "guaranteedStopLossOrderExecutionPremium":
			out.GuaranteedStopLossOrderExecutionPremium = DecimalNumber(in.String())
		case "guaranteedStopLossOrderLevelRestriction":
			if in.IsNull() {
				in.Skip()
				out.GuaranteedStopLossOrderLevelRestriction = nil
			} else {
				if out.GuaranteedStopLossOrderLevelRestriction == nil {
					out.GuaranteedStopLossOrderLevelRestriction = new(GuaranteedStopLossOrderLevelRestriction)
				}
				(*out.GuaranteedStopLossOrderLevelRestriction).UnmarshalEasyJSON(in)
			}
		case "financing":
			(out.Financing).UnmarshalEasyJSON(in)
		case "tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]Tag, 0, 2)
					} else {
						out.Tags = []Tag{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Tag
					(v4).UnmarshalEasyJSON(in)
					out.Tags = append(out.Tags, v4)
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
func easyjson821530a6EncodeGithubComKamaiuOandaGoModel3(out *jwriter.Writer, in Instrument) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix)
		out.String(string(in.Type))
	}
	{
		const prefix string = ",\"displayName\":"
		out.RawString(prefix)
		out.String(string(in.DisplayName))
	}
	{
		const prefix string = ",\"pipLocation\":"
		out.RawString(prefix)
		out.Int64(int64(in.PipLocation))
	}
	{
		const prefix string = ",\"displayPrecision\":"
		out.RawString(prefix)
		out.Int64(int64(in.DisplayPrecision))
	}
	{
		const prefix string = ",\"tradeUnitsPrecision\":"
		out.RawString(prefix)
		out.Int64(int64(in.TradeUnitsPrecision))
	}
	{
		const prefix string = ",\"minimumTradeSize\":"
		out.RawString(prefix)
		out.String(string(in.MinimumTradeSize))
	}
	{
		const prefix string = ",\"maximumTrailingStopDistance\":"
		out.RawString(prefix)
		out.String(string(in.MaximumTrailingStopDistance))
	}
	{
		const prefix string = ",\"minimumGuaranteedStopLossDistance\":"
		out.RawString(prefix)
		out.String(string(in.MinimumGuaranteedStopLossDistance))
	}
	{
		const prefix string = ",\"minimumTrailingStopDistance\":"
		out.RawString(prefix)
		out.String(string(in.MinimumTrailingStopDistance))
	}
	{
		const prefix string = ",\"maximumPositionSize\":"
		out.RawString(prefix)
		out.String(string(in.MaximumPositionSize))
	}
	{
		const prefix string = ",\"maximumOrderUnits\":"
		out.RawString(prefix)
		out.String(string(in.MaximumOrderUnits))
	}
	{
		const prefix string = ",\"marginRate\":"
		out.RawString(prefix)
		out.String(string(in.MarginRate))
	}
	{
		const prefix string = ",\"commission\":"
		out.RawString(prefix)
		out.String(string(in.Commission))
	}
	{
		const prefix string = ",\"guaranteedStopLossOrderMode\":"
		out.RawString(prefix)
		out.String(string(in.GuaranteedStopLossOrderMode))
	}
	{
		const prefix string = ",\"guaranteedStopLossOrderExecutionPremium\":"
		out.RawString(prefix)
		out.String(string(in.GuaranteedStopLossOrderExecutionPremium))
	}
	{
		const prefix string = ",\"guaranteedStopLossOrderLevelRestriction\":"
		out.RawString(prefix)
		if in.GuaranteedStopLossOrderLevelRestriction == nil {
			out.RawString("null")
		} else {
			(*in.GuaranteedStopLossOrderLevelRestriction).MarshalEasyJSON(out)
		}
	}
	{
		const prefix string = ",\"financing\":"
		out.RawString(prefix)
		(in.Financing).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Tags {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Instrument) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Instrument) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Instrument) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Instrument) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel3(l, v)
}
func easyjson821530a6DecodeGithubComKamaiuOandaGoModel4(in *jlexer.Lexer, out *HomeConversionFactors) {
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
		case "gainQuoteHome":
			(out.GainQuoteHome).UnmarshalEasyJSON(in)
		case "lossQuoteHome":
			(out.LossQuoteHome).UnmarshalEasyJSON(in)
		case "gainBaseHome":
			(out.GainBaseHome).UnmarshalEasyJSON(in)
		case "lossBaseHome":
			(out.LossBaseHome).UnmarshalEasyJSON(in)
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
func easyjson821530a6EncodeGithubComKamaiuOandaGoModel4(out *jwriter.Writer, in HomeConversionFactors) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"gainQuoteHome\":"
		out.RawString(prefix[1:])
		(in.GainQuoteHome).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"lossQuoteHome\":"
		out.RawString(prefix)
		(in.LossQuoteHome).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"gainBaseHome\":"
		out.RawString(prefix)
		(in.GainBaseHome).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"lossBaseHome\":"
		out.RawString(prefix)
		(in.LossBaseHome).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v HomeConversionFactors) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v HomeConversionFactors) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *HomeConversionFactors) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *HomeConversionFactors) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel4(l, v)
}
func easyjson821530a6DecodeGithubComKamaiuOandaGoModel5(in *jlexer.Lexer, out *GuaranteedStopLossOrderLevelRestriction) {
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
		case "volume":
			out.Volume = DecimalNumber(in.String())
		case "priceRange":
			out.PriceRange = DecimalNumber(in.String())
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
func easyjson821530a6EncodeGithubComKamaiuOandaGoModel5(out *jwriter.Writer, in GuaranteedStopLossOrderLevelRestriction) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"volume\":"
		out.RawString(prefix[1:])
		out.String(string(in.Volume))
	}
	{
		const prefix string = ",\"priceRange\":"
		out.RawString(prefix)
		out.String(string(in.PriceRange))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v GuaranteedStopLossOrderLevelRestriction) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v GuaranteedStopLossOrderLevelRestriction) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GuaranteedStopLossOrderLevelRestriction) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GuaranteedStopLossOrderLevelRestriction) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel5(l, v)
}
func easyjson821530a6DecodeGithubComKamaiuOandaGoModel6(in *jlexer.Lexer, out *FinancingDayOfWeek) {
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
		case "dayOfWeek":
			out.DayOfWeek = DayOfWeek(in.String())
		case "daysCharged":
			out.DaysCharged = int64(in.Int64())
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
func easyjson821530a6EncodeGithubComKamaiuOandaGoModel6(out *jwriter.Writer, in FinancingDayOfWeek) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"dayOfWeek\":"
		out.RawString(prefix[1:])
		out.String(string(in.DayOfWeek))
	}
	{
		const prefix string = ",\"daysCharged\":"
		out.RawString(prefix)
		out.Int64(int64(in.DaysCharged))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v FinancingDayOfWeek) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v FinancingDayOfWeek) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *FinancingDayOfWeek) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *FinancingDayOfWeek) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel6(l, v)
}
func easyjson821530a6DecodeGithubComKamaiuOandaGoModel7(in *jlexer.Lexer, out *ConversionFactor) {
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
		case "factor":
			out.Factor = DecimalNumber(in.String())
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
func easyjson821530a6EncodeGithubComKamaiuOandaGoModel7(out *jwriter.Writer, in ConversionFactor) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"factor\":"
		out.RawString(prefix[1:])
		out.String(string(in.Factor))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v ConversionFactor) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v ConversionFactor) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson821530a6EncodeGithubComKamaiuOandaGoModel7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *ConversionFactor) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *ConversionFactor) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson821530a6DecodeGithubComKamaiuOandaGoModel7(l, v)
}