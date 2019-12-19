package main

import "fmt"

type duration int

func (d *duration) pretty() string {
	return fmt.Sprintf("duration: %d", d)
}

func main() {
	// println((duration(10)).pretty())
	//cannot take the address of duration(10)
}
