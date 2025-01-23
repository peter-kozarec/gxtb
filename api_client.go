package gxtb

import (
	"context"
	"encoding/json"
	"fmt"
)

type ApiClient struct {
	conn *Conn
	url  string
}

func NewApiClient() *ApiClient {
	return &ApiClient{
		url: "wss://ws.xtb.com/real",
	}
}

func NewApiDemoClient() *ApiClient {
	return &ApiClient{
		url: "wss://ws.xtb.com/demo",
	}
}

func (c *ApiClient) Connect(ctx context.Context) error {
	var err error
	c.conn, err = Dial(ctx, c.url)
	if err != nil {
		return fmt.Errorf("failed to connect to API server at %s: %w", c.url, err)
	}
	return nil
}

func (c *ApiClient) Disconnect() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("failed to disconnect from API server: %w", err)
	}
	return nil
}

func (c *ApiClient) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {

	var loginResponse LoginResponse

	requestData, err := marshalRequest("login", req)
	if err != nil {
		return loginResponse, fmt.Errorf("failed to serialize API request for login: %w", err)
	}

	responseData, err := c.conn.WriteRead(ctx, requestData)
	if err != nil {
		return loginResponse, fmt.Errorf("failed to read login response: %w", err)
	}

	if err := json.Unmarshal(responseData, &loginResponse); err != nil {
		return loginResponse, fmt.Errorf("failed to parse login response JSON: %w", err)
	}

	return loginResponse, nil
}

func (c *ApiClient) Logout(ctx context.Context) error {

	requestData, err := marshalRequest("logout", nil)
	if err != nil {
		return fmt.Errorf("failed to serialize API request for logout: %w", err)
	}

	responseData, err := c.conn.WriteRead(ctx, requestData)
	if err != nil {
		return fmt.Errorf("failed to read logout response: %w", err)
	}

	if err := unmarshalResponse(responseData, nil); err != nil {
		return fmt.Errorf("unable to unmarshal logout response: %w", err)
	}

	return nil

}

func (c *ApiClient) GetAllSymbols(ctx context.Context) ([]SymbolRecord, error) {

	var symbolRecords []SymbolRecord

	requestData, err := marshalRequest("getAllSymbols", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetAllSymbols request: %w", err)
	}

	responseData, err := c.conn.WriteRead(ctx, requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GetAllSymbols: %w", err)
	}

	if err := unmarshalResponse(responseData, &symbolRecords); err != nil {
		return nil, fmt.Errorf("unable to unmarshal GetAllSymbols response: %w", err)
	}

	return symbolRecords, nil
}

func (c *ApiClient) GetCalendar(ctx context.Context) ([]CalendarRecord, error) {

	var calendarRecords []CalendarRecord

	requestData, err := marshalRequest("getCalendar", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetCalendar request: %w", err)
	}

	responseData, err := c.conn.WriteRead(ctx, requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GetCalendar: %w", err)
	}

	if err := unmarshalResponse(responseData, &calendarRecords); err != nil {
		return nil, fmt.Errorf("unable to unmarshal GetCalendar response: %w", err)
	}

	return calendarRecords, nil
}

func (c *ApiClient) GetChartLastRequest(ctx context.Context, req ChartLastRequest) ([]RateInfoRecord, error) {

	var rates []RateInfoRecord

	requestData, err := marshalRequest("getChartLastRequest", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetChartLastRequest request: %w", err)
	}

	responseData, err := c.conn.WriteRead(ctx, requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GetChartLastRequest: %w", err)
	}

	if err := unmarshalResponse(responseData, &rates); err != nil {
		return nil, fmt.Errorf("unable to unmarshal GetChartLastRequest response: %w", err)
	}

	return rates, nil
}

func (c *ApiClient) GetChartRangeRequest(ctx context.Context, req ChartRangeRequest) ([]RateInfoRecord, error) {

	var rates []RateInfoRecord

	requestData, err := marshalRequest("getChartRangeRequest", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GetChartRangeRequest request: %w", err)
	}

	responseData, err := c.conn.WriteRead(ctx, requestData)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GetChartRangeRequest: %w", err)
	}

	if err := unmarshalResponse(responseData, &rates); err != nil {
		return nil, fmt.Errorf("unable to unmarshal GetChartRangeRequest response: %w", err)
	}

	return rates, nil
}

func (c *ApiClient) GetCommissionDef(ctx context.Context, req CommissionRequest) (CommissionData, error) {

	var commissionData CommissionData

	requestData, err := marshalRequest("getCommissionDef", nil)
	if err != nil {
		return commissionData, fmt.Errorf("failed to marshal GetCommissionDef request: %w", err)
	}

	responseData, err := c.conn.WriteRead(ctx, requestData)
	if err != nil {
		return commissionData, fmt.Errorf("failed to perform GetCommissionDef: %w", err)
	}

	if err := unmarshalResponse(responseData, &commissionData); err != nil {
		return commissionData, fmt.Errorf("unable to unmarshal GetCommissionDef response: %w", err)
	}

	return commissionData, nil
}

func (c *ApiClient) GetCurrentUserData(ctx context.Context) (UserData, error) {

	var userData UserData

	requestData, err := marshalRequest("getCurrentUserData", nil)
	if err != nil {
		return userData, fmt.Errorf("failed to marshal GetCurrentUserData request: %w", err)
	}

	responseData, err := c.conn.WriteRead(ctx, requestData)
	if err != nil {
		return userData, fmt.Errorf("failed to perform GetCurrentUserData: %w", err)
	}

	if err := unmarshalResponse(responseData, &userData); err != nil {
		return userData, fmt.Errorf("unable to unmarshal GetCurrentUserData response: %w", err)
	}

	return userData, nil
}

func (c *ApiClient) GetMarginLevel(ctx context.Context) (MarginData, error) {

	var marginData MarginData

	requestData, err := marshalRequest("getMarginLevel", nil)
	if err != nil {
		return marginData, fmt.Errorf("failed to marshal GetMarginLevel request: %w", err)
	}

	responseData, err := c.conn.WriteRead(ctx, requestData)
	if err != nil {
		return marginData, fmt.Errorf("failed to perform GetMarginLevel: %w", err)
	}

	if err := unmarshalResponse(responseData, &marginData); err != nil {
		return marginData, fmt.Errorf("unable to unmarshal GetMarginLevel response: %w", err)
	}

	return marginData, nil
}
