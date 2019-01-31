package main

import (
	"fmt"
)

type Person struct {
	name string
}
type MyInterface interface {
	Read(message string) string
}

func (t Person) Read(message string) string {
	return fmt.Sprintln(t.name + " is reading " + message)
}

func main() {

	t1 := 64
	fmt.Printf("%d is %T\n", t1, t1)

	t2 := "mike"
	fmt.Printf("%s is %T\n", t2, t2)

	t3 := false
	fmt.Printf("%t is %T\n", t3, t3)

	t4 := &t1
	fmt.Printf("%v is %T\n", t4, t4)

	t5 := []string{"i", "really", "like", "golang"}
	fmt.Printf("%v is %T\n", t5, t5)

	t6 := [4]string{"i", "really", "like", "golang"}
	fmt.Printf("%v is %T\n", t6, t6)

	t7 := map[int]string{
		1: "mike",
		2: "alice",
		3: "anjoue",
	}
	fmt.Printf("%v is %T\n", t7, t7)

	t8 := struct {
		name string
		age  int
	}{"mike", 12}
	fmt.Printf("%v is %T\n", t8, t8)

	t9 := func(message string) {
		fmt.Println(message)
	}
	fmt.Printf("t9 is %T\n", t9)

	type Mytype int8
	var mytype Mytype
	fmt.Printf("mytype is %T\n", mytype)

	p := Person{"Mike"}
	fmt.Println(p.Read("novel"))
}
