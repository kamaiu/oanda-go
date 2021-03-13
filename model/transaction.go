//go:generate easyjson -all $GOFILE
package model

// The unique Transaction identifier within each Account.
// String representation of the numerical OANDA-assigned TransactionID
type TransactionID string

// The base Transaction specification. Specifies properties that are common between
// all Transaction.
type Transaction struct {
	// The Transaction’s Identifier.
	Id TransactionID `json:"id"`
	// The date/time when the Transaction was created.
	Time DateTime `json:"time"`
	// The ID of the user that initiated the creation of the Transaction.
	UserID int64 `json:"userID"`
	// The ID of the Account the Transaction was created for.
	AccountID AccountID `json:"accountID"`
	// The ID of the “batch” that the Transaction belongs to. Transactions in the same
	// batch are applied to the Account simultaneously.
	BatchID TransactionID `json:"batchID"`
	// The Request ID of the request which generated the transaction.
	RequestID RequestID `json:"requestID"`
}

// A CreateTransaction represents the creation of an Account.
type CreateTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “CREATE” in a CreateTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Division that the Account is in
	DivisionID int64 `json:"divisionID"`
	// The ID of the Site that the Account was created at
	SiteID int64 `json:"siteID"`
	// The ID of the user that the Account was created for
	AccountUserID int64 `json:"accountUserID"`
	// The number of the Account within the site/division/user
	AccountNumber int64 `json:"accountNumber"`
	// The home currency of the Account
	HomeCurrency Currency `json:"homeCurrency"`
}

// A CloseTransaction represents the closing of an Account.
type CloseTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “CLOSE” in a CloseTransaction.
	Type TransactionType `json:"type"`
}

// A ReopenTransaction represents the re-opening of a closed Account.
type ReopenTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “REOPEN” in a ReopenTransaction.
	Type TransactionType `json:"type"`
}

// A ClientConfigureTransaction represents the configuration of an Account by a client.
type ClientConfigureTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “CLIENT_CONFIGURE” in a ClientConfigureTransaction.
	Type TransactionType `json:"type"`
	// The client-provided alias for the Account.
	Alias string `json:"alias"`
	// The margin rate override for the Account.
	MarginRate DecimalNumber `json:"marginRate"`
}

// A ClientConfigureRejectTransaction represents the reject of configuration of an Account
// by a client.
type ClientConfigureRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “CLIENT_CONFIGURE_REJECT” in a ClientConfigureRejectTransaction.
	Type TransactionType `json:"type"`
	// The client-provided alias for the Account.
	Alias string `json:"alias"`
	// The margin rate override for the Account.
	MarginRate DecimalNumber `json:"marginRate"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A TransferFundsTransaction represents the transfer of funds in/out of an Account.
type TransferFundsTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “TRANSFER_FUNDS” in a TransferFundsTransaction.
	Type TransactionType `json:"type"`
	// The amount to deposit/withdraw from the Account in the Account’s home currency.
	// A positive value indicates a deposit, a negative value indicates a withdrawal.
	Amount AccountUnits `json:"amount"`
	// The reason that an Account is being funded.
	FundingReason FundingReason `json:"fundingReason"`
	// An optional comment that may be attached to a fund transfer for audit purposes
	Comment string `json:"comment"`
	// The Account’s balance after funds are transferred.
	AccountBalance AccountUnits `json:"accountBalance"`
}

// A TransferFundsRejectTransaction represents the rejection of the transfer of funds
// in/out of an Account.
type TransferFundsRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “TRANSFER_FUNDS_REJECT” in a TransferFundsRejectTransaction.
	Type TransactionType `json:"type"`
	// The amount to deposit/withdraw from the Account in the Account’s home currency.
	// A positive value indicates a deposit, a negative value indicates a withdrawal.
	Amount AccountUnits `json:"amount"`
	// The reason that an Account is being funded.
	FundingReason FundingReason `json:"fundingReason"`
	// An optional comment that may be attached to a fund transfer for audit purposes
	Comment string `json:"comment"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A MarketOrderTransaction represents the creation of a Market Order in the user’s
// account. A Market Order is an Order that is filled immediately at the current market
// price. Market Orders can be specialized when they are created to accomplish a specific
// task: to close a Trade, to closeout a Position or to participate in in a Margin closeout.
type MarketOrderTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “MARKET_ORDER” in a MarketOrderTransaction.
	Type TransactionType `json:"type"`
	// The Market Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The quantity requested to be filled by the Market Order. A positive number of units
	// results in a long Order, and a negative number of units results in a short Order.
	Units DecimalNumber `json:"units"`
	// The time-in-force requested for the Market Order. Restricted to FOK or IOC for a
	// MarketOrder.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The worst price that the client is willing to have the Market Order filled at.
	PriceBound PriceValue `json:"priceBound"`
	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`
	// Details of the Trade requested to be closed, only provided when the Market Order
	// is being used to explicitly close a Trade.
	TradeClose *MarketOrderTradeClose `json:"tradeClose"`
	// Details of the long Position requested to be closed out, only provided when a Market
	// Order is being used to explicitly closeout a long Position.
	LongPositionCloseout *MarketOrderPositionCloseout `json:"longPositionCloseout"`
	// Details of the short Position requested to be closed out, only provided when a Market
	// Order is being used to explicitly closeout a short Position.
	ShortPositionCloseout *MarketOrderPositionCloseout `json:"shortPositionCloseout"`
	// Details of the Margin Closeout that this Market Order was created for
	MarginCloseout *MarketOrderMarginCloseout `json:"marginCloseout"`
	// Details of the delayed Trade close that this Market Order was created for
	DelayedTradeClose *MarketOrderDelayedTradeClose `json:"delayedTradeClose"`
	// The reason that the Market Order was created
	Reason MarketOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The specification of the Take Profit Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// The specification of the Stop Loss Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// The specification of the Trailing Stop Loss Order that should be created for a Trade
	// that is opened when the Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// The specification of the Guaranteed Stop Loss Order that should be created for a
	// Trade that is opened when the Order is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created).  Do not set, modify, delete tradeClientExtensions if your account
	// is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
}

// A MarketOrderRejectTransaction represents the rejection of the creation of a Market
// Order.
type MarketOrderRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “MARKET_ORDER_REJECT” in a MarketOrderRejectTransaction.
	Type TransactionType `json:"type"`
	// The Market Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The quantity requested to be filled by the Market Order. A positive number of units
	// results in a long Order, and a negative number of units results in a short Order.
	Units DecimalNumber `json:"units"`
	// The time-in-force requested for the Market Order. Restricted to FOK or IOC for a
	// MarketOrder.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The worst price that the client is willing to have the Market Order filled at.
	PriceBound PriceValue `json:"priceBound"`
	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`
	// Details of the Trade requested to be closed, only provided when the Market Order
	// is being used to explicitly close a Trade.
	TradeClose *MarketOrderTradeClose `json:"tradeClose"`
	// Details of the long Position requested to be closed out, only provided when a Market
	// Order is being used to explicitly closeout a long Position.
	LongPositionCloseout *MarketOrderPositionCloseout `json:"longPositionCloseout"`
	// Details of the short Position requested to be closed out, only provided when a Market
	// Order is being used to explicitly closeout a short Position.
	ShortPositionCloseout *MarketOrderPositionCloseout `json:"shortPositionCloseout"`
	// Details of the Margin Closeout that this Market Order was created for
	MarginCloseout *MarketOrderMarginCloseout `json:"marginCloseout"`
	// Details of the delayed Trade close that this Market Order was created for
	DelayedTradeClose *MarketOrderDelayedTradeClose `json:"delayedTradeClose"`
	// The reason that the Market Order was created
	Reason MarketOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The specification of the Take Profit Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// The specification of the Stop Loss Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// The specification of the Trailing Stop Loss Order that should be created for a Trade
	// that is opened when the Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// The specification of the Guaranteed Stop Loss Order that should be created for a
	// Trade that is opened when the Order is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created).  Do not set, modify, delete tradeClientExtensions if your account
	// is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A FixedPriceOrderTransaction represents the creation of a Fixed Price Order in the
