package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"server_crm/internal/auxiliary"
	rest_errors "server_crm/internal/modules/rest/utils"
	"strings"

	"github.com/gorilla/mux"
)

type UserDecrypter interface {
	Validate(ctx context.Context, tokenStr string) (userId int64, role string, err error)
}

func AuthMiddleware(usC UserDecrypter) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := strings.Split(r.Header.Get("Authorization"), " ")
			if len(tokenStr) < 2 {
				next.ServeHTTP(w, r)
				return
			}
			fmt.Println(tokenStr)
			userId,role, err := usC.Validate(r.Context(), tokenStr[1])
			if err != nil {
				rest_errors.ErrorsHandler(w, http.StatusUnauthorized, "Unauthorized")
				return
			}
			oldContext := r.Context()

			auxiliary.SetUserInfo(&oldContext, userId,role)
			next.ServeHTTP(w, r.WithContext(oldContext))
		})
	}
}
