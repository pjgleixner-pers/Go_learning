package main

import "fmt"

func main() {
	//as a group, many variables in one go
	var (
		m1 = 2
		m2 = 3
	)
	//directly
	m3 := 5
	//variable then data
	var m4 int
	m4 = 6

	fmt.Println(m1, m2, m3, m4)
	//output: 2 3 5 6

}
