package endpoint

import (
	. "github.com/kamaiu/oanda-go/model"
	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
)

// GET /v3/accounts/{accountID}/positions
// List all Positions for an Account. The Positions returned are for every
// instrument that has had a position during the lifetime of an the Account.
func (c *Connection) Positions(
	accountID AccountID,
) (*PositionsResponse, error) {
	resp := &PositionsResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.hostname)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/positions")
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/accounts/{accountID}/openPositions
// List all open Positions for an Account. An open Position is a Position in
// an Account that currently has a Trade opened for it.
func (c *Connection) PositionsOpen(
	accountID AccountID,
) (*PositionsResponse, error) {
	resp := &PositionsResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.hostname)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/openPositions")
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/accounts/{accountID}/positions/{instrument}
// Get the details of a single Instrument’s Position in an Account. The Position
// may by open or not.
func (c *Connection) Position(
	accountID AccountID,
	instrument InstrumentName,
) (*PositionResponse, error) {
	resp := &PositionResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.hostname)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/positions/")
	_, _ = url.WriteString((string)(instrument))
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// PUT /v3/accounts/{accountID}/positions/{instrument}/close
// Closeout the open Position for a specific instrument in an Account.
func (c *Connection) PositionClose(
	accountID AccountID,
	instrument InstrumentName,
	request *PositionCloseRequest,
) (*PositionCloseResponse, *PositionCloseError, error) {
	if request == nil {
		return nil, nil, ErrNilRequest
	}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.hostname)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/positions/")
	_, _ = url.WriteString((string)(instrument))
	_, _ = url.WriteString("/close")
	ctx := newCall(c, fasthttp.MethodPut, url, AcceptDatetimeFormat_RFC3339)
	defer ctx.release()

	w := &jwriter.Writer{}
	request.MarshalEasyJSON(w)
	ctx.req.SetBody(w.Buffer.BuildBytes())

	err := fasthttp.DoRedirects(ctx.req, ctx.resp, maxRedirectsCount)
	if err != nil {
		return nil, nil, err
	}
	statusCode := ctx.resp.StatusCode()

	switch statusCode {
	// HTTP 200 – The Position closeout request has been successfully processed.
	case fasthttp.StatusOK:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &PositionCloseResponse{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return resp, nil, nil

	// HTTP 400 – The Parameters provided that describe the Position closeout are invalid.
	// HTTP 404 – The Account or one or more of the Positions specified does not exist.
	case fasthttp.StatusBadRequest, fasthttp.StatusNotFound:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &PositionCloseError{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return nil, resp, StatusCodeError{Code: statusCode}

	default:
		return nil, nil, StatusCodeError{Code: statusCode}
	}
}
