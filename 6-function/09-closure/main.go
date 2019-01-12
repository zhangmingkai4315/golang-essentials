package main

import "fmt"

func incfactor(base int) func() int {
	i := base
	return func() int {
		i = i + 1
		return i
	}
}


func fibFactoty() func() int {
	x, y := 0, 1
	return func() (r int) {
		r = x
		x, y = y, x+y
		return
	}
}

func main() {
	inc := incfactor(10)
	fmt.Println(inc())
	fmt.Println(inc())
	fmt.Println(inc())
	// 11
	// 12
	// 13

	fib := fibFactoty()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fib())
	}
	// 0 1 1 2 3 5 8 13 21 34

}
