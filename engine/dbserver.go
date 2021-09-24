package engine

import (
	"github.com/gin-gonic/gin"
	"github.com/teramono/utilities/pkg/setup"
)

// DBServer ...
type DBServer struct {
	setup setup.Setup
	dbs   []DB
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

	// Use GRPC instead.
	router.POST("/", func(c *gin.Context) {})

	return router.Run()
}
