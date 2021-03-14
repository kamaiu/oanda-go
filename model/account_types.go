//go:generate easyjson -all $GOFILE
package model

// The financing mode of an Account
type AccountFinancingMode string

const (
	// No financing is paid/charged for open Trades in the Account
	FinancingMode_NO_FINANCING AccountFinancingMode = "NO_FINANCING"
	// Second-by-second financing is paid/charged for open Trades in the Account,
	// both daily and when the the Trade is closed
	FinancingMode_SECOND_BY_SECOND AccountFinancingMode = "SECOND_BY_SECOND"
	// A full day’s worth of financing is paid/charged for open Trades in the Account daily at 5pm New York time
	FinancingMode_DAILY AccountFinancingMode = "DAILY"
)

// The current mutability and hedging settings related to guaranteed Stop Loss orders.
type GuaranteedStopLossOrderParameters struct {
	// The current guaranteed Stop Loss Order mutability setting of the Account when market is open.
	MutabilityMarketOpen GuaranteedStopLossOrderMutability `json:"mutabilityMarketOpen"`
	// The current guaranteed Stop Loss Order mutability setting of the Account when market is halted.
	MutabilityMarketHalted GuaranteedStopLossOrderMutability `json:"mutabilityMarketHalted"`
}

// The overall behaviour of the Account regarding guaranteed Stop Loss Orders.
type GuaranteedStopLossOrderMode string

const (
	// The Account is not permitted to create guaranteed Stop Loss Orders.
	GuaranteedStopLossOrderMode_DISABLED GuaranteedStopLossOrderMode = "DISABLED"
	// The Account is able, but not required to have guaranteed Stop Loss Orders for open Trades.
	GuaranteedStopLossOrderMode_ALLOWED GuaranteedStopLossOrderMode = "ALLOWED"
	// The Account is required to have guaranteed Stop Loss Orders for all open Trades.
	GuaranteedStopLossOrderMode_REQUIRED GuaranteedStopLossOrderMode = "REQUIRED"
)

// For Accounts that support guaranteed Stop Loss Orders, describes the actions
// that can be be performed on guaranteed Stop Loss Orders.
type GuaranteedStopLossOrderMutability string

const (
	// Once a guaranteed Stop Loss Order has been created it cannot be replaced or cancelled.
	GuaranteedStopLossOrderMutability_FIXED GuaranteedStopLossOrderMutability = "FIXED"
	// An existing guaranteed Stop Loss Order can only be replaced, not cancelled.
	GuaranteedStopLossOrderMutability_REPLACEABLE GuaranteedStopLossOrderMutability = "REPLACEABLE"
	// Once a guaranteed Stop Loss Order has been created it can be either replaced or cancelled.
	GuaranteedStopLossOrderMutability_CANCELABLE GuaranteedStopLossOrderMutability = "CANCELABLE"
	// An existing guaranteed Stop Loss Order can only be replaced to widen
	// the gap from the current price, not cancelled.
	GuaranteedStopLossOrderMutability_PRICE_WIDEN_ONLY GuaranteedStopLossOrderMutability = "PRICE_WIDEN_ONLY"
)

// The way that position values for an Account are calculated and aggregated.
type PositionAggregationMode string

const (
	// The Position value or margin for each side (long and short) of the Position are
	// computed independently and added together.
	PositionAggretationMode_ABSOLUTE_SUM PositionAggregationMode = "ABSOLUTE_SUM"
	// The Position value or margin for each side (long and short) of the Position are computed independently.
	// The Position value or margin chosen is the maximal absolute value of the two.
	PositionAggretationMode_MAXIMAL_SIDE PositionAggregationMode = "MAXIMAL_SIDE"
	// The units for each side (long and short) of the Position are netted together and the resulting value
	// (long or short) is used to compute the Position value or margin.
	PositionAggretationMode_NET_SUM PositionAggregationMode = "NET_SUM"
)
