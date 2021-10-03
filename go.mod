module github.com/teramono/engine-db

go 1.16

require (
	github.com/gin-gonic/gin v1.7.4
	github.com/go-pg/pg/v10 v10.10.6
	github.com/teramono/utilities v0.0.0-20210919081101-b247dd3f53c0
)

replace github.com/teramono/utilities v0.0.0-20210919081101-b247dd3f53c0 => ../utilities
