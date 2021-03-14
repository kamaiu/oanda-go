package endpoint

import (
	. "github.com/kamaiu/oanda-go/model"
	"github.com/valyala/bytebufferpool"
)

// GET /v3/accounts/{accountID}/candles/latest
// Get dancing bears and most recently completed candles within an Account
// for specified combinations of instrument, granularity, and price component.
func (c *Connection) CandlesLatest(
	accountID AccountID,
	request *CandlesLatestRequest,
) (*CandlesLatestResponse, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	resp := &CandlesLatestResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.hostname)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/candles/latest?")
	request.AppendQuery(url)

	_, err := doGET(c, url, c.headers.DateFormat, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/accounts/{accountID}/pricing
// Get pricing information for a specified list of Instruments within an Account.
func (c *Connection) Pricing(
	accountID AccountID,
	request *PricingRequest,
) (*PricingResponse, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	resp := &PricingResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.hostname)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/pricing?")
	request.AppendQuery(url)

	_, err := doGET(c, url, c.headers.DateFormat, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/accounts/{accountID}/instruments/{instrument}/candles
// Fetch candlestick data for an instrument.
func (c *Connection) PricingCandles(
	accountID AccountID,
	instrument InstrumentName,
	request *PricingCandlesRequest,
) (*PricingCandlesResponse, error) {
	if request == nil {
		return nil, ErrNilRequest
	}
	resp := &PricingCandlesResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.hostname)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/instruments/")
	_, _ = url.WriteString((string)(instrument))
	_, _ = url.WriteString("/candles?")
	request.AppendQuery(url)

	_, err := doGET(c, url, c.headers.DateFormat, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
