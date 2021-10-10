module github.com/teramono/engine-db

go 1.16

require (
	github.com/go-pg/pg/v10 v10.10.6
	github.com/nats-io/nats.go v1.12.3
	github.com/teramono/utilities v0.0.0-20210919081101-b247dd3f53c0
)

replace github.com/teramono/utilities v0.0.0-20210919081101-b247dd3f53c0 => ../utilities
