package storage_autoservice

import (
	"context"
	"fmt"
)

func (s AutoserviceStorage) Delete(ctx context.Context, id int64) error {

	const op = "postgres.autoservice.Delete"

	stmt, err := s.db.PrepareContext(ctx, `
		DELETE FROM autoservices
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
