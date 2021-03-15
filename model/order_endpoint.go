//go:generate easyjson -all $GOFILE
package model

import (
	"github.com/valyala/bytebufferpool"
	"strconv"
)

type OrdersRequest struct {
	// List of Order IDs to retrieve
	IDs []OrderID `json:"ids"`
	// The state to filter the requested Orders by
	// [default=PENDING]
	State OrderStateFilter `json:"state"`
	// The instrument to filter the requested orders by
	Instrument InstrumentName `json:"instrument"`
	// The maximum number of Orders to return
	// [default=50, maximum=500]
	Count int `json:"count"`
	// The maximum Order ID to return. If not provided the most
	// recent Orders in the Account are returned
	BeforeID OrderID `json:"beforeID"`
}

func NewOrdersRequest() *OrdersRequest {
	return &OrdersRequest{
		State: OrderStateFilter_PENDING,
		Count: 50,
	}
}

// List of Order IDs to retrieve
func (g *OrdersRequest) WithIDs(ids ...OrderID) *OrdersRequest {
	g.IDs = ids
	return g
}

// The state to filter the requested Orders by
// [default=PENDING]
func (g *OrdersRequest) WithState(state OrderStateFilter) *OrdersRequest {
	switch state {
	case OrderStateFilter_PENDING,
		OrderStateFilter_CANCELLED,
		OrderStateFilter_FILLED,
		OrderStateFilter_TRIGGERED,
		OrderStateFilter_ALL:
		g.State = state
	default:
		g.State = OrderStateFilter_PENDING
	}
	return g
}

// The instrument to filter the requested orders by
func (g *OrdersRequest) WithInstrument(instrument InstrumentName) *OrdersRequest {
	g.Instrument = instrument
	return g
}

// The maximum number of Orders to return
// [default=50, maximum=500]
func (g *OrdersRequest) WithCount(count int) *OrdersRequest {
	if count < 1 {
		count = 50
	} else if count > 500 {
		count = 500
	} else {
		g.Count = count
	}
	return g
}

// The maximum Order ID to return. If not provided the most
// recent Orders in the Account are returned
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
		for i, id := range g.IDs {
			if i > 0 {
				_, _ = b.WriteString(UrlEncodedComma)
			}
			_, _ = b.WriteString((string)(id))
		}
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
