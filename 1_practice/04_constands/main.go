package main

import "fmt"

const PI = 0.8086
func main (){
	const NAME string = "golang"

	// name = "JS"

	fmt.Println(NAME)
	fmt.Println(PI)

	const (
		PORT = 8080
		HOST = "localhost"
	)

	fmt.Print(PORT)
	fmt.Print(HOST)
}