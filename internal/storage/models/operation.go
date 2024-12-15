package storage_models

import "context"

type OperationDomain interface {
	Add(ctx context.Context, dto AddOperationDto) (id int64, err error)
	GetByCarId(ctx context.Context, carId int64) (operations []Operation, err error)
	GetByAutoserviceId(ctx context.Context, autoserviceId int64) (operations []Operation, err error)
	GetByWorkId(ctx context.Context, workId int64) (operations []Operation, err error)
	GetById(ctx context.Context, id int64) (catalog Operation, err error)
	Update(ctx context.Context, dto UpdateOperationDto) error
	Delete(ctx context.Context, id int64) error
}
type Operation struct {
	Id            int64
	Description   string
	CarId         int64
	WorkId        int64
	AutoserviceId int64
}

type AddOperationDto struct {
	Description   string
	CarId         int64
	WorkId        int64
	AutoserviceId int64
}
type UpdateOperationDto struct {
	Id int64
	Description   string
	CarId         int64
	WorkId        int64
	AutoserviceId int64
}
