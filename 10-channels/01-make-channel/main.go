package main

import "fmt"

// func main() {
// 	c := make(chan int)
// 	c <- 1
// 	fmt.Println("never show this message")
// 	// all goroutines are asleep - deadlock!
// }

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
	fmt.Println("exit success")
}
