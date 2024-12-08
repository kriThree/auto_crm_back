package storage_models

import "context"

type CatalogDomain interface {
	Add(ctx context.Context, dto AddCatalogDto) (id int64, err error)
	GetByAdminId(ctx context.Context, adminId int64) (catalogs []Catalog, err error)
	Get(ctx context.Context) (catalogs []Catalog, err error)
	GetById(ctx context.Context, id int64) (catalog Catalog, err error)
	Update(ctx context.Context, dto UpdateCatalogDto) error
	Delete(ctx context.Context, id int64) error
}
type Catalog struct {
	Id        int64
	Name      string
	CreaterId int64
}
type AddCatalogDto struct {
	AdminId int64
	Name string
}
type UpdateCatalogDto struct {
	Id      int64
	Name    string
	AdminId int64
}
