package models

import (
	"context"
	storage_models "server_crm/internal/storage/models"
)

type CarRepo interface {
	Add(ctx context.Context, dto storage_models.AddCarDto) (id int64, err error)
	Update(ctx context.Context, dto storage_models.UpdateCarDto) error
	Get(ctx context.Context, userId int64) (cars []Car, err error)
	GetById(ctx context.Context, id int64) (car Car, err error)
	Delete(ctx context.Context, id int64) error
}

type Car struct {
	Id          int64
	Number      string
	Description string
	ClientId    int64
}