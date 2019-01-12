package main

import "fmt"

type dog struct {
	name  string
	age   uint8
	color string
	dogDetail
}

func (d dog) Info() string {
	return fmt.Sprintf("Dog name is %s, age %d and color is %s", d.name, d.age, d.color)
}

type dogDetail struct {
	weight    float32
	height    float32
	favorfood string
}

func (dd dogDetail) ShowFoverFood() string {
	return dd.favorfood
}
func main() {
	d := dog{
		name:  "litterOne",
		age:   2,
		color: "black",
		dogDetail: dogDetail{
			weight:    25.1,
			height:    0.9,
			favorfood: "tomato",
		},
	}
	fmt.Printf("%s, favor food is %s", d.Info(), d.ShowFoverFood())
	// Dog name is litterOne, age 2 and color is black, favor food is tomato
}
