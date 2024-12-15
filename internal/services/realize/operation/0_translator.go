package operation_service

import (
	"context"
	"fmt"
	"server_crm/internal/auxiliary"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

func (s OperationService) fromStorageToDomain(storage_operation storage_models.Operation,
	Car storage_models.Car,
	Work storage_models.Work,
	Autoservice storage_models.Autoservice,
) models.Operation {
	return models.Operation{
		Id:          storage_operation.Id,
		Description: storage_operation.Description,
		Car: models.Car{
			Id:          Car.Id,
			Number:      Car.Number,
			Description: Car.Description,
			ClientId:    Car.ClientId,
		},
		Work: models.Work{
			Id:        Work.Id,
			Name:      Work.Name,
			Cost:      Work.Cost,
			CatalogId: Work.CatalogId,
		},
		Autoservice: models.Autoservice{
			Id:        Autoservice.Id,
			Name:      Autoservice.Name,
			Address:   Autoservice.Address,
			Phone:     Autoservice.Phone,
			Email:     Autoservice.Email,
			OwnerId:   Autoservice.OwnerId,
			CreatedAt: Autoservice.CreatedAt,
		},
	}
}
func (s OperationService) fromStorageToDomainWithQuery(ctx context.Context, storage_operation storage_models.Operation) (models.Operation, error) {

	const op = "service.operation.fromStorageToDomainWithQuery"

	storage_car, err := s.crP.GetById(ctx, storage_operation.CarId)

	if err != nil {
		return models.Operation{}, fmt.Errorf("%s: %w", op, err)
	}

	storage_work, err := s.wkP.GetById(ctx, storage_operation.WorkId)

	if err != nil {

		return models.Operation{}, fmt.Errorf("%s: %w", op, err)
	}

	storage_autoservice, err := s.auP.GetById(ctx, storage_operation.AutoserviceId)

	if err != nil {
		return models.Operation{}, fmt.Errorf("%s: %w", op, err)
	}

	return s.fromStorageToDomain(storage_operation, storage_car, storage_work, storage_autoservice), nil
}

func (s OperationService) fromStorageToDomainArray(ctx context.Context, storage_operations []storage_models.Operation) ([]models.Operation, error) {
	var operations []models.Operation
	operations, err := auxiliary.NewWorker(
		ctx,
		storage_operations,
		func(ctx auxiliary.Context, storage_operation storage_models.Operation) models.Operation {
			operation, err := s.fromStorageToDomainWithQuery(ctx, storage_operation)
			if err != nil {
				ctx.PushError(err)
				return models.Operation{}
			}
			return operation
		},
	)

	if err != nil {
		return nil, err
	}

	return operations, nil
}
