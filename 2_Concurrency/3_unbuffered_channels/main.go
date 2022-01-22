package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	//<name> chan <datatype>
	//var c chan int
	//send
	go func() {
		c <- 1
	}()
	//sniff
	val := <-c
	fmt.Println(val)

	go func() {
		c <- 2
	}()

	time.Sleep(time.Second * 2)
	val = <-c
	fmt.Println(val)

	fmt.Println(c)
}
