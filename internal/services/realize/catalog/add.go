package catalog_service

import (
	"context"
	"errors"
	"log/slog"
	storage_errors "server_crm/internal/storage/errors"
	storage_models "server_crm/internal/storage/models"
)

func (s CatalogService) Add(ctx context.Context, dto storage_models.AddCatalogDto) (id int64, err error) {

	const op = "service.catalog.Add"

	log := s.l.With(
		slog.String("op", op),
		slog.String("name", dto.Name),
		slog.Int64("admin_id", dto.AdminId),
	)

	log.Info("Start adding catalog")

	_, err = s.adP.GetOne(ctx, dto.AdminId)

	if err != nil {
		if errors.Is(err, storage_errors.ErrAdminRoleNotFound) {
			log.Error("Admin not found", slog.Any("error", err.Error()))
			return 0, err
		}
		log.Error("Get admin error", slog.Any("error", err.Error()))
		return 0, err
	}

	id, err = s.ctP.Add(ctx, dto)

	log.Info("Add catalog success",
		slog.Int64("id", id),
	)

	if err != nil {
		log.Error("Add catalog error", slog.Any("error", err.Error()))
		return 0, err
	}

	return id, nil
}
