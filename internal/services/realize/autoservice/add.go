package autoservice_service

import (
	"context"
	"errors"
	"log/slog"
	storage_errors "server_crm/internal/storage/errors"
	storage_models "server_crm/internal/storage/models"
)

func (s AutoserviceService) Add(ctx context.Context, dto storage_models.AddAutoserviceDto) (id int64, err error) {

	const op = "service.autoservice.Add"

	log := s.l.With(
		slog.String("op", op),
	)

	log.Info("Start adding autoservice",
		slog.String("name", dto.Name),
		slog.String("address", dto.Address),
		slog.String("phone", dto.Phone),
		slog.Int64("owner_id", dto.Owner_id),
		slog.String("email", dto.Email),
	)

	_, err = s.owP.GetOne(ctx, dto.Owner_id)

	if err != nil {
		if errors.Is(err, storage_errors.ErrOwnerRoleNotFound) {
			log.Error("Owner not found", slog.Any("error", err.Error()))
			return 0, err
		}
		log.Error("Get owner error", slog.Any("error", err.Error()))
		return 0, err
	}

	id, err = s.auP.Add(ctx, dto)

	log.Info("Add autoservice success",
		slog.Int64("id", id),
	)

	if err != nil {
		log.Error("Add autoservice error", slog.Any("error", err.Error()))
		return 0, err
	}

	return id, nil

}
