package main

import "fmt"

// SumWithVariadicArgs with variadic parameter
func SumWithVariadicArgs(a ...int) int {
	sum := 0
	for _, i := range a {
		sum = sum + i
	}
	return sum
}

func main() {
	fmt.Printf("sum(10,20)=%d\n", SumWithVariadicArgs(10, 20))
	fmt.Printf("sum(10,20,30)=%d\n", SumWithVariadicArgs(10, 20, 30))
	arr := []int{10, 20, 30, 23, 23}
	fmt.Printf("sum(10,20,30,23,23) = %d\n", SumWithVariadicArgs(arr...))
	// sum(10,20)=30
	// sum(10,20,30)=60
	// sum(10,20,30,23,23) = 106
}
