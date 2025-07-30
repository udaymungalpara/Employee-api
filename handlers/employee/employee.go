package employee

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

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

		respones.WriteJson(w, http.StatusOK, map[string]string{
			"status": "ok",
		})

	}

}
