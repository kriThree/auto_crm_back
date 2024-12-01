package user_controller

import (
	"errors"
	"net/http"
	rest_models "server_crm/internal/modules/rest/models"
	rest_utils "server_crm/internal/modules/rest/utils"
	services_errors "server_crm/internal/services/errors"
	"strings"
)

func (h UserController) Authorize(w http.ResponseWriter, r *http.Request) {

	accessTokenArr := strings.Split(r.Header.Get("Authorization"), " ")

	accessToken := ""

	if len(accessTokenArr) > 1 {
		accessToken = accessTokenArr[1]
	}

	refreshToken := r.Header.Get("refresh_token")

	accessToken, refreshToken, user, err := h.uc.Authorize(r.Context(), accessToken, refreshToken)

	if err != nil {
		if errors.Is(err, services_errors.ErrIncorrectAuthToken) {
			rest_utils.ErrorsHandler(w, http.StatusBadRequest, "Not valid token")
			return
		}
		if errors.Is(err, services_errors.ErrNoAuthorizationTokens) {
			rest_utils.ErrorsHandler(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		rest_utils.ErrorsHandler(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	h.setRefreshCookie(w, refreshToken)

	rest_utils.RequestReturn(w, http.StatusOK, rest_models.AuthorizeResDto{
		Token: accessToken,
		User:  h.fromDomainToRest(user),
	})
}
