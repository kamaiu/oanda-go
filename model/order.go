//go:generate easyjson -all $GOFILE
package model

// The Order’s identifier, unique within the Order’s Account.
// The string representation of the OANDA-assigned OrderID. OANDA-assigned OrderIDs are positive integers,
// and are derived from the TransactionID of the Transaction that created the Order.
type OrderID string

// The base Order definition specifies the properties that are common to all Orders.
type Order struct {
	// The Order’s identifier, unique within the Order’s Account.
	Id OrderID `json:"id"`
	// The time when the Order was created.
	CreateTime DateTime `json:"createTime"`
	// The current state of the Order.
	State OrderState `json:"state"`
	// The client extensions of the Order. Do not set, modify, or delete clientExtensions
	// if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

// A MarketOrder is an order that is filled immediately upon creation using the current
// market price.
type MarketOrder struct {
	Order
	// The type of the Order. Always set to “MARKET” for Market Orders.
	Type OrderType `json:"type"`
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
	// TakeProfitDetails specifies the details of a Take Profit Order to be created on
	// behalf of a client. This may happen when an Order is filled that opens a Trade requiring
	// a Take Profit, or when a Trade’s dependent Take Profit Order is modified directly
	// through the Trade.
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// StopLossDetails specifies the details of a Stop Loss Order to be created on behalf
	// of a client. This may happen when an Order is filled that opens a Trade requiring
	// a Stop Loss, or when a Trade’s dependent Stop Loss Order is modified directly through
	// the Trade.
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss Order
	// to be created on behalf of a client. This may happen when an Order is filled that
	// opens a Trade requiring a Guaranteed Stop Loss, or when a Trade’s dependent Guaranteed
	// Stop Loss Order is modified directly through the Trade.
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss Order to be
	// created on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Trailing Stop Loss, or when a Trade’s dependent Trailing Stop
	// Loss Order is modified directly through the Trade.
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created). Do not set, modify, or delete tradeClientExtensions if your
	// account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// ID of the Transaction that filled this Order (only provided when the Order’s state
	// is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`
	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime DateTime `json:"filledTime"`
	// Trade ID of Trade opened when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`
	// Trade ID of Trade reduced when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID TradeID `json:"tradeReducedID"`
	// Trade IDs of Trades closed when the Order was filled (only provided when the Order’s
	// state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`
	// ID of the Transaction that cancelled the Order (only provided when the Order’s
	// state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
	// Date/time when the Order was cancelled (only provided when the state of the Order
	// is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
}

// A FixedPriceOrder is an order that is filled immediately upon creation using a fixed
// price.
type FixedPriceOrder struct {
	Order
	// The type of the Order. Always set to “FIXED_PRICE” for Fixed Price Orders.
	Type OrderType `json:"type"`
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
	// TakeProfitDetails specifies the details of a Take Profit Order to be created on
	// behalf of a client. This may happen when an Order is filled that opens a Trade requiring
	// a Take Profit, or when a Trade’s dependent Take Profit Order is modified directly
	// through the Trade.
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// StopLossDetails specifies the details of a Stop Loss Order to be created on behalf
	// of a client. This may happen when an Order is filled that opens a Trade requiring
	// a Stop Loss, or when a Trade’s dependent Stop Loss Order is modified directly through
	// the Trade.
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss Order
	// to be created on behalf of a client. This may happen when an Order is filled that
	// opens a Trade requiring a Guaranteed Stop Loss, or when a Trade’s dependent Guaranteed
	// Stop Loss Order is modified directly through the Trade.
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss Order to be
	// created on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Trailing Stop Loss, or when a Trade’s dependent Trailing Stop
	// Loss Order is modified directly through the Trade.
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created). Do not set, modify, or delete tradeClientExtensions if your
	// account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// ID of the Transaction that filled this Order (only provided when the Order’s state
	// is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`
	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime DateTime `json:"filledTime"`
	// Trade ID of Trade opened when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`
	// Trade ID of Trade reduced when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID TradeID `json:"tradeReducedID"`
	// Trade IDs of Trades closed when the Order was filled (only provided when the Order’s
	// state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`
	// ID of the Transaction that cancelled the Order (only provided when the Order’s
	// state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
	// Date/time when the Order was cancelled (only provided when the state of the Order
	// is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
}

