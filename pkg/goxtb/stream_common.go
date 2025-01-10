package goxtb

import (
	"encoding/json"
	"fmt"
)

type RecordType string

const (
	StreamBalanceRecordType   RecordType = "BalanceRecord"
	StreamCandleRecordType    RecordType = "CandleRecord"
	StreamKeepAliveRecordType RecordType = "KeepAliveRecord"
	StreamNewsRecordType      RecordType = "NewsRecord"
	StreamProfitRecordType    RecordType = "ProfitRecord"
)

type StreamRecord interface {
	Type() RecordType
}

func UnmarshalStreamRecord(message []byte) (StreamRecord, error) {

	var raw map[string]string

	if err := json.Unmarshal(message, &raw); err != nil {
		return nil, fmt.Errorf("unable to unmarshal message: %v", err)
	}

	messageData := []byte(raw["data"])

	switch raw["command"] {
	case "balance":
		var balanceRecord BalanceRecord
		if err := json.Unmarshal(messageData, &balanceRecord); err != nil {
			return nil, fmt.Errorf("unable to unmarshal balance record: %v", err)
		}
		return balanceRecord, nil
	case "candle":
		var candleRecord CandleRecord
		if err := json.Unmarshal(messageData, &candleRecord); err != nil {
			return nil, fmt.Errorf("unable to unmarshal candle record: %v", err)
		}
		return candleRecord, nil
	case "keepAlive":
		var keepAliveRecord KeepAliveRecord
		if err := json.Unmarshal(messageData, &keepAliveRecord); err != nil {
			return nil, fmt.Errorf("unable to unmarshal keep alive record: %v", err)
		}
		return keepAliveRecord, nil
	case "news":
		var newsRecord NewsRecord
		if err := json.Unmarshal(messageData, &newsRecord); err != nil {
			return nil, fmt.Errorf("unable to unmarshal news record: %v", err)
		}
		return newsRecord, nil
	case "profit":
		var profitRecord ProfitRecord
		if err := json.Unmarshal(messageData, &profitRecord); err != nil {
			return nil, fmt.Errorf("unable to unmarshal profit record: %v", err)
		}
		return profitRecord, nil
	default:
		return nil, fmt.Errorf("unknown command: %s", raw["data"])
	}
}

type BalanceRecord struct {
	Balance     float32 `json:"balance"`
	Credit      float32 `json:"credit"`
	Equity      float32 `json:"equity"`
	Margin      float32 `json:"margin"`
	MarginFree  float32 `json:"marginFree"`
	MarginLevel float32 `json:"marginLevel"`
}

func (_ BalanceRecord) Type() RecordType {
	return StreamBalanceRecordType
}

type CandleRecord struct {
	Close     float32 `json:"close"`
	Ctm       int64   `json:"ctm"`
	CtmString string  `json:"ctmString"`
	High      float32 `json:"high"`
	Low       float32 `json:"low"`
	Open      float32 `json:"open"`
	QuoteId   int     `json:"quoteId"`
	Symbol    string  `json:"symbol"`
	Volume    float32 `json:"vol"`
}

func (_ CandleRecord) Type() RecordType {
	return StreamCandleRecordType
}

type KeepAliveRecord struct {
	Timestamp int64 `json:"timestamp"`
}

func (_ KeepAliveRecord) Type() RecordType {
	return StreamKeepAliveRecordType
}

type NewsRecord struct {
	Body  string `json:"body"`
	Key   string `json:"key"`
	Time  int64  `json:"time"`
	Title string `json:"title"`
}

func (_ NewsRecord) Type() RecordType {
	return StreamNewsRecordType
}

type ProfitRecord struct {
	Order    int     `json:"order"`
	Order2   int     `json:"order2"`
	Position int     `json:"position"`
	Profit   float32 `json:"profit"`
}

func (_ ProfitRecord) Type() RecordType {
	return StreamProfitRecordType
}
