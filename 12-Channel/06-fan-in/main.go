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
			time.Sleep(time.Millisecond * 1000)
		}

	}()
	return c
}

func merge(producers ...<-chan int) <-chan int {
	all := make(chan int)
	for _, p := range producers {
		go func(p <-chan int) {
			for {
				all <- <-p
			}
		}(p)
	}
	return all
}
func main() {
	pl := []<-chan int{}
	for i := 0; i < 10; i++ {
		p := producer()
		pl = append(pl, p)
	}
	all := merge(pl...)
	t := time.Tick(time.Second * 5)
	for {
		select {
		case v := <-all:
			fmt.Printf("%d\t", v)
		case <-t:
			fmt.Println("timeout")
			return
		}
	}
}