// user’s account. A Fixed Price Order is an Order that is filled immediately at a
// specified price.
type FixedPriceOrderTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “FIXED_PRICE_ORDER” in a FixedPriceOrderTransaction.
	Type TransactionType `json:"type"`
	// The Fixed Price Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The quantity requested to be filled by the Fixed Price Order. A positive number
	// of units results in a long Order, and a negative number of units results in a short
	// Order.
	Units DecimalNumber `json:"units"`
	// The price specified for the Fixed Price Order. This price is the exact price that
	// the Fixed Price Order will be filled at.
	Price PriceValue `json:"price"`
	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`
	// The state that the trade resulting from the Fixed Price Order should be set to.
	TradeState string `json:"tradeState"`
	// The reason that the Fixed Price Order was created
	Reason FixedPriceOrderReason `json:"reason"`
	// The client extensions for the Fixed Price Order.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The specification of the Take Profit Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// The specification of the Stop Loss Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// The specification of the Trailing Stop Loss Order that should be created for a Trade
	// that is opened when the Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// The specification of the Guaranteed Stop Loss Order that should be created for a
	// Trade that is opened when the Order is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created).  Do not set, modify, delete tradeClientExtensions if your account
	// is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
}

// A LimitOrderTransaction represents the creation of a Limit Order in the user’s
// Account.
type LimitOrderTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “LIMIT_ORDER” in a LimitOrderTransaction.
	Type TransactionType `json:"type"`
	// The Limit Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The quantity requested to be filled by the Limit Order. A positive number of units
	// results in a long Order, and a negative number of units results in a short Order.
	Units DecimalNumber `json:"units"`
	// The price threshold specified for the Limit Order. The Limit Order will only be
	// filled by a market price that is equal to or better than this price.
	Price PriceValue `json:"price"`
	// The time-in-force requested for the Limit Order.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the Limit Order will be cancelled if its timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Limit Order was initiated
	Reason LimitOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The specification of the Take Profit Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// The specification of the Stop Loss Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// The specification of the Trailing Stop Loss Order that should be created for a Trade
	// that is opened when the Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// The specification of the Guaranteed Stop Loss Order that should be created for a
	// Trade that is opened when the Order is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created).  Do not set, modify, delete tradeClientExtensions if your account
	// is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// The ID of the Order that this Order replaces (only provided if this Order replaces
	// an existing Order).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Transaction that cancels the replaced Order (only provided if this
	// Order replaces an existing Order).
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
}

// A LimitOrderRejectTransaction represents the rejection of the creation of a Limit
// Order.
type LimitOrderRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “LIMIT_ORDER_REJECT” in a LimitOrderRejectTransaction.
	Type TransactionType `json:"type"`
	// The Limit Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The quantity requested to be filled by the Limit Order. A positive number of units
	// results in a long Order, and a negative number of units results in a short Order.
	Units DecimalNumber `json:"units"`
	// The price threshold specified for the Limit Order. The Limit Order will only be
	// filled by a market price that is equal to or better than this price.
	Price PriceValue `json:"price"`
	// The time-in-force requested for the Limit Order.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the Limit Order will be cancelled if its timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Limit Order was initiated
	Reason LimitOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The specification of the Take Profit Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// The specification of the Stop Loss Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// The specification of the Trailing Stop Loss Order that should be created for a Trade
	// that is opened when the Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// The specification of the Guaranteed Stop Loss Order that should be created for a
	// Trade that is opened when the Order is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created).  Do not set, modify, delete tradeClientExtensions if your account
	// is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// The ID of the Order that this Order was intended to replace (only provided if this
	// Order was intended to replace an existing Order).
	IntendedReplacesOrderID OrderID `json:"intendedReplacesOrderID"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A StopOrderTransaction represents the creation of a Stop Order in the user’s Account.
type StopOrderTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “STOP_ORDER” in a StopOrderTransaction.
	Type TransactionType `json:"type"`
	// The Stop Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The quantity requested to be filled by the Stop Order. A positive number of units
	// results in a long Order, and a negative number of units results in a short Order.
	Units DecimalNumber `json:"units"`
	// The price threshold specified for the Stop Order. The Stop Order will only be filled
	// by a market price that is equal to or worse than this price.
	Price PriceValue `json:"price"`
	// The worst market price that may be used to fill this Stop Order. If the market gaps
	// and crosses through both the price and the priceBound, the Stop Order will be cancelled
	// instead of being filled.
	PriceBound PriceValue `json:"priceBound"`
	// The time-in-force requested for the Stop Order.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the Stop Order will be cancelled if its timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Stop Order was initiated
	Reason StopOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The specification of the Take Profit Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// The specification of the Stop Loss Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// The specification of the Trailing Stop Loss Order that should be created for a Trade
	// that is opened when the Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// The specification of the Guaranteed Stop Loss Order that should be created for a
	// Trade that is opened when the Order is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created).  Do not set, modify, delete tradeClientExtensions if your account
	// is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// The ID of the Order that this Order replaces (only provided if this Order replaces
	// an existing Order).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Transaction that cancels the replaced Order (only provided if this
	// Order replaces an existing Order).
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
}

// A StopOrderRejectTransaction represents the rejection of the creation of a Stop Order.
type StopOrderRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “STOP_ORDER_REJECT” in a StopOrderRejectTransaction.
	Type TransactionType `json:"type"`
	// The Stop Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The quantity requested to be filled by the Stop Order. A positive number of units
	// results in a long Order, and a negative number of units results in a short Order.
	Units DecimalNumber `json:"units"`
	// The price threshold specified for the Stop Order. The Stop Order will only be filled
	// by a market price that is equal to or worse than this price.
	Price PriceValue `json:"price"`
	// The worst market price that may be used to fill this Stop Order. If the market gaps
	// and crosses through both the price and the priceBound, the Stop Order will be cancelled
	// instead of being filled.
	PriceBound PriceValue `json:"priceBound"`
	// The time-in-force requested for the Stop Order.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the Stop Order will be cancelled if its timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Stop Order was initiated
	Reason StopOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The specification of the Take Profit Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// The specification of the Stop Loss Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// The specification of the Trailing Stop Loss Order that should be created for a Trade
	// that is opened when the Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// The specification of the Guaranteed Stop Loss Order that should be created for a
	// Trade that is opened when the Order is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created).  Do not set, modify, delete tradeClientExtensions if your account
	// is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// The ID of the Order that this Order was intended to replace (only provided if this
	// Order was intended to replace an existing Order).
	IntendedReplacesOrderID OrderID `json:"intendedReplacesOrderID"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A MarketIfTouchedOrderTransaction represents the creation of a MarketIfTouched Order
// in the user’s Account.
type MarketIfTouchedOrderTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “MARKET_IF_TOUCHED_ORDER” in a MarketIfTouchedOrderTransaction.
	Type TransactionType `json:"type"`
	// The MarketIfTouched Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The quantity requested to be filled by the MarketIfTouched Order. A positive number
	// of units results in a long Order, and a negative number of units results in a short
	// Order.
	Units DecimalNumber `json:"units"`
	// The price threshold specified for the MarketIfTouched Order. The MarketIfTouched
	// Order will only be filled by a market price that crosses this price from the direction
	// of the market price at the time when the Order was created (the initialMarketPrice).
	// Depending on the value of the Order’s price and initialMarketPrice, the MarketIfTouchedOrder
	// will behave like a Limit or a Stop Order.
	Price PriceValue `json:"price"`
	// The worst market price that may be used to fill this MarketIfTouched Order.
	PriceBound PriceValue `json:"priceBound"`
	// The time-in-force requested for the MarketIfTouched Order. Restricted to “GTC”,
	// “GFD” and “GTD” for MarketIfTouched Orders.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the MarketIfTouched Order will be cancelled if its timeInForce
	// is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Market-if-touched Order was initiated
	Reason MarketIfTouchedOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The specification of the Take Profit Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// The specification of the Stop Loss Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// The specification of the Trailing Stop Loss Order that should be created for a Trade
	// that is opened when the Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// The specification of the Guaranteed Stop Loss Order that should be created for a
	// Trade that is opened when the Order is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created).  Do not set, modify, delete tradeClientExtensions if your account
	// is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// The ID of the Order that this Order replaces (only provided if this Order replaces
	// an existing Order).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Transaction that cancels the replaced Order (only provided if this
	// Order replaces an existing Order).
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
}

// A MarketIfTouchedOrderRejectTransaction represents the rejection of the creation
// of a MarketIfTouched Order.
type MarketIfTouchedOrderRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “MARKET_IF_TOUCHED_ORDER_REJECT”
	// in a MarketIfTouchedOrderRejectTransaction.
	Type TransactionType `json:"type"`
	// The MarketIfTouched Order’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The quantity requested to be filled by the MarketIfTouched Order. A positive number
	// of units results in a long Order, and a negative number of units results in a short
	// Order.
	Units DecimalNumber `json:"units"`
	// The price threshold specified for the MarketIfTouched Order. The MarketIfTouched
	// Order will only be filled by a market price that crosses this price from the direction
	// of the market price at the time when the Order was created (the initialMarketPrice).
	// Depending on the value of the Order’s price and initialMarketPrice, the MarketIfTouchedOrder
	// will behave like a Limit or a Stop Order.
	Price PriceValue `json:"price"`
	// The worst market price that may be used to fill this MarketIfTouched Order.
	PriceBound PriceValue `json:"priceBound"`
	// The time-in-force requested for the MarketIfTouched Order. Restricted to “GTC”,
	// “GFD” and “GTD” for MarketIfTouched Orders.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the MarketIfTouched Order will be cancelled if its timeInForce
	// is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of how Positions in the Account are modified when the Order is filled.
	PositionFill OrderPositionFill `json:"positionFill"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Market-if-touched Order was initiated
	Reason MarketIfTouchedOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The specification of the Take Profit Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// The specification of the Stop Loss Order that should be created for a Trade opened
	// when the Order is filled (if such a Trade is created).
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// The specification of the Trailing Stop Loss Order that should be created for a Trade
	// that is opened when the Order is filled (if such a Trade is created).
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// The specification of the Guaranteed Stop Loss Order that should be created for a
	// Trade that is opened when the Order is filled (if such a Trade is created).
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created).  Do not set, modify, delete tradeClientExtensions if your account
	// is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// The ID of the Order that this Order was intended to replace (only provided if this
	// Order was intended to replace an existing Order).
	IntendedReplacesOrderID OrderID `json:"intendedReplacesOrderID"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A TakeProfitOrderTransaction represents the creation of a TakeProfit Order in the
// user’s Account.
type TakeProfitOrderTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “TAKE_PROFIT_ORDER” in a TakeProfitOrderTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`
	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`
	// The price threshold specified for the TakeProfit Order. The associated Trade will
	// be closed by a market price that is equal to or better than this threshold.
	Price PriceValue `json:"price"`
	// The time-in-force requested for the TakeProfit Order. Restricted to “GTC”, “GFD”
	// and “GTD” for TakeProfit Orders.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the TakeProfit Order will be cancelled if its timeInForce is
	// “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Take Profit Order was initiated
	Reason TakeProfitOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The ID of the OrderFill Transaction that caused this Order to be created (only provided
	// if this Order was created automatically when another Order was filled).
	OrderFillTransactionID TransactionID `json:"orderFillTransactionID"`
	// The ID of the Order that this Order replaces (only provided if this Order replaces
	// an existing Order).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Transaction that cancels the replaced Order (only provided if this
	// Order replaces an existing Order).
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
}

// A TakeProfitOrderRejectTransaction represents the rejection of the creation of a
// TakeProfit Order.
type TakeProfitOrderRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “TAKE_PROFIT_ORDER_REJECT” in a TakeProfitOrderRejectTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`
	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`
	// The price threshold specified for the TakeProfit Order. The associated Trade will
	// be closed by a market price that is equal to or better than this threshold.
	Price PriceValue `json:"price"`
	// The time-in-force requested for the TakeProfit Order. Restricted to “GTC”, “GFD”
	// and “GTD” for TakeProfit Orders.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the TakeProfit Order will be cancelled if its timeInForce is
	// “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Take Profit Order was initiated
	Reason TakeProfitOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The ID of the OrderFill Transaction that caused this Order to be created (only provided
	// if this Order was created automatically when another Order was filled).
	OrderFillTransactionID TransactionID `json:"orderFillTransactionID"`
	// The ID of the Order that this Order was intended to replace (only provided if this
	// Order was intended to replace an existing Order).
	IntendedReplacesOrderID OrderID `json:"intendedReplacesOrderID"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A StopLossOrderTransaction represents the creation of a StopLoss Order in the user’s
