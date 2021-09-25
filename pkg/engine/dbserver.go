package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/teramono/utilities/pkg/setup"
)

// DBServer ...
type DBServer struct {
	setup setup.Setup
	// dbs   []DBInfo // Not scalable
}

// NewDBServer ...
func NewDBServer(setup setup.Setup) (DBServer, error) {
	return DBServer{
		setup: setup,
	}, nil
}

// Listen ...
func (server *DBServer) Listen() error {
	router := gin.Default()

	router.POST("/query", server.Query)

	return router.Run(":5053") // TODO: Get from setup.Config
}
