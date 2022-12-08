package main

import "fmt"

// The channel type `chan<-` is a "send only" channel
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

// The channel type `<-chan` is a "receive only" channel
func squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for squares := range in {
		fmt.Println(squares)
	}
}

// prints infinite numbers
func main() {
	naturalNums := make(chan int)
	squareNums := make(chan int)

	go counter(naturalNums)
	go squarer(naturalNums, squareNums)
	// Printer
	printer(squareNums)
}
