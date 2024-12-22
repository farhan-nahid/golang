package sqlite

import (
	"database/sql"

	"github.com/farhan-nahid/golang/students-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type SQLite struct {
	DB *sql.DB
}

// GetStudent implements storage.Storage.
func (s *SQLite) GetStudent() error {
	panic("unimplemented")
}

func New(cfg config.Config) (*SQLite, error) {
	db, err := sql.Open("sqlite3", cfg.Storage_path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		firstName TEXT NOT NULL,
		lastName TEXT NOT NULL,
		email TEXT NOT NULL,
		age INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP 
	)`)

	if err != nil {
		return nil, err
	}

	return &SQLite{DB: db}, nil
}

func (s *SQLite) CreateStudent(firstName string, lastName string, age int, email string) (int64, error) {

	stmt, err := s.DB.Prepare(`INSERT INTO students (firstName, lastName, age, email) VALUES (?, ?, ?, ?)`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(firstName, lastName, age, email)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}
