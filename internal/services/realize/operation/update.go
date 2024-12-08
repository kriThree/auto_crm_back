package operation_service

import (
	"context"
	"log/slog"
	storage_models "server_crm/internal/storage/models"
)

func (s OperationService) Update(ctx context.Context, dto storage_models.UpdateOperationDto) error {

	const op = "service.operation.Update"

	log := s.l.With(
		slog.String("op", op),
	)
	log.Info("Start updating operation",
		slog.String("description", dto.Description),
		slog.Int64("car_id", dto.CarId),
		slog.Int64("work_id", dto.WorkId),
		slog.Int64("autoservice_id", dto.AutoserviceId),
		slog.Int64("id", dto.Id),
	)

	err := s.opP.Update(ctx, dto)

	if err != nil {
		log.Error("Update operation error", slog.Any("error", err.Error()))
		return err
	}

	log.Info("Update operation success",
		slog.Int64("id", dto.Id))

	return nil
}
