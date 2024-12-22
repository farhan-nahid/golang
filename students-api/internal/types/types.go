package types


type Student struct {
	ID        int    
	FirstName string  `validate:"required"`
	LastName  string  `validate:"required"`
	Age       int     `validate:"required"`
	Email     string  `validate:"required,email"`
}