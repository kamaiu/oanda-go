//go:generate easyjson -all $GOFILE
package model

import (
	"github.com/valyala/bytebufferpool"
	"strconv"
	"strings"
	"unsafe"
)

type OrdersRequest struct {
	IDs        []OrderID
	State      OrderStateFilter
	Instrument InstrumentName
	Count      int
	BeforeID   OrderID
}

func (g *OrdersRequest) WithIDs(ids ...OrderID) *OrdersRequest {
	g.IDs = ids
	return g
}
func (g *OrdersRequest) WithState(state OrderStateFilter) *OrdersRequest {
	g.State = state
	return g
}
func (g *OrdersRequest) WithInstrument(instrument InstrumentName) *OrdersRequest {
	g.Instrument = instrument
	return g
}
func (g *OrdersRequest) WithCount(count int) *OrdersRequest {
	g.Count = count
	return g
}
func (g *OrdersRequest) WithBeforeID(beforeID OrderID) *OrdersRequest {
	g.BeforeID = beforeID
	return g
}
func (g *OrdersRequest) AppendQuery(b *bytebufferpool.ByteBuffer) {
	_, _ = b.WriteString("count=")
	if g.Count <= 0 {
		g.Count = 50
	} else if g.Count > 500 {
		g.Count = 500
	}
	_, _ = b.WriteString(strconv.Itoa(g.Count))

	if len(g.State) == 0 {
		g.State = OrderStateFilter_PENDING
	}
	_, _ = b.WriteString("&state=")
	_, _ = b.WriteString((string)(g.State))

	if len(g.BeforeID) > 0 {
		_, _ = b.WriteString("&beforeID=")
		_, _ = b.WriteString((string)(g.BeforeID))
	}
	if len(g.IDs) > 0 {
		_, _ = b.WriteString("&ids=")
		_, _ = b.WriteString(strings.Join(*(*[]string)(unsafe.Pointer(&g.IDs)), ","))
	}
}

type OrdersResponse struct {
	// The list of pending Order details
	Orders []*Order `json:"orders"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type CreateOrderRequest struct {
	Order *OrderRequestParser `json:"order"`
}

type CreateOrderResponse struct {
	// The Transaction that created the Order specified by the request.
	OrderCreateTransaction *TransactionParser `json:"orderCreateTransaction"`
	// The Transaction that filled the newly created Order. Only provided when
	// the Order was immediately filled.
	OrderFillTransaction *OrderFillTransaction `json:"orderFillTransaction"`
	// The Transaction that cancelled the newly created Order. Only provided
	// when the Order was immediately cancelled.
	OrderCancelTransaction *OrderCancelTransaction `json:"orderCancelTransaction"`
	// The Transaction that reissues the Order. Only provided when the Order is
	// configured to be reissued for its remaining units after a partial fill
	// and the reissue was successful.
	OrderReissueTransaction *TransactionParser `json:"orderReissueTransaction"`
	// The Transaction that rejects the reissue of the Order. Only provided when
	// the Order is configured to be reissued for its remaining units after a
	// partial fill and the reissue was rejected.
	OrderReissueRejectTransaction *TransactionParser `json:"orderReissueRejectTransaction"`
	// The IDs of all Transactions that were created while satisfying the
	// request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type CreateOrderError struct {
	// The Transaction that rejected the creation of the Order as requested.
	// Only present if the Account exists.
	OrderRejectTransaction *TransactionParser `json:"orderRejectTransaction"`
	// The IDs of all Transactions that were created while satisfying the
	// request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
	// The code of the error that has occurred. This field may not be returned
	// for some errors.
	ErrorCode string `json:"errorCode"`
	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

type CancelOrderResponse struct {
	// The Transaction that cancelled the Order
	OrderCancelTransaction *OrderCancelTransaction `json:"orderCancelTransaction"`
	// The IDs of all Transactions that were created while satisfying the
	// request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type CancelOrderError struct {
	// The Transaction that rejected the cancellation of the Order. Only present
	// if the Account exists.
	OrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"orderCancelRejectTransaction"`
	// The IDs of all Transactions that were created while satisfying the
	// request. Only present if the Account exists.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
	// The ID of the most recent Transaction created for the Account. Only
	// present if the Account exists.
	ErrorCode string `json:"errorCode"`
	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

type OrderClientExtensionsRequest struct {
	// The Client Extensions to update for the Order. Do not set, modify, or
	// delete clientExtensions if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The Client Extensions to update for the Trade created when the Order is
	// filled. Do not set, modify, or delete clientExtensions if your account is
	// associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
}

type OrderClientExtensionsResponse struct {
	// The Transaction that modified the Client Extensions for the Order
	OrderClientExtensionsModifyTransaction *OrderClientExtensionsModifyTransaction `json:"orderClientExtensionsModifyTransaction"`
	// The IDs of all Transactions that were created while satisfying the
	// request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type OrderClientExtensionsError struct {
	// The Transaction that rejected the modification of the Client Extensions
	// for the Order
	OrderClientExtensionsModifyRejectTransaction *OrderClientExtensionsModifyRejectTransaction `json:"orderClientExtensionsModifyRejectTransaction"`
	// The IDs of all Transactions that were created while satisfying the
	// request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
	// The code of the error that has occurred. This field may not be returned
	// for some errors.
	ErrorCode string `json:"errorCode"`
	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}
