package endpoint

import (
	. "github.com/kamaiu/oanda-go/model"
	"github.com/valyala/bytebufferpool"
)

type TxStreamHandler interface {
	OnMessage(msg interface{}) error

	OnHeartbeat(last TransactionID, time DateTime) error

	OnClose()
}

func (c *Connection) StartTransactionStream(
	accountID AccountID,
	handler TxStreamHandler,
) (*Stream, error) {
	if handler == nil {
		return nil, ErrNilRequest
	}
	url := bytebufferpool.Get()
	_, _ = url.WriteString(c.streamingHost)
	_, _ = url.WriteString("/v3/accounts/")
	_, _ = url.WriteString((string)(accountID))
	_, _ = url.WriteString("/transactions/stream")
	return c.doStream(url, &txHandler{handler: handler})
}

type txHandler struct {
	tx      TransactionParser
	handler TxStreamHandler
}

func (t *txHandler) handle(msg []byte) error {
	err := t.tx.UnmarshalJSON(msg)
	if err != nil {
		return err
	}
	if t.tx.Type == "HEARTBEAT" {
		return t.handler.OnHeartbeat(t.tx.LastTransactionID, t.tx.Time)
	} else {
		return t.handler.OnMessage(t.tx.Parse())
	}
}

func (t *txHandler) onClose() {
	t.handler.OnClose()
}
