package types


type Student struct {
	ID        int     `json:"id"`  
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	Age       int     `json:"age" validate:"required"`
	Email     string  `json:"email" validate:"required,email"`
	CreatedAt string  `json:"created_at"`
}