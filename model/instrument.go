//go:generate easyjson -all $GOFILE
package model

// Supporting OANDA docs - http://developer.com/rest-live-v20/instrument-ep/

// The granularity of the candlestick
type CandlestickGranularity string

const (
	CandlestickGranularity_S5  CandlestickGranularity = "S5"  // 5 second candlesticks, minute alignment
	CandlestickGranularity_S10 CandlestickGranularity = "S10" // 10 second candlesticks, minute alignment
	CandlestickGranularity_S15 CandlestickGranularity = "S15" // 15 second candlesticks, minute alignment
	CandlestickGranularity_S30 CandlestickGranularity = "S30" // 30 second candlesticks, minute alignment
	CandlestickGranularity_M1  CandlestickGranularity = "M1"  // 1 minute candlesticks, minute alignment
	CandlestickGranularity_M2  CandlestickGranularity = "M2"  // 2 minute candlesticks, hour alignment
	CandlestickGranularity_M4  CandlestickGranularity = "M4"  // 4 minute candlesticks, hour alignment
	CandlestickGranularity_M5  CandlestickGranularity = "M5"  // 5 minute candlesticks, hour alignment
	CandlestickGranularity_M10 CandlestickGranularity = "M10" // 10 minute candlesticks, hour alignment
	CandlestickGranularity_M15 CandlestickGranularity = "M15" // 15 minute candlesticks, hour alignment
	CandlestickGranularity_M30 CandlestickGranularity = "M30" // 30 minute candlesticks, hour alignment
	CandlestickGranularity_H1  CandlestickGranularity = "H1"  // 1 hour candlesticks, hour alignment
	CandlestickGranularity_H2  CandlestickGranularity = "H2"  // 2 hour candlesticks, day alignment
	CandlestickGranularity_H3  CandlestickGranularity = "H3"  // 3 hour candlesticks, day alignment
	CandlestickGranularity_H4  CandlestickGranularity = "H4"  // 4 hour candlesticks, day alignment
	CandlestickGranularity_H6  CandlestickGranularity = "H6"  // 6 hour candlesticks, day alignment
	CandlestickGranularity_H8  CandlestickGranularity = "H8"  // 8 hour candlesticks, day alignment
	CandlestickGranularity_H12 CandlestickGranularity = "H12" // 12 hour candlesticks, day alignment
	CandlestickGranularity_D   CandlestickGranularity = "D"   // 1 day candlesticks, day alignment
	CandlestickGranularity_W   CandlestickGranularity = "W"   // 1 week candlesticks, aligned to start of week
	CandlestickGranularity_M   CandlestickGranularity = "M"   // 1 month candlesticks, aligned to first day of the month
)

// The day of the week to use for candlestick granularity with weekly alignment.
type WeeklyAlignment string

const (
	WeeklyAlignment_Monday    WeeklyAlignment = "Monday"
	WeeklyAlignment_Tuesday   WeeklyAlignment = "Tuesday"
	WeeklyAlignment_Wednesday WeeklyAlignment = "Wednesday"
	WeeklyAlignment_Thursday  WeeklyAlignment = "Thursday"
	WeeklyAlignment_Friday    WeeklyAlignment = "Friday"
	WeeklyAlignment_Saturday  WeeklyAlignment = "Saturday"
	WeeklyAlignment_Sunday    WeeklyAlignment = "Sunday"
)

// The candlestick representation
type Candlestick struct {
	// The start time of the candlestick.
	Time DateTime `json:"time"`
	// The candlestick data based on bids. Only provided if bid-based candles were requested.
	Bid *CandlestickData `json:"bid"`
	// The candlestick data based on asks. Only provided if ask-based candles were requested.
	Ask *CandlestickData `json:"ask"`
	// The candlestick data based on midpoints. Only provided if midpoint-based candles were requested.
	Mid *CandlestickData `json:"mid"`
	// The number of prices created during the time-range represented by the candlestick.
	Volume int64 `json:"volume"`
	// A flag indicating if the candlestick is complete. A complete candlestick
	// is one whose ending time is not in the future.
	Complete bool `json:"complete"`
}

// The price data (open, high, low, close) for the Candlestick representation.
type CandlestickData struct {
	// The first (open) price in the time-range represented by the candlestick.
	Open PriceValue `json:"o"`
	//  The highest price in the time-range represented by the candlestick.
	High PriceValue `json:"h"`
	// The lowest price in the time-range represented by the candlestick.
	Low PriceValue `json:"l"`
	// The last (closing) price in the time-range represented by the candlestick.
	Close PriceValue `json:"c"`
}

// Response containing instrument, granularity, and list of candles.
type CandlestickResponse struct {
	// The instrument whose Prices are represented by the candlesticks.
	Instrument InstrumentName `json:"instrument"`
	// The granularity of the candlesticks provided.
	Granularity CandlestickGranularity `json:"granularity"`
	// The list of candlesticks that satisfy the request.
	Candles []*Candlestick `json:"candles"`
}

// The representation of an instrument’s order book at a point in time
type OrderBook struct {
	// The order book’s instrument.
	Instrument InstrumentName `json:"instrument"`
	// The time when the order book snapshot was created.
	Time DateTime `json:"time"`
	// The price (midpoint) for the order book’s instrument at the time of the
	// order book snapshot
	Price PriceValue `json:"price"`
	// The price width for each bucket. Each bucket covers the price range from
	// the bucket’s price to the bucket’s price + bucketWidth.
	BucketWidth PriceValue `json:"bucketWidth"`
	// The partitioned order book, divided into buckets using a default bucket
	// width. These buckets are only provided for price ranges which actually
	// contain order or position data.
	Buckets []*OrderBookBucket `json:"buckets"`
}

// The order book data for a partition of the instrument’s prices.
type OrderBookBucket struct {
	// The lowest price (inclusive) covered by the bucket. The bucket covers the
	// price range from the price to price + the order book’s bucketWidth.
	Price PriceValue `json:"price"`
	// The percentage of the total number of orders represented by the long
	// orders found in this bucket.
	LongCountPercent DecimalNumber `json:"longCountPercent"`
	// The percentage of the total number of orders represented by the short
	// orders found in this bucket.
	ShortCountPercent DecimalNumber `json:"shortCountPercent"`
}

// The representation of an instrument’s position book at a point in time
type PositionBook struct {
	// The position book’s instrument.
	Instrument InstrumentName `json:"instrument"`
	// The time when the position book snapshot was created.
	Time DateTime `json:"time"`
	// The price (midpoint) for the position book’s instrument at the time of the
	// position book snapshot
	Price PriceValue `json:"price"`
	// The price width for each bucket. Each bucket covers the price range from
	// the bucket’s price to the bucket’s price + bucketWidth.
	BucketWidth PriceValue `json:"bucketWidth"`
	// The partitioned position book, divided into buckets using a default bucket
	// width. These buckets are only provided for price ranges which actually
	// contain order or position data.
	Buckets []*OrderBookBucket `json:"buckets"`
}

// The position book data for a partition of the instrument’s prices.
type PositionBookBucket struct {
	// The lowest price (inclusive) covered by the bucket. The bucket covers the
	// price range from the price to price + the position book’s bucketWidth.
	Price PriceValue `json:"price"`
	// The percentage of the total number of positions represented by the long
	// orders found in this bucket.
	LongCountPercent DecimalNumber `json:"longCountPercent"`
	// The percentage of the total number of positions represented by the short
	// orders found in this bucket.
	ShortCountPercent DecimalNumber `json:"shortCountPercent"`
}
