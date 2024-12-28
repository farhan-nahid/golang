package main

import "fmt"

func main(){
	// shorthand
	// map1 := map[string]int {"phone1":12, "phone2":1}

	m:= make(map[string]string)


	m["name"] = "golang"
	m["area"] = "backend"

	fmt.Println(m["name"])
	fmt.Println(m["area"])
	fmt.Println(len(m))
	delete(m, "area")
	fmt.Println(m)
	clear(m)
	v, ok := m["phone"]

	if ok {
		fmt.Println("All Ok", v)
	} else {
		fmt.Println("Not Ok", v)
	}
}