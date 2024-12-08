package work_service

import (
	"context"
	"log/slog"
	storage_models "server_crm/internal/storage/models"
)

func (s WorkService) Update(ctx context.Context, dto storage_models.UpdateWorkDto) error {

	const op = "service.work.Update"

	log := s.l.With(
		slog.String("op", op),
	)
	log.Info("Start updating work",
		slog.String("name", dto.Name),
		slog.Int("cost", dto.Cost),
		slog.Int64("catalog_id", dto.CatalogId),
		slog.Int64("id", dto.Id),
	)

	err := s.wkP.Update(ctx, dto)

	if err != nil {
		log.Error("Update work error", slog.Any("error", err.Error()))
		return err
	}

	log.Info("Update work success",
		slog.Int64("id", dto.Id))

	return nil
}
