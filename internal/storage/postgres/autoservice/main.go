package storage_autoservice

import (
	"database/sql"
)

type AutoserviceStorage struct {
	db *sql.DB
}

func New(db *sql.DB) AutoserviceStorage {
	return AutoserviceStorage{db: db}
}
