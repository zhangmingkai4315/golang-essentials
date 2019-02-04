package main

import (
	"fmt"
	"time"
)

func producer() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
			time.Sleep(time.Second * 1)
		}
		close(c)
	}()
	return c
}

func consumer(c <-chan int) {
	for x := range c {
		fmt.Println(x)
	}
}

func main() {
	p := producer()
	consumer(p)

}
