package main

import (
	"log"

	"github.com/teramono/engine-db/pkg/engine"
	"github.com/teramono/utilities/pkg/setup"
)

func main() {
	// ...
	setup, err := setup.NewSetup(false)
	if err != nil {
		log.Fatalln(err)
	}

	// ...
	server, err := engine.NewFSServer(setup)
	if err != nil {
		log.Fatalln(err)
	}

	server.Listen()
}
