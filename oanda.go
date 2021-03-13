package oanda

import (
	"errors"
	"fmt"
	"github.com/kamaiu/oanda-go/endpoint"
	"github.com/kamaiu/oanda-go/model"
	"sync"
)

type Client struct {
	token        string
	conn         *endpoint.Connection
	accounts     []*Account
	accountsByID map[model.AccountID]*Account

	mu sync.RWMutex
}

func NewClient(token string, live bool) (*Client, error) {
	conn := endpoint.NewConnection(token, live)
	client := &Client{
		token:        token,
		conn:         conn,
		accounts:     nil,
		accountsByID: make(map[model.AccountID]*Account),
	}

	accounts, err := conn.Accounts()
	if err != nil {
		return nil, err
	}
	if accounts == nil {
		return nil, errors.New("no accounts")
	}
	client.accounts = make([]*Account, 0, len(accounts.Accounts))
	for _, props := range accounts.Accounts {
		details, err := conn.Account(props.ID)
		if err != nil {
			return nil, err
		}
		account := newAccount(props, details)
		if details == nil {
			account.err = fmt.Errorf("could not retrieve account: %s", props.ID)
		}
		client.accounts = append(client.accounts, account)
		client.accountsByID[props.ID] = account
	}
	return client, nil
}

type Pricing struct {
}

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
