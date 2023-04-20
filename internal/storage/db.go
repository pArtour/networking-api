package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	Db *pgxpool.Pool
)

func ConnectDB(connectionString string) error {
	var err error
	Db, err = pgxpool.Connect(context.Background(), connectionString)
	if err != nil {
		return fmt.Errorf("unable to connect to the database: %w", err)
	}

	return nil
}

func CloseDB() {
	Db.Close()
}
