package main

import "fmt"

func main() {
	var i = 10
	fmt.Printf("the address of i=%d is %v\n", i, &i)
	// the address of i=10 is 0xc000062058

	var j = &i
	fmt.Printf("the address of j=%d is %v\n", *j, j)
	// the address of j=10 is 0xc000062058

	fmt.Printf("the type of i is %T, the type of j is %T\n", i, j)
	// the type of i is int, the type of j is *int

	fmt.Printf("the value of i is %d\n", *&i)
	// the value of i is 10

	*j = 20
	fmt.Printf("the value of i is %d\n", *&i)
	// the value of i is 20
}
