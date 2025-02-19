package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vsrtferrum/OzonIntro/internal/errors"
)

type DatabaseConnection interface {
	Connect() error
	Ping() error
	Close()
}

func (db *Database) Connect() error {
	config, err := pgxpool.ParseConfig(db.connStr)
	if err != nil {
		return errors.ErrCreateConfig
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return errors.ErrCreateConnection
	}

	db.pool = pool
	return db.Ping()
}

func (db *Database) Ping() error {
	err := db.pool.Ping(context.Background())
	if err != nil {
		return errors.ErrConnectionTimeLimit
	}
	return nil
}

func (db *Database) Close() {
	db.pool.Close()
}
