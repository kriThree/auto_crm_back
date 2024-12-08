package storage_admin

import (
	"database/sql"
)

type AdminStorage struct {
	db *sql.DB
}

func New(db *sql.DB) AdminStorage {
	return AdminStorage{db: db}
}