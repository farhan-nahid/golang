package main

import "fmt"


func sum (numbers ...int)int{
	total :=0;

	for _, num :=range numbers{
		total += num
	}

	return total
}

func anyFunc (params ...interface{}){}

func main(){
	fmt.Println(sum(1,2,3,4,5))
	
	numbers := []int{1,2,3,4,5}
	fmt.Println(sum(numbers...))
}