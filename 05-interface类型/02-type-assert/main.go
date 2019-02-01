package main

import (
	"fmt"
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

func main() {
	p := Programmer{
		name: "mike",
		age:  24,
	}
	Say(p)
	d := Doctor("alice")
	Say(d)

}
