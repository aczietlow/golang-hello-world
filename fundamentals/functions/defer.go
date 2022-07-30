package main

import "fmt"

func main() {
	// defer will invoke a function whose execution is deferred until
	// the surrounding function reaches a return statement
	// the surrounding function reaches the end of the function body
	// or b/c goroutine is panicking

	// Can be used similar to __destruct() in php
	defer foo() // <-- deferred until main() reached the end of the func body
	bar()       // <-- executes before foo()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
