package work_service

import (
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

func (s WorkService) fromStorageToDomain(storage_work storage_models.Work) models.Work {
	return models.Work{
		Id:          storage_work.Id,
		Cost:        storage_work.Cost,
		Name:        storage_work.Name,
		CatalogId:   storage_work.CatalogId,
	}
}
