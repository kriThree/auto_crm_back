package storage_operation

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s OperationStorage) Update(ctx context.Context, dto storage_models.UpdateOperationDto) error {

	const op = "postgres.operation.Update"

	stmt, err := s.db.PrepareContext(ctx, `
		UPDATE operations
		SET car_id = $1, work_id = $2, autoservice_id = $3, description = $4
		WHERE id = $5;
	`)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.ExecContext(ctx, dto.CarId, dto.WorkId, dto.AutoserviceId, dto.Description, dto.Id)

	if err != nil {
		return fmt.Errorf("%s : %w", op, err)
	}

	return nil
}
