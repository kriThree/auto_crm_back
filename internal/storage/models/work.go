package storage_models

import "context"

type WorkDomain interface {
	Add(ctx context.Context, dto AddWorkDto) (id int64, err error)
	Get(ctx context.Context) (work []Work, err error)
	GetByCatalogId(ctx context.Context, catalogId int64) (works []Work, err error)
	GetById(ctx context.Context, id int64) (work Work, err error)
	Update(ctx context.Context, dto UpdateWorkDto) error
	Delete(ctx context.Context, id int64) error
}

type Work struct {
	Id        int64
	Cost      int
	Name      string
	CatalogId int64
}

type AddWorkDto struct {
	Cost      int
	Name      string
	CatalogId int64
}
type UpdateWorkDto struct {
	Id        int64
	Cost      int
	Name      string
	CatalogId int64
}
