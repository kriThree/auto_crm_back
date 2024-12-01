package storage_user

import (
	"database/sql"
)

type UserStorage struct {
	db *sql.DB
}

func New(db *sql.DB) UserStorage {
	return UserStorage{db: db}
}
