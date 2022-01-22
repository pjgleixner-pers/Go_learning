package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super Heavy")
	}
}

func main() {
	go heavy() //run in background
	/*
		runs go heavy, finds fmt.Println("Fin") and finds out its done
	*/
	go superHeavy()
	fmt.Println("Fin")
	time.Sleep(time.Second * 5) //will print 4 Heavy and 2 Super Heavy
	//select{} //will run forever
}
