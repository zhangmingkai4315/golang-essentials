package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	print("begin")
	rand.Seed(time.Now().UTC().UnixNano())
}
func Timeit(f func()) func() {
	return func() {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed: %v ---n", time.Since(t))
		}(time.Now())
		f()
	}
}

func main() {
	f := func() {
		sleep := rand.Int31n(10)
		time.Sleep(time.Second * time.Duration(sleep))
	}
	profileFunc := Timeit(f)
	profileFunc()
}
