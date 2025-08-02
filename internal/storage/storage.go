package storage

import types "github.com/udaymungalpara/employee-api/internal/Types"

type Storage interface {
	CreateEmp(name string, email string, gender string, department string, age int) (int64, error)
	GetbyId(id int) (types.Employee, error)
	GetList() ([]types.Employee, error)

	DeleteById(id int) error
}
