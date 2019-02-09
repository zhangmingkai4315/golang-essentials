package main

import (
	"fmt"
)

func main() {
	fmt.Printf("true && true = %v \n", true && true)
	fmt.Printf("true && false = %v\n", true && false)
	fmt.Printf("true || true = %v\n", true || true)
	fmt.Printf("true || false = %v\n", true || false)
	fmt.Printf("!true = %v\n", !true)
}

// true && true = true
// true && false = false
// true || true = true
// true || false = true
// !true = false
