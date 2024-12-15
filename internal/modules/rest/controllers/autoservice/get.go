package autoservice_controller

import (
	"net/http"
	"server_crm/internal/auxiliary"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	"server_crm/internal/services/models"
)

// Handler for get autoservices
// @Summary Get
// @Description Get autoservices
// @Tags Autoservices
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /autoservices/ [GET]
func (h AutoserviceController) Get(w http.ResponseWriter, r *http.Request) {

	userId, _ := auxiliary.GetUserInfo(r.Context())

	domainAutoservices, err := h.autUc.Get(r.Context(), userId)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}


	autoservices,err := auxiliary.NewWorker[models.Autoservice, rest_models.Autoservice](
		r.Context(),
		domainAutoservices,
		func(ctx auxiliary.Context, a models.Autoservice) rest_models.Autoservice {
			return h.fromDomainToRest(a)
		},
	)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	
	rest_utils.RequestReturn(w, http.StatusOK, rest_models.GetAutoserviceRes{
		Autoservices: autoservices,
	})
}
