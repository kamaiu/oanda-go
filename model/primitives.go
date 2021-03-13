//go:generate easyjson -all $GOFILE
package model

import (
	"strconv"
	"time"
)

// The string representation of a decimal number.
// A decimal number encoded as a string. The amount of precision provided depends on what the number represents.
type DecimalNumber string

func (d DecimalNumber) AsFloat64(or float64) float64 {
	if len(d) == 0 {
		return or
	}
	f, err := strconv.ParseFloat((string)(d), 64)
	if err != nil {
		return or
	}
	return f
}

// The string representation of a Price for a Bucket.
// A decimal number encodes as a string. The amount of precision provided depends on the Instrument.
type PriceValue DecimalNumber

// The string representation of a quantity of an Account’s home currency.
// A decimal number encoded as a string. The amount of precision provided depends on the Account’s home currency.
type AccountUnits DecimalNumber

// DateTime	A date and time value using either RFC3339 or UNIX time representation.
// The RFC 3339 representation is a string conforming to https://tools.ietf.org/rfc/rfc3339.txt.
// The Unix representation is a string representing the number of seconds since the
// Unix Epoch (January 1st, 1970 at UTC). The value is a fractional number, where the fractional
// part represents a fraction of a second (up to nine decimal places).
type DateTime string

func (d DateTime) UNIX() (time.Time, error) {
	if len(d) == 0 {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, (string)(d))
}

func (d DateTime) RFC3339() (time.Time, error) {
	if len(d) == 0 {
		return time.Time{}, nil
	}
	return time.Parse(time.RFC3339, (string)(d))
}

func (d DateTime) Parse() (time.Time, error) {
	if len(d) <= 1 {
		return time.Time{}, nil
	}

	for i := 0; i < len(d); i++ {
		c := d[i]
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '.':
			secs, err := strconv.ParseInt((string)(d[0:i]), 10, 64)
			if err != nil {
				return time.Time{}, err
			}
			f := (string)(d[i+1:])
			if len(f) == 0 {
				return time.Unix(secs, 0), nil
			}
			nanos, err := strconv.ParseInt(f, 10, 64)
			if err != nil {
				return time.Time{}, err
			}

			switch len(f) {
			case 1:
				nanos *= 100000000
			case 2:
				nanos *= 10000000
			case 3:
				nanos *= 1000000
			case 4:
				nanos *= 100000
			case 5:
				nanos *= 10000
			case 6:
				nanos *= 1000
			case 7:
				nanos *= 100
			case 8:
				nanos *= 10
			default:
				if nanos > 1000000000 {
					nanos = 1000000000
				}
			}
			return time.Unix(secs, nanos), nil

		default:
			return time.Parse(time.RFC3339, (string)(d))
		}
	}
	secs, err := strconv.ParseInt((string)(d), 10, 64)
	if err != nil {
		return time.Time{}, nil
	}
	return time.Unix(secs, 0), nil
}

// The request identifier
type RequestID string

// A client provided request identifier
type ClientRequestID string

// Currency name identifier. Used by clients to refer to currencies.
// A string containing an ISO 4217 currency (http://en.wikipedia.org/wiki/ISO_4217)
type Currency string

// A tag associated with an entity.
type Tag struct {
	Type string `json:"type"` // The type of the tag.
	Name string `json:"name"` // The name of the tag.
}

// A string containing the base currency and quote currency delimited by a “_”.
type InstrumentName string

// InstrumentType	The type of an Instrument.
type InstrumentType string

const (
	InstrumentType_CURRENCY InstrumentType = "CURRENCY" // Currency
	InstrumentType_CFD                     = "CFD"      // Contract For Difference
	InstrumentType_METAL                   = "METAL"    // Metal
)

// The DayOfWeek provides a representation of the day of the week.
type DayOfWeek string

const (
	DayOfWeek_SUNDAY    DayOfWeek = "SUNDAY"    // Sunday
	DayOfWeek_MONDAY    DayOfWeek = "MONDAY"    // Monday
	DayOfWeek_TUESDAY   DayOfWeek = "TUESDAY"   // Tuesday
	DayOfWeek_WEDNESDAY DayOfWeek = "WEDNESDAY" // Wednesday
	DayOfWeek_THURSDAY  DayOfWeek = "THURSDAY"  // Thursday
	DayOfWeek_FRIDAY    DayOfWeek = "FRIDAY"    // Friday
	DayOfWeek_SATURDAY  DayOfWeek = "SATURDAY"  // Sunday
)

// DateTime header
type AcceptDatetimeFormat string

