package user_serivce

import (
	"context"
	"fmt"
	"log/slog"
	storage_models "server_crm/internal/storage/models"
)

func (s UserService) Update(
	ctx context.Context,
	userId int64,
	dto storage_models.UpdateUserDto,
) error {

	const op = "service.user.Update"

	log := s.l.With(
		slog.String("op", op),
		slog.String("email", dto.Email),
		slog.String("password", dto.Password),
		slog.String("name", dto.Name),
	)

	err := s.usP.Update(ctx, dto)

	if err != nil {
		log.Error("Update user error", slog.Any("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
