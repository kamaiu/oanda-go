//go:generate easyjson -all $GOFILE
package model

// The type of the Order.
type OrderType string

const (
	OrderType_MARKET               OrderType = "MARKET"               // Market Order
	OrderType_LIMIT                OrderType = "LIMIT"                // Limit Order
	OrderType_STOP                 OrderType = "STOP"                 // Stop Order
	OrderType_MARKET_IF_TOUCHED    OrderType = "MARKET_IF_TOUCHED"    // Market-if-touched Order
	OrderType_TAKE_PROFIT          OrderType = "TAKE_PROFIT"          // Take Profit Order
	OrderType_STOP_LOSS            OrderType = "STOP_LOSS"            // Stop Loss Order
	OrderType_GUARANTEED_STOP_LOSS OrderType = "GUARANTEED_STOP_LOSS" // Guaranteed Stop Loss Order
	OrderType_TRAILING_STOP_LOSS   OrderType = "TRAILING_STOP_LOSS"   // Trailing Stop Loss Order
	OrderType_FIXED_PRICE          OrderType = "FIXED_PRICE"          // Fixed Price Order
)

// The type of the Order.
type CancellableOrderType string

const (
	CancellableOrderType_LIMIT                OrderType = "LIMIT"                // Limit Order
	CancellableOrderType_STOP                 OrderType = "STOP"                 // Stop Order
	CancellableOrderType_MARKET_IF_TOUCHED    OrderType = "MARKET_IF_TOUCHED"    // Market-if-touched Order
	CancellableOrderType_TAKE_PROFIT          OrderType = "TAKE_PROFIT"          // Take Profit Order
	CancellableOrderType_STOP_LOSS            OrderType = "STOP_LOSS"            // Stop Loss Order
	CancellableOrderType_GUARANTEED_STOP_LOSS OrderType = "GUARANTEED_STOP_LOSS" // Guaranteed Stop Loss Order
	CancellableOrderType_TRAILING_STOP_LOSS   OrderType = "TRAILING_STOP_LOSS"   // Trailing Stop Loss Order
)

// The current state of the Order.
type OrderState string

const (
	OrderState_PENDING   OrderState = "PENDING"   // The Order is currently pending execution
	OrderState_FILLED    OrderState = "FILLED"    // The Order has been filled
	OrderState_TRIGGERED OrderState = "TRIGGERED" // The Order has been triggered
	OrderState_CANCELLED OrderState = "CANCELLED" // The Order has been cancelled
)

// The state to filter the requested Orders by.
type OrderStateFilter string

const (
	OrderStateFilter_PENDING   OrderStateFilter = "PENDING"   // The Orders that are currently pending execution
	OrderStateFilter_FILLED    OrderStateFilter = "FILLED"    // The Orders that have been filled
	OrderStateFilter_TRIGGERED OrderStateFilter = "TRIGGERED" // The Orders that have been triggered
	OrderStateFilter_CANCELLED OrderStateFilter = "CANCELLED" // The Orders that have been cancelled
	OrderStateFilter_ALL       OrderStateFilter = "ALL"       // The Orders that are in any of the possible states
)

// An OrderIdentifier is used to refer to an Order, and contains both the OrderID and the ClientOrderID.
type OrderIdentifier struct {
	// The OANDA-assigned Order ID
	OrderID OrderID `json:"orderID"`
	// The client-provided client Order ID
	ClientOrderID ClientID `json:"clientOrderID"`
}

// The specification of an Order as referred to by clients
// Either the Order’s OANDA-assigned OrderID or the Order’s client-provided ClientID prefixed by the “@” symbol
type OrderSpecifier string

func (o OrderSpecifier) IsClient() bool {
	return len(o) > 0 && o[0] == '@'
}

func (o OrderSpecifier) OrderID() OrderID {
	if len(o) == 0 || o[0] == '@' {
		return ""
	}
	return OrderID(o)
}

func (o OrderSpecifier) ClientID() ClientID {
	if len(o) == 0 || o[0] != '@' {
		return ""
	}
	return ClientID(o[1:])
}

// The time-in-force of an Order. TimeInForce describes how long an Order should remain pending
// before being automatically cancelled by the execution system.
type TimeInForce string

const (
	// The Order is “Good unTil Cancelled”
	TimeInForce_GTC TimeInForce = "GTC"
	// The Order is “Good unTil Date” and will be cancelled at the provided time
	TimeInForce_GTD TimeInForce = "GTD"
	// The Order is “Good For Day” and will be cancelled at 5pm New York time
	TimeInForce_GFD TimeInForce = "GFD"
	// The Order must be immediately “Filled Or Killed”
	TimeInForce_FOK TimeInForce = "FOK"
	// The Order must be “Immediately partially filled Or Cancelled”
	TimeInForce_IOC TimeInForce = "IOC"
)

// Specification of how Positions in the Account are modified when the Order is filled.
type OrderPositionFill string

