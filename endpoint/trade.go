package endpoint

import (
	. "github.com/kamaiu/oanda-go/model"
	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
)

func (c *Connection) Trades(
	accountID AccountID,
	request *TradesRequest,
) (*TradesResponse, error) {
	if request == nil {
		request = &TradesRequest{}
	}
	resp := &TradesResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/trades?")
	request.AppendQuery(url)
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/accounts/{accountID}/openTrades
// Get the list of open Trades for an Account
func (c *Connection) TradesOpen(
	accountID AccountID,
) (*TradesResponse, error) {
	resp := &TradesResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/openTrades")
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/accounts/{accountID}/trades/{tradeSpecifier}
// Get the details of a specific Trade in an Account
func (c *Connection) Trade(
	accountID AccountID,
	specifier TradeSpecifier,
) (*TradeResponse, error) {
	resp := &TradeResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/trades/")
	_, _ = url.WriteString((string)(specifier))
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PUT /v3/accounts/{accountID}/trades/{tradeSpecifier}/close
// Close (partially or fully) a specific open Trade in an Account
func (c *Connection) TradeClose(
	accountID AccountID,
	specifier TradeSpecifier,
	units DecimalNumber,
) (*TradeCloseResponse, *TradeCloseError, error) {
	if len(units) == 0 {
		units = "ALL"
	}

	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/trades/")
	_, _ = url.WriteString((string)(specifier))
	_, _ = url.WriteString("/close")
	ctx := newCall(c, fasthttp.MethodPut, url, AcceptDatetimeFormat_RFC3339)
	defer ctx.release()

	b := bytebufferpool.Get()
	_, _ = b.WriteString("{\"units\":\"")
	_, _ = b.WriteString((string)(units))
	_, _ = b.WriteString("\"}")
	ctx.req.SetBody(b.Bytes())
	bytebufferpool.Put(b)

	err := fasthttp.DoRedirects(ctx.req, ctx.resp, maxRedirectsCount)
	if err != nil {
		return nil, nil, err
	}
	statusCode := ctx.resp.StatusCode()

	switch statusCode {
	// HTTP 200 – The Trade has been closed as requested
	case fasthttp.StatusOK:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &TradeCloseResponse{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return resp, nil, nil

	// HTTP 400 – The Trade cannot be closed as requested.
	// HTTP 404 – The Account or Trade specified does not exist.
	case fasthttp.StatusBadRequest, fasthttp.StatusNotFound:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &TradeCloseError{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return nil, resp, StatusCodeError{Code: statusCode}

	default:
		return nil, nil, StatusCodeError{Code: statusCode}
	}
}

// PUT /v3/accounts/{accountID}/trades/{tradeSpecifier}/clientExtensions
// Update the Client Extensions for a Trade. Do not add, update, or delete
// the Client Extensions if your account is associated with MT4.
func (c *Connection) TradeClientExtensions(
	accountID AccountID,
	specifier TradeSpecifier,
	request *OrderClientExtensionsRequest,
) (*OrderClientExtensionsResponse, *OrderClientExtensionsError, error) {
	if request == nil {
		return nil, nil, ErrNilRequest
	}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/trades/")
	_, _ = url.WriteString((string)(specifier))
	_, _ = url.WriteString("/clientExtensions")
	ctx := newCall(c, fasthttp.MethodPut, url, AcceptDatetimeFormat_RFC3339)
	defer ctx.release()

	w := &jwriter.Writer{}
	request.MarshalEasyJSON(w)
	// Set body
	ctx.req.SetBody(w.Buffer.BuildBytes())

	err := fasthttp.DoRedirects(ctx.req, ctx.resp, maxRedirectsCount)
	if err != nil {
		return nil, nil, err
	}
	statusCode := ctx.resp.StatusCode()

	switch statusCode {
	// HTTP 200 – The Order’s Client Extensions were successfully modified
	case fasthttp.StatusOK:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &OrderClientExtensionsResponse{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return resp, nil, nil

	// HTTP 400 – The Order specification was invalid
	// HTTP 404 – The Order or Account specified does not exist
	case fasthttp.StatusBadRequest, fasthttp.StatusNotFound:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &OrderClientExtensionsError{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return nil, resp, StatusCodeError{Code: statusCode}

	default:
		return nil, nil, StatusCodeError{Code: statusCode}
	}
}

// PUT /v3/accounts/{accountID}/trades/{tradeSpecifier}/orders
// Create, replace and cancel a Trade’s dependent Orders
//		Take Profit,
//		Stop Loss
//		Trailing Stop Loss
// through the Trade itself
func (c *Connection) TradeModify(
	accountID AccountID,
	specifier TradeSpecifier,
	request *TradeModifyRequest,
) (*TradeModifyResponse, *TradeModifyError, error) {
	if request == nil {
		return nil, nil, ErrNilRequest
	}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/trades/")
	_, _ = url.WriteString((string)(specifier))
	_, _ = url.WriteString("/orders")
	ctx := newCall(c, fasthttp.MethodPut, url, AcceptDatetimeFormat_RFC3339)
	defer ctx.release()

	w := &jwriter.Writer{}
	request.MarshalEasyJSON(w)
	// Set body
	ctx.req.SetBody(w.Buffer.BuildBytes())

	err := fasthttp.DoRedirects(ctx.req, ctx.resp, maxRedirectsCount)
	if err != nil {
		return nil, nil, err
	}
	statusCode := ctx.resp.StatusCode()

	switch statusCode {
	// HTTP 200 – The Trade’s dependent Orders have been modified as requested.
	case fasthttp.StatusOK:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &TradeModifyResponse{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return resp, nil, nil

	// HTTP 400 – The Trade’s dependent Orders cannot be modified as requested.
	case fasthttp.StatusBadRequest:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &TradeModifyError{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return nil, resp, StatusCodeError{Code: statusCode}

	default:
		return nil, nil, StatusCodeError{Code: statusCode}
	}
}
