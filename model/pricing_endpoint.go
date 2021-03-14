//go:generate easyjson -all $GOFILE
package model

import (
	"github.com/valyala/bytebufferpool"
	"net/url"
	"strconv"
	"time"
)

type CandlesLatestRequest struct {
	// List of candle specifications to get pricing for.
	// [required]
	CandleSpecifications []CandleSpecification `json:"candleSpecifications"`
	// The number of units used to calculate the volume-weighted average
	// bid and ask prices in the returned candles.
	// [default=1]
	Units DecimalNumber `json:"units"`
	// A flag that controls whether the candlestick is “smoothed” or not.
	// A smoothed candlestick uses the previous candle’s close price as its
	// open price, while an unsmoothed candlestick uses the first price from
	// its time range as its open price.
	// [default=False]
	Smooth bool `json:"smooth"`
	// The hour of the day (in the specified timezone) to use for granularities
	// that have daily alignments.
	// [default=17, minimum=0, maximum=23]
	DailyAlignment int `json:"dailyAlignment"`
	// The timezone to use for the dailyAlignment parameter. Candlesticks with
	// daily alignment will be aligned to the dailyAlignment hour within the
	// alignmentTimezone.
	// Note that the returned times will still be represented in UTC.
	// [default=America/New_York]
	AlignmentTimezone string `json:"alignmentTimezone"`
	// The day of the week used for granularities that have weekly alignment.
	// [default=Friday]
	WeeklyAlignment WeeklyAlignment `json:"weeklyAlignment"`
}

func NewCandlesLatestRequest(specs ...CandleSpecification) *CandlesLatestRequest {
	return &CandlesLatestRequest{
		CandleSpecifications: specs,
		Units:                "1",
		Smooth:               false,
		DailyAlignment:       17,
		AlignmentTimezone:    "America/New_York",
		WeeklyAlignment:      WeeklyAlignment_Friday,
	}
}

// List of candle specifications to get pricing for.
// [required]
func (c *CandlesLatestRequest) AddCandleSpecification(specification CandleSpecification) *CandlesLatestRequest {
	c.CandleSpecifications = append(c.CandleSpecifications, specification)
	return c
}

// The number of units used to calculate the volume-weighted average
// bid and ask prices in the returned candles.
// [default=1]
func (c *CandlesLatestRequest) WithUnits(units DecimalNumber) *CandlesLatestRequest {
	c.Units = units
	return c
}

// A flag that controls whether the candlestick is “smoothed” or not.
// A smoothed candlestick uses the previous candle’s close price as its
// open price, while an unsmoothed candlestick uses the first price from
// its time range as its open price.
// [default=False]
func (c *CandlesLatestRequest) WithSmooth(smooth bool) *CandlesLatestRequest {
	c.Smooth = smooth
	return c
}

// The hour of the day (in the specified timezone) to use for granularities
// that have daily alignments.
// [default=17, minimum=0, maximum=23]
func (c *CandlesLatestRequest) WithDailyAlignment(dailyAlignment int) *CandlesLatestRequest {
	if dailyAlignment < 0 {
		c.DailyAlignment = 17
	} else if dailyAlignment > 23 {
		c.DailyAlignment = 23
	} else {
		c.DailyAlignment = dailyAlignment
	}
	return c
}

// The timezone to use for the dailyAlignment parameter. Candlesticks with
// daily alignment will be aligned to the dailyAlignment hour within the
// alignmentTimezone.
// Note that the returned times will still be represented in UTC.
// [default=America/New_York]
func (c *CandlesLatestRequest) WithAlignmentTimezone(timezone string) *CandlesLatestRequest {
	c.AlignmentTimezone = timezone
	return c
}

// The day of the week used for granularities that have weekly alignment.
// [default=Friday]
func (c *CandlesLatestRequest) WithWeeklyAlignment(alignment WeeklyAlignment) *CandlesLatestRequest {
	switch alignment {
	case WeeklyAlignment_Friday,
		WeeklyAlignment_Saturday,
		WeeklyAlignment_Sunday,
		WeeklyAlignment_Monday,
		WeeklyAlignment_Tuesday,
		WeeklyAlignment_Wednesday,
		WeeklyAlignment_Thursday:
		c.WeeklyAlignment = alignment
	default:
		c.WeeklyAlignment = WeeklyAlignment_Friday
	}
	return c
}

