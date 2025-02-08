package gxtb

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type websocketConnection struct {
	ws     *websocket.Conn
	url    string
	logger *zap.Logger
}

func (c *websocketConnection) AttachLogger(logger *zap.Logger) {
	c.logger = logger
}

func (c *websocketConnection) connect(ctx context.Context, url url.URL) error {

	var err error

	c.url = url.String()
	c.ws, _, err = websocket.DefaultDialer.DialContext(ctx, c.url, nil)
	if err != nil {
		return fmt.Errorf("unable to dial %v: %w", url, err)
	}

	return nil
}

func (c *websocketConnection) disconnect() error {

	return c.ws.Close()
}

func (c *websocketConnection) write(ctx context.Context, data []byte) error {

	commChan := make(chan goCommChan)

	go func() {
		err := c.ws.WriteMessage(websocket.TextMessage, data)
		commChan <- goCommChan{nil, err}

		if err == nil && c.logger != nil {
			c.logger.Debug(c.url, zap.String("write", string(data)))
		}
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("write canceled: %w", ctx.Err())
	case resp := <-commChan:
		return resp.err
	}
}

func (c *websocketConnection) read(ctx context.Context) ([]byte, error) {

	commChan := make(chan goCommChan)

	go func() {
		_, p, err := c.ws.ReadMessage()
		commChan <- goCommChan{p, err}

		if err == nil && c.logger != nil {
			c.logger.Debug(c.url, zap.String("read", string(p)))
		}
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("read canceled: %w", ctx.Err())
	case resp := <-commChan:
		return resp.data.([]byte), resp.err
	}
}
