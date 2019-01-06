package main

import "fmt"

func main() {
	m1 := struct {
		Name, Address string
	}{
		Name:    "Mike",
		Address: "Beijing",
	}
	fmt.Printf("%+v", m1)
}
