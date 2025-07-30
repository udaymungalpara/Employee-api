package sqlite

import (
	"database/sql"

	_ "modernc.org/sqlite"

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
