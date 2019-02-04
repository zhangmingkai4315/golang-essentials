package main

import "fmt"

// everything in go pass by value!！！

func addValue(i int) {
	i = i + 1
}

func addValueByPointer(i *int) {
	*i = *i + 1
}

func main() {
	a := 10
	addValue(a)
	fmt.Printf("a = %d\n", a)
	// a = 10

	addValueByPointer(&a)
	fmt.Printf("a = %d\n", a)
	// a = 11
}
