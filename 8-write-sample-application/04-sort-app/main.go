package main

import (
	"fmt"
	"sort"
)

type student struct {
	name  string
	age   int
	score int
}

type ByScore []student

func (b ByScore) Len() int { return len(b) }
func (b ByScore) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByScore) Less(i, j int) bool {
	return b[i].score < b[j].score
}

func main() {
	arr := []int{10, 2, 3, 42, 4, 3, 2, 4, 21}
	sort.Ints(arr)
	fmt.Println(arr)
	// [2 2 3 3 4 4 10 21 42]

	arrString := []string{"mike", "alice", "bob"}
	sort.Strings(arrString)
	fmt.Println(arrString)
	// [alice bob mike]

	school := []student{
		{"Mike", 12, 96},
		{"Alice", 13, 100},
		{"Bob", 12, 60},
	}
	fmt.Printf("Unsorted student array is %v\n", school)
	// Unsorted student array is [{Mike 12 96} {Alice 13 100} {Bob 12 60}]
	sort.Sort(ByScore(school))

	fmt.Printf("Sorted student array is %v\n", school)
	// Sorted student array is [{Bob 12 60} {Mike 12 96} {Alice 13 100}]

	school2 := []student{
		{"Mike", 12, 96},
		{"Alice", 13, 100},
		{"Bob", 12, 60},
	}

	sort.Slice(school2, func(i, j int) bool {
		return school2[i].score > school2[j].score
	})
	fmt.Printf("Sorted student2 array by score is %v\n", school)
	// Sorted student2 array by score is [{Bob 12 60} {Mike 12 96} {Alice 13 100}]
}
