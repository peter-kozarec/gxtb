package goxtb

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type streamCommand struct {
	Command         string `json:"command"`
	StreamSessionId string `json:"streamSessionId,omitempty"`
	Symbol          string `json:"symbol,omitempty"`
	MinArrivalTime  int    `json:"minArrivalTime,omitempty"`
	MaxLevel        int    `json:"maxLevel,omitempty"`
}

type streamRecord struct {
	Command string          `json:"command"`
	Data    json.RawMessage `json:"data"`
}

type StreamClient struct {
	conn               *websocket.Conn
	url                string
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
		conn:               nil,
		url:                "wss://ws.xtb.com/realStream",
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
		conn:               nil,
		url:                "wss://ws.xtb.com/demoStream",
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

func (c *StreamClient) Connect() error {
	var err error
	c.conn, _, err = websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return fmt.Errorf("websocket failed to connect: %w", err)
	}
	return nil
}

func (c *StreamClient) Disconnect() {
	c.conn.Close()
}

func (c *StreamClient) GetBalance() (chan BalanceRecord, error) {

	cmd := streamCommand{
		Command:         "getBalance",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getBalance: %w", err)
	}

	return c.getBalanceChan, c.write(data)
}

func (c *StreamClient) StopBalance() error {

	cmd := streamCommand{
		Command: "stopBalance",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopBalance: %w", err)
	}

	return c.write(data)
}

func (c *StreamClient) GetCandles(symbol string) (chan CandleRecord, error) {

	cmd := streamCommand{
		Command:         "getCandles",
		StreamSessionId: c.StreamSessionId,
		Symbol:          symbol,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getCandles: %w", err)
	}

	return c.getCandlesChan, c.write(data)
}

func (c *StreamClient) StopCandles(symbol string) error {

	cmd := streamCommand{
		Command: "stopCandles",
		Symbol:  symbol,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopCandles: %w", err)
	}

	return c.write(data)
}

func (c *StreamClient) GetKeepAlive() (chan KeepAliveRecord, error) {

	cmd := streamCommand{
		Command:         "getKeepAlive",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getKeepAlive: %w", err)
	}

	return c.getKeepAliveChan, c.write(data)
}

func (c *StreamClient) StopKeepAlive() error {

	cmd := streamCommand{
		Command: "stopKeepAlive",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopKeepAlive: %w", err)
	}

	return c.write(data)
}

func (c *StreamClient) GetNews() (chan NewsRecord, error) {

	cmd := streamCommand{
		Command:         "getNews",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getNews: %w", err)
	}

	return c.getNewsChan, c.write(data)
}

func (c *StreamClient) StopNews() error {

	cmd := streamCommand{
		Command: "stopNews",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopNews: %w", err)
	}

	return c.write(data)
}

func (c *StreamClient) GetProfits() (chan ProfitsRecord, error) {

	cmd := streamCommand{
		Command:         "getProfits",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getProfits: %w", err)
	}

	return c.getProfitsChan, c.write(data)
}

func (c *StreamClient) StopProfits() error {

	cmd := streamCommand{
		Command: "stopProfits",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopProfits: %w", err)
	}

	return c.write(data)
}

func (c *StreamClient) GetTickPrices(symbol string, minArrivalTime, maxLevel int) (chan TickPricesRecord, error) {

	cmd := streamCommand{
		Command:         "getTickPrices",
		StreamSessionId: c.StreamSessionId,
		Symbol:          symbol,
		MinArrivalTime:  minArrivalTime,
		MaxLevel:        maxLevel,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getTickPrices: %w", err)
	}

	return c.getTickPricesChan, c.write(data)
}

func (c *StreamClient) StopTickPrices(symbol string) error {

	cmd := streamCommand{
		Command: "stopTickPrices",
		Symbol:  symbol,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopTickPrices: %w", err)
	}

	return c.write(data)
}

func (c *StreamClient) GetTrades() (chan TradesRecord, error) {

	cmd := streamCommand{
		Command:         "getTrades",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getTrades: %w", err)
	}

	return c.getTradesChan, c.write(data)
}

func (c *StreamClient) StopTrades() error {

	cmd := streamCommand{
		Command: "stopTrades",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopTrades: %w", err)
	}

	return c.write(data)
}

func (c *StreamClient) GetTradeStatus() (chan TradeStatusRecord, error) {

	cmd := streamCommand{
		Command:         "getTradeStatus",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal getTradeStatus: %w", err)
	}

	return c.getTradeStatusChan, c.write(data)
}

func (c *StreamClient) StopTradeStatus() error {

	cmd := streamCommand{
		Command: "stopTradeStatus",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopTradeStatus: %w", err)
	}

	return c.write(data)
}

func (c *StreamClient) Ping() error {

	cmd := streamCommand{
		Command:         "ping",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal ping: %w", err)
	}

	return c.write(data)
}

func (c *StreamClient) Listen(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			_, rawBytes, err := c.conn.ReadMessage()
			if err != nil {
				return fmt.Errorf("error reading message: %v", err)
			}

			var streamRecord streamRecord
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
				c.getBalanceChan <- balanceRecord
			case "candle":
				var candleRecord CandleRecord
				if err := json.Unmarshal(streamRecord.Data, &candleRecord); err != nil {
					return fmt.Errorf("unable to unmarshal candle record: %w", err)
				}
				c.getCandlesChan <- candleRecord
			case "keepAlive":
				var keepAliveRecord KeepAliveRecord
				if err := json.Unmarshal(streamRecord.Data, &keepAliveRecord); err != nil {
					return fmt.Errorf("unable to unmarshal keep alive record: %w", err)
				}
				c.getKeepAliveChan <- keepAliveRecord
			case "news":
				var newsRecord NewsRecord
				if err := json.Unmarshal(streamRecord.Data, &newsRecord); err != nil {
					return fmt.Errorf("unable to unmarshal news record: %w", err)
				}
				c.getNewsChan <- newsRecord
			case "profit":
				var profitsRecord ProfitsRecord
				if err := json.Unmarshal(streamRecord.Data, &profitsRecord); err != nil {
					return fmt.Errorf("unable to unmarshal profits record: %w", err)
				}
				c.getProfitsChan <- profitsRecord
			case "tickPrices":
				var tickPricesRecord TickPricesRecord
				if err := json.Unmarshal(streamRecord.Data, &tickPricesRecord); err != nil {
					return fmt.Errorf("unable to unmarshal tick prices record: %w", err)
				}
				c.getTickPricesChan <- tickPricesRecord
			case "trade":
				var tradesRecord TradesRecord
				if err := json.Unmarshal(streamRecord.Data, &tradesRecord); err != nil {
					return fmt.Errorf("unable to unmarshal trades record: %w", err)
				}
				c.getTradesChan <- tradesRecord
			case "tradeStatus":
				var tradeStatusRecord TradeStatusRecord
				if err := json.Unmarshal(streamRecord.Data, &tradeStatusRecord); err != nil {
					return fmt.Errorf("unable to unmarshal trade status record: %w", err)
				}
				c.getTradeStatusChan <- tradeStatusRecord
			default:
				return fmt.Errorf("invalid command: %s", streamRecord.Command)
			}
		}
	}
}

func (c *StreamClient) write(data []byte) error {

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}

	return nil
}
