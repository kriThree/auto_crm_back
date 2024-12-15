package car_controller

import (
	"net/http"
	"server_crm/internal/auxiliary"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	"server_crm/internal/services/models"
)

// Handler for get cars
// @Summary Get
// @Description Get cars
// @Tags Cars
// @Accept  json
// @Produce  json
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /cars/ [GET]
func (h CarController) Get(w http.ResponseWriter, r *http.Request) {

	userId, _ := auxiliary.GetUserInfo(r.Context())

	domainCars, err := h.crUc.Get(r.Context(), userId)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	cars, err := auxiliary.NewWorker[models.Car, rest_models.Car](
		r.Context(),
		domainCars,
		func(ctx auxiliary.Context, a models.Car) rest_models.Car {
			return h.fromDomainToRest(a)
		},
	)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.GetCarsRes{
		Cars: cars,
	})
}
