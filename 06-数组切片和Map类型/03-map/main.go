package main

import "fmt"

func main() {
	person := map[string]uint16{
		"mike":  25,
		"alice": 20,
		"bob":   24,
	}
	fmt.Printf("mike age is %d\n", person["mike"])
	// mike age is 25

	// fmt.Printf("tom age is %d\n", person["tom"])
	if v, isExist := person["tom"]; isExist {
		fmt.Printf("tom age is %d\n", v)
	} else {
		fmt.Println("tom is not in list")
	}
	// tom is not in list

	for k, v := range person {
		fmt.Printf("%s age is %d\n", k, v)
	}
	// mike age is 25
	// alice age is 20
	// bob age is 24

	delete(person, "mike")
	fmt.Printf("%+v", person) //  map[alice:20 bob:24]
	delete(person, "nickey")
	fmt.Println(len(person)) // 2

}
