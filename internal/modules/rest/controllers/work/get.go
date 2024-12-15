package work_controller

import (
	"net/http"
	"server_crm/internal/auxiliary"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	"server_crm/internal/services/models"
)

// Handler for get works
// @Summary Get
// @Description Get works
// @Tags Works
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /works/ [GET]
func (h WorkController) Get(w http.ResponseWriter, r *http.Request) {

	domainWorks, err := h.wkUc.Get(r.Context())

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	works, err := auxiliary.NewWorker[models.Work, rest_models.Work](
		r.Context(),
		domainWorks,
		func(ctx auxiliary.Context, a models.Work) rest_models.Work {
			return h.fromDomainToRest(a)
		},
	)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.GetWorksRes{
		Works: works,
	})
}
