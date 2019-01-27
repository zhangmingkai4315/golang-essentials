package main

import (
	"errors"
	"fmt"
	"log"
)

func badFunc() (string, error) {
	return "", errors.New("error occure")
}

func main() {
	defer fmt.Println("panic will do defer before exit")
	_, err := badFunc()
	if err != nil {
		fmt.Println(err)
	}

	_, err = badFunc()
	if err != nil {
		log.Println(err)
	}
	_, err = badFunc()
	if err != nil {
		log.Panicln(err)
		// print code location and exit with code 2 without defer
	}
}
