package storage_user

import (
	"context"
	"fmt"
	storage_errors "server_crm/internal/storage/errors"
	storage_models "server_crm/internal/storage/models"
)

func (s UserStorage) Read(ctx context.Context) ([]storage_models.User, error) {
	return nil, nil
}

func (s UserStorage) FindByEmail(ctx context.Context, email string) (storage_models.User, error) {

	const op = "postgres.user.FindByEmail"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id, name, email, password, created_at
		FROM users
		WHERE
	`)

	if err != nil {
		return storage_models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	var user storage_models.User

	err = stmt.QueryRowContext(ctx, email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return storage_models.User{}, fmt.Errorf("%s : %w", op, storage_errors.ErrUserNotFound)
		}
		return storage_models.User{}, fmt.Errorf("%s : %w", op, err)
	}

	return user, nil
}
