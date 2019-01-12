package main

import (
	"errors"
	"fmt"
)

// Foo is a zero arguments function
func Foo() {
	fmt.Println("Foo: function")
}

// FooWithArgs function in go pass by value
func FooWithArgs(s string) {
	fmt.Printf("FooWithArgs: %s\n", s)
}

// FooWithMultiArgs function in go pass by value
func FooWithMultiArgs(s string, prefix bool) {
	if prefix == true {
		fmt.Printf("FooWithArgs: %s\n", s)
	} else {
		fmt.Printf("%s\n", s)
	}

}

// FooWithArgsAndReturn with string return
func FooWithArgsAndReturn(s string) string {
	return fmt.Sprintf("FooWithArgsAndReturn: %s\n", s)
}

// FooWithArgsAndMultiReturn will return multi return
func FooWithArgsAndMultiReturn(s string) (string, error) {
	if s == "error" {
		return "", errors.New("FooWithArgsAndMultiReturn: Error")
	}
	return fmt.Sprintf("FooWithArgsAndMultiReturn: %s\n", s), nil
}

// FooWithArgsAndMultiDefaultReturn will return multi return with some default
func FooWithArgsAndMultiDefaultReturn(s string) (message string, err error) {
	if s == "error" {
		err = errors.New("FooWithArgsAndMultiDefaultReturn: Error")
		return
	}
	message = fmt.Sprintf("FooWithArgsAndMultiDefaultReturn: %s\n", s)
	return
}

func main() {
	Foo()
	// Foo: function

	FooWithArgs("Hello")
	// FooWithArgs: Hello
	FooWithMultiArgs("Hello", false)
	// Hello

	fmt.Print(FooWithArgsAndReturn("Hello world"))
	// FooWithArgsAndReturn: Hello world

	_, err := FooWithArgsAndMultiReturn("error")
	if err != nil {
		fmt.Println(err)
	}
	// FooWithArgsAndMultiReturn: Error

	_, err = FooWithArgsAndMultiDefaultReturn("error")
	if err != nil {
		fmt.Println(err)
	}
	// FooWithArgsAndMultiReturn: Error

	f := func(name string) string {
		return fmt.Sprintf("My name is %s", name)
	}
	f("mike")
}
