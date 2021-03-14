//go:generate easyjson -all $GOFILE
package model

import (
	"errors"
	"strings"
)

// The string representation of an Account Identifier.
// “-“-delimited string with format “{siteID}-{divisionID}-{userID}-{accountNumber}”
// Example: 001-011-5838423-001
type AccountID string

func (s AccountID) Parse() (siteID string, divisionID string, userID string, accountNumber string, err error) {
	p := strings.Split(string(s), "-")
	if len(p) != 4 {
		err = errors.New("invalid format")
		return
	}
	return p[0], p[1], p[2], p[3], nil
}

// The full details of a client’s Account. This includes full open Trade, open Position
// and pending Order representation.
type Account struct {
	// The Account’s identifier
	Id AccountID `json:"id"`
	// Client-assigned alias for the Account. Only provided if the Account has an alias set
	Alias string `json:"alias"`
	// The home currency of the Account
	Currency Currency `json:"currency"`
	// AccountID of the user that created the Account.
	CreatedByUserID int64 `json:"createdByUserID"`
	// The date/time when the Account was created.
	CreatedTime DateTime `json:"createdTime"`
	// The current guaranteed Stop Loss Order settings of the Account. This
	// field will only be present if the guaranteedStopLossOrderMode is not ‘DISABLED’.
	GuaranteedStopLossOrderParameters *GuaranteedStopLossOrderParameters `json:"guaranteedStopLossOrderParameters"`
	// The current guaranteed Stop Loss Order mode of the Account.
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderMode `json:"guaranteedStopLossOrderMode"`
	// The current guaranteed Stop Loss Order mutability setting of the Account.
	// This field will only be present if the guaranteedStopLossOrderMode is not ‘DISABLED’.
	// DEPREACTED
	GuaranteedStopLossOrderMutability GuaranteedStopLossOrderMutability `json:"guaranteedStopLossOrderMutability"`
	// The date/time that the Account’s resettablePL was last reset.
	ResettablePLTime DateTime `json:"resettablePLTime"`
	// Client-provided margin rate override for the Account. The effective
	// margin rate of the Account is the lesser of this value and the OANDA
	// margin rate for the Account’s division. This value is only provided if a
	// margin rate override exists for the Account.
	MarginRate DecimalNumber `json:"marginRate"`
	// The number of Trades currently open in the Account.
	OpenTradeCount int64 `json:"openTradeCount"`
	// The number of Positions currently open in the Account.
	OpenPositionCount int64 `json:"openPositionCount"`
	// The number of Orders currently pending in the Account.
	PendingOrderCount int64 `json:"pendingOrderCount"`
	// Flag indicating that the Account has hedging enabled.
	HedgingEnabled bool `json:"hedgingEnabled"`
	// The total unrealized profit/loss for all Trades currently open in the Account.
	UnrealizedPL AccountUnits `json:"unrealizedPL"`
	// The net asset value of the Account. Equal to Account balance + unrealizedPL.
	NAV AccountUnits `json:"NAV"`
	// Margin currently used for the Account.
	MarginUsed AccountUnits `json:"marginUsed"`
	// Margin available for Account currency.
	MarginAvailable AccountUnits `json:"marginAvailable"`
	// The value of the Account’s open positions represented in the Account’s home currency.
	PositionValue AccountUnits `json:"positionValue"`
	// The Account’s margin closeout unrealized PL.
	MarginCloseoutUnrealizedPL AccountUnits `json:"marginCloseoutUnrealizedPL"`
	// The Account’s margin closeout NAV.
	MarginCloseoutNAV AccountUnits `json:"marginCloseoutNAV"`
	// The Account’s margin closeout margin used.
	MarginCloseoutMarginUsed AccountUnits `json:"marginCloseoutMarginUsed"`
	// The Account’s margin closeout percentage. When this value is 1.0 or above
	// the Account is in a margin closeout situation.
	MarginCloseoutPercent DecimalNumber `json:"marginCloseoutPercent"`
	// The value of the Account’s open positions as used for margin closeout
	// calculations represented in the Account’s home currency.
	MarginCloseoutPositionValue DecimalNumber `json:"marginCloseoutPositionValue"`
	// The current WithdrawalLimit for the account which will be zero or a
	// positive value indicating how much can be withdrawn from the account.
	WithdrawalLimit AccountUnits `json:"withdrawalLimit"`
	// The Account’s margin call margin used.
	MarginCallMarginUsed AccountUnits `json:"marginCallMarginUsed"`
	// The Account’s margin call percentage. When this value is 1.0 or above the
	// Account is in a margin call situation.
	MarginCallPercent DecimalNumber `json:"marginCallPercent"`
	// The current balance of the account.
	Balance AccountUnits `json:"balance"`
	// The total profit/loss realized over the lifetime of the Account.
	PL AccountUnits `json:"pl"`
	// The total realized profit/loss for the account since it was last reset by the client.
	ResettablePL AccountUnits `json:"resettablePL"`
	// The total amount of financing paid/collected over the lifetime of the account.
	Financing AccountUnits `json:"financing"`
	// The total amount of commission paid over the lifetime of the Account.
	Commission AccountUnits `json:"commission"`
	// The total amount of dividend adjustment paid over the lifetime of the Account in the Account’s home currency.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`
	// The total amount of fees charged over the lifetime of the Account for the
	// execution of guaranteed Stop Loss Orders.
	GuaranteedExecutionFees AccountUnits `json:"guaranteedExecutionFees"`
	// The date/time when the Account entered a margin call state. Only provided
	// if the Account is in a margin call.
	MarginCallEnterTime DateTime `json:"marginCallEnterTime"`
	// The number of times that the Account’s current margin call was extended.
	MarginCallExtensionCount int64 `json:"marginCallExtensionCount"`
	// The date/time of the Account’s last margin call extension.
	LastMarginCallExtensionTime DateTime `json:"lastMarginCallExtensionTime"`
	// The AccountID of the last Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`
	// The details of the Trades currently open in the Account.
	Trades []*TradeSummary `json:"trades"`
	// The details all Account Positions.
	Positions []*Position `json:"positions"`
	// The details of the Orders currently pending in the Account.
	Orders []*Order `json:"orders"`
}

