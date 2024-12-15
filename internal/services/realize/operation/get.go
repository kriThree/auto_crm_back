package operation_service

import (
	"context"
	"log/slog"
	"server_crm/internal/services/models"
)

func (s OperationService) GetForCar(ctx context.Context, carId int64) ([]models.Operation, error) {

	const op = "service.operation.Get"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("car_id", carId),
	)

	log.Info("Start getting operation for car")

	storageOperations, err := s.opP.GetByCarId(ctx, carId)
	if err != nil {
		return nil, err
	}
	
	operations, err := s.fromStorageToDomainArray(ctx, storageOperations)

	if err != nil {
		log.Error("Failed to convert storage to domain", slog.Any("error", err.Error()))
		return nil, err
	}

	log.Info("Get operations for car success",
		slog.Int("count", len(operations)),
	)

	return operations, nil
}
func (s OperationService) GetForWork(ctx context.Context, workId int64) ([]models.Operation, error) {

	const op = "service.operation.Get"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("work_id", workId),
	)

	log.Info("Start getting operation for work")

	storageOperations, err := s.opP.GetByWorkId(ctx, workId)
	if err != nil {
		log.Error("Get operations error", slog.Any("error", err.Error()))
		return nil, err
	}

	operations, err := s.fromStorageToDomainArray(ctx, storageOperations)

	if err != nil {
		log.Error("Failed to convert storage to domain", slog.Any("error", err.Error()))
		return nil, err
	}

	log.Info("Get operations for work success",
		slog.Int("count", len(operations)),
	)

	return operations, nil
}
func (s OperationService) GetForAutoservice(ctx context.Context, autoserviceId int64) ([]models.Operation, error) {

	const op = "service.operation.Get"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("autoservice_id", autoserviceId),
	)

	log.Info("Start getting operation for autoservice")

	storageOperations, err := s.opP.GetByAutoserviceId(ctx, autoserviceId)
	if err != nil {
		return nil, err
	}
	operations, err := s.fromStorageToDomainArray(ctx, storageOperations)

	if err != nil {
		log.Error("Failed to convert storage to domain", slog.Any("error", err.Error()))
		return nil, err
	}

	log.Info("Get operations for autoservice success",
		slog.Int("count", len(operations)),
	)

	return operations, nil
}
func (s OperationService) GetById(ctx context.Context, id int64) (models.Operation, error) {

	const op = "service.operation.GetById"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("id", id),
	)

	log.Info("Start getting operation by id")

	storageOperation, err := s.opP.GetById(ctx, id)
	if err != nil {
		return models.Operation{}, err
	}

	log.Info("Get operation by id success",
		slog.Int64("id", id),
	)

	return s.fromStorageToDomainWithQuery(ctx,storageOperation)
}