func (c *CandlesLatestRequest) AppendQuery(b *bytebufferpool.ByteBuffer) string {
	_, _ = b.WriteString("smooth=")
	_, _ = b.WriteString(strconv.FormatBool(c.Smooth))

	if len(c.Units) == 0 {
		c.Units = "1"
	}
	_, _ = b.WriteString("&units=")
	_, _ = b.WriteString((string)(c.Units))

	if c.DailyAlignment < 0 {
		c.DailyAlignment = 17
	} else if c.DailyAlignment > 23 {
		c.DailyAlignment = 23
	}
	_, _ = b.WriteString("&dailyAlignment=")
	_, _ = b.WriteString(strconv.Itoa(c.DailyAlignment))

	if len(c.AlignmentTimezone) > 0 {
		_, _ = b.WriteString("&alignmentTimezone=")
		_, _ = b.WriteString(url.PathEscape(c.AlignmentTimezone))
	}

	if len(c.WeeklyAlignment) == 0 {
		c.WeeklyAlignment = WeeklyAlignment_Friday
	}
	_, _ = b.WriteString("&weeklyAlignment=")
	_, _ = b.WriteString((string)(c.WeeklyAlignment))

	if len(c.CandleSpecifications) > 0 {
		_, _ = b.WriteString("&candleSpecifications=")
		for i, t := range c.CandleSpecifications {
			if i > 0 {
				_, _ = b.WriteString(",")
			}
			_, _ = b.WriteString((string)(t))
		}
	}

	return b.String()
}

type CandlesLatestResponse struct {
	// The latest candle sticks.
	LatestCandles []*CandlestickResponse `json:"latestCandles"`
}

type PricingRequest struct {
	// List of Instruments to get pricing for.
	// [required]
	Instruments []InstrumentName `json:"instruments"`
	// Date/Time filter to apply to the response. Only prices and home conversions
	// (if requested) with a time later than this filter (i.e. the price has changed
	// after the since time) will be provided, and are filtered independently.
	Since DateTime `json:"since"`
	// Flag that enables the inclusion of the unitsAvailable field in the returned
	// Price objects.
	// [default=True]
	// Deprecated: Will be removed in a future API update.
	IncludeUnitsAvailable bool `json:"includeUnitsAvailable"`
	// Flag that enables the inclusion of the homeConversions field in the returned response.
	// An entry will be returned for each currency in the set of all base and quote currencies
	// present in the requested instruments list.
	// [default=False]
	IncludeHomeConversions bool `json:"includeHomeConversions"`
}

// List of Instruments to get pricing for.
// [required]
func (p *PricingRequest) WithInstruments(instruments ...InstrumentName) *PricingRequest {
	p.Instruments = instruments
	return p
}

// Date/Time filter to apply to the response. Only prices and home conversions
// (if requested) with a time later than this filter (i.e. the price has changed
// after the since time) will be provided, and are filtered independently.
func (p *PricingRequest) WithSince(since time.Time) *PricingRequest {
	p.Since = DateTime(since.Format(time.RFC3339))
	return p
}

// Flag that enables the inclusion of the homeConversions field in the returned response.
// An entry will be returned for each currency in the set of all base and quote currencies
// present in the requested instruments list.
// [default=False]
func (p *PricingRequest) WithIncludeHomeConversions(includeHomeConversions bool) *PricingRequest {
	p.IncludeHomeConversions = includeHomeConversions
	return p
}

func NewPricingRequest() *PricingRequest {
	return &PricingRequest{}
}

func (p *PricingRequest) AppendQuery(b *bytebufferpool.ByteBuffer) {
	if len(p.Since) == 0 {
		p.WithSince(time.Now())
	}
	_, _ = b.WriteString("since=")
	_, _ = b.WriteString((string)(p.Since))
	_, _ = b.WriteString("includeHomeConversions=")
	_, _ = b.WriteString(strconv.FormatBool(p.IncludeHomeConversions))
	if len(p.Instruments) > 0 {
		for i, instrument := range p.Instruments {
			if i > 0 {
				_ = b.WriteByte(',')
			}
			_, _ = b.WriteString((string)(instrument))
		}
	}
}

type PricingResponse struct {
	// The list of Price objects requested.
	Prices []*ClientPrice `json:"prices"`
	// The list of home currency conversion factors requested. This field will
	// only be present if includeHomeConversions was set to true in the request.
	HomeConversions []*HomeConversions `json:"homeConversions"`
	// The DateTime value to use for the “since” parameter in the next poll request.
	Time DateTime `json:"time"`
}

