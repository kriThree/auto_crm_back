package storage_models

import "context"

type CarDomain interface {
	Add(ctx context.Context, dto AddCarDto) (id int64, err error)
	GetById(ctx context.Context, id int64) (car Car, err error)
	GetByClientId(ctx context.Context, userId int64) (cars []Car, err error)
	Update(ctx context.Context, dto UpdateCarDto) error
	Delete(ctx context.Context, id int64) error
}

type AddCarDto struct {
	Number      string
	Description string
	ClientId    int64
}
type UpdateCarDto struct {
	Id          int64
	Number      string
	Description string
}
type Car struct {
	Id          int64
	Number      string
	Description string
	ClientId    int64
}
