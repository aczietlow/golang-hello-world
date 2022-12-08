package main

import "fmt"

// prints infinite numbers
func main() {
	naturalNums := make(chan int)
	squareNums := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturalNums <- x
		}
		// Communicate to other channels using this channel, that we're done; No more values will be sent, and it can
		// safely be "drained"
		close(naturalNums)
	}()

	// Squarer
	go func() {
		// channels can be ranged over. Which will receive any input sent down the channel. Once the close signal
		// is sent the range iterator will be closed.
		for x := range naturalNums {
			squareNums <- x * x
		}
		close(squareNums)
	}()

	// Printer
	for squares := range squareNums {
		fmt.Println(squares)
	}
}
