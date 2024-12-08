package autoservice_service

import (
	"context"
	"log/slog"
	storage_models "server_crm/internal/storage/models"
)

func (s AutoserviceService) Update(ctx context.Context, dto storage_models.UpdateAutoserviceDto) error {

	const op = "service.autoservice.Update"

	log := s.l.With(
		slog.String("op", op),
	)
	log.Info("Start updating autoservice",
		slog.String("name", dto.Name),
		slog.String("address", dto.Address),
		slog.String("phone", dto.Phone),
		slog.String("email", dto.Email),
	)

	err := s.auP.Update(ctx, dto)

	if err != nil {
		log.Error("Update autoservice error", slog.Any("error", err.Error()))
		return err
	}

	log.Info("Update autoservice success",
		slog.Int64("id", dto.Id))

	return nil
}
