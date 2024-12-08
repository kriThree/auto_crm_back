package storage_car

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s CarStorage) Update(ctx context.Context, dto storage_models.UpdateCarDto) error {

	const op = "postgres.car.Update"

	stmt, err := s.db.PrepareContext(ctx, `
		UPDATE cars
		SET number = $1, description = $2
		WHERE id = $3;
	`)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, dto.Number, dto.Description, dto.Id)

	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}
