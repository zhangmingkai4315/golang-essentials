package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}

}

func main() {
	fmt.Printf("OS %v\n", runtime.GOOS)
	fmt.Printf("ARCH %v\n", runtime.GOARCH)
	fmt.Printf("CPU %v\n", runtime.NumCPU())
	fmt.Printf("GOROUTINES %v\n", runtime.NumGoroutine())

	go boring("boring...")
	fmt.Printf("GOROUTINES %v\n", runtime.NumGoroutine())
	fmt.Println("I am listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}

// OS windows
// ARCH amd64
// CPU 12
// GOROUTINES 1

// GOROUTINES 2
// I am listening
// boring... 0
// boring... 1
// boring... 2
// boring... 3
// boring... 4
// boring... 5
// You're boring; I'm leaving.
