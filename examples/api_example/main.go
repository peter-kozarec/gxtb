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

	// Connect to api client
	if err := c.Connect(ctx); err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	defer c.Disconnect()

	// Login to api client
	if _, err := c.Login(ctx, os.Getenv("XTB_UserId"), os.Getenv("XTB_Password"), "testApp"); err != nil {
		log.Fatalf("unable to login: %v", err)
	}
	defer c.Logout(ctx)

	// Get api version
	version, err := c.GetVersion(ctx)
	if err != nil {
		log.Fatalf("unable to retrieve api version: %v", err)
	}
	log.Printf("Api version: %s\n", version)

	// Get broker server time
	serverTime, err := c.GetServerTime(ctx)
	if err != nil {
		log.Fatalf("unable to retrieve server time: %v", err)
	}
	log.Printf("Server time: %s\n", serverTime.TimeString)
}
