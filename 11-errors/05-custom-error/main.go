package main

import (
	"fmt"
)

type MyError struct {
	funcName    string
	funcLine    int
	description string
}

func (myerror *MyError) Error() string {
	return fmt.Sprintf("Error: funcname=%s line=%d", myerror.funcName, myerror.funcLine)
}

func myFunc() error {
	return &MyError{funcName: "myFunc", funcLine: 1}
}

func main() {
	err := myFunc()
	if err != nil {
		fmt.Println(err)
		// Error: funcname=myFunc line=1
	}

}
