package main

import (
	"fmt"
)

func main() {
	//if
	f := true
	flag := &f
	if flag == nil {
		fmt.Println("nil")
	} else if *flag {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	//for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	arr := []string{"first", "second", "third"}
	for i, value := range arr {
		fmt.Println(i, ": ", value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "ALJ33PO235"
	mymap["age"] = "20"

	for k, v := range mymap {
		fmt.Printf("Key: %s - Value: %v \n", k, v)
	}
}
