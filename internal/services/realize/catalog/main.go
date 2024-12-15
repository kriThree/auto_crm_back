package catalog_service

import (
	"log/slog"
	"server_crm/internal/services/models"
	storage_models "server_crm/internal/storage/models"
)

type CatalogService struct {
	ctP CatalogProvider
	adP AdminProvider
	wkP workProvider
	l   *slog.Logger
}
type AdminProvider = storage_models.RoleDomain
type CatalogProvider = storage_models.CatalogDomain
type workProvider = storage_models.WorkDomain

func New(catalogProvider CatalogProvider, adminProvider AdminProvider,workProvider workProvider,l *slog.Logger) models.CatalogRepo {
	return CatalogService{
		ctP: catalogProvider,
		l:   l,
		adP: adminProvider,
		wkP: workProvider,
	}
}
