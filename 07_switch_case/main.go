package main

import (
	"fmt"
	"time"
)

func main (){
	i:=5

	// single condition switch
	switch i {
		case 1:
			fmt.Println("One")

		case 2:
			fmt.Println("Two")

		case 3:
			fmt.Println("Three")

		case 4:
			fmt.Println("Four")

		case 5:
			fmt.Println("Five")

		default:
			fmt.Println("Unhandled")
		
	}





	// multiple condition switch
	switch time.Now().Weekday(){
		case time.Friday , time.Saturday:
			fmt.Println("It's Weekend")
		
		default:
			fmt.Println("It's Workday")
	}





	// type switch
	whoAmI := func  (i interface{})  {
		// switch t:= i.(type){
		switch i.(type){
			case int:
				fmt.Println("Integer value")
			
			case string:
				fmt.Println("String value")

			case bool:
				fmt.Println("Boolean value")

			default:
				// fmt.Println("Unhandled", t)
				fmt.Println("Unhandled")
		}		
	}

	whoAmI("golang")
}