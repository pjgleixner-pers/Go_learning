package main

import "fmt"

//structure encapsulation
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println(c.Name, "is Driving")
}

func main() {
	c1 := Car{"chevy", 1, 2}
	c2 := Car{
		Name:    "toyota",
		Age:     3,
		ModelNo: 4,
	}
	c1.Print()
	//output: {chevy 1 2}
	c1.Drive()
	//output: chevy is Driving
	c2.Print()
	//output: {toyota 3 4}
	c2.Drive()
	//outputtoyota is Driving

}
