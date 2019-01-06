package main

import "fmt"

func deferedFunc() {
	fmt.Println("this function is defered!")
}

func funcWithDefer() {
	defer deferedFunc()
	fmt.Println("this is function with defer")
}

func main() {
	funcWithDefer()
	// this is function with defer
	// this function is defered!
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Println("done")
	// counting
	// done
	// 9 8 7 6 5 4 3 2 1 0
}
