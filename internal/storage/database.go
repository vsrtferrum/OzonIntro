package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	connStr string
	pool    *pgxpool.Pool
}

func NewDatabase(connStr string) (*Database, error) {
	db := Database{connStr: connStr}

	err := db.Connect()
	if err != nil {
		return &Database{}, err
	}

	return &db, nil
}
