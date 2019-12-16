package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	workersNumber = 5
	tasksNumber   = 100
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, tasksNumber)
	wg.Add(workersNumber)

	for i := 0; i < workersNumber; i++ {
		go func(workerID int) {
			defer wg.Done()
			for {
				task, ok := <-tasks
				if ok == false {
					fmt.Println("task queue is empty , worker ", workerID, " will quit")
					return
				}
				time.Sleep(time.Duration(rand.Int63n(100)) * time.Millisecond)
				fmt.Printf("worker done the task : %s\n", task)
			}
		}(i)
	}

	for i := 0; i < tasksNumber; i++ {
		tasks <- fmt.Sprintf("Tasks: %d", i)
	}
	close(tasks)

	wg.Wait()
}
