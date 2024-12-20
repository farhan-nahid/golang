package main

import "fmt"

func main(){
	var name string = "golang"
	var age int = 24
	var isMarried bool = false
	var isComplete bool

	fmt.Println("isComplete", isComplete)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)

	// shorthand syntax
	university:= "NUB"
	fmt.Println(university)

	isComplete = true
	fmt.Println("isComplete", isComplete)

}