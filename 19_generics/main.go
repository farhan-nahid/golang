package main

import "fmt"

// Sum function with generics
func Sum[T int | float64](numbers []T) T {
    var total T
    for _, number := range numbers {
        total += number
    }
    return total
}

func print[T comparable](data []T)  {
	for _, d := range data {
       fmt.Println(d)
    }
}

// Define the struct
type User struct {
    ID        int
    FirstName string
    LastName  string
    Email     string
}

// Method to display the full name
func (u User) FullName() string {
    return u.FirstName + " " + u.LastName
}

// Method to update the email
func (u *User) UpdateEmail(newEmail string) {
    u.Email = newEmail
}

func main() {
	  // Create a new user
	user := User{
        ID:        1,
        FirstName: "John",
        LastName:  "Doe",
        Email:     "john.doe@example.com",
    }

    intNumbers := []int{1, 2, 3, 4, 5}
    floatNumbers := []float64{1.5, 2.3, 3.7, 4.1}

    fmt.Printf("Sum of integers: %d\n", Sum(intNumbers))
    fmt.Printf("Sum of floats: %.2f\n", Sum(floatNumbers))
	

    // Access and display user details
    fmt.Printf("User ID: %d\n", user.ID)
    fmt.Printf("User Name: %s\n", user.FullName())
    fmt.Printf("User Email: %s\n", user.Email)

    // Update the user's email
    user.UpdateEmail("john.newemail@example.com")
    fmt.Printf("Updated User Email: %s\n", user.Email)
}
