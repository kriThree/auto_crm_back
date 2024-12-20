package storage_client

import (
	"context"
	"fmt"
	storage_errors "server_crm/internal/storage/errors"
)

func (s ClientStorage) GetOne(ctx context.Context, userId int64) (id int64, err error) {

	const op = "postgres.roles.client.GetById"

	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id
		FROM clients
		WHERE id = $1;
	`)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowContext(ctx, userId).Scan(&id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return 0, fmt.Errorf("%s : %w", op, storage_errors.ErrClientRoletNotFound)
		}
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return id, nil
}
