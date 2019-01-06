package main

import (
	"fmt"
	"strconv"
)

const (
	a1 = 1
	a2 = 2
	a3 = 3
)

const (
	b1 = iota
	b2
	b3
)

type Season uint8

const (
	Spring = Season(iota)
	Summer
	Autumn
	Winner
)

func (s Season) String() string {
	name := []string{"spring", "summer", "autumn", "winner"}
	i := uint8(s)
	switch {
	case i <= uint8(Winner):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}

func main() {
	const name = "mike"
	const age = 26

	fmt.Println(name, age)
	// name = "alice"
	//cannot assign to name

	fmt.Println(a1, a2, a3)
	// 1 2 3
	fmt.Println(b1, b2, b3)
	// 0 1 2

	fmt.Println(Summer)
	fmt.Println(Season(0))
	// summer
	// spring
}
