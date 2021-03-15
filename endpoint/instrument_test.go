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
	apiToken         = ""
	apiPracticeToken = ""
)

func init() {
	b, err := ioutil.ReadFile(".env")
	if err != nil {
		panic(err)
	}
	s := strings.TrimSpace(string(b))
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		index := strings.IndexByte(line, '=')
		if index < 0 {
			continue
		}
		switch strings.TrimSpace(line[0:index]) {
		case "OANDA_API_KEY":
			apiToken = strings.TrimSpace(line[index+1:])
		case "OANDA_PRACTICE_API_KEY":
			apiPracticeToken = strings.TrimSpace(line[index+1:])
		}
	}
}

func newLiveConnection() *Connection {
	return NewConnection(apiToken, true)
}

func newPracticeConnection() *Connection {
	return NewConnection(apiPracticeToken, false)
}

func TestConnection_InstrumentCandles(t *testing.T) {
	c := newPracticeConnection()
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
