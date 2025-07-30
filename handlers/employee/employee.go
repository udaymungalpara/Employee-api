package employee

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
	types "github.com/udaymungalpara/employee-api/internal/Types"
	"github.com/udaymungalpara/employee-api/internal/utils/respones"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var employee types.Employee

		err := json.NewDecoder(r.Body).Decode(&employee)

		if errors.Is(err, io.EOF) {
			respones.WriteJson(w, http.StatusBadRequest, respones.JsonError(err))
			return
		}

		if err != nil {
			respones.WriteJson(w, http.StatusBadRequest, respones.JsonError(err))
			return
		}

		//validate request
		if err := validator.New().Struct(employee); err != nil {

			verr := err.(validator.ValidationErrors)
			respones.WriteJson(w, http.StatusBadRequest, respones.ValidationError(verr))
			return
		}

		respones.WriteJson(w, http.StatusOK, map[string]string{
			"status": "ok",
		})

	}

}
