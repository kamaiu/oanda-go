//go:generate easyjson -all $GOFILE
package model

type PositionsResponse struct {
	// The list of Account Positions.
	Positions []*Position `json:"positions"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

// The requested Position.
type PositionResponse struct {
	Position *Position `json:"position"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type PositionCloseRequest struct {
	// Indication of how much of the long Position to closeout. Either the string “ALL”,
	// the string “NONE”, or a DecimalNumber representing how many units of the long
	// position to close using a PositionCloseout MarketOrder. The units specified must
	// always be positive.
	LongUnits string `json:"longUnits"`
	// The client extensions to add to the MarketOrder used to close the long position.
	LongClientExtensions *ClientExtensions `json:"longClientExtensions"`
	// Indication of how much of the short Position to closeout. Either the string “ALL”,
	// the string “NONE”, or a DecimalNumber representing how many units of the short
	// position to close using a PositionCloseout MarketOrder. The units specified must
	// always be positive.
	ShortUnits string `json:"shortUnits"`
	// The client extensions to add to the MarketOrder used to close the short position.
	ShortClientExtensions *ClientExtensions `json:"shortClientExtensions"`
}

type PositionCloseResponse struct {
	// The MarketOrderTransaction created to close the long Position.
	LongOrderCreateTransaction *MarketOrderTransaction `json:"longOrderCreateTransaction"`
	// OrderFill Transaction that closes the long Position
	LongOrderFillTransaction *OrderFillTransaction `json:"longOrderFillTransaction"`
	// OrderCancel Transaction that cancels the MarketOrder created to close the long Position
	LongOrderCancelTransaction *OrderCancelTransaction `json:"longOrderCancelTransaction"`
	// The MarketOrderTransaction created to close the short Position.
	ShortOrderCreateTransaction *MarketOrderTransaction `json:"shortOrderCreateTransaction"`
	// OrderFill Transaction that closes the short Position
	ShortOrderFillTransaction *OrderFillTransaction `json:"shortOrderFillTransaction"`
	// OrderCancel Transaction that cancels the MarketOrder created to close the short
	// Position
	ShortOrderCancelTransaction *OrderCancelTransaction `json:"shortOrderCancelTransaction"`
	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type PositionCloseError struct {
	// The Transaction created that rejects the creation of a MarketOrder to close the
	// long Position.
	LongOrderRejectTransaction *MarketOrderRejectTransaction `json:"longOrderRejectTransaction"`
	// The Transaction created that rejects the creation of a MarketOrder to close the
	// short Position.
	ShortOrderRejectTransaction *MarketOrderRejectTransaction `json:"shortOrderRejectTransaction"`
	// The IDs of all Transactions that were created while satisfying the request.
	RelatedTransactionIDs []TransactionID `json:"relatedTransactionIDs"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
	// The code of the error that has occurred. This field may not be returned for some
	// errors.
	ErrorCode string `json:"errorCode"`
	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}