type AccountChangesState struct {
	// The total unrealized profit/loss for all Trades currently open in the Account.
	UnrealizedPL AccountUnits `json:"unrealizedPL"`
	// The net asset value of the Account. Equal to Account balance + unrealizedPL.
	NAV AccountUnits `json:"NAV"`
	// Margin currently used for the Account.
	MarginUsed AccountUnits `json:"marginUsed"`
	// Margin available for Account currency.
	MarginAvailable AccountUnits `json:"marginAvailable"`
	// The value of the Account’s open positions represented in the Account’s home currency.
	PositionValue AccountUnits `json:"positionValue"`
	// The Account’s margin closeout unrealized PL.
	MarginCloseoutUnrealizedPL AccountUnits `json:"marginCloseoutUnrealizedPL"`
	// The Account’s margin closeout NAV.
	MarginCloseoutNAV AccountUnits `json:"marginCloseoutNAV"`
	// The Account’s margin closeout margin used.
	MarginCloseoutMarginUsed AccountUnits `json:"marginCloseoutMarginUsed"`
	// The Account’s margin closeout percentage. When this value is 1.0 or above
	// the Account is in a margin closeout situation.
	MarginCloseoutPercent DecimalNumber `json:"marginCloseoutPercent"`
	// The value of the Account’s open positions as used for margin closeout
	// calculations represented in the Account’s home currency.
	MarginCloseoutPositionValue DecimalNumber `json:"marginCloseoutPositionValue"`
	// The current WithdrawalLimit for the account which will be zero or a
	// positive value indicating how much can be withdrawn from the account.
	WithdrawalLimit AccountUnits `json:"withdrawalLimit"`
	// The Account’s margin call margin used.
	MarginCallMarginUsed AccountUnits `json:"marginCallMarginUsed"`
	// The Account’s margin call percentage. When this value is 1.0 or above the
	// Account is in a margin call situation.
	MarginCallPercent DecimalNumber `json:"marginCallPercent"`
	// The current balance of the account.
	Balance AccountUnits `json:"balance"`
	// The total profit/loss realized over the lifetime of the Account.
	PL AccountUnits `json:"pl"`
	// The total realized profit/loss for the account since it was last reset by the client.
	ResettablePL AccountUnits `json:"resettablePL"`
	// The total amount of financing paid/collected over the lifetime of the account.
	Financing AccountUnits `json:"financing"`
	// The total amount of commission paid over the lifetime of the Account.
	Commission AccountUnits `json:"commission"`
	// The total amount of dividend adjustment paid over the lifetime of the Account in the Account’s home currency.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`
	// The total amount of fees charged over the lifetime of the Account for the
	// execution of guaranteed Stop Loss Orders.
	GuaranteedExecutionFees AccountUnits `json:"guaranteedExecutionFees"`
	// The date/time when the Account entered a margin call state. Only provided
	// if the Account is in a margin call.
	MarginCallEnterTime DateTime `json:"marginCallEnterTime"`
	// The number of times that the Account’s current margin call was extended.
	MarginCallExtensionCount int64 `json:"marginCallExtensionCount"`
	// The date/time of the Account’s last margin call extension.
	LastMarginCallExtensionTime DateTime `json:"lastMarginCallExtensionTime"`
	// The price-dependent state of each pending Order in the Account.
	Orders []*DynamicOrderState `json:"orders"`
	// The price-dependent state for each open Trade in the Account.
	Trades []*CalculatedTradeState `json:"trades"`
	// The price-dependent state for each open Position in the Account.
	Positions []*CalculatedPositionState `json:"positions"`
}

