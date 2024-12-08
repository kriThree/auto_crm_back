package user_serivce

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s UserService) Delete(ctx context.Context, id int64) error {

	const op = "service.user.Delete"

	role, roleId, err := s.identifyRole(ctx, id)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if role == storage_models.ROLE_ADMIN {
		if err := s.rsP.Admin.Delete(ctx, roleId); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	} else if role == storage_models.ROLE_OWNER {
		if err := s.rsP.Owner.Delete(ctx, roleId); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	} else if role == storage_models.ROLE_CLIENT {
		if err := s.rsP.Client.Delete(ctx, roleId); err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	err = s.usP.Delete(ctx, id)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
