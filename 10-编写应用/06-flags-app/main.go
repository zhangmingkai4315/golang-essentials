package main

import (
	"flag"
	"fmt"
)

var (
	word    string
	num     int
	boolVal bool
)

func init() {
	flag.StringVar(&word, "word", "foo", "an example of string")
	flag.IntVar(&num, "number", 42, "an example of int")
	flag.BoolVar(&boolVal, "bool", false, "an example of bool")
}

func main() {
	flag.Parse()
	fmt.Printf("word = %s\n", word)
	fmt.Printf("number = %d\n", num)
	fmt.Printf("bool = %t\n", boolVal)
}
