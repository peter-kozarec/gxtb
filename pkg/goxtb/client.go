package goxtb

type Client struct {
	Stream *StreamClient
}

func NewClient() *Client {
	return &Client{
		Stream: NewStreamClient(),
	}
}

func NewDemoClient() *Client {
	return &Client{
		Stream: NewStreamDemoClient(),
	}
}
