package oanda

import (
	"github.com/kamaiu/oanda-go/model"
	"github.com/kamaiu/oanda-go/wal"
)

type TxLog struct {
	accountID model.AccountID
	first     model.TransactionID
	last      model.TransactionID
	wal       *wal.Log
}

// Truncate the log from the first record up to and including the record closest
// to or inclusive of but not greater than the last TransactionID supplied.
func (t *TxLog) Truncate(last model.TransactionID) {

}
