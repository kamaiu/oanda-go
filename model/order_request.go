//go:generate easyjson -all $GOFILE
package model

type OrderRequest interface {
}

// The base Order specification used when requesting that an Order be created.
// Each specific Order-type extends this definition.
// 		MarketOrderRequest
//		LimitOrderRequest
//		StopOrderRequest
//		MarketIfTouchedOrderRequest
//		TakeProfitOrderRequest
//		StopLossOrderRequest
//		GuaranteedStopLossOrderRequest
//		TrailingStopLossOrderRequest
// The base Order specification used when requesting that an Order be created. Each
// specific Order-type extends this definition.
type BaseOrderRequest struct {
	OrderRequest
}

// A MarketOrderRequest specifies the parameters that may be set when creating a Market
// Order.
type MarketOrderRequest struct {
	BaseOrderRequest
	// The type of the Order to Create. Must be set to “MARKET” when creating a Market
	// Order.
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
	// The client extensions to add to the Order. Do not set, modify, or delete clientExtensions
	// if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
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
}

// A LimitOrderRequest specifies the parameters that may be set when creating a Limit
// Order.
type LimitOrderRequest struct {
	BaseOrderRequest
	// The type of the Order to Create. Must be set to “LIMIT” when creating a Market
	// Order.
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
	// The client extensions to add to the Order. Do not set, modify, or delete clientExtensions
	// if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
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
}

// A StopOrderRequest specifies the parameters that may be set when creating a Stop
// Order.
type StopOrderRequest struct {
	BaseOrderRequest
	// The type of the Order to Create. Must be set to “STOP” when creating a Stop
	// Order.
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
	// The client extensions to add to the Order. Do not set, modify, or delete clientExtensions
	// if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
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
}

// A MarketIfTouchedOrderRequest specifies the parameters that may be set when creating
// a Market-if-Touched Order.
type MarketIfTouchedOrderRequest struct {
	BaseOrderRequest
	// The type of the Order to Create. Must be set to “MARKET_IF_TOUCHED” when creating
	// a Market If Touched Order.
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
	// The client extensions to add to the Order. Do not set, modify, or delete clientExtensions
	// if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
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
}