type PricingCandlesRequest struct {
	// The Price component(s) to get candlestick data for.
	// [default=M]
	Price PricingComponent `json:"price"`
	// The granularity of the candlesticks to fetch.
	// [default=S5]
	Granularity CandlestickGranularity `json:"granularity"`
	// The number of candlesticks to return in the response. Count should
	// not be specified if both the start and end parameters are provided,
	// as the time range combined with the granularity will determine the
	// number of candlesticks to return. [default=500, maximum=5000]
	Count int `json:"count"`
	// The start of the time range to fetch candlesticks for.
	From DateTime `json:"from"`
	// The end of the time range to fetch candlesticks for.
	To DateTime `json:"to"`
	// A flag that controls whether the candlestick is “smoothed” or not.
	// A smoothed candlestick uses the previous candle’s close price as its
	// open price, while an un-smoothed candlestick uses the first price from
	// its time range as its open price.
	// [default=False]
	Smooth bool `json:"smooth"`
	// A flag that controls whether the candlestick that is covered by the from
	// time should be included in the results. This flag enables clients to use
	// the timestamp of the last completed candlestick received to poll for future
	// candlesticks but avoid receiving the previous candlestick repeatedly.
	// [default=True]
	IncludeFirst bool `json:"includeFirst"`
	// The hour of the day (in the specified timezone) to use for granularities
	// that have daily alignments.
	// [default=17, minimum=0, maximum=23]
	DailyAlignment int `json:"dailyAlignment"`
	// The timezone to use for the dailyAlignment parameter. Candlesticks with daily
	// alignment will be aligned to the dailyAlignment hour within the alignmentTimezone.
	// Note that the returned times will still be represented in UTC.
	// [default=America/New_York]
	AlignmentTimezone string `json:"alignmentTimezone"`
	// The day of the week used for granularities that have weekly alignment.
	// [default=Friday]
	WeeklyAlignment WeeklyAlignment `json:"weeklyAlignment"`
	// The number of units used to calculate the volume-weighted average bid and ask
	// prices in the returned candles.
	// [default=1]
	Units DecimalNumber `json:"units"`
}

// The Price component(s) to get candlestick data for.
// [default=M]
func (s *PricingCandlesRequest) WithPrice(price PricingComponent) *PricingCandlesRequest {
	if len(price) == 0 {
		price = PricingComponent_MID
	}
	s.Price = price
	return s
}

// The granularity of the candlesticks to fetch.
// [default=S5]
func (s *PricingCandlesRequest) WithGranularity(granularity CandlestickGranularity) *PricingCandlesRequest {
	if len(granularity) == 0 {
		granularity = CandlestickGranularity_S5
	}
	s.Granularity = granularity
	return s
}

// The number of candlesticks to return in the response. Count should
// not be specified if both the start and end parameters are provided,
// as the time range combined with the granularity will determine the
// number of candlesticks to return.
//[default=500, maximum=5000]
func (s *PricingCandlesRequest) WithCount(count int) *PricingCandlesRequest {
	if count < 1 {
		s.Count = 500
	} else if count > 5000 {
		s.Count = 5000
	} else {
		s.Count = count
	}
	return s
}

// The start of the time range to fetch candlesticks for.
func (s *PricingCandlesRequest) WithFrom(from time.Time) *PricingCandlesRequest {
	if from.IsZero() {
		s.From = ""
	} else {
		s.From = DateTime(from.Format(time.RFC3339))
	}
	return s
}

// The end of the time range to fetch candlesticks for.
func (s *PricingCandlesRequest) WithTo(to time.Time) *PricingCandlesRequest {
	if to.IsZero() {
		s.To = ""
	} else {
		s.To = DateTime(to.Format(time.RFC3339))
	}
	return s
}

// The start of the time range to fetch candlesticks for.
// The end of the time range to fetch candlesticks for.
func (s *PricingCandlesRequest) WithRange(from, to time.Time) *PricingCandlesRequest {
	return s.WithFrom(from).WithTo(to)
}

// A flag that controls whether the candlestick is “smoothed” or not.
// A smoothed candlestick uses the previous candle’s close price as its
// open price, while an un-smoothed candlestick uses the first price from
// its time range as its open price.
// [default=False]
func (s *PricingCandlesRequest) WithSmooth(v bool) *PricingCandlesRequest {
	s.Smooth = v
	return s
}

// A flag that controls whether the candlestick that is covered by the from
// time should be included in the results. This flag enables clients to use
// the timestamp of the last completed candlestick received to poll for future
// candlesticks but avoid receiving the previous candlestick repeatedly.
// [default=True]
func (s *PricingCandlesRequest) WithIncludeFirst(includeFirst bool) *PricingCandlesRequest {
	s.IncludeFirst = includeFirst
	return s
}

// The hour of the day (in the specified timezone) to use for granularities
// that have daily alignments.
// [default=17, minimum=0, maximum=23]
func (s *PricingCandlesRequest) WithDailyAlignment(dailyAlignment int) *PricingCandlesRequest {
	if dailyAlignment < 0 {
		s.DailyAlignment = 0
	} else if dailyAlignment > 23 {
		s.DailyAlignment = 23
	} else {
		s.DailyAlignment = dailyAlignment
	}
	return s
}

