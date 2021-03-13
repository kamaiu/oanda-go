//go:generate easyjson -all $GOFILE
package model

type AccountsResponse struct {
	Accounts []*AccountProperties `json:"accounts"`
}

type AccountResponse struct {
	Account *Account `json:"account"`
}

type AccountConfigurationRequest struct {
	// Client-defined alias (name) for the Account
	Alias string `json:"alias"`
	// Margin rate
	MarginRate DecimalNumber `json:"marginRate"`
}
type AccountConfigurationResponse struct {
	ClientConfigureTransaction *ClientConfigureTransaction `json:"clientConfigureTransaction"`
	LastTransactionID          TransactionID               `json:"lastTransactionID"`
}

type AccountSummaryResponse struct {
	Account           *AccountSummary `json:"account"`
	LastTransactionID TransactionID   `json:"lastTransactionID"`
}

type AccountInstrumentsResponse struct {
	Instruments       []*Instrument `json:"account"`
	LastTransactionID TransactionID `json:"lastTransactionID"`
}

type AccountChangesResponse struct {
	Changes           []*AccountChanges    `json:"account"`
	LastTransactionID TransactionID        `json:"lastTransactionID"`
	State             *AccountChangesState `json:"state"`
}
