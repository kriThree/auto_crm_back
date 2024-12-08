package storage_owner

import (
	"context"
	"fmt"
	storage_errors "server_crm/internal/storage/errors"
)

func (s OwnerStorage) GetByUserId(ctx context.Context, userId int64) (id int64, err error) {

	const op = "postgres.roles.owner.GetByUserId"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id
		FROM owners
		WHERE user_id = $1;
	`)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowContext(ctx, userId).Scan(&id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return 0, fmt.Errorf("%s : %w", op, storage_errors.ErrOwnerRoleNotFound)
		}
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return id, nil
}
