package operation_controller

import (
	"encoding/json"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"
)

// Handler for update operation
// @Summary Update
// @Description Update operation
// @Tags Operations
// @Accept  json
// @Produce  json
// @Param dto body rest_models.UpdateOperationReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /operations/ [PATCH]
func (h OperationController) Update(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.UpdateOperationReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	if dto.CarId == 0 ||
		dto.Id == 0 ||
		dto.WorkId == 0 ||
		dto.AutoserviceId == 0 {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	err = h.opUc.Update(r.Context(), storage_models.UpdateOperationDto{
		Id:   dto.Id,
		Description: dto.Description,
		CarId:       dto.CarId,
		WorkId:      dto.WorkId,
		AutoserviceId: dto.AutoserviceId,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	operation, err := h.opUc.GetById(r.Context(), dto.Id)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, h.fromDomainToRest(operation))
}
