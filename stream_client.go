package gxtb

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type streamCommand struct {
	Command         string `json:"command"`
	StreamSessionId string `json:"streamSessionId"`
	Symbol          string `json:"symbol,omitempty"`
	MinArrivalTime  int    `json:"minArrivalTime,omitempty"`
	MaxLevel        int    `json:"maxLevel,omitempty"`
}

type streamData struct {
	Command string          `json:"command"`
	Data    json.RawMessage `json:"data"`
}

type GetBalanceCb func(Balance)
type GetCandlesCb func(Candle)
type GetKeepAliveCb func(KeepAlive)
type GetNewsCb func(News)
type GetProfitsCb func(Profit)
type GetTickPricesCb func(TickPrice)
type GetTradesCb func(Trade)
type GetTradeStatusCb func(TradeStatus)

type StreamClient struct {
	websocketConnection

	sessionId       string
	opts            StreamOptions
	listenCtxCancel context.CancelFunc
	mu              sync.Mutex

	balanceCb     GetBalanceCb
	candlesCb     map[string]GetCandlesCb
	keepAliveCb   GetKeepAliveCb
	newsCb        GetNewsCb
	profitsCb     GetProfitsCb
	tickPricesCb  map[string]GetTickPricesCb
	tradesCb      GetTradesCb
	tradeStatusCb GetTradeStatusCb
}

func NewStreamClient(opts StreamOptions) *StreamClient {

	return &StreamClient{
		opts:         opts,
		candlesCb:    make(map[string]GetCandlesCb),
		tickPricesCb: make(map[string]GetTickPricesCb),
	}
}

func (c *StreamClient) Connect(ctx context.Context) error {

	if err := c.connect(ctx, c.opts.GetUrl()); err != nil {
		return err
	}

	return nil
}

func (c *StreamClient) Disconnect() error {

	if c.listenCtxCancel != nil {
		c.listenCtxCancel()
	}

	return c.disconnect()
}

func (c *StreamClient) SetSessionId(sessionId string) {
	c.sessionId = sessionId
}

func (c *StreamClient) GetBalance(ctx context.Context, cb GetBalanceCb) error {

	c.balanceCb = cb

	return c.sendCommand(ctx, streamCommand{
		Command:         "getBalance",
		StreamSessionId: c.sessionId,
	})
}

func (c *StreamClient) StopBalance(ctx context.Context) error {

	c.balanceCb = nil

	return c.sendCommand(ctx, streamCommand{
		Command: "stopBalance",
	})
}

func (c *StreamClient) GetCandles(ctx context.Context, symbol string, cb GetCandlesCb) error {

	c.candlesCb[symbol] = cb

	return c.sendCommand(ctx, streamCommand{
		Command:         "getCandles",
		StreamSessionId: c.sessionId,
		Symbol:          symbol,
	})
}

func (c *StreamClient) StopCandles(ctx context.Context, symbol string) error {

	delete(c.candlesCb, symbol)

	return c.sendCommand(ctx, streamCommand{
		Command: "stopCandles",
		Symbol:  symbol,
	})
}

func (c *StreamClient) GetKeepAlive(ctx context.Context, cb GetKeepAliveCb) error {

	c.keepAliveCb = cb

	return c.sendCommand(ctx, streamCommand{
		Command:         "getKeepAlive",
		StreamSessionId: c.sessionId,
	})
}

func (c *StreamClient) StopKeepAlive(ctx context.Context) error {

	c.keepAliveCb = nil

	return c.sendCommand(ctx, streamCommand{
		Command: "stopKeepAlive",
	})
}

func (c *StreamClient) GetNews(ctx context.Context, cb GetNewsCb) error {

	c.newsCb = cb

	return c.sendCommand(ctx, streamCommand{
		Command:         "getNews",
		StreamSessionId: c.sessionId,
	})
}

