package main

import "fmt"

func add (a, b int) int{
	return a + b
}

func returnMultiple ()(string, int, bool){
	return "golang", 5, true
}

func process (fn func(a int) int) int{
	return fn(0) 
}

func processIt () func (a int) int{
	return func(a int) int{
		return 3
	}
}

func main(){
	one, two, _ := returnMultiple()
	add(1, 3)
	fmt.Println(one, two)
}