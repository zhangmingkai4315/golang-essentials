package main

import "fmt"

// User ...
type User struct {
	Name  string
	Email string
}

type Admin struct {
	User
	Level string
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

func (admin *Admin) Notify() error {
	fmt.Printf("Admin send notify to %s(%s)\n", admin.Name, admin.Email)
	return nil
}

func main() {
	admin := &Admin{
		User: User{
			"Mike", "mike@example.com",
		},
		Level: "superadmin",
	}
	SendNotification(admin)
	admin.Notify()
	admin.User.Notify()
}

// Admin send notify to Mike(mike@example.com)
// Admin send notify to Mike(mike@example.com)
// Send notify to Mike(mike@example.com)
