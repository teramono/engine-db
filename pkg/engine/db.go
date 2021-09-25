// Package engine ..
package engine

import "github.com/teramono/utilities/pkg/setup"

// DB ...
type DB struct {
	dbServer DBServerConn
}

// DBServerConn ...
type DBServerConn struct {
	address string
	port    string
}

// NewDB ...
func NewDB(setup *setup.Setup) DB {
	return DB{}
}

// Query ...
func Query(workspaceID string, query string) []byte {
	return []byte{}
}
