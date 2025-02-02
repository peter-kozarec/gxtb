package gxtb

import (
	"net/url"
	"time"
)

type ApiPath string

const (
	RealApi ApiPath = "/real"
	DemoApi ApiPath = "/demo"
)

type ApiOptions struct {
	EndpointPath      ApiPath
	ApiCallTimeout    time.Duration
	KeepAliveInterval time.Duration
	PollingInterval   time.Duration
}

func (o ApiOptions) GetUrl() url.URL {
	return url.URL{Scheme: "wss", Host: "ws.xtb.com", Path: string(o.EndpointPath)}
}

func DefaultApiOptions() ApiOptions {
	return ApiOptions{
		EndpointPath:      RealApi,
		ApiCallTimeout:    time.Millisecond * 250,
		KeepAliveInterval: time.Second * 10,
		PollingInterval:   time.Millisecond * 10,
	}
}

func DefaultDemoApiOptions() ApiOptions {
	return ApiOptions{
		EndpointPath:      DemoApi,
		ApiCallTimeout:    time.Millisecond * 250,
		KeepAliveInterval: time.Second * 10,
		PollingInterval:   time.Millisecond * 10,
	}
}
