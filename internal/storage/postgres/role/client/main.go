package storage_client

import (
	"database/sql"
)

type ClientStorage struct {
	db *sql.DB
}

func New(db *sql.DB) ClientStorage {
	return ClientStorage{db: db}
}
