package main

import "fmt"

func main() {
	var x bool
	fmt.Println(x)
	x = true
	fmt.Println(x)
	// false
	// true

	x1 := 12
	x2 := 20
	fmt.Println(x1 == x2)
	fmt.Println(x1 <= x2)
	fmt.Println(x1 >= x2)

	// false
	// true
	// false
}
