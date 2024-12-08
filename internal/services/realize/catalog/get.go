package catalog_service

import (
	"context"
	"log/slog"
	"server_crm/internal/services/models"
)

func (s CatalogService) Get(ctx context.Context) ([]models.Catalog, error) {

	const op = "service.catalog.Get"

	log := s.l.With(
		slog.String("op", op),
	)

	log.Info("Start getting catalog")

	storageCatalogs, err := s.ctP.Get(ctx)
	if err != nil {
		log.Error("Get catalogs error", slog.Any("error", err.Error()))
		return nil, err
	}

	catalogs := make([]models.Catalog, 0, len(storageCatalogs))

	for _, catalog := range storageCatalogs {
		works,err := s.wkP.GetByCatalogId(ctx, catalog.Id) 

		if err != nil {
			log.Error("Get works error", slog.Int64("id", catalog.Id), slog.Any("error", err.Error()))
			return nil, err
		}

		catalogs = append(catalogs, s.fromStorageToDomain(catalog,works))
	}

	log.Info("Get catalogs success",
		slog.Int("count", len(catalogs)),
	)

	return catalogs, nil
}
func (s CatalogService) GetForAdmin(ctx context.Context, adminId int64) ([]models.Catalog, error) {
	return []models.Catalog{}, nil
}
