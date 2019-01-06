package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Printf("len(x)=%d\n", len(x))
	// len(x)=5
	for index, value := range x {
		fmt.Printf("index=%d, value=%d\n", index, value)
	}

	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:len(x)-2])
	// [2 3 4 5]
	// [2 3]
	// [1 2 3]

	y := []int{100, 200, 300}

	x = append(x, y...)
	fmt.Println(x) //[1 2 3 4 5 100 200 300]

	x = append(x[:2], x[4:]...)
	fmt.Println(x) //[1 2 5 100 200 300]

	z := make([]int, 10, 11)
	fmt.Printf("len of z is %d ,cap of z is %d\n", len(z), cap(z))
	// len of z is 10 ,cap of z is 11
	z = append(z, x...)
	fmt.Printf("len of z is %d ,cap of z is %d\n", len(z), cap(z))
	// len of z is 16 ,cap of z is 22

	persons := []string{"mike", "alice", "bob"}
	city := []string{"beijing", "sanjun", "tokyo"}
	info := [][]string{persons, city}

	fmt.Println(info)
	// [[mike alice bob] [beijing sanjun tokyo]]

}
