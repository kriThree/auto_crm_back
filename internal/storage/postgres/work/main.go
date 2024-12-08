package storage_work

import (
	"database/sql"
	storage_models "server_crm/internal/storage/models"
)

type workStorage struct {
	db *sql.DB
}

func New(db *sql.DB) storage_models.WorkDomain {
	return workStorage{db: db}
}
