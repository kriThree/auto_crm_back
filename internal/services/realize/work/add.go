package work_service

import (
	"context"
	"log/slog"
	storage_models "server_crm/internal/storage/models"
)

func (s WorkService) Add(ctx context.Context, dto storage_models.AddWorkDto) (id int64, err error) {

	const op = "service.work.Add"

	log := s.l.With(
		slog.String("op", op),
		slog.String("name", dto.Name),
		slog.Int("cost", dto.Cost),
		slog.Int64("catalog_id", dto.CatalogId),
	)

	log.Info("Start adding work")

	id, err = s.wkP.Add(ctx, dto)

	log.Info("Add work success",
		slog.Int64("id", id),
	)

	if err != nil {
		log.Error("Add work error", slog.Any("error", err.Error()))
		return 0, err
	}

	return id, nil
}
