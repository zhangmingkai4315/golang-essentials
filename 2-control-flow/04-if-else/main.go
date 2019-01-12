package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("Always print")
	}
	if false {
		fmt.Println("Never print")
	}
	if num := 9; num < 0 {
		fmt.Println("Number is negative")
	} else if num > 0 {
		fmt.Println("Number is positive")
	} else {
		fmt.Println("Number is zero")
	}
	// Always print
	// Number is positive
}
