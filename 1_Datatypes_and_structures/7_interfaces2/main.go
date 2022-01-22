package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	Anything(2.44)
	Anything("asdasfg")
	Anything(struct{}{})

	mymap := make(map[string]interface{})

	mymap["name"] = "L9784MMF09W3"
	mymap["age"] = 10
	fmt.Println(mymap)

	// output:
	// 		2.44
	// 		asdasfg
	// 		{}
	// 		map[age:10 name:L9784MMF09W3]

}
