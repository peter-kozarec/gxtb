package gxtb

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type apiCommand struct {
	Command   string      `json:"command"`
	Arguments interface{} `json:"arguments,omitempty"`
}

type extendedArguments struct {
	Info           interface{} `json:"info,omitempty"`
	TradeTransInfo interface{} `json:"tradeTransInfo,omitempty"`
}

type apiResponse struct {
	Status          bool            `json:"status"`
	ReturnData      json.RawMessage `json:"returnData,omitempty"`
	StreamSessionId string          `json:"streamSessionId,omitempty"`
	ErrorCode       string          `json:"errorCode,omitempty"`
	ErrorDescr      string          `json:"errorDescr,omitempty"`
}

type ApiClient struct {
	websocketConnection

	opts          ApiOptions
	sessionId     string
	mu            sync.Mutex
	keepAliveCncl context.CancelFunc
}

func NewApiClient(opts ApiOptions) *ApiClient {

	return &ApiClient{
		opts: opts,
	}
}

func (c *ApiClient) Connect(ctx context.Context) error {

	if err := c.connect(ctx, c.opts.GetUrl()); err != nil {
		return err
	}

	return nil
}

func (c *ApiClient) Disconnect() error {

	if c.keepAliveCncl != nil {
		c.keepAliveCncl()
	}

	return c.disconnect()
}

func (c *ApiClient) Login(ctx context.Context, userId, password, appName string) (string, error) {

	args := struct {
		UserId   string `json:"userId"`
		Password string `json:"password"`
		AppName  string `json:"appName"`
	}{userId, password, appName}

	resp, err := c.sendRecieve(ctx, apiCommand{Command: "login", Arguments: args})
	if err != nil {
		return "", fmt.Errorf("unable to process login api call: %w", err)
	}

	ctx, c.keepAliveCncl = context.WithCancel(ctx)

	// Orphan goroutine to refresh connection, until canceled
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(c.opts.KeepAliveInterval):
				c.Ping(ctx)
			default:
				time.Sleep(c.opts.PollingInterval)
			}
		}
	}()

	c.sessionId = resp.StreamSessionId
	return resp.StreamSessionId, nil
}

func (c *ApiClient) Logout(ctx context.Context) error {

	_, err := c.sendRecieve(ctx, apiCommand{Command: "logout"})
	if err != nil {
		return fmt.Errorf("unable to process logout api call: %w", err)
	}

	c.keepAliveCncl()

	return nil
}

func (c *ApiClient) Ping(ctx context.Context) error {

	_, err := c.sendRecieve(ctx, apiCommand{Command: "ping"})
	if err != nil {
		return fmt.Errorf("unable to process ping api call: %w", err)
	}

	return nil
}