// Account.
type StopLossOrderTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “STOP_LOSS_ORDER” in a StopLossOrderTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`
	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`
	// The price threshold specified for the Stop Loss Order. The associated Trade will
	// be closed by a market price that is equal to or worse than this threshold.
	Price PriceValue `json:"price"`
	// Specifies the distance (in price units) from the Account’s current price to use
	// as the Stop Loss Order price. If the Trade is short the Instrument’s bid price
	// is used, and for long Trades the ask is used.
	Distance DecimalNumber `json:"distance"`
	// The time-in-force requested for the StopLoss Order. Restricted to “GTC”, “GFD”
	// and “GTD” for StopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the StopLoss Order will be cancelled if its timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// Flag indicating that the Stop Loss Order is guaranteed. The default value depends
	// on the GuaranteedStopLossOrderMode of the account, if it is REQUIRED, the default
	// will be true, for DISABLED or ENABLED the default is false.   Deprecated: Will be
	// removed in a future API update.
	Guaranteed bool `json:"guaranteed"`
	// The fee that will be charged if the Stop Loss Order is guaranteed and the Order
	// is filled at the guaranteed price. The value is determined at Order creation time.
	// It is in price units and is charged for each unit of the Trade.   Deprecated: Will
	// be removed in a future API update.
	GuaranteedExecutionPremium DecimalNumber `json:"guaranteedExecutionPremium"`
	// The reason that the Stop Loss Order was initiated
	Reason StopLossOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The ID of the OrderFill Transaction that caused this Order to be created (only provided
	// if this Order was created automatically when another Order was filled).
	OrderFillTransactionID TransactionID `json:"orderFillTransactionID"`
	// The ID of the Order that this Order replaces (only provided if this Order replaces
	// an existing Order).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Transaction that cancels the replaced Order (only provided if this
	// Order replaces an existing Order).
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
}

