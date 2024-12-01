package user_controller

import (
	"errors"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	services_errors "server_crm/internal/services/errors"
)

// Auth Handler for login
// @Summary Login
// @Description Login with email and password
// @Tags User
// @Accept  json
// @Produce  json
// @Param dto body rest_models.LoginReqDto true "Login data"
// @Success 200  "OK"
// @Failure 500  "Error"
// @Router /user/login [post]
func (h UserController) Login(w http.ResponseWriter, r *http.Request) {

	var dto rest_models.LoginReqDto

	aToken, rToken, user, err := h.uc.Login(r.Context(), dto.Email, dto.Password)

	if err != nil {
		if errors.Is(err, services_errors.ErrIncorrectPassword) {
			rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid email or password")
			return
		}
		rest_utils.ErrorsHandler(w, http.StatusInternalServerError, "Internal server error")
	}

	h.setRefreshCookie(w, rToken)

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.LoginResDto{
		Token: aToken,
		User:  h.fromDomainToRest(user),
	})
}
