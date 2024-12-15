package catalog_controller

import (
	"encoding/json"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
)

// Handler for delete catalog
// @Summary Delete
// @Description Delete catalog
// @Tags Catalogs
// @Accept  json
// @Produce  json
// @Param dto body rest_models.DeleteCatalogReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /catalogs/ [DELETE]
func (h CatalogController) Delete(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.DeleteCatalogReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	if dto.Id == 0 {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	err = h.catUc.Delete(r.Context(), dto.Id)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, nil)
}
