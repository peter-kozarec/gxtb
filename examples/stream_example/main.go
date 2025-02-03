package main

import (
	"context"
	"log"
	"os"
	
	"github.com/peter-kozarec/gxtb"
)

func tickPriceUpdate(tickPrice gxtb.TickPrice) {
	log.Printf("Tick update - %v\n", tickPrice)
}

func keepAliveUpdate(keepAlive gxtb.KeepAlive) {
	log.Printf("Keep alive update - %v\n", keepAlive)
}

func main() {
	apiClient := gxtb.NewApiClient(gxtb.DefaultDemoApiOptions())
	streamClient := gxtb.NewStreamClient(gxtb.DefaultDemoStreamOptions())
	ctx := context.Background()

	// Connect to api client
	if err := apiClient.Connect(ctx); err != nil {
		log.Fatalf("unable to connect api client: %v", err)
	}
	defer apiClient.Disconnect()

	// Connect to stream client
	if err := streamClient.Connect(ctx); err != nil {
		log.Fatalf("unable to connect stream client: %v", err)
	}
	defer streamClient.Disconnect()

	// Login and retrieve stream session id
	streamSessionId, err := apiClient.Login(ctx, os.Getenv("XTB_UserId"), os.Getenv("XTB_Password"), "testApp")
	if err != nil {
		log.Fatalf("unable to login: %v", err)
	}
	defer apiClient.Logout(ctx)

	// Set stream session id
	streamClient.SetSessionId(streamSessionId)

	// Subscribe to keep alive updates
	if err := streamClient.GetKeepAlive(ctx, keepAliveUpdate); err != nil {
		log.Fatalf("unable to subscribe to keep alive updates: %v", err)
	}

	// Subscribe to tick price updates
	if err := streamClient.GetTickPrices(ctx, "BITCOIN", 100, 1, tickPriceUpdate); err != nil {
		log.Fatalf("unable to subscribe to tick updates: %v", err)
	}

	// Create a timer for 30 seconds to cancel the listen executed bellow
	// ctx, ctxCancel := context.WithTimeout(ctx, time.Second*30)
	// defer ctxCancel()

	// Listen forever or until context is canceled
	if err := streamClient.Listen(ctx); err != nil {
		log.Fatalf("unexpected: %v", err)
	}

	log.Print("Gracefull exit")
}
