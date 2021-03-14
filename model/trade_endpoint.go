//go:generate easyjson -all $GOFILE
package model

import (
	"github.com/valyala/bytebufferpool"
	"strconv"
)

type TradesRequest struct {
	// List of Trade IDs to retrieve.
	IDs []TradeID `json:"ids"`
	// The state to filter the requested Trades by.
	// [default=OPEN]
	State TradeStateFilter `json:"state"`
	// The instrument to filter the requested Trades by.
	Instrument InstrumentName `json:"instrument"`
	// The maximum number of Trades to return. [default=50, maximum=500]
	Count int `json:"count"`
	// The maximum Trade ID to return. If not provided the most recent
	// Trades in the Account are returned.
	BeforeID TradeID `json:"beforeID"`
}

func (g *TradesRequest) AppendQuery(b *bytebufferpool.ByteBuffer) {
	_, _ = b.WriteString("count=")
	if g.Count <= 0 {
		g.Count = 50
	} else if g.Count > 500 {
		g.Count = 500
	}
	_, _ = b.WriteString(strconv.Itoa(g.Count))

	if len(g.State) == 0 {
		g.State = TradeStateFilter_OPEN
	}
	_, _ = b.WriteString("&state=")
	_, _ = b.WriteString((string)(g.State))

	if len(g.Instrument) > 0 {
		_, _ = b.WriteString("&instrument=")
		_, _ = b.WriteString((string)(g.Instrument))
	}

	if len(g.BeforeID) > 0 {
		_, _ = b.WriteString("&beforeID=")
		_, _ = b.WriteString((string)(g.BeforeID))
	}
	if len(g.IDs) > 0 {
		_, _ = b.WriteString("&ids=")
		for i, id := range g.IDs {
			if i > 0 {
				_, _ = b.WriteString(",")
			}
			_, _ = b.WriteString((string)(id))
		}
	}
}

