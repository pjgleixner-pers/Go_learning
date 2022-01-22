package main

import (
	"fmt"
)

func main() {
	//continue
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
		// output:
		// 		1
		// 		3
		// 		5
		// 		7
		// 		9
	}

	//break
	flag := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag = false
			break
		}
		fmt.Println(i)
		//output: 1
	}
	fmt.Println(flag)
	//output: false

	//switch
	day := "fri"
	switch day {
	case "fri":
		fmt.Println("TGIF")
	case "mon", "tue", "wen":
		fmt.Println("Boring")
	default:
		fmt.Println("default")
	}
	//output: TGIF

	switch {
	case day == "fri":
		fmt.Println("TGIFF")
		break
	}
	//output: TGIFF
}
