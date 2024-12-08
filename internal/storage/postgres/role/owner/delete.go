package storage_owner

import (
	"context"
	"fmt"
)

func (s OwnerStorage) Delete(ctx context.Context, userId int64) (err error) {

	const op = "postgres.roles.owner.Delete"

	stmt, err := s.db.PrepareContext(ctx, `
		DELETE FROM owners
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
