package catalog_controller

import (
	"encoding/json"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"
)

// Handler for update catalog
// @Summary Update
// @Description Update catalog
// @Tags Catalogs
// @Accept  json
// @Produce  json
// @Param dto body rest_models.UpdateCatalogReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /catalogs/ [PATCH]
func (h CatalogController) Update(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.UpdateCatalogReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	if dto.Id == 0 ||
		dto.Name == "" {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	err = h.catUc.Update(r.Context(), storage_models.UpdateCatalogDto{
		Id:      dto.Id,
		Name:    dto.Name,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	catalog, err := h.catUc.GetById(r.Context(), dto.Id)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, h.fromDomainToRest(catalog))
}
