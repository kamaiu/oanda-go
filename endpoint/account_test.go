package endpoint

import (
	"fmt"
	. "github.com/kamaiu/oanda-go/model"
	"testing"
)

func TestConnection_GetAccounts(t *testing.T) {
	c := NewConnection(apiToken, true)
	resp, err := c.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	detailList := make([]*Account, 0, len(resp.Accounts))
	for _, acc := range resp.Accounts {
		details, err := c.Account(acc.ID)
		if err != nil {
			t.Fatal(err)
		}
		tm, err := details.CreatedTime.Parse()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(tm)
		detailList = append(detailList, details)
	}

	fmt.Println(resp)
}
