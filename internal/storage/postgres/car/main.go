package storage_car

import (
	"database/sql"
	storage_models "server_crm/internal/storage/models"
)

type CarStorage struct {
	db *sql.DB
}

func New(db *sql.DB) storage_models.CarDomain {
	return CarStorage{db: db}
}
