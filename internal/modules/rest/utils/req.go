package rest_utils

import (
	"encoding/json"
	"net/http"
)

func RequestReturn(w http.ResponseWriter, status int, body any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(body)

} 