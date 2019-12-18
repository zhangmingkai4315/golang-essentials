package main

import "fmt"

func showArray(arr [5]int) {
	fmt.Println(arr)
}

func main() {
	x := [10]int{}
	x[5] = 3
	fmt.Println(x)
	// [0 0 0 0 0 3 0 0 0 0]

	// showArray(x)
	//  cannotuse x (type [10]int) as type [5]int

	y := [...]int{10, 20, 30, 40}
	z := [5]int{1: 20, 2: 20}

	fmt.Println(y, z)
	// [10 20 30 40] [0 20 20 0 0]

	zCopy := [5]int{}

	zCopy = z
	zCopy[0] = 1
	fmt.Println(z, zCopy)
	// [0 20 20 0 0] [1 20 20 0 0]

}