const (
	// If “UNIX” is specified DateTime fields will be specified or returned
	// in the “12345678.000000123” format.
	AcceptDatetimeFormat_UNIX AcceptDatetimeFormat = "UNIX"
	// If “RFC3339” is specified DateTime will be specified or returned in
	// “YYYY-MM-DDTHH:MM:SS.nnnnnnnnnZ” format.
	AcceptDatetimeFormat_RFC3339 AcceptDatetimeFormat = "RFC3339"
)

// A FinancingDayOfWeek message defines a day of the week when financing charges are debited or credited.
type FinancingDayOfWeek struct {
	// The day of the week to charge the financing.
	DayOfWeek DayOfWeek `json:"dayOfWeek"`
	// The number of days worth of financing to be charged on dayOfWeek.
	DaysCharged int64 `json:"daysCharged"`
}

// Financing data for the instrument.
type InstrumentFinancing struct {
	// The financing rate to be used for a long position for the instrument. The
	// value is in decimal rather than percentage points, i.e. 5% is represented
	// as 0.05.
	LongRate DecimalNumber `json:"longRate"`
	// The financing rate to be used for a short position for the instrument.
	// The value is in decimal rather than percentage points, i.e. 5% is
	// represented as 0.05.
	ShortRate DecimalNumber `json:"shortRate"`
	// The days of the week to debit or credit financing charges; the exact time
	// of day at which to charge the financing is set in the
	// DivisionTradingGroup for the client’s account.
	FinancingDaysOfWeek []FinancingDayOfWeek `json:"financingDaysOfWeek"`
}

type Instrument struct {
	// The name of the Instrument
	Name InstrumentName `json:"name"`
	// The type of the Instrument
	Type InstrumentType `json:"type"`
	// The display name of the Instrument
	DisplayName string `json:"displayName"`
	// The location of the “pip” for this instrument. The decimal position of
	// the pip in this Instrument’s price can be found at 10 ^ pipLocation (e.g.
	// -4 pipLocation results in a decimal pip position of 10 ^ -4 = 0.0001).
	PipLocation int64 `json:"pipLocation"`
	// The number of decimal places that should be used to display prices for
	// this instrument. (e.g. a displayPrecision of 5 would result in a price of
	// “1” being displayed as “1.00000”)
	DisplayPrecision int64 `json:"displayPrecision"`
	// The amount of decimal places that may be provided when specifying the
	// number of units traded for this instrument.
	TradeUnitsPrecision int64 `json:"tradeUnitsPrecision"`
	// The smallest number of units allowed to be traded for this instrument.
	MinimumTradeSize DecimalNumber `json:"minimumTradeSize"`
	// The maximum trailing stop distance allowed for a trailing stop loss
	// created for this instrument. Specified in price units.
	MaximumTrailingStopDistance DecimalNumber `json:"maximumTrailingStopDistance"`
	// The minimum distance allowed between the Trade’s fill price and the
	// configured price for guaranteed Stop Loss Orders created for this
	// instrument. Specified in price units.
	MinimumGuaranteedStopLossDistance DecimalNumber `json:"minimumGuaranteedStopLossDistance"`
	// The minimum trailing stop distance allowed for a trailing stop loss
	// created for this instrument. Specified in price units.
	MinimumTrailingStopDistance DecimalNumber `json:"minimumTrailingStopDistance"`
	// The maximum position size allowed for this instrument. Specified in units.
	MaximumPositionSize DecimalNumber `json:"maximumPositionSize"`
	// The maximum units allowed for an Order placed for this instrument.
	// Specified in units.
	MaximumOrderUnits DecimalNumber `json:"maximumOrderUnits"`
	// The margin rate for this instrument.
	MarginRate DecimalNumber `json:"marginRate"`
	// The commission structure for this instrument.
	Commission DecimalNumber `json:"commission"`
	// The current Guaranteed Stop Loss Order mode of the Account for this Instrument.
	GuaranteedStopLossOrderMode DecimalNumber `json:"guaranteedStopLossOrderMode"`
	// The amount that is charged to the account if a guaranteed Stop Loss Order
	// is triggered and filled. The value is in price units and is charged for
	// each unit of the Trade. This field will only be present if the Account’s
	// guaranteedStopLossOrderMode for this Instrument is not ‘DISABLED’.
	GuaranteedStopLossOrderExecutionPremium DecimalNumber `json:"guaranteedStopLossOrderExecutionPremium"`
	// The guaranteed Stop Loss Order level restriction for this instrument.
	// This field will only be present if the Account’s
	// guaranteedStopLossOrderMode for this Instrument is not ‘DISABLED’.
	GuaranteedStopLossOrderLevelRestriction *GuaranteedStopLossOrderLevelRestriction `json:"guaranteedStopLossOrderLevelRestriction"`
	// Financing data for this instrument.
	Financing InstrumentFinancing `json:"financing"`
	// The tags associated with this instrument.
	Tags []Tag `json:"tags"`
}

