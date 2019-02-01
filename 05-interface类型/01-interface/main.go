package main

import "fmt"

type human interface {
	speak() string
}

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

func main() {
	var p human
	p = Programmer{
		name: "mike",
		age:  20,
	}
	p.speak()
}
