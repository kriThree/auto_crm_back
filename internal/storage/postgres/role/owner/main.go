package storage_owner

import (
	"database/sql"
)

type OwnerStorage struct {
	db *sql.DB
}

func New(db *sql.DB) OwnerStorage {
	return OwnerStorage{db: db}
}
