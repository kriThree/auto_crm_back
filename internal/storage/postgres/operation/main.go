package storage_operation

import (
	"database/sql"
	storage_models "server_crm/internal/storage/models"
)

type OperationStorage struct {
	db *sql.DB
}

func New(db *sql.DB) storage_models.OperationDomain {
	return OperationStorage{db: db}
}
