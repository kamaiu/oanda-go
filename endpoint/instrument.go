package endpoint

import (
	"errors"
	. "github.com/kamaiu/oanda-go/model"
	"github.com/valyala/bytebufferpool"
	"time"
)

// GET  /v3/instruments/{instrument}/candles
// Fetch candlestick data for an instrument.
func (c *Connection) InstrumentCandles(
	request *InstrumentCandlesRequest,
) (*CandlestickResponse, error) {
	if request == nil || len(request.Instrument) == 0 {
		return nil, errors.New("instrument name required")
	}
	resp := &CandlestickResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/instruments/")
	_, _ = url.WriteString((string)(request.Instrument))
	_, _ = url.WriteString("/candles?")
	request.AppendQuery(url)

	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/instruments/{instrument}/orderBook
// Fetch an order book for an instrument.
func (c *Connection) InstrumentOrderBook(
	// Name of the Instrument [required]
	instrument InstrumentName,
	// The time of the snapshot to fetch. If not specified, then the most recent snapshot is fetched.
	t time.Time,
) (*OrderBook, error) {
	if len(instrument) == 0 {
		return nil, errors.New("instrument name required")
	}
	if t.IsZero() {
		t = time.Now()
	}
	resp := &OrderBookResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/instruments/")
	_, _ = url.WriteString((string)(instrument))
	_, _ = url.WriteString("/orderBook?time=")
	_, _ = url.WriteString(t.Format(time.RFC3339))

	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp.OrderBook, nil
}

// GET /v3/instruments/{instrument}/positionBook
// Fetch a position book for an instrument.
func (c *Connection) InstrumentPositionBook(
	// Name of the Instrument [required]
	instrument InstrumentName,
	// The time of the snapshot to fetch. If not specified, then the most recent snapshot is fetched.
	t time.Time,
) (*PositionBook, error) {
	if len(instrument) == 0 {
		return nil, errors.New("instrument name required")
	}
	if t.IsZero() {
		t = time.Now()
	}
	resp := &PositionBookResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/instruments/")
	_, _ = url.WriteString((string)(instrument))
	_, _ = url.WriteString("/positionBook?time=")
	_, _ = url.WriteString(t.Format(time.RFC3339))

	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp.PositionBook, nil
}
