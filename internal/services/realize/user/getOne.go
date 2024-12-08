package user_serivce

import (
	"context"
	"fmt"
	"server_crm/internal/services/models"
)

func (s UserService) GetOne(ctx context.Context, userId int64) (models.User, error) {

	const op = "service.user.Get"

	user, err := s.usP.GetOne(ctx, userId)

	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	role, roleId, err := s.identifyRole(ctx, userId)

	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return s.fromStorageToDomain(user, s.writeUserRoles(roleId, role)), err
}
