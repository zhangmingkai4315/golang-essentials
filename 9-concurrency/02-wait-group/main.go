package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("No %d goroutine begin\n", i)
			fmt.Printf("No %d goroutine is done\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

// No 1 goroutine begin
// No 3 goroutine begin
// No 0 goroutine begin
// No 9 goroutine begin
// No 4 goroutine begin
// No 7 goroutine begin
// No 8 goroutine begin
// No 2 goroutine begin
// No 5 goroutine begin
// No 6 goroutine begin
// No 9 goroutine is done
// No 1 goroutine is done
// No 4 goroutine is done
// No 6 goroutine is done
// No 7 goroutine is done
// No 8 goroutine is done
// No 5 goroutine is done
// No 2 goroutine is done
// No 0 goroutine is done
// No 3 goroutine is done
