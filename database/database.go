package database

import (
	"database/sql"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var pool *bun.DB

func NewConnection() *bun.DB {
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("POSTGRES_URL") + "?sslmode=disable")))
	dbpool := bun.NewDB(sqldb, pgdialect.New())

	pool = dbpool
	return pool
}

func Instance() (*bun.DB, error) {
	if pool != nil {
		if err := pool.Ping(); err != nil {
			return nil, err
		}
		return pool, nil
	}
	pool = NewConnection()
	return pool, nil
}
