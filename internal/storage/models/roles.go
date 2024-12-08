package storage_models

import (
	"context"
)

type RoleDomain interface {
	Add(ctx context.Context, userId int64) (roleId int64, err error)
	GetByUserId(ctx context.Context, userId int64) (roleId int64, err error)
	GetOne(ctx context.Context, roleId int64) (userId int64, err error)
	Delete(ctx context.Context, userId int64) (err error)
}

var (
	ROLE_ADMIN  = "admin"
	ROLE_OWNER  = "owner"
	ROLE_CLIENT = "client"
)

