package storage_operation

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s OperationStorage) Add(ctx context.Context, dto storage_models.AddOperationDto) (id int64, err error) {

	const op = "postgres.operation.Add"

	stmt, err := s.db.PrepareContext(ctx, `
		INSERT INTO operations (car_id ,work_id, autoservice_id, description)
		VALUES ($1, $2, $3, $4)
		Returning id;
	`)

	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	var resId int64 = 0

	err = stmt.QueryRowContext(ctx, dto.CarId, dto.WorkId, dto.AutoserviceId, dto.Description).Scan(&resId)

	if err != nil {
		return 0, fmt.Errorf("%s : %w", op, err)
	}

	return resId, nil
}
