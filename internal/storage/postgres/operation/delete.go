package storage_operation

import (
	"context"
	"fmt"
)

func (s OperationStorage) Delete(ctx context.Context, id int64) error {

	const op = "postgres.operation.Delete"

	stmt, err := s.db.PrepareContext(ctx, `
		DELETE FROM operations
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
