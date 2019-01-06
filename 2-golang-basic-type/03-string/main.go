package main

import "fmt"

func main() {
	welcome := "hello world"
	welcomeByte := []byte(welcome)
	fmt.Println(welcomeByte)
	fmt.Printf("welcomeByte is %T and size is %d", welcomeByte, len(welcomeByte))
	// [104 101 108 108 111 32 119 111 114 108 100]
	// welcomeByte is []uint8

	welcomeCN := "hello 世界"
	welcomeByteCN := []byte(welcomeCN)
	fmt.Println(welcomeByteCN)
	fmt.Printf("welcomeByteCN is %T and size is %d", welcomeByteCN, len(welcomeByteCN))
	// welcomeByte is []uint8 and size is 11[104 101 108 108 111 32 228 184 150231 149 140]
	// welcomeByteCN is []uint8 and size is 12

	welcomeRuneCN := []rune(welcomeCN)
	fmt.Println(welcomeRuneCN)
	fmt.Printf("welcomeRuneCN is %T and size is %d", welcomeRuneCN, len(welcomeRuneCN))
	// [104 101 108 108 111 32 19990 30028]
	// welcomeRuneCN is []int32 and size is 8
}
