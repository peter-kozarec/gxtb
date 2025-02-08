package gxtb

import (
	"context"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
)

type websocketConnection struct {
	ws *websocket.Conn
}

func (c *websocketConnection) connect(ctx context.Context, url url.URL) error {

	var err error

	c.ws, _, err = websocket.DefaultDialer.DialContext(ctx, url.String(), nil)
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
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("read canceled: %w", ctx.Err())
	case resp := <-commChan:
		return resp.data.([]byte), resp.err
	}
}
