package goxtb

import (
	"encoding/json"
	"fmt"
)

type RecordType string

const (
	StreamBalanceRecordType     RecordType = "BalanceRecord"
	StreamCandleRecordType      RecordType = "CandleRecord"
	StreamKeepAliveRecordType   RecordType = "KeepAliveRecord"
	StreamNewsRecordType        RecordType = "NewsRecord"
	StreamProfitRecordType      RecordType = "ProfitRecord"
	StreamTickRecordType        RecordType = "TickRecord"
	StreamTradeRecordType       RecordType = "TradeRecord"
	StreamTradeStatusRecordType RecordType = "TradeStatusRecord"
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
	case "tickPrices":
		var tickRecord TickRecord
		if err := json.Unmarshal(messageData, &tickRecord); err != nil {
			return nil, fmt.Errorf("unable to unmarshal tick record: %v", err)
		}
		return tickRecord, nil
	case "trade":
		var tradeRecord TradeRecord
		if err := json.Unmarshal(messageData, &tradeRecord); err != nil {
			return nil, fmt.Errorf("unable to unmarshal trade record: %v", err)
		}
		return tradeRecord, nil
	case "tradeStatus":
		var tradeStatusRecord TradeStatusRecord
		if err := json.Unmarshal(messageData, &tradeStatusRecord); err != nil {
			return nil, fmt.Errorf("unable to unmarshal trade status record: %v", err)
		}
		return tradeStatusRecord, nil
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

type TickRecord struct {
	Ask         float32 `json:"ask"`
	AskVolume   int     `json:"askVolume"`
	Bid         float32 `json:"bid"`
	BidVolume   int     `json:"bidVolume"`
	High        float32 `json:"high"`
	Level       int     `json:"level"`
	Low         float32 `json:"low"`
	QuoteId     int     `json:"quoteId"`
	SpreadRaw   float32 `json:"spreadRaw"`
	SpreadTable float32 `json:"spreadTable"`
	Symbol      string  `json:"symbol"`
	Timestamp   int64   `json:"timestamp"`
}

func (_ TickRecord) Type() RecordType {
	return StreamTickRecordType
}

type TradeRecord struct {
	ClosePrice    float32 `json:"close_price"`
	CloseTime     int64   `json:"close_time"`
	Closed        bool    `json:"closed"`
	Cmd           int     `json:"cmd"`
	Comment       string  `json:"comment"`
	Commission    float32 `json:"commission"`
	CustomComment string  `json:"customComment"`
	Digits        int     `json:"digits"`
	Expiration    int64   `json:"expiration"`
	MarginRate    float32 `json:"margin_rate"`
	Offset        int     `json:"offset"`
	OpenPrice     float32 `json:"open_price"`
	OpenTime      int64   `json:"open_time"`
	Order         int     `json:"order"`
	Order2        int     `json:"order2"`
	Position      int     `json:"position"`
	Profit        float32 `json:"profit"`
	StopLoss      float32 `json:"sl"`
	State         string  `json:"state"`
	Storage       float32 `json:"storage"`
	TakeProfit    float32 `json:"tp"`
	TradeType     int     `json:"type"`
	Volume        float32 `json:"volume"`
}

func (_ TradeRecord) Type() RecordType {
	return StreamTradeRecordType
}

type TradeStatusRecord struct {
	CustomComment string  `json:"customComment"`
	Message       string  `json:"message"`
	Order         int     `json:"order"`
	Price         float32 `json:"price"`
	RequestStatus int     `json:"requestStatus"`
}

func (_ TradeStatusRecord) Type() RecordType {
	return StreamTradeStatusRecordType
}
