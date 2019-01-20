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

type studentWithLabel struct {
	Name        string `json:"name"`
	Age         int    `json:"age"`
	privateInfo string `json:"private"` // useless
}

func main() {
	school := []student{}
	rawdata := `[{"Name":"mike","Age":26},{"Name":"alice","Age":30}]`

	err := json.Unmarshal([]byte(rawdata), &school)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("school infomation is %v\n", school)
	}
	// school infomation is [{mike 26 } {alice 30 }]

	rawdata2 := `[{"name":"mike","age":26,"private":"secret information"},{"name":"alice","age":30}]`

	studentLabelArray := []studentWithLabel{}
	err = json.Unmarshal([]byte(rawdata2), &studentLabelArray)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("school infomation is %v\n", studentLabelArray)
	}

	// school infomation is [{mike 26 } {alice 30 }]
}
