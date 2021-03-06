//go:generate easyjson -all $GOFILE
package model

import (
	"fmt"
	"github.com/mailru/easyjson/jlexer"
	"strings"
	"time"
)

// OANDA docs - http://developer.oanda.com/rest-live-v20/pricing-ep/

// A PriceBucket represents a price available for an amount of liquidity
type PriceBucket struct {
	// The Price offered by the PriceBucket
	Price PriceValue `json:"price"`
	// The amount of liquidity offered by the PriceBucket
	Liquidity int64 `json:"liquidity"`
}

// The specification of an Account-specific Price.
type ClientPrice struct {
	// The string “PRICE”. Used to identify the a Price object when found in a stream.
	Type string `json:"type"`
	// The Price’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The date/time when the Price was created
	Time DateTime `json:"time"`
	// Flag indicating if the Price is tradeable or not
	Tradeable bool `json:"tradeable"`
	// The list of prices and liquidity available on the Instrument’s bid side.
	// It is possible for this list to be empty if there is no bid liquidity
	// currently available for the Instrument in the Account.
	Bids []PriceBucket `json:"bids"`
	// The list of prices and liquidity available on the Instrument’s ask side.
	// It is possible for this list to be empty if there is no ask liquidity
	// currently available for the Instrument in the Account.
	Asks []PriceBucket `json:"asks"`
	// The closeout bid Price. This Price is used when a bid is required to
	// closeout a Position (margin closeout or manual) yet there is no bid
	// liquidity. The closeout bid is never used to open a new position.
	CloseoutBid PriceValue `json:"closeoutBid"`
	// The closeout ask Price. This Price is used when a ask is required to
	// closeout a Position (margin closeout or manual) yet there is no ask
	// liquidity. The closeout ask is never used to open a new position.
	CloseoutAsk PriceValue `json:"closeoutAsk"`
}

// QuoteHomeConversionFactors represents the factors that can be used used to convert
// quantities of a Price’s Instrument’s quote currency into the Account’s home currency.
type QuoteHomeConversionFactors struct {
	// The factor used to convert a positive amount of the Price’s Instrument’s
	// quote currency into a positive amount of the Account’s home currency.
	// Conversion is performed by multiplying the quote units by the conversion
	// factor.
	PositiveUnits DecimalNumber `json:"positiveUnits"`
	// The factor used to convert a negative amount of the Price’s Instrument’s
	// quote currency into a negative amount of the Account’s home currency.
	// Conversion is performed by multiplying the quote units by the conversion
	// factor.
	NegativeUnits DecimalNumber `json:"negativeUnits"`
}

// A PricingHeartbeat object is injected into the Pricing stream to ensure that the
// HTTP connection remains active.
type PricingHeartbeat struct {
	// The string “HEARTBEAT”
	Type string `json:"type"`
	// The date/time when the PricingHeartbeat was created.
	Time DateTime `json:"time"`
}

// An instrument name, a granularity, and a price component to get candlestick data for.
// A string containing the following, all delimited by “:” characters:
// 1) InstrumentName
// 2) CandlestickGranularity
// 3) PricingComponent
// 		e.g. EUR_USD:S10:BM
type CandleSpecification string

func NewCandleSpecification(
	instrument InstrumentName,
	granularity CandlestickGranularity,
	price PricingComponent,
) CandleSpecification {
	return CandleSpecification(fmt.Sprintf("%s:%s:%s", instrument, granularity, price))
}

func (c CandleSpecification) Parse() (InstrumentName, string, PricingComponent) {
	var (
		s           = string(c)
		instrument  InstrumentName
		granularity string
		component   PricingComponent
	)
	i := 0
	for ; i < len(s); i++ {
		if s[i] == ':' {
			instrument = InstrumentName(strings.TrimSpace(s[0:i]))
			s = s[i+1:]
			i = 0
			break
		}
	}
	for ; i < len(c); i++ {
		if c[i] == ':' {
			granularity = strings.TrimSpace(s[0:i])
			component = PricingComponent(strings.TrimSpace(s[i+1:]))
			break
		}
	}
	return instrument, granularity, component
}

