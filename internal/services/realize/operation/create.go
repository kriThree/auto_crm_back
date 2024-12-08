package operation_service

import (
	"context"
	"log/slog"
	storage_models "server_crm/internal/storage/models"
)

func (s OperationService) Create(ctx context.Context, dto storage_models.AddOperationDto) (id int64, err error) {

	const op = "service.operation.Add"

	log := s.l.With(
		slog.String("op", op),
		slog.String("description", dto.Description),
		slog.Int64("car_id", dto.CarId),
		slog.Int64("work_id", dto.WorkId),
		slog.Int64("autoservice_id", dto.AutoserviceId),
	)

	log.Info("Start adding operation")

	id, err = s.opP.Add(ctx, dto)

	log.Info("Add operation success",
		slog.Int64("id", id),
	)

	if err != nil {
		log.Error("Add operation error", slog.Any("error", err.Error()))
		return 0, err
	}

	return id, nil
}