type TradesResponse struct {
	// The list of Trade detail objects
	Trades []*Trade `json:"trades"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type TradeResponse struct {
	// The details of the requested trade
	Trade *Trade `json:"trade"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type TradeCloseResponse struct {
	// The MarketOrder Transaction created to close the Trade.
	OrderCreateTransaction *MarketOrderTransaction `json:"orderCreateTransaction"`
	// The OrderFill Transaction that fills the Trade-closing MarketOrder and
	// closes the Trade.
	OrderFillTransaction *OrderFillTransaction `json:"orderFillTransaction"`
	// The OrderCancel Transaction that immediately cancelled the Trade-closing MarketOrder.
	OrderCancelTransaction *OrderCancelTransaction `json:"orderCancelTransaction"`
	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type TradeCloseError struct {
	// The MarketOrderReject Transaction that rejects the creation of the Trade-closing MarketOrder.
	OrderRejectTransaction *MarketOrderRejectTransaction `json:"orderRejectTransaction"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The code of the error that has occurred. This field may not be returned
	// for some errors.
	ErrorCode string `json:"errorCode"`
	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

type TradeClientExtensionsRequest struct {
	// The Client Extensions to update the Trade with. Do not add, update, or
	// delete the Client Extensions if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

type TradeClientExtensionsResponse struct {
	// The Transaction that updates the Trade’s Client Extensions.
	TradeClientExtensionsModifyTransaction *TradeClientExtensionsModifyTransaction `json:"tradeClientExtensionsModifyTransaction"`
	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type TradeClientExtensionsError struct {
	// The Transaction that rejects the modification of the Trade’s Client Extensions.
	TradeClientExtensionsModifyRejectTransaction *TradeClientExtensionsModifyRejectTransaction `json:"tradeClientExtensionsModifyRejectTransaction"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The code of the error that has occurred. This field may not be returned
	// for some errors.
	ErrorCode string `json:"errorCode"`
	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

type TradeModifyRequest struct {
	// The specification of the Take Profit to create/modify/cancel. If
	// takeProfit is set to null, the Take Profit Order will be cancelled if it
	// exists. If takeProfit is not provided, the existing Take Profit Order
	// will not be modified. If a sub-field of takeProfit is not specified, that
	// field will be set to a default value on create, and be inherited by the
	// replacing order on modify.
	TakeProfit *TakeProfitDetails `json:"takeProfit"`
	// The specification of the Stop Loss to create/modify/cancel. If stopLoss
	// is set to null, the Stop Loss Order will be cancelled if it exists. If
	// stopLoss is not provided, the existing Stop Loss Order will not be
	// modified. If a sub-field of stopLoss is not specified, that field will be
	// set to a default value on create, and be inherited by the replacing order
	// on modify.
	StopLoss *StopLossDetails `json:"stopLoss"`
	// The specification of the Trailing Stop Loss to create/modify/cancel. If
	// trailingStopLoss is set to null, the Trailing Stop Loss Order will be
	// cancelled if it exists. If trailingStopLoss is not provided, the existing
	// Trailing Stop Loss Order will not be modified. If a sub-field of
	// trailingStopLoss is not specified, that field will be set to a default
	// value on create, and be inherited by the replacing order on modify.
	TrailingStopLoss *TrailingStopLossDetails `json:"trailingStopLoss"`
	// The specification of the Guaranteed Stop Loss to create/modify/cancel. If
	// guaranteedStopLoss is set to null, the Guaranteed Stop Loss Order will be
	// cancelled if it exists. If guaranteedStopLoss is not provided, the
	// existing Guaranteed Stop Loss Order will not be modified. If a sub-field
	// of guaranteedStopLoss is not specified, that field will be set to a
	// default value on create, and be inherited by the replacing order on
	// modify.
	GuaranteedStopLoss *GuaranteedStopLossDetails `json:"guaranteedStopLoss"`
}

type TradeModifyResponse struct {
	// The Transaction created that cancels the Trade’s existing Take Profit Order.
	TakeProfitOrderCancelTransaction *OrderCancelTransaction `json:"takeProfitOrderCancelTransaction"`
	// The Transaction created that creates a new Take Profit Order for the Trade.
	TakeProfitOrderTransaction *TakeProfitOrderTransaction `json:"takeProfitOrderTransaction"`
	// The Transaction created that immediately fills the Trade’s new Take Profit Order.
	// Only provided if the new Take Profit Order was immediately filled.
	TakeProfitOrderFillTransaction *OrderFillTransaction `json:"takeProfitOrderFillTransaction"`
	// The Transaction created that immediately cancels the Trade’s new Take Profit Order.
	// Only provided if the new Take Profit Order was immediately cancelled.
	TakeProfitOrderCreatedCancelTransaction *OrderCancelTransaction `json:"takeProfitOrderCreatedCancelTransaction"`
	// The Transaction created that cancels the Trade’s existing Stop Loss Order.
	StopLossOrderCancelTransaction *OrderCancelTransaction `json:"stopLossOrderCancelTransaction"`
	// The Transaction created that creates a new Stop Loss Order for the Trade.
	StopLossOrderTransaction *StopLossOrderTransaction `json:"stopLossOrderTransaction"`
	// The Transaction created that immediately fills the Trade’s new Stop Order. Only
	// provided if the new Stop Loss Order was immediately filled.
	StopLossOrderFillTransaction *OrderFillTransaction `json:"stopLossOrderFillTransaction"`
	// The Transaction created that immediately cancels the Trade’s new Stop Loss Order.
	// Only provided if the new Stop Loss Order was immediately cancelled.
	StopLossOrderCreatedCancelTransaction *OrderCancelTransaction `json:"stopLossOrderCreatedCancelTransaction"`
	// The Transaction created that cancels the Trade’s existing Trailing Stop Loss Order.
	TrailingStopLossOrderCancelTransaction *OrderCancelTransaction `json:"trailingStopLossOrderCancelTransaction"`
	// The Transaction created that creates a new Trailing Stop Loss Order for the Trade.
	TrailingStopLossOrderTransaction *TrailingStopLossOrderTransaction `json:"trailingStopLossOrderTransaction"`
	// The Transaction created that cancels the Trade’s existing Guaranteed Stop Loss
	// Order.
	GuaranteedStopLossOrderCancelTransaction *OrderCancelTransaction `json:"guaranteedStopLossOrderCancelTransaction"`
	// The Transaction created that creates a new Guaranteed Stop Loss Order for the Trade.
	GuaranteedStopLossOrderTransaction *GuaranteedStopLossOrderTransaction `json:"guaranteedStopLossOrderTransaction"`
	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type TradeModifyError struct {
	// An OrderCancelRejectTransaction represents the rejection of the cancellation of
	// an Order in the client’s Account.
	TakeProfitOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"takeProfitOrderCancelRejectTransaction"`
	// A TakeProfitOrderRejectTransaction represents the rejection of the creation of a
	// TakeProfit Order.
	TakeProfitOrderRejectTransaction *TakeProfitOrderRejectTransaction `json:"takeProfitOrderRejectTransaction"`
	// An OrderCancelRejectTransaction represents the rejection of the cancellation of
	// an Order in the client’s Account.
	StopLossOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"stopLossOrderCancelRejectTransaction"`
	// A StopLossOrderRejectTransaction represents the rejection of the creation of a StopLoss
	// Order.
	StopLossOrderRejectTransaction *StopLossOrderRejectTransaction `json:"stopLossOrderRejectTransaction"`
	// An OrderCancelRejectTransaction represents the rejection of the cancellation of
	// an Order in the client’s Account.
	TrailingStopLossOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"trailingStopLossOrderCancelRejectTransaction"`
	// A TrailingStopLossOrderRejectTransaction represents the rejection of the creation
	// of a TrailingStopLoss Order.
	TrailingStopLossOrderRejectTransaction *TrailingStopLossOrderRejectTransaction `json:"trailingStopLossOrderRejectTransaction"`
	// An OrderCancelRejectTransaction represents the rejection of the cancellation of
	// an Order in the client’s Account.
	GuaranteedStopLossOrderCancelRejectTransaction *OrderCancelRejectTransaction `json:"guaranteedStopLossOrderCancelRejectTransaction"`
	// A GuaranteedStopLossOrderRejectTransaction represents the rejection of the creation
	// of a GuaranteedStopLoss Order.
	GuaranteedStopLossOrderRejectTransaction *GuaranteedStopLossOrderRejectTransaction `json:"guaranteedStopLossOrderRejectTransaction"`
	// The code of the error that has occurred. This field may not be returned for some
	// errors.
	ErrorCode string `json:"errorCode"`
	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}
