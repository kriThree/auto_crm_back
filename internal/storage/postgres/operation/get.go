package storage_operation

import (
	"context"
	"fmt"
	storage_models "server_crm/internal/storage/models"
)

func (s OperationStorage) GetByCarId(ctx context.Context, operationId int64) ([]storage_models.Operation, error) {

	const op = "postgres.operation.Get"

	return s.getByQuery(ctx, op, "operation_id = $1", operationId)
}
func (s OperationStorage) GetByAutoserviceId(ctx context.Context, autoserviceId int64) ([]storage_models.Operation, error) {

	const op = "postgres.operation.GetByAutoserviceId"
	return s.getByQuery(ctx, op, "autoservice_id = $1", autoserviceId)
}
func (s OperationStorage) GetByWorkId(ctx context.Context, workId int64) ([]storage_models.Operation, error) {
	const op = "postgres.operation.GetByWorkId"
	return s.getByQuery(ctx, op, "work_id = $1", workId)
}
func (s OperationStorage) getByQuery(ctx context.Context, op string, query string, args ...any) ([]storage_models.Operation, error) {

	stmt, err := s.db.PrepareContext(ctx, fmt.Sprintf(
		`SELECT id, description, car_id, work_id, autoservice_id
		FROM operations
		WHERE %s `, query))

	if err != nil {
		return []storage_models.Operation{}, err
	}
	rows, err := stmt.QueryContext(ctx, args)
	if err != nil {
		return []storage_models.Operation{}, err
	}

	var operations []storage_models.Operation
	for rows.Next() {
		var operation storage_models.Operation
		err = rows.Scan(&operation.Id, &operation.Description, &operation.CarId, &operation.WorkId, &operation.AutoserviceId)

		if err != nil {
			return nil, fmt.Errorf("%s : %w", op, err)
		}

		operations = append(operations, operation)
	}
	return operations, nil
}
func (s OperationStorage) GetById(ctx context.Context, id int64) (storage_models.Operation, error) {
	const op = "postgres.operation.GetById"

	stmt, err := s.db.PrepareContext(ctx,
		`SELECT id, description, car_id, work_id, autoservice_id
		FROM operations
		WHERE id = $1;`,
	)

	if err != nil {
		return storage_models.Operation{}, fmt.Errorf("%s: %w", op, err)
	}

	var operation storage_models.Operation

	err = stmt.QueryRowContext(ctx, id).Scan(&operation.Id, &operation.Description, &operation.CarId, &operation.WorkId, &operation.AutoserviceId)

	if err != nil {
		return storage_models.Operation{}, fmt.Errorf("%s : %w", op, err)
	}

	return operation, nil
}