// A TakeProfitOrderRequest specifies the parameters that may be set when creating a
// Take Profit Order.
type TakeProfitOrderRequest struct {
	BaseOrderRequest
	// The type of the Order to Create. Must be set to “TAKE_PROFIT” when creating
	// a Take Profit Order.
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
	// The client extensions to add to the Order. Do not set, modify, or delete clientExtensions
	// if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

// A StopLossOrderRequest specifies the parameters that may be set when creating a Stop
// Loss Order. Only one of the price and distance fields may be specified.
type StopLossOrderRequest struct {
	BaseOrderRequest
	// The type of the Order to Create. Must be set to “STOP_LOSS” when creating a
	// Stop Loss Order.
	Type OrderType `json:"type"`
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
	// The client extensions to add to the Order. Do not set, modify, or delete clientExtensions
	// if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

// A GuaranteedStopLossOrderRequest specifies the parameters that may be set when creating
// a Guaranteed Stop Loss Order. Only one of the price and distance fields may be specified.
type GuaranteedStopLossOrderRequest struct {
	BaseOrderRequest
	// The type of the Order to Create. Must be set to “GUARANTEED_STOP_LOSS” when
	// creating a Guaranteed Stop Loss Order.
	Type OrderType `json:"type"`
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
	// The client extensions to add to the Order. Do not set, modify, or delete clientExtensions
	// if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

// A TrailingStopLossOrderRequest specifies the parameters that may be set when creating
// a Trailing Stop Loss Order.
type TrailingStopLossOrderRequest struct {
	BaseOrderRequest
	// The type of the Order to Create. Must be set to “TRAILING_STOP_LOSS” when creating
	// a Trailing Stop Loss Order.
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
	// The client extensions to add to the Order. Do not set, modify, or delete clientExtensions
	// if your account is associated with MT4.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

type OrderRequestParser struct {
	ClientExtensions         *ClientExtensions          `json:"clientExtensions"`
	ClientTradeID            string                     `json:"clientTradeID"`
	Distance                 DecimalNumber              `json:"distance"`
	GtdTime                  DateTime                   `json:"gtdTime"`
	Guaranteed               bool                       `json:"guaranteed"`
	GuaranteedStopLossOnFill *GuaranteedStopLossDetails `json:"guaranteedStopLossOnFill"`
	Instrument               InstrumentName             `json:"instrument"`
	PositionFill             OrderPositionFill          `json:"positionFill"`
	Price                    PriceValue                 `json:"price"`
	PriceBound               PriceValue                 `json:"priceBound"`
	StopLossOnFill           *StopLossDetails           `json:"stopLossOnFill"`
	TakeProfitOnFill         *TakeProfitDetails         `json:"takeProfitOnFill"`
	TimeInForce              TimeInForce                `json:"timeInForce"`
	TradeClientExtensions    *ClientExtensions          `json:"tradeClientExtensions"`
	TradeID                  string                     `json:"tradeID"`
	TrailingStopLossOnFill   *TrailingStopLossDetails   `json:"trailingStopLossOnFill"`
	TriggerCondition         string                     `json:"triggerCondition"`
	Type                     string                     `json:"type"`
	Units                    DecimalNumber              `json:"units"`
}

// Example
/*
r := parser.Parse()
switch v := r.(type) {
case *MarketOrderRequest:
case *LimitOrderRequest:
case *StopOrderRequest:
case *MarketIfTouchedOrderRequest:
case *TakeProfitOrderRequest:
case *StopLossOrderRequest:
case *GuaranteedStopLossOrderRequest:
case *TrailingStopLossOrderRequest:
}
*/
func (p *OrderRequestParser) Parse() interface{} {
	switch p.Type {
	case "MARKET":
		return &MarketOrderRequest{
			BaseOrderRequest:         BaseOrderRequest{},
			Type:                     OrderType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			TimeInForce:              p.TimeInForce,
			PriceBound:               p.PriceBound,
			PositionFill:             p.PositionFill,
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
		}
	case "LIMIT":
		return &LimitOrderRequest{
			BaseOrderRequest:         BaseOrderRequest{},
			Type:                     OrderType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
		}
	case "STOP":
		return &StopOrderRequest{
			BaseOrderRequest:         BaseOrderRequest{},
			Type:                     OrderType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			PriceBound:               p.PriceBound,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
		}
	case "MARKET_IF_TOUCHED":
		return &MarketIfTouchedOrderRequest{
			BaseOrderRequest:         BaseOrderRequest{},
			Type:                     OrderType(p.Type),
			Instrument:               p.Instrument,
			Units:                    p.Units,
			Price:                    p.Price,
			PriceBound:               p.PriceBound,
			TimeInForce:              p.TimeInForce,
			GtdTime:                  p.GtdTime,
			PositionFill:             p.PositionFill,
			TriggerCondition:         OrderTriggerCondition(p.TriggerCondition),
			ClientExtensions:         p.ClientExtensions,
			TakeProfitOnFill:         p.TakeProfitOnFill,
			StopLossOnFill:           p.StopLossOnFill,
			GuaranteedStopLossOnFill: p.GuaranteedStopLossOnFill,
			TrailingStopLossOnFill:   p.TrailingStopLossOnFill,
			TradeClientExtensions:    p.TradeClientExtensions,
		}
	case "TAKE_PROFIT":
		return &TakeProfitOrderRequest{
			BaseOrderRequest: BaseOrderRequest{},
			Type:             OrderType(p.Type),
			TradeID:          TradeID(p.TradeID),
			ClientTradeID:    ClientID(p.ClientTradeID),
			Price:            p.Price,
			TimeInForce:      p.TimeInForce,
			GtdTime:          p.GtdTime,
			TriggerCondition: OrderTriggerCondition(p.TriggerCondition),
			ClientExtensions: p.ClientExtensions,
		}
	case "STOP_LOSS":
		return &StopLossOrderRequest{
			BaseOrderRequest: BaseOrderRequest{},
			Type:             OrderType(p.Type),
			TradeID:          TradeID(p.TradeID),
			ClientTradeID:    ClientID(p.ClientTradeID),
			Price:            p.Price,
			Distance:         p.Distance,
			TimeInForce:      p.TimeInForce,
			GtdTime:          p.GtdTime,
			TriggerCondition: OrderTriggerCondition(p.TriggerCondition),
			Guaranteed:       p.Guaranteed,
			ClientExtensions: p.ClientExtensions,
		}
	case "GUARANTEED_STOP_LOSS":
		return &GuaranteedStopLossOrderRequest{
			BaseOrderRequest: BaseOrderRequest{},
			Type:             OrderType(p.Type),
			TradeID:          TradeID(p.TradeID),
			ClientTradeID:    ClientID(p.ClientTradeID),
			Price:            p.Price,
			Distance:         p.Distance,
			TimeInForce:      p.TimeInForce,
			GtdTime:          p.GtdTime,
			TriggerCondition: OrderTriggerCondition(p.TriggerCondition),
			ClientExtensions: p.ClientExtensions,
		}
	case "TRAILING_STOP_LOSS":
		return &TrailingStopLossOrderRequest{
			BaseOrderRequest: BaseOrderRequest{},
			Type:             OrderType(p.Type),
			TradeID:          TradeID(p.TradeID),
			ClientTradeID:    ClientID(p.ClientTradeID),
			Distance:         p.Distance,
			TimeInForce:      p.TimeInForce,
			GtdTime:          p.GtdTime,
			TriggerCondition: OrderTriggerCondition(p.TriggerCondition),
			ClientExtensions: p.ClientExtensions,
		}
	}
	return p
}
