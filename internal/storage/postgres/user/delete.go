package storage_user

import (
	"context"
	"fmt"
)

func (s UserStorage) Delete(ctx context.Context, userId int64) (err error) {

	const op = "postgres.user.Delete"

	stmt, err := s.db.PrepareContext(ctx, `
		DELETE FROM users
		WHERE id = $1;
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