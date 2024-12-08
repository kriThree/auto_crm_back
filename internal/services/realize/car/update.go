package car_service

import (
	"context"
	"log/slog"
	storage_models "server_crm/internal/storage/models"
)

func (s CarService) Update(ctx context.Context, dto storage_models.UpdateCarDto) error {

	const op = "service.car.Update"

	log := s.l.With(
		slog.String("op", op),
	)
	log.Info("Start updating car",
		slog.String("number", dto.Number),
		slog.String("description", dto.Description),
		slog.Int64("id", dto.Id),
	)

	err := s.crP.Update(ctx, dto)

	if err != nil {
		log.Error("Update car error", slog.Any("error", err.Error()))
		return err
	}

	log.Info("Update car success",
		slog.Int64("id", dto.Id))

	return nil
}
