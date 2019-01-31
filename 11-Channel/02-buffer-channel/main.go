package main

import "fmt"

func main() {
	c := make(chan int, 1)
	c <- 1
	fmt.Println("not block this message")
	c <- 2
	fmt.Println("block again, never show this message")
}
