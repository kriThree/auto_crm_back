package storage_catalog

import (
	"database/sql"
	storage_models "server_crm/internal/storage/models"
)

type catalogStorage struct {
	db *sql.DB
}

func New(db *sql.DB) storage_models.CatalogDomain {
	return catalogStorage{db: db}
}