// A LimitOrder is an order that is created with a price threshold, and will only be
// filled by a price that is equal to or better than the threshold.
type LimitOrder struct {
	Order
	// The type of the Order. Always set to “LIMIT” for Limit Orders.
	Type OrderType `json:"type"`
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
	// TakeProfitDetails specifies the details of a Take Profit Order to be created on
	// behalf of a client. This may happen when an Order is filled that opens a Trade requiring
	// a Take Profit, or when a Trade’s dependent Take Profit Order is modified directly
	// through the Trade.
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// StopLossDetails specifies the details of a Stop Loss Order to be created on behalf
	// of a client. This may happen when an Order is filled that opens a Trade requiring
	// a Stop Loss, or when a Trade’s dependent Stop Loss Order is modified directly through
	// the Trade.
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss Order
	// to be created on behalf of a client. This may happen when an Order is filled that
	// opens a Trade requiring a Guaranteed Stop Loss, or when a Trade’s dependent Guaranteed
	// Stop Loss Order is modified directly through the Trade.
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss Order to be
	// created on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Trailing Stop Loss, or when a Trade’s dependent Trailing Stop
	// Loss Order is modified directly through the Trade.
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created). Do not set, modify, or delete tradeClientExtensions if your
	// account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// ID of the Transaction that filled this Order (only provided when the Order’s state
	// is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`
	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime DateTime `json:"filledTime"`
	// Trade ID of Trade opened when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`
	// Trade ID of Trade reduced when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID TradeID `json:"tradeReducedID"`
	// Trade IDs of Trades closed when the Order was filled (only provided when the Order’s
	// state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`
	// ID of the Transaction that cancelled the Order (only provided when the Order’s
	// state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
	// Date/time when the Order was cancelled (only provided when the state of the Order
	// is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
	// The ID of the Order that was replaced by this Order (only provided if this Order
	// was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Order that replaced this Order (only provided if this Order was cancelled
	// as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

// A StopOrder is an order that is created with a price threshold, and will only be
// filled by a price that is equal to or worse than the threshold.
type StopOrder struct {
	Order
	// The type of the Order. Always set to “STOP” for Stop Orders.
	Type OrderType `json:"type"`
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
	// TakeProfitDetails specifies the details of a Take Profit Order to be created on
	// behalf of a client. This may happen when an Order is filled that opens a Trade requiring
	// a Take Profit, or when a Trade’s dependent Take Profit Order is modified directly
	// through the Trade.
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// StopLossDetails specifies the details of a Stop Loss Order to be created on behalf
	// of a client. This may happen when an Order is filled that opens a Trade requiring
	// a Stop Loss, or when a Trade’s dependent Stop Loss Order is modified directly through
	// the Trade.
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss Order
	// to be created on behalf of a client. This may happen when an Order is filled that
	// opens a Trade requiring a Guaranteed Stop Loss, or when a Trade’s dependent Guaranteed
	// Stop Loss Order is modified directly through the Trade.
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss Order to be
	// created on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Trailing Stop Loss, or when a Trade’s dependent Trailing Stop
	// Loss Order is modified directly through the Trade.
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created). Do not set, modify, or delete tradeClientExtensions if your
	// account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// ID of the Transaction that filled this Order (only provided when the Order’s state
	// is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`
	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime DateTime `json:"filledTime"`
	// Trade ID of Trade opened when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`
	// Trade ID of Trade reduced when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID TradeID `json:"tradeReducedID"`
	// Trade IDs of Trades closed when the Order was filled (only provided when the Order’s
	// state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`
	// ID of the Transaction that cancelled the Order (only provided when the Order’s
	// state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
	// Date/time when the Order was cancelled (only provided when the state of the Order
	// is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
	// The ID of the Order that was replaced by this Order (only provided if this Order
	// was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Order that replaced this Order (only provided if this Order was cancelled
	// as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