// HomeConversions represents the factors to use to convert quantities
// of a given currency into the Account’s home currency. The conversion factor depends
// on the scenario the conversion is required for.
type HomeConversions struct {
	// The currency to be converted into the home currency.
	Currency Currency `json:"currency"`
	// The factor used to convert any gains for an Account in the specified
	// currency into the Account’s home currency. This would include positive
	// realized P/L and positive financing amounts. Conversion is performed by
	// multiplying the positive P/L by the conversion factor.
	AccountGain DecimalNumber `json:"accountGain"`
	// The factor used to convert any losses for an Account in the specified
	// currency into the Account’s home currency. This would include negative
	// realized P/L and negative financing amounts. Conversion is performed by
	// multiplying the positive P/L by the conversion factor.
	AccountLoss DecimalNumber `json:"accountLoss"`
	// The factor used to convert a Position or Trade Value in the specified
	// currency into the Account’s home currency. Conversion is performed by
	// multiplying the Position or Trade Value by the conversion factor.
	PositionValue DecimalNumber `json:"positionValue"`
}

//easyjson:skip
type StreamClientPrice struct {
	instrument  [32]byte
	Instrument  []byte
	Time        time.Time
	bids        [8]StreamPriceBucket
	Bids        []StreamPriceBucket
	asks        [8]StreamPriceBucket
	Asks        []StreamPriceBucket
	CloseoutBid float64
	CloseoutAsk float64
	Tradeable   bool
	IsHeartbeat bool
}

//easyjson:skip
type StreamPriceBucket struct {
	// The Price offered by the PriceBucket
	Price float64
	// The amount of liquidity offered by the PriceBucket
	Liquidity int64
}

func (out *StreamPriceBucket) UnmarshalEasyJSON(in *jlexer.Lexer) {
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
			out.Price = (PriceValue)(in.UnsafeString()).AsFloat64(0)
		case "liquidity":
			out.Liquidity = in.Int64()
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

func (out *StreamClientPrice) UnmarshalJSON(b []byte) {
	in := &jlexer.Lexer{Data: b}
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
			s := string(in.String())
			out.IsHeartbeat = s == "HEARTBEAT"
			//out.Type = string(in.String())
		case "instrument":
			s := in.UnsafeString()
			if len(s) > len(out.instrument) {
				s = s[0:len(out.Instrument)]
			}
			out.Instrument = out.instrument[0:len(s)]
			copy(out.Instrument, s)
		case "time":
			out.Time, _ = (DateTime)(in.UnsafeString()).Parse()
		case "tradeable":
			out.Tradeable = bool(in.Bool())
		case "bids":
			if in.IsNull() {
				in.Skip()
				out.Bids = nil
			} else {
				in.Delim('[')
				out.Bids = out.bids[:0]
				for !in.IsDelim(']') {
					var v1 StreamPriceBucket
					(v1).UnmarshalEasyJSON(in)
					if len(out.Bids)+1 == cap(out.bids) {
						bids := make([]StreamPriceBucket, len(out.Bids), len(out.Bids)+1)
						copy(bids, out.Bids)
						out.Bids = bids
					}
					out.Bids = append(out.Bids, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "asks":
			if in.IsNull() {
				in.Skip()
				out.Asks = nil
			} else {
				in.Delim('[')
				out.Asks = out.asks[:0]
				for !in.IsDelim(']') {
					var v2 StreamPriceBucket
					(v2).UnmarshalEasyJSON(in)
					if len(out.Asks)+1 == cap(out.asks) {
						asks := make([]StreamPriceBucket, len(out.Asks), len(out.Asks)+1)
						copy(asks, out.Asks)
						out.Asks = asks
					}
					out.Asks = append(out.Asks, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "closeoutBid":
			out.CloseoutBid = (PriceValue)(in.UnsafeString()).AsFloat64(0)
		case "closeoutAsk":
			out.CloseoutAsk = (PriceValue)(in.UnsafeString()).AsFloat64(0)
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
