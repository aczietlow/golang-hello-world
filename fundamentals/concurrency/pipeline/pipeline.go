package main

import "fmt"

// [counter] ----> [squarer] ----> [printer]

// prints infinite numbers
func main() {
	// create 2 channels

	// when used to connect the output of 1 goroutine as the input of another go routine, they're called "pipelines"
	naturalNums := make(chan int)
	squareNums := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturalNums <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturalNums
			squareNums <- x * x
		}
	}()

	// Printer
	for {
		fmt.Println(<-squareNums)
	}
}