// Properties related to an Account.
type AccountProperties struct {
	// The Account’s identifier
	ID AccountID `json:"id"`
	// The Account’s associated MT4 Account ID. This field will not be present
	// if the Account is not an MT4 account.
	MT4AccountID int64 `json:"mt4AccountID"`
	// The Account’s tags
	Tags []string `json:"tags"`
}

// A summary representation of a client’s Account. The AccountSummary does not provide to full
// specification of pending Orders, open Trades and Positions.
type AccountSummary struct {
	// The Account’s identifier
	Id AccountID `json:"id"`
	// Client-assigned alias for the Account. Only provided if the Account has an alias set
	Alias string `json:"alias"`
	// The home currency of the Account
	Currency Currency `json:"currency"`
	// AccountID of the user that created the Account.
	CreatedByUserID int64 `json:"createdByUserID"`
	// The date/time when the Account was created.
	CreatedTime DateTime `json:"createdTime"`
	// The current guaranteed Stop Loss Order settings of the Account. This
	// field will only be present if the guaranteedStopLossOrderMode is not ‘DISABLED’.
	GuaranteedStopLossOrderParameters *GuaranteedStopLossOrderParameters `json:"guaranteedStopLossOrderParameters"`
	// The current guaranteed Stop Loss Order mode of the Account.
	GuaranteedStopLossOrderMode GuaranteedStopLossOrderMode `json:"guaranteedStopLossOrderMode"`
	// The current guaranteed Stop Loss Order mutability setting of the Account.
	// This field will only be present if the guaranteedStopLossOrderMode is not ‘DISABLED’.
	// DEPREACTED
	GuaranteedStopLossOrderMutability GuaranteedStopLossOrderMutability `json:"guaranteedStopLossOrderMutability"`
	// The date/time that the Account’s resettablePL was last reset.
	ResettablePLTime DateTime `json:"resettablePLTime"`
	// Client-provided margin rate override for the Account. The effective
	// margin rate of the Account is the lesser of this value and the OANDA
	// margin rate for the Account’s division. This value is only provided if a
	// margin rate override exists for the Account.
	MarginRate DecimalNumber `json:"marginRate"`
	// The number of Trades currently open in the Account.
	OpenTradeCount int64 `json:"openTradeCount"`
	// The number of Positions currently open in the Account.
	OpenPositionCount int64 `json:"openPositionCount"`
	// The number of Orders currently pending in the Account.
	PendingOrderCount int64 `json:"pendingOrderCount"`
	// Flag indicating that the Account has hedging enabled.
	HedgingEnabled bool `json:"hedgingEnabled"`
	// The total unrealized profit/loss for all Trades currently open in the Account.
	UnrealizedPL AccountUnits `json:"unrealizedPL"`
	// The net asset value of the Account. Equal to Account balance + unrealizedPL.
	NAV AccountUnits `json:"NAV"`
	// Margin currently used for the Account.
	MarginUsed AccountUnits `json:"marginUsed"`
	// Margin available for Account currency.
	MarginAvailable AccountUnits `json:"marginAvailable"`
	// The value of the Account’s open positions represented in the Account’s home currency.
	PositionValue AccountUnits `json:"positionValue"`
	// The Account’s margin closeout unrealized PL.
	MarginCloseoutUnrealizedPL AccountUnits `json:"marginCloseoutUnrealizedPL"`
	// The Account’s margin closeout NAV.
	MarginCloseoutNAV AccountUnits `json:"marginCloseoutNAV"`
	// The Account’s margin closeout margin used.
	MarginCloseoutMarginUsed AccountUnits `json:"marginCloseoutMarginUsed"`
	// The Account’s margin closeout percentage. When this value is 1.0 or above
	// the Account is in a margin closeout situation.
	MarginCloseoutPercent DecimalNumber `json:"marginCloseoutPercent"`
	// The value of the Account’s open positions as used for margin closeout
	// calculations represented in the Account’s home currency.
	MarginCloseoutPositionValue DecimalNumber `json:"marginCloseoutPositionValue"`
	// The current WithdrawalLimit for the account which will be zero or a
	// positive value indicating how much can be withdrawn from the account.
	WithdrawalLimit AccountUnits `json:"withdrawalLimit"`
	// The Account’s margin call margin used.
	MarginCallMarginUsed AccountUnits `json:"marginCallMarginUsed"`
	// The Account’s margin call percentage. When this value is 1.0 or above the
	// Account is in a margin call situation.
	MarginCallPercent DecimalNumber `json:"marginCallPercent"`
	// The current balance of the account.
	Balance AccountUnits `json:"balance"`
	// The total profit/loss realized over the lifetime of the Account.
	PL AccountUnits `json:"pl"`
	// The total realized profit/loss for the account since it was last reset by the client.
	ResettablePL AccountUnits `json:"resettablePL"`
	// The total amount of financing paid/collected over the lifetime of the account.
	Financing AccountUnits `json:"financing"`
	// The total amount of commission paid over the lifetime of the Account.
	Commission AccountUnits `json:"commission"`
	// The total amount of dividend adjustment paid over the lifetime of the Account in the Account’s home currency.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`
	// The total amount of fees charged over the lifetime of the Account for the
	// execution of guaranteed Stop Loss Orders.
	GuaranteedExecutionFees AccountUnits `json:"guaranteedExecutionFees"`
	// The date/time when the Account entered a margin call state. Only provided
	// if the Account is in a margin call.
	MarginCallEnterTime DateTime `json:"marginCallEnterTime"`
	// The number of times that the Account’s current margin call was extended.
	MarginCallExtensionCount int64 `json:"marginCallExtensionCount"`
	// The date/time of the Account’s last margin call extension.
	LastMarginCallExtensionTime DateTime `json:"lastMarginCallExtensionTime"`
	// The AccountID of the last Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

// The mutable state of a client’s Account.
type AccumulatedAccountState struct {
	// The current balance of the account.
	Balance AccountUnits `json:"balance"`
	// The total profit/loss realized over the lifetime of the Account.
	PL AccountUnits `json:"pl"`
	// The total realized profit/loss for the account since it was last reset by the client.
	ResettablePL AccountUnits `json:"resettablePL"`
	// The total amount of financing paid/collected over the lifetime of the account.
	Financing AccountUnits `json:"financing"`
	// The total amount of commission paid over the lifetime of the Account.
	Commission AccountUnits `json:"commission"`
	// The total amount of dividend adjustment paid over the lifetime of the Account in the Account’s home currency.
	DividendAdjustment AccountUnits `json:"dividendAdjustment"`
	// The total amount of fees charged over the lifetime of the Account for the
	// execution of guaranteed Stop Loss Orders.
	GuaranteedExecutionFees AccountUnits `json:"guaranteedExecutionFees"`
	// The date/time when the Account entered a margin call state. Only provided
	// if the Account is in a margin call.
	MarginCallEnterTime DateTime `json:"marginCallEnterTime"`
	// The number of times that the Account’s current margin call was extended.
	MarginCallExtensionCount int64 `json:"marginCallExtensionCount"`
	// The date/time of the Account’s last margin call extension.
	LastMarginCallExtensionTime DateTime `json:"lastMarginCallExtensionTime"`
}

// The dynamically calculated state of a client’s Account.
type CalculatedAccountState struct {
	// The total unrealized profit/loss for all Trades currently open in the Account.
	UnrealizedPL AccountUnits `json:"unrealizedPL"`
	// The net asset value of the Account. Equal to Account balance + unrealizedPL.
	NAV AccountUnits `json:"NAV"`
	// Margin currently used for the Account.
	MarginUsed AccountUnits `json:"marginUsed"`
	// Margin available for Account currency.
	MarginAvailable AccountUnits `json:"marginAvailable"`
	// The value of the Account’s open positions represented in the Account’s home currency.
	PositionValue AccountUnits `json:"positionValue"`
	// The Account’s margin closeout unrealized PL.
	MarginCloseoutUnrealizedPL AccountUnits `json:"marginCloseoutUnrealizedPL"`
	// The Account’s margin closeout NAV.
	MarginCloseoutNAV AccountUnits `json:"marginCloseoutNAV"`
	// The Account’s margin closeout margin used.
	MarginCloseoutMarginUsed AccountUnits `json:"marginCloseoutMarginUsed"`
	// The Account’s margin closeout percentage. When this value is 1.0 or above
	// the Account is in a margin closeout situation.
	MarginCloseoutPercent DecimalNumber `json:"marginCloseoutPercent"`
	// The value of the Account’s open positions as used for margin closeout
	// calculations represented in the Account’s home currency.
	MarginCloseoutPositionValue DecimalNumber `json:"marginCloseoutPositionValue"`
	// The current WithdrawalLimit for the account which will be zero or a
	// positive value indicating how much can be withdrawn from the account.
	WithdrawalLimit AccountUnits `json:"withdrawalLimit"`
	// The Account’s margin call margin used.
	MarginCallMarginUsed AccountUnits `json:"marginCallMarginUsed"`
	// The Account’s margin call percentage. When this value is 1.0 or above the
	// Account is in a margin call situation.
	MarginCallPercent DecimalNumber `json:"marginCallPercent"`
}

// An AccountChanges Object is used to represent the changes to an Account’s Orders, Trades and
// Positions since a specified Account TransactionID in the past.
type AccountChanges struct {
	// The Orders created. These Orders may have been filled, cancelled or triggered in the same period.
	OrdersCreated []*Order `json:"ordersCreated"`
	// The Orders cancelled.
	OrdersCancelled []*Order `json:"ordersCancelled"`
	// The Orders filled.
	OrdersFilled []*Order `json:"ordersFilled"`
	// The Orders triggered.
	OrdersTriggered []*Order `json:"ordersTriggered"`
	// The Trades opened.
	TradesOpened []*TradeSummary `json:"tradesOpened"`
	// The Trades reduced.
	TradesReduced []*TradeSummary `json:"tradesReduced"`
	// The Trades closed.
	TradesClosed []*TradeSummary `json:"tradesClosed"`
	// The Positions changed.
	Positions []*Position `json:"positions"`
	// The Transactions that have been generated.
	Transactions []*Transaction `json:"transactions"`
}

// Contains the attributes of a user.
type UserAttributes struct {
	// The user’s OANDA-assigned user AccountID.
	UserID int64 `json:"userID"`
	// The user-provided username.
	Username string `json:"username"`
	// The user’s title.
	Title string `json:"title"`
	// The user’s name.
	Name string `json:"name"`
	// The user’s email address.
	Email string `json:"email"`
	// The OANDA division the user belongs to.
	DivisionAbbreviation string `json:"divisionAbbreviation"`
	// The user’s preferred language.
	LanguageAbbreviation string `json:"languageAbbreviation"`
	// The home currency of the Account.
	HomeCurrency Currency `json:"homeCurrency"`
}
