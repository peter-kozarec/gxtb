package goxtb

import (
	"encoding/json"
	"fmt"
)

type StreamRecordType string

const (
	StreamBalanceRecordType     StreamRecordType = "BalanceRecord"
	StreamCandleRecordType      StreamRecordType = "CandleRecord"
	StreamKeepAliveRecordType   StreamRecordType = "KeepAliveRecord"
	StreamNewsRecordType        StreamRecordType = "NewsRecord"
	StreamProfitRecordType      StreamRecordType = "ProfitRecord"
	StreamTickRecordType        StreamRecordType = "TickRecord"
	StreamTradeRecordType       StreamRecordType = "TradeRecord"
	StreamTradeStatusRecordType StreamRecordType = "TradeStatusRecord"
)

type StreamRecord interface {
	Type() StreamRecordType
}

func UnmarshalStreamRecord(message []byte) (StreamRecord, error) {
	var raw map[string]interface{}

	if err := json.Unmarshal(message, &raw); err != nil {
		return nil, fmt.Errorf("unable to unmarshal message: %v", err)
	}

	data, hasData := raw["data"]
	if !hasData {
		return nil, fmt.Errorf("stream record does not contain 'data' field")
	}

	var messageData []byte
	switch d := data.(type) {
	case string:
		messageData = []byte(d)
	case map[string]interface{}:
		var err error
		messageData, err = json.Marshal(d)
		if err != nil {
			return nil, fmt.Errorf("unable to re-marshal data field: %v", err)
		}
	default:
		return nil, fmt.Errorf("'data' field is of an unsupported type")
	}

	command, ok := raw["command"].(string)
	if !ok {
		return nil, fmt.Errorf("stream record does not contain a valid 'command' field")
	}

	return unmarshalRecordByCommand(command, messageData)
}

func unmarshalRecordByCommand(command string, data []byte) (StreamRecord, error) {
	switch command {
	case "balance":
		var record StreamBalanceRecord
		if err := json.Unmarshal(data, &record); err != nil {
			return nil, fmt.Errorf("unable to unmarshal balance record: %v", err)
		}
		return record, nil
	case "candle":
		var record StreamCandleRecord
		if err := json.Unmarshal(data, &record); err != nil {
			return nil, fmt.Errorf("unable to unmarshal candle record: %v", err)
		}
		return record, nil
	case "keepAlive":
		var record StreamKeepAliveRecord
		if err := json.Unmarshal(data, &record); err != nil {
			return nil, fmt.Errorf("unable to unmarshal keep alive record: %v", err)
		}
		return record, nil
	case "news":
		var record StreamNewsRecord
		if err := json.Unmarshal(data, &record); err != nil {
			return nil, fmt.Errorf("unable to unmarshal news record: %v", err)
		}
		return record, nil
	case "profit":
		var record StreamProfitRecord
		if err := json.Unmarshal(data, &record); err != nil {
			return nil, fmt.Errorf("unable to unmarshal profit record: %v", err)
		}
		return record, nil
	case "tickPrices":
		var record StreamTickRecord
		if err := json.Unmarshal(data, &record); err != nil {
			return nil, fmt.Errorf("unable to unmarshal tick record: %v", err)
		}
		return record, nil
	case "trade":
		var record StreamTradeRecord
		if err := json.Unmarshal(data, &record); err != nil {
			return nil, fmt.Errorf("unable to unmarshal trade record: %v", err)
		}
		return record, nil
	case "tradeStatus":
		var record StreamTradeStatusRecord
		if err := json.Unmarshal(data, &record); err != nil {
			return nil, fmt.Errorf("unable to unmarshal trade status record: %v", err)
		}
		return record, nil
	default:
		return nil, fmt.Errorf("unknown command: %s", command)
	}
}

type StreamBalanceRecord struct {
	Balance     float32 `json:"balance"`
	Credit      float32 `json:"credit"`
	Equity      float32 `json:"equity"`
	Margin      float32 `json:"margin"`
	MarginFree  float32 `json:"marginFree"`
	MarginLevel float32 `json:"marginLevel"`
}

func (_ StreamBalanceRecord) Type() StreamRecordType {
	return StreamBalanceRecordType
}

type StreamCandleRecord struct {
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

func (_ StreamCandleRecord) Type() StreamRecordType {
	return StreamCandleRecordType
}

type StreamKeepAliveRecord struct {
	Timestamp int64 `json:"timestamp"`
}

func (_ StreamKeepAliveRecord) Type() StreamRecordType {
	return StreamKeepAliveRecordType
}

type StreamNewsRecord struct {
	Body  string `json:"body"`
	Key   string `json:"key"`
	Time  int64  `json:"time"`
	Title string `json:"title"`
}

func (_ StreamNewsRecord) Type() StreamRecordType {
	return StreamNewsRecordType
}

type StreamProfitRecord struct {
	Order    int     `json:"order"`
	Order2   int     `json:"order2"`
	Position int     `json:"position"`
	Profit   float32 `json:"profit"`
}

func (_ StreamProfitRecord) Type() StreamRecordType {
	return StreamProfitRecordType
}

type StreamTickRecord struct {
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

func (_ StreamTickRecord) Type() StreamRecordType {
	return StreamTickRecordType
}

type StreamTradeRecord struct {
	ClosePrice    float32  `json:"close_price"`
	CloseTime     *int64   `json:"close_time"`
	Closed        bool     `json:"closed"`
	Cmd           int      `json:"cmd"`
	Comment       string   `json:"comment"`
	Commission    float32  `json:"commission"`
	CustomComment string   `json:"customComment"`
	Digits        int      `json:"digits"`
	Expiration    *int64   `json:"expiration"`
	MarginRate    float32  `json:"margin_rate"`
	Offset        int      `json:"offset"`
	OpenPrice     float32  `json:"open_price"`
	OpenTime      int64    `json:"open_time"`
	Order         int      `json:"order"`
	Order2        int      `json:"order2"`
	Position      int      `json:"position"`
	Profit        *float32 `json:"profit"`
	StopLoss      float32  `json:"sl"`
	State         string   `json:"state"`
	Storage       float32  `json:"storage"`
	Symbol        string   `json:"symbol"`
	TakeProfit    float32  `json:"tp"`
	TradeType     int      `json:"type"`
	Volume        float32  `json:"volume"`
}

func (_ StreamTradeRecord) Type() StreamRecordType {
	return StreamTradeRecordType
}

type StreamTradeStatusRecord struct {
	CustomComment string  `json:"customComment"`
	Message       *string `json:"message"`
	Order         int     `json:"order"`
	Price         float32 `json:"price"`
	RequestStatus int     `json:"requestStatus"`
}

func (_ StreamTradeStatusRecord) Type() StreamRecordType {
	return StreamTradeStatusRecordType
}
