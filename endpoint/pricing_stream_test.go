package endpoint

import (
	"fmt"
	. "github.com/kamaiu/oanda-go/model"
	"testing"
	"time"
)

func TestPricingStream(t *testing.T) {
	c := newPracticeConnection()
	accounts, err := c.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	conn := PricingHandler{}
	stream, err := c.StartPricingStream(accounts.Accounts[0].ID, NewPricingStreamRequest(
		"EUR_USD", "USD_CAD", "USD_CHF", "GBP_USD",
	), conn)
	if err != nil {
		t.Fatal(err)
	}
	conn.s = stream
	select {
	case <-stream.Done():
	case <-time.After(time.Second * 30000):
	}
	_ = stream.Close()
}

type PricingHandler struct {
	s *Stream
}

func (p PricingHandler) OnMessage(price *ClientPrice) error {
	//b, err := price.MarshalJSON()
	//if err != nil {
	//	return err
	//}

	if len(price.Bids) == 0 || len(price.Asks) == 0 {
		return nil
	}

	bid := price.Bids[0].Price.AsFloat64(0.0)
	ask := price.Asks[0].Price.AsFloat64(0.0)

	//tm, err := price.Time.Parse()
	fmt.Println(price.Instrument+"   bid ", fmt.Sprintf("%.5f", bid), "  ask ", fmt.Sprintf("%.5f", ask), "    ", fmt.Sprintf("%.1f", (ask-bid)*10000), " pips")
	//fmt.Println("PRICE: " + tm.String())
	//fmt.Println("\t" + string(b))
	ReleaseClientPrice(price)
	return nil
}

func (p PricingHandler) OnHeartbeat(time DateTime) {
	//unix, _ := time.Parse()
	//fmt.Println("HEARTBEAT: " + unix.String())
}

func (p PricingHandler) OnClose() {

}
