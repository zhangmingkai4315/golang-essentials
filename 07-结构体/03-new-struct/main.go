package main

import "fmt"

func main() {
	type person struct {
		name   string
		age    int
		isMale bool
		phones []string
	}

	p := new(person)
	fmt.Printf("the type of p = %T\n", p)
	// the type of p = *main.person
	fmt.Printf("the value of p = %+v", p)
	// the value of p = &{name: age:0 isMale:false phones:[]}
	p.phones = append(p.phones, "12345678")
	fmt.Printf("the value of p = %+v", p)
	// the value of p = &{name: age:0 isMale:false phones:[12345678]}
}

