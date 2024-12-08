package storage_admin

import (
	"context"
	"fmt"
)

func (s AdminStorage) Delete(ctx context.Context, userId int64) (err error) {

	const op = "postgres.roles.admin.Delete"

	stmt, err := s.db.PrepareContext(ctx, `
		DELETE FROM admins
		WHERE user_id = $1;
	`)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, userId)

	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}
