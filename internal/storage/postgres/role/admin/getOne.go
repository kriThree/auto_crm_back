package storage_admin

import (
	"context"
	"fmt"
	storage_errors "server_crm/internal/storage/errors"
)

func (s AdminStorage) GetOne(ctx context.Context, userId int64) (id int64, err error){

	const op = "postgres.roles.admin.GetById"
	
	stmt, err := s.db.PrepareContext(ctx, `
		SELECT id
		FROM admins
		WHERE id = $1;
	`)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRowContext(ctx, userId).Scan(&id)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return 0, fmt.Errorf("%s : %w", op, storage_errors.ErrAdminRoleNotFound)
		}
		return 0, fmt.Errorf("%s : %w", op, err)
	}
	return id, nil
}