// A StopLossOrderRejectTransaction represents the rejection of the creation of a StopLoss
// Order.
type StopLossOrderRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “STOP_LOSS_ORDER_REJECT” in a StopLossOrderRejectTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`
	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`
	// The price threshold specified for the Stop Loss Order. The associated Trade will
	// be closed by a market price that is equal to or worse than this threshold.
	Price PriceValue `json:"price"`
	// Specifies the distance (in price units) from the Account’s current price to use
	// as the Stop Loss Order price. If the Trade is short the Instrument’s bid price
	// is used, and for long Trades the ask is used.
	Distance DecimalNumber `json:"distance"`
	// The time-in-force requested for the StopLoss Order. Restricted to “GTC”, “GFD”
	// and “GTD” for StopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the StopLoss Order will be cancelled if its timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// Flag indicating that the Stop Loss Order is guaranteed. The default value depends
	// on the GuaranteedStopLossOrderMode of the account, if it is REQUIRED, the default
	// will be true, for DISABLED or ENABLED the default is false.   Deprecated: Will be
	// removed in a future API update.
	Guaranteed bool `json:"guaranteed"`
	// The reason that the Stop Loss Order was initiated
	Reason StopLossOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The ID of the OrderFill Transaction that caused this Order to be created (only provided
	// if this Order was created automatically when another Order was filled).
	OrderFillTransactionID TransactionID `json:"orderFillTransactionID"`
	// The ID of the Order that this Order was intended to replace (only provided if this
	// Order was intended to replace an existing Order).
	IntendedReplacesOrderID OrderID `json:"intendedReplacesOrderID"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A GuaranteedStopLossOrderTransaction represents the creation of a GuaranteedStopLoss
// Order in the user’s Account.
type GuaranteedStopLossOrderTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “GUARANTEED_STOP_LOSS_ORDER” in a
	// GuaranteedStopLossOrderTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`
	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`
	// The price threshold specified for the Guaranteed Stop Loss Order. The associated
	// Trade will be closed at this price.
	Price PriceValue `json:"price"`
	// Specifies the distance (in price units) from the Account’s current price to use
	// as the Guaranteed Stop Loss Order price. If the Trade is short the Instrument’s
	// bid price is used, and for long Trades the ask is used.
	Distance DecimalNumber `json:"distance"`
	// The time-in-force requested for the GuaranteedStopLoss Order. Restricted to “GTC”,
	// “GFD” and “GTD” for GuaranteedStopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the GuaranteedStopLoss Order will be cancelled if its timeInForce
	// is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The fee that will be charged if the Guaranteed Stop Loss Order is filled at the
	// guaranteed price. The value is determined at Order creation time. It is in price
	// units and is charged for each unit of the Trade.
	GuaranteedExecutionPremium DecimalNumber `json:"guaranteedExecutionPremium"`
	// The reason that the Guaranteed Stop Loss Order was initiated
	Reason GuaranteedStopLossOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The ID of the OrderFill Transaction that caused this Order to be created (only provided
	// if this Order was created automatically when another Order was filled).
	OrderFillTransactionID TransactionID `json:"orderFillTransactionID"`
	// The ID of the Order that this Order replaces (only provided if this Order replaces
	// an existing Order).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Transaction that cancels the replaced Order (only provided if this
	// Order replaces an existing Order).
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
}

// A GuaranteedStopLossOrderRejectTransaction represents the rejection of the creation
// of a GuaranteedStopLoss Order.
type GuaranteedStopLossOrderRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “GUARANTEED_STOP_LOSS_ORDER_REJECT”
	// in a GuaranteedStopLossOrderRejectTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`
	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`
	// The price threshold specified for the Guaranteed Stop Loss Order. The associated
	// Trade will be closed at this price.
	Price PriceValue `json:"price"`
	// Specifies the distance (in price units) from the Account’s current price to use
	// as the Guaranteed Stop Loss Order price. If the Trade is short the Instrument’s
	// bid price is used, and for long Trades the ask is used.
	Distance DecimalNumber `json:"distance"`
	// The time-in-force requested for the GuaranteedStopLoss Order. Restricted to “GTC”,
	// “GFD” and “GTD” for GuaranteedStopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the GuaranteedStopLoss Order will be cancelled if its timeInForce
	// is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Guaranteed Stop Loss Order was initiated
	Reason GuaranteedStopLossOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The ID of the OrderFill Transaction that caused this Order to be created (only provided
	// if this Order was created automatically when another Order was filled).
	OrderFillTransactionID TransactionID `json:"orderFillTransactionID"`
	// The ID of the Order that this Order was intended to replace (only provided if this
	// Order was intended to replace an existing Order).
	IntendedReplacesOrderID OrderID `json:"intendedReplacesOrderID"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A TrailingStopLossOrderTransaction represents the creation of a TrailingStopLoss
// Order in the user’s Account.
type TrailingStopLossOrderTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “TRAILING_STOP_LOSS_ORDER” in a TrailingStopLossOrderTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`
	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`
	// The price distance (in price units) specified for the TrailingStopLoss Order.
	Distance DecimalNumber `json:"distance"`
	// The time-in-force requested for the TrailingStopLoss Order. Restricted to “GTC”,
	// “GFD” and “GTD” for TrailingStopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the StopLoss Order will be cancelled if its timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Trailing Stop Loss Order was initiated
	Reason TrailingStopLossOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The ID of the OrderFill Transaction that caused this Order to be created (only provided
	// if this Order was created automatically when another Order was filled).
	OrderFillTransactionID TransactionID `json:"orderFillTransactionID"`
	// The ID of the Order that this Order replaces (only provided if this Order replaces
	// an existing Order).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Transaction that cancels the replaced Order (only provided if this
	// Order replaces an existing Order).
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
}

// A TrailingStopLossOrderRejectTransaction represents the rejection of the creation
// of a TrailingStopLoss Order.
type TrailingStopLossOrderRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “TRAILING_STOP_LOSS_ORDER_REJECT”
	// in a TrailingStopLossOrderRejectTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Trade to close when the price threshold is breached.
	TradeID TradeID `json:"tradeID"`
	// The client ID of the Trade to be closed when the price threshold is breached.
	ClientTradeID ClientID `json:"clientTradeID"`
	// The price distance (in price units) specified for the TrailingStopLoss Order.
	Distance DecimalNumber `json:"distance"`
	// The time-in-force requested for the TrailingStopLoss Order. Restricted to “GTC”,
	// “GFD” and “GTD” for TrailingStopLoss Orders.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date/time when the StopLoss Order will be cancelled if its timeInForce is “GTD”.
	GtdTime DateTime `json:"gtdTime"`
	// Specification of which price component should be used when determining if an Order
	// should be triggered and filled. This allows Orders to be triggered based on the bid,
	// ask, mid, default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy)
	// price depending on the desired behaviour. Orders are always filled using their default
	// price component. This feature is only provided through the REST API. Clients who
	// choose to specify a non-default trigger condition will not see it reflected in any
	// of OANDA’s proprietary or partner trading platforms, their transaction history
	// or their account statements. OANDA platforms always assume that an Order’s trigger
	// condition is set to the default value when indicating the distance from an Order’s
	// trigger price, and will always provide the default trigger condition when creating
	// or modifying an Order. A special restriction applies when creating a Guaranteed Stop
	// Loss Order. In this case the TriggerCondition value must either be “DEFAULT”,
	// or the “natural” trigger side “DEFAULT” results in. So for a Guaranteed Stop
	// Loss Order for a long trade valid values are “DEFAULT” and “BID”, and for
	// short trades “DEFAULT” and “ASK” are valid.
	TriggerCondition OrderTriggerCondition `json:"triggerCondition"`
	// The reason that the Trailing Stop Loss Order was initiated
	Reason TrailingStopLossOrderReason `json:"reason"`
	// Client Extensions to add to the Order (only provided if the Order is being created
	// with client extensions).
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The ID of the OrderFill Transaction that caused this Order to be created (only provided
	// if this Order was created automatically when another Order was filled).
	OrderFillTransactionID TransactionID `json:"orderFillTransactionID"`
	// The ID of the Order that this Order was intended to replace (only provided if this
	// Order was intended to replace an existing Order).
	IntendedReplacesOrderID OrderID `json:"intendedReplacesOrderID"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// An OrderFillTransaction represents the filling of an Order in the client’s Account.
type OrderFillTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “ORDER_FILL” for an OrderFillTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Order filled.
	OrderID OrderID `json:"orderID"`
	// The client Order ID of the Order filled (only provided if the client has assigned
	// one).
	ClientOrderID ClientID `json:"clientOrderID"`
	// The name of the filled Order’s instrument.
	Instrument InstrumentName `json:"instrument"`
	// The number of units filled by the OrderFill.
	Units DecimalNumber `json:"units"`
	// This is the conversion factor in effect for the Account at the time of the OrderFill
	// for converting any gains realized in Instrument quote units into units of the Account’s
	// home currency.   Deprecated: Will be removed in a future API update.
	GainQuoteHomeConversionFactor DecimalNumber `json:"gainQuoteHomeConversionFactor"`
	// This is the conversion factor in effect for the Account at the time of the OrderFill
	// for converting any losses realized in Instrument quote units into units of the Account’s
	// home currency.   Deprecated: Will be removed in a future API update.
	LossQuoteHomeConversionFactor DecimalNumber `json:"lossQuoteHomeConversionFactor"`
	// The HomeConversionFactors in effect at the time of the OrderFill.
	HomeConversionFactors *HomeConversionFactors `json:"homeConversionFactors"`
	// This field is now deprecated and should no longer be used. The individual tradesClosed,
	// tradeReduced and tradeOpened fields contain the exact/official price each unit was
	// filled at.   Deprecated: Will be removed in a future API update.
	Price PriceValue `json:"price"`
	// The price that all of the units of the OrderFill should have been filled at, in
	// the absence of guaranteed price execution. This factors in the Account’s current
	// ClientPrice, used liquidity and the units of the OrderFill only. If no Trades were
	// closed with their price clamped for guaranteed stop loss enforcement, then this value
	// will match the price fields of each Trade opened, closed, and reduced, and they will
	// all be the exact same.
	FullVWAP PriceValue `json:"fullVWAP"`
	// The price in effect for the account at the time of the Order fill.
	FullPrice ClientPrice `json:"fullPrice"`
	// The reason that an Order was filled
	Reason OrderFillReason `json:"reason"`
	// The profit or loss incurred when the Order was filled.
	Pl AccountUnits `json:"pl"`
	// The profit or loss incurred when the Order was filled, in the Instrument’s quote
	// currency.
	QuotePL DecimalNumber `json:"quotePL"`
	// The financing paid or collected when the Order was filled.
	Financing AccountUnits `json:"financing"`
	// The financing paid or collected when the Order was filled, in the Instrument’s
	// base currency.
	BaseFinancing DecimalNumber `json:"baseFinancing"`
	// The financing paid or collected when the Order was filled, in the Instrument’s
	// quote currency.
	QuoteFinancing DecimalNumber `json:"quoteFinancing"`
	// The commission charged in the Account’s home currency as a result of filling the
	// Order. The commission is always represented as a positive quantity of the Account’s
	// home currency, however it reduces the balance in the Account.
	Commission AccountUnits `json:"commission"`
	// The total guaranteed execution fees charged for all Trades opened, closed or reduced
	// with guaranteed Stop Loss Orders.
	GuaranteedExecutionFee AccountUnits `json:"guaranteedExecutionFee"`
	// The total guaranteed execution fees charged for all Trades opened, closed or reduced
	// with guaranteed Stop Loss Orders, expressed in the Instrument’s quote currency.
	QuoteGuaranteedExecutionFee DecimalNumber `json:"quoteGuaranteedExecutionFee"`
	// The Account’s balance after the Order was filled.
	AccountBalance AccountUnits `json:"accountBalance"`
	// The Trade that was opened when the Order was filled (only provided if filling the
	// Order resulted in a new Trade).
	TradeOpened *TradeOpen `json:"tradeOpened"`
	// The Trades that were closed when the Order was filled (only provided if filling
	// the Order resulted in a closing open Trades).
	TradesClosed []*TradeReduce `json:"tradesClosed"`
	// The Trade that was reduced when the Order was filled (only provided if filling the
	// Order resulted in reducing an open Trade).
	TradeReduced *TradeReduce `json:"tradeReduced"`
	// The half spread cost for the OrderFill, which is the sum of the halfSpreadCost values
	// in the tradeOpened, tradesClosed and tradeReduced fields. This can be a positive
	// or negative value and is represented in the home currency of the Account.
	HalfSpreadCost AccountUnits `json:"halfSpreadCost"`
}

// An OrderCancelTransaction represents the cancellation of an Order in the client’s
// Account.
type OrderCancelTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “ORDER_CANCEL” for an OrderCancelTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Order cancelled
	OrderID OrderID `json:"orderID"`
	// The client ID of the Order cancelled (only provided if the Order has a client Order
	// ID).
	ClientOrderID OrderID `json:"clientOrderID"`
	// The reason that the Order was cancelled.
	Reason OrderCancelReason `json:"reason"`
	// The ID of the Order that replaced this Order (only provided if this Order was cancelled
	// for replacement).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

// An OrderCancelRejectTransaction represents the rejection of the cancellation of an
// Order in the client’s Account.
type OrderCancelRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “ORDER_CANCEL_REJECT” for an OrderCancelRejectTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Order intended to be cancelled
	OrderID OrderID `json:"orderID"`
	// The client ID of the Order intended to be cancelled (only provided if the Order
	// has a client Order ID).
	ClientOrderID OrderID `json:"clientOrderID"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A OrderClientExtensionsModifyTransaction represents the modification of an Order’s
// Client Extensions.
type OrderClientExtensionsModifyTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “ORDER_CLIENT_EXTENSIONS_MODIFY”
	// for a OrderClientExtensionsModifyTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Order who’s client extensions are to be modified.
	OrderID OrderID `json:"orderID"`
	// The original Client ID of the Order who’s client extensions are to be modified.
	ClientOrderID ClientID `json:"clientOrderID"`
	// The new Client Extensions for the Order.
	ClientExtensionsModify *ClientExtensions `json:"clientExtensionsModify"`
	// The new Client Extensions for the Order’s Trade on fill.
	TradeClientExtensionsModify *ClientExtensions `json:"tradeClientExtensionsModify"`
}

// A OrderClientExtensionsModifyRejectTransaction represents the rejection of the modification
// of an Order’s Client Extensions.
type OrderClientExtensionsModifyRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT”
	// for a OrderClientExtensionsModifyRejectTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Order who’s client extensions are to be modified.
	OrderID OrderID `json:"orderID"`
	// The original Client ID of the Order who’s client extensions are to be modified.
	ClientOrderID ClientID `json:"clientOrderID"`
	// The new Client Extensions for the Order.
	ClientExtensionsModify *ClientExtensions `json:"clientExtensionsModify"`
	// The new Client Extensions for the Order’s Trade on fill.
	TradeClientExtensionsModify *ClientExtensions `json:"tradeClientExtensionsModify"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A TradeClientExtensionsModifyTransaction represents the modification of a Trade’s
// Client Extensions.
type TradeClientExtensionsModifyTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “TRADE_CLIENT_EXTENSIONS_MODIFY”
	// for a TradeClientExtensionsModifyTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Trade who’s client extensions are to be modified.
	TradeID TradeID `json:"tradeID"`
	// The original Client ID of the Trade who’s client extensions are to be modified.
	ClientTradeID ClientID `json:"clientTradeID"`
	// The new Client Extensions for the Trade.
	TradeClientExtensionsModify *ClientExtensions `json:"tradeClientExtensionsModify"`
}

// A TradeClientExtensionsModifyRejectTransaction represents the rejection of the modification
// of a Trade’s Client Extensions.
type TradeClientExtensionsModifyRejectTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT”
	// for a TradeClientExtensionsModifyRejectTransaction.
	Type TransactionType `json:"type"`
	// The ID of the Trade who’s client extensions are to be modified.
	TradeID TradeID `json:"tradeID"`
	// The original Client ID of the Trade who’s client extensions are to be modified.
	ClientTradeID ClientID `json:"clientTradeID"`
	// The new Client Extensions for the Trade.
	TradeClientExtensionsModify *ClientExtensions `json:"tradeClientExtensionsModify"`
	// The reason that the Reject Transaction was created
	RejectReason TransactionRejectReason `json:"rejectReason"`
}

// A MarginCallEnterTransaction is created when an Account enters the margin call state.
type MarginCallEnterTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “MARGIN_CALL_ENTER” for an MarginCallEnterTransaction.
	Type TransactionType `json:"type"`
}

// A MarginCallExtendTransaction is created when the margin call state for an Account
// has been extended.
type MarginCallExtendTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “MARGIN_CALL_EXTEND” for an MarginCallExtendTransaction.
	Type TransactionType `json:"type"`
	// The number of the extensions to the Account’s current margin call that have been
	// applied. This value will be set to 1 for the first MarginCallExtend Transaction
	ExtensionNumber int64 `json:"extensionNumber"`
}

// A MarginCallExitTransaction is created when an Account leaves the margin call state.
type MarginCallExitTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “MARGIN_CALL_EXIT” for an MarginCallExitTransaction.
	Type TransactionType `json:"type"`
}

// A DelayedTradeClosure Transaction is created administratively to indicate open trades
// that should have been closed but weren’t because the open trades’ instruments
// were untradeable at the time. Open trades listed in this transaction will be closed
// once their respective instruments become tradeable.
type DelayedTradeClosureTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “DELAYED_TRADE_CLOSURE” for an DelayedTradeClosureTransaction.
	Type TransactionType `json:"type"`
	// The reason for the delayed trade closure
	Reason MarketOrderReason `json:"reason"`
	// List of Trade ID’s identifying the open trades that will be closed when their
	// respective instruments become tradeable
	TradeIDs TradeID `json:"tradeIDs"`
}

// A DailyFinancingTransaction represents the daily payment/collection of financing
// for an Account.
type DailyFinancingTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “DAILY_FINANCING” for a DailyFinancingTransaction.
	Type TransactionType `json:"type"`
	// The amount of financing paid/collected for the Account.
	Financing AccountUnits `json:"financing"`
	// The Account’s balance after daily financing.
	AccountBalance AccountUnits `json:"accountBalance"`
	// The account financing mode at the time of the daily financing. This field is no
	// longer in use moving forward and was replaced by accountFinancingMode in individual
	// positionFinancings since the financing mode could differ between instruments.   Deprecated:
	// Will be removed in a future API update.
	AccountFinancingMode AccountFinancingMode `json:"accountFinancingMode"`
	// The financing paid/collected for each Position in the Account.
	PositionFinancings []*PositionFinancing `json:"positionFinancings"`
}

// A DividendAdjustment Transaction is created administratively to pay or collect dividend
// adjustment mounts to or from an Account.
type DividendAdjustmentTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “DIVIDEND_ADJUSTMENT” for a DividendAdjustmentTransaction.
	Type TransactionType `json:"type"`
	// The name of the instrument for the dividendAdjustment transaction
	Instrument InstrumentName `json:"instrument"`
	// The total dividend adjustment amount paid or collected in the Account’s home currency
	// for the Account as a result of applying the DividendAdjustment Transaction. This
	// is the sum of the dividend adjustments paid/collected for each OpenTradeDividendAdjustment
	// found within the Transaction.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`
	// The total dividend adjustment amount paid or collected in the Instrument’s quote
	// currency for the Account as a result of applying the DividendAdjustment Transaction.
	// This is the sum of the quote dividend adjustments paid/collected for each OpenTradeDividendAdjustment
	// found within the Transaction.
	QuoteDividendAdjustment DecimalNumber `json:"quoteDividendAdjustment"`
	// The HomeConversionFactors in effect at the time of the DividendAdjustment.
	HomeConversionFactors *HomeConversionFactors `json:"homeConversionFactors"`
	// The Account balance after applying the DividendAdjustment Transaction
	AccountBalance AccountUnits `json:"accountBalance"`
	// The dividend adjustment payment/collection details for each open Trade, within the
	// Account, for which a dividend adjustment is to be paid or collected.
	OpenTradeDividendAdjustments []*OpenTradeDividendAdjustment `json:"openTradeDividendAdjustments"`
}

