package main

import "fmt"

// Append will append the newSlice to the old slice and
// return the whole slice
func Append(oldSlice []int, newSlice []int) []int {
	if len(newSlice) == 0 {
		return oldSlice
	}
	length := len(oldSlice)
	if length+len(newSlice) > cap(oldSlice) {
		temSlice := make([]int, (length + cap(newSlice)*2))
		copy(temSlice, oldSlice)
		oldSlice = temSlice
	}
	oldSlice = oldSlice[0 : length+len(newSlice)]
	copy(oldSlice[length:], newSlice)
	return oldSlice
}

func main() {
	oldSlice := []int{1, 2, 3, 4}
	fmt.Printf("len(oldslice)=%d, cap(oldslice)=%d\n", len(oldSlice), cap(oldSlice))
	//len(oldslice)=4, cap(oldslice)=4

	newSlice := []int{5, 6, 7, 8, 9, 10}
	oldSlice = Append(oldSlice, newSlice)
	fmt.Println(oldSlice)
	//[1 2 3 4 5 6 7 8 9 10]

	fmt.Printf("len(oldslice)=%d, cap(oldslice)=%d\n", len(oldSlice), cap(oldSlice))
	// len(oldslice)=10, cap(oldslice)=16
}
