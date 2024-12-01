package user_controller

import (
	"net/http"
	"time"
)

func ( h UserController) setRefreshCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		HttpOnly: true,
	})
}
