package main

import "fmt"

func main(){
	var numbers = make([]int, 0, 5)
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	numbers = append(numbers, 3)
	numbers = append(numbers, 4)

	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))

	// shorthand
	// numbers := []int{}

	//duplicate array 
	var numbers2 = make([]int, len(numbers))
	copy(numbers2, numbers)
	fmt.Println(numbers, numbers2)

	// slice operator
	fmt.Println(numbers2[0:2])
	fmt.Println(numbers2[:2]) // shorthand if we want slice start from 0
	fmt.Println(numbers2[2:]) // after 2 index return all

	// 2D slices
	var twoDSlices = [][]int{{1,2,3}, {4,5,6}}
	fmt.Println(twoDSlices)
}