func (c *StreamClient) StopNews(ctx context.Context) error {

	c.newsCb = nil

	return c.sendCommand(ctx, streamCommand{
		Command: "stopNews",
	})
}

func (c *StreamClient) GetProfits(ctx context.Context, cb GetProfitsCb) error {

	c.profitsCb = cb

	return c.sendCommand(ctx, streamCommand{
		Command:         "getProfits",
		StreamSessionId: c.sessionId,
	})
}

func (c *StreamClient) StopProfits(ctx context.Context) error {

	c.profitsCb = nil

	return c.sendCommand(ctx, streamCommand{
		Command: "stopProfits",
	})
}

func (c *StreamClient) GetTickPrices(ctx context.Context, symbol string, minArrivalTime, maxLevel int, cb GetTickPricesCb) error {

	c.tickPricesCb[symbol] = cb

	return c.sendCommand(ctx, streamCommand{
		Command:         "getTickPrices",
		StreamSessionId: c.sessionId,
		Symbol:          symbol,
		MinArrivalTime:  minArrivalTime,
		MaxLevel:        maxLevel,
	})
}

func (c *StreamClient) StopTickPrices(ctx context.Context, symbol string) error {

	delete(c.tickPricesCb, symbol)

	return c.sendCommand(ctx, streamCommand{
		Command: "stopTickPrices",
		Symbol:  symbol,
	})
}

func (c *StreamClient) GetTrades(ctx context.Context, cb GetTradesCb) error {

	c.tradesCb = cb

	return c.sendCommand(ctx, streamCommand{
		Command:         "getTrades",
		StreamSessionId: c.sessionId,
	})
}

func (c *StreamClient) StopTrades(ctx context.Context) error {

	c.tradesCb = nil

	return c.sendCommand(ctx, streamCommand{
		Command: "stopTrades",
	})
}

func (c *StreamClient) GetTradeStatus(ctx context.Context, cb GetTradeStatusCb) error {

	c.tradeStatusCb = cb

	return c.sendCommand(ctx, streamCommand{
		Command:         "getTradeStatus",
		StreamSessionId: c.sessionId,
	})
}

func (c *StreamClient) StopTradeStatus(ctx context.Context) error {

	c.tradeStatusCb = nil

	return c.sendCommand(ctx, streamCommand{
		Command: "stopTradeStatus",
	})
}

func (c *StreamClient) Ping(ctx context.Context) error {

	return c.sendCommand(ctx, streamCommand{
		Command:         "ping",
		StreamSessionId: c.sessionId,
	})
}

func (c *StreamClient) Listen(ctx context.Context) error {

	commChan := make(chan goCommChan, c.opts.IncommingBufferSize)

	ctx, c.listenCtxCancel = context.WithCancel(ctx)
	defer c.listenCtxCancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				msg, err := c.read(ctx)
				if err != nil {
					commChan <- goCommChan{nil, err}
					return
				}
				commChan <- goCommChan{msg, err}
			}
		}
	}()

	ticker := time.NewTicker(c.opts.KeepAliveInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			c.Ping(ctx)
		case resp := <-commChan:
			if resp.err == nil {
				if err := c.handleMessage(resp.data.([]byte)); err != nil {
					return err
				}
				continue
			}
			return resp.err
		default:
			time.Sleep(c.opts.PollingInterval)
		}
	}
}

func (c *StreamClient) sendCommand(ctx context.Context, cmd streamCommand) error {

	c.mu.Lock()
	defer c.mu.Unlock()

	data, err := json.Marshal(cmd)
	if err != nil {
		return fmt.Errorf("unable to marshal %s command: %w", cmd.Command, err)
	}

	ctx, ctxCancel := context.WithTimeout(ctx, c.opts.WriteTimeout)
	defer ctxCancel()

	if err := c.write(ctx, data); err != nil {
		return fmt.Errorf("unable to send %s command: %w", cmd.Command, err)
	}

	return nil
}

