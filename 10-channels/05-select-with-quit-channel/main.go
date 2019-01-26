package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(quit chan<- bool) <-chan int {
	c := make(chan int)
	go func() {
		for {
			v := rand.Int() % 10
			if v == 9 {
				close(quit)
			} else {
				c <- v
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	return c
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	quit := make(chan bool)
	p := producer(quit)
	for {
		select {
		case v := <-p:
			fmt.Printf("%d\t", v)
		case q, ok := <-quit:
			fmt.Printf("receive data=%v and status=%v", q, ok)
			return
		}
	}
}
