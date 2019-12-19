package main

import (
	"fmt"

	"github.com/zhangmingkai4315/golang-essentials/07-结构体/06-struct-export/lib"
)

func main() {
	admin := &lib.Admin{
		Level: "superadmin",
	}
	admin.Name = "Mike"
	admin.Email = "example@example.com"
	fmt.Printf("%v", admin)
	//{{Mike example@example.com} superadmin}
}

// Admin send notify to Mike(mike@example.com)
// Admin send notify to Mike(mike@example.com)
// Send notify to Mike(mike@example.com)
