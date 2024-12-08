package storage_catalog

import (
	"context"
	"fmt"
)

func (s catalogStorage) Delete(ctx context.Context, id int64) error {

	const op = "postgres.catalog.Delete"

	stmt, err := s.db.PrepareContext(ctx, `
		DELETE FROM catalogs
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
