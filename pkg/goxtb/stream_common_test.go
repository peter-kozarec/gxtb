package goxtb

import (
	"reflect"
	"testing"
)

const (
	balanceRecord = `{
		"balance": 995800269.43,
		"credit": 1000.00,
		"equity": 995985397.56,
		"margin": 572634.43,
		"marginFree": 995227635.00,
		"marginLevel": 173930.41
	}`
	candleRecord = `{
		"close": 4.1849,
		"ctm": 1378369375000,
		"ctmString": "Sep 05, 2013 10:22:55 AM",
		"high": 4.1854,
		"low": 4.1848,
		"open": 4.1848,
		"quoteId": 2,
		"symbol": "EURUSD",
		"vol": 0.0
	}`
	keepAliveRecord = `{
		"timestamp": 1362944112000
	}`
	newsRecord = `{
		"body": "<html>...</html>",
		"key": "1f6da766abd29927aa854823f0105c23",
		"time": 1262944112000,
		"title": "Breaking trend"
	}`
	profitRecord = `{
		"order": 7497776,
		"order2": 7497777,
		"position": 7497776,
		"profit": 7076.52
	}`
	tickRecord = `{
		"ask": 4000.0,
		"askVolume": 15000,
		"bid": 4000.0,
		"bidVolume": 16000,
		"high": 4000.0,
		"level": 0,
		"low": 3500.0,
		"quoteId": 0,
		"spreadRaw": 0.000003,
		"spreadTable": 0.00042,
		"symbol": "KOMB.CZ",
		"timestamp": 1272529161605
	}`
	tradeRecord = `{
		"close_price": 1.3256,
		"close_time": null,
		"closed": false,
		"cmd": 0,
		"comment": "Web Trader",
		"commission": 0.0,
		"customComment": "Some text",
		"digits": 4,
		"expiration": null,
		"margin_rate": 3.9149000,
		"offset": 0,
		"open_price": 1.4,
		"open_time": 1272380927000,
		"order": 7497776,
		"order2": 1234567,
		"position": 1234567,
		"sl": 0.0,
		"state": "Modified",
		"storage": -4.46,
		"symbol": "EURUSD",
		"tp": 0.0,
		"type": 0,
		"volume": 0.10
	}`
	tradeStatusRecord = `{
		"customComment": "Some text",
		"message": null,
		"order": 43,
		"price": 1.392,
		"requestStatus": 3
	}`
)

func Test_unmarshalRecordByCommand(t *testing.T) {
	tests := []struct {
		name         string
		inputCommand string
		inputData    string
		outputRecord StreamRecord
	}{
		{"UnmarshalBalanceRecord", "balance", balanceRecord, StreamBalanceRecord{
			Balance:     995800269.43,
			Credit:      1000.00,
			Equity:      995985397.56,
			Margin:      572634.43,
			MarginFree:  995227635.00,
			MarginLevel: 173930.41,
		}},
		{"UnamrshalCandleRecord", "candle", candleRecord, StreamCandleRecord{
			Close:     4.1849,
			Ctm:       1378369375000,
			CtmString: "Sep 05, 2013 10:22:55 AM",
			High:      4.1854,
			Low:       4.1848,
			Open:      4.1848,
			QuoteId:   2,
			Symbol:    "EURUSD",
			Volume:    0.0,
		}},
		{"UnmarshalKeepAliveRecord", "keepAlive", keepAliveRecord, StreamKeepAliveRecord{
			Timestamp: 1362944112000,
		}},
		{"UnmarshalNewsRecord", "news", newsRecord, StreamNewsRecord{
			Body:  "<html>...</html>",
			Key:   "1f6da766abd29927aa854823f0105c23",
			Time:  1262944112000,
			Title: "Breaking trend",
		}},
		{"UnmarshalProfitRecord", "profit", profitRecord, StreamProfitRecord{
			Order:    7497776,
			Order2:   7497777,
			Position: 7497776,
			Profit:   7076.52,
		}},
		{"UnmarshalTickRecord", "tickPrices", tickRecord, StreamTickRecord{
			Ask:         4000.0,
			AskVolume:   15000,
			Bid:         4000.0,
			BidVolume:   16000,
			High:        4000.0,
			Level:       0,
			Low:         3500.0,
			QuoteId:     0,
			SpreadRaw:   0.000003,
			SpreadTable: 0.00042,
			Symbol:      "KOMB.CZ",
			Timestamp:   1272529161605,
		}},
		{"UnmarshalTradeRecord", "trade", tradeRecord, StreamTradeRecord{
			ClosePrice:    1.3256,
			CloseTime:     nil,
			Closed:        false,
			Cmd:           0,
			Comment:       "Web Trader",
			Commission:    0.0,
			CustomComment: "Some text",
			Digits:        4,
			Expiration:    nil,
			MarginRate:    3.9149000,
			Offset:        0,
			OpenPrice:     1.4,
			OpenTime:      1272380927000,
			Order:         7497776,
			Order2:        1234567,
			Position:      1234567,
			//Profit:        68.392, -- ToDo: Handle this case
			StopLoss:   0.0,
			State:      "Modified",
			Storage:    -4.46,
			Symbol:     "EURUSD",
			TakeProfit: 0.0,
			TradeType:  0,
			Volume:     0.10,
		}},
		{"UnmarshalTradeStatusRecord", "tradeStatus", tradeStatusRecord, StreamTradeStatusRecord{
			CustomComment: "Some text",
			Message:       nil,
			Order:         43,
			Price:         1.392,
			RequestStatus: 3,
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			record, err := unmarshalRecordByCommand(test.inputCommand, []byte(test.inputData))
			if err != nil {
				t.Fatalf("unmarshalRecordByCommand(%q, %q) returned an error: %v",
					test.inputCommand, test.inputData, err)
			}
			if !reflect.DeepEqual(record, test.outputRecord) {
				t.Errorf("unmarshalRecordByCommand(%q, %q) = %+v; want %+v",
					test.inputCommand, test.inputData, record, test.outputRecord)
			}
		})
	}
}
