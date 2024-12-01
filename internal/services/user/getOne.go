package user_serivce

import (
	"context"
	"server_crm/internal/services/models"
)

func (s UserService) GetOne(ctx context.Context, userId int64) (models.User, error) {

	const op = "service.user.Get"

	user, err := s.usP.ReadOne(ctx, userId)
	return s.fromStorageToDomain(user), err
}
