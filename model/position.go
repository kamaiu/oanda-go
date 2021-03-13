//go:generate easyjson -all $GOFILE
package model

// The specification of a Position within an Account.
type Position struct {
	// The Position’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// Profit/loss realized by the Position over the lifetime of the Account.
	Pl AccountUnits `json:"pl"`
	// The unrealized profit/loss of all open Trades that contribute to this Position.
	UnrealizedPL AccountUnits `json:"unrealizedPL"`
	// Margin currently used by the Position.
	MarginUsed AccountUnits `json:"marginUsed"`
	// Profit/loss realized by the Position since the Account’s resettablePL was last
	// reset by the client.
	ResettablePL AccountUnits `json:"resettablePL"`
	// The total amount of financing paid/collected for this instrument over the lifetime
	// of the Account.
	Financing AccountUnits `json:"financing"`
	// The total amount of commission paid for this instrument over the lifetime of the
	// Account.
	Commission AccountUnits `json:"commission"`
	// The total amount of dividend adjustment paid for this instrument over the lifetime
	// of the Account.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`
	// The total amount of fees charged over the lifetime of the Account for the execution
	// of guaranteed Stop Loss Orders for this instrument.
	GuaranteedExecutionFees AccountUnits `json:"guaranteedExecutionFees"`
	// The details of the long side of the Position.
	Long *PositionSide `json:"long"`
	// The details of the short side of the Position.
	Short *PositionSide `json:"short"`
}

// The representation of a Position for a single direction (long or short).
type PositionSide struct {
	// Number of units in the position (negative value indicates short position, positive
	// indicates long position).
	Units DecimalNumber `json:"units"`
	// Volume-weighted average of the underlying Trade open prices for the Position.
	AveragePrice PriceValue `json:"averagePrice"`
	// List of the open Trade IDs which contribute to the open Position.
	TradeIDs []TradeID `json:"tradeIDs"`
	// Profit/loss realized by the PositionSide over the lifetime of the Account.
	Pl AccountUnits `json:"pl"`
	// The unrealized profit/loss of all open Trades that contribute to this PositionSide.
	UnrealizedPL AccountUnits `json:"unrealizedPL"`
	// Profit/loss realized by the PositionSide since the Account’s resettablePL was
	// last reset by the client.
	ResettablePL AccountUnits `json:"resettablePL"`
	// The total amount of financing paid/collected for this PositionSide over the lifetime
	// of the Account.
	Financing AccountUnits `json:"financing"`
	// The total amount of dividend adjustment paid for the PositionSide over the lifetime
	// of the Account.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`
	// The total amount of fees charged over the lifetime of the Account for the execution
	// of guaranteed Stop Loss Orders attached to Trades for this PositionSide.
	GuaranteedExecutionFees AccountUnits `json:"guaranteedExecutionFees"`
}

// The dynamic (calculated) state of a Position
type CalculatedPositionState struct {
	// The Position’s Instrument.
	Instrument InstrumentName `json:"instrument"`
	// The Position’s net unrealized profit/loss
	NetUnrealizedPL AccountUnits `json:"netUnrealizedPL"`
	// The unrealized profit/loss of the Position’s long open Trades
	LongUnrealizedPL AccountUnits `json:"longUnrealizedPL"`
	// The unrealized profit/loss of the Position’s short open Trades
	ShortUnrealizedPL AccountUnits `json:"shortUnrealizedPL"`
	// Margin currently used by the Position.
	MarginUsed AccountUnits `json:"marginUsed"`
}
