package catalog_service

import (
	"context"
	"log/slog"
	"server_crm/internal/auxiliary"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
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

	catalogs, err := auxiliary.NewWorker(
		ctx,
		storageCatalogs,
		func(aCtx auxiliary.Context, storageCatalog storage_models.Catalog) models.Catalog {
			a, err := s.fromStorageToDomainWithQuery(ctx, storageCatalog)
			if err != nil {
				aCtx.PushError(err)
				return models.Catalog{}
			}
			return a
		},
	)

	if err != nil {
		log.Error("Get works error", slog.Any("error", err.Error()))
		return nil, err
	}

	log.Info("Get catalogs success",
		slog.Int("count", len(catalogs)),
	)

	return catalogs, nil
}
func (s CatalogService) GetForAdmin(ctx context.Context, adminId int64) ([]models.Catalog, error) {

	const op = "service.catalog.GetForAdmin"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("admin_id", adminId),
	)

	log.Info("Start getting catalog for admin")

	storageCatalogs, err := s.ctP.GetByAdminId(ctx, adminId)

	if err != nil {
		log.Error("Get catalogs for admin error", slog.Int64("admin_id", adminId), slog.Any("error", err.Error()))
		return nil, err
	}

	catalogs, err := auxiliary.NewWorker(
		ctx,
		storageCatalogs,
		func(aCtx auxiliary.Context, storageCatalog storage_models.Catalog) models.Catalog {
			a, err := s.fromStorageToDomainWithQuery(ctx, storageCatalog)
			if err != nil {
				aCtx.PushError(err)
				return models.Catalog{}
			}
			return a
		},
	)

	if err != nil {
		log.Error("Get works error", slog.Any("error", err.Error()))
		return nil, err
	}

	log.Info("Get catalogs for admin success",
		slog.Int("count", len(storageCatalogs)),
	)

	return catalogs, nil
}

func (s CatalogService) GetById(ctx context.Context, id int64) (models.Catalog, error) {

	const op = "service.catalog.GetById"

	log := s.l.With(
		slog.String("op", op),
		slog.Int64("id", id),
	)

	log.Info("Start getting catalog by id")

	catalog, err := s.ctP.GetById(ctx, id)

	if err != nil {
		log.Error("Get catalog by id error", slog.Int64("id", id), slog.Any("error", err.Error()))
		return models.Catalog{}, err
	}

	log.Info("Get catalog by id success",
		slog.Int64("id", id),
	)

	return s.fromStorageToDomainWithQuery(ctx, catalog)

}
