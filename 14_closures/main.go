package main

import "fmt"

func counter () func() int{
	count :=0

	return func() int {
		count+=1

		return count
	}
}

func main(){
	increment1 := counter()
	increment2 := counter()

	fmt.Println(increment1())
	fmt.Println(increment1())
	fmt.Println(increment1())

	fmt.Println(increment2())
	fmt.Println(increment2())
	fmt.Println(increment2())

}