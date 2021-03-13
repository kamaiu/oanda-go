//go:generate easyjson -all $GOFILE
package model

// A client-provided identifier, used by clients to refer to their Orders or Trades
// with an identifier that they have provided.
type ClientID string

// A client-provided tag that can contain any data and may be assigned to their Orders or Trades.
// Tags are typically used to associate groups of Trades and/or Orders together.
type ClientTag string

// A client-provided comment that can contain any data and may be assigned to their Orders or Trades.
// Comments are typically used to provide extra context or meaning to an Order or Trade.
type ClientComment string

// A ClientExtensions object allows a client to attach a clientID, tag and comment to Orders
// and Trades in their Account. Do not set, modify, or delete this field if your account is
// associated with MT4.
type ClientExtensions struct {
	// The Client ID of the Order/Trade
	ID ClientID `json:"id"`
	// A tag associated with the Order/Trade
	Tag ClientTag `json:"tag"`
	// A comment associated with the Order/Trade
	Comment ClientComment `json:"comment"`
}

// A MarketOrderTradeClose specifies the extensions to a Market Order that has been
// created specifically to close a Trade.
type MarketOrderTradeClose struct {
	// The ID of the Trade requested to be closed
	TradeID TradeID `json:"tradeId"`
	// The client ID of the Trade requested to be closed
	ClientTradeID string `json:"clientTradeId"`
	// Indication of how much of the Trade to close. Either “ALL”, or a
	// DecimalNumber reflection a partial close of the Trade.
	Units string `json:"units"`
}

// Details for the Market Order extensions specific to a Market Order placed that is part of a
// Market Order Margin Closeout in a client’s account
type MarketOrderMarginCloseout struct {
	Reason MarketOrderMarginCloseoutReason `json:"reason"`
}

// The reason that the Market Order was created to perform a margin closeout
type MarketOrderMarginCloseoutReason string

const (
	// Trade closures resulted from violating OANDA’s margin policy
	MarketOrderMarginCloseoutReason_MARGIN_CHECK_VIOLATION MarketOrderMarginCloseoutReason = "MARGIN_CHECK_VIOLATION"
	// Trade closures came from a margin closeout event resulting from regulatory conditions placed on the Account’s margin call
	MarketOrderMarginCloseoutReason_REGULATORY_MARGIN_CALL_VIOLATION MarketOrderMarginCloseoutReason = "REGULATORY_MARGIN_CALL_VIOLATION"
	// Trade closures resulted from violating the margin policy imposed by regulatory requirements
	MarketOrderMarginCloseoutReason_REGULATORY_MARGIN_CHECK_VIOLATION MarketOrderMarginCloseoutReason = "REGULATORY_MARGIN_CHECK_VIOLATION"
)

// Details for the Market Order extensions specific to a Market Order placed with the intent
// of fully closing a specific open trade that should have already been closed but wasn’t due
// to halted market conditions
type MarketOrderDelayedTradeClose struct {
	// The ID of the Trade being closed
	TradeID TradeID `json:"tradeId"`
	// The client ID of the Trade being closed
	ClientTradeID TradeID `json:"clientTradeId"`
	// The Transaction ID of the DelayedTradeClosure transaction to which this
	// Delayed Trade Close belongs to
	SourceTransactionID TransactionID `json:"sourceTransactionID"`
}

// A MarketOrderPositionCloseout specifies the extensions to a Market Order when it has
// been created to closeout a specific Position.
type MarketOrderPositionCloseout struct {
	// The instrument of the Position being closed out.
	Instrument InstrumentName `json:"instrument"`
	// Indication of how much of the Position to close. Either “ALL”, or a
	// DecimalNumber reflection a partial close of the Trade. The DecimalNumber
	// must always be positive, and represent a number that doesn’t exceed the
	// absolute size of the Position.
	Units string `json:"units"`
}

// TakeProfitDetails specifies the details of a Take Profit Order to be created on behalf of a client.
// This may happen when an Order is filled that opens a Trade requiring a Take Profit,
// or when a Trade’s dependent Take Profit Order is modified directly through the Trade.
type TakeProfitDetails struct {
	// The price that the Take Profit Order will be triggered at. Only one of
	// the price and distance fields may be specified.
	Price PriceValue `json:"price"`
	// The time in force for the created Take Profit Order. This may only be GTC, GTD or GFD.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date when the Take Profit Order will be cancelled on if timeInForce is GTD.
	GtdTime DateTime `json:"gtdTime"`
	// The Client Extensions to add to the Take Profit Order when created.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

// StopLossDetails specifies the details of a Stop Loss Order to be created on behalf of a client.
// This may happen when an Order is filled that opens a Trade requiring a Stop Loss,
// or when a Trade’s dependent Stop Loss Order is modified directly through the Trade.
type StopLossDetails struct {
	// The price that the Stop Loss Order will be triggered at. Only one of the
	// price and distance fields may be specified.
	Price PriceValue `json:"price"`
	// Specifies the distance (in price units) from the Trade’s open price to
	// use as the Stop Loss Order price. Only one of the distance and price
	// fields may be specified.
	Distance DecimalNumber `json:"distance"`
	// The time in force for the created Stop Loss Order. This may only be GTC, GTD or GFD.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date when the Stop Loss Order will be cancelled on if timeInForce is GTD.
	GtdTime DateTime `json:"gtdTime"`
	// The Client Extensions to add to the Stop Loss Order when created.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

// GuaranteedStopLossDetails specifies the details of a Guaranteed Stop Loss Order to be
// created on behalf of a client. This may happen when an Order is filled that opens a
// Trade requiring a Guaranteed Stop Loss, or when a Trade’s dependent Guaranteed Stop Loss Order
// is modified directly through the Trade.
type GuaranteedStopLossDetails struct {
	// The price that the Guaranteed Stop Loss Order will be triggered at. Only
	// one of the price and distance fields may be specified.
	Price PriceValue `json:"price"`
	// Specifies the distance (in price units) from the Trade’s open price to
	// use as the Guaranteed Stop Loss Order price. Only one of the distance and
	// price fields may be specified.
	Distance DecimalNumber `json:"distance"`
	// The time in force for the created Guaranteed Stop Loss Order. This may only be GTC, GTD or GFD.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date when the Guaranteed Stop Loss Order will be cancelled on if timeInForce is GTD.
	GtdTime DateTime `json:"gtdTime"`
	// The Client Extensions to add to the Guaranteed Stop Loss Order when created.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

// TrailingStopLossDetails specifies the details of a Trailing Stop Loss Order to be created on
// behalf of a client. This may happen when an Order is filled that opens a Trade requiring a
// Trailing Stop Loss, or when a Trade’s dependent Trailing Stop Loss Order is modified directly
// through the Trade.
type TrailingStopLossDetails struct {
	// The distance (in price units) from the Trade’s fill price that the
	// Trailing Stop Loss Order will be triggered at.
	Distance DecimalNumber `json:"distance"`
	// The time in force for the created Trailing Stop Loss Order. This may only be GTC, GTD or GFD.
	TimeInForce TimeInForce `json:"timeInForce"`
	// The date when the Trailing Stop Loss Order will be cancelled on if timeInForce is GTD.
	GtdTime DateTime `json:"gtdTime"`
	// The Client Extensions to add to the Trailing Stop Loss Order when created.
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
}

// A TradeOpen object represents a Trade for an instrument that was opened in an Account.
// It is found embedded in Transactions that affect the position of an instrument in the
// Account, specifically the OrderFill Transaction.
type TradeOpen struct {
	// The ID of the Trade that was opened
	TradeID TradeID `json:"tradeID"`
	// The number of units opened by the Trade
	Units DecimalNumber `json:"units"`
	// The average price that the units were opened at.
	Price PriceValue `json:"price"`
	// This is the fee charged for opening the trade if it has a guaranteed Stop
	// Loss Order attached to it.
	GuaranteedExecutionFee AccountUnits `json:"guaranteedExecutionFee"`
	// This is the fee charged for opening the trade if it has a guaranteed Stop
	// Loss Order attached to it, expressed in the Instrument’s quote currency.
	QuoteGuaranteedExecutionFee DecimalNumber `json:"quoteGuaranteedExecutionFee"`
	// The client extensions for the newly opened Trade
	ClientExtensions *ClientExtensions `json:"clientExtensions"`
	// The half spread cost for the trade open. This can be a positive or
	// negative value and is represented in the home currency of the Account.
	HalfSpreadCost AccountUnits `json:"halfSpreadCost"`
	// The margin required at the time the Trade was created. Note, this is the
	// ‘pure’ margin required, it is not the ‘effective’ margin used that
	// factors in the trade risk if a GSLO is attached to the trade.
	InitialMarginRequired AccountUnits `json:"initialMarginRequired"`
}

// A TradeReduce object represents a Trade for an instrument that was reduced
// (either partially or fully) in an Account. It is found embedded in Transactions that affect
// the position of an instrument in the account, specifically the OrderFill Transaction.
type TradeReduce struct {
	// The ID of the Trade that was reduced or closed
	TradeID TradeID `json:"tradeID"`
	// The number of units that the Trade was reduced by
	Units DecimalNumber `json:"units"`
	// The average price that the units were closed at. This price may be
	// clamped for guaranteed Stop Loss Orders.
	Price PriceValue `json:"price"`
	// The PL realized when reducing the Trade
	RealizedPL AccountUnits `json:"realizedPL"`
	// The financing paid/collected when reducing the Trade
	Financing AccountUnits `json:"financing"`
	// The base financing paid/collected when reducing the Trade
	BaseFinancing DecimalNumber `json:"baseFinancing"`
	// The quote financing paid/collected when reducing the Trade
	QuoteFinancing DecimalNumber `json:"quoteFinancing"`
	// The financing rate in effect for the instrument used to calculate the
	// amount of financing paid/collected when reducing the Trade. This field
	// will only be set if the AccountFinancingMode at the time of the order
	// fill is SECOND_BY_SECOND_INSTRUMENT. The value is in decimal rather than
	// percentage points, e.g. 5% is represented as 0.05.
	FinancingRate DecimalNumber `json:"financingRate"`
	// This is the fee that is charged for closing the Trade if it has a
	// guaranteed Stop Loss Order attached to it.
	GuaranteedExecutionFee AccountUnits `json:"guaranteedExecutionFee"`
	// This is the fee that is charged for closing the Trade if it has a
	// guaranteed Stop Loss Order attached to it, expressed in the Instrument’s
	// quote currency.
	QuoteGuaranteedExecutionFee DecimalNumber `json:"quoteGuaranteedExecutionFee"`
	//
	HalfSpreadCost AccountUnits `json:"halfSpreadCost"`
}

// The possible types of a Transaction
type TransactionType string

const (
	// Account Create Transaction
	TransactionType_CREATE TransactionType = "CREATE"
	// Account Close Transaction
	TransactionType_CLOSE TransactionType = "CLOSE"
	// Account Reopen Transaction
	TransactionType_REOPEN TransactionType = "REOPEN"
	// Client Configuration Transaction
	TransactionType_CLIENT_CONFIGURE TransactionType = "CLIENT_CONFIGURE"
	// Client Configuration Reject Transaction
	TransactionType_CLIENT_CONFIGURE_REJECT TransactionType = "CLIENT_CONFIGURE_REJECT"
	// Transfer Funds Transaction
	TransactionType_TRANSFER_FUNDS TransactionType = "TRANSFER_FUNDS"
	// Transfer Funds Reject Transaction
	TransactionType_TRANSFER_FUNDS_REJECT TransactionType = "TRANSFER_FUNDS_REJECT"
	// Market Order Transaction
	TransactionType_MARKET_ORDER TransactionType = "MARKET_ORDER"
	// Market Order Reject Transaction
	TransactionType_MARKET_ORDER_REJECT TransactionType = "MARKET_ORDER_REJECT"
	// Fixed Price Order Transaction
	TransactionType_FIXED_PRICE_ORDER TransactionType = "FIXED_PRICE_ORDER"
	// Order Transaction
	TransactionType_LIMIT_ORDER TransactionType = "LIMIT_ORDER"
	// Limit Order Reject Transaction
	TransactionType_LIMIT_ORDER_REJECT TransactionType = "LIMIT_ORDER_REJECT"
	// Stop Order Transaction
	TransactionType_STOP_ORDER TransactionType = "STOP_ORDER"
	// Stop Order Reject Transaction
	TransactionType_STOP_ORDER_REJECT TransactionType = "STOP_ORDER_REJECT"
	// Market if Touched Order Transaction
	TransactionType_MARKET_IF_TOUCHED_ORDER TransactionType = "MARKET_IF_TOUCHED_ORDER"
	// Market if Touched Order Reject Transaction
	TransactionType_MARKET_IF_TOUCHED_ORDER_REJECT TransactionType = "MARKET_IF_TOUCHED_ORDER_REJECT"
	// Take Profit Order Transaction
	TransactionType_TAKE_PROFIT_ORDER TransactionType = "TAKE_PROFIT_ORDER"
	// Take Profit Order Reject Transaction
	TransactionType_TAKE_PROFIT_ORDER_REJECT TransactionType = "TAKE_PROFIT_ORDER_REJECT"
	// Stop Loss Order Transaction
	TransactionType_STOP_LOSS_ORDER TransactionType = "STOP_LOSS_ORDER"
	// Stop Loss Order Reject Transaction
	TransactionType_STOP_LOSS_ORDER_REJECT TransactionType = "STOP_LOSS_ORDER_REJECT"
	// Guaranteed Stop Loss Order Transaction
	TransactionType_GUARANTEED_STOP_LOSS_ORDER TransactionType = "GUARANTEED_STOP_LOSS_ORDER"
	// Guaranteed Stop Loss Order Reject Transaction
	TransactionType_GUARANTEED_STOP_LOSS_ORDER_REJECT TransactionType = "GUARANTEED_STOP_LOSS_ORDER_REJECT"
	// Trailing Stop Loss Order Transaction
	TransactionType_TRAILING_STOP_LOSS_ORDER TransactionType = "TRAILING_STOP_LOSS_ORDER"
	// Trailing Stop Loss Order Reject Transaction
	TransactionType_TRAILING_STOP_LOSS_ORDER_REJECT TransactionType = "TRAILING_STOP_LOSS_ORDER_REJECT"
	// Order Fill Transaction
	TransactionType_ORDER_FILL TransactionType = "ORDER_FILL"
	// Order Cancel Transaction
	TransactionType_ORDER_CANCEL TransactionType = "ORDER_CANCEL"
	// Order Cancel Reject Transaction
	TransactionType_ORDER_CANCEL_REJECT TransactionType = "ORDER_CANCEL_REJECT"
	// Order Client Extensions Modify Transaction
	TransactionType_ORDER_CLIENT_EXTENSIONS_MODIFY TransactionType = "ORDER_CLIENT_EXTENSIONS_MODIFY"
	// Order Client Extensions Modify Reject Transaction
	TransactionType_ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT TransactionType = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	// Trade Client Extensions Modify Transaction
	TransactionType_TRADE_CLIENT_EXTENSIONS_MODIFY TransactionType = "TRADE_CLIENT_EXTENSIONS_MODIFY"
	// Trade Client Extensions Modify Reject Transaction
	TransactionType_TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT TransactionType = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	// Margin Call Enter Transaction
	TransactionType_MARGIN_CALL_ENTER TransactionType = "MARGIN_CALL_ENTER"
	// Margin Call Extend Transaction
	TransactionType_MARGIN_CALL_EXTEND TransactionType = "MARGIN_CALL_EXTEND"
	// Margin Call Exit Transaction
	TransactionType_MARGIN_CALL_EXIT TransactionType = "MARGIN_CALL_EXIT"
	// Delayed Trade Closure Transaction
	TransactionType_DELAYED_TRADE_CLOSURE TransactionType = "DELAYED_TRADE_CLOSURE"
	// Daily Financing Transaction
	TransactionType_DAILY_FINANCING TransactionType = "DAILY_FINANCING"
	// Dividend Adjustment Transaction
	TransactionType_DIVIDEND_ADJUSTMENT TransactionType = "DIVIDEND_ADJUSTMENT"
	// Reset Resettable PL Transaction
	TransactionType_RESET_RESETTABLE_PL TransactionType = "RESET_RESETTABLE_PL"
)

// The reason that an Account is being funded.
type FundingReason string

const (
	// The client has initiated a funds transfer
	FundingReason_CLIENT_FUNDING FundingReason = "CLIENT_FUNDING"
	// Funds are being transferred between two Accounts
	FundingReason_ACCOUNT_TRANSFER FundingReason = "ACCOUNT_TRANSFER"
	// Funds are being transferred as part of a Division migration
	FundingReason_DIVISION_MIGRATION FundingReason = "DIVISION_MIGRATION"
	// Funds are being transferred as part of a Site migration
	FundingReason_SITE_MIGRATION FundingReason = "SITE_MIGRATION"
	// Funds are being transferred as part of an Account adjustment
	FundingReason_ADJUSTMENT FundingReason = "ADJUSTMENT"
)

// The reason that the Market Order was created
type MarketOrderReason string

const (
	// The Market Order was created at the request of a client
	MarketOrderReason_CLIENT_ORDER MarketOrderReason = "CLIENT_ORDER"
	//  The Market Order was created to close a Trade at the request of a client
	MarketOrderReason_TRADE_CLOSE MarketOrderReason = "TRADE_CLOSE"
	// The Market Order was created to close a Position at the request of a client
	MarketOrderReason_POSITION_CLOSEOUT MarketOrderReason = "POSITION_CLOSEOUT"
	// The Market Order was created as part of a Margin Closeout
	MarketOrderReason_MARGIN_CLOSEOUT MarketOrderReason = "MARGIN_CLOSEOUT"
	// The Market Order was created to close a trade marked for delayed closure
	MarketOrderReason_DELAYED_TRADE_CLOSE MarketOrderReason = "DELAYED_TRADE_CLOSE"
)

// The reason that the Fixed Price Order was created
type FixedPriceOrderReason string

const (
	// The Fixed Price Order was created as part of a platform account migration
	FixedPriceOrderReason_PLATFORM_ACCOUNT_MIGRATION FixedPriceOrderReason = "PLATFORM_ACCOUNT_MIGRATION"
	// The Fixed Price Order was created to close a Trade as part of division account migration
	FixedPriceOrderReason_TRADE_CLOSE_DIVISION_ACCOUNT_MIGRATION FixedPriceOrderReason = "TRADE_CLOSE_DIVISION_ACCOUNT_MIGRATION"
	// The Fixed Price Order was created to close a Trade administratively
	FixedPriceOrderReason_TRADE_CLOSE_ADMINISTRATIVE_ACTION FixedPriceOrderReason = "TRADE_CLOSE_ADMINISTRATIVE_ACTION"
)

// The reason that the Limit Order was initiated
type LimitOrderReason string

const (
	// The Limit Order was initiated at the request of a client
	LimitOrderReason_CLIENT_ORDER LimitOrderReason = "CLIENT_ORDER"
	// The Limit Order was initiated as a replacement for an existing Order
	LimitOrderReason_REPLACEMENT LimitOrderReason = "REPLACEMENT"
)

// The reason that the Stop Order was initiated
type StopOrderReason string

const (
	// The Stop Order was initiated at the request of a client
	StopOrderReason_CLIENT_ORDER LimitOrderReason = "CLIENT_ORDER"
	// The Stop Order was initiated as a replacement for an existing Order
	StopOrderReason_REPLACEMENT LimitOrderReason = "REPLACEMENT"
)

// The reason that the Market-if-touched Order was initiated
type MarketIfTouchedOrderReason string

const (
	// The Market-if-touched Order was initiated at the request of a client
	MarketIfTouchedOrderReason_CLIENT_ORDER LimitOrderReason = "CLIENT_ORDER"
	// The Market-if-touched Order was initiated as a replacement for an existing Order
	MarketIfTouchedOrderReason_REPLACEMENT LimitOrderReason = "REPLACEMENT"
)

// The reason that the Take Profit Order was initiated
type TakeProfitOrderReason string

const (
	// The Take Profit Order was initiated at the request of a client
	TakeProfitOrderReason_CLIENT_ORDER LimitOrderReason = "CLIENT_ORDER"
	// The Take Profit Order was initiated as a replacement for an existing Order
	TakeProfitOrderReason_REPLACEMENT LimitOrderReason = "REPLACEMENT"
	// The Take Profit Order was initiated automatically when an Order was filled that opened a new
	// Trade requiring a Take Profit Order.
	TakeProfitOrderReason_ON_FILL LimitOrderReason = "ON_FILL"
)

// The reason that the Stop Loss Order was initiated
type StopLossOrderReason string

const (
	// The Stop Loss Order was initiated at the request of a client
	StopLossOrderReason_CLIENT_ORDER LimitOrderReason = "CLIENT_ORDER"
	// The Stop Loss Order was initiated as a replacement for an existing Order
	StopLossOrderReason_REPLACEMENT LimitOrderReason = "REPLACEMENT"
	// The Stop Loss Order was initiated automatically when an Order was filled that opened a new
	// Trade requiring a Stop Loss Order.
	StopLossOrderReason_ON_FILL LimitOrderReason = "ON_FILL"
)

// The reason that the Guaranteed Stop Loss Order was initiated
type GuaranteedStopLossOrderReason string

const (
	// The Guaranteed Stop Loss Order was initiated at the request of a client
	GuaranteedStopLossOrderReason_CLIENT_ORDER LimitOrderReason = "CLIENT_ORDER"
	// The Guaranteed Stop Loss Order was initiated as a replacement for an existing Order
	GuaranteedStopLossOrderReason_REPLACEMENT LimitOrderReason = "REPLACEMENT"
	// The Guaranteed Stop Loss Order was initiated automatically when an Order was filled that opened a new
	// Trade requiring a Guaranteed Stop Loss Order.
	GuaranteedStopLossOrderReason_ON_FILL LimitOrderReason = "ON_FILL"
)

// The reason that the Trailing Stop Loss Order was initiated
type TrailingStopLossOrderReason string

const (
	// The Trailing Stop Loss Order was initiated at the request of a client
	TrailingStopLossOrderReason_CLIENT_ORDER LimitOrderReason = "CLIENT_ORDER"
	// The Trailing Stop Loss Order was initiated as a replacement for an existing Order
	TrailingStopLossOrderReason_REPLACEMENT LimitOrderReason = "REPLACEMENT"
	// The Trailing Stop Loss Order was initiated automatically when an Order was filled that opened a new
	// Trade requiring a Trailing Stop Loss Order.
	TrailingStopLossOrderReason_ON_FILL LimitOrderReason = "ON_FILL"
)

// The reason that an Order was filled
type OrderFillReason string

const (
	// The Order filled was a Limit Order
	OrderFillReason_LIMIT_ORDER OrderFillReason = "LIMIT_ORDER"
	// The Order filled was a Stop Order
	OrderFillReason_STOP_ORDER OrderFillReason = "STOP_ORDER"
	// The Order filled was a Market-if-touched Order
	OrderFillReason_MARKET_IF_TOUCHED_ORDER OrderFillReason = "MARKET_IF_TOUCHED_ORDER"
	// The Order filled was a Take Profit Order
	OrderFillReason_TAKE_PROFIT_ORDER OrderFillReason = "TAKE_PROFIT_ORDER"
	// The Order filled was a Stop Loss Order
	OrderFillReason_STOP_LOSS_ORDER OrderFillReason = "STOP_LOSS_ORDER"
	// The Order filled was a Guaranteed Stop Loss Order
	OrderFillReason_GUARANTEED_STOP_LOSS_ORDER OrderFillReason = "GUARANTEED_STOP_LOSS_ORDER"
	// The Order filled was a Trailing Stop Loss Order
	OrderFillReason_TRAILING_STOP_LOSS_ORDER OrderFillReason = "TRAILING_STOP_LOSS_ORDER"
	// The Order filled was a Market Order
	OrderFillReason_MARKET_ORDER OrderFillReason = "MARKET_ORDER"
	// The Order filled was a Market Order used to explicitly close a Trade
	OrderFillReason_MARKET_ORDER_TRADE_CLOSE OrderFillReason = "MARKET_ORDER_TRADE_CLOSE"
	// The Order filled was a Market Order used to explicitly close a Position
	OrderFillReason_MARKET_ORDER_POSITION_CLOSEOUT OrderFillReason = "MARKET_ORDER_POSITION_CLOSEOUT"
	// The Order filled was a Market Order used for a Margin Closeout
	OrderFillReason_MARKET_ORDER_MARGIN_CLOSEOUT OrderFillReason = "MARKET_ORDER_MARGIN_CLOSEOUT"
	// The Order filled was a Market Order used for a delayed Trade close
	OrderFillReason_MARKET_ORDER_DELAYED_TRADE_CLOSE OrderFillReason = "MARKET_ORDER_DELAYED_TRADE_CLOSE"
	// The Order filled was a Fixed Price Order
	OrderFillReason_FIXED_PRICE_ORDER OrderFillReason = "FIXED_PRICE_ORDER"
	// The Order filled was a Fixed Price Order created as part of a platform account migration
	OrderFillReason_FIXED_PRICE_ORDER_PLATFORM_ACCOUNT_MIGRATION OrderFillReason = "FIXED_PRICE_ORDER_PLATFORM_ACCOUNT_MIGRATION"
	// The Order filled was a Fixed Price Order created to close a Trade as part of division account migration
	OrderFillReason_FIXED_PRICE_ORDER_DIVISION_ACCOUNT_MIGRATION OrderFillReason = "FIXED_PRICE_ORDER_DIVISION_ACCOUNT_MIGRATION"
	// The Order filled was a Fixed Price Order created to close a Trade administratively
	OrderFillReason_FIXED_PRICE_ORDER_ADMINISTRATIVE_ACTION OrderFillReason = "FIXED_PRICE_ORDER_ADMINISTRATIVE_ACTION"
)

// The reason that an Order was cancelled.
type OrderCancelReason string

const (
	// The Order was cancelled because at the time of filling, an unexpected internal server
	// error occurred.
	OrderCancelReason_INTERNAL_SERVER_ERROR OrderCancelReason = "INTERNAL_SERVER_ERROR"
	// The Order was cancelled because at the time of filling the account was locked.
	OrderCancelReason_ACCOUNT_LOCKED OrderCancelReason = "ACCOUNT_LOCKED"
	// The order was to be filled, however the account is configured to not allow new positions
	// to be created.
	OrderCancelReason_ACCOUNT_NEW_POSITIONS_LOCKED OrderCancelReason = "ACCOUNT_NEW_POSITIONS_LOCKED"
	// Filling the Order wasn’t possible because it required the creation of a dependent
	// Order and the Account is locked for Order creation.
	OrderCancelReason_ACCOUNT_ORDER_CREATION_LOCKED OrderCancelReason = "ACCOUNT_ORDER_CREATION_LOCKED"
	// Filling the Order was not possible because the Account is locked for filling Orders.
	OrderCancelReason_ACCOUNT_ORDER_FILL_LOCKED OrderCancelReason = "ACCOUNT_ORDER_FILL_LOCKED"
	// The Order was cancelled explicitly at the request of the client.
	OrderCancelReason_CLIENT_REQUEST OrderCancelReason = "CLIENT_REQUEST"
	// The Order cancelled because it is being migrated to another account.
	OrderCancelReason_MIGRATION OrderCancelReason = "MIGRATION"
	// Filling the Order wasn’t possible because the Order’s instrument was halted.
	OrderCancelReason_MARKET_HALTED OrderCancelReason = "MARKET_HALTED"
	// The Order is linked to an open Trade that was closed.
	OrderCancelReason_LINKED_TRADE_CLOSED OrderCancelReason = "LINKED_TRADE_CLOSED"
	// The time in force specified for this order has passed.
	OrderCancelReason_TIME_IN_FORCE_EXPIRED OrderCancelReason = "TIME_IN_FORCE_EXPIRED"
	// Filling the Order wasn’t possible because the Account had insufficient margin.
	OrderCancelReason_INSUFFICIENT_MARGIN OrderCancelReason = "INSUFFICIENT_MARGIN"
	// Filling the Order would have resulted in a a FIFO violation.
	OrderCancelReason_FIFO_VIOLATION OrderCancelReason = "FIFO_VIOLATION"
	// Filling the Order would have violated the Order’s price bound.
	OrderCancelReason_BOUNDS_VIOLATION OrderCancelReason = "BOUNDS_VIOLATION"
	// The Order was cancelled for replacement at the request of the client.
	OrderCancelReason_CLIENT_REQUEST_REPLACED OrderCancelReason = "CLIENT_REQUEST_REPLACED"
	// The Order was cancelled for replacement with an adjusted fillPrice to accommodate
	// for the price movement caused by a dividendAdjustment.
	OrderCancelReason_DIVIDEND_ADJUSTMENT_REPLACED OrderCancelReason = "DIVIDEND_ADJUSTMENT_REPLACED"
	// Filling the Order wasn’t possible because enough liquidity available.
	OrderCancelReason_INSUFFICIENT_LIQUIDITY OrderCancelReason = "INSUFFICIENT_LIQUIDITY"
	// Filling the Order would have resulted in the creation of a Take Profit Order with
	// a GTD time in the past.
	OrderCancelReason_TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_IN_PAST OrderCancelReason = "TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	// Filling the Order would result in the creation of a Take Profit Order that would
	// have been filled immediately, closing the new Trade at a loss.
	OrderCancelReason_TAKE_PROFIT_ON_FILL_LOSS OrderCancelReason = "TAKE_PROFIT_ON_FILL_LOSS"
	// Filling the Order would result in the creation of a Take Profit Loss Order that would
	// close the new Trade at a loss when filled.
	OrderCancelReason_LOSING_TAKE_PROFIT OrderCancelReason = "LOSING_TAKE_PROFIT"
	// Filling the Order would have resulted in the creation of a Stop Loss Order with a
	// GTD time in the past.
	OrderCancelReason_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST OrderCancelReason = "STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	// Filling the Order would result in the creation of a Stop Loss Order that would have
	// been filled immediately, closing the new Trade at a loss.
	OrderCancelReason_STOP_LOSS_ON_FILL_LOSS OrderCancelReason = "STOP_LOSS_ON_FILL_LOSS"
	// Filling the Order would result in the creation of a Stop Loss Order whose price would
	// be zero or negative due to the specified distance.
	OrderCancelReason_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED OrderCancelReason = "STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED"
	// Filling the Order would not result in the creation of Stop Loss Order, however the
	// Account’s configuration requires that all Trades have a Stop Loss Order attached
	// to them.
	OrderCancelReason_STOP_LOSS_ON_FILL_REQUIRED OrderCancelReason = "STOP_LOSS_ON_FILL_REQUIRED"
	// Filling the Order would not result in the creation of a guaranteed Stop Loss Order,
	// however the Account’s configuration requires that all Trades have a guaranteed
	// Stop Loss Order attached to them.
	OrderCancelReason_STOP_LOSS_ON_FILL_GUARANTEED_REQUIRED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_REQUIRED"
	// Filling the Order would result in the creation of a guaranteed Stop Loss Order, however
	// the Account’s configuration does not allow guaranteed Stop Loss Orders.
	OrderCancelReason_STOP_LOSS_ON_FILL_GUARANTEED_NOT_ALLOWED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_NOT_ALLOWED"
	// Filling the Order would result in the creation of a guaranteed Stop Loss Order with
	// a distance smaller than the configured minimum distance.
	OrderCancelReason_STOP_LOSS_ON_FILL_GUARANTEED_MINIMUM_DISTANCE_NOT_MET OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_MINIMUM_DISTANCE_NOT_MET"
	// Filling the Order would result in the creation of a guaranteed Stop Loss Order with
	// trigger price and number of units that that violates the account’s guaranteed Stop
	// Loss Order level restriction.
	OrderCancelReason_STOP_LOSS_ON_FILL_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED"
	// Filling the Order would result in the creation of a guaranteed Stop Loss Order for
	// a hedged Trade, however the Account’s configuration does not allow guaranteed Stop
	// Loss Orders for hedged Trades/Positions.
	OrderCancelReason_STOP_LOSS_ON_FILL_GUARANTEED_HEDGING_NOT_ALLOWED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_HEDGING_NOT_ALLOWED"
	// Filling the Order would result in the creation of a Stop Loss Order whose TimeInForce
	// value is invalid. A likely cause would be if the Account requires guaranteed stop
	// loss orders and the TimeInForce value were not GTC.
	OrderCancelReason_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID OrderCancelReason = "STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID"
	// Filling the Order would result in the creation of a Stop Loss Order whose TriggerCondition
	// value is invalid. A likely cause would be if the stop loss order is guaranteed and
	// the TimeInForce is not TRIGGER_DEFAULT or TRIGGER_BID for a long trade, or not TRIGGER_DEFAULT
	// or TRIGGER_ASK for a short trade.
	OrderCancelReason_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID OrderCancelReason = "STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID"
	// Filling the Order would have resulted in the creation of a Guaranteed Stop Loss Order
	// with a GTD time in the past.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order that
	// would have been filled immediately, closing the new Trade at a loss.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_LOSS OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_LOSS"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order whose
	// price would be zero or negative due to the specified distance.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED"
	// Filling the Order would not result in the creation of a Guaranteed Stop Loss Order,
	// however the Account’s configuration requires that all Trades have a Guaranteed
	// Stop Loss Order attached to them.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_REQUIRED OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_REQUIRED"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order, however
	// the Account’s configuration does not allow Guaranteed Stop Loss Orders.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_NOT_ALLOWED OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_NOT_ALLOWED"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order with
	// a distance smaller than the configured minimum distance.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_MINIMUM_DISTANCE_NOT_MET OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_MINIMUM_DISTANCE_NOT_MET"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order with
	// trigger number of units that violates the account’s Guaranteed Stop Loss Order
	// level restriction volume.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_VOLUME_EXCEEDED OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_VOLUME_EXCEEDED"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order with
	// trigger price that violates the account’s Guaranteed Stop Loss Order level restriction
	// price range.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_PRICE_RANGE_EXCEEDED OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_PRICE_RANGE_EXCEEDED"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order for
	// a hedged Trade, however the Account’s configuration does not allow Guaranteed Stop
	// Loss Orders for hedged Trades/Positions.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_HEDGING_NOT_ALLOWED OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_HEDGING_NOT_ALLOWED"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order whose
	// TimeInForce value is invalid. A likely cause would be if the Account requires guaranteed
	// stop loss orders and the TimeInForce value were not GTC.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order whose
	// TriggerCondition value is invalid. A likely cause would be the TimeInForce is not
	// TRIGGER_DEFAULT or TRIGGER_BID for a long trade, or not TRIGGER_DEFAULT or TRIGGER_ASK
	// for a short trade.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID"
	// Filling the Order would result in the creation of a Take Profit Order whose price
	// would be zero or negative due to the specified distance.
	OrderCancelReason_TAKE_PROFIT_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED OrderCancelReason = "TAKE_PROFIT_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED"
	// Filling the Order would have resulted in the creation of a Trailing Stop Loss Order
	// with a GTD time in the past.
	OrderCancelReason_TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST OrderCancelReason = "TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	// Filling the Order would result in the creation of a new Open Trade with a client
	// Trade ID already in use.
	OrderCancelReason_CLIENT_TRADE_ID_ALREADY_EXISTS OrderCancelReason = "CLIENT_TRADE_ID_ALREADY_EXISTS"
	// Closing out a position wasn’t fully possible.
	OrderCancelReason_POSITION_CLOSEOUT_FAILED OrderCancelReason = "POSITION_CLOSEOUT_FAILED"
	// Filling the Order would cause the maximum open trades allowed for the Account to
	// be exceeded.
	OrderCancelReason_OPEN_TRADES_ALLOWED_EXCEEDED OrderCancelReason = "OPEN_TRADES_ALLOWED_EXCEEDED"
	// Filling the Order would have resulted in exceeding the number of pending Orders allowed
	// for the Account.
	OrderCancelReason_PENDING_ORDERS_ALLOWED_EXCEEDED OrderCancelReason = "PENDING_ORDERS_ALLOWED_EXCEEDED"
	// Filling the Order would have resulted in the creation of a Take Profit Order with
	// a client Order ID that is already in use.
	OrderCancelReason_TAKE_PROFIT_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS OrderCancelReason = "TAKE_PROFIT_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS"
	// Filling the Order would have resulted in the creation of a Stop Loss Order with a
	// client Order ID that is already in use.
	OrderCancelReason_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS OrderCancelReason = "STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS"
	// Filling the Order would have resulted in the creation of a Guaranteed Stop Loss Order
	// with a client Order ID that is already in use.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS"
	// Filling the Order would have resulted in the creation of a Trailing Stop Loss Order
	// with a client Order ID that is already in use.
	OrderCancelReason_TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS OrderCancelReason = "TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_ALREADY_EXISTS"
	// Filling the Order would have resulted in the Account’s maximum position size limit
	// being exceeded for the Order’s instrument.
	OrderCancelReason_POSITION_SIZE_EXCEEDED OrderCancelReason = "POSITION_SIZE_EXCEEDED"
	// Filling the Order would result in the creation of a Trade, however there already
	// exists an opposing (hedged) Trade that has a guaranteed Stop Loss Order attached
	// to it. Guaranteed Stop Loss Orders cannot be combined with hedged positions.
	OrderCancelReason_HEDGING_GSLO_VIOLATION OrderCancelReason = "HEDGING_GSLO_VIOLATION"
	// Filling the order would cause the maximum position value allowed for the account
	// to be exceeded. The Order has been cancelled as a result.
	OrderCancelReason_ACCOUNT_POSITION_VALUE_LIMIT_EXCEEDED OrderCancelReason = "ACCOUNT_POSITION_VALUE_LIMIT_EXCEEDED"
	// Filling the order would require the creation of a short trade, however the instrument
	// is configured such that orders being filled using bid prices can only reduce existing
	// positions. New short positions cannot be created, but existing long positions may
	// be reduced or closed.
	OrderCancelReason_INSTRUMENT_BID_REDUCE_ONLY OrderCancelReason = "INSTRUMENT_BID_REDUCE_ONLY"
	// Filling the order would require the creation of a long trade, however the instrument
	// is configured such that orders being filled using ask prices can only reduce existing
	// positions. New long positions cannot be created, but existing short positions may
	// be reduced or closed.
	OrderCancelReason_INSTRUMENT_ASK_REDUCE_ONLY OrderCancelReason = "INSTRUMENT_ASK_REDUCE_ONLY"
	// Filling the order would require using the bid, however the instrument is configured
	// such that the bids are halted, and so no short orders may be filled.
	OrderCancelReason_INSTRUMENT_BID_HALTED OrderCancelReason = "INSTRUMENT_BID_HALTED"
	// Filling the order would require using the ask, however the instrument is configured
	// such that the asks are halted, and so no long orders may be filled.
	OrderCancelReason_INSTRUMENT_ASK_HALTED OrderCancelReason = "INSTRUMENT_ASK_HALTED"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order (GSLO).
	// Since the trade is long the GSLO would be short, however the bid side is currently
	// halted. GSLOs cannot be created in this situation.
	OrderCancelReason_STOP_LOSS_ON_FILL_GUARANTEED_BID_HALTED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_BID_HALTED"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order (GSLO).
	// Since the trade is short the GSLO would be long, however the ask side is currently
	// halted. GSLOs cannot be created in this situation.
	OrderCancelReason_STOP_LOSS_ON_FILL_GUARANTEED_ASK_HALTED OrderCancelReason = "STOP_LOSS_ON_FILL_GUARANTEED_ASK_HALTED"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order (GSLO).
	// Since the trade is long the GSLO would be short, however the bid side is currently
	// halted. GSLOs cannot be created in this situation.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_BID_HALTED OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_BID_HALTED"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order (GSLO).
	// Since the trade is short the GSLO would be long, however the ask side is currently
	// halted. GSLOs cannot be created in this situation.
	OrderCancelReason_GUARANTEED_STOP_LOSS_ON_FILL_ASK_HALTED OrderCancelReason = "GUARANTEED_STOP_LOSS_ON_FILL_ASK_HALTED"
	// Filling the Order would have resulted in a new Trade that violates the FIFO violation
	// safeguard constraints.
	OrderCancelReason_FIFO_VIOLATION_SAFEGUARD_VIOLATION OrderCancelReason = "FIFO_VIOLATION_SAFEGUARD_VIOLATION"
	// Filling the Order would have reduced an existing Trade such that the reduced Trade
	// violates the FIFO violation safeguard constraints.
	OrderCancelReason_FIFO_VIOLATION_SAFEGUARD_PARTIAL_CLOSE_VIOLATION OrderCancelReason = "FIFO_VIOLATION_SAFEGUARD_PARTIAL_CLOSE_VIOLATION"
	// The Orders on fill would be in violation of the risk management Order mutual exclusivity
	// configuration specifying that only one risk management Order can be attached to a
	// Trade.
	OrderCancelReason_ORDERS_ON_FILL_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION OrderCancelReason = "ORDERS_ON_FILL_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION"
)

// Used to pay or collect a dividend adjustment amount for an open Trade within the Account.
type OpenTradeDividendAdjustment struct {
	// The ID of the Trade for which the dividend adjustment is to be paid or collected.
	TradeID TradeID `json:"tradeId"`
	// The dividend adjustment amount to pay or collect for the Trade.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`
	// The dividend adjustment amount to pay or collect for the Trade, in the Instrument’s quote currency.
	QuoteDividendAdjustment DecimalNumber `json:"quoteDividendAdjustment"`
}

// A LiquidityRegenerationSchedule indicates how liquidity that is used when filling an Order
// for an instrument is regenerated following the fill. A liquidity regeneration schedule will
// be in effect until the timestamp of its final step, but may be replaced by a schedule created
// for an Order of the same instrument that is filled while it is still in effect.
type LiquidityRegenerationSchedule struct {
	Steps []*LiquidityRegenerationScheduleStep `json:"steps"`
}

// A liquidity regeneration schedule Step indicates the amount of bid and ask liquidity that is
// used by the Account at a certain time. These amounts will only change at the timestamp of the
// following step.
type LiquidityRegenerationScheduleStep struct {
	// The timestamp of the schedule step.
	Timestamp DateTime `json:"timestamp"`
	// The amount of bid liquidity used at this step in the schedule.
	BidLiquidityUsed DecimalNumber `json:"bidLiquidityUsed"`
	// The amount of ask liquidity used at this step in the schedule.
	AskLiquidityUsed DecimalNumber `json:"askLiquidityUsed"`
}

// OpenTradeFinancing is used to pay/collect daily financing charge for an open Trade within an Account
type OpenTradeFinancing struct {
	// The ID of the Trade that financing is being paid/collected for.
	TradeID TradeID `json:"tradeID"`
	// The amount of financing paid/collected for the Trade.
	Financing AccountUnits `json:"financing"`
	// The amount of financing paid/collected in the Instrument’s base currency for the Trade.
	BaseFinancing DecimalNumber `json:"baseFinancing"`
	// The amount of financing paid/collected in the Instrument’s quote currency for the Trade.
	QuoteFinancing DecimalNumber `json:"quoteFinancing"`
	// The financing rate in effect for the instrument used to calculate the the
	// amount of financing paid/collected for the Trade. This field will only be
	// set if the AccountFinancingMode at the time of the daily financing is
	// DAILY_INSTRUMENT or SECOND_BY_SECOND_INSTRUMENT. The value is in decimal
	// rather than percentage points, e.g. 5% is represented as 0.05.
	FinancingRate DecimalNumber `json:"financingRate"`
}

// PositionFinancing is used to pay/collect daily financing charge for a Position within an Account
type PositionFinancing struct {
	// The instrument of the Position that financing is being paid/collected for.
	Instrument InstrumentName `json:"instrument"`
	// The amount of financing paid/collected for the Position.
	Financing AccountUnits `json:"financing"`
	// The amount of base financing paid/collected for the Position.
	BaseFinancing DecimalNumber `json:"baseFinancing"`
	// The amount of quote financing paid/collected for the Position.
	QuoteFinancing DecimalNumber `json:"quoteFinancing"`
	// The HomeConversionFactors in effect for the Position’s Instrument at the
	// time of the DailyFinancing.
	HomeConversionFactors *HomeConversionFactors `json:"homeConversionFactors"`
	// The financing paid/collected for each open Trade within the Position.
	OpenTradeFinancings []*OpenTradeFinancing `json:"openTradeFinancings"`
	// The account financing mode at the time of the daily financing.
	AccountFinancingMode AccountFinancingMode `json:"accountFinancingMode"`
}

// The reason that a Transaction was rejected.
type TransactionRejectReason string

const (
	// An unexpected internal server error has occurred
	TransactionRejectReason_INTERNAL_SERVER_ERROR TransactionRejectReason = "INTERNAL_SERVER_ERROR"
	// The system was unable to determine the current price for the Order’s instrument
	TransactionRejectReason_INSTRUMENT_PRICE_UNKNOWN TransactionRejectReason = "INSTRUMENT_PRICE_UNKNOWN"
	// The Account is not active
	TransactionRejectReason_ACCOUNT_NOT_ACTIVE TransactionRejectReason = "ACCOUNT_NOT_ACTIVE"
	// The Account is locked
	TransactionRejectReason_ACCOUNT_LOCKED TransactionRejectReason = "ACCOUNT_LOCKED"
	// The Account is locked for Order creation
	TransactionRejectReason_ACCOUNT_ORDER_CREATION_LOCKED TransactionRejectReason = "ACCOUNT_ORDER_CREATION_LOCKED"
	// The Account is locked for configuration
	TransactionRejectReason_ACCOUNT_CONFIGURATION_LOCKED TransactionRejectReason = "ACCOUNT_CONFIGURATION_LOCKED"
	// The Account is locked for deposits
	TransactionRejectReason_ACCOUNT_DEPOSIT_LOCKED TransactionRejectReason = "ACCOUNT_DEPOSIT_LOCKED"
	// The Account is locked for withdrawals
	TransactionRejectReason_ACCOUNT_WITHDRAWAL_LOCKED TransactionRejectReason = "ACCOUNT_WITHDRAWAL_LOCKED"
	// The Account is locked for Order cancellation
	TransactionRejectReason_ACCOUNT_ORDER_CANCEL_LOCKED TransactionRejectReason = "ACCOUNT_ORDER_CANCEL_LOCKED"
	// The instrument specified is not tradeable by the Account
	TransactionRejectReason_INSTRUMENT_NOT_TRADEABLE TransactionRejectReason = "INSTRUMENT_NOT_TRADEABLE"
	// Creating the Order would result in the maximum number of allowed pending Orders being
	// exceeded
	TransactionRejectReason_PENDING_ORDERS_ALLOWED_EXCEEDED TransactionRejectReason = "PENDING_ORDERS_ALLOWED_EXCEEDED"
	// Neither the Order ID nor client Order ID are specified
	TransactionRejectReason_ORDER_ID_UNSPECIFIED TransactionRejectReason = "ORDER_ID_UNSPECIFIED"
	// The Order specified does not exist
	TransactionRejectReason_ORDER_DOESNT_EXIST TransactionRejectReason = "ORDER_DOESNT_EXIST"
	// The Order ID and client Order ID specified do not identify the same Order
	TransactionRejectReason_ORDER_IDENTIFIER_INCONSISTENCY TransactionRejectReason = "ORDER_IDENTIFIER_INCONSISTENCY"
	// Neither the Trade ID nor client Trade ID are specified
	TransactionRejectReason_TRADE_ID_UNSPECIFIED TransactionRejectReason = "TRADE_ID_UNSPECIFIED"
	// The Trade specified does not exist
	TransactionRejectReason_TRADE_DOESNT_EXIST TransactionRejectReason = "TRADE_DOESNT_EXIST"
	// The Trade ID and client Trade ID specified do not identify the same Trade
	TransactionRejectReason_TRADE_IDENTIFIER_INCONSISTENCY TransactionRejectReason = "TRADE_IDENTIFIER_INCONSISTENCY"
	// The Account had insufficient margin to perform the action specified. One possible
	// reason for this is due to the creation or modification of a guaranteed StopLoss Order.
	TransactionRejectReason_INSUFFICIENT_MARGIN TransactionRejectReason = "INSUFFICIENT_MARGIN"
	// Order instrument has not been specified
	TransactionRejectReason_INSTRUMENT_MISSING TransactionRejectReason = "INSTRUMENT_MISSING"
	// The instrument specified is unknown
	TransactionRejectReason_INSTRUMENT_UNKNOWN TransactionRejectReason = "INSTRUMENT_UNKNOWN"
	// Order units have not been not specified
	TransactionRejectReason_UNITS_MISSING TransactionRejectReason = "UNITS_MISSING"
	// Order units specified are invalid
	TransactionRejectReason_UNITS_INVALID TransactionRejectReason = "UNITS_INVALID"
	// The units specified contain more precision than is allowed for the Order’s instrument
	TransactionRejectReason_UNITS_PRECISION_EXCEEDED TransactionRejectReason = "UNITS_PRECISION_EXCEEDED"
	// The units specified exceeds the maximum number of units allowed
	TransactionRejectReason_UNITS_LIMIT_EXCEEDED TransactionRejectReason = "UNITS_LIMIT_EXCEEDED"
	// The units specified is less than the minimum number of units required
	TransactionRejectReason_UNITS_MINIMUM_NOT_MET TransactionRejectReason = "UNITS_MINIMUM_NOT_MET"
	// The price has not been specified
	TransactionRejectReason_PRICE_MISSING TransactionRejectReason = "PRICE_MISSING"
	// The price specified is invalid
	TransactionRejectReason_PRICE_INVALID TransactionRejectReason = "PRICE_INVALID"
	// The price specified contains more precision than is allowed for the instrument
	TransactionRejectReason_PRICE_PRECISION_EXCEEDED TransactionRejectReason = "PRICE_PRECISION_EXCEEDED"
	// The price distance has not been specified
	TransactionRejectReason_PRICE_DISTANCE_MISSING TransactionRejectReason = "PRICE_DISTANCE_MISSING"
	// The price distance specified is invalid
	TransactionRejectReason_PRICE_DISTANCE_INVALID TransactionRejectReason = "PRICE_DISTANCE_INVALID"
	// The price distance specified contains more precision than is allowed for the instrument
	TransactionRejectReason_PRICE_DISTANCE_PRECISION_EXCEEDED TransactionRejectReason = "PRICE_DISTANCE_PRECISION_EXCEEDED"
	// The price distance exceeds that maximum allowed amount
	TransactionRejectReason_PRICE_DISTANCE_MAXIMUM_EXCEEDED TransactionRejectReason = "PRICE_DISTANCE_MAXIMUM_EXCEEDED"
	// The price distance does not meet the minimum allowed amount
	TransactionRejectReason_PRICE_DISTANCE_MINIMUM_NOT_MET TransactionRejectReason = "PRICE_DISTANCE_MINIMUM_NOT_MET"
	// The TimeInForce field has not been specified
	TransactionRejectReason_TIME_IN_FORCE_MISSING TransactionRejectReason = "TIME_IN_FORCE_MISSING"
	// The TimeInForce specified is invalid
	TransactionRejectReason_TIME_IN_FORCE_INVALID TransactionRejectReason = "TIME_IN_FORCE_INVALID"
	// The TimeInForce is GTD but no GTD timestamp is provided
	TransactionRejectReason_TIME_IN_FORCE_GTD_TIMESTAMP_MISSING TransactionRejectReason = "TIME_IN_FORCE_GTD_TIMESTAMP_MISSING"
	// The TimeInForce is GTD but the GTD timestamp is in the past
	TransactionRejectReason_TIME_IN_FORCE_GTD_TIMESTAMP_IN_PAST TransactionRejectReason = "TIME_IN_FORCE_GTD_TIMESTAMP_IN_PAST"
	// The price bound specified is invalid
	TransactionRejectReason_PRICE_BOUND_INVALID TransactionRejectReason = "PRICE_BOUND_INVALID"
	// The price bound specified contains more precision than is allowed for the Order’s
	// instrument
	TransactionRejectReason_PRICE_BOUND_PRECISION_EXCEEDED TransactionRejectReason = "PRICE_BOUND_PRECISION_EXCEEDED"
	// Multiple Orders on fill share the same client Order ID
	TransactionRejectReason_ORDERS_ON_FILL_DUPLICATE_CLIENT_ORDER_IDS TransactionRejectReason = "ORDERS_ON_FILL_DUPLICATE_CLIENT_ORDER_IDS"
	// The Order does not support Trade on fill client extensions because it cannot create
	// a new Trade
	TransactionRejectReason_TRADE_ON_FILL_CLIENT_EXTENSIONS_NOT_SUPPORTED TransactionRejectReason = "TRADE_ON_FILL_CLIENT_EXTENSIONS_NOT_SUPPORTED"
	// The client Order ID specified is invalid
	TransactionRejectReason_CLIENT_ORDER_ID_INVALID TransactionRejectReason = "CLIENT_ORDER_ID_INVALID"
	// The client Order ID specified is already assigned to another pending Order
	TransactionRejectReason_CLIENT_ORDER_ID_ALREADY_EXISTS TransactionRejectReason = "CLIENT_ORDER_ID_ALREADY_EXISTS"
	// The client Order tag specified is invalid
	TransactionRejectReason_CLIENT_ORDER_TAG_INVALID TransactionRejectReason = "CLIENT_ORDER_TAG_INVALID"
	// The client Order comment specified is invalid
	TransactionRejectReason_CLIENT_ORDER_COMMENT_INVALID TransactionRejectReason = "CLIENT_ORDER_COMMENT_INVALID"
	// The client Trade ID specified is invalid
	TransactionRejectReason_CLIENT_TRADE_ID_INVALID TransactionRejectReason = "CLIENT_TRADE_ID_INVALID"
	// The client Trade ID specified is already assigned to another open Trade
	TransactionRejectReason_CLIENT_TRADE_ID_ALREADY_EXISTS TransactionRejectReason = "CLIENT_TRADE_ID_ALREADY_EXISTS"
	// The client Trade tag specified is invalid
	TransactionRejectReason_CLIENT_TRADE_TAG_INVALID TransactionRejectReason = "CLIENT_TRADE_TAG_INVALID"
	// The client Trade comment is invalid
	TransactionRejectReason_CLIENT_TRADE_COMMENT_INVALID TransactionRejectReason = "CLIENT_TRADE_COMMENT_INVALID"
	// The OrderFillPositionAction field has not been specified
	TransactionRejectReason_ORDER_FILL_POSITION_ACTION_MISSING TransactionRejectReason = "ORDER_FILL_POSITION_ACTION_MISSING"
	// The OrderFillPositionAction specified is invalid
	TransactionRejectReason_ORDER_FILL_POSITION_ACTION_INVALID TransactionRejectReason = "ORDER_FILL_POSITION_ACTION_INVALID"
	// The TriggerCondition field has not been specified
	TransactionRejectReason_TRIGGER_CONDITION_MISSING TransactionRejectReason = "TRIGGER_CONDITION_MISSING"
	// The TriggerCondition specified is invalid
	TransactionRejectReason_TRIGGER_CONDITION_INVALID TransactionRejectReason = "TRIGGER_CONDITION_INVALID"
	// The OrderFillPositionAction field has not been specified
	TransactionRejectReason_ORDER_PARTIAL_FILL_OPTION_MISSING TransactionRejectReason = "ORDER_PARTIAL_FILL_OPTION_MISSING"
	// The OrderFillPositionAction specified is invalid.
	TransactionRejectReason_ORDER_PARTIAL_FILL_OPTION_INVALID TransactionRejectReason = "ORDER_PARTIAL_FILL_OPTION_INVALID"
	// When attempting to reissue an order (currently only a MarketIfTouched) that was immediately
	// partially filled, it is not possible to create a correct pending Order.
	TransactionRejectReason_INVALID_REISSUE_IMMEDIATE_PARTIAL_FILL TransactionRejectReason = "INVALID_REISSUE_IMMEDIATE_PARTIAL_FILL"
	// The Orders on fill would be in violation of the risk management Order mutual exclusivity
	// configuration specifying that only one risk management Order can be attached to a
	// Trade.
	TransactionRejectReason_ORDERS_ON_FILL_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION TransactionRejectReason = "ORDERS_ON_FILL_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION"
	// The Orders on fill would be in violation of the risk management Order mutual exclusivity
	// configuration specifying that if a GSLO is already attached to a Trade, no other
	// risk management Order can be attached to a Trade.
	TransactionRejectReason_ORDERS_ON_FILL_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION TransactionRejectReason = "ORDERS_ON_FILL_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION"
	// A Take Profit Order for the specified Trade already exists
	TransactionRejectReason_TAKE_PROFIT_ORDER_ALREADY_EXISTS TransactionRejectReason = "TAKE_PROFIT_ORDER_ALREADY_EXISTS"
	// The Take Profit Order would cause the associated Trade to be in violation of the
	// FIFO violation safeguard constraints.
	TransactionRejectReason_TAKE_PROFIT_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD TransactionRejectReason = "TAKE_PROFIT_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD"
	// The Take Profit on fill specified does not provide a price
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_PRICE_MISSING TransactionRejectReason = "TAKE_PROFIT_ON_FILL_PRICE_MISSING"
	// The Take Profit on fill specified contains an invalid price
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_PRICE_INVALID TransactionRejectReason = "TAKE_PROFIT_ON_FILL_PRICE_INVALID"
	// The Take Profit on fill specified contains a price with more precision than is allowed
	// by the Order’s instrument
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_PRICE_PRECISION_EXCEEDED TransactionRejectReason = "TAKE_PROFIT_ON_FILL_PRICE_PRECISION_EXCEEDED"
	// The Take Profit on fill specified does not provide a TimeInForce
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_TIME_IN_FORCE_MISSING TransactionRejectReason = "TAKE_PROFIT_ON_FILL_TIME_IN_FORCE_MISSING"
	// The Take Profit on fill specifies an invalid TimeInForce
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_TIME_IN_FORCE_INVALID TransactionRejectReason = "TAKE_PROFIT_ON_FILL_TIME_IN_FORCE_INVALID"
	// The Take Profit on fill specifies a GTD TimeInForce but does not provide a GTD timestamp
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_MISSING TransactionRejectReason = "TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_MISSING"
	// The Take Profit on fill specifies a GTD timestamp that is in the past
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_IN_PAST TransactionRejectReason = "TAKE_PROFIT_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	// The Take Profit on fill client Order ID specified is invalid
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_CLIENT_ORDER_ID_INVALID TransactionRejectReason = "TAKE_PROFIT_ON_FILL_CLIENT_ORDER_ID_INVALID"
	// The Take Profit on fill client Order tag specified is invalid
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_CLIENT_ORDER_TAG_INVALID TransactionRejectReason = "TAKE_PROFIT_ON_FILL_CLIENT_ORDER_TAG_INVALID"
	// The Take Profit on fill client Order comment specified is invalid
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_CLIENT_ORDER_COMMENT_INVALID TransactionRejectReason = "TAKE_PROFIT_ON_FILL_CLIENT_ORDER_COMMENT_INVALID"
	// The Take Profit on fill specified does not provide a TriggerCondition
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_TRIGGER_CONDITION_MISSING TransactionRejectReason = "TAKE_PROFIT_ON_FILL_TRIGGER_CONDITION_MISSING"
	// The Take Profit on fill specifies an invalid TriggerCondition
	TransactionRejectReason_TAKE_PROFIT_ON_FILL_TRIGGER_CONDITION_INVALID TransactionRejectReason = "TAKE_PROFIT_ON_FILL_TRIGGER_CONDITION_INVALID"
	// A Stop Loss Order for the specified Trade already exists
	TransactionRejectReason_STOP_LOSS_ORDER_ALREADY_EXISTS TransactionRejectReason = "STOP_LOSS_ORDER_ALREADY_EXISTS"
	// An attempt was made to to create a non-guaranteed stop loss order in an account that
	// requires all stop loss orders to be guaranteed.
	TransactionRejectReason_STOP_LOSS_ORDER_GUARANTEED_REQUIRED TransactionRejectReason = "STOP_LOSS_ORDER_GUARANTEED_REQUIRED"
	// An attempt to create a guaranteed stop loss order with a price that is within the
	// current tradeable spread.
	TransactionRejectReason_STOP_LOSS_ORDER_GUARANTEED_PRICE_WITHIN_SPREAD TransactionRejectReason = "STOP_LOSS_ORDER_GUARANTEED_PRICE_WITHIN_SPREAD"
	// An attempt was made to create a guaranteed Stop Loss Order, however the Account’s
	// configuration does not allow guaranteed Stop Loss Orders.
	TransactionRejectReason_STOP_LOSS_ORDER_GUARANTEED_NOT_ALLOWED TransactionRejectReason = "STOP_LOSS_ORDER_GUARANTEED_NOT_ALLOWED"
	// An attempt was made to create a guaranteed Stop Loss Order when the market was halted.
	TransactionRejectReason_STOP_LOSS_ORDER_GUARANTEED_HALTED_CREATE_VIOLATION TransactionRejectReason = "STOP_LOSS_ORDER_GUARANTEED_HALTED_CREATE_VIOLATION"
	// An attempt was made to re-create a guaranteed Stop Loss Order with a tighter fill
	// price when the market was halted.
	TransactionRejectReason_STOP_LOSS_ORDER_GUARANTEED_HALTED_TIGHTEN_VIOLATION TransactionRejectReason = "STOP_LOSS_ORDER_GUARANTEED_HALTED_TIGHTEN_VIOLATION"
	// An attempt was made to create a guaranteed Stop Loss Order on a hedged Trade (ie
	// there is an existing open Trade in the opposing direction), however the Account’s
	// configuration does not allow guaranteed Stop Loss Orders for hedged Trades/Positions.
	TransactionRejectReason_STOP_LOSS_ORDER_GUARANTEED_HEDGING_NOT_ALLOWED TransactionRejectReason = "STOP_LOSS_ORDER_GUARANTEED_HEDGING_NOT_ALLOWED"
	// An attempt was made to create a guaranteed Stop Loss Order, however the distance
	// between the current price and the trigger price does not meet the Account’s configured
	// minimum Guaranteed Stop Loss distance.
	TransactionRejectReason_STOP_LOSS_ORDER_GUARANTEED_MINIMUM_DISTANCE_NOT_MET TransactionRejectReason = "STOP_LOSS_ORDER_GUARANTEED_MINIMUM_DISTANCE_NOT_MET"
	// An attempt was made to cancel a Stop Loss Order, however the Account’s configuration
	// requires every Trade have an associated Stop Loss Order.
	TransactionRejectReason_STOP_LOSS_ORDER_NOT_CANCELABLE TransactionRejectReason = "STOP_LOSS_ORDER_NOT_CANCELABLE"
	// An attempt was made to cancel and replace a Stop Loss Order, however the Account’s
	// configuration prevents the modification of Stop Loss Orders.
	TransactionRejectReason_STOP_LOSS_ORDER_NOT_REPLACEABLE TransactionRejectReason = "STOP_LOSS_ORDER_NOT_REPLACEABLE"
	// An attempt was made to create a guaranteed Stop Loss Order, however doing so would
	// exceed the Account’s configured guaranteed StopLoss Order level restriction volume.
	TransactionRejectReason_STOP_LOSS_ORDER_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED TransactionRejectReason = "STOP_LOSS_ORDER_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED"
	// The Stop Loss Order request contains both the price and distance fields.
	TransactionRejectReason_STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_SPECIFIED TransactionRejectReason = "STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_SPECIFIED"
	// The Stop Loss Order request contains neither the price nor distance fields.
	TransactionRejectReason_STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_MISSING TransactionRejectReason = "STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_MISSING"
	// The Stop Loss Order would cause the associated Trade to be in violation of the FIFO
	// violation safeguard constraints
	TransactionRejectReason_STOP_LOSS_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD TransactionRejectReason = "STOP_LOSS_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD"
	// The Stop Loss Order would be in violation of the risk management Order mutual exclusivity
	// configuration specifying that only one risk management order can be attached to a
	// Trade.
	TransactionRejectReason_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION TransactionRejectReason = "STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION"
	// The Stop Loss Order would be in violation of the risk management Order mutual exclusivity
	// configuration specifying that if a GSLO is already attached to a Trade, no other
	// risk management Order can be attached to the same Trade.
	TransactionRejectReason_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION TransactionRejectReason = "STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION"
	// An attempt to create a pending Order was made with no Stop Loss Order on fill specified
	// and the Account’s configuration requires that every Trade have an associated Stop
	// Loss Order.
	TransactionRejectReason_STOP_LOSS_ON_FILL_REQUIRED_FOR_PENDING_ORDER TransactionRejectReason = "STOP_LOSS_ON_FILL_REQUIRED_FOR_PENDING_ORDER"
	// An attempt to create a pending Order was made with a Stop Loss Order on fill that
	// was explicitly configured to be guaranteed, however the Account’s configuration
	// does not allow guaranteed Stop Loss Orders.
	TransactionRejectReason_STOP_LOSS_ON_FILL_GUARANTEED_NOT_ALLOWED TransactionRejectReason = "STOP_LOSS_ON_FILL_GUARANTEED_NOT_ALLOWED"
	// An attempt to create a pending Order was made with a Stop Loss Order on fill that
	// was explicitly configured to be not guaranteed, however the Account’s configuration
	// requires guaranteed Stop Loss Orders.
	TransactionRejectReason_STOP_LOSS_ON_FILL_GUARANTEED_REQUIRED TransactionRejectReason = "STOP_LOSS_ON_FILL_GUARANTEED_REQUIRED"
	// The Stop Loss on fill specified does not provide a price
	TransactionRejectReason_STOP_LOSS_ON_FILL_PRICE_MISSING TransactionRejectReason = "STOP_LOSS_ON_FILL_PRICE_MISSING"
	// The Stop Loss on fill specifies an invalid price
	TransactionRejectReason_STOP_LOSS_ON_FILL_PRICE_INVALID TransactionRejectReason = "STOP_LOSS_ON_FILL_PRICE_INVALID"
	// The Stop Loss on fill specifies a price with more precision than is allowed by the
	// Order’s instrument
	TransactionRejectReason_STOP_LOSS_ON_FILL_PRICE_PRECISION_EXCEEDED TransactionRejectReason = "STOP_LOSS_ON_FILL_PRICE_PRECISION_EXCEEDED"
	// An attempt to create a pending Order was made with the distance between the guaranteed
	// Stop Loss Order on fill’s price and the pending Order’s price is less than the
	// Account’s configured minimum guaranteed stop loss distance.
	TransactionRejectReason_STOP_LOSS_ON_FILL_GUARANTEED_MINIMUM_DISTANCE_NOT_MET TransactionRejectReason = "STOP_LOSS_ON_FILL_GUARANTEED_MINIMUM_DISTANCE_NOT_MET"
	// An attempt to create a pending Order was made with a guaranteed Stop Loss Order on
	// fill configured, and the Order’s units exceed the Account’s configured guaranteed
	// StopLoss Order level restriction volume.
	TransactionRejectReason_STOP_LOSS_ON_FILL_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED TransactionRejectReason = "STOP_LOSS_ON_FILL_GUARANTEED_LEVEL_RESTRICTION_EXCEEDED"
	// The Stop Loss on fill distance is invalid
	TransactionRejectReason_STOP_LOSS_ON_FILL_DISTANCE_INVALID TransactionRejectReason = "STOP_LOSS_ON_FILL_DISTANCE_INVALID"
	// The Stop Loss on fill price distance exceeds the maximum allowed amount
	TransactionRejectReason_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED TransactionRejectReason = "STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED"
	// The Stop Loss on fill distance contains more precision than is allowed by the instrument
	TransactionRejectReason_STOP_LOSS_ON_FILL_DISTANCE_PRECISION_EXCEEDED TransactionRejectReason = "STOP_LOSS_ON_FILL_DISTANCE_PRECISION_EXCEEDED"
	// The Stop Loss on fill contains both the price and distance fields.
	TransactionRejectReason_STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_SPECIFIED TransactionRejectReason = "STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_SPECIFIED"
	// The Stop Loss on fill contains neither the price nor distance fields.
	TransactionRejectReason_STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_MISSING TransactionRejectReason = "STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_MISSING"
	// The Stop Loss on fill specified does not provide a TimeInForce
	TransactionRejectReason_STOP_LOSS_ON_FILL_TIME_IN_FORCE_MISSING TransactionRejectReason = "STOP_LOSS_ON_FILL_TIME_IN_FORCE_MISSING"
	// The Stop Loss on fill specifies an invalid TimeInForce
	TransactionRejectReason_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID TransactionRejectReason = "STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID"
	// The Stop Loss on fill specifies a GTD TimeInForce but does not provide a GTD timestamp
	TransactionRejectReason_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_MISSING TransactionRejectReason = "STOP_LOSS_ON_FILL_GTD_TIMESTAMP_MISSING"
	// The Stop Loss on fill specifies a GTD timestamp that is in the past
	TransactionRejectReason_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST TransactionRejectReason = "STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	// The Stop Loss on fill client Order ID specified is invalid
	TransactionRejectReason_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_INVALID TransactionRejectReason = "STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_INVALID"
	// The Stop Loss on fill client Order tag specified is invalid
	TransactionRejectReason_STOP_LOSS_ON_FILL_CLIENT_ORDER_TAG_INVALID TransactionRejectReason = "STOP_LOSS_ON_FILL_CLIENT_ORDER_TAG_INVALID"
	// The Stop Loss on fill client Order comment specified is invalid
	TransactionRejectReason_STOP_LOSS_ON_FILL_CLIENT_ORDER_COMMENT_INVALID TransactionRejectReason = "STOP_LOSS_ON_FILL_CLIENT_ORDER_COMMENT_INVALID"
	// The Stop Loss on fill specified does not provide a TriggerCondition
	TransactionRejectReason_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_MISSING TransactionRejectReason = "STOP_LOSS_ON_FILL_TRIGGER_CONDITION_MISSING"
	// The Stop Loss on fill specifies an invalid TriggerCondition
	TransactionRejectReason_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID TransactionRejectReason = "STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID"
	// A Guaranteed Stop Loss Order for the specified Trade already exists
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_ALREADY_EXISTS TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_ALREADY_EXISTS"
	// An attempt was made to to create a non-guaranteed stop loss order in an account that
	// requires all stop loss orders to be guaranteed.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_REQUIRED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_REQUIRED"
	// An attempt to create a guaranteed stop loss order with a price that is within the
	// current tradeable spread.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_PRICE_WITHIN_SPREAD TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_PRICE_WITHIN_SPREAD"
	// An attempt was made to create a Guaranteed Stop Loss Order, however the Account’s
	// configuration does not allow Guaranteed Stop Loss Orders.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_NOT_ALLOWED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_NOT_ALLOWED"
	// An attempt was made to create a Guaranteed Stop Loss Order when the market was halted.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_HALTED_CREATE_VIOLATION TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_HALTED_CREATE_VIOLATION"
	// An attempt was made to create a Guaranteed Stop Loss Order when the market was open.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_CREATE_VIOLATION TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_CREATE_VIOLATION"
	// An attempt was made to re-create a Guaranteed Stop Loss Order with a tighter fill
	// price when the market was halted.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_HALTED_TIGHTEN_VIOLATION TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_HALTED_TIGHTEN_VIOLATION"
	// An attempt was made to re-create a Guaranteed Stop Loss Order with a tighter fill
	// price when the market was open.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_TIGHTEN_VIOLATION TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_TIGHTEN_VIOLATION"
	// An attempt was made to create a Guaranteed Stop Loss Order on a hedged Trade (ie
	// there is an existing open Trade in the opposing direction), however the Account’s
	// configuration does not allow Guaranteed Stop Loss Orders for hedged Trades/Positions.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_HEDGING_NOT_ALLOWED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_HEDGING_NOT_ALLOWED"
	// An attempt was made to create a Guaranteed Stop Loss Order, however the distance
	// between the current price and the trigger price does not meet the Account’s configured
	// minimum Guaranteed Stop Loss distance.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_MINIMUM_DISTANCE_NOT_MET TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_MINIMUM_DISTANCE_NOT_MET"
	// An attempt was made to cancel a Guaranteed Stop Loss Order when the market is open,
	// however the Account’s configuration requires every Trade have an associated Guaranteed
	// Stop Loss Order.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_NOT_CANCELABLE TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_NOT_CANCELABLE"
	// An attempt was made to cancel a Guaranteed Stop Loss Order when the market is halted,
	// however the Account’s configuration requires every Trade have an associated Guaranteed
	// Stop Loss Order.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_HALTED_NOT_CANCELABLE TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_HALTED_NOT_CANCELABLE"
	// An attempt was made to cancel and replace a Guaranteed Stop Loss Order when the market
	// is open, however the Account’s configuration prevents the modification of Guaranteed
	// Stop Loss Orders.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_NOT_REPLACEABLE TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_NOT_REPLACEABLE"
	// An attempt was made to cancel and replace a Guaranteed Stop Loss Order when the market
	// is halted, however the Account’s configuration prevents the modification of Guaranteed
	// Stop Loss Orders.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_HALTED_NOT_REPLACEABLE TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_HALTED_NOT_REPLACEABLE"
	// An attempt was made to create a Guaranteed Stop Loss Order, however doing so would
	// exceed the Account’s configured guaranteed StopLoss Order level restriction volume.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_LEVEL_RESTRICTION_VOLUME_EXCEEDED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_LEVEL_RESTRICTION_VOLUME_EXCEEDED"
	// An attempt was made to create a Guaranteed Stop Loss Order, however doing so would
	// exceed the Account’s configured guaranteed StopLoss Order level restriction price
	// range.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_LEVEL_RESTRICTION_PRICE_RANGE_EXCEEDED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_LEVEL_RESTRICTION_PRICE_RANGE_EXCEEDED"
	// The Guaranteed Stop Loss Order request contains both the price and distance fields.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_SPECIFIED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_SPECIFIED"
	// The Guaranteed Stop Loss Order request contains neither the price nor distance fields.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_MISSING TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_PRICE_AND_DISTANCE_BOTH_MISSING"
	// The Guaranteed Stop Loss Order would cause the associated Trade to be in violation
	// of the FIFO violation safeguard constraints
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD"
	// The Guaranteed Stop Loss Order would be in violation of the risk management Order
	// mutual exclusivity configuration specifying that only one risk management order can
	// be attached to a Trade.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION"
	// The Guaranteed Stop Loss Order would be in violation of the risk management Order
	// mutual exclusivity configuration specifying that if a GSLO is already attached to
	// a Trade, no other risk management Order can be attached to the same Trade.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION TransactionRejectReason = "GUARANTEED_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION"
	// An attempt to create a pending Order was made with no Guaranteed Stop Loss Order
	// on fill specified and the Account’s configuration requires that every Trade have
	// an associated Guaranteed Stop Loss Order.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_REQUIRED_FOR_PENDING_ORDER TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_REQUIRED_FOR_PENDING_ORDER"
	// An attempt to create a pending Order was made with a Guaranteed Stop Loss Order on
	// fill that was explicitly configured to be guaranteed, however the Account’s configuration
	// does not allow guaranteed Stop Loss Orders.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_NOT_ALLOWED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_NOT_ALLOWED"
	// An attempt to create a pending Order was made with a Guaranteed Stop Loss Order on
	// fill that was explicitly configured to be not guaranteed, however the Account’s
	// configuration requires Guaranteed Stop Loss Orders.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_REQUIRED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_REQUIRED"
	// The Guaranteed Stop Loss on fill specified does not provide a price
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_PRICE_MISSING TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_PRICE_MISSING"
	// The Guaranteed Stop Loss on fill specifies an invalid price
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_PRICE_INVALID TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_PRICE_INVALID"
	// The Guaranteed Stop Loss on fill specifies a price with more precision than is allowed
	// by the Order’s instrument
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_PRICE_PRECISION_EXCEEDED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_PRICE_PRECISION_EXCEEDED"
	// An attempt to create a pending Order was made with the distance between the Guaranteed
	// Stop Loss Order on fill’s price and the pending Order’s price is less than the
	// Account’s configured minimum guaranteed stop loss distance.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_MINIMUM_DISTANCE_NOT_MET TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_MINIMUM_DISTANCE_NOT_MET"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order with
	// trigger number of units that violates the account’s Guaranteed Stop Loss Order
	// level restriction volume.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_VOLUME_EXCEEDED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_VOLUME_EXCEEDED"
	// Filling the Order would result in the creation of a Guaranteed Stop Loss Order with
	// trigger price that violates the account’s Guaranteed Stop Loss Order level restriction
	// price range.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_PRICE_RANGE_EXCEEDED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_LEVEL_RESTRICTION_PRICE_RANGE_EXCEEDED"
	// The Guaranteed Stop Loss on fill distance is invalid
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_DISTANCE_INVALID TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_DISTANCE_INVALID"
	// The Guaranteed Stop Loss on fill price distance exceeds the maximum allowed amount.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED"
	// The Guaranteed Stop Loss on fill distance contains more precision than is allowed
	// by the instrument
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_DISTANCE_PRECISION_EXCEEDED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_DISTANCE_PRECISION_EXCEEDED"
	// The Guaranteed Stop Loss on fill contains both the price and distance fields.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_SPECIFIED TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_SPECIFIED"
	// The Guaranteed Stop Loss on fill contains neither the price nor distance fields.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_MISSING TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_PRICE_AND_DISTANCE_BOTH_MISSING"
	// The Guaranteed Stop Loss on fill specified does not provide a TimeInForce
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_TIME_IN_FORCE_MISSING TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_TIME_IN_FORCE_MISSING"
	// The Guaranteed Stop Loss on fill specifies an invalid TimeInForce
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID"
	// The Guaranteed Stop Loss on fill specifies a GTD TimeInForce but does not provide
	// a GTD timestamp
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_MISSING TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_MISSING"
	// The Guaranteed Stop Loss on fill specifies a GTD timestamp that is in the past.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	// The Guaranteed Stop Loss on fill client Order ID specified is invalid
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_INVALID TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_INVALID"
	// The Guaranteed Stop Loss on fill client Order tag specified is invalid
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_TAG_INVALID TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_TAG_INVALID"
	// The Guaranteed Stop Loss on fill client Order comment specified is invalid.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_COMMENT_INVALID TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_CLIENT_ORDER_COMMENT_INVALID"
	// The Guaranteed Stop Loss on fill specified does not provide a TriggerCondition.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_MISSING TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_MISSING"
	// The Guaranteed Stop Loss on fill specifies an invalid TriggerCondition.
	TransactionRejectReason_GUARANTEED_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID TransactionRejectReason = "GUARANTEED_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID"
	// A Trailing Stop Loss Order for the specified Trade already exists
	TransactionRejectReason_TRAILING_STOP_LOSS_ORDER_ALREADY_EXISTS TransactionRejectReason = "TRAILING_STOP_LOSS_ORDER_ALREADY_EXISTS"
	// The Trailing Stop Loss Order would cause the associated Trade to be in violation
	// of the FIFO violation safeguard constraints
	TransactionRejectReason_TRAILING_STOP_LOSS_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD TransactionRejectReason = "TRAILING_STOP_LOSS_ORDER_WOULD_VIOLATE_FIFO_VIOLATION_SAFEGUARD"
	// The Trailing Stop Loss Order would be in violation of the risk management Order mutual
	// exclusivity configuration specifying that only one risk management order can be attached
	// to a Trade.
	TransactionRejectReason_TRAILING_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION TransactionRejectReason = "TRAILING_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_MUTUALLY_EXCLUSIVE_VIOLATION"
	// The Trailing Stop Loss Order would be in violation of the risk management Order mutual
	// exclusivity configuration specifying that if a GSLO is already attached to a Trade,
	// no other risk management Order can be attached to the same Trade.
	TransactionRejectReason_TRAILING_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION TransactionRejectReason = "TRAILING_STOP_LOSS_ORDER_RMO_MUTUAL_EXCLUSIVITY_GSLO_EXCLUDES_OTHERS_VIOLATION"
	// The Trailing Stop Loss on fill specified does not provide a distance
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MISSING TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MISSING"
	// The Trailing Stop Loss on fill distance is invalid
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_INVALID TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_INVALID"
	// The Trailing Stop Loss on fill distance contains more precision than is allowed by
	// the instrument
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_PRECISION_EXCEEDED TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_PRECISION_EXCEEDED"
	// The Trailing Stop Loss on fill price distance exceeds the maximum allowed amount
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MAXIMUM_EXCEEDED"
	// The Trailing Stop Loss on fill price distance does not meet the minimum allowed amount
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MINIMUM_NOT_MET TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_PRICE_DISTANCE_MINIMUM_NOT_MET"
	// The Trailing Stop Loss on fill specified does not provide a TimeInForce
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_TIME_IN_FORCE_MISSING TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_TIME_IN_FORCE_MISSING"
	// The Trailing Stop Loss on fill specifies an invalid TimeInForce
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_TIME_IN_FORCE_INVALID"
	// The Trailing Stop Loss on fill TimeInForce is specified as GTD but no GTD timestamp
	// is provided
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_MISSING TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_MISSING"
	// The Trailing Stop Loss on fill GTD timestamp is in the past
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_GTD_TIMESTAMP_IN_PAST"
	// The Trailing Stop Loss on fill client Order ID specified is invalid
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_INVALID TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_ID_INVALID"
	// The Trailing Stop Loss on fill client Order tag specified is invalid
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_TAG_INVALID TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_TAG_INVALID"
	// The Trailing Stop Loss on fill client Order comment specified is invalid
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_COMMENT_INVALID TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_CLIENT_ORDER_COMMENT_INVALID"
	// A client attempted to create either a Trailing Stop Loss order or an order with a
	// Trailing Stop Loss On Fill specified, which may not yet be supported.
	TransactionRejectReason_TRAILING_STOP_LOSS_ORDERS_NOT_SUPPORTED TransactionRejectReason = "TRAILING_STOP_LOSS_ORDERS_NOT_SUPPORTED"
	// The Trailing Stop Loss on fill specified does not provide a TriggerCondition
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_MISSING TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_MISSING"
	// The Tailing Stop Loss on fill specifies an invalid TriggerCondition
	TransactionRejectReason_TRAILING_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID TransactionRejectReason = "TRAILING_STOP_LOSS_ON_FILL_TRIGGER_CONDITION_INVALID"
	// The request to close a Trade does not specify a full or partial close
	TransactionRejectReason_CLOSE_TRADE_TYPE_MISSING TransactionRejectReason = "CLOSE_TRADE_TYPE_MISSING"
	// The request to close a Trade partially did not specify the number of units to close
	TransactionRejectReason_CLOSE_TRADE_PARTIAL_UNITS_MISSING TransactionRejectReason = "CLOSE_TRADE_PARTIAL_UNITS_MISSING"
	// The request to partially close a Trade specifies a number of units that exceeds the
	// current size of the given Trade
	TransactionRejectReason_CLOSE_TRADE_UNITS_EXCEED_TRADE_SIZE TransactionRejectReason = "CLOSE_TRADE_UNITS_EXCEED_TRADE_SIZE"
	// The Position requested to be closed out does not exist
	TransactionRejectReason_CLOSEOUT_POSITION_DOESNT_EXIST TransactionRejectReason = "CLOSEOUT_POSITION_DOESNT_EXIST"
	// The request to closeout a Position was specified incompletely
	TransactionRejectReason_CLOSEOUT_POSITION_INCOMPLETE_SPECIFICATION TransactionRejectReason = "CLOSEOUT_POSITION_INCOMPLETE_SPECIFICATION"
	// A partial Position closeout request specifies a number of units that exceeds the
	// current Position
	TransactionRejectReason_CLOSEOUT_POSITION_UNITS_EXCEED_POSITION_SIZE TransactionRejectReason = "CLOSEOUT_POSITION_UNITS_EXCEED_POSITION_SIZE"
	// The request to closeout a Position could not be fully satisfied
	TransactionRejectReason_CLOSEOUT_POSITION_REJECT TransactionRejectReason = "CLOSEOUT_POSITION_REJECT"
	// The request to partially closeout a Position did not specify the number of units
	// to close.
	TransactionRejectReason_CLOSEOUT_POSITION_PARTIAL_UNITS_MISSING TransactionRejectReason = "CLOSEOUT_POSITION_PARTIAL_UNITS_MISSING"
	// The markup group ID provided is invalid
	TransactionRejectReason_MARKUP_GROUP_ID_INVALID TransactionRejectReason = "MARKUP_GROUP_ID_INVALID"
	// The PositionAggregationMode provided is not supported/valid.
	TransactionRejectReason_POSITION_AGGREGATION_MODE_INVALID TransactionRejectReason = "POSITION_AGGREGATION_MODE_INVALID"
	// No configuration parameters provided
	TransactionRejectReason_ADMIN_CONFIGURE_DATA_MISSING TransactionRejectReason = "ADMIN_CONFIGURE_DATA_MISSING"
	// The margin rate provided is invalid
	TransactionRejectReason_MARGIN_RATE_INVALID TransactionRejectReason = "MARGIN_RATE_INVALID"
	// The margin rate provided would cause an immediate margin closeout
	TransactionRejectReason_MARGIN_RATE_WOULD_TRIGGER_CLOSEOUT TransactionRejectReason = "MARGIN_RATE_WOULD_TRIGGER_CLOSEOUT"
	// The account alias string provided is invalid
	TransactionRejectReason_ALIAS_INVALID TransactionRejectReason = "ALIAS_INVALID"
	// No configuration parameters provided
	TransactionRejectReason_CLIENT_CONFIGURE_DATA_MISSING TransactionRejectReason = "CLIENT_CONFIGURE_DATA_MISSING"
	// The margin rate provided would cause the Account to enter a margin call state.
	TransactionRejectReason_MARGIN_RATE_WOULD_TRIGGER_MARGIN_CALL TransactionRejectReason = "MARGIN_RATE_WOULD_TRIGGER_MARGIN_CALL"
	// Funding is not possible because the requested transfer amount is invalid
	TransactionRejectReason_AMOUNT_INVALID TransactionRejectReason = "AMOUNT_INVALID"
	// The Account does not have sufficient balance to complete the funding request
	TransactionRejectReason_INSUFFICIENT_FUNDS TransactionRejectReason = "INSUFFICIENT_FUNDS"
	// Funding amount has not been specified
	TransactionRejectReason_AMOUNT_MISSING TransactionRejectReason = "AMOUNT_MISSING"
	// Funding reason has not been specified
	TransactionRejectReason_FUNDING_REASON_MISSING TransactionRejectReason = "FUNDING_REASON_MISSING"
	// The list of Order Identifiers provided for a One Cancels All Order contains an Order
	// Identifier that refers to a Stop Loss Order. OCA groups cannot contain Stop Loss
	// Orders.
	TransactionRejectReason_OCA_ORDER_IDS_STOP_LOSS_NOT_ALLOWED TransactionRejectReason = "OCA_ORDER_IDS_STOP_LOSS_NOT_ALLOWED"
	// Neither Order nor Trade on Fill client extensions were provided for modification
	TransactionRejectReason_CLIENT_EXTENSIONS_DATA_MISSING TransactionRejectReason = "CLIENT_EXTENSIONS_DATA_MISSING"
	// The Order to be replaced has a different type than the replacing Order.
	TransactionRejectReason_REPLACING_ORDER_INVALID TransactionRejectReason = "REPLACING_ORDER_INVALID"
	// The replacing Order refers to a different Trade than the Order that is being replaced.
	TransactionRejectReason_REPLACING_TRADE_ID_INVALID TransactionRejectReason = "REPLACING_TRADE_ID_INVALID"
	// Canceling the order would cause an immediate margin closeout.
	TransactionRejectReason_ORDER_CANCEL_WOULD_TRIGGER_CLOSEOUT TransactionRejectReason = "ORDER_CANCEL_WOULD_TRIGGER_CLOSEOUT"
)

// A filter that can be used when fetching Transactions
type TransactionFilter string

const (
	// Order-related Transactions. These are the Transactions that create, cancel, fill
	// or trigger Orders
	TransactionFilter_ORDER TransactionFilter = "ORDER"
	// Funding-related Transactions
	TransactionFilter_FUNDING TransactionFilter = "FUNDING"
	// Administrative Transactions
	TransactionFilter_ADMIN TransactionFilter = "ADMIN"
	// Account Create Transaction
	TransactionFilter_CREATE TransactionFilter = "CREATE"
	// Account Close Transaction
	TransactionFilter_CLOSE TransactionFilter = "CLOSE"
	// Account Reopen Transaction
	TransactionFilter_REOPEN TransactionFilter = "REOPEN"
	// Client Configuration Transaction
	TransactionFilter_CLIENT_CONFIGURE TransactionFilter = "CLIENT_CONFIGURE"
	// Client Configuration Reject Transaction
	TransactionFilter_CLIENT_CONFIGURE_REJECT TransactionFilter = "CLIENT_CONFIGURE_REJECT"
	// Transfer Funds Transaction
	TransactionFilter_TRANSFER_FUNDS TransactionFilter = "TRANSFER_FUNDS"
	// Transfer Funds Reject Transaction
	TransactionFilter_TRANSFER_FUNDS_REJECT TransactionFilter = "TRANSFER_FUNDS_REJECT"
	// Market Order Transaction
	TransactionFilter_MARKET_ORDER TransactionFilter = "MARKET_ORDER"
	// Market Order Reject Transaction
	TransactionFilter_MARKET_ORDER_REJECT TransactionFilter = "MARKET_ORDER_REJECT"
	// Limit Order Transaction
	TransactionFilter_LIMIT_ORDER TransactionFilter = "LIMIT_ORDER"
	// Limit Order Reject Transaction
	TransactionFilter_LIMIT_ORDER_REJECT TransactionFilter = "LIMIT_ORDER_REJECT"
	// Stop Order Transaction
	TransactionFilter_STOP_ORDER TransactionFilter = "STOP_ORDER"
	// Stop Order Reject Transaction
	TransactionFilter_STOP_ORDER_REJECT TransactionFilter = "STOP_ORDER_REJECT"
	// Market if Touched Order Transaction
	TransactionFilter_MARKET_IF_TOUCHED_ORDER TransactionFilter = "MARKET_IF_TOUCHED_ORDER"
	// Market if Touched Order Reject Transaction
	TransactionFilter_MARKET_IF_TOUCHED_ORDER_REJECT TransactionFilter = "MARKET_IF_TOUCHED_ORDER_REJECT"
	// Take Profit Order Transaction
	TransactionFilter_TAKE_PROFIT_ORDER TransactionFilter = "TAKE_PROFIT_ORDER"
	// Take Profit Order Reject Transaction
	TransactionFilter_TAKE_PROFIT_ORDER_REJECT TransactionFilter = "TAKE_PROFIT_ORDER_REJECT"
	// Stop Loss Order Transaction
	TransactionFilter_STOP_LOSS_ORDER TransactionFilter = "STOP_LOSS_ORDER"
	// Stop Loss Order Reject Transaction
	TransactionFilter_STOP_LOSS_ORDER_REJECT TransactionFilter = "STOP_LOSS_ORDER_REJECT"
	// Guaranteed Stop Loss Order Transaction
	TransactionFilter_GUARANTEED_STOP_LOSS_ORDER TransactionFilter = "GUARANTEED_STOP_LOSS_ORDER"
	// Guaranteed Stop Loss Order Reject Transaction
	TransactionFilter_GUARANTEED_STOP_LOSS_ORDER_REJECT TransactionFilter = "GUARANTEED_STOP_LOSS_ORDER_REJECT"
	// Trailing Stop Loss Order Transaction
	TransactionFilter_TRAILING_STOP_LOSS_ORDER TransactionFilter = "TRAILING_STOP_LOSS_ORDER"
	// Trailing Stop Loss Order Reject Transaction
	TransactionFilter_TRAILING_STOP_LOSS_ORDER_REJECT TransactionFilter = "TRAILING_STOP_LOSS_ORDER_REJECT"
	// One Cancels All Order Transaction
	TransactionFilter_ONE_CANCELS_ALL_ORDER TransactionFilter = "ONE_CANCELS_ALL_ORDER"
	// One Cancels All Order Reject Transaction
	TransactionFilter_ONE_CANCELS_ALL_ORDER_REJECT TransactionFilter = "ONE_CANCELS_ALL_ORDER_REJECT"
	// One Cancels All Order Trigger Transaction
	TransactionFilter_ONE_CANCELS_ALL_ORDER_TRIGGERED TransactionFilter = "ONE_CANCELS_ALL_ORDER_TRIGGERED"
	// Order Fill Transaction
	TransactionFilter_ORDER_FILL TransactionFilter = "ORDER_FILL"
	// Order Cancel Transaction
	TransactionFilter_ORDER_CANCEL TransactionFilter = "ORDER_CANCEL"
	// Order Cancel Reject Transaction
	TransactionFilter_ORDER_CANCEL_REJECT TransactionFilter = "ORDER_CANCEL_REJECT"
	// Order Client Extensions Modify Transaction
	TransactionFilter_ORDER_CLIENT_EXTENSIONS_MODIFY TransactionFilter = "ORDER_CLIENT_EXTENSIONS_MODIFY"
	// Order Client Extensions Modify Reject Transaction
	TransactionFilter_ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT TransactionFilter = "ORDER_CLIENT_EXTENSIONS_MODIFY_REJECT"
	// Trade Client Extensions Modify Transaction
	TransactionFilter_TRADE_CLIENT_EXTENSIONS_MODIFY TransactionFilter = "TRADE_CLIENT_EXTENSIONS_MODIFY"
	// Trade Client Extensions Modify Reject Transaction
	TransactionFilter_TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT TransactionFilter = "TRADE_CLIENT_EXTENSIONS_MODIFY_REJECT"
	// Margin Call Enter Transaction
	TransactionFilter_MARGIN_CALL_ENTER TransactionFilter = "MARGIN_CALL_ENTER"
	// Margin Call Extend Transaction
	TransactionFilter_MARGIN_CALL_EXTEND TransactionFilter = "MARGIN_CALL_EXTEND"
	// Margin Call Exit Transaction
	TransactionFilter_MARGIN_CALL_EXIT TransactionFilter = "MARGIN_CALL_EXIT"
	// Delayed Trade Closure Transaction
	TransactionFilter_DELAYED_TRADE_CLOSURE TransactionFilter = "DELAYED_TRADE_CLOSURE"
	// Daily Financing Transaction
	TransactionFilter_DAILY_FINANCING TransactionFilter = "DAILY_FINANCING"
	// Reset Resettable PL Transaction
	TransactionFilter_RESET_RESETTABLE_PL TransactionFilter = "RESET_RESETTABLE_PL"
)

type TransactionHeartbeat struct {
	// The string “HEARTBEAT”
	Type string `json:"type"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
	// The date/time when the TransactionHeartbeat was created
	Time DateTime `json:"time"`
}
