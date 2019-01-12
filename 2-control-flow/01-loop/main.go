package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func main() {
	i := 1
	for i < 3 {
		fmt.Printf("i=%d \n", i)
		i = i + 1
	}
	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i=%d,j=%d\n", i, j)
		}
	}

	users := []user{
		{"Mike", "mike@example.com"},
		{"Alice", "alice@example.com"},
	}

	for _, user := range users {
		fmt.Printf("name = %s ; email = %s \n", user.name, user.email)
	}
	// name = Mike ; email = mike@example.com
	// name = Alice ; email = alice@example.com

	persons := map[int]user{
		1: user{name: "mike", email: "mike@example.com"},
		2: user{name: "alice", email: "alice@example.com"},
	}
	for k, v := range persons {
		fmt.Printf("No%d :name = %s ; email = %s \n", k, v.name, v.email)
	}
	// No1 :name = mike ; email = mike@example.com
	// No2 :name = alice ; email = alice@example.com
}
