package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("hello world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("print counter:%d", n)

	f, err := os.Create("temp.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("hello world")
	_, err = io.Copy(f, r)
	if err != nil {
		fmt.Println(err)
	}
}