// InstrumentCommission represents an instrument-specific commission
type InstrumentCommission struct {
	// The commission amount (in the Account’s home currency) charged per
	// unitsTraded of the instrument
	Commission DecimalNumber `json:"commission"`
	// The number of units traded that the commission amount is based on.
	UnitsTraded DecimalNumber `json:"unitsTraded"`
	// The minimum commission amount (in the Account’s home currency) that is
	// charged when an Order is filled for this instrument.
	MinimumCommission DecimalNumber `json:"minimumCommission"`
}

// The overall behaviour of the Account regarding Guaranteed Stop Loss Orders for a specific Instrument.
type GuaranteedStopLossOrderModeForInstrument string

const (
	// The Account is not permitted to create Guaranteed Stop Loss Orders for this Instrument.
	GuaranteedStopLossOrderModeForInstrument_DISABLED GuaranteedStopLossOrderModeForInstrument = "DISABLED"
	// 	The Account is able, but not required to have Guaranteed Stop Loss Orders for open Trades for this Instrument.
	GuaranteedStopLossOrderModeForInstrument_ALLOWED = "ALLOWED"
	// The Account is required to have Guaranteed Stop Loss Orders for all open Trades for this Instrument.
	GuaranteedStopLossOrderModeForInstrument_REQUIRED = "REQUIRED"
)

// A GuaranteedStopLossOrderLevelRestriction represents the total position size that can exist within a
// given price window for Trades with guaranteed Stop Loss Orders attached for a specific Instrument.
type GuaranteedStopLossOrderLevelRestriction struct {
	// Applies to Trades with a guaranteed Stop Loss Order attached for the
	// specified Instrument. This is the total allowed Trade volume that can
	// exist within the priceRange based on the trigger prices of the guaranteed
	// Stop Loss Orders.
	Volume DecimalNumber `json:"volume"`
	// The price range the volume applies to. This value is in price units.
	PriceRange DecimalNumber `json:"priceRange"`
}

// Direction	In the context of an Order or a Trade, defines whether the units are positive or negative.
type Direction string

const (
	// A long Order is used to to buy units of an Instrument. A Trade is long when it has
	// bought units of an Instrument.
	Direction_LONG Direction = "LONG"
	// A short Order is used to to sell units of an Instrument. A Trade is short when it has
	// sold units of an Instrument.
	Direction_SHORT = "SHORT"
)

// The Price component(s) to get candlestick data for.
type PricingComponent string

const (
	PricingComponent_BID         PricingComponent = "B"   // Bid candles
	PricingComponent_ASK         PricingComponent = "A"   // Ask candles
	PricingComponent_MID         PricingComponent = "M"   // Midpoint candles
	PricingComponent_BID_ASK     PricingComponent = "BA"  // Bid/Ask candles
	PricingComponent_BID_ASK_MID PricingComponent = "BAM" // Bid/Ask/Midpoint candles
)

// A ConversionFactor contains information used to convert an amount, from an Instrument’s base or
// quote currency, to the home currency of an Account.
type ConversionFactor struct {
	// The factor by which to multiply the amount in the given currency to
	// obtain the amount in the home currency of the Account.
	Factor DecimalNumber `json:"factor"`
}

// A HomeConversionFactors message contains information used to convert amounts, from an
// Instrument’s base or quote currency, to the home currency of an Account.
type HomeConversionFactors struct {
	// The ConversionFactor in effect for the Account for converting any gains
	// realized in Instrument quote units into units of the Account’s home currency.
	GainQuoteHome ConversionFactor `json:"gainQuoteHome"`
	// The ConversionFactor in effect for the Account for converting any losses
	// realized in Instrument quote units into units of the Account’s home currency.
	LossQuoteHome ConversionFactor `json:"lossQuoteHome"`
	// The ConversionFactor in effect for the Account for converting any gains
	// realized in Instrument base units into units of the Account’s home currency.
	GainBaseHome ConversionFactor `json:"gainBaseHome"`
	// The ConversionFactor in effect for the Account for converting any losses
	// realized in Instrument base units into units of the Account’s home currency.
	LossBaseHome ConversionFactor `json:"lossBaseHome"`
}
