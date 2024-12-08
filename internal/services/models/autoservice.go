package models

import (
	"context"
	storage_models "server_crm/internal/storage/models"
	"time"
)

type Autoservice struct {
	Id        int64
	Name      string
	Address   string
	Phone     string
	Email     string
	OwnerId   int64
	CreatedAt time.Time
}
type AutoserviceRepo interface {
	Add(ctx context.Context, dto storage_models.AddAutoserviceDto) (id int64, err error)
	Update(ctx context.Context, dto storage_models.UpdateAutoserviceDto) error
	Get(ctx context.Context, userId int64) (autoservices []Autoservice, err error)
	GetById(ctx context.Context, id int64) (autoservice Autoservice, err error)
	Delete(ctx context.Context, id int64) error
}
