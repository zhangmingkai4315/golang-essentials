package main

import "fmt"

func main() {
	x := 10
	y := 10.2
	fmt.Printf("x = 10 is %T\n", x)
	fmt.Printf("y= 10.2 is %T\n", y)
	// x = 10 is int
	// y= 10.2 is float64

	var a1 uint8 // [0,255]
	a1 = 128
	fmt.Printf("a1=128 is %T\n", a1)

}
