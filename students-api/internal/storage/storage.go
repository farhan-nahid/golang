package storage

type Storage interface {
	CreateStudent(firstName string, lastName string, age int, email string) (int64, error)
	// GetStudent() error
}

