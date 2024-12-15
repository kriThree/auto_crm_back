package catalog_controller

import (
	"encoding/json"
	"net/http"
	"server_crm/internal/auxiliary"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"

)

// Handler for add catalog
// @Summary Add
// @Description Add catalog with
// @Tags Catalogs
// @Accept  json
// @Produce  json
// @Param dto body rest_models.AddCatalogReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /catalogs/ [post]
func (h CatalogController) Add(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.AddCatalogReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	adminId, _ := auxiliary.GetUserInfo(r.Context())

	if dto.Name == "" {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	id, err := h.catUc.Add(r.Context(), storage_models.AddCatalogDto{
		Name:    dto.Name,
		AdminId: adminId,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.AddCatalogRes{
		Id: id,
	})
}
