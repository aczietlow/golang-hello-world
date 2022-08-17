package main

import "fmt"

func main() {
	fmt.Println(foo5())
	fmt.Println(bar5())
}

func foo5() int {
	return 5
}

func bar5() (int, string) {
	return 5, "random string"
}
