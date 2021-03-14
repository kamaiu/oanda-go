//go:generate easyjson -all $GOFILE
package model

import (
	"github.com/valyala/bytebufferpool"
	"net/url"
	"strconv"
)

type TransactionsRequest struct {
	// The starting time (inclusive) of the time range for the
	// Transactions being queried.
	// [default=Account Creation Time]
	From DateTime `json:"from"`
	// The ending time (inclusive) of the time range for the Transactions
	// being queried.
	// [default=Request Time]
	To DateTime `json:"to"`
	// The number of Transactions to include in each page of the results.
	// [default=100, maximum=1000]
	PageSize int `json:"pageSize"`
	// A filter for restricting the types of Transactions to retrieve.
	Type []TransactionFilter `json:"type"`
}

func NewTransactionsRequest(types ...TransactionFilter) *TransactionsRequest {
	return &TransactionsRequest{
		PageSize: 100,
		Type:     types,
	}
}

// The starting time (inclusive) of the time range for the
// Transactions being queried.
// [default=Account Creation Time]
func (t *TransactionsRequest) WithFrom(from DateTime) *TransactionsRequest {
	t.From = from
	return t
}

// The ending time (inclusive) of the time range for the Transactions
// being queried.
// [default=Request Time]
func (t *TransactionsRequest) WithTo(to DateTime) *TransactionsRequest {
	t.To = to
	return t
}

// The starting time (inclusive) of the time range for the
// Transactions being queried.
// [default=Account Creation Time]
//
// The ending time (inclusive) of the time range for the Transactions
// being queried.
// [default=Request Time]
func (t *TransactionsRequest) WithRange(from, to DateTime) *TransactionsRequest {
	t.From = from
	t.To = to
	return t
}

// The number of Transactions to include in each page of the results.
// [default=100, maximum=1000]
func (t *TransactionsRequest) WithPageSize(pageSize int) *TransactionsRequest {
	if pageSize < 1 {
		t.PageSize = 100
	} else if pageSize > 1000 {
		t.PageSize = 1000
	} else {
		t.PageSize = pageSize
	}
	return t
}

// A filter for restricting the types of Transactions to retrieve.
func (t *TransactionsRequest) WithType(types ...TransactionFilter) *TransactionsRequest {
	t.Type = types
	return t
}

func (g *TransactionsRequest) AppendQuery(b *bytebufferpool.ByteBuffer) {
	_, _ = b.WriteString("pageSize=")
	if g.PageSize <= 0 {
		g.PageSize = 50
	} else if g.PageSize > 1000 {
		g.PageSize = 1000
	}
	_, _ = b.WriteString(strconv.Itoa(g.PageSize))

	if len(g.From) > 0 {
		_, _ = b.WriteString("&from=")
		_, _ = b.WriteString(url.PathEscape((string)(g.From)))
	}
	if len(g.To) > 0 {
		_, _ = b.WriteString("&to=")
		_, _ = b.WriteString(url.PathEscape((string)(g.To)))
	}
	if len(g.Type) > 0 {
		_, _ = b.WriteString("&type=")
		for i, t := range g.Type {
			if i > 0 {
				_, _ = b.WriteString(",")
			}
			_, _ = b.WriteString((string)(t))
		}
	}
}

type TransactionsPagesResponse struct {
	// The starting time provided in the request.
	From DateTime `json:"from"`
	// The ending time provided in the request.
	To DateTime `json:"to"`
	// The pageSize provided in the request
	PageSize int64 `json:"pageSize"`
	// The Transaction-type filter provided in the request
	Type []TransactionFilter `json:"type"`
	// The number of Transactions that are contained in the pages returned
	Count int64 `json:"count"`
	// The list of URLs that represent idrange queries providing the data for
	// each page in the query results
	Pages []string `json:"pages"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type TransactionResponse struct {
	// The details of the Transaction requested
	Transaction *TransactionParser `json:"transaction"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type TransactionsResponse struct {
	// The list of Transactions that satisfy the request.
	Transactions []*TransactionParser `json:"transactions"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type TransactionsIDRangeRequest struct {
	// The starting Transaction ID (inclusive) to fetch.
	// [required]
	From TransactionID `json:"from"`
	// The ending Transaction ID (inclusive) to fetch.
	// [required]
	To TransactionID `json:"to"`
	// A filter for restricting the types of Transactions to retrieve.
	Type []TransactionFilter `json:"type"`
}

func NewTransactionsIDRangeRequest(from, to TransactionID, types ...TransactionFilter) *TransactionsIDRangeRequest {
	return &TransactionsIDRangeRequest{
		From: from,
		To:   to,
		Type: types,
	}
}

// The starting Transaction ID (inclusive) to fetch.
// [required]
func (t *TransactionsIDRangeRequest) WithFrom(from TransactionID) *TransactionsIDRangeRequest {
	t.From = from
	return t
}

// The ending Transaction ID (inclusive) to fetch.
// [required]
func (t *TransactionsIDRangeRequest) WithTo(to TransactionID) *TransactionsIDRangeRequest {
	t.To = to
	return t
}

// The starting Transaction ID (inclusive) to fetch.
// The ending Transaction ID (inclusive) to fetch.
// [required]
func (t *TransactionsIDRangeRequest) WithRange(from, to TransactionID) *TransactionsIDRangeRequest {
	t.From = from
	t.To = to
	return t
}

// A filter for restricting the types of Transactions to retrieve.
func (t *TransactionsIDRangeRequest) WithType(types ...TransactionFilter) *TransactionsIDRangeRequest {
	t.Type = types
	return t
}

func (g *TransactionsIDRangeRequest) AppendQuery(b *bytebufferpool.ByteBuffer) {
	_, _ = b.WriteString("from=")
	_, _ = b.WriteString((string)(g.From))
	_, _ = b.WriteString("&from=")
	_, _ = b.WriteString((string)(g.To))
	if len(g.Type) > 0 {
		_, _ = b.WriteString("&type=")
		for i, t := range g.Type {
			if i > 0 {
				_, _ = b.WriteString(",")
			}
			_, _ = b.WriteString((string)(t))
		}
	}
}

type TransactionsSinceIDRequest struct {
	// The ID of the last Transaction fetched. This query will return all
	// Transactions newer than the TransactionID.
	// [required]
	ID TransactionID `json:"id"`
	// A filter for restricting the types of Transactions to retrieve.
	Type []TransactionFilter `json:"type"`
}

func NewTransactionsSinceIDRequest(id TransactionID, types ...TransactionFilter) *TransactionsSinceIDRequest {
	return &TransactionsSinceIDRequest{ID: id, Type: types}
}

// The ID of the last Transaction fetched. This query will return all
// Transactions newer than the TransactionID.
// [required]
func (t *TransactionsSinceIDRequest) WithID(id TransactionID) *TransactionsSinceIDRequest {
	t.ID = id
	return t
}

// A filter for restricting the types of Transactions to retrieve.
func (t *TransactionsSinceIDRequest) WithType(types ...TransactionFilter) *TransactionsSinceIDRequest {
	t.Type = types
	return t
}

func (g *TransactionsSinceIDRequest) AppendQuery(b *bytebufferpool.ByteBuffer) {
	_, _ = b.WriteString("id=")
	_, _ = b.WriteString((string)(g.ID))
	if len(g.Type) > 0 {
		_, _ = b.WriteString("&type=")
		for i, t := range g.Type {
			if i > 0 {
				_, _ = b.WriteString(",")
			}
			_, _ = b.WriteString((string)(t))
		}
	}
}
