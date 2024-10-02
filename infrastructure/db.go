package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	pool    *pgxpool.Pool
	sqlDb   *sql.DB
	builder *goqu.Database
}

func (db *DB) Connect(ctx context.Context, connStr string) {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		// TODO: replace by logger
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	db.pool = pool
	db.sqlDb = stdlib.OpenDBFromPool(pool)
	db.builder = goqu.New("postgres", db.sqlDb)
}

// dbPool, err := pgxpool.New(context.Background(), cfg.Postgres.ConfigString())
// 	if err != nil {
// 		// TODO: replace by logger
// 		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
// 		os.Exit(1)
// 	}
// 	conn, err := dbPool.Acquire(context.Background());
// 	if err != nil {
// 		return nil
// 	}
// 	conn.

// 	return dbPool
