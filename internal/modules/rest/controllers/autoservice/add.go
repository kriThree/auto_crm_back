package autoservice_controller

import (
	"encoding/json"
	"net/http"
	"server_crm/internal/auxiliary"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"
)

// Handler for add autoservice
// @Summary Add
// @Description Add autoservice with
// @Tags Autoservices
// @Accept  json
// @Produce  json
// @Param dto body rest_models.AddAutoserviceReq true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /autoservices/ [post]
func (h AutoserviceController) Add(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.AddAutoserviceReq

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	ownerId, _ := auxiliary.GetUserInfo(r.Context())


	if dto.Email == "" || dto.Name == "" || dto.Address == "" || dto.Phone == "" {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	id, err := h.autUc.Add(r.Context(), storage_models.AddAutoserviceDto{
		Name:    dto.Name,
		Email:   dto.Email,
		Phone:   dto.Phone,
		Address: dto.Address,
		Owner_id: ownerId,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.AddAutoserviceRes{
		Id: id,
	})
}
