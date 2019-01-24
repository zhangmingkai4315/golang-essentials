package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type sharedWithMutex struct {
	shared int
	mu     sync.Mutex
}

func main() {
	wg.Add(2)
	s := sharedWithMutex{
		shared: 0,
	}
	go func() {
		defer func() {

			wg.Done()
		}()
		for i := 0; i < 10000; i++ {
			s.mu.Lock()
			s.shared++
			s.mu.Unlock()
		}
	}()
	go func() {
		defer func() {

			wg.Done()
		}()
		for i := 0; i < 10000; i++ {
			s.mu.Lock()
			s.shared++
			s.mu.Unlock()
		}
	}()
	wg.Wait()
	fmt.Printf("the value of shared = %d , expected=20000", s.shared)
	// the value of shared = 20000 , expected=20000
}
