package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// go get golang.org/x/crypto/bcrypt
func main() {
	s := "password"
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bs))
	// $2a$04$48e/0SGKpjsJ/nD6G4S83ey65T1.kCXdplusBEX/iIkBzlegGi.b6
	// 每次即便是输入相同，输出也不同

	confirmPassword := "password"

	err = bcrypt.CompareHashAndPassword(bs, []byte(confirmPassword))
	if err != nil {
		panic(err)
	}
	fmt.Println("password is correct")

}
