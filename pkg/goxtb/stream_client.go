package goxtb

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type StreamClient struct {
	wsConn             *websocket.Conn
	wsUrl              string
	getBalanceChan     chan BalanceRecord
	getCandlesChan     chan CandleRecord
	getKeepAliveChan   chan KeepAliveRecord
	getNewsChan        chan NewsRecord
	getProfitsChan     chan ProfitsRecord
	getTickPricesChan  chan TickPricesRecord
	getTradesChan      chan TradesRecord
	getTradeStatusChan chan TradeStatusRecord
	mutex              sync.Mutex

	StreamSessionId string
}

func NewStreamClient() *StreamClient {
	return &StreamClient{
		wsConn:             nil,
		wsUrl:              "wss://wss.xtb.com:5113/realStream",
		getBalanceChan:     make(chan BalanceRecord),
		getCandlesChan:     make(chan CandleRecord),
		getKeepAliveChan:   make(chan KeepAliveRecord),
		getNewsChan:        make(chan NewsRecord),
		getProfitsChan:     make(chan ProfitsRecord),
		getTickPricesChan:  make(chan TickPricesRecord),
		getTradesChan:      make(chan TradesRecord),
		getTradeStatusChan: make(chan TradeStatusRecord),
	}
}

func NewStreamDemoClient() *StreamClient {
	return &StreamClient{
		wsConn:             nil,
		wsUrl:              "wss://wss.xtb.com:5125/demoStream",
		getBalanceChan:     make(chan BalanceRecord),
		getCandlesChan:     make(chan CandleRecord),
		getKeepAliveChan:   make(chan KeepAliveRecord),
		getNewsChan:        make(chan NewsRecord),
		getProfitsChan:     make(chan ProfitsRecord),
		getTickPricesChan:  make(chan TickPricesRecord),
		getTradesChan:      make(chan TradesRecord),
		getTradeStatusChan: make(chan TradeStatusRecord),
	}
}

func (streamClient *StreamClient) Connect() error {
	var err error
	streamClient.wsConn, _, err = websocket.DefaultDialer.Dial(streamClient.wsUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	return nil
}

func (streamClient *StreamClient) Close() {
	streamClient.wsConn.Close()
}

func (streamClient *StreamClient) SubGetBalance() (chan BalanceRecord, error) {
	msg := struct {
		Command   string `json:"command"`
		SessionId string `json:"streamSessionId"`
	}{Command: "getBalance", SessionId: streamClient.StreamSessionId}

	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getBalance message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, fmt.Errorf("failed to write getBalance message: %w", err)
	}

	return streamClient.getBalanceChan, nil
}

func (streamClient *StreamClient) UnsubGetBalance() error {
	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	data := []byte("{\"command\":\"stopBalance\"}")
	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write stopBalance message: %w", err)
	}

	return nil
}

func (streamClient *StreamClient) SubGetCandles(symbol string) (chan CandleRecord, error) {
	msg := struct {
		Command   string `json:"command"`
		SessionId string `json:"streamSessionId"`
		Symbol    string `json:"symbol"`
	}{Command: "getCandles", SessionId: streamClient.StreamSessionId, Symbol: symbol}

	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getCandles message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, fmt.Errorf("failed to write getCandles message: %w", err)
	}

	return streamClient.getCandlesChan, nil
}

func (streamClient *StreamClient) UnsubGetCandles(symbol string) error {
	msg := struct {
		Command string `json:"command"`
		Symbol  string `json:"symbol"`
	}{Command: "stopCandles", Symbol: symbol}

	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal stopCandles message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write stopCandles message: %w", err)
	}

	return nil
}

func (streamClient *StreamClient) SubGetKeepAlive() (chan KeepAliveRecord, error) {
	msg := struct {
		Command   string `json:"command"`
		SessionId string `json:"streamSessionId"`
	}{Command: "getKeepAlive", SessionId: streamClient.StreamSessionId}

	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getKeepAlive message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, fmt.Errorf("failed to write getKeepAlive message: %w", err)
	}

	return streamClient.getKeepAliveChan, nil
}

func (streamClient *StreamClient) UnsubGetKeepAlive() error {
	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	data := []byte("{\"command\":\"stopKeepAlive\"}")
	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write stopKeepAlive message: %w", err)
	}

	return nil
}

func (streamClient *StreamClient) SubNews() (chan NewsRecord, error) {
	msg := struct {
		Command   string `json:"command"`
		SessionId string `json:"streamSessionId"`
	}{Command: "getNews", SessionId: streamClient.StreamSessionId}

	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getNews message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, fmt.Errorf("failed to write getNews message: %w", err)
	}

	return streamClient.getNewsChan, nil
}

func (streamClient *StreamClient) UnsubNews() error {
	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	data := []byte("{\"command\":\"stopNews\"}")
	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write stopNews message: %w", err)
	}

	return nil
}

func (streamClient *StreamClient) SubProfits() (chan ProfitsRecord, error) {
	msg := struct {
		Command   string `json:"command"`
		SessionId string `json:"streamSessionId"`
	}{Command: "getProfits", SessionId: streamClient.StreamSessionId}

	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getProfits message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, fmt.Errorf("failed to write getProfits message: %w", err)
	}

	return streamClient.getProfitsChan, nil
}

func (streamClient *StreamClient) UnsubProfits() error {
	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	data := []byte("{\"command\":\"stopProfits\"}")
	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write stopProfits message: %w", err)
	}

	return nil
}

