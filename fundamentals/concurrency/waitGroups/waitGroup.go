package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	/*
		Since the program will exit when it reaches the end of main(), this go routine can be left before it's completed its
		execution. This is the purpose of the WaitGroup.
		@SeeAlso Mutex
	*/
	go foo()

	bar()
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
