package catalog_service

import (
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

func (s CatalogService) fromStorageToDomain(storage_catalog storage_models.Catalog, storage_works []storage_models.Work) models.Catalog {
	
	works := make([]models.Work, 0, len(storage_works))

	for i, storageWork := range storage_works {
		works[i] = models.Work{
			Id:        storageWork.Id,
			Cost:      storageWork.Cost,
			Name:      storageWork.Name,
			CatalogId: storageWork.CatalogId,
		}
	}

	return models.Catalog{
		Id:        storage_catalog.Id,
		Name:      storage_catalog.Name,
		CreaterId: storage_catalog.CreaterId,
		Works:     works,
	}
}
