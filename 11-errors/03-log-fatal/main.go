package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func badFunc() (string, error) {
	return "", errors.New("error occure")
}

func main() {
	f, err := os.Create("temp.log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		f.Close()
		fmt.Println("close file success?")
	}()

	log.SetOutput(f)
	defer fmt.Println("not working")
	_, err = badFunc()
	if err != nil {
		log.Fatal(err)
		// print and exit with code 1 without defer  (Fatal = println + os.Exit())
		// Fatal always call os.Exit , so all defer function will not work
	}
}
