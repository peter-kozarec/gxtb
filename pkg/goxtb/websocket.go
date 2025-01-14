package goxtb

import (
	"context"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type wsConn interface {
	connect(url string) error
	disconnect() error
	write(ctx context.Context, data []byte) error
	read(ctx context.Context) ([]byte, error)
}

type wsConnImpl struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func (w *wsConnImpl) connect(url string) error {
	var err error
	w.conn, _, err = websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return fmt.Errorf("websocket failed to connect: %w", err)
	}
	return nil
}

func (w *wsConnImpl) disconnect() error {
	if err := w.conn.Close(); err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}
	return nil
}

func (w *wsConnImpl) write(ctx context.Context, data []byte) error {
	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)

		w.mu.Lock()
		defer w.mu.Unlock()

		if err := w.conn.WriteMessage(websocket.TextMessage, data); err != nil {
			errChan <- fmt.Errorf("failed to write: %w", err)
			return
		}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errChan:
		return err
	}
}

func (w *wsConnImpl) read(ctx context.Context) ([]byte, error) {
	dataChan := make(chan []byte, 1)
	errChan := make(chan error, 1)

	go func() {
		defer close(dataChan)
		defer close(errChan)

		w.mu.Lock()
		defer w.mu.Unlock()

		_, data, err := w.conn.ReadMessage()
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
