package endpoint

import (
	"errors"
	. "github.com/kamaiu/oanda-go/model"
	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
)

// POST /v3/accounts/{accountID}/orders
// Create an Order for an Account
func (c *Connection) OrderCreate(
	accountID AccountID,
	request OrderRequest,
) (*CreateOrderResponse, *CreateOrderError, error) {
	w := &jwriter.Writer{}
	w.RawByte('{')
	w.RawString("\"order\":")

	switch v := request.(type) {
	case *MarketOrderRequest:
		v.MarshalEasyJSON(w)
	case *LimitOrderRequest:
		v.MarshalEasyJSON(w)
	case *StopOrderRequest:
		v.MarshalEasyJSON(w)
	case *MarketIfTouchedOrderRequest:
		v.MarshalEasyJSON(w)
	case *TakeProfitOrderRequest:
		v.MarshalEasyJSON(w)
	case *StopLossOrderRequest:
		v.MarshalEasyJSON(w)
	case *GuaranteedStopLossOrderRequest:
		v.MarshalEasyJSON(w)
	case *TrailingStopLossOrderRequest:
		v.MarshalEasyJSON(w)
	default:
		return nil, nil, errors.New("unsupported order request type")
	}
	w.RawByte('}')

	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/orders")
	ctx := newCall(c, fasthttp.MethodPost, url, AcceptDatetimeFormat_RFC3339)
	defer ctx.release()

	// Set body
	ctx.req.SetBody(w.Buffer.BuildBytes())

	err := fasthttp.DoRedirects(ctx.req, ctx.resp, maxRedirectsCount)
	if err != nil {
		return nil, nil, err
	}
	statusCode := ctx.resp.StatusCode()

	switch statusCode {
	// HTTP 201 – The Order was created as specified
	case fasthttp.StatusCreated:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &CreateOrderResponse{}
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
		resp := &CreateOrderError{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return nil, resp, StatusCodeError{Code: statusCode}

	default:
		return nil, nil, StatusCodeError{Code: statusCode}
	}
}

// GET  /v3/accounts/{accountID}/orders
// Get a list of Orders for an Account
func (c *Connection) Orders(
	accountID AccountID,
	request *OrdersRequest,
) (*OrdersResponse, error) {
	if request == nil {
		request = &OrdersRequest{}
	}
	resp := &OrdersResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/orders?")
	request.AppendQuery(url)
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET	/v3/accounts/{accountID}/pendingOrders
// List all pending Orders in an Account
func (c *Connection) OrdersPending(
	accountID AccountID,
) (*OrdersResponse, error) {
	resp := &OrdersResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/pendingOrders")
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/accounts/{accountID}/orders/{orderSpecifier}
// Get details for a single Order in an Account
func (c *Connection) OrdersBySpecifier(
	accountID AccountID,
	specifier OrderSpecifier,
) (*OrdersResponse, error) {
	resp := &OrdersResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/orders/")
	_, _ = url.WriteString((string)(specifier))
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PUT /v3/accounts/{accountID}/orders/{orderSpecifier}
// Replace an Order in an Account by simultaneously cancelling
// it and creating a replacement Order
func (c *Connection) OrderReplace(
	accountID AccountID,
	specifier OrderSpecifier,
	order OrderRequest,
) (*CreateOrderResponse, *CreateOrderError, error) {
	w := &jwriter.Writer{}
	w.RawByte('{')
	w.RawString("\"order\":")

	switch v := order.(type) {
	case *MarketOrderRequest:
		v.MarshalEasyJSON(w)
	case *LimitOrderRequest:
		v.MarshalEasyJSON(w)
	case *StopOrderRequest:
		v.MarshalEasyJSON(w)
	case *MarketIfTouchedOrderRequest:
		v.MarshalEasyJSON(w)
	case *TakeProfitOrderRequest:
		v.MarshalEasyJSON(w)
	case *StopLossOrderRequest:
		v.MarshalEasyJSON(w)
	case *GuaranteedStopLossOrderRequest:
		v.MarshalEasyJSON(w)
	case *TrailingStopLossOrderRequest:
		v.MarshalEasyJSON(w)
	default:
		return nil, nil, errors.New("unsupported order request type")
	}
	w.RawByte('}')

	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/orders/")
	_, _ = url.WriteString((string)(specifier))
	ctx := newCall(c, fasthttp.MethodPut, url, AcceptDatetimeFormat_RFC3339)
	defer ctx.release()

	// Set body
	ctx.req.SetBody(w.Buffer.BuildBytes())

	err := fasthttp.DoRedirects(ctx.req, ctx.resp, maxRedirectsCount)
	if err != nil {
		return nil, nil, err
	}
	statusCode := ctx.resp.StatusCode()

	switch statusCode {
	// HTTP 201 – The Order was created as specified
	case fasthttp.StatusCreated:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &CreateOrderResponse{}
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
		resp := &CreateOrderError{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return nil, resp, StatusCodeError{Code: statusCode}

	default:
		return nil, nil, StatusCodeError{Code: statusCode}
	}
}

// PUT /v3/accounts/{accountID}/orders/{orderSpecifier}
// Replace an Order in an Account by simultaneously cancelling
// it and creating a replacement Order
func (c *Connection) OrderCancel(
	accountID AccountID,
	specifier OrderSpecifier,
) (*CancelOrderResponse, *CancelOrderError, error) {
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/orders/")
	_, _ = url.WriteString((string)(specifier))
	_, _ = url.WriteString("/cancel")
	ctx := newCall(c, fasthttp.MethodPut, url, AcceptDatetimeFormat_RFC3339)
	defer ctx.release()

	err := fasthttp.DoRedirects(ctx.req, ctx.resp, maxRedirectsCount)
	if err != nil {
		return nil, nil, err
	}
	statusCode := ctx.resp.StatusCode()

	switch statusCode {
	// HTTP 200 – The Order was cancelled as specified
	case fasthttp.StatusOK:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &CancelOrderResponse{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return resp, nil, nil

	// HTTP 404 – The Account or Order specified does not exist.
	case fasthttp.StatusNotFound:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &CancelOrderError{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return nil, resp, StatusCodeError{Code: statusCode}

	default:
		return nil, nil, StatusCodeError{Code: statusCode}
	}
}

// PUT /v3/accounts/{accountID}/orders/{orderSpecifier}/clientExtensions
// Update the Client Extensions for an Order in an Account. Do not set,
// modify, or delete clientExtensions if your account is associated with MT4.
func (c *Connection) OrderClientExtensions(
	accountID AccountID,
	specifier OrderSpecifier,
	request *OrderClientExtensionsRequest,
) (*OrderClientExtensionsResponse, *OrderClientExtensionsError, error) {
	if request == nil {
		return nil, nil, ErrNilRequest
	}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/orders/")
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
