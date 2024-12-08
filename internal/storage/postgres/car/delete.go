package storage_car

import (
	"context"
	"fmt"
)

func (s CarStorage) Delete(ctx context.Context, id int64) error {

	const op = "postgres.car.Delete"

	stmt, err := s.db.PrepareContext(ctx, `
		DELETE FROM cars
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
