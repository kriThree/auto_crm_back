package models

import (
	"context"
	storage_models "server_crm/internal/storage/models"
)

type OperationRepo interface {
	Create(ctx context.Context, operation storage_models.AddOperationDto) (id int64, err error)
	GetForCar(ctx context.Context, carId int64) (operations []Operation, err error)
	GetForWork(ctx context.Context, workId int64) ([]Operation, error)
	GetForAutoservice(ctx context.Context, autoserviceId int64) (operations []Operation, err error)
	GetById(ctx context.Context, id int64) (operation Operation, err error)
	Update(ctx context.Context, dto storage_models.UpdateOperationDto) error
	Delete(ctx context.Context, id int64) error
}

type Operation struct {
	Id          int64
	Description string
	Car         Car
	Work        Work
	Autoservice Autoservice
}
