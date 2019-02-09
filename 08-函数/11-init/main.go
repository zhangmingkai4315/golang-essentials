package main

import "fmt"

var (
	a = 1
	b = 2
)

func init() {
	fmt.Println("Start init1 function")
	fmt.Println(a, b)
	a++
}
func init() {
	fmt.Println("Start init2 function")
	fmt.Println(a, b)
	a++
}
func main() {
	fmt.Println("Start main function ")
	fmt.Println(a, b)
}

// Start init1 function
// 1 2
// Start init2 function
// 2 2
// Start main function
// 3 2
