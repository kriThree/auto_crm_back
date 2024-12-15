package catalog_controller

import (
	"net/http"
	"server_crm/internal/auxiliary"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	"server_crm/internal/services/models"
)

// Handler for get catalogs
// @Summary Get
// @Description Get catalogs
// @Tags Catalogs
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /catalogs/ [GET]
func (h CatalogController) Get(w http.ResponseWriter, r *http.Request) {


	domainCatalogs, err := h.catUc.Get(r.Context())

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	catalogs, err := auxiliary.NewWorker(
		r.Context(),
		domainCatalogs,
		func(ctx auxiliary.Context, a models.Catalog) rest_models.Catalog {
			return h.fromDomainToRest(a)
		},
	)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.GetCatalogRes{
		Catalogs: catalogs,
	})
}
