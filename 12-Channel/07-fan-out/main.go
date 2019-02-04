package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func producer(counter *int) <-chan int {
	c := make(chan int)
	go func() {
		t := time.Tick(time.Second * 10)
		for {
			c <- rand.Int() % 5
			*counter++
			select {
			case <-t:
				close(c)
				return
			default:
				time.Sleep(time.Millisecond)
			}
		}
	}()
	return c
}

func worker(id int, job <-chan int, counter *int32) {
	defer wg.Done()
	for x := range job {
		fmt.Printf("ID=%d, receive job %d\n", id, x)
		time.Sleep(time.Duration(x) * time.Second)
		fmt.Printf("ID=%d, Done job %d\n", id, x)
		atomic.AddInt32(counter, 1)
	}
}
func main() {
	var producerJobCounter int
	producer := producer(&producerJobCounter)
	var doneCounter int32
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, producer, &doneCounter)
	}
	wg.Wait()
	fmt.Printf("Total create job %d, done %d jobs\n", producerJobCounter, doneCounter)
	// Total create job 5398, done 5398 jobs
}
