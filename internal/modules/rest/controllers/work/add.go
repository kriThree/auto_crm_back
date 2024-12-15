package work_controller

import (
	"encoding/json"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"
)

// Handler for add work
// @Summary Add
// @Description Add work with
// @Tags Works
// @Accept  json
// @Produce  json
// @Param dto body rest_models.AddWorkReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /works/ [post]
func (h WorkController) Add(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.AddWorkReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	if dto.Name == "" || dto.Cost == 0 || dto.CatalogId == 0 {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	id, err := h.wkUc.Add(r.Context(), storage_models.AddWorkDto{
		Name:      dto.Name,
		Cost:      dto.Cost,
		CatalogId: dto.CatalogId,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.AddWorkRes{
		Id: id,
	})
}
