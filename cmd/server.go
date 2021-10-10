package main

import (
	"log"

	"github.com/teramono/engine-db/pkg/server"
	"github.com/teramono/utilities/pkg/setup"
)

func main() {
	// Establish common setup.
	setup, err := setup.NewCommonSetup()
	if err != nil {
		log.Panic(err)
	}

	// Create server.
	server, err := server.NewDBServer(setup)
	if err != nil {
		log.Panic(err)
	}

	defer server.BrokerClient.Close()

	// Activate subscriptions.
	err = server.ActivateSubscriptions()
	if err != nil {
		log.Panic(err)
	}
}
