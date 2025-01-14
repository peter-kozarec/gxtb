package goxtb

import (
	"context"
	"encoding/json"
	"fmt"
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
	conn wsConn
	url  string

	StreamSessionId string
}

func NewStreamClient() *StreamClient {
	return &StreamClient{
		conn: new(wsImpl),
		url:  "wss://ws.xtb.com/realStream",
	}
}

func NewStreamDemoClient() *StreamClient {
	return &StreamClient{
		conn: new(wsImpl),
		url:  "wss://ws.xtb.com/demoStream",
	}
}

func (c *StreamClient) Connect() error {
	return c.conn.connect(c.url)
}

func (c *StreamClient) Disconnect() error {
	return c.conn.disconnect()
}

func (c *StreamClient) GetBalance(ctx context.Context) error {

	cmd := streamCommand{
		Command:         "getBalance",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getBalance: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) StopBalance(ctx context.Context) error {

	cmd := streamCommand{
		Command: "stopBalance",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopBalance: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) GetCandles(ctx context.Context, symbol string) error {

	cmd := streamCommand{
		Command:         "getCandles",
		StreamSessionId: c.StreamSessionId,
		Symbol:          symbol,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getCandles: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) StopCandles(ctx context.Context, symbol string) error {

	cmd := streamCommand{
		Command: "stopCandles",
		Symbol:  symbol,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopCandles: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) GetKeepAlive(ctx context.Context) error {

	cmd := streamCommand{
		Command:         "getKeepAlive",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getKeepAlive: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) StopKeepAlive(ctx context.Context) error {

	cmd := streamCommand{
		Command: "stopKeepAlive",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopKeepAlive: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) GetNews(ctx context.Context) error {

	cmd := streamCommand{
		Command:         "getNews",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getNews: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) StopNews(ctx context.Context) error {

	cmd := streamCommand{
		Command: "stopNews",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopNews: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) GetProfits(ctx context.Context) error {

	cmd := streamCommand{
		Command:         "getProfits",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getProfits: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) StopProfits(ctx context.Context) error {

	cmd := streamCommand{
		Command: "stopProfits",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopProfits: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) GetTickPrices(ctx context.Context, symbol string, minArrivalTime, maxLevel int) error {

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

	return c.conn.write(ctx, data)
}

func (c *StreamClient) StopTickPrices(ctx context.Context, symbol string) error {

	cmd := streamCommand{
		Command: "stopTickPrices",
		Symbol:  symbol,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopTickPrices: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) GetTrades(ctx context.Context) error {

	cmd := streamCommand{
		Command:         "getTrades",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getTrades: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) StopTrades(ctx context.Context) error {

	cmd := streamCommand{
		Command: "stopTrades",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopTrades: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) GetTradeStatus(ctx context.Context) error {

	cmd := streamCommand{
		Command:         "getTradeStatus",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal getTradeStatus: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) StopTradeStatus(ctx context.Context) error {

	cmd := streamCommand{
		Command: "stopTradeStatus",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal stopTradeStatus: %w", err)
	}

	return c.conn.write(ctx, data)
}

func (c *StreamClient) Ping(ctx context.Context) error {

	cmd := streamCommand{
		Command:         "ping",
		StreamSessionId: c.StreamSessionId,
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("failed to marshal ping: %w", err)
	}

	return c.conn.write(ctx, data)
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
				recordMessage := RecordMessage{}

				rawBytes, err := c.conn.read(ctx)
				if err != nil {
					recordMessage.Err = fmt.Errorf("error reading message: %v", err)
					recordChan <- recordMessage
					return
				}

				var streamRecord streamRecord
				err = json.Unmarshal(rawBytes, &streamRecord)
				if err != nil {
					recordMessage.Err = fmt.Errorf("unable to unmarshal raw message: %w", err)
					recordChan <- recordMessage
					continue
				}

				switch streamRecord.Command {
				case "balance":
					var balanceRecord BalanceRecord
					if err := json.Unmarshal(streamRecord.Data, &balanceRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal balance record: %w", err)
					} else {
						recordMessage.Value = balanceRecord
					}
				case "candle":
					var candleRecord CandleRecord
					if err := json.Unmarshal(streamRecord.Data, &candleRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal candle record: %w", err)
					} else {
						recordMessage.Value = candleRecord
					}
				case "keepAlive":
					var keepAliveRecord KeepAliveRecord
					if err := json.Unmarshal(streamRecord.Data, &keepAliveRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal keep alive record: %w", err)
					} else {
						recordMessage.Value = keepAliveRecord
					}
				case "news":
					var newsRecord NewsRecord
					if err := json.Unmarshal(streamRecord.Data, &newsRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal news record: %w", err)
					} else {
						recordMessage.Value = newsRecord
					}
				case "profit":
					var profitsRecord ProfitsRecord
					if err := json.Unmarshal(streamRecord.Data, &profitsRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal profits record: %w", err)
					} else {
						recordMessage.Value = profitsRecord
					}
				case "tickPrices":
					var tickPricesRecord TickPricesRecord
					if err := json.Unmarshal(streamRecord.Data, &tickPricesRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal tick prices record: %w", err)
					} else {
						recordMessage.Value = tickPricesRecord
					}
				case "trade":
					var tradesRecord TradesRecord
					if err := json.Unmarshal(streamRecord.Data, &tradesRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal trades record: %w", err)
					} else {
						recordMessage.Value = tradesRecord
					}
				case "tradeStatus":
					var tradeStatusRecord TradeStatusRecord
					if err := json.Unmarshal(streamRecord.Data, &tradeStatusRecord); err != nil {
						recordMessage.Err = fmt.Errorf("unable to unmarshal trade status record: %w", err)
					} else {
						recordMessage.Value = tradeStatusRecord
					}
				default:
					recordMessage.Err = fmt.Errorf("invalid command: %s", streamRecord.Command)
				}

				recordChan <- recordMessage
			}
		}
	}()

	return recordChan
}
