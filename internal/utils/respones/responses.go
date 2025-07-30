package respones

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
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

func ValidationError(errs validator.ValidationErrors) ErrorResponse {

	var errmsg []string

	for _, err := range errs {
		switch err.ActualTag() {
		case "required":
			errmsg = append(errmsg, fmt.Sprintf("the %s field is required", err.Field()))

		default:
			errmsg = append(errmsg, fmt.Sprintf("the %s field is invalid", err.Field()))
		}

	}

	return ErrorResponse{
		Status: StatusError,
		Error:  strings.Join(errmsg, ","),
	}

}
