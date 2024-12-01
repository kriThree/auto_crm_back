package storage_user

import (
	"context"
	storage_models "server_crm/internal/storage/models"
)

func (s UserStorage) ReadOne(ctx context.Context, userId int64) (storage_models.User, error) {
	return storage_models.User{}, nil
}
