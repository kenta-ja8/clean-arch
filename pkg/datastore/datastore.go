package datastore

import (
	"context"
	"database/sql"

	"github.com/kenta-ja8/clean-arch/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

type Datastore struct {
	DB *sql.DB
}

func NewDataStore(ctx context.Context, config config.Config) (Datastore, error) {

	db, err := sql.Open(config.DBDriver, config.DBName)
	if err != nil {
		panic(err)
	}
	return Datastore{
		DB: db,
	}, nil
}

func (d Datastore) Close() {
	d.DB.Close()
}
