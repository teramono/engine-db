package main

import (
	"log"

	"github.com/teramono/engine-db/pkg/server"
	"github.com/teramono/utilities/pkg/setup"
)

// TODO:
const Port = 5051
const PGURI = ""

func main() {
	// ...
	setup, err := setup.NewCommonSetup()
	if err != nil {
		log.Fatalln(err)
	}

	// ...
	server, err := server.NewDBServer(setup, PGURI, Port)
	if err != nil {
		log.Fatalln(err)
	}

	server.Listen()
}
