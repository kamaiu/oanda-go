package endpoint

import (
	. "github.com/kamaiu/oanda-go/model"
	"github.com/mailru/easyjson/jwriter"
	"github.com/valyala/bytebufferpool"
	"github.com/valyala/fasthttp"
)

// GET  /v3/accounts
// Get a list of all Accounts authorized for the provided token.
func (c *Connection) Accounts() (*AccountsResponse, error) {
	// Build URL
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts")

	resp := &AccountsResponse{}
	if _, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GET  /v3/accounts
// Get a list of all Accounts authorized for the provided token.
func (c *Connection) Account(
	id AccountID,
) (*Account, error) {
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(id))
	resp := &AccountResponse{}
	if _, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp); err != nil {
		return nil, err
	}
	return resp.Account, nil
}

// GET  /v3/accounts/{accountID}/summary
// Get a summary for a single Account that a client has access to.
func (c *Connection) AccountSummary(
	id AccountID,
) (*AccountSummaryResponse, error) {
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(id))
	_, _ = url.WriteString("/summary")
	resp := &AccountSummaryResponse{}
	if _, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// GET  /v3/accounts/{accountID}/instruments
// Get the list of tradeable instruments for the given Account. The list of tradeable
// instruments is dependent on the regulatory division that the Account is located in,
// thus should be the same for all Accounts owned by a single user.
func (c *Connection) AccountInstruments(
	id AccountID,
	filter ...string,
) (*AccountInstrumentsResponse, error) {
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(id))
	if len(filter) > 0 {
		_, _ = url.WriteString("/instruments?instruments=")
		for i, s := range filter {
			if i > 0 {
				_, _ = url.WriteString(UrlEncodedComma)
			}
			_, _ = url.WriteString(s)
		}
	} else {
		_, _ = url.WriteString("/instruments")
	}

	resp := &AccountInstrumentsResponse{}
	if _, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// PATCH  /v3/accounts/{accountID}/configuration
// Set the client-configurable portions of an Account.
func (c *Connection) AccountConfigure(
	id AccountID,
	config *AccountConfigurationRequest,
) (*AccountConfigurationResponse, *AccountConfigurationError, error) {
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(id))
	_, _ = url.WriteString("/v3/configuration")

	ctx := newCall(c, fasthttp.MethodPatch, url, AcceptDatetimeFormat_RFC3339)
	defer ctx.release()

	w := &jwriter.Writer{}
	config.MarshalEasyJSON(w)
	ctx.req.SetBody(w.Buffer.BuildBytes())

	err := fasthttp.DoRedirects(ctx.req, ctx.resp, maxRedirectsCount)
	if err != nil {
		return nil, nil, err
	}
	statusCode := ctx.resp.StatusCode()

	switch statusCode {
	// HTTP 200 – The Account was configured successfully.
	case fasthttp.StatusOK:
		body, err := readBody(ctx.resp)
		resp := &AccountConfigurationResponse{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return resp, nil, nil

	// HTTP 400 – The configuration specification was invalid.
	// HTTP 403 – The configuration operation was forbidden on the Account.
	case fasthttp.StatusBadRequest, fasthttp.StatusForbidden:
		body, err := readBody(ctx.resp)
		if err != nil {
			return nil, nil, err
		}
		resp := &AccountConfigurationError{}
		err = resp.UnmarshalJSON(body)
		if err != nil {
			return nil, nil, err
		}
		return nil, resp, StatusCodeError{Code: statusCode}

	default:
		return nil, nil, StatusCodeError{Code: statusCode}
	}
}

// GET  /v3/accounts/{accountID}/changes
//Endpoint used to poll an Account for its current state and changes since a specified TransactionID.
func (c *Connection) AccountChanges(
	id AccountID,
	sinceTransactionID TransactionID,
) (*AccountChangesResponse, error) {
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(id))
	_, _ = url.WriteString("/changes?sinceTransactionID=")
	_, _ = url.WriteString((string)(sinceTransactionID))

	resp := &AccountChangesResponse{}
	if _, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
