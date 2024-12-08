package catalog_service

import (
	"context"
	"fmt"
	"log/slog"
)

func (s CatalogService) Delete(ctx context.Context, id int64) error {

	const op = "service.catalog.Delete"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("id", id),
	)

	log.Info("Start to delete catalog")

	err := s.ctP.Delete(ctx, id)

	if err != nil {
		log.Error("Delete catalog error", slog.Any("error", err.Error()))
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("Delete catalog success")

	return nil
}
