package main

import "fmt"

type city struct {
	name      string
	provience string
	country   string
}

func (c city) ShowCountry() string {
	return c.country
}

func (c *city) Info() string {
	return fmt.Sprintf("%s is in %s", c.name, c.ShowCountry())
}

func (c city) ChangeCountryName(country string) {
	// fmt.Printf("the address of c is %v\n", &c)
	c.country = country
}

func (c *city) ChangeCountryNameByPointer(country string) {
	// fmt.Printf("the address of c is %v\n", &c)
	c.country = country
}

type information interface {
	Info() string
}

func show(i information) {
	fmt.Println(i.Info())
}

func main() {
	c := city{
		name:      "beijing",
		provience: "beijing",
		country:   "china",
	}

	fmt.Println(c.Info())
	fmt.Println((&c).Info())
	// beijing is in china
	// beijing is in china

	(&c).ChangeCountryName("China")
	fmt.Println(c.Info())
	fmt.Println((&c).Info())
	// beijing is in china
	// beijing is in china

	c.ChangeCountryNameByPointer("China")
	fmt.Println(c.Info())
	fmt.Println((&c).Info())

	// beijing is in China
	// beijing is in China

	// show(c)
	//  cannot use c (type city) as type information in argument to show:
	// city does not implement information (Info method has pointer receiver)

	show(&c)
	// beijing is in China
}
