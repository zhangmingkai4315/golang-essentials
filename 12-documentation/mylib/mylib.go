// Package mylib offer some basic lib for welcome
package mylib

import "fmt"

// ShowWelcome receive a username value and
// print welcome message
func ShowWelcome(name string) {
	fmt.Printf("Hi %s, welcome", name)
}
