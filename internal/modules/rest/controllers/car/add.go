package car_controller

import (
	"encoding/json"
	"net/http"
	"server_crm/internal/auxiliary"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"
)

// Handler for add car
// @Summary Add
// @Description Add car with
// @Tags Cars
// @Accept  json
// @Produce  json
// @Param dto body rest_models.AddCarReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /cars/ [post]
func (h CarController) Add(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.AddCarReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}
	clientId,_ := auxiliary.GetUserInfo(r.Context())


	if dto.Number == "" {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	id, err := h.crUc.Add(r.Context(), storage_models.AddCarDto{
		Number:      dto.Number,
		Description: dto.Description,
		ClientId:    clientId,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.AddCarRes{
		Id: id,
	})
}
