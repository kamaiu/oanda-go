package endpoint

import (
	"fmt"
	. "github.com/kamaiu/oanda-go/model"
	"testing"
	"time"
)

func TestPricingStream(t *testing.T) {
	c := NewConnection(apiToken, true)
	accounts, err := c.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	conn := PricingHandler{}
	stream, err := c.StartPricingStream(accounts.Accounts[3].ID, NewPricingStreamRequest(
		"EUR_USD", "USD_CAD", "USD_CHF",
	), conn)
	if err != nil {
		t.Fatal(err)
	}
	conn.s = stream
	select {
	case <-stream.Done():
	case <-time.After(time.Second * 30):
	}
	_ = stream.Close()
}

type PricingHandler struct {
	s *Stream
}

func (p PricingHandler) OnMessage(price *ClientPrice) error {
	b, err := price.MarshalJSON()
	if err != nil {
		return err
	}

	tm, err := price.Time.Parse()
	fmt.Println("PRICE: " + tm.String())
	fmt.Println("\t" + string(b))
	ReleaseClientPrice(price)
	return nil
}

func (p PricingHandler) OnHeartbeat(time DateTime) {
	unix, _ := time.Parse()
	fmt.Println("HEARTBEAT: " + unix.String())
}

func (p PricingHandler) OnClose() {

}
