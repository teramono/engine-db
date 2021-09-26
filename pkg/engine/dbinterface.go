// Package engine ..
package engine

import "github.com/teramono/utilities/pkg/setup"

// DBInterface ...
type DBInterface struct {
	dbServer DBInterfaceServerConn
}

// DBInterfaceServerConn ...
type DBInterfaceServerConn struct {
	address string
	port    string
}

// NewDBInterface ...
func NewDBInterface(setup *setup.Setup) DBInterface {
	return DBInterface{}
}

// Query ...
func Query(workspaceID string, query string) []byte {
	return []byte{}
}
