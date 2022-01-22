package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quits chan struct{}) {
	//switch case
	//select for channel async
	for {
		time.Sleep(time.Second * 1)
		select {
		case <-c:
			fmt.Println("recived")
		case <-quits:
			fmt.Println("quit")
			os.Exit(0)
		}
	}

}

func main() {
	c := make(chan int, 2)
	quits := make(chan struct{})
	go func() {
		c <- 1
		//quits <- struct{}{}
	}()

	go Select(c, quits)
	select {}
}
