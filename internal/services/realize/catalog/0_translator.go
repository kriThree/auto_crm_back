package catalog_service

import (
	"context"
	"fmt"
	"server_crm/internal/auxiliary"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

func (s CatalogService) fromStorageToDomainWithQuery(
	ctx context.Context,
	storageCatalog storage_models.Catalog,
) (models.Catalog, error) {

	const op = "service.catalog.fromStorageToDomainWithQuery"

	storageWorks, err := s.wkP.GetByCatalogId(ctx, storageCatalog.Id)

	if err != nil {
		return models.Catalog{}, fmt.Errorf("%s: %w", op, err)
	}

	return s.fromStorageToDomain(storageCatalog, storageWorks), nil
}
func (s CatalogService) fromStorageToDomain(storageCatalog storage_models.Catalog, storageWorks []storage_models.Work) models.Catalog {

	works, _ := auxiliary.NewWorker(
		context.Background(),
		storageWorks,
		func(aCtx auxiliary.Context, storageWork storage_models.Work) models.Work {
			return models.Work{
				Id:   storageWork.Id,
				Cost: storageWork.Cost,
				Name: storageWork.Name,
			}
		},
	)

	return models.Catalog{
		Id:        storageCatalog.Id,
		Name:      storageCatalog.Name,
		CreaterId: storageCatalog.CreaterId,
		Works:     works,
	}
}
