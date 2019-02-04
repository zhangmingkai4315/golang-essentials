package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Canceled by someone, i am leaving")
				return
			default:
				fmt.Println("working....")
				time.Sleep(time.Second * 1)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}
