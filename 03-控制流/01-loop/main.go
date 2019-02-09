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
	// 同 while
	for i < 3 {
		fmt.Printf("i=%d \n", i)
		i = i + 1
	}
	// 进入无限循环，直到内部代码使用break主动退出
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
	// 如果仅仅使用第二个对象（对于数组或者slice的话每次获得一对[index, value]）
	for _, user := range users {
		fmt.Printf("name = %s ; email = %s \n", user.name, user.email)
	}
	// name = Mike ; email = mike@example.com
	// name = Alice ; email = alice@example.com

	//
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
