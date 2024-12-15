package catalog_controller

import (
	rest_models "server_crm/internal/modules/rest/models"
	"server_crm/internal/services/models"
)

func (h CatalogController) fromDomainToRest(catalog models.Catalog) rest_models.Catalog {
	worksInCatalog := make([]rest_models.WorkInCatalog, 0, len(catalog.Works))

	for _, work := range catalog.Works {
		worksInCatalog = append(worksInCatalog, rest_models.WorkInCatalog{
			Id:   work.Id,
			Cost: work.Cost,
			Name: work.Name,
		})
	}
	return rest_models.Catalog{
		Id:        catalog.Id,
		Name:      catalog.Name,
		Works:     worksInCatalog,
	}
}
