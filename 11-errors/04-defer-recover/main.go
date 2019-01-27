package main

import "fmt"

func c() (i int) {
	defer func() {
		i++
	}()
	return 1
}

func recoveredFunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("function panic is somewhere：%v", r)
			// function panic is somewhere：kernal panic
		}
	}()

	panic("kernal panic")
	fmt.Println("function will not print this line")
}

func main() {
	x := c()
	fmt.Println(x)
	// 2

	recoveredFunc()

}
