package models

import (
	"context"
	storage_models "server_crm/internal/storage/models"
)

type WorkRepo interface {
	Add(ctx context.Context, dto storage_models.AddWorkDto) (id int64, err error)
	Update(ctx context.Context, dto storage_models.UpdateWorkDto) error
	Get(ctx context.Context) (works []Work, err error)
	GetById(ctx context.Context, id int64) (work Work, err error)
	Delete(ctx context.Context, id int64) error
}


type Work struct {
	Id        int64
	Cost      int
	Name      string
	CatalogId int64
}