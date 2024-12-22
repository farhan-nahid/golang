package storage

import "github.com/farhan-nahid/golang/students-api/internal/types"

type Storage interface {
	CreateStudent(firstName string, lastName string, age int, email string) (int64, error)
	GetStudentByID(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudentByID(id int64, student types.Student) (types.Student, error)
}

