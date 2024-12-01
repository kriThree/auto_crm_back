package rest_utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ReqError struct {
	Status  int               `json:"status"`
	Text    string            `json:"text"`
	Details map[string]string `json:"details,omitempty"`
}

func ValidateErrorsHandler(w http.ResponseWriter, err error, status int, text string) error {
	var details map[string]string
	if ve, ok := err.(validator.ValidationErrors); ok {
		details = make(map[string]string)
		for _, fe := range ve {
			field := fe.Field()
			tag := fe.Tag()
			details[field] = "failed on the " + tag + " tag"
		}
	} else {
		details = map[string]string{"error": err.Error()}
	}

	return json.NewEncoder(w).Encode(ReqError{
		Status: status,
		Text:   text,
		Details: map[string]string{
			"error": text,
		},
	})
}
func ErrorsHandler(w http.ResponseWriter, status int, text string) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(ReqError{
		Status: status,
		Text:   text,
		Details: map[string]string{
			"error": text,
		},
	})
}
