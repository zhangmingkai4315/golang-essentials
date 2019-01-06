package main

import (
	"fmt"
)

const (
	_  = iota
	kb = 1 << (10 * iota)
	mb
	gb
	tb
)

func main() {
	x := 10
	fmt.Printf("%d\t\t%b\n", x, x)
	// 10              1010
	fmt.Printf("%d\t\t%b\n", x<<1, x<<1)
	// 20              10100
	fmt.Printf("%d\n", kb)
	fmt.Printf("%d\n", mb)
	fmt.Printf("%d\n", gb)
	fmt.Printf("%d\n", tb)
	// 1024
	// 1048576
	// 1073741824
	// 1099511627776
}