const (
	// When the Order is filled, only allow Positions to be opened or extended.
	OrderPositionFill_OPEN_ONLY OrderPositionFill = "OPEN_ONLY"
	// When the Order is filled, always fully reduce an existing Position before opening a new Position.
	OrderPositionFill_REDUCE_FIRST OrderPositionFill = "REDUCE_FIRST"
	// When the Order is filled, only reduce an existing Position.
	OrderPositionFill_REDUCE_ONLY OrderPositionFill = "REDUCE_ONLY"
	// When the Order is filled, use REDUCE_FIRST behaviour for non-client hedging Accounts,
	// and OPEN_ONLY behaviour for client hedging Accounts.
	OrderPositionFill_DEFAULT OrderPositionFill = "DEFAULT"
)

// Specification of which price component should be used when determining if an Order should
// be triggered and filled. This allows Orders to be triggered based on the bid, ask, mid,
// default (ask for buy, bid for sell) or inverse (ask for sell, bid for buy) price depending
// on the desired behaviour. Orders are always filled using their default price component.
// This feature is only provided through the REST API. Clients who choose to specify a non-default
// trigger condition will not see it reflected in any of OANDA’s proprietary or partner trading
// platforms, their transaction history or their account statements. OANDA platforms always assume
// that an Order’s trigger condition is set to the default value when indicating the distance from
// an Order’s trigger price, and will always provide the default trigger condition when creating or
// modifying an Order. A special restriction applies when creating a Guaranteed Stop Loss Order.
// In this case the TriggerCondition value must either be “DEFAULT”, or the “natural” trigger side
// “DEFAULT” results in. So for a Guaranteed Stop Loss Order for a long trade valid values are
// “DEFAULT” and “BID”, and for short trades “DEFAULT” and “ASK” are valid.
type OrderTriggerCondition string

const (
	// Trigger an Order the “natural” way: compare its price to the ask for long Orders and bid for short Orders.
	OrderTriggerCondition_DEFAULT OrderTriggerCondition = "DEFAULT"
	// Trigger an Order the opposite of the “natural” way: compare its price the bid for long Orders and ask
	// for short Orders.
	OrderTriggerCondition_INVERSE OrderTriggerCondition = "INVERSE"
	// Trigger an Order by comparing its price to the bid regardless of whether it is long or short.
	OrderTriggerCondition_BID OrderTriggerCondition = "BID"
	// Trigger an Order by comparing its price to the ask regardless of whether it is long or short.
	OrderTriggerCondition_ASK OrderTriggerCondition = "ASK"
	// Trigger an Order by comparing its price to the midpoint regardless of whether it is long or short.
	OrderTriggerCondition_MID OrderTriggerCondition = "MID"
)

// The dynamic state of an Order. This is only relevant to TrailingStopLoss Orders,
// as no other Order type has dynamic state.
type DynamicOrderState struct {
	// The Order’s ID
	ID OrderID `json:"id"`
	// The Order’s calculated trailing stop value.
	TrailingStopValue PriceValue `json:"trailingStopValue"`
	// The distance between the Trailing Stop Loss Order’s trailingStopValue and
	// the current Market Price. This represents the distance (in price units)
	// of the Order from a triggering price. If the distance could not be
	// determined, this value will not be set.
	TriggerDistance PriceValue `json:"triggerDistance"`
	// True if an exact trigger distance could be calculated. If false, it means
	// the provided trigger distance is a best estimate. If the distance could
	// not be determined, this value will not be set.
	IsTriggerDistanceExact bool `json:"isTriggerDistanceExact"`
}

// Representation of many units of an Instrument are available to be traded for both long and short Orders.
type UnitsAvailableDetails struct {
	// The units available for long Orders.
	Long DecimalNumber `json:"long"`
	// The units available for short Orders.
	Short DecimalNumber `json:"short"`
}

// Representation of how many units of an Instrument are available to be traded by an
// Order depending on its positionFill option.
type UnitsAvailable struct {
	// The number of units that are available to be traded using an Order with a
	// positionFill option of “DEFAULT”. For an Account with hedging enabled,
	// this value will be the same as the “OPEN_ONLY” value. For an Account
	// without hedging enabled, this value will be the same as the
	// “REDUCE_FIRST” value.
	Default *UnitsAvailableDetails `json:"default"`
	// The number of units that may are available to be traded with an Order
	// with a positionFill option of “REDUCE_FIRST”.
	ReduceFirst *UnitsAvailableDetails `json:"reduceFirst"`
	// The number of units that may are available to be traded with an Order
	// with a positionFill option of “REDUCE_ONLY”.
	ReduceOnly *UnitsAvailableDetails `json:"reduceOnly"`
	// The number of units that may are available to be traded with an Order
	// with a positionFill option of “OPEN_ONLY”.
	OpenOnly *UnitsAvailableDetails `json:"openOnly"`
}

// Details required by clients creating a Guaranteed Stop Loss Order
type GuaranteedStopLossOrderEntryData struct {
	// The minimum distance allowed between the Trade’s fill price and the
	// configured price for guaranteed Stop Loss Orders created for this
	// instrument. Specified in price units.
	MinimumDistance DecimalNumber `json:"minimumDistance"`
	// The amount that is charged to the account if a guaranteed Stop Loss Order
	// is triggered and filled. The value is in price units and is charged for
	// each unit of the Trade.
	Premium DecimalNumber `json:"premium"`
	// The guaranteed Stop Loss Order level restriction for this instrument.
	LevelRestriction GuaranteedStopLossOrderLevelRestriction `json:"levelRestriction"`
}
