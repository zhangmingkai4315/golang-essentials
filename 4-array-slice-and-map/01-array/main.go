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
}
