package endpoint

import (
	"errors"
	. "github.com/kamaiu/oanda-go/model"
	"github.com/valyala/bytebufferpool"
	"sync"
)

var (
	clientPricePool = sync.Pool{New: func() interface{} {
		return &ClientPrice{}
	}}
	ErrInstrumentsRequired = errors.New("instruments required")
)

// GET /v3/accounts/{accountID}/pricing/stream
// Get a stream of Account Prices starting from when the request is made.
// This pricing stream does not include every single price created for the
// Account, but instead will provide at most 4 prices per second
// (every 250 milliseconds) for each instrument being requested. If more than
// one price is created for an instrument during the 250 millisecond window,
// only the price in effect at the end of the window is sent. This means that
// during periods of rapid price movement, subscribers to this stream will not
// be sent every price. Pricing windows for different connections to the price
// stream are not all aligned in the same way (i.e. they are not all aligned to
// the top of the second). This means that during periods of rapid price movement,
// different subscribers may observe different prices depending on their alignment.
//
// Note: This endpoint is served by the streaming URLs.
func (c *Connection) StartPricingStream(
	accountID AccountID,
	request *PricingStreamRequest,
	handler PricingStreamHandler,
) (*Stream, error) {
	if handler == nil {
		return nil, ErrNilRequest
	}
	if request == nil {
		return nil, ErrNilRequest
	}
	if len(request.Instruments) == 0 {
		return nil, ErrInstrumentsRequired
	}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.streamingHost)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/pricing/stream?")
	request.AppendQuery(url)
	return c.doStream(url, &pricingHandler{handler: handler})
}

func AcquireClientPrice() *ClientPrice {
	return clientPricePool.Get().(*ClientPrice)
}

func ReleaseClientPrice(price *ClientPrice) {
	clientPricePool.Put(price)
}

type PricingStreamHandler interface {
	OnMessage(price *ClientPrice) error

	OnHeartbeat(time DateTime)

	OnClose()
}

type pricingHandler struct {
	handler PricingStreamHandler
}

func (t *pricingHandler) handle(b []byte) error {
	price := AcquireClientPrice()
	err := price.UnmarshalJSON(b)
	if err != nil {
		return err
	}
	if price.Type == "HEARTBEAT" {
		t.handler.OnHeartbeat(price.Time)
		return nil
	} else {
		return t.handler.OnMessage(price)
	}
}

func (t *pricingHandler) onClose() {
	t.handler.OnClose()
}
