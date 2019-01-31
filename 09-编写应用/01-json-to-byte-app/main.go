package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	Name        string
	Age         int
	privateInfo string
}

func main() {
	s1 := student{
		Name:        "mike",
		Age:         26,
		privateInfo: "i like music",
	}
	s2 := student{
		Name:        "alice",
		Age:         30,
		privateInfo: "i hate mike",
	}

	school := []student{s1, s2}

	schoolInfo, err := json.Marshal(school)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("school infomation is %v", string(schoolInfo))
	}

	// school infomation is [{"Name":"mike","Age":26},{"Name":"alice","Age":30}]
}
