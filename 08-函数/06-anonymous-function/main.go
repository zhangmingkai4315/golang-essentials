package main

import "fmt"

func defered() {
	defer func() {
		fmt.Println("this is defered function")
	}()
	fmt.Println("normal function call")
}

func main() {
	func() {
		fmt.Println("this is a anonymous function")
	}()
	defered()
}

// this is a anonymous function
// normal function call
// this is defered function
