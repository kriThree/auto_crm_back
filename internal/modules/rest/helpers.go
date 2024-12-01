package rest

import "net/http"

func OptionsOK(w http.ResponseWriter, r *http.Request) {
	// Установка заголовков CORS

	w.WriteHeader(http.StatusOK)
}
