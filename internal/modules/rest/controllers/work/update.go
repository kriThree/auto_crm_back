package work_controller

import (
	"encoding/json"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"
)

// Handler for update work
// @Summary Update
// @Description Update work
// @Tags Works
// @Accept  json
// @Produce  json
// @Param dto body rest_models.UpdateWorkReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /works/ [PATCH]
func (h WorkController) Update(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.UpdateWorkReq

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

	err = h.wkUc.Update(r.Context(), storage_models.UpdateWorkDto{
		Id:   dto.Id,
		Name: dto.Name,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	work, err := h.wkUc.GetById(r.Context(), dto.Id)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, h.fromDomainToRest(work))
}
