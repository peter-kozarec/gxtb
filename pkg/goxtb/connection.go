package goxtb

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

type Conn struct {
	conn *websocket.Conn
}

func Dial(ctx context.Context, url string) (*Conn, error) {
	c := new(Conn)
	var err error
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)
		c.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err != nil {
			errChan <- fmt.Errorf("websocket failed to connect: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		return c, ctx.Err()
	case err := <-errChan:
		return nil, err
	}
}

func (c *Conn) Close() error {
	if err := c.conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}
	return nil
}

func (c *Conn) Write(ctx context.Context, data []byte) error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)
		if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			errChan <- fmt.Errorf("failed to write: %w", err)
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func (c *Conn) Read(ctx context.Context) ([]byte, error) {
	dataChan := make(chan []byte, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(dataChan)
		defer close(errChan)

		_, data, err := c.conn.ReadMessage()
		if err != nil {
			errChan <- fmt.Errorf("failed to read: %w", err)
			return
		}

		dataChan <- data
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errChan:
		return nil, err
	case data := <-dataChan:
		return data, nil
	}
}
