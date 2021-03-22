//go:generate easyjson -all $GOFILE
package model

type AccountsResponse struct {
	// The list of Accounts the client is authorized to access and their
	// associated properties.
	Accounts []*AccountProperties `json:"accounts"`
}

type AccountResponse struct {
	// The full details of the requested Account.
	Account *Account `json:"account"`
	// The ID of the most recent Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type AccountSummaryResponse struct {
	// The summary of the requested Account.
	Account *AccountSummary `json:"account"`
	// The ID of the most recent Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type AccountConfigurationRequest struct {
	// Client-defined alias (name) for the Account
	Alias string `json:"alias"`
	// Margin rate
	MarginRate DecimalNumber `json:"marginRate"`
}

type AccountConfigurationResponse struct {
	// The transaction that configures the Account.
	ClientConfigureTransaction *ClientConfigureTransaction `json:"clientConfigureTransaction"`
	// The ID of the last Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type AccountConfigurationError struct {
	// The transaction that rejects the configuration of the Account.
	ClientConfigureRejectTransaction *ClientConfigureRejectTransaction `json:"clientConfigureRejectTransaction"`
	// The ID of the most recent Transaction created for the Account
	LastTransactionID TransactionID `json:"lastTransactionID"`
	// The code of the error that has occurred. This field may not be returned
	// for some errors.
	ErrorCode string `json:"errorCode"`
	// The human-readable description of the error that has occurred.
	ErrorMessage string `json:"errorMessage"`
}

type AccountInstrumentsResponse struct {
	// The requested list of instruments.
	Instruments []*Instrument `json:"instruments"`
	// The ID of the most recent Transaction created for the Account.
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type AccountChangesResponse struct {
	// The changes to the Account’s Orders, Trades and Positions since the
	// specified Transaction ID. Only provided if the sinceTransactionID is
	// supplied to the poll request.
	Changes *AccountChanges `json:"changes"`
	// The ID of the last Transaction created for the Account.  This Transaction
	// ID should be used for future poll requests, as the client has already
	// observed all changes up to and including it.
	LastTransactionID TransactionID `json:"lastTransactionID"`
	// The Account’s current price-dependent state.
	State *AccountChangesState `json:"state"`
}
