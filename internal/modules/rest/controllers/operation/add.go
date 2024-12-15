package operation_controller

import (
	"encoding/json"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"
)

// Handler for add operation
// @Summary Add
// @Description Add operation with
// @Tags Operations
// @Accept  json
// @Produce  json
// @Param dto body rest_models.AddOperationReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /operations/ [post]
func (h OperationController) Add(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.AddOperationReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	if dto.CarID == 0 || dto.WorkID == 0 || dto.AutoserviceId == 0 {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	id, err := h.opUc.Create(r.Context(), storage_models.AddOperationDto{
		CarId:      dto.CarID,
		WorkId:     dto.WorkID,
		AutoserviceId: dto.AutoserviceId,
		Description: dto.Description,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.AddOperationRes{
		Id: id,
	})
}
