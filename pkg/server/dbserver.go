package server

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/nats-io/nats.go"
	"github.com/teramono/utilities/pkg/broker"
	"github.com/teramono/utilities/pkg/setup"
)

type DBServer struct {
	setup.CommonSetup
	DB *pg.DB
}

func NewDBServer(setup setup.CommonSetup) (DBServer, error) {
	opts, err := pg.ParseURL(setup.Config.Engines.DB.DBURL)
	if err != nil {
		return DBServer{}, err
	}

	db := pg.Connect(opts)

	return DBServer{
		CommonSetup: setup,
		DB:          db,
	}, nil
}

func (server *DBServer) LogsVersion() uint {
	return server.Config.Broker.Subscriptions.Logs.Version
}

func (server *DBServer) ActivateSubscriptions() error {
	// Create channels for subscribed subjects.
	queryCh := make(chan *nats.Msg, server.BrokerClient.Opts.SubChanLen)
	querySub, err := server.BrokerClient.ChanQueueSubscribe(
		broker.GetWorkspacesSubjectByEngine(broker.EngineDB, &server.Config, "query"),
		broker.GetWorkspacesResponderGroupByEngine(broker.EngineDB, &server.Config, "query"),
		queryCh,
	)
	if err != nil {
		return err
	}

	defer querySub.Unsubscribe()
	fmt.Println(">> Subscriptions set up!")

	// Listen to subscribed messages.
	for msg := range queryCh {
		go broker.PanicWrap(msg, server.Query)
	}

	return nil
}
