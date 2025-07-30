package respones

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOk    = "ok"
	StatusError = "error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func JsonError(err error) ErrorResponse {

	return ErrorResponse{
		Status: StatusError,
		Error:  err.Error(),
	}

}