package endpoint

import (
	"fmt"
	. "github.com/kamaiu/oanda-go/model"
	"testing"
)

func TestConnection_Accounts(t *testing.T) {
	c := newPracticeConnection()
	resp, err := c.Accounts()
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("nil accounts response")
	}
	if len(resp.Accounts) == 0 {
		t.Fatal("no accounts in response")
	}

	detailList := make([]*Account, 0, len(resp.Accounts))
	for _, acc := range resp.Accounts {
		details, err := c.Account(acc.ID)
		if err != nil {
			t.Fatal(err)
		}
		if details == nil {
			t.Fatal("nil account details response")
		}
		_, err = details.CreatedTime.Parse()
		if err != nil {
			t.Fatal(err)
		}
		detailList = append(detailList, details)
	}

	fmt.Println(resp)
}

func TestConnection_AccountInstruments(t *testing.T) {
	c := newPracticeConnection()
	resp, err := c.Accounts()
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("nil accounts response")
	}
	if len(resp.Accounts) == 0 {
		t.Fatal("no accounts in response")
	}
	instruments, err := c.AccountInstruments(resp.Accounts[0].ID)
	if err != nil {
		t.Fatal(err)
	}
	if instruments == nil {
		t.Fatal("nil instruments response")
	}
	if len(instruments.Instruments) == 0 {
		t.Fatal("no instruments in response")
	}
	j, err := instruments.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(j))
}

func TestConnection_AccountChanges(t *testing.T) {
	c := newPracticeConnection()
	resp, err := c.Accounts()
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatal("nil accounts response")
	}
	if len(resp.Accounts) == 0 {
		t.Fatal("no accounts in response")
	}
	//account, err := c.Account(resp.Accounts[0].ID)
	//if err != nil {
	//	t.Fatal(err)
	//}
	changes, err := c.AccountChanges(resp.Accounts[0].ID, "10")
	if err != nil {
		t.Fatal(err)
	}
	if changes == nil || changes.Changes == nil {
		t.Fatal("nil changes response")
	}
}
