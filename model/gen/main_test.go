package main

import (
	"github.com/mailru/easyjson/jlexer"
	"testing"
)

func TestParse(t *testing.T) {
	in := &jlexer.Lexer{
		Data: []byte(`
{
	"transactions": [
		{
			"type": "MARKET_ORDER"
		},
		{
			"type": "LIMIT_ORDER"
		},
	],
}
`),
	}
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "transactions":
			if in.IsNull() {
				in.Skip()
				//out.Prices = nil
			} else {
				in.Delim('[')
				//if out.Prices == nil {
				//	if !in.IsDelim(']') {
				//		out.Prices = make([]struct {
				//			Asks []struct {
				//				Liquidity int    `json:"liquidity"`
				//				Price     string `json:"price"`
				//			} `json:"asks"`
				//			Bids []struct {
				//				Liquidity int    `json:"liquidity"`
				//				Price     string `json:"price"`
				//			} `json:"bids"`
				//			CloseoutAsk                string `json:"closeoutAsk"`
				//			CloseoutBid                string `json:"closeoutBid"`
				//			Instrument                 string `json:"instrument"`
				//			QuoteHomeConversionFactors struct {
				//				NegativeUnits string `json:"negativeUnits"`
				//				PositiveUnits string `json:"positiveUnits"`
				//			} `json:"quoteHomeConversionFactors"`
				//			Status         string    `json:"status"`
				//			Time           time.Time `json:"time"`
				//			UnitsAvailable struct {
				//				Default struct {
				//					Long  string `json:"long"`
				//					Short string `json:"short"`
				//				} `json:"default"`
				//				OpenOnly struct {
				//					Long  string `json:"long"`
				//					Short string `json:"short"`
				//				} `json:"openOnly"`
				//				ReduceFirst struct {
				//					Long  string `json:"long"`
				//					Short string `json:"short"`
				//				} `json:"reduceFirst"`
				//				ReduceOnly struct {
				//					Long  string `json:"long"`
				//					Short string `json:"short"`
				//				} `json:"reduceOnly"`
				//			} `json:"unitsAvailable"`
				//		}, 0, 0)
				//	} else {
				//		out.Prices = []struct {
				//			Asks []struct {
				//				Liquidity int    `json:"liquidity"`
				//				Price     string `json:"price"`
				//			} `json:"asks"`
				//			Bids []struct {
				//				Liquidity int    `json:"liquidity"`
				//				Price     string `json:"price"`
				//			} `json:"bids"`
				//			CloseoutAsk                string `json:"closeoutAsk"`
				//			CloseoutBid                string `json:"closeoutBid"`
				//			Instrument                 string `json:"instrument"`
				//			QuoteHomeConversionFactors struct {
				//				NegativeUnits string `json:"negativeUnits"`
				//				PositiveUnits string `json:"positiveUnits"`
				//			} `json:"quoteHomeConversionFactors"`
				//			Status         string    `json:"status"`
				//			Time           time.Time `json:"time"`
				//			UnitsAvailable struct {
				//				Default struct {
				//					Long  string `json:"long"`
				//					Short string `json:"short"`
				//				} `json:"default"`
				//				OpenOnly struct {
				//					Long  string `json:"long"`
				//					Short string `json:"short"`
				//				} `json:"openOnly"`
				//				ReduceFirst struct {
				//					Long  string `json:"long"`
				//					Short string `json:"short"`
				//				} `json:"reduceFirst"`
				//				ReduceOnly struct {
				//					Long  string `json:"long"`
				//					Short string `json:"short"`
				//				} `json:"reduceOnly"`
				//			} `json:"unitsAvailable"`
				//		}{}
				//	}
				//} else {
				//	out.Prices = (out.Prices)[:0]
				//}
				for !in.IsDelim(']') {
					//var v1 struct {
					//	Asks []struct {
					//		Liquidity int    `json:"liquidity"`
					//		Price     string `json:"price"`
					//	} `json:"asks"`
					//	Bids []struct {
					//		Liquidity int    `json:"liquidity"`
					//		Price     string `json:"price"`
					//	} `json:"bids"`
					//	CloseoutAsk                string `json:"closeoutAsk"`
					//	CloseoutBid                string `json:"closeoutBid"`
					//	Instrument                 string `json:"instrument"`
					//	QuoteHomeConversionFactors struct {
					//		NegativeUnits string `json:"negativeUnits"`
					//		PositiveUnits string `json:"positiveUnits"`
					//	} `json:"quoteHomeConversionFactors"`
					//	Status         string    `json:"status"`
					//	Time           time.Time `json:"time"`
					//	UnitsAvailable struct {
					//		Default struct {
					//			Long  string `json:"long"`
					//			Short string `json:"short"`
					//		} `json:"default"`
					//		OpenOnly struct {
					//			Long  string `json:"long"`
					//			Short string `json:"short"`
					//		} `json:"openOnly"`
					//		ReduceFirst struct {
					//			Long  string `json:"long"`
					//			Short string `json:"short"`
					//		} `json:"reduceFirst"`
					//		ReduceOnly struct {
					//			Long  string `json:"long"`
					//			Short string `json:"short"`
					//		} `json:"reduceOnly"`
					//	} `json:"unitsAvailable"`
					//}
					//easyjson9f04e93cDecode(in, &v1)
					//out.Prices = append(out.Prices, v1)
					in.SkipRecursive()
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
