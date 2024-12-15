package operation_controller

import (
	"encoding/json"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
)

// Handler for delete operation
// @Summary Delete
// @Description Delete operation
// @Tags Operations
// @Accept  json
// @Produce  json
// @Param dto body rest_models.DeleteOperationReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /operations/ [DELETE]
func (h OperationController) Delete(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.DeleteOperationReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	if dto.Id == 0 {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	err = h.opUc.Delete(r.Context(), dto.Id)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, nil)
}
