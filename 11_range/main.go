package main

import "fmt"

func main(){
	var numbers = []int{1,2,3}

	for index, num := range numbers{
		fmt.Println("num",num, "Index", index)
	}

	m := map[string]string{"firstName": "Farhan", "lastName": "Nahid"}

	// for k, v := range m{
	for key, value := range m{
		fmt.Println(key, value)
	}

	// c -> unicode code point rune
	// i -> staring byte of rune
	// 255 -> 1 byte
	for i, c := range "golang"{
		fmt.Println(i, string(c), c)
	}
}