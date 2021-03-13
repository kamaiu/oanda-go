//go:generate easyjson -all $GOFILE
package model

import (
	"strings"
)

// Supporting OANDA docs - http://developer.oanda.com/rest-live-v20/pricing-ep/

// A PriceBucket represents a price available for an amount of liquidity
type PriceBucket struct {
	// The Price offered by the PriceBucket
	Price PriceValue `json:"price"`
	// The amount of liquidity offered by the PriceBucket
	Liquidity int64 `json:"liquidity"`
}

// The status of the Price.
// DEPRECATED
//type PriceStatus string
//
//const (
//	// The Instrument’s price is tradeable.
//	PriceStatusTradeable PriceStatus = "tradeable"
//	// The Instrument’s price is not tradeable.
//	PriceStatusNonTradeable PriceStatus = "non-tradeable"
//	// The Instrument of the price is invalid or there is no valid Price for the Instrument.
//	PriceStatusNonInvalid PriceStatus = "invalid"
//)

type ClientPrice struct {
	Type       string         `json:"type"`
	Instrument InstrumentName `json:"instrument"`
	Time       DateTime       `json:"time"`
	//Status     PriceStatus          `json:"status"` // deprecated
	Tradeable   bool          `json:"tradeable"`
	Bids        []PriceBucket `json:"bids"`
	Asks        []PriceBucket `json:"asks"`
	CloseoutBid PriceValue    `json:"closeoutBid"`
	CloseoutAsk PriceValue    `json:"closeoutAsk"`
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

// A PricingHeartbeat object is injected into the Pricing stream to ensure that the HTTP connection remains active.
type Heartbeat struct {
	// The string “HEARTBEAT”
	Type string `json:"type"`
	// The date/time when the Heartbeat was created.
	Time DateTime `json:"time"`
}

// An instrument name, a granularity, and a price component to get candlestick data for.
// A string containing the following, all delimited by “:” characters:
// 1) InstrumentName
// 2) CandlestickGranularity
// 3) PricingComponent
// 		e.g. EUR_USD:S10:BM
type CandleSpecification string

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
