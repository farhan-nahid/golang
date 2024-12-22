package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/farhan-nahid/golang/students-api/internal/config"
	"github.com/farhan-nahid/golang/students-api/internal/types"
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

func (s *SQLite) GetStudentByID(id int64) (types.Student, error) {
	var student types.Student

	stmt, err := s.DB.Prepare(`SELECT * FROM students WHERE id = ? LIMIT 1`)

	if err != nil {
		return student, err
	}

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Age, &student.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return student, fmt.Errorf("student not found with id: %d", id)
		}

		return student, fmt.Errorf("failed to get student: %v", err)
	}

	return student, nil
}

func (s *SQLite) GetStudents() ([]types.Student, error) {
	var students []types.Student

	rows, err := s.DB.Query(`SELECT id, firstName, lastName,  email, age, created_at FROM students`)

	if err != nil {
		return students, err
	}

	for rows.Next() {
		var student types.Student

		err = rows.Scan(&student.ID, &student.FirstName, &student.LastName, &student.Email, &student.Age, &student.CreatedAt)

		if err != nil {
			return students, err
		}

		students = append(students, student)
	}

	return students, nil
}