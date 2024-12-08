package storage_owner

import (
	"context"
	"fmt"
	storage_errors "server_crm/internal/storage/errors"
)

func (s OwnerStorage) GetOne(ctx context.Context, id int64) (int64, error) {
	
	const op = "postgres.roles.client.GetById"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id
		FROM owners
		WHERE id = $1;
	`)


	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowContext(ctx, id).Scan(&id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return 0, fmt.Errorf("%s : %w", op, storage_errors.ErrOwnerRoleNotFound)
		}
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return 0, nil
}