// The timezone to use for the dailyAlignment parameter. Candlesticks with daily
// alignment will be aligned to the dailyAlignment hour within the alignmentTimezone.
// Note that the returned times will still be represented in UTC.
// [default=America/New_York]
func (s *PricingCandlesRequest) WithAlignmentTimezone(timezone string) *PricingCandlesRequest {
	if len(timezone) == 0 {
		s.AlignmentTimezone = "America/New_York"
	} else {
		s.AlignmentTimezone = timezone
	}
	return s
}

// The day of the week used for granularities that have weekly alignment.
// [default=Friday]
func (s *PricingCandlesRequest) WithWeeklyAlignment(alignment WeeklyAlignment) *PricingCandlesRequest {
	switch alignment {
	case WeeklyAlignment_Friday,
		WeeklyAlignment_Saturday,
		WeeklyAlignment_Sunday,
		WeeklyAlignment_Monday,
		WeeklyAlignment_Tuesday,
		WeeklyAlignment_Wednesday,
		WeeklyAlignment_Thursday:
		s.WeeklyAlignment = alignment
	default:
		s.WeeklyAlignment = WeeklyAlignment_Friday
	}
	return s
}

// The number of units used to calculate the volume-weighted average bid and ask
// prices in the returned candles.
// [default=1]
func (s *PricingCandlesRequest) WithUnits(units DecimalNumber) *PricingCandlesRequest {
	if len(units) == 0 {
		s.Units = "1"
	} else {
		s.Units = units
	}
	return s
}

func NewAccountInstrumentCandlesRequest(instrument InstrumentName, from time.Time) *PricingCandlesRequest {
	if from.IsZero() {
		from = time.Now().UTC()
	}
	r := &PricingCandlesRequest{
		Price:             PricingComponent_MID,
		Granularity:       CandlestickGranularity_S5,
		Count:             500,
		Smooth:            false,
		IncludeFirst:      true,
		DailyAlignment:    17,
		AlignmentTimezone: "America/New_York",
		WeeklyAlignment:   WeeklyAlignment_Friday,
		Units:             "1",
	}
	return r.WithFrom(from)
}

func (s *PricingCandlesRequest) AppendQuery(b *bytebufferpool.ByteBuffer) string {
	// Price
	_, _ = b.WriteString("price=")
	_, _ = b.WriteString((string)(s.Price))

	_, _ = b.WriteString("&includeFirst=")
	_, _ = b.WriteString(strconv.FormatBool(s.IncludeFirst))

	_, _ = b.WriteString("&smooth=")
	_, _ = b.WriteString(strconv.FormatBool(s.Smooth))

	_, _ = b.WriteString("&count=")
	_, _ = b.WriteString(strconv.Itoa(s.Count))

	if len(s.Granularity) == 0 {
		s.Granularity = CandlestickGranularity_S5
	}
	_, _ = b.WriteString("&granularity=")
	_, _ = b.WriteString((string)(s.Granularity))

	if s.DailyAlignment < 0 {
		s.DailyAlignment = 0
	} else if s.DailyAlignment > 23 {
		s.DailyAlignment = 23
	}
	_, _ = b.WriteString("&dailyAlignment=")
	_, _ = b.WriteString(strconv.Itoa(s.DailyAlignment))

	if len(s.AlignmentTimezone) == 0 {
		s.AlignmentTimezone = "America/New_York"
	}
	_, _ = b.WriteString("&alignmentTimezone=")
	_, _ = b.WriteString(url.PathEscape(s.AlignmentTimezone))

	if len(s.WeeklyAlignment) == 0 {
		s.WeeklyAlignment = WeeklyAlignment_Friday
	}
	_, _ = b.WriteString("&weeklyAlignment=")
	_, _ = b.WriteString((string)(s.WeeklyAlignment))

	if len(s.From) > 0 {
		_, _ = b.WriteString("&from=")
		_, _ = b.WriteString(url.PathEscape((string)(s.From)))
	}

	if len(s.To) > 0 {
		_, _ = b.WriteString("&to=")
		_, _ = b.WriteString(url.PathEscape((string)(s.To)))
	}

	if len(s.Units) > 0 {
		_, _ = b.WriteString("&units=")
		_, _ = b.WriteString((string)(s.Units))
	}

	return b.String()
}

type PricingCandlesResponse struct {
	// The instrument whose Prices are represented by the candlesticks.
	Instrument InstrumentName `json:"instrument"`
	// The granularity of the candlesticks provided.
	Granularity CandlestickGranularity `json:"granularity"`
	// The list of candlesticks that satisfy the request.
	Candles []*Candlestick `json:"candles"`
}
