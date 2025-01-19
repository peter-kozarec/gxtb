package goxtb

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

func (c *ApiClient) Connect(ctx context.Context) (err error) {
	c.conn, err = Dial(ctx, c.url)
	return err
}

func (c *ApiClient) Disconnect() error {
	return c.conn.Close()
}

func (c *ApiClient) Login(ctx context.Context, r LoginRequest) (sessionId string, err error) {

	request, err := createLoginRequest(r)
	if err != nil {
		return "", fmt.Errorf("")
	}

	if err := c.conn.Write(ctx, request); err != nil {
		return "", fmt.Errorf("login request: %w", err)
	}

	response, err := c.conn.Read(ctx)
	if err != nil {
		return "", fmt.Errorf("login response: %w", err)
	}

	return parseLoginResponse(response)
}

func createLoginRequest(r LoginRequest) ([]byte, error) {

	arguments, err := json.Marshal(r)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to marshal login request arguments: %w", err)
	}

	request := ApiRequest{Command: "login", Arguments: arguments}
	requestData, err := json.Marshal(request)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to marshal API request: %w", err)
	}
	return requestData, nil
}

func parseLoginResponse(responseData []byte) (sessionId string, err error) {

	var r LoginResponse
	if err := json.Unmarshal(responseData, &r); err != nil {
		return "", fmt.Errorf("failed to unmarshal login response: %w", err)
	}

	if !r.Status {
		return "", fmt.Errorf("login failed with response: %s", string(responseData))
	}

	return r.StreamSessionId, nil
}
