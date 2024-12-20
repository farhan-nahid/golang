package main

import "fmt"

func main(){
	// while loop syntax
	const COUNT = 3

	i:=0

	fmt.Println("while loop syntax")
	for  i <= COUNT {
		fmt.Print(i)
		i++
	}

	// classic for loop

	fmt.Println("\nclassic for loop")
	for i:=0; i<=COUNT; i+=1 {

		if i == 2 {
			continue
		}

		fmt.Println(i)
	}


	// range loop

	fmt.Println("range loop")
	for i := range 5 {
		fmt.Println(i)
	}
}