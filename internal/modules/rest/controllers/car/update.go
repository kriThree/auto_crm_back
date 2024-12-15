package car_controller

import (
	"encoding/json"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"
)

// Handler for update car
// @Summary Update
// @Description Update car
// @Tags Cars
// @Accept  json
// @Produce  json
// @Param dto body rest_models.UpdateCarReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /cars/ [PATCH]
func (h CarController) Update(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.UpdateCarReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	if dto.Id == 0 ||
		dto.Description == "" ||
		dto.Number == "" {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	err = h.crUc.Update(r.Context(), storage_models.UpdateCarDto{
		Id:   dto.Id,
		Number: dto.Number,
		Description: dto.Description,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	car, err := h.crUc.GetById(r.Context(), dto.Id)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, h.fromDomainToRest(car))
}