func (c *StreamClient) handleMessage(msg []byte) error {

	var s streamData
	if err := json.Unmarshal(msg, &s); err != nil {
		return fmt.Errorf("failed to unmarshal message - %s: %w", msg, err)
	}

	switch s.Command {
	case "balance":
		return c.handleBalance(s)
	case "candle":
		return c.handleCandle(s)
	case "keepAlive":
		return c.handleKeepAlive(s)
	case "news":
		return c.handleNews(s)
	case "profit":
		return c.handleProfits(s)
	case "tickPrices":
		return c.handleTickPrices(s)
	case "trade":
		return c.handleTrades(s)
	case "tradeStatus":
		return c.handleTradeStatus(s)
	default:
		return fmt.Errorf("invalid command recieved %s in %s", s.Command, msg)
	}
}

func (c *StreamClient) handleBalance(s streamData) error {

	if c.balanceCb != nil {
		var balance Balance
		if err := unmarshalRecord(s, &balance); err != nil {
			return fmt.Errorf("failed to handle balance message: %w", err)
		}
		c.balanceCb(balance)
	}

	return nil
}

func (c *StreamClient) handleCandle(s streamData) error {

	if c.candlesCb != nil {
		var candle Candle
		if err := unmarshalRecord(s, &candle); err != nil {
			return fmt.Errorf("failed to handle candle message: %w", err)
		}
		if cb, exists := c.candlesCb[candle.Symbol]; exists && cb != nil {
			cb(candle)
		}
	}

	return nil
}

func (c *StreamClient) handleKeepAlive(s streamData) error {

	if c.keepAliveCb != nil {
		var keepAlive KeepAlive
		if err := unmarshalRecord(s, &keepAlive); err != nil {
			return fmt.Errorf("failed to handle keepAlive message: %w", err)
		}
		c.keepAliveCb(keepAlive)
	}

	return nil
}

func (c *StreamClient) handleNews(s streamData) error {

	if c.newsCb != nil {
		var news News
		if err := unmarshalRecord(s, &news); err != nil {
			return fmt.Errorf("failed to handle news message: %w", err)
		}
		c.newsCb(news)
	}

	return nil
}

func (c *StreamClient) handleProfits(s streamData) error {

	if c.profitsCb != nil {
		var profit Profit
		if err := unmarshalRecord(s, &profit); err != nil {
			return fmt.Errorf("failed to handle profit message: %w", err)
		}
		c.profitsCb(profit)
	}

	return nil
}

func (c *StreamClient) handleTickPrices(s streamData) error {

	if c.tickPricesCb != nil {
		var tickPrice TickPrice
		if err := unmarshalRecord(s, &tickPrice); err != nil {
			return fmt.Errorf("failed to handle tickPrice message: %w", err)
		}
		if cb, exists := c.tickPricesCb[tickPrice.Symbol]; exists && cb != nil {
			cb(tickPrice)
		}
	}

	return nil
}

func (c *StreamClient) handleTrades(s streamData) error {

	if c.tradesCb != nil {
		var trade Trade
		if err := unmarshalRecord(s, &trade); err != nil {
			return fmt.Errorf("failed to handle trade message: %w", err)
		}
		c.tradesCb(trade)
	}

	return nil
}

func (c *StreamClient) handleTradeStatus(s streamData) error {

	if c.tradeStatusCb != nil {
		var tradeStatus TradeStatus
		if err := unmarshalRecord(s, &tradeStatus); err != nil {
			return fmt.Errorf("failed to handle tradeStatus message: %w", err)
		}
		c.tradeStatusCb(tradeStatus)
	}

	return nil
}

func unmarshalRecord(s streamData, rec interface{}) error {

	if err := json.Unmarshal(s.Data, &rec); err != nil {
		return fmt.Errorf("failed to unmarshal %s - %s: %w", s.Command, s.Data, err)
	}

	return nil
}
