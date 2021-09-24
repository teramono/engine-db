package tests

import (
	"testing"

	"github.com/teramono/utilities/pkg/setup"
	"github.com/teramono/engine-db/pkg/engine"
)

func TestFSServer(t *testing.T) {
	setup, err := setup.NewSetup(false)
	if err != nil {
		// ...
	}

	server, err := engine.NewFSServer(setup)
	if err != nil {
		// ...
	}

	server.Listen()
}
