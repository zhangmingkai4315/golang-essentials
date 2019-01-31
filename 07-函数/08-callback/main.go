package main

import "fmt"

func sumInt(arr ...int) int {
	var total = 0
	for _, i := range arr {
		total += i
	}
	return total
}

func even(sum func(...int) int, arr ...int) int {
	newArray := []int{}
	for _, i := range arr {
		if i%2 == 0 {
			newArray = append(newArray, i)
		}
	}
	return sumInt(newArray...)
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sumInt(a...))
	// 55
	fmt.Println(even(sumInt, a...))
	// 30
}
