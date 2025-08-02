package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"

	types "github.com/udaymungalpara/employee-api/internal/Types"
	"github.com/udaymungalpara/employee-api/internal/config"
)

type Sqlite struct {
	db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {

	db, err := sql.Open("sqlite", cfg.Storage_path)

	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS employee(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			email TEXT,
			gender TEXT,
			department TEXT,
			age INTEGER
		)`)

	if err != nil {
		return nil, err
	}
	return &Sqlite{
		db: db,
	}, nil
}

func (s *Sqlite) CreateEmp(name string, email string, gender string, department string, age int) (int64, error) {

	stmt, err := s.db.Prepare("INSERT INTO employee (name,email,gender,department,age) VALUES (? ,? ,? , ? ,?)")

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(name, email, gender, department, age)

	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()

	return id, nil

}

func (s *Sqlite) GetbyId(id int) (types.Employee, error) {
	stmt, err := s.db.Prepare("SELECT * FROM employee by WHERE id=?")

	if err != nil {
		return types.Employee{}, err
	}

	defer stmt.Close()

	res := stmt.QueryRow(id)

	var emp types.Employee

	err = res.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Gender, &emp.Department, &emp.Age)
	if err != nil {
		return types.Employee{}, err
	}

	return emp, nil

}

func (s *Sqlite) GetList() ([]types.Employee, error) {

	employee := []types.Employee{}
	stmt, err := s.db.Prepare(`SELECT * FROM employee`)

	if err != nil {
		return employee, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return employee, err
	}
	defer rows.Close()
	for rows.Next() {
		var emp types.Employee
		err := rows.Scan(&emp.Id, &emp.Name, &emp.Email, &emp.Gender, &emp.Department, &emp.Age)

		if err != nil {
			return employee, err
		}
		employee = append(employee, emp)
	}

	return employee, nil

}