func (c *ApiClient) GetAllSymbols(ctx context.Context) ([]SymbolInfo, error) {

	var records []SymbolInfo

	resp, err := c.sendRecieve(ctx, apiCommand{Command: "getAllSymbols"})
	if err != nil {
		return records, fmt.Errorf("unable to process getAllSymbols api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &records); err != nil {
		return records, fmt.Errorf("unable to unmarshal getAllSymbols response: %w", err)
	}

	return records, nil
}

func (c *ApiClient) GetCalendar(ctx context.Context) ([]Calendar, error) {

	var records []Calendar

	resp, err := c.sendRecieve(ctx, apiCommand{Command: "getCalendar"})
	if err != nil {
		return records, fmt.Errorf("unable to process getCalendar api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &records); err != nil {
		return records, fmt.Errorf("unable to unmarshal getCalendar response: %w", err)
	}

	return records, nil
}

func (c *ApiClient) GetChartLastRequest(ctx context.Context, info ChartLastInfo) (ChartData, error) {

	var chartData ChartData

	resp, err := c.sendRecieve(ctx, apiCommand{"getChartLastRequest", extendedArguments{Info: info}})
	if err != nil {
		return chartData, fmt.Errorf("unable to process getChartLastRequest api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &chartData); err != nil {
		return chartData, fmt.Errorf("unable to unmarshal getChartLastRequest response: %w", err)
	}

	return chartData, nil
}

func (c *ApiClient) GetChartRangeRequest(ctx context.Context, info ChartRangeInfo) (ChartData, error) {

	var chartData ChartData

	resp, err := c.sendRecieve(ctx, apiCommand{"getChartRangeRequest", extendedArguments{Info: info}})
	if err != nil {
		return chartData, fmt.Errorf("unable to process getChartRangeRequest api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &chartData); err != nil {
		return chartData, fmt.Errorf("unable to unmarshal getChartRangeRequest response: %w", err)
	}

	return chartData, nil
}

func (c *ApiClient) GetCommissionDef(ctx context.Context, symbol string, volume float32) (CommissionData, error) {

	args := struct {
		Symbol string  `json:"symbol"`
		Volume float32 `json:"volume"`
	}{symbol, volume}

	var commission CommissionData

	resp, err := c.sendRecieve(ctx, apiCommand{"getCommissionDef", args})
	if err != nil {
		return commission, fmt.Errorf("unable to process GetCommissionDef api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &commission); err != nil {
		return commission, fmt.Errorf("unable to unmarshal GetCommissionDef response: %w", err)
	}

	return commission, nil
}

func (c *ApiClient) GetCurrentUserData(ctx context.Context) (UserData, error) {

	var userData UserData

	resp, err := c.sendRecieve(ctx, apiCommand{Command: "getCurrentUserData"})
	if err != nil {
		return userData, fmt.Errorf("unable to process getCurrentUserData api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &userData); err != nil {
		return userData, fmt.Errorf("unable to unmarshal getCurrentUserData response: %w", err)
	}

	return userData, nil
}

func (c *ApiClient) GetMarginLevel(ctx context.Context) (MarginData, error) {

	var marginData MarginData

	resp, err := c.sendRecieve(ctx, apiCommand{Command: "getMarginLevel"})
	if err != nil {
		return marginData, fmt.Errorf("unable to process getMarginLevel api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &marginData); err != nil {
		return marginData, fmt.Errorf("unable to unmarshal getMarginLevel response: %w", err)
	}

	return marginData, nil
}

func (c *ApiClient) GetMarginTrade(ctx context.Context, symbol string, volume float32) (float32, error) {

	args := struct {
		Symbol string  `json:"symbol"`
		Volume float32 `json:"volume"`
	}{symbol, volume}

	marginData := struct {
		Margin float32 `json:"margin"`
	}{}

	resp, err := c.sendRecieve(ctx, apiCommand{"getMarginTrade", args})
	if err != nil {
		return marginData.Margin, fmt.Errorf("unable to process getMarginTrade api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &marginData); err != nil {
		return marginData.Margin, fmt.Errorf("unable to unmarshal getMarginTrade response: %w", err)
	}

	return marginData.Margin, nil
}

func (c *ApiClient) GetNews(ctx context.Context, end, start int) ([]NewsTopic, error) {

	args := struct {
		End   int `json:"end"`
		Start int `json:"start"`
	}{end, start}

	var news []NewsTopic

	resp, err := c.sendRecieve(ctx, apiCommand{"getNews", args})
	if err != nil {
		return news, fmt.Errorf("unable to process getNews api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &news); err != nil {
		return news, fmt.Errorf("unable to unmarshal getNews response: %w", err)
	}

	return news, nil
}

func (c *ApiClient) GetServerTime(ctx context.Context) (ServerTime, error) {

	var serverTime ServerTime

	resp, err := c.sendRecieve(ctx, apiCommand{Command: "getServerTime"})
	if err != nil {
		return serverTime, fmt.Errorf("unable to process getServerTime api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &serverTime); err != nil {
		return serverTime, fmt.Errorf("unable to unmarshal getServerTime response: %w", err)
	}

	return serverTime, nil
}

func (c *ApiClient) GetSymbol(ctx context.Context, symbol string) (SymbolInfo, error) {

	args := struct {
		Symbol string `json:"symbol"`
	}{symbol}

	var symbolInfo SymbolInfo

	resp, err := c.sendRecieve(ctx, apiCommand{"getSymbol", args})
	if err != nil {
		return symbolInfo, fmt.Errorf("unable to process getSymbol api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &symbolInfo); err != nil {
		return symbolInfo, fmt.Errorf("unable to unmarshal getSymbol response: %w", err)
	}

	return symbolInfo, nil
}

func (c *ApiClient) GetTickPrices(ctx context.Context, symbols []string, level, ts int) ([]TickRecord, error) {

	args := struct {
		Level     int      `json:"level"`
		Symbols   []string `json:"symbols"`
		Timestamp int      `json:"timestamp"`
	}{level, symbols, ts}

	tickRecordData := struct {
		Quotations []TickRecord `json:"quotations"`
	}{}

	resp, err := c.sendRecieve(ctx, apiCommand{"getTickPrices", args})
	if err != nil {
		return tickRecordData.Quotations, fmt.Errorf("unable to process getTickPrices api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &tickRecordData); err != nil {
		return tickRecordData.Quotations, fmt.Errorf("unable to unmarshal getTickPrices response: %w", err)
	}

	return tickRecordData.Quotations, nil
}

func (c *ApiClient) GetTradeRecords(ctx context.Context, orders []int) ([]TradeRecord, error) {

	args := struct {
		Orders []int `json:"orders"`
	}{orders}

	var tradeRecords []TradeRecord

	resp, err := c.sendRecieve(ctx, apiCommand{"getTradeRecords", args})
	if err != nil {
		return tradeRecords, fmt.Errorf("unable to process getTradeRecords api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &tradeRecords); err != nil {
		return tradeRecords, fmt.Errorf("unable to unmarshal getTradeRecords response: %w", err)
	}

	return tradeRecords, nil
}

func (c *ApiClient) GetTrades(ctx context.Context, openedOnly bool) ([]TradeRecord, error) {

	args := struct {
		OpenedOnly bool `json:"openedOnly"`
	}{openedOnly}

	var tradeRecords []TradeRecord

	resp, err := c.sendRecieve(ctx, apiCommand{"getTrades", args})
	if err != nil {
		return tradeRecords, fmt.Errorf("unable to process getTrades api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &tradeRecords); err != nil {
		return tradeRecords, fmt.Errorf("unable to unmarshal getTrades response: %w", err)
	}

	return tradeRecords, nil
}

func (c *ApiClient) GetTradesHistory(ctx context.Context, end, start int) ([]TradeRecord, error) {

	args := struct {
		End   int `json:"end"`
		Start int `json:"start"`
	}{end, start}

	var tradeRecords []TradeRecord

	resp, err := c.sendRecieve(ctx, apiCommand{"getTradesHistory", args})
	if err != nil {
		return tradeRecords, fmt.Errorf("unable to process getTradesHistory api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &tradeRecords); err != nil {
		return tradeRecords, fmt.Errorf("unable to unmarshal getTradesHistory response: %w", err)
	}

	return tradeRecords, nil
}

func (c *ApiClient) GetVersion(ctx context.Context) (string, error) {

	versionData := struct {
		Version string `json:"version"`
	}{}

	resp, err := c.sendRecieve(ctx, apiCommand{Command: "getVersion"})
	if err != nil {
		return versionData.Version, fmt.Errorf("unable to process getVersion api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &versionData); err != nil {
		return versionData.Version, fmt.Errorf("unable to unmarshal getVersion response: %w", err)
	}

	return versionData.Version, nil
}

func (c *ApiClient) TradeTransaction(ctx context.Context, txnInfo TransactionInfo) (OrderId, error) {

	var orderId OrderId

	resp, err := c.sendRecieve(ctx, apiCommand{"tradeTransaction", extendedArguments{TradeTransInfo: txnInfo}})
	if err != nil {
		return orderId, fmt.Errorf("unable to process tradeTransaction api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &orderId); err != nil {
		return orderId, fmt.Errorf("unable to unmarshal tradeTransaction response: %w", err)
	}

	return orderId, nil
}

func (c *ApiClient) TradeTransactionStatus(ctx context.Context, orderId int) (TransactionStatus, error) {

	args := struct {
		Order int `json:"order"`
	}{orderId}

	var txnStatus TransactionStatus

	resp, err := c.sendRecieve(ctx, apiCommand{"tradeTransactionStatus", args})
	if err != nil {
		return txnStatus, fmt.Errorf("unable to process tradeTransactionStatus api call: %w", err)
	}

	if err := json.Unmarshal(resp.ReturnData, &txnStatus); err != nil {
		return txnStatus, fmt.Errorf("unable to unmarshal tradeTransactionStatus response: %w", err)
	}

	return txnStatus, nil
}

func (c *ApiClient) sendRecieve(ctx context.Context, cmd apiCommand) (apiResponse, error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	var resp apiResponse

	req, err := json.Marshal(cmd)
	if err != nil {
		return resp, fmt.Errorf("failed to marshal %v: %w", cmd, err)
	}

	ctx, ctxCancel := context.WithTimeout(ctx, c.opts.ApiCallTimeout)
	defer ctxCancel()

	if err := c.write(ctx, req); err != nil {
		return resp, fmt.Errorf("failed to send %v: %w", req, err)
	}

	respData, err := c.read(ctx)
	if err != nil {
		return resp, fmt.Errorf("failed to read: %w", err)
	}

	if err := json.Unmarshal(respData, &resp); err != nil {
		return resp, fmt.Errorf("failed to unmarshal %s: %w", respData, err)
	}

	if !resp.Status {
		return resp, fmt.Errorf("%s - %s", resp.ErrorCode, resp.ErrorDescr)
	}

	return resp, nil
}
