package endpoint

import (
	"fmt"
	. "github.com/kamaiu/oanda-go/model"
	"testing"
)

func TestConnection_CandlesLatest(t *testing.T) {
	c := NewConnection(apiToken, true)
	accounts, err := c.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CandlesLatest(accounts.Accounts[0].ID, NewCandlesLatestRequest(
		NewCandleSpecification("EUR_USD", CandlestickGranularity_M10, PricingComponent_BID_ASK_MID),
	))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(resp)
}
