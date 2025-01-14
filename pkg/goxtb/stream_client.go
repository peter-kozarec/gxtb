package goxtb

import (
	"context"
	"encoding/json"
	"fmt"

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

type RecordMessage struct {
	Value interface{}
	Err   error
}

type StreamClient struct {
	conn *websocket.Conn
	url  string

	StreamSessionId string
}

func NewStreamClient() *StreamClient {
	return &StreamClient{
		conn: nil,
		url:  "wss://ws.xtb.com/realStream",
	}
}

func NewStreamDemoClient() *StreamClient {
	return &StreamClient{
		conn: nil,
		url:  "wss://ws.xtb.com/demoStream",
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

func (c *StreamClient) GetBalance() error {

	cmd := streamCommand{
		Command:         "getBalance",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getBalance: %w", err)
	}

	return c.write(data)
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

func (c *StreamClient) GetCandles(symbol string) error {

	cmd := streamCommand{
		Command:         "getCandles",
		StreamSessionId: c.StreamSessionId,
		Symbol:          symbol,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getCandles: %w", err)
	}

	return c.write(data)
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

func (c *StreamClient) GetKeepAlive() error {

	cmd := streamCommand{
		Command:         "getKeepAlive",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getKeepAlive: %w", err)
	}

	return c.write(data)
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

func (c *StreamClient) GetNews() error {

	cmd := streamCommand{
		Command:         "getNews",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getNews: %w", err)
	}

	return c.write(data)
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

func (c *StreamClient) GetProfits() error {

	cmd := streamCommand{
		Command:         "getProfits",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getProfits: %w", err)
	}

	return c.write(data)
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

func (c *StreamClient) GetTickPrices(symbol string, minArrivalTime, maxLevel int) error {

	cmd := streamCommand{
		Command:         "getTickPrices",
		StreamSessionId: c.StreamSessionId,
		Symbol:          symbol,
		MinArrivalTime:  minArrivalTime,
		MaxLevel:        maxLevel,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getTickPrices: %w", err)
	}

	return c.write(data)
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

func (c *StreamClient) GetTrades() error {

	cmd := streamCommand{
		Command:         "getTrades",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getTrades: %w", err)
	}

	return c.write(data)
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

func (c *StreamClient) GetTradeStatus() error {

	cmd := streamCommand{
		Command:         "getTradeStatus",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getTradeStatus: %w", err)
	}

	return c.write(data)
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

func (c *StreamClient) Listen(ctx context.Context) <-chan RecordMessage {

	recordChan := make(chan RecordMessage)

	go func() {
		defer close(recordChan)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				recordMessage := RecordMessage{Value: nil, Err: nil}

				_, rawBytes, err := c.conn.ReadMessage()
				if err != nil {
					if !websocket.IsCloseError(err, websocket.CloseNormalClosure) {
						recordMessage.Err = fmt.Errorf("error reading message: %v", err)
						recordChan <- recordMessage
					}
					return
				}

				var streamRecord streamRecord
				err = json.Unmarshal(rawBytes, &streamRecord)
				if err != nil {
					recordMessage.Err = fmt.Errorf("unable to unmarshal raw message: %w", err)
					recordChan <- recordMessage
					return
				}

				switch streamRecord.Command {
				case "balance":
					var balanceRecord BalanceRecord
					if err := json.Unmarshal(streamRecord.Data, &balanceRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal balance record: %w", err)
						recordChan <- recordMessage
						return
					}
					recordMessage.Value = balanceRecord
				case "candle":
					var candleRecord CandleRecord
					if err := json.Unmarshal(streamRecord.Data, &candleRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal candle record: %w", err)
						recordChan <- recordMessage
						return
					}
					recordMessage.Value = candleRecord
				case "keepAlive":
					var keepAliveRecord KeepAliveRecord
					if err := json.Unmarshal(streamRecord.Data, &keepAliveRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal keep alive record: %w", err)
						recordChan <- recordMessage
						return
					}
					recordMessage.Value = keepAliveRecord
				case "news":
					var newsRecord NewsRecord
					if err := json.Unmarshal(streamRecord.Data, &newsRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal news record: %w", err)
						recordChan <- recordMessage
						return
					}
					recordMessage.Value = newsRecord
				case "profit":
					var profitsRecord ProfitsRecord
					if err := json.Unmarshal(streamRecord.Data, &profitsRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal profits record: %w", err)
						recordChan <- recordMessage
						return
					}
					recordMessage.Value = profitsRecord
				case "tickPrices":
					var tickPricesRecord TickPricesRecord
					if err := json.Unmarshal(streamRecord.Data, &tickPricesRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal tick prices record: %w", err)
						recordChan <- recordMessage
						return
					}
					recordMessage.Value = tickPricesRecord
				case "trade":
					var tradesRecord TradesRecord
					if err := json.Unmarshal(streamRecord.Data, &tradesRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal trades record: %w", err)
						recordChan <- recordMessage
						return
					}
					recordMessage.Value = tradesRecord
				case "tradeStatus":
					var tradeStatusRecord TradeStatusRecord
					if err := json.Unmarshal(streamRecord.Data, &tradeStatusRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal trade status record: %w", err)
						recordChan <- recordMessage
						return
					}
					recordMessage.Value = tradeStatusRecord
				default:
					recordMessage.Err = fmt.Errorf("invalid command: %s", streamRecord.Command)
					recordChan <- recordMessage
					return
				}

				recordChan <- recordMessage
			}
		}
	}()

	return recordChan
}

func (c *StreamClient) write(data []byte) error {

	if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return fmt.Errorf("failed to write: %w", err)
	}

	return nil
}
