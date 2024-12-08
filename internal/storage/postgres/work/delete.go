package storage_work

import (
	"context"
	"fmt"
)

func (s workStorage) Delete(ctx context.Context, id int64) error {

	const op = "postgres.work.Delete"

	stmt, err := s.db.PrepareContext(ctx, `
		DELETE FROM works
		WHERE id = $1;
	`)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, id)

	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}
