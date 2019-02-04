// Package mylib offer some basic lib for welcome
package mylib

import (
	"errors"
	"fmt"
)

// Error codes returned by failures to parse an expression.
var (
	ErrInternal      = errors.New("regexp: internal error")
	ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
	ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
)

// ShowWelcome receive a username value and
// print welcome message
func ShowWelcome(name string) {
	fmt.Printf("Hi %s, welcome", name)
}
