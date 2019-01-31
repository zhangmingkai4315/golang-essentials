package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer() <-chan int {
	c := make(chan int)
	go func() {
		for {
			c <- rand.Int() % 1000
			time.Sleep(time.Millisecond * 100)
		}
	}()
	return c
}

func main() {
	p := producer()
	t := time.Tick(time.Second * 5)
	for {
		select {
		case v := <-p:
			fmt.Printf("%d\t", v)
		case <-t:
			fmt.Println("time out")
			return
		}
	}
}
