package model

import (
	"encoding/json"
	"github.com/mailru/easyjson"
)

type Message interface {
	easyjson.MarshalerUnmarshaler
	json.Marshaler
	json.Unmarshaler
}
