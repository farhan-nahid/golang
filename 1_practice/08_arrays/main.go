package main

import "fmt"


func main(){
	// single line
	// num:= [5]int {0,1,2,3,4,}

	var num [5]int

	// array length
	fmt.Println(len(num))

	// add in array
	num[0] = 1
	num[4] = 6

	// get value
	fmt.Println(num)
	fmt.Println(num[0])

	// 2D arrays
	numbers := [3][3]int{{1, 2, 3}, {4,5,6}, {7,8,9}}
	fmt.Println(numbers)
}