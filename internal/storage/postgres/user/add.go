package storage_user

import (
	"context"
	"fmt"
	storage_errors "server_crm/internal/storage/errors"
	storage_models "server_crm/internal/storage/models"
	"strings"
)

func (s UserStorage) Add(ctx context.Context, dto storage_models.AddUserDto) (int64, error) {

	const op = "postgres.user.Add"
	stmt, err := s.db.PrepareContext(ctx, `
		INSERT INTO users (name, email, password)
		VALUES ($1, $2, $3)
		Returning id;
	`)

	if err != nil {

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var resId int64 = 0

	err = stmt.QueryRowContext(ctx, dto.Name, dto.Email, dto.Password).Scan(&resId)

	if err != nil {
		if strings.Contains(err.Error(), "users_email_key") && strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return 0, fmt.Errorf("%s : %w", op, storage_errors.ErrEmailAlreadyExist)
		}
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return resId, nil
}
