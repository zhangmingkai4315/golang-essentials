package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var shared int64

	wg.Add(2)
	go func() {
		defer func() {

			wg.Done()
		}()
		for i := 0; i < 10000; i++ {
			atomic.AddInt64(&shared, 1)
		}
	}()
	go func() {
		defer func() {

			wg.Done()
		}()
		for i := 0; i < 10000; i++ {
			atomic.AddInt64(&shared, 1)
		}
	}()
	wg.Wait()
	fmt.Printf("the value of shared = %d , expected=20000", shared)
	// the value of shared = 20000 , expected=20000

}
