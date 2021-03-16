package endpoint

import (
	"fmt"
	. "github.com/kamaiu/oanda-go/model"
	"testing"
	"time"
)

func TestTransactionStream(t *testing.T) {
	c := newPracticeConnection()
	accounts, err := c.Accounts()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("Account: " + accounts.Accounts[0].ID)

	handler := TxHandler{}
	stream, err := c.StartTransactionStream(accounts.Accounts[0].ID, handler)
	if err != nil {
		t.Fatal(err)
	}
	handler.s = stream
	select {
	case <-stream.Done():
	case <-time.After(time.Second * 3000):
	}
	_ = stream.Close()
}

type TxHandler struct {
	s *Stream
}

func (p TxHandler) OnMessage(message TransactionMessage) error {
	b, err := message.MarshalJSON()
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func (p TxHandler) OnHeartbeat(time DateTime, last TransactionID) error {
	unix, _ := time.Parse()
	fmt.Println("HEARTBEAT: " + unix.String())
	return nil
}

func (p TxHandler) OnClose() {

}
