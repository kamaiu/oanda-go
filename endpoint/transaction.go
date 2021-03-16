package endpoint

import (
	. "github.com/kamaiu/oanda-go/model"
	"github.com/valyala/bytebufferpool"
)

// GET /v3/accounts/{accountID}/transactions
// Get a list of Transactions pages that satisfy a time-based Transaction query.
func (c *Connection) Transactions(
	accountID AccountID,
	request *TransactionsRequest,
) (*TransactionsPagesResponse, error) {
	resp := &TransactionsPagesResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/transactions?")
	request.AppendQuery(url)
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/accounts/{accountID}/transactions/{transactionID}
// Get the details of a single Account Transaction.
func (c *Connection) Transaction(
	accountID AccountID,
	id TransactionID,
) (*TransactionResponse, error) {
	resp := &TransactionResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/transactions/")
	_, _ = url.WriteString((string)(id))
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/accounts/{accountID}/transactions/idrange
// Get a range of Transactions for an Account based on the Transaction IDs.
func (c *Connection) TransactionsIDRange(
	accountID AccountID,
	request *TransactionsIDRangeRequest,
) (*TransactionsResponse, error) {
	resp := &TransactionsResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/transactions/")
	_, _ = url.WriteString("idrange?")
	request.AppendQuery(url)
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GET /v3/accounts/{accountID}/transactions/sinceid
// Get a range of Transactions for an Account starting at (but not including)
// a provided Transaction ID.
func (c *Connection) TransactionsSinceID(
	accountID AccountID,
	request *TransactionsSinceIDRequest,
) (*TransactionsResponse, error) {
	resp := &TransactionsResponse{}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.host)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/transactions/")
	_, _ = url.WriteString("sinceid?")
	request.AppendQuery(url)
	_, err := doGET(c, url, AcceptDatetimeFormat_RFC3339, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// TODO:
// GET /v3/accounts/{accountID}/transactions/Stream
// Get a Stream of Transactions for an Account starting from when the
// request is made.
// Note: This endpoint is served by the streaming URLs.
//
// Response Body Schema (application/octet-Stream)
// The response body for the Transaction Stream uses chunked transfer encoding.
// Each chunk contains Transaction and/or TransactionHeartbeat objects encoded
// as JSON. Each JSON object is serialized into a single line of text, and multiple
// objects found in the same chunk are separated by newlines.
// TransactionHeartbeats are sent every 5 seconds.
//
// The specification for the objects found in response Stream are as follows:
//
//		Transaction
//		TransactionHeartbeat
