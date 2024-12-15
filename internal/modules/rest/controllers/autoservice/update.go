package autoservice_controller

import (
	"encoding/json"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"
)

// Handler for update autoservice
// @Summary Update
// @Description Update autoservice
// @Tags Autoservices
// @Accept  json
// @Produce  json
// @Param dto body rest_models.UpdateAutoserviceReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /autoservices/ [PATCH]
func (h AutoserviceController) Update(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.UpdateAutoserviceReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	if dto.Id == 0 ||
		dto.Name == "" ||
		dto.Address == "" ||
		dto.Phone == "" ||
		dto.Email == "" {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	err = h.autUc.Update(r.Context(), storage_models.UpdateAutoserviceDto{
		Id:      dto.Id,
		Name:    dto.Name,
		Address: dto.Address,
		Phone:   dto.Phone,
		Email:   dto.Email,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	autoservice, err := h.autUc.GetById(r.Context(), dto.Id)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, h.fromDomainToRest(autoservice))
}
