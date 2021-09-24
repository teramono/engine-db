// Package engine ..
package engine

// DB ...
type DB struct {
	dbServer DBServerConn
}

// DBServerConn ...
type DBServerConn struct{}

// NewDB ...
func NewDB() DB {
	return DB{}
}

// Model ...
func (db *DB) Model(path string) Model {
	return Model{db: db}
}
