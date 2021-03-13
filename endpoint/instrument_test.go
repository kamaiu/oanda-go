package endpoint

import (
	"fmt"
	. "github.com/kamaiu/oanda-go/model"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

var (
	apiToken = ""
)

func init() {
	b, err := ioutil.ReadFile(".env")
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(b))
	lines := strings.Split(s, "\n")

	tokenName := "OANDA_API_KEY"
	for _, line := range lines {
		line = strings.TrimSpace(line)
		index := strings.IndexByte(line, '=')
		if index < 0 {
			continue
		}
		switch strings.TrimSpace(line[0:index]) {
		case tokenName:
			apiToken = strings.TrimSpace(line[index+1:])
		}
	}
}

func TestConnection_InstrumentCandles(t *testing.T) {
	c := NewConnection(apiToken, true)
	resp, err := c.InstrumentCandles(
		NewInstrumentCandlesRequest("EUR_USD", time.Now().Add(-(time.Hour * 760))).
			WithGranularity(CandlestickGranularity_H1).
			WithPrice(PricingComponent_BID_ASK_MID).
			WithCount(10),
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(resp)
}
