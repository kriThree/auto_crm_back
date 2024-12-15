package user_controller

import (
	"encoding/json"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	"server_crm/internal/services/models"
)

// Auth Handler for registration
// @Summary Registration
// @Description Register new user
// @Tags User
// @Accept  json
// @Produce  json
// @Param dto body rest_models.RegisterReqDto true "Registration data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /user/register [post]
func (h UserController) Register(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.RegisterReqDto

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid format of request body")
		return
	}

	if dto.Email == "" || dto.Password == "" || dto.Name == "" || dto.Role == ""{
		rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid values of request body")
		return
	}

	aToken, rToken, user, err := h.uc.Register(r.Context(), models.RegisterUserDto{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
		Role:     dto.Role,
	})

	if err != nil {
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	h.setRefreshCookie(w, rToken)

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.RegisterResDto{
		Token: aToken,
		User : h.fromDomainToRest(user),
	})
}