// A ResetResettablePLTransaction represents the resetting of the Account’s resettable
// PL counters.
type ResetResettablePLTransaction struct {
	Transaction
	// The Type of the Transaction. Always set to “RESET_RESETTABLE_PL” for a ResetResettablePLTransaction.
	Type TransactionType `json:"type"`
}

type TransactionParser struct {
	AccountBalance                AccountUnits                   `json:"accountBalance"`
	AccountFinancingMode          AccountFinancingMode           `json:"accountFinancingMode"`
	AccountID                     string                         `json:"accountID"`
	AccountNumber                 int64                          `json:"accountNumber"`
	AccountUserID                 int64                          `json:"accountUserID"`
	Alias                         string                         `json:"alias"`
	Amount                        AccountUnits                   `json:"amount"`
	BaseFinancing                 DecimalNumber                  `json:"baseFinancing"`
	BatchID                       string                         `json:"batchID"`
	CancellingTransactionID       string                         `json:"cancellingTransactionID"`
	ClientExtensions              *ClientExtensions              `json:"clientExtensions"`
	ClientExtensionsModify        *ClientExtensions              `json:"clientExtensionsModify"`
	ClientOrderID                 string                         `json:"clientOrderID"`
	ClientTradeID                 string                         `json:"clientTradeID"`
	Comment                       string                         `json:"comment"`
	Commission                    AccountUnits                   `json:"commission"`
	DelayedTradeClose             *MarketOrderDelayedTradeClose  `json:"delayedTradeClose"`
	Distance                      DecimalNumber                  `json:"distance"`
	DividendAdjustment            AccountUnits                   `json:"dividendAdjustment"`
	DivisionID                    int64                          `json:"divisionID"`
	ExtensionNumber               int64                          `json:"extensionNumber"`
	Financing                     AccountUnits                   `json:"financing"`
	FullPrice                     ClientPrice                    `json:"fullPrice"`
	FullVWAP                      PriceValue                     `json:"fullVWAP"`
	FundingReason                 string                         `json:"fundingReason"`
	GainQuoteHomeConversionFactor DecimalNumber                  `json:"gainQuoteHomeConversionFactor"`
	GtdTime                       DateTime                       `json:"gtdTime"`
	Guaranteed                    bool                           `json:"guaranteed"`
	GuaranteedExecutionFee        AccountUnits                   `json:"guaranteedExecutionFee"`
	GuaranteedExecutionPremium    DecimalNumber                  `json:"guaranteedExecutionPremium"`
	GuaranteedStopLossOnFill      *GuaranteedStopLossDetails     `json:"guaranteedStopLossOnFill"`
	HalfSpreadCost                AccountUnits                   `json:"halfSpreadCost"`
	HomeConversionFactors         *HomeConversionFactors         `json:"homeConversionFactors"`
	HomeCurrency                  Currency                       `json:"homeCurrency"`
	Id                            string                         `json:"id"`
	Instrument                    InstrumentName                 `json:"instrument"`
	IntendedReplacesOrderID       string                         `json:"intendedReplacesOrderID"`
	LongPositionCloseout          *MarketOrderPositionCloseout   `json:"longPositionCloseout"`
	LossQuoteHomeConversionFactor DecimalNumber                  `json:"lossQuoteHomeConversionFactor"`
	MarginCloseout                *MarketOrderMarginCloseout     `json:"marginCloseout"`
	MarginRate                    DecimalNumber                  `json:"marginRate"`
	OpenTradeDividendAdjustments  []*OpenTradeDividendAdjustment `json:"openTradeDividendAdjustments"`
	OrderFillTransactionID        string                         `json:"orderFillTransactionID"`
	OrderID                       string                         `json:"orderID"`
	Pl                            AccountUnits                   `json:"pl"`
	PositionFill                  OrderPositionFill              `json:"positionFill"`
	PositionFinancings            []*PositionFinancing           `json:"positionFinancings"`
	Price                         PriceValue                     `json:"price"`
	PriceBound                    PriceValue                     `json:"priceBound"`
	QuoteDividendAdjustment       DecimalNumber                  `json:"quoteDividendAdjustment"`
	QuoteFinancing                DecimalNumber                  `json:"quoteFinancing"`
	QuoteGuaranteedExecutionFee   DecimalNumber                  `json:"quoteGuaranteedExecutionFee"`
	QuotePL                       DecimalNumber                  `json:"quotePL"`
	Reason                        string                         `json:"reason"`
	RejectReason                  string                         `json:"rejectReason"`
	ReplacedByOrderID             string                         `json:"replacedByOrderID"`
	ReplacesOrderID               string                         `json:"replacesOrderID"`
	RequestID                     string                         `json:"requestID"`
	ShortPositionCloseout         *MarketOrderPositionCloseout   `json:"shortPositionCloseout"`
	SiteID                        int64                          `json:"siteID"`
	StopLossOnFill                *StopLossDetails               `json:"stopLossOnFill"`
	TakeProfitOnFill              *TakeProfitDetails             `json:"takeProfitOnFill"`
	Time                          DateTime                       `json:"time"`
	TimeInForce                   TimeInForce                    `json:"timeInForce"`
	TradeClientExtensions         *ClientExtensions              `json:"tradeClientExtensions"`
	TradeClientExtensionsModify   *ClientExtensions              `json:"tradeClientExtensionsModify"`
	TradeClose                    *MarketOrderTradeClose         `json:"tradeClose"`
	TradeID                       string                         `json:"tradeID"`
	TradeIDs                      string                         `json:"tradeIDs"`
	TradeOpened                   *TradeOpen                     `json:"tradeOpened"`
	TradeReduced                  *TradeReduce                   `json:"tradeReduced"`
	TradeState                    string                         `json:"tradeState"`
	TradesClosed                  []*TradeReduce                 `json:"tradesClosed"`
	TrailingStopLossOnFill        *TrailingStopLossDetails       `json:"trailingStopLossOnFill"`
	TriggerCondition              string                         `json:"triggerCondition"`
	Type                          string                         `json:"type"`
	Units                         DecimalNumber                  `json:"units"`
	UserID                        int64                          `json:"userID"`
}

