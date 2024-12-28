package main

import "fmt"

func changeNum (num *int){
	*num  =100;

	fmt.Println("changeNum func", *num)
}

func main(){
	num :=10;

	fmt.Println("Before change", num)

	changeNum(&num)

	fmt.Println("Main func", num)
}