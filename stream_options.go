package gxtb

import (
	"net/url"
	"time"
)

type StreamPath string

const (
	RealStream StreamPath = "/realStream"
	DemoStream StreamPath = "/demoStream"
)

type StreamOptions struct {
	EndpointPath        StreamPath
	WriteTimeout        time.Duration // Timeout for the websocket write operation
	KeepAliveInterval   time.Duration // Interval for sending keep-alive pings
	IncommingBufferSize int           // Size of the channel for incoming messages
	PollingInterval     time.Duration // Frequency of polling operations
}

func (o StreamOptions) GetUrl() url.URL {
	return url.URL{Scheme: "wss", Host: "ws.xtb.com", Path: string(o.EndpointPath)}
}

func DefaultStreamOptions() StreamOptions {
	return StreamOptions{
		EndpointPath:        RealStream,
		WriteTimeout:        time.Millisecond * 500,
		KeepAliveInterval:   time.Second * 10,
		IncommingBufferSize: 10,
		PollingInterval:     time.Millisecond * 10,
	}
}

func DefaultDemoStreamOptions() StreamOptions {
	return StreamOptions{
		EndpointPath:        DemoStream,
		WriteTimeout:        time.Millisecond * 500,
		KeepAliveInterval:   time.Second * 10,
		IncommingBufferSize: 10,
		PollingInterval:     time.Millisecond * 10,
	}
}