// Example
/*
r := parser.Parse()
switch v := r.(type) {
case *CreateTransaction:
case *CloseTransaction:
case *ReopenTransaction:
case *ClientConfigureTransaction:
case *ClientConfigureRejectTransaction:
case *TransferFundsTransaction:
case *TransferFundsRejectTransaction:
case *MarketOrderTransaction:
case *MarketOrderRejectTransaction:
case *FixedPriceOrderTransaction:
case *LimitOrderTransaction:
case *LimitOrderRejectTransaction:
case *StopOrderTransaction:
case *StopOrderRejectTransaction:
case *MarketIfTouchedOrderTransaction:
case *MarketIfTouchedOrderRejectTransaction:
case *TakeProfitOrderTransaction:
case *TakeProfitOrderRejectTransaction:
case *StopLossOrderTransaction:
case *StopLossOrderRejectTransaction:
case *GuaranteedStopLossOrderTransaction:
case *GuaranteedStopLossOrderRejectTransaction:
case *TrailingStopLossOrderTransaction:
case *TrailingStopLossOrderRejectTransaction:
case *OrderFillTransaction:
case *OrderCancelTransaction:
case *OrderCancelRejectTransaction:
case *OrderClientExtensionsModifyTransaction:
case *OrderClientExtensionsModifyRejectTransaction:
case *TradeClientExtensionsModifyTransaction:
case *TradeClientExtensionsModifyRejectTransaction:
case *MarginCallEnterTransaction:
case *MarginCallExtendTransaction:
case *MarginCallExitTransaction:
case *DelayedTradeClosureTransaction:
case *DailyFinancingTransaction:
case *DividendAdjustmentTransaction:
case *ResetResettablePLTransaction:
}
*/
func (p *TransactionParser) Parse() interface{} {
	switch p.Type {
	case "CREATE":
		return &CreateTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:          TransactionType(p.Type),
			DivisionID:    p.DivisionID,
			SiteID:        p.SiteID,
			AccountUserID: p.AccountUserID,
			AccountNumber: p.AccountNumber,
			HomeCurrency:  p.HomeCurrency,
		}
	case "CLOSE":
		return &CloseTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type: TransactionType(p.Type),
		}
	case "REOPEN":
		return &ReopenTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type: TransactionType(p.Type),
		}
	case "CLIENT_CONFIGURE":
		return &ClientConfigureTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:       TransactionType(p.Type),
			Alias:      p.Alias,
			MarginRate: p.MarginRate,
		}
	case "CLIENT_CONFIGURE_REJECT":
		return &ClientConfigureRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:         TransactionType(p.Type),
			Alias:        p.Alias,
			MarginRate:   p.MarginRate,
			RejectReason: TransactionRejectReason(p.RejectReason),
		}
	case "TRANSFER_FUNDS":
		return &TransferFundsTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:           TransactionType(p.Type),
			Amount:         p.Amount,
			FundingReason:  FundingReason(p.FundingReason),
			Comment:        p.Comment,
			AccountBalance: p.AccountBalance,
		}
	case "TRANSFER_FUNDS_REJECT":
		return &TransferFundsRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:          TransactionType(p.Type),
			Amount:        p.Amount,
			FundingReason: FundingReason(p.FundingReason),
			Comment:       p.Comment,
			RejectReason:  TransactionRejectReason(p.RejectReason),
		}
	case "MARKET_ORDER":
		return &MarketOrderTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                     TransactionType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			TimeInForce:              p.TimeInForce,
			PriceBound:               p.PriceBound,
			PositionFill:             p.PositionFill,
			TradeClose:               p.TradeClose,
			LongPositionCloseout:     p.LongPositionCloseout,
			ShortPositionCloseout:    p.ShortPositionCloseout,
			MarginCloseout:           p.MarginCloseout,
			DelayedTradeClose:        p.DelayedTradeClose,
			Reason:                   MarketOrderReason(p.Reason),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
		}
	case "MARKET_ORDER_REJECT":
		return &MarketOrderRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                     TransactionType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			TimeInForce:              p.TimeInForce,
			PriceBound:               p.PriceBound,
			PositionFill:             p.PositionFill,
			TradeClose:               p.TradeClose,
			LongPositionCloseout:     p.LongPositionCloseout,
			ShortPositionCloseout:    p.ShortPositionCloseout,
			MarginCloseout:           p.MarginCloseout,
			DelayedTradeClose:        p.DelayedTradeClose,
			Reason:                   MarketOrderReason(p.Reason),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			RejectReason:             TransactionRejectReason(p.RejectReason),
		}
	case "FIXED_PRICE_ORDER":
		return &FixedPriceOrderTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                     TransactionType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			PositionFill:             p.PositionFill,
			TradeState:               p.TradeState,
			Reason:                   FixedPriceOrderReason(p.Reason),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
		}
	case "LIMIT_ORDER":
		return &LimitOrderTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                     TransactionType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			Reason:                   LimitOrderReason(p.Reason),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			ReplacesOrderID:          OrderID(p.ReplacesOrderID),
			CancellingTransactionID:  TransactionID(p.CancellingTransactionID),
		}
	case "LIMIT_ORDER_REJECT":
		return &LimitOrderRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                     TransactionType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			Reason:                   LimitOrderReason(p.Reason),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			IntendedReplacesOrderID:  OrderID(p.IntendedReplacesOrderID),
			RejectReason:             TransactionRejectReason(p.RejectReason),
		}
	case "STOP_ORDER":
		return &StopOrderTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                     TransactionType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			PriceBound:               p.PriceBound,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			Reason:                   StopOrderReason(p.Reason),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			ReplacesOrderID:          OrderID(p.ReplacesOrderID),
			CancellingTransactionID:  TransactionID(p.CancellingTransactionID),
		}
	case "STOP_ORDER_REJECT":
		return &StopOrderRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                     TransactionType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			PriceBound:               p.PriceBound,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			Reason:                   StopOrderReason(p.Reason),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			IntendedReplacesOrderID:  OrderID(p.IntendedReplacesOrderID),
			RejectReason:             TransactionRejectReason(p.RejectReason),
		}
	case "MARKET_IF_TOUCHED_ORDER":
		return &MarketIfTouchedOrderTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                     TransactionType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			PriceBound:               p.PriceBound,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			Reason:                   MarketIfTouchedOrderReason(p.Reason),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			ReplacesOrderID:          OrderID(p.ReplacesOrderID),
			CancellingTransactionID:  TransactionID(p.CancellingTransactionID),
		}
	case "MARKET_IF_TOUCHED_ORDER_REJECT":
		return &MarketIfTouchedOrderRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                     TransactionType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			PriceBound:               p.PriceBound,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			Reason:                   MarketIfTouchedOrderReason(p.Reason),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			IntendedReplacesOrderID:  OrderID(p.IntendedReplacesOrderID),
			RejectReason:             TransactionRejectReason(p.RejectReason),
		}
	case "TAKE_PROFIT_ORDER":
		return &TakeProfitOrderTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                    TransactionType(p.Type),
			TradeID:                 TradeID(p.TradeID),
			ClientTradeID:           ClientID(p.ClientTradeID),
			Price:                   p.Price,
			TimeInForce:             p.TimeInForce,
			GtdTime:                 p.GtdTime,
			TriggerCondition:        OrderTriggerCondition(p.TriggerCondition),
			Reason:                  TakeProfitOrderReason(p.Reason),
			ClientExtensions:        p.ClientExtensions,
			OrderFillTransactionID:  TransactionID(p.OrderFillTransactionID),
			ReplacesOrderID:         OrderID(p.ReplacesOrderID),
			CancellingTransactionID: TransactionID(p.CancellingTransactionID),
		}
	case "TAKE_PROFIT_ORDER_REJECT":
		return &TakeProfitOrderRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                    TransactionType(p.Type),
			TradeID:                 TradeID(p.TradeID),
			ClientTradeID:           ClientID(p.ClientTradeID),
			Price:                   p.Price,
			TimeInForce:             p.TimeInForce,
			GtdTime:                 p.GtdTime,
			TriggerCondition:        OrderTriggerCondition(p.TriggerCondition),
			Reason:                  TakeProfitOrderReason(p.Reason),
			ClientExtensions:        p.ClientExtensions,
			OrderFillTransactionID:  TransactionID(p.OrderFillTransactionID),
			IntendedReplacesOrderID: OrderID(p.IntendedReplacesOrderID),
			RejectReason:            TransactionRejectReason(p.RejectReason),
		}
	case "STOP_LOSS_ORDER":
		return &StopLossOrderTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                       TransactionType(p.Type),
			TradeID:                    TradeID(p.TradeID),
			ClientTradeID:              ClientID(p.ClientTradeID),
			Price:                      p.Price,
			Distance:                   p.Distance,
			TimeInForce:                p.TimeInForce,
			GtdTime:                    p.GtdTime,
			TriggerCondition:           OrderTriggerCondition(p.TriggerCondition),
			Guaranteed:                 p.Guaranteed,
			GuaranteedExecutionPremium: p.GuaranteedExecutionPremium,
			Reason:                     StopLossOrderReason(p.Reason),
			ClientExtensions:           p.ClientExtensions,
			OrderFillTransactionID:     TransactionID(p.OrderFillTransactionID),
			ReplacesOrderID:            OrderID(p.ReplacesOrderID),
			CancellingTransactionID:    TransactionID(p.CancellingTransactionID),
		}
	case "STOP_LOSS_ORDER_REJECT":
		return &StopLossOrderRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                    TransactionType(p.Type),
			TradeID:                 TradeID(p.TradeID),
			ClientTradeID:           ClientID(p.ClientTradeID),
			Price:                   p.Price,
			Distance:                p.Distance,
			TimeInForce:             p.TimeInForce,
			GtdTime:                 p.GtdTime,
			TriggerCondition:        OrderTriggerCondition(p.TriggerCondition),
			Guaranteed:              p.Guaranteed,
			Reason:                  StopLossOrderReason(p.Reason),
			ClientExtensions:        p.ClientExtensions,
			OrderFillTransactionID:  TransactionID(p.OrderFillTransactionID),
			IntendedReplacesOrderID: OrderID(p.IntendedReplacesOrderID),
			RejectReason:            TransactionRejectReason(p.RejectReason),
		}
	case "GUARANTEED_STOP_LOSS_ORDER":
		return &GuaranteedStopLossOrderTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                       TransactionType(p.Type),
			TradeID:                    TradeID(p.TradeID),
			ClientTradeID:              ClientID(p.ClientTradeID),
			Price:                      p.Price,
			Distance:                   p.Distance,
			TimeInForce:                p.TimeInForce,
			GtdTime:                    p.GtdTime,
			TriggerCondition:           OrderTriggerCondition(p.TriggerCondition),
			GuaranteedExecutionPremium: p.GuaranteedExecutionPremium,
			Reason:                     GuaranteedStopLossOrderReason(p.Reason),
			ClientExtensions:           p.ClientExtensions,
			OrderFillTransactionID:     TransactionID(p.OrderFillTransactionID),
			ReplacesOrderID:            OrderID(p.ReplacesOrderID),
			CancellingTransactionID:    TransactionID(p.CancellingTransactionID),
		}
	case "GUARANTEED_STOP_LOSS_ORDER_REJECT":
		return &GuaranteedStopLossOrderRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                    TransactionType(p.Type),
			TradeID:                 TradeID(p.TradeID),
			ClientTradeID:           ClientID(p.ClientTradeID),
			Price:                   p.Price,
			Distance:                p.Distance,
			TimeInForce:             p.TimeInForce,
			GtdTime:                 p.GtdTime,
			TriggerCondition:        OrderTriggerCondition(p.TriggerCondition),
			Reason:                  GuaranteedStopLossOrderReason(p.Reason),
			ClientExtensions:        p.ClientExtensions,
			OrderFillTransactionID:  TransactionID(p.OrderFillTransactionID),
			IntendedReplacesOrderID: OrderID(p.IntendedReplacesOrderID),
			RejectReason:            TransactionRejectReason(p.RejectReason),
		}
	case "TRAILING_STOP_LOSS_ORDER":
		return &TrailingStopLossOrderTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                    TransactionType(p.Type),
			TradeID:                 TradeID(p.TradeID),
			ClientTradeID:           ClientID(p.ClientTradeID),
			Distance:                p.Distance,
			TimeInForce:             p.TimeInForce,
			GtdTime:                 p.GtdTime,
			TriggerCondition:        OrderTriggerCondition(p.TriggerCondition),
			Reason:                  TrailingStopLossOrderReason(p.Reason),
			ClientExtensions:        p.ClientExtensions,
			OrderFillTransactionID:  TransactionID(p.OrderFillTransactionID),
			ReplacesOrderID:         OrderID(p.ReplacesOrderID),
			CancellingTransactionID: TransactionID(p.CancellingTransactionID),
		}
	case "TRAILING_STOP_LOSS_ORDER_REJECT":
		return &TrailingStopLossOrderRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                    TransactionType(p.Type),
			TradeID:                 TradeID(p.TradeID),
			ClientTradeID:           ClientID(p.ClientTradeID),
			Distance:                p.Distance,
			TimeInForce:             p.TimeInForce,
			GtdTime:                 p.GtdTime,
			TriggerCondition:        OrderTriggerCondition(p.TriggerCondition),
			Reason:                  TrailingStopLossOrderReason(p.Reason),
			ClientExtensions:        p.ClientExtensions,
			OrderFillTransactionID:  TransactionID(p.OrderFillTransactionID),
			IntendedReplacesOrderID: OrderID(p.IntendedReplacesOrderID),
			RejectReason:            TransactionRejectReason(p.RejectReason),
		}
	case "ORDER_FILL":
		return &OrderFillTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                          TransactionType(p.Type),
			OrderID:                       OrderID(p.OrderID),
			ClientOrderID:                 ClientID(p.ClientOrderID),
			Instrument:                    p.Instrument,
			Units:                         p.Units,
			GainQuoteHomeConversionFactor: p.GainQuoteHomeConversionFactor,
			LossQuoteHomeConversionFactor: p.LossQuoteHomeConversionFactor,
			HomeConversionFactors:         p.HomeConversionFactors,
			Price:                         p.Price,
			FullVWAP:                      p.FullVWAP,
			FullPrice:                     p.FullPrice,
			Reason:                        OrderFillReason(p.Reason),
			Pl:                            p.Pl,
			QuotePL:                       p.QuotePL,
			Financing:                     p.Financing,
			BaseFinancing:                 p.BaseFinancing,
			QuoteFinancing:                p.QuoteFinancing,
			Commission:                    p.Commission,
			GuaranteedExecutionFee:        p.GuaranteedExecutionFee,
			QuoteGuaranteedExecutionFee:   p.QuoteGuaranteedExecutionFee,
			AccountBalance:                p.AccountBalance,
			TradeOpened:                   p.TradeOpened,
			TradesClosed:                  p.TradesClosed,
			TradeReduced:                  p.TradeReduced,
			HalfSpreadCost:                p.HalfSpreadCost,
		}
	case "ORDER_CANCEL":
		return &OrderCancelTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:              TransactionType(p.Type),
			OrderID:           OrderID(p.OrderID),
			ClientOrderID:     OrderID(p.ClientOrderID),
			Reason:            OrderCancelReason(p.Reason),
			ReplacedByOrderID: OrderID(p.ReplacedByOrderID),
		}
	case "ORDER_CANCEL_REJECT":
		return &OrderCancelRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:          TransactionType(p.Type),
			OrderID:       OrderID(p.OrderID),
			ClientOrderID: OrderID(p.ClientOrderID),
			RejectReason:  TransactionRejectReason(p.RejectReason),
		}
	case "ORDER_CLIENT_EXTENSIONS_MODIFY":
		return &OrderClientExtensionsModifyTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                        TransactionType(p.Type),
			OrderID:                     OrderID(p.OrderID),
			ClientOrderID:               ClientID(p.ClientOrderID),
			ClientExtensionsModify:      p.ClientExtensionsModify,
			TradeClientExtensionsModify: p.TradeClientExtensionsModify,
		}
	case "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT":
		return &OrderClientExtensionsModifyRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                        TransactionType(p.Type),
			OrderID:                     OrderID(p.OrderID),
			ClientOrderID:               ClientID(p.ClientOrderID),
			ClientExtensionsModify:      p.ClientExtensionsModify,
			TradeClientExtensionsModify: p.TradeClientExtensionsModify,
			RejectReason:                TransactionRejectReason(p.RejectReason),
		}
	case "TRADE_CLIENT_EXTENSIONS_MODIFY":
		return &TradeClientExtensionsModifyTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                        TransactionType(p.Type),
			TradeID:                     TradeID(p.TradeID),
			ClientTradeID:               ClientID(p.ClientTradeID),
			TradeClientExtensionsModify: p.TradeClientExtensionsModify,
		}
	case "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT":
		return &TradeClientExtensionsModifyRejectTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                        TransactionType(p.Type),
			TradeID:                     TradeID(p.TradeID),
			ClientTradeID:               ClientID(p.ClientTradeID),
			TradeClientExtensionsModify: p.TradeClientExtensionsModify,
			RejectReason:                TransactionRejectReason(p.RejectReason),
		}
	case "MARGIN_CALL_ENTER":
		return &MarginCallEnterTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type: TransactionType(p.Type),
		}
	case "MARGIN_CALL_EXTEND":
		return &MarginCallExtendTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:            TransactionType(p.Type),
			ExtensionNumber: p.ExtensionNumber,
		}
	case "MARGIN_CALL_EXIT":
		return &MarginCallExitTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type: TransactionType(p.Type),
		}
	case "DELAYED_TRADE_CLOSURE":
		return &DelayedTradeClosureTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:     TransactionType(p.Type),
			Reason:   MarketOrderReason(p.Reason),
			TradeIDs: TradeID(p.TradeIDs),
		}
	case "DAILY_FINANCING":
		return &DailyFinancingTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                 TransactionType(p.Type),
			Financing:            p.Financing,
			AccountBalance:       p.AccountBalance,
			AccountFinancingMode: p.AccountFinancingMode,
			PositionFinancings:   p.PositionFinancings,
		}
	case "DIVIDEND_ADJUSTMENT":
		return &DividendAdjustmentTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type:                         TransactionType(p.Type),
			Instrument:                   p.Instrument,
			DividendAdjustment:           p.DividendAdjustment,
			QuoteDividendAdjustment:      p.QuoteDividendAdjustment,
			HomeConversionFactors:        p.HomeConversionFactors,
			AccountBalance:               p.AccountBalance,
			OpenTradeDividendAdjustments: p.OpenTradeDividendAdjustments,
		}
	case "RESET_RESETTABLE_PL":
		return &ResetResettablePLTransaction{
			Transaction: Transaction{
				Id:        TransactionID(p.Id),
				Time:      p.Time,
				UserID:    p.UserID,
				AccountID: AccountID(p.AccountID),
				BatchID:   TransactionID(p.BatchID),
				RequestID: RequestID(p.RequestID),
			},
			Type: TransactionType(p.Type),
		}
	}
	return p
}
