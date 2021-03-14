//go:generate easyjson -all $GOFILE
package model

import (
	"github.com/valyala/bytebufferpool"
	"net/url"
	"strconv"
	"time"
)

type InstrumentCandlesRequest struct {
	// Name of the Instrument.
	// [required]
	Instrument InstrumentName `json:"instrument"`
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
}

// Name of the Instrument.
// [required]
func (s *InstrumentCandlesRequest) WithInstrument(instrument InstrumentName) *InstrumentCandlesRequest {
	s.Instrument = instrument
	return s
}

// The Price component(s) to get candlestick data for.
// [default=M]
func (s *InstrumentCandlesRequest) WithPrice(price PricingComponent) *InstrumentCandlesRequest {
	if len(price) == 0 {
		price = PricingComponent_MID
	}
	s.Price = price
	return s
}

// The granularity of the candlesticks to fetch.
// [default=S5]
func (s *InstrumentCandlesRequest) WithGranularity(granularity CandlestickGranularity) *InstrumentCandlesRequest {
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
func (s *InstrumentCandlesRequest) WithCount(count int) *InstrumentCandlesRequest {
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
func (s *InstrumentCandlesRequest) WithFrom(from time.Time) *InstrumentCandlesRequest {
	if from.IsZero() {
		s.From = ""
	} else {
		s.From = DateTime(from.Format(time.RFC3339))
	}
	return s
}

// The end of the time range to fetch candlesticks for.
func (s *InstrumentCandlesRequest) WithTo(to time.Time) *InstrumentCandlesRequest {
	if to.IsZero() {
		s.To = ""
	} else {
		s.To = DateTime(to.Format(time.RFC3339))
	}
	return s
}

// The start of the time range to fetch candlesticks for.
// The end of the time range to fetch candlesticks for.
func (s *InstrumentCandlesRequest) WithRange(from, to time.Time) *InstrumentCandlesRequest {
	return s.WithFrom(from).WithTo(to)
}

// A flag that controls whether the candlestick is “smoothed” or not.
// A smoothed candlestick uses the previous candle’s close price as its
// open price, while an un-smoothed candlestick uses the first price from
// its time range as its open price.
// [default=False]
func (s *InstrumentCandlesRequest) WithSmooth(v bool) *InstrumentCandlesRequest {
	s.Smooth = v
	return s
}

// A flag that controls whether the candlestick that is covered by the from
// time should be included in the results. This flag enables clients to use
// the timestamp of the last completed candlestick received to poll for future
// candlesticks but avoid receiving the previous candlestick repeatedly.
// [default=True]
func (s *InstrumentCandlesRequest) WithIncludeFirst(v bool) *InstrumentCandlesRequest {
	s.IncludeFirst = v
	return s
}

// The hour of the day (in the specified timezone) to use for granularities
// that have daily alignments.
// [default=17, minimum=0, maximum=23]
func (s *InstrumentCandlesRequest) WithDailyAlignment(dailyAlignment int) *InstrumentCandlesRequest {
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
func (s *InstrumentCandlesRequest) WithAlignmentTimezone(timezone string) *InstrumentCandlesRequest {
	if len(timezone) == 0 {
		s.AlignmentTimezone = "America/New_York"
	} else {
		s.AlignmentTimezone = timezone
	}
	return s
}

// The day of the week used for granularities that have weekly alignment.
// [default=Friday]
func (s *InstrumentCandlesRequest) WithWeeklyAlignment(alignment WeeklyAlignment) *InstrumentCandlesRequest {
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

func NewInstrumentCandlesRequest(instrument InstrumentName, from time.Time) *InstrumentCandlesRequest {
	if from.IsZero() {
		from = time.Now().UTC()
	}
	r := &InstrumentCandlesRequest{
		Instrument:        instrument,
		Price:             PricingComponent_MID,
		Granularity:       CandlestickGranularity_S5,
		Count:             500,
		Smooth:            false,
		IncludeFirst:      true,
		DailyAlignment:    17,
		AlignmentTimezone: "America/New_York",
		WeeklyAlignment:   WeeklyAlignment_Friday,
	}
	return r.WithFrom(from)
}

func (s *InstrumentCandlesRequest) AppendQuery(b *bytebufferpool.ByteBuffer) string {
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

	return b.String()
}

type OrderBookResponse struct {
	OrderBook *OrderBook `json:"orderBook"`
}

type PositionBookResponse struct {
	PositionBook *PositionBook `json:"positionBook"`
}
