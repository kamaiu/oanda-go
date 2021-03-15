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

func easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel(in *jlexer.Lexer, out *PricingStreamRequest) {
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
		case "instruments":
			if in.IsNull() {
				in.Skip()
				out.Instruments = nil
			} else {
				in.Delim('[')
				if out.Instruments == nil {
					if !in.IsDelim(']') {
						out.Instruments = make([]InstrumentName, 0, 4)
					} else {
						out.Instruments = []InstrumentName{}
					}
				} else {
					out.Instruments = (out.Instruments)[:0]
				}
				for !in.IsDelim(']') {
					var v1 InstrumentName
					v1 = InstrumentName(in.String())
					out.Instruments = append(out.Instruments, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "snapshot":
			out.Snapshot = bool(in.Bool())
		case "includeHomeConversions":
			out.IncludeHomeConversions = bool(in.Bool())
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
func easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel(out *jwriter.Writer, in PricingStreamRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"instruments\":"
		out.RawString(prefix[1:])
		if in.Instruments == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Instruments {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"snapshot\":"
		out.RawString(prefix)
		out.Bool(bool(in.Snapshot))
	}
	{
		const prefix string = ",\"includeHomeConversions\":"
		out.RawString(prefix)
		out.Bool(bool(in.IncludeHomeConversions))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PricingStreamRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PricingStreamRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PricingStreamRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PricingStreamRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel(l, v)
}
func easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel1(in *jlexer.Lexer, out *PricingResponse) {
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
		case "prices":
			if in.IsNull() {
				in.Skip()
				out.Prices = nil
			} else {
				in.Delim('[')
				if out.Prices == nil {
					if !in.IsDelim(']') {
						out.Prices = make([]*ClientPrice, 0, 8)
					} else {
						out.Prices = []*ClientPrice{}
					}
				} else {
					out.Prices = (out.Prices)[:0]
				}
				for !in.IsDelim(']') {
					var v4 *ClientPrice
					if in.IsNull() {
						in.Skip()
						v4 = nil
					} else {
						if v4 == nil {
							v4 = new(ClientPrice)
						}
						(*v4).UnmarshalEasyJSON(in)
					}
					out.Prices = append(out.Prices, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "homeConversions":
			if in.IsNull() {
				in.Skip()
				out.HomeConversions = nil
			} else {
				in.Delim('[')
				if out.HomeConversions == nil {
					if !in.IsDelim(']') {
						out.HomeConversions = make([]*HomeConversions, 0, 8)
					} else {
						out.HomeConversions = []*HomeConversions{}
					}
				} else {
					out.HomeConversions = (out.HomeConversions)[:0]
				}
				for !in.IsDelim(']') {
					var v5 *HomeConversions
					if in.IsNull() {
						in.Skip()
						v5 = nil
					} else {
						if v5 == nil {
							v5 = new(HomeConversions)
						}
						(*v5).UnmarshalEasyJSON(in)
					}
					out.HomeConversions = append(out.HomeConversions, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "time":
			out.Time = DateTime(in.String())
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
func easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel1(out *jwriter.Writer, in PricingResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"prices\":"
		out.RawString(prefix[1:])
		if in.Prices == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Prices {
				if v6 > 0 {
					out.RawByte(',')
				}
				if v7 == nil {
					out.RawString("null")
				} else {
					(*v7).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"homeConversions\":"
		out.RawString(prefix)
		if in.HomeConversions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.HomeConversions {
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
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix)
		out.String(string(in.Time))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PricingResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PricingResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PricingResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PricingResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel1(l, v)
}
func easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel2(in *jlexer.Lexer, out *PricingRequest) {
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
		case "instruments":
			if in.IsNull() {
				in.Skip()
				out.Instruments = nil
			} else {
				in.Delim('[')
				if out.Instruments == nil {
					if !in.IsDelim(']') {
						out.Instruments = make([]InstrumentName, 0, 4)
					} else {
						out.Instruments = []InstrumentName{}
					}
				} else {
					out.Instruments = (out.Instruments)[:0]
				}
				for !in.IsDelim(']') {
					var v10 InstrumentName
					v10 = InstrumentName(in.String())
					out.Instruments = append(out.Instruments, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "since":
			out.Since = DateTime(in.String())
		case "includeUnitsAvailable":
			out.IncludeUnitsAvailable = bool(in.Bool())
		case "includeHomeConversions":
			out.IncludeHomeConversions = bool(in.Bool())
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
func easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel2(out *jwriter.Writer, in PricingRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"instruments\":"
		out.RawString(prefix[1:])
		if in.Instruments == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Instruments {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"since\":"
		out.RawString(prefix)
		out.String(string(in.Since))
	}
	{
		const prefix string = ",\"includeUnitsAvailable\":"
		out.RawString(prefix)
		out.Bool(bool(in.IncludeUnitsAvailable))
	}
	{
		const prefix string = ",\"includeHomeConversions\":"
		out.RawString(prefix)
		out.Bool(bool(in.IncludeHomeConversions))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PricingRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PricingRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PricingRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PricingRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel2(l, v)
}
func easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel3(in *jlexer.Lexer, out *PricingCandlesResponse) {
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
		case "instrument":
			out.Instrument = InstrumentName(in.String())
		case "granularity":
			out.Granularity = CandlestickGranularity(in.String())
		case "candles":
			if in.IsNull() {
				in.Skip()
				out.Candles = nil
			} else {
				in.Delim('[')
				if out.Candles == nil {
					if !in.IsDelim(']') {
						out.Candles = make([]*Candlestick, 0, 8)
					} else {
						out.Candles = []*Candlestick{}
					}
				} else {
					out.Candles = (out.Candles)[:0]
				}
				for !in.IsDelim(']') {
					var v13 *Candlestick
					if in.IsNull() {
						in.Skip()
						v13 = nil
					} else {
						if v13 == nil {
							v13 = new(Candlestick)
						}
						(*v13).UnmarshalEasyJSON(in)
					}
					out.Candles = append(out.Candles, v13)
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
func easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel3(out *jwriter.Writer, in PricingCandlesResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"instrument\":"
		out.RawString(prefix[1:])
		out.String(string(in.Instrument))
	}
	{
		const prefix string = ",\"granularity\":"
		out.RawString(prefix)
		out.String(string(in.Granularity))
	}
	{
		const prefix string = ",\"candles\":"
		out.RawString(prefix)
		if in.Candles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v14, v15 := range in.Candles {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					(*v15).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PricingCandlesResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PricingCandlesResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PricingCandlesResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PricingCandlesResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel3(l, v)
}
func easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel4(in *jlexer.Lexer, out *PricingCandlesRequest) {
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
		case "price":
			out.Price = PricingComponent(in.String())
		case "granularity":
			out.Granularity = CandlestickGranularity(in.String())
		case "count":
			out.Count = int(in.Int())
		case "from":
			out.From = DateTime(in.String())
		case "to":
			out.To = DateTime(in.String())
		case "smooth":
			out.Smooth = bool(in.Bool())
		case "includeFirst":
			out.IncludeFirst = bool(in.Bool())
		case "dailyAlignment":
			out.DailyAlignment = int(in.Int())
		case "alignmentTimezone":
			out.AlignmentTimezone = string(in.String())
		case "weeklyAlignment":
			out.WeeklyAlignment = WeeklyAlignment(in.String())
		case "units":
			out.Units = DecimalNumber(in.String())
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
func easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel4(out *jwriter.Writer, in PricingCandlesRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"price\":"
		out.RawString(prefix[1:])
		out.String(string(in.Price))
	}
	{
		const prefix string = ",\"granularity\":"
		out.RawString(prefix)
		out.String(string(in.Granularity))
	}
	{
		const prefix string = ",\"count\":"
		out.RawString(prefix)
		out.Int(int(in.Count))
	}
	{
		const prefix string = ",\"from\":"
		out.RawString(prefix)
		out.String(string(in.From))
	}
	{
		const prefix string = ",\"to\":"
		out.RawString(prefix)
		out.String(string(in.To))
	}
	{
		const prefix string = ",\"smooth\":"
		out.RawString(prefix)
		out.Bool(bool(in.Smooth))
	}
	{
		const prefix string = ",\"includeFirst\":"
		out.RawString(prefix)
		out.Bool(bool(in.IncludeFirst))
	}
	{
		const prefix string = ",\"dailyAlignment\":"
		out.RawString(prefix)
		out.Int(int(in.DailyAlignment))
	}
	{
		const prefix string = ",\"alignmentTimezone\":"
		out.RawString(prefix)
		out.String(string(in.AlignmentTimezone))
	}
	{
		const prefix string = ",\"weeklyAlignment\":"
		out.RawString(prefix)
		out.String(string(in.WeeklyAlignment))
	}
	{
		const prefix string = ",\"units\":"
		out.RawString(prefix)
		out.String(string(in.Units))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v PricingCandlesRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PricingCandlesRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PricingCandlesRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PricingCandlesRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel4(l, v)
}
func easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel5(in *jlexer.Lexer, out *CandlesLatestResponse) {
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
		case "latestCandles":
			if in.IsNull() {
				in.Skip()
				out.LatestCandles = nil
			} else {
				in.Delim('[')
				if out.LatestCandles == nil {
					if !in.IsDelim(']') {
						out.LatestCandles = make([]*CandlestickResponse, 0, 8)
					} else {
						out.LatestCandles = []*CandlestickResponse{}
					}
				} else {
					out.LatestCandles = (out.LatestCandles)[:0]
				}
				for !in.IsDelim(']') {
					var v16 *CandlestickResponse
					if in.IsNull() {
						in.Skip()
						v16 = nil
					} else {
						if v16 == nil {
							v16 = new(CandlestickResponse)
						}
						(*v16).UnmarshalEasyJSON(in)
					}
					out.LatestCandles = append(out.LatestCandles, v16)
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
func easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel5(out *jwriter.Writer, in CandlesLatestResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"latestCandles\":"
		out.RawString(prefix[1:])
		if in.LatestCandles == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.LatestCandles {
				if v17 > 0 {
					out.RawByte(',')
				}
				if v18 == nil {
					out.RawString("null")
				} else {
					(*v18).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CandlesLatestResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CandlesLatestResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CandlesLatestResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CandlesLatestResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel5(l, v)
}
func easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel6(in *jlexer.Lexer, out *CandlesLatestRequest) {
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
		case "candleSpecifications":
			if in.IsNull() {
				in.Skip()
				out.CandleSpecifications = nil
			} else {
				in.Delim('[')
				if out.CandleSpecifications == nil {
					if !in.IsDelim(']') {
						out.CandleSpecifications = make([]CandleSpecification, 0, 4)
					} else {
						out.CandleSpecifications = []CandleSpecification{}
					}
				} else {
					out.CandleSpecifications = (out.CandleSpecifications)[:0]
				}
				for !in.IsDelim(']') {
					var v19 CandleSpecification
					v19 = CandleSpecification(in.String())
					out.CandleSpecifications = append(out.CandleSpecifications, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "units":
			out.Units = DecimalNumber(in.String())
		case "smooth":
			out.Smooth = bool(in.Bool())
		case "dailyAlignment":
			out.DailyAlignment = int(in.Int())
		case "alignmentTimezone":
			out.AlignmentTimezone = string(in.String())
		case "weeklyAlignment":
			out.WeeklyAlignment = WeeklyAlignment(in.String())
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
func easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel6(out *jwriter.Writer, in CandlesLatestRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"candleSpecifications\":"
		out.RawString(prefix[1:])
		if in.CandleSpecifications == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.CandleSpecifications {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"units\":"
		out.RawString(prefix)
		out.String(string(in.Units))
	}
	{
		const prefix string = ",\"smooth\":"
		out.RawString(prefix)
		out.Bool(bool(in.Smooth))
	}
	{
		const prefix string = ",\"dailyAlignment\":"
		out.RawString(prefix)
		out.Int(int(in.DailyAlignment))
	}
	{
		const prefix string = ",\"alignmentTimezone\":"
		out.RawString(prefix)
		out.String(string(in.AlignmentTimezone))
	}
	{
		const prefix string = ",\"weeklyAlignment\":"
		out.RawString(prefix)
		out.String(string(in.WeeklyAlignment))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v CandlesLatestRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v CandlesLatestRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson18f31bf6EncodeGithubComKamaiuOandaGoModel6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *CandlesLatestRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *CandlesLatestRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson18f31bf6DecodeGithubComKamaiuOandaGoModel6(l, v)
}