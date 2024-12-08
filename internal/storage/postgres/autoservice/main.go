package storage_autoservice

import (
	"database/sql"
	storage_models "server_crm/internal/storage/models"
)

type AutoserviceStorage struct {
	db *sql.DB
}

func New(db *sql.DB) storage_models.AutoserviceDomain {
	return AutoserviceStorage{db: db}
}
