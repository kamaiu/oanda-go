package endpoint

import (
	. "github.com/kamaiu/oanda-go/model"
	"github.com/valyala/bytebufferpool"
)

type TxStreamHandler interface {
	/*
		switch v := msg.(type) {
		case *TransactionHeartbeat:
		case *CreateTransaction:
		case *CloseTransaction:
		case *ReopenTransaction:
		case *ClientConfigureTransaction:
		case *ClientConfigureRejectTransaction:
		case *TransferFundsTransaction:
		case *TransferFundsRejectTransaction:
		case *MarketOrderTransaction:
		case *MarketOrderRejectTransaction:
		case *FixedPriceOrderTransaction:
		case *LimitOrderTransaction:
		case *LimitOrderRejectTransaction:
		case *StopOrderTransaction:
		case *StopOrderRejectTransaction:
		case *MarketIfTouchedOrderTransaction:
		case *MarketIfTouchedOrderRejectTransaction:
		case *TakeProfitOrderTransaction:
		case *TakeProfitOrderRejectTransaction:
		case *StopLossOrderTransaction:
		case *StopLossOrderRejectTransaction:
		case *GuaranteedStopLossOrderTransaction:
		case *GuaranteedStopLossOrderRejectTransaction:
		case *TrailingStopLossOrderTransaction:
		case *TrailingStopLossOrderRejectTransaction:
		case *OrderFillTransaction:
		case *OrderCancelTransaction:
		case *OrderCancelRejectTransaction:
		case *OrderClientExtensionsModifyTransaction:
		case *OrderClientExtensionsModifyRejectTransaction:
		case *TradeClientExtensionsModifyTransaction:
		case *TradeClientExtensionsModifyRejectTransaction:
		case *MarginCallEnterTransaction:
		case *MarginCallExtendTransaction:
		case *MarginCallExitTransaction:
		case *DelayedTradeClosureTransaction:
		case *DailyFinancingTransaction:
		case *DividendAdjustmentTransaction:
		case *ResetResettablePLTransaction:
		}
	*/
	OnMessage(msg interface{}) error

	OnHeartbeat(time DateTime, last TransactionID) error

	OnClose()
}

// GET /v3/accounts/{accountID}/transactions/stream
// Get a stream of Transactions for an Account starting from when the request is made.
//
// Note: This endpoint is served by the streaming URLs.
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
		return t.handler.OnHeartbeat(t.tx.Time, t.tx.LastTransactionID)
	} else {
		return t.handler.OnMessage(t.tx.Parse())
	}
}

func (t *txHandler) onClose() {
	t.handler.OnClose()
}
