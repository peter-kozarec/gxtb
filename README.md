# gxtb

[![Go](https://github.com/peter-kozarec/gxtb/actions/workflows/go.yml/badge.svg)](https://github.com/peter-kozarec/gxtb/actions/workflows/go.yml)

`gxtb` is a Go package that provides API and streaming clients for interacting with the XTB trading platform.

## Installation

To install the package, run:

```sh
go get -u github.com/peter-kozarec/gxtb@latest
```

## Usage

### API Client Example

The following example demonstrates how to use the API client to connect, authenticate, and retrieve server details.

```go
package main

import (
	"context"
	"log"
	"os"
	"github.com/peter-kozarec/gxtb"
)

func main() {
	c := gxtb.NewApiClient(gxtb.DefaultDemoApiOptions())
	ctx := context.Background()

	// Connect to API client
	if err := c.Connect(ctx); err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer c.Disconnect()

	userId := os.Getenv("XTB_UserId")
	password := os.Getenv("XTB_Password")

	// Login to API client
	if _, err := c.Login(ctx, userId, password, "testApp"); err != nil {
		log.Fatalf("unable to login: %v", err)
	}
	defer c.Logout(ctx)

	// Get API version
	version, err := c.GetVersion(ctx)
	if err != nil {
		log.Fatalf("unable to retrieve API version: %v", err)
	}
	log.Printf("API version: %s\n", version)

	// Get broker server time
	serverTime, err := c.GetServerTime(ctx)
	if err != nil {
		log.Fatalf("unable to retrieve server time: %v", err)
	}
	log.Printf("Server time: %s\n", serverTime.TimeString)
}
```

### Streaming Client Example

This example demonstrates how to connect to the streaming client and subscribe to real-time market updates.

```go
package main

import (
	"context"
	"log"
	"os"
	"github.com/peter-kozarec/gxtb"
)

func tUpdates(tickPrice gxtb.TickPrice) {
	log.Printf("Tick update - %v\n", tickPrice)
}

func main() {
	apiClient := gxtb.NewApiClient(gxtb.DefaultDemoApiOptions())
	streamClient := gxtb.NewStreamClient(gxtb.DefaultDemoStreamOptions())
	ctx := context.Background()

	// Connect to API client
	if err := apiClient.Connect(ctx); err != nil {
		log.Fatalf("unable to connect API client: %v", err)
	}
	defer apiClient.Disconnect()

	// Connect to stream client
	if err := streamClient.Connect(ctx); err != nil {
		log.Fatalf("unable to connect stream client: %v", err)
	}
	defer streamClient.Disconnect()

	userId := os.Getenv("XTB_UserId")
	password := os.Getenv("XTB_Password")

	// Login and retrieve stream session ID
	streamSessionId, err := apiClient.Login(ctx, userId, password, "testApp")
	if err != nil {
		log.Fatalf("unable to login: %v", err)
	}
	defer apiClient.Logout(ctx)

	// Set stream session ID
	streamClient.SetSessionId(streamSessionId)

	// Subscribe to tick price updates
	if err := streamClient.GetTickPrices(ctx, "BITCOIN", 100, 1, tUpdates); err != nil {
		log.Fatalf("unable to subscribe to tick updates: %v", err)
	}

	// Listen for updates
	if err := streamClient.Listen(ctx); err != nil {
		log.Fatalf("unexpected error: %v", err)
	}

	log.Print("Graceful exit")
}
```

## License

This project is licensed under the MIT License.

