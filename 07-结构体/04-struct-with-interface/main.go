package main

import "fmt"

// User ...
type User struct {
	Name  string
	Email string
}

// Notify will send a email to the caller
func (u *User) Notify() error {
	fmt.Printf("Send notify to %s(%s)\n", u.Name, u.Email)
	return nil
}

// Notifier include a simple Notify function
type Notifier interface {
	Notify() error
}

func SendNotification(n Notifier) error {
	return n.Notify()
}

func main() {
	mike := User{"Mike", "mike@example.com"}
	// mike.Notify()
	SendNotification(&mike)
	tom := &User{"tom", "tom@example.com"}
	// tom.Notify()
	SendNotification(tom)
}
