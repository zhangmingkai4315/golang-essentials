// $ go doc fmt.Println
// func Println(a ...interface{}) (n int, err error)
//     Println formats using the default formats for its operands and writes to
//     standard output. Spaces are always added between operands and a newline is
//     appended. It returns the number of bytes written and any write error
//     encountered.

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello", 12, false)

	n, err := fmt.Println("hello world")
	fmt.Println(n, err)

	n1, _ := fmt.Println(12)
	fmt.Println(n1)

}

// hello 12 false
// hello world
// 12 <nil>
