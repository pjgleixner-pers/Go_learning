package main

import "fmt"

//array
func todo() {
	//var arr []int
	arr := []int{1, 2, 3, 4}
	arr2 := []string{"hi", "my", "name"}
	fmt.Println(arr, arr2)

	arr3 := []string{"hi", "my", "name"}
	arr3 = append(arr3, "is", "pit")
	fmt.Println(arr3)

}

func main() {
	todo()
	//output: [1 2 3 4] [hi my name]
	//        [hi my name is pit]

}
