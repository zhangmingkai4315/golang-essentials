package main

import (
	"fmt"
)

func ShowMessage(message interface{}) {
	switch _message := message.(type) {
	case string:
		fmt.Printf("string message: %s\n", _message)
	case int:
		fmt.Printf("int message: %d\n", _message)
	default:
		fmt.Println("Unknown type")
	}
}

func ShowStringMessage(message interface{}) {
	str, ok := message.(string)
	if ok == true {
		fmt.Printf("this is a string message :%s \n", str)
	} else {
		fmt.Println("value is not a string")
	}
}

func main() {
	ShowMessage("hello")
	ShowMessage(12)
	ShowMessage(1.00)

	ShowStringMessage("this is mike")
}

// string message: hello
// int message: 12
// Unknown type
// this is a string message :this is mike
