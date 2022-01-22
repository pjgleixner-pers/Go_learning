package main

import (
	"fmt"
	"sync"
)

func main() {
	//wait groups
	wg := &sync.WaitGroup{}
	wg.Add(2)
	//mutexes (critical section)
	/* mut := &sync.Mutex{}
	mut.Lock()
	mut.Unlock() */

	//Lambda functions
	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()

	fmt.Println("Fin")
	wg.Wait()
	fmt.Println("Exit")

}
