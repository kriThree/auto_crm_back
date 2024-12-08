package storage_car

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s CarStorage) Add(ctx context.Context, dto storage_models.AddCarDto) (id int64, err error) {

	const op = "postgres.car.Add"

	stmt, err := s.db.PrepareContext(ctx, `
		INSERT INTO cars (number, description, client_id)
		VALUES ($1, $2, $3)
		Returning id;
	`)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var resId int64 = 0

	err = stmt.QueryRowContext(ctx, dto.Number, dto.Description, dto.ClientId).Scan(&resId)

	if err != nil {
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return resId, nil
}
