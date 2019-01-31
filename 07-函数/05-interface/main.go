package main

import (
	"fmt"
	"sort"
)

type Programmer struct {
	name string
	age  int
}

func (p Programmer) speak() string {
	return fmt.Sprintf("i am a programmer")
}

func (p Programmer) Doing() string {
	return "Coding..."
}

type Doctor string

func (b Doctor) speak() string {
	return "i am a doctor"
}

type Human interface {
	speak() string
}

func Say(h Human) {
	switch h.(type) {
	case Programmer:
		fmt.Printf("Programmer say: %s and right now i am doing %s \n", h.speak(), h.(Programmer).Doing())
	case Doctor:
		fmt.Printf("Doctor say: %s\n", h.speak())
	default:
		fmt.Printf("Some one say: %s\n", h.speak())
	}
}

type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
	sort.Sort(s)
	str := "["
	for i, elem := range s {
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

func main() {
	p := Programmer{
		name: "mike",
		age:  24,
	}
	Say(p)
	d := Doctor("alice")
	Say(d)

	s := Sequence([]int{10, 32, 4, 2, 4, 2, 24, 221, 3, 23})
	fmt.Println(s.String())
	// [2 2 3 4 4 10 23 24 32 221]
}
