package main

import "fmt"

func main(){
	age:=16

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 && age <= 18{
		fmt.Println("Person is an teenager")
	} else {
		fmt.Println("Person is a kid")
	}


	var role = "admin"
	var permission = false
	
	if role == "admin" || permission{
		fmt.Println("Admin or permission true")
	}


	// variable inside if statement
	if age:=16; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 && age <= 18{
		fmt.Println("Person is an teenager", age)
	} else {
		fmt.Println("Person is a kid",age)
	}
}