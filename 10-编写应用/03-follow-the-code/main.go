package main

import (
	"fmt"
	"io"
	"os"
)

// func Println(a ...interface{}) (n int, err error) {
// 	return Fprintln(os.Stdout, a...)
// }

// $ go doc fmt.Fprintln
// func Fprintln(w io.Writer, a ...interface{}) (n int, err error)

// package io
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// var (
// 	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
// 	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
// 	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
// )

// $ go doc os.NewFile
// func NewFile(fd uintptr, name string) *File

// func (f *File) Write(b []byte) (n int, err error)

func main() {
	fmt.Println("hello world")
	fmt.Fprintf(os.Stdout, "hello world\n")
	io.WriteString(os.Stdout, "hello world")
}
