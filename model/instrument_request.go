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
	From time.Time `json:"from"`
	// The end of the time range to fetch candlesticks for.
	To time.Time `json:"to"`
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

func (i *InstrumentCandlesRequest) WithPrice(v PricingComponent) *InstrumentCandlesRequest {
	i.Price = v
	return i
}
func (i *InstrumentCandlesRequest) WithGranularity(v CandlestickGranularity) *InstrumentCandlesRequest {
	i.Granularity = v
	return i
}
func (i *InstrumentCandlesRequest) WithCount(v int) *InstrumentCandlesRequest {
	i.Count = v
	return i
}
func (i *InstrumentCandlesRequest) WithFrom(v time.Time) *InstrumentCandlesRequest {
	i.From = v
	return i
}
func (i *InstrumentCandlesRequest) WithTo(v time.Time) *InstrumentCandlesRequest {
	i.To = v
	return i
}
func (i *InstrumentCandlesRequest) WithSmooth(v bool) *InstrumentCandlesRequest {
	i.Smooth = v
	return i
}
func (i *InstrumentCandlesRequest) WithIncludeFirst(v bool) *InstrumentCandlesRequest {
	i.IncludeFirst = v
	return i
}
func (i *InstrumentCandlesRequest) WithDailyAlignment(v int) *InstrumentCandlesRequest {
	if v < 0 {
		v = 0
	} else if v > 23 {
		v = 23
	}
	i.DailyAlignment = v
	return i
}
func (i *InstrumentCandlesRequest) WithAlignmentTimezone(v string) *InstrumentCandlesRequest {
	i.AlignmentTimezone = v
	return i
}
func (i *InstrumentCandlesRequest) WithWeeklyAlignment(v string) *InstrumentCandlesRequest {
	i.AlignmentTimezone = v
	return i
}

func NewInstrumentCandlesRequest(instrument InstrumentName, from time.Time) *InstrumentCandlesRequest {
	if from.IsZero() {
		from = time.Now().UTC()
	}
	return &InstrumentCandlesRequest{
		Instrument:        instrument,
		Price:             PricingComponent_MID,
		Granularity:       CandlestickGranularity_S5,
		Count:             500,
		Smooth:            false,
		IncludeFirst:      true,
		From:              from,
		DailyAlignment:    17,
		AlignmentTimezone: "America/New_York",
		WeeklyAlignment:   WeeklyAlignment_Friday,
	}
}
func (i *InstrumentCandlesRequest) AppendQuery(b *bytebufferpool.ByteBuffer) string {
	// Price
	_, _ = b.WriteString("price=")
	_, _ = b.WriteString((string)(i.Price))

	_, _ = b.WriteString("&includeFirst=")
	_, _ = b.WriteString(strconv.FormatBool(i.IncludeFirst))

	_, _ = b.WriteString("&smooth=")
	_, _ = b.WriteString(strconv.FormatBool(i.Smooth))

	_, _ = b.WriteString("&count=")
	_, _ = b.WriteString(strconv.Itoa(i.Count))

	if len(i.Granularity) == 0 {
		i.Granularity = CandlestickGranularity_S5
	}
	_, _ = b.WriteString("&granularity=")
	_, _ = b.WriteString((string)(i.Granularity))

	_, _ = b.WriteString("&dailyAlignment=")
	_, _ = b.WriteString(strconv.Itoa(i.DailyAlignment))

	_, _ = b.WriteString("&alignmentTimezone=")
	_, _ = b.WriteString(url.PathEscape(i.AlignmentTimezone))

	_, _ = b.WriteString("&weeklyAlignment=")
	_, _ = b.WriteString((string)(i.WeeklyAlignment))

	if i.From.IsZero() {
		i.From = time.Now()
	}
	_, _ = b.WriteString("&from=")
	_, _ = b.WriteString(i.From.Format(time.RFC3339))

	if !i.To.IsZero() {
		_, _ = b.WriteString("&to=")
		_, _ = b.WriteString(i.To.Format(time.RFC3339))
	}

	return b.String()
}

type OrderBookResponse struct {
	OrderBook *OrderBook `json:"orderBook"`
}

type PositionBookResponse struct {
	PositionBook *PositionBook `json:"positionBook"`
}
