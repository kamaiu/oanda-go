package oanda

import (
	"github.com/kamaiu/oanda-go/model"
	"sync"
)

type Account struct {
	props   *model.AccountProperties
	details *model.Account
	err     error
	mu      sync.RWMutex
}

func newAccount(props *model.AccountProperties, details *model.Account) *Account {
	return &Account{
		props:   props,
		details: details,
	}
}
