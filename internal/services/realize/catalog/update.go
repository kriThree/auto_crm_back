package catalog_service

import (
	"context"
	"log/slog"
	storage_models "server_crm/internal/storage/models"
)

func (s CatalogService) Update(ctx context.Context, dto storage_models.UpdateCatalogDto) error {

	const op = "service.catalog.Update"

	log := s.l.With(
		slog.String("op", op),
	)
	log.Info("Start updating catalog",
		slog.String("number", dto.Name),
		slog.Int64("admin_id", dto.AdminId),
		slog.Int64("id", dto.Id),
	)

	err := s.ctP.Update(ctx, dto)

	if err != nil {
		log.Error("Update catalog error", slog.Any("error", err.Error()))
		return err
	}

	log.Info("Update catalog success",
		slog.Int64("id", dto.Id))

	return nil
}
