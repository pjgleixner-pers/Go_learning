package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := "my name"
	m1 = "asfg"
	fmt.Println(m1)
	//output: asfg

	m2 := "my name"
	m3 := "name"
	fmt.Println(strings.Contains(m2, m3))
	//output: true
	fmt.Println(strings.ReplaceAll(m2, "m", "NO"))
	//output: NOy naNOe
	fmt.Println(strings.Split(m2, "m"), m2+m3)
	//output: [ y na e] my namename

}