// A MarketIfTouchedOrder is an order that is created with a price threshold, and will
// only be filled by a market price that is touches or crosses the threshold.
type MarketIfTouchedOrder struct {
	Order
	// The type of the Order. Always set to “MARKET_IF_TOUCHED” for Market If Touched
	// Orders.
	Type OrderType `json:"type"`
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
	// The Market price at the time when the MarketIfTouched Order was created.
	InitialMarketPrice PriceValue `json:"initialMarketPrice"`
	// TakeProfitDetails specifies the details of a Take Profit Order to be created on
	// behalf of a client. This may happen when an Order is filled that opens a Trade requiring
	// a Take Profit, or when a Trade’s dependent Take Profit Order is modified directly
	// through the Trade.
	TakeProfitOnFill *TakeProfitDetails `json:"takeProfitOnFill"`
	// StopLossDetails specifies the details of a Stop Loss Order to be created on behalf
	// of a client. This may happen when an Order is filled that opens a Trade requiring
	// a Stop Loss, or when a Trade’s dependent Stop Loss Order is modified directly through
	// the Trade.
	StopLossOnFill *StopLossDetails `json:"stopLossOnFill"`
	// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss Order
	// to be created on behalf of a client. This may happen when an Order is filled that
	// opens a Trade requiring a Guaranteed Stop Loss, or when a Trade’s dependent Guaranteed
	// Stop Loss Order is modified directly through the Trade.
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	// TrailingStopLossDetails specifies the details of a Trailing Stop Loss Order to be
	// created on behalf of a client. This may happen when an Order is filled that opens
	// a Trade requiring a Trailing Stop Loss, or when a Trade’s dependent Trailing Stop
	// Loss Order is modified directly through the Trade.
	TrailingStopLossOnFill *TrailingStopLossDetails `json:"trailingStopLossOnFill"`
	// Client Extensions to add to the Trade created when the Order is filled (if such
	// a Trade is created). Do not set, modify, or delete tradeClientExtensions if your
	// account is associated with MT4.
	TradeClientExtensions *ClientExtensions `json:"tradeClientExtensions"`
	// ID of the Transaction that filled this Order (only provided when the Order’s state
	// is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`
	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime DateTime `json:"filledTime"`
	// Trade ID of Trade opened when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`
	// Trade ID of Trade reduced when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID TradeID `json:"tradeReducedID"`
	// Trade IDs of Trades closed when the Order was filled (only provided when the Order’s
	// state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`
	// ID of the Transaction that cancelled the Order (only provided when the Order’s
	// state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
	// Date/time when the Order was cancelled (only provided when the state of the Order
	// is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
	// The ID of the Order that was replaced by this Order (only provided if this Order
	// was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Order that replaced this Order (only provided if this Order was cancelled
	// as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

// A TakeProfitOrder is an order that is linked to an open Trade and created with a
// price threshold. The Order will be filled (closing the Trade) by the first price
// that is equal to or better than the threshold. A TakeProfitOrder cannot be used to
// open a new Position.
type TakeProfitOrder struct {
	Order
	// The type of the Order. Always set to “TAKE_PROFIT” for Take Profit Orders.
	Type OrderType `json:"type"`
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
	// ID of the Transaction that filled this Order (only provided when the Order’s state
	// is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`
	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime DateTime `json:"filledTime"`
	// Trade ID of Trade opened when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`
	// Trade ID of Trade reduced when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID TradeID `json:"tradeReducedID"`
	// Trade IDs of Trades closed when the Order was filled (only provided when the Order’s
	// state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`
	// ID of the Transaction that cancelled the Order (only provided when the Order’s
	// state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
	// Date/time when the Order was cancelled (only provided when the state of the Order
	// is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
	// The ID of the Order that was replaced by this Order (only provided if this Order
	// was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Order that replaced this Order (only provided if this Order was cancelled
	// as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

// A StopLossOrder is an order that is linked to an open Trade and created with a price
// threshold. The Order will be filled (closing the Trade) by the first price that is
// equal to or worse than the threshold. A StopLossOrder cannot be used to open a new
// Position.
type StopLossOrder struct {
	Order
	// The type of the Order. Always set to “STOP_LOSS” for Stop Loss Orders.
	Type OrderType `json:"type"`
	// The premium that will be charged if the Stop Loss Order is guaranteed and the Order
	// is filled at the guaranteed price. It is in price units and is charged for each unit
	// of the Trade.   Deprecated: Will be removed in a future API update.
	GuaranteedExecutionPremium DecimalNumber `json:"guaranteedExecutionPremium"`
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
	// ID of the Transaction that filled this Order (only provided when the Order’s state
	// is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`
	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime DateTime `json:"filledTime"`
	// Trade ID of Trade opened when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`
	// Trade ID of Trade reduced when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID TradeID `json:"tradeReducedID"`
	// Trade IDs of Trades closed when the Order was filled (only provided when the Order’s
	// state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`
	// ID of the Transaction that cancelled the Order (only provided when the Order’s
	// state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
	// Date/time when the Order was cancelled (only provided when the state of the Order
	// is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
	// The ID of the Order that was replaced by this Order (only provided if this Order
	// was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Order that replaced this Order (only provided if this Order was cancelled
	// as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

// A GuaranteedStopLossOrder is an order that is linked to an open Trade and created
// with a price threshold which is guaranteed against slippage that may occur as the
// market crosses the price set for that order. The Order will be filled (closing the
// Trade) by the first price that is equal to or worse than the threshold. The price
// level specified for the GuaranteedStopLossOrder must be at least the configured minimum
// distance (in price units) away from the entry price for the traded instrument. A
// GuaranteedStopLossOrder cannot be used to open a new Position.
type GuaranteedStopLossOrder struct {
	Order
	// The type of the Order. Always set to “GUARANTEED_STOP_LOSS” for Guaranteed Stop
	// Loss Orders.
	Type OrderType `json:"type"`
	// The premium that will be charged if the Guaranteed Stop Loss Order is filled at
	// the guaranteed price. It is in price units and is charged for each unit of the Trade.
	GuaranteedExecutionPremium DecimalNumber `json:"guaranteedExecutionPremium"`
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
	// ID of the Transaction that filled this Order (only provided when the Order’s state
	// is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`
	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime DateTime `json:"filledTime"`
	// Trade ID of Trade opened when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`
	// Trade ID of Trade reduced when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID TradeID `json:"tradeReducedID"`
	// Trade IDs of Trades closed when the Order was filled (only provided when the Order’s
	// state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`
	// ID of the Transaction that cancelled the Order (only provided when the Order’s
	// state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
	// Date/time when the Order was cancelled (only provided when the state of the Order
	// is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
	// The ID of the Order that was replaced by this Order (only provided if this Order
	// was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Order that replaced this Order (only provided if this Order was cancelled
	// as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

// A TrailingStopLossOrder is an order that is linked to an open Trade and created with
// a price distance. The price distance is used to calculate a trailing stop value for
// the order that is in the losing direction from the market price at the time of the
// order’s creation. The trailing stop value will follow the market price as it moves
// in the winning direction, and the order will filled (closing the Trade) by the first
// price that is equal to or worse than the trailing stop value. A TrailingStopLossOrder
// cannot be used to open a new Position.
type TrailingStopLossOrder struct {
	Order
	// The type of the Order. Always set to “TRAILING_STOP_LOSS” for Trailing Stop
	// Loss Orders.
	Type OrderType `json:"type"`
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
	// The trigger price for the Trailing Stop Loss Order. The trailing stop value will
	// trail (follow) the market price by the TSL order’s configured “distance” as
	// the market price moves in the winning direction. If the market price moves to a level
	// that is equal to or worse than the trailing stop value, the order will be filled
	// and the Trade will be closed.
	TrailingStopValue PriceValue `json:"trailingStopValue"`
	// ID of the Transaction that filled this Order (only provided when the Order’s state
	// is FILLED)
	FillingTransactionID TransactionID `json:"fillingTransactionID"`
	// Date/time when the Order was filled (only provided when the Order’s state is FILLED)
	FilledTime DateTime `json:"filledTime"`
	// Trade ID of Trade opened when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was opened as a result of the fill)
	TradeOpenedID TradeID `json:"tradeOpenedID"`
	// Trade ID of Trade reduced when the Order was filled (only provided when the Order’s
	// state is FILLED and a Trade was reduced as a result of the fill)
	TradeReducedID TradeID `json:"tradeReducedID"`
	// Trade IDs of Trades closed when the Order was filled (only provided when the Order’s
	// state is FILLED and one or more Trades were closed as a result of the fill)
	TradeClosedIDs []TradeID `json:"tradeClosedIDs"`
	// ID of the Transaction that cancelled the Order (only provided when the Order’s
	// state is CANCELLED)
	CancellingTransactionID TransactionID `json:"cancellingTransactionID"`
	// Date/time when the Order was cancelled (only provided when the state of the Order
	// is CANCELLED)
	CancelledTime DateTime `json:"cancelledTime"`
	// The ID of the Order that was replaced by this Order (only provided if this Order
	// was created as part of a cancel/replace).
	ReplacesOrderID OrderID `json:"replacesOrderID"`
	// The ID of the Order that replaced this Order (only provided if this Order was cancelled
	// as part of a cancel/replace).
	ReplacedByOrderID OrderID `json:"replacedByOrderID"`
}

type OrderParser struct {
	CancelledTime              DateTime                      `json:"cancelledTime"`
	CancellingTransactionID    string                        `json:"cancellingTransactionID"`
	ClientExtensions           *ClientExtensions             `json:"clientExtensions"`
	ClientTradeID              string                        `json:"clientTradeID"`
	CreateTime                 DateTime                      `json:"createTime"`
	DelayedTradeClose          *MarketOrderDelayedTradeClose `json:"delayedTradeClose"`
	Distance                   DecimalNumber                 `json:"distance"`
	FilledTime                 DateTime                      `json:"filledTime"`
	FillingTransactionID       string                        `json:"fillingTransactionID"`
	GtdTime                    DateTime                      `json:"gtdTime"`
	Guaranteed                 bool                          `json:"guaranteed"`
	GuaranteedExecutionPremium DecimalNumber                 `json:"guaranteedExecutionPremium"`
	GuaranteedStopLossOnFill   *GuaranteedStopLossDetails    `json:"guaranteedStopLossOnFill"`
	Id                         string                        `json:"id"`
	InitialMarketPrice         PriceValue                    `json:"initialMarketPrice"`
	Instrument                 InstrumentName                `json:"instrument"`
	LongPositionCloseout       *MarketOrderPositionCloseout  `json:"longPositionCloseout"`
	MarginCloseout             *MarketOrderMarginCloseout    `json:"marginCloseout"`
	PositionFill               OrderPositionFill             `json:"positionFill"`
	Price                      PriceValue                    `json:"price"`
	PriceBound                 PriceValue                    `json:"priceBound"`
	ReplacedByOrderID          string                        `json:"replacedByOrderID"`
	ReplacesOrderID            string                        `json:"replacesOrderID"`
	ShortPositionCloseout      *MarketOrderPositionCloseout  `json:"shortPositionCloseout"`
	State                      OrderState                    `json:"state"`
	StopLossOnFill             *StopLossDetails              `json:"stopLossOnFill"`
	TakeProfitOnFill           *TakeProfitDetails            `json:"takeProfitOnFill"`
	TimeInForce                TimeInForce                   `json:"timeInForce"`
	TradeClientExtensions      *ClientExtensions             `json:"tradeClientExtensions"`
	TradeClose                 *MarketOrderTradeClose        `json:"tradeClose"`
	TradeClosedIDs             []TradeID                     `json:"tradeClosedIDs"`
	TradeID                    string                        `json:"tradeID"`
	TradeOpenedID              string                        `json:"tradeOpenedID"`
	TradeReducedID             string                        `json:"tradeReducedID"`
	TradeState                 string                        `json:"tradeState"`
	TrailingStopLossOnFill     *TrailingStopLossDetails      `json:"trailingStopLossOnFill"`
	TrailingStopValue          PriceValue                    `json:"trailingStopValue"`
	TriggerCondition           string                        `json:"triggerCondition"`
	Type                       string                        `json:"type"`
	Units                      DecimalNumber                 `json:"units"`
}

// Example
/*
r := parser.Parse()
switch v := r.(type) {
case *MarketOrder:
case *FixedPriceOrder:
case *LimitOrder:
case *StopOrder:
case *MarketIfTouchedOrder:
case *TakeProfitOrder:
case *StopLossOrder:
case *GuaranteedStopLossOrder:
case *TrailingStopLossOrder:
}
*/
func (p *OrderParser) Parse() interface{} {
	switch p.Type {
	case "MARKET":
		return &MarketOrder{
			Order: Order{
				Id:               OrderID(p.Id),
				CreateTime:       p.CreateTime,
				State:            p.State,
				ClientExtensions: p.ClientExtensions,
			},
			Type:                     OrderType(p.Type),
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
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			FillingTransactionID:     TransactionID(p.FillingTransactionID),
			FilledTime:               p.FilledTime,
			TradeOpenedID:            TradeID(p.TradeOpenedID),
			TradeReducedID:           TradeID(p.TradeReducedID),
			TradeClosedIDs:           p.TradeClosedIDs,
			CancellingTransactionID:  TransactionID(p.CancellingTransactionID),
			CancelledTime:            p.CancelledTime,
		}
	case "FIXED_PRICE":
		return &FixedPriceOrder{
			Order: Order{
				Id:               OrderID(p.Id),
				CreateTime:       p.CreateTime,
				State:            p.State,
				ClientExtensions: p.ClientExtensions,
			},
			Type:                     OrderType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			PositionFill:             p.PositionFill,
			TradeState:               p.TradeState,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			FillingTransactionID:     TransactionID(p.FillingTransactionID),
			FilledTime:               p.FilledTime,
			TradeOpenedID:            TradeID(p.TradeOpenedID),
			TradeReducedID:           TradeID(p.TradeReducedID),
			TradeClosedIDs:           p.TradeClosedIDs,
			CancellingTransactionID:  TransactionID(p.CancellingTransactionID),
			CancelledTime:            p.CancelledTime,
		}
	case "LIMIT":
		return &LimitOrder{
			Order: Order{
				Id:               OrderID(p.Id),
				CreateTime:       p.CreateTime,
				State:            p.State,
				ClientExtensions: p.ClientExtensions,
			},
			Type:                     OrderType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			FillingTransactionID:     TransactionID(p.FillingTransactionID),
			FilledTime:               p.FilledTime,
			TradeOpenedID:            TradeID(p.TradeOpenedID),
			TradeReducedID:           TradeID(p.TradeReducedID),
			TradeClosedIDs:           p.TradeClosedIDs,
			CancellingTransactionID:  TransactionID(p.CancellingTransactionID),
			CancelledTime:            p.CancelledTime,
			ReplacesOrderID:          OrderID(p.ReplacesOrderID),
			ReplacedByOrderID:        OrderID(p.ReplacedByOrderID),
		}
	case "STOP":
		return &StopOrder{
			Order: Order{
				Id:               OrderID(p.Id),
				CreateTime:       p.CreateTime,
				State:            p.State,
				ClientExtensions: p.ClientExtensions,
			},
			Type:                     OrderType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			PriceBound:               p.PriceBound,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			FillingTransactionID:     TransactionID(p.FillingTransactionID),
			FilledTime:               p.FilledTime,
			TradeOpenedID:            TradeID(p.TradeOpenedID),
			TradeReducedID:           TradeID(p.TradeReducedID),
			TradeClosedIDs:           p.TradeClosedIDs,
			CancellingTransactionID:  TransactionID(p.CancellingTransactionID),
			CancelledTime:            p.CancelledTime,
			ReplacesOrderID:          OrderID(p.ReplacesOrderID),
			ReplacedByOrderID:        OrderID(p.ReplacedByOrderID),
		}
	case "MARKET_IF_TOUCHED":
		return &MarketIfTouchedOrder{
			Order: Order{
				Id:               OrderID(p.Id),
				CreateTime:       p.CreateTime,
				State:            p.State,
				ClientExtensions: p.ClientExtensions,
			},
			Type:                     OrderType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			PriceBound:               p.PriceBound,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			InitialMarketPrice:       p.InitialMarketPrice,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
			FillingTransactionID:     TransactionID(p.FillingTransactionID),
			FilledTime:               p.FilledTime,
			TradeOpenedID:            TradeID(p.TradeOpenedID),
			TradeReducedID:           TradeID(p.TradeReducedID),
			TradeClosedIDs:           p.TradeClosedIDs,
			CancellingTransactionID:  TransactionID(p.CancellingTransactionID),
			CancelledTime:            p.CancelledTime,
			ReplacesOrderID:          OrderID(p.ReplacesOrderID),
			ReplacedByOrderID:        OrderID(p.ReplacedByOrderID),
		}
	case "TAKE_PROFIT":
		return &TakeProfitOrder{
			Order: Order{
				Id:               OrderID(p.Id),
				CreateTime:       p.CreateTime,
				State:            p.State,
				ClientExtensions: p.ClientExtensions,
			},
			Type:                    OrderType(p.Type),
			TradeID:                 TradeID(p.TradeID),
			ClientTradeID:           ClientID(p.ClientTradeID),
			Price:                   p.Price,
			TimeInForce:             p.TimeInForce,
			GtdTime:                 p.GtdTime,
			TriggerCondition:        OrderTriggerCondition(p.TriggerCondition),
			FillingTransactionID:    TransactionID(p.FillingTransactionID),
			FilledTime:              p.FilledTime,
			TradeOpenedID:           TradeID(p.TradeOpenedID),
			TradeReducedID:          TradeID(p.TradeReducedID),
			TradeClosedIDs:          p.TradeClosedIDs,
			CancellingTransactionID: TransactionID(p.CancellingTransactionID),
			CancelledTime:           p.CancelledTime,
			ReplacesOrderID:         OrderID(p.ReplacesOrderID),
			ReplacedByOrderID:       OrderID(p.ReplacedByOrderID),
		}
	case "STOP_LOSS":
		return &StopLossOrder{
			Order: Order{
				Id:               OrderID(p.Id),
				CreateTime:       p.CreateTime,
				State:            p.State,
				ClientExtensions: p.ClientExtensions,
			},
			Type:                       OrderType(p.Type),
			GuaranteedExecutionPremium: p.GuaranteedExecutionPremium,
			TradeID:                    TradeID(p.TradeID),
			ClientTradeID:              ClientID(p.ClientTradeID),
			Price:                      p.Price,
			Distance:                   p.Distance,
			TimeInForce:                p.TimeInForce,
			GtdTime:                    p.GtdTime,
			TriggerCondition:           OrderTriggerCondition(p.TriggerCondition),
			Guaranteed:                 p.Guaranteed,
			FillingTransactionID:       TransactionID(p.FillingTransactionID),
			FilledTime:                 p.FilledTime,
			TradeOpenedID:              TradeID(p.TradeOpenedID),
			TradeReducedID:             TradeID(p.TradeReducedID),
			TradeClosedIDs:             p.TradeClosedIDs,
			CancellingTransactionID:    TransactionID(p.CancellingTransactionID),
			CancelledTime:              p.CancelledTime,
			ReplacesOrderID:            OrderID(p.ReplacesOrderID),
			ReplacedByOrderID:          OrderID(p.ReplacedByOrderID),
		}
	case "GUARANTEED_STOP_LOSS":
		return &GuaranteedStopLossOrder{
			Order: Order{
				Id:               OrderID(p.Id),
				CreateTime:       p.CreateTime,
				State:            p.State,
				ClientExtensions: p.ClientExtensions,
			},
			Type:                       OrderType(p.Type),
			GuaranteedExecutionPremium: p.GuaranteedExecutionPremium,
			TradeID:                    TradeID(p.TradeID),
			ClientTradeID:              ClientID(p.ClientTradeID),
			Price:                      p.Price,
			Distance:                   p.Distance,
			TimeInForce:                p.TimeInForce,
			GtdTime:                    p.GtdTime,
			TriggerCondition:           OrderTriggerCondition(p.TriggerCondition),
			FillingTransactionID:       TransactionID(p.FillingTransactionID),
			FilledTime:                 p.FilledTime,
			TradeOpenedID:              TradeID(p.TradeOpenedID),
			TradeReducedID:             TradeID(p.TradeReducedID),
			TradeClosedIDs:             p.TradeClosedIDs,
			CancellingTransactionID:    TransactionID(p.CancellingTransactionID),
			CancelledTime:              p.CancelledTime,
			ReplacesOrderID:            OrderID(p.ReplacesOrderID),
			ReplacedByOrderID:          OrderID(p.ReplacedByOrderID),
		}
	case "TRAILING_STOP_LOSS":
		return &TrailingStopLossOrder{
			Order: Order{
				Id:               OrderID(p.Id),
				CreateTime:       p.CreateTime,
				State:            p.State,
				ClientExtensions: p.ClientExtensions,
			},
			Type:                    OrderType(p.Type),
			TradeID:                 TradeID(p.TradeID),
			ClientTradeID:           ClientID(p.ClientTradeID),
			Distance:                p.Distance,
			TimeInForce:             p.TimeInForce,
			GtdTime:                 p.GtdTime,
			TriggerCondition:        OrderTriggerCondition(p.TriggerCondition),
			TrailingStopValue:       p.TrailingStopValue,
			FillingTransactionID:    TransactionID(p.FillingTransactionID),
			FilledTime:              p.FilledTime,
			TradeOpenedID:           TradeID(p.TradeOpenedID),
			TradeReducedID:          TradeID(p.TradeReducedID),
			TradeClosedIDs:          p.TradeClosedIDs,
			CancellingTransactionID: TransactionID(p.CancellingTransactionID),
			CancelledTime:           p.CancelledTime,
			ReplacesOrderID:         OrderID(p.ReplacesOrderID),
			ReplacedByOrderID:       OrderID(p.ReplacedByOrderID),
		}
	}
	return p
}
