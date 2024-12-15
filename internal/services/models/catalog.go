package models

import (
	"context"
	storage_models "server_crm/internal/storage/models"
)

type CatalogRepo interface {
	Add(ctx context.Context, dto storage_models.AddCatalogDto) (id int64, err error)
	Update(ctx context.Context, dto storage_models.UpdateCatalogDto) error
	GetForAdmin(ctx context.Context, adminId int64) (catalogs []Catalog, err error)
	Get(ctx context.Context) (catalogs []Catalog, err error)
	GetById(ctx context.Context, id int64) (catalog Catalog, err error)
	Delete(ctx context.Context, id int64) error
}

type Catalog struct {
	Id        int64
	Name      string
	CreaterId int64
	Works     []Work
}
