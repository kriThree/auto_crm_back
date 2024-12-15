package operation_controller

import (
	"net/http"
	"server_crm/internal/auxiliary"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	"server_crm/internal/services/models"

)

// Handler for get operations
// @Summary Get
// @Description Get operations
// @Tags Operations
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /operations/ [GET]
func (h OperationController) Get(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.GetOperationsReq

	var domainOperations []models.Operation
	err := error(nil)
	if dto.AutoserviceId != 0 {
		domainOperations,err = h.opUc.GetForAutoservice(r.Context(), dto.AutoserviceId)
	} else if dto.WorkId != 0 {
		domainOperations, err = h.opUc.GetForWork(r.Context(), dto.WorkId)
	} else if dto.CarId != 0 {
		domainOperations, err = h.opUc.GetForCar(r.Context(), dto.CarId)
	} else {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}


	operations, err := auxiliary.NewWorker(
		r.Context(),
		domainOperations,
		func(ctx auxiliary.Context, a models.Operation) rest_models.Operation {
			return h.fromDomainToRest(a)
		},
	)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.GetOperationsRes{
		Operations: operations,
	})
}
