package middlewares

import (
	"errors"
	"net/http"
	"server_crm/internal/auxiliary"
	"server_crm/internal/modules/rest/utils"
	storage_models "server_crm/internal/storage/models"

	"github.com/gorilla/mux"
)

func CheckRole(roles []string) (mux.MiddlewareFunc, error) {

	for _, role := range roles {
		if !(role == storage_models.ROLE_ADMIN ||
			role == storage_models.ROLE_OWNER ||
			role == storage_models.ROLE_CLIENT) {
			return nil, errors.New("not valid role")
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			_, role := auxiliary.GetUserInfo(r.Context())

			for _, roleI := range roles {
				if role == roleI {
					next.ServeHTTP(w, r)
					return
				}
			}
			rest_utils.ErrorsHandler(w, http.StatusUnauthorized, "Unauthorized")
		})
	}, nil
}
