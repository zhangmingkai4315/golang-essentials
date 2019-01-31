package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var shared = 0
	wg.Add(2)
	go func() {

		for i := 0; i < 10000; i++ {
			shared++
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000; i++ {
			shared++
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("the value of shared = %d , expected=20000", shared)
}

// the value of shared = 11649 , expected=20000
