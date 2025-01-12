package goxtb

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	wsConn *websocket.Conn
	wsUrl  string
	mutex  sync.Mutex
}

func NewClient() *Client {
	return &Client{
		wsConn: nil,
		wsUrl:  "wss://ws.xtb.com/real",
	}
}

func NewDemoClient() *Client {
	return &Client{
		wsConn: nil,
		wsUrl:  "wss://ws.xtb.com/demo",
	}
}

func (client *Client) Connect() error {
	var err error
	client.wsConn, _, err = websocket.DefaultDialer.Dial(client.wsUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	return nil
}

func (client *Client) Close() {
	client.wsConn.Close()
}

func (c *Client) Login(r LoginRequest) (sessionId string, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.wsConn == nil {
		return "", fmt.Errorf("WebSocket connection is not initialized")
	}

	data, err := createLoginRequest(r)
	if err != nil {
		return "", fmt.Errorf("")
	}

	if err := c.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
		return "", fmt.Errorf("failed to write login request to WebSocket: %w", err)
	}

	_, responseData, err := c.wsConn.ReadMessage()
	if err != nil {
		return "", fmt.Errorf("failed to read login response from WebSocket: %w", err)
	}

	return parseLoginResponse(responseData)
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
