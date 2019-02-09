package main

import (
	"fmt"
	"time"
)

func main() {
	switch {
	case (2 == 2):
		fmt.Println("2==2 is true")
		fallthrough
	case (3 == 3):
		fmt.Println("3==3 is true too")
	default:
		fmt.Println("default case")
	}
	// 2==2 is true
	// 3==3 is true too

	score := 7
	switch score {
	case 0, 1, 3:
		fmt.Println("Terrible")
	case 4, 5:
		fmt.Println("Mediocre")
	case 6, 7:
		fmt.Println("Not bad")
	case 8, 9:
		fmt.Println("Almost perfect")
	case 10:
		fmt.Println("hmm did you cheat?")
	default:
		fmt.Println(score, " off the chart")
	}
	// Not bad

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's a weekday")
	}

	typePrintFunc := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("args is a bool type")
		case string:
			fmt.Println("args is a string")
		default:
			fmt.Println("args type unknown")
		}
	}
	typePrintFunc("mike")
	typePrintFunc(12)
	// args is a string
	// args type unknown
}
