package rest_utils

import (
	"encoding/json"
	"net/http"
)

func RequestReturn( w http.ResponseWriter, status int, body any) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(body)

} 