package car_service

import (
	"context"
	"errors"
	"log/slog"
	storage_errors "server_crm/internal/storage/errors"
	storage_models "server_crm/internal/storage/models"
)

func (s CarService) Add(ctx context.Context, dto storage_models.AddCarDto) (id int64, err error) {

	const op = "service.car.Add"

	log := s.l.With(
		slog.String("op", op),
		slog.String("number", dto.Number),
		slog.String("description", dto.Description),
		slog.Int64("client_id", dto.ClientId),
	)

	log.Info("Start adding car")

	_, err = s.clP.GetOne(ctx, dto.ClientId)

	if err != nil {
		if errors.Is(err, storage_errors.ErrClientRoletNotFound) {
			log.Error("Client not found", slog.Any("error", err.Error()))
			return 0, err
		}
		log.Error("Get client error", slog.Any("error", err.Error()))
		return 0, err
	}

	id, err = s.crP.Add(ctx, dto)

	if err != nil {
		log.Error("Add car error", slog.Any("error", err.Error()))
		return 0, err
	}

	log.Info("Add car success",
		slog.Int64("id", id),
	)
	return id, nil
}