func (streamClient *StreamClient) SubGetTickPrices(symbol string, minArrivalTime, maxLevel int) (chan TickPricesRecord, error) {
	msg := struct {
		Command        string `json:"command"`
		SessionId      string `json:"streamSessionId"`
		Symbol         string `json:"symbol"`
		MinArrivalTime int    `json:"minArrivalTime"`
		MaxLevel       int    `json:"maxLevel"`
	}{Command: "getTickPrices", SessionId: streamClient.StreamSessionId,
		Symbol: symbol, MinArrivalTime: minArrivalTime, MaxLevel: maxLevel}

	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getTickPrices message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, fmt.Errorf("failed to write getTickPrices message: %w", err)
	}

	return streamClient.getTickPricesChan, nil
}

func (streamClient *StreamClient) UnsubGetTickPrices(symbol string) error {
	msg := struct {
		Command string `json:"command"`
		Symbol  string `json:"symbol"`
	}{Command: "stopTickPrices", Symbol: symbol}

	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal stopTickPrices message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write stopTickPrices message: %w", err)
	}

	return nil
}

func (streamClient *StreamClient) SubTrades() (chan TradesRecord, error) {
	msg := struct {
		Command   string `json:"command"`
		SessionId string `json:"streamSessionId"`
	}{Command: "getTrades", SessionId: streamClient.StreamSessionId}

	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getTrades message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, fmt.Errorf("failed to write getTrades message: %w", err)
	}

	return streamClient.getTradesChan, nil
}

func (streamClient *StreamClient) UnsubTrades() error {
	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	data := []byte("{\"command\":\"stopTrades\"}")
	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write stopTrades message: %w", err)
	}

	return nil
}

func (streamClient *StreamClient) SubTradeStatus() (chan TradeStatusRecord, error) {
	msg := struct {
		Command   string `json:"command"`
		SessionId string `json:"streamSessionId"`
	}{Command: "getTradeStatus", SessionId: streamClient.StreamSessionId}

	data, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getTradeStatus message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return nil, fmt.Errorf("failed to write getTradeStatus message: %w", err)
	}

	return streamClient.getTradeStatusChan, nil
}

func (streamClient *StreamClient) UnsubTradeStatus() error {
	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	data := []byte("{\"command\":\"stopTradeStatus\"}")
	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write stopTradeStatus message: %w", err)
	}

	return nil
}

func (streamClient *StreamClient) Ping() error {
	msg := struct {
		Command   string `json:"command"`
		SessionId string `json:"streamSessionId"`
	}{Command: "ping", SessionId: streamClient.StreamSessionId}

	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal ping message: %w", err)
	}

	streamClient.mutex.Lock()
	defer streamClient.mutex.Unlock()

	if err := streamClient.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write ping message: %w", err)
	}

	return nil
}

func (streamClient *StreamClient) Listen(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, rawBytes, err := streamClient.wsConn.ReadMessage()
			if err != nil {
				return fmt.Errorf("error reading message: %v", err)
			}

			var streamRecord StreamRecord
			err = json.Unmarshal(rawBytes, &streamRecord)
			if err != nil {
				return fmt.Errorf("unable to unmarshal raw message: %w", err)
			}

			switch streamRecord.Command {
			case "balance":
				var balanceRecord BalanceRecord
				if err := json.Unmarshal(streamRecord.Data, &balanceRecord); err != nil {
					return fmt.Errorf("unable to unmarshal balance record: %w", err)
				}
				streamClient.getBalanceChan <- balanceRecord
			case "candle":
				var candleRecord CandleRecord
				if err := json.Unmarshal(streamRecord.Data, &candleRecord); err != nil {
					return fmt.Errorf("unable to unmarshal candle record: %w", err)
				}
				streamClient.getCandlesChan <- candleRecord
			case "keepAlive":
				var keepAliveRecord KeepAliveRecord
				if err := json.Unmarshal(streamRecord.Data, &keepAliveRecord); err != nil {
					return fmt.Errorf("unable to unmarshal keep alive record: %w", err)
				}
				streamClient.getKeepAliveChan <- keepAliveRecord
			case "news":
				var newsRecord NewsRecord
				if err := json.Unmarshal(streamRecord.Data, &newsRecord); err != nil {
					return fmt.Errorf("unable to unmarshal news record: %w", err)
				}
				streamClient.getNewsChan <- newsRecord
			case "profit":
				var profitsRecord ProfitsRecord
				if err := json.Unmarshal(streamRecord.Data, &profitsRecord); err != nil {
					return fmt.Errorf("unable to unmarshal profits record: %w", err)
				}
				streamClient.getProfitsChan <- profitsRecord
			case "tickPrices":
				var tickPricesRecord TickPricesRecord
				if err := json.Unmarshal(streamRecord.Data, &tickPricesRecord); err != nil {
					return fmt.Errorf("unable to unmarshal tick prices record: %w", err)
				}
				streamClient.getTickPricesChan <- tickPricesRecord
			case "trade":
				var tradesRecord TradesRecord
				if err := json.Unmarshal(streamRecord.Data, &tradesRecord); err != nil {
					return fmt.Errorf("unable to unmarshal trades record: %w", err)
				}
				streamClient.getTradesChan <- tradesRecord
			case "tradeStatus":
				var tradeStatusRecord TradeStatusRecord
				if err := json.Unmarshal(streamRecord.Data, &tradeStatusRecord); err != nil {
					return fmt.Errorf("unable to unmarshal trade status record: %w", err)
				}
				streamClient.getTradeStatusChan <- tradeStatusRecord
			default:
				return fmt.Errorf("invalid command: %s", streamRecord.Command)
			}
		}
	}
}
