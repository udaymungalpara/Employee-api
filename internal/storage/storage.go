package storage

type Storage interface {
	CreateEmp(name string, email string, gender string, department string, age int) (int64, error)
}
