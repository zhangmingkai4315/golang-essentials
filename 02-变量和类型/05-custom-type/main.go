package main

import (
	"fmt"
)

type MyCustomType string

var mct MyCustomType

func main() {
	mct = "hello"

	fmt.Println(mct)

	temp := "world"

	// mct = temp
	// cannot use temp (type string) as type MyCustomType in assignment

	mct = MyCustomType(temp)
	
	fmt.Println(mct)
}
