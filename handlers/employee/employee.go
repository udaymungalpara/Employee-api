package employee

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	types "github.com/udaymungalpara/employee-api/internal/Types"
	"github.com/udaymungalpara/employee-api/internal/storage"
	"github.com/udaymungalpara/employee-api/internal/utils/respones"
)

func New(storage storage.Storage) http.HandlerFunc {
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

		id, err := storage.CreateEmp(employee.Name, employee.Email, employee.Gender, employee.Department, employee.Age)

		if err != nil {
			respones.WriteJson(w, http.StatusInternalServerError, respones.JsonError(err))
			return
		}

		respones.WriteJson(w, http.StatusOK, respones.DoneJson(id))

	}

}

func GetId(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var employee types.Employee

		getid := r.PathValue("id")

		id, _ := strconv.Atoi(getid)

		employee, err := storage.Storage.GetbyId(s, id)
		if err != nil {
			respones.WriteJson(w, http.StatusInternalServerError, respones.JsonError(err))
			return
		}
		respones.WriteJson(w, http.StatusOK, &employee)

	}
}

func GetList(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		emps, err := storage.Storage.GetList(s)

		if err != nil {
			respones.WriteJson(w, http.StatusInternalServerError, respones.JsonError(err))
			return
		}

		respones.WriteJson(w, http.StatusOK, &emps)

	}
}

func DeleteById(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		eid, err := strconv.Atoi(id)

		if err != nil {
			respones.WriteJson(w, http.StatusBadRequest, respones.JsonError(err))
		}

		err = storage.Storage.DeleteById(s, eid)
		
		if err != nil {
			respones.WriteJson(w, http.StatusInternalServerError, respones.JsonError(err))
			return
		}
		respones.WriteJson(w, http.StatusOK, map[string]string{"id": id, "Respone": "Employee Deleted"})

	}
}
