package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/teramono/utilities/pkg/broker"
	"github.com/teramono/utilities/pkg/setup"
)

// TODO:
const ConstPort = ":5052"

// DBServer ...
type DBServer struct {
	setup.CommonSetup
	backends []broker.Address
	pgURI    string
	pgConn   pg.Conn
	port     uint
}

// NewDBServer ...
func NewDBServer(setup setup.CommonSetup, pgURI string, port uint) (DBServer, error) {
	// TODO: Create and connect database
	return DBServer{
		CommonSetup: setup,
		pgURI: pgURI,
		port:  port,
	}, nil
}

// Listen ...
func (server *DBServer) Listen() error {
	router := gin.Default()

	router.POST("/query", server.Query)

	return router.Run(fmt.Sprintf(":%d", server.port))
}
