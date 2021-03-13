//go:generate easyjson -all $GOFILE
package model

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
