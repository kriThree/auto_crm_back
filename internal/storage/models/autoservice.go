package storage_models

import (
	"context"
	"time"
)

type AutoserviceDomain interface {
	Add(ctx context.Context, dto AddAutoserviceDto) (id int64, err error)
	Update(ctx context.Context, dto UpdateAutoserviceDto) error
	GetByOwnerId(ctx context.Context, userId int64) (autoservices []Autoservice, err error)
	GetById(ctx context.Context, id int64) (autoservice Autoservice, err error)
	Delete(ctx context.Context, id int64) error
}
type Autoservice struct {
	Id        int64
	Name      string
	Address   string
	Phone     string
	Email     string
	OwnerId   int64
	CreatedAt time.Time
}

type AddAutoserviceDto struct {
	Name       string
	Address    string
	Phone      string
	Email      string
	Owner_id   int64
	Created_at time.Time
}
type UpdateAutoserviceDto struct {
	Id      int64
	Name    string
	Address string
	Phone   string
	Email   string
}
