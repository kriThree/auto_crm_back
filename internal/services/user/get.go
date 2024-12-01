package user_serivce

import (
	"context"
	storage_models "server_crm/internal/storage/models"
)

func (s UserService) Get(ctx context.Context) ([]storage_models.User, error) {

	const op = "service.user.Get"
	return s.usP.Read(ctx)
}

