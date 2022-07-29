package main

import "fmt"

// func (r receiver) identifier(parameters) (returns(s)) {...}
// when defining a function, it's called parameters
// when calling a defined function and passing variables, their "arguments"

func main() {
	sum := addNum(2, 6)
	println(sum)
	functions()
	println(rectProps2(1.8, 4.0))
	variadic(1, 2, 3, 4)
}

func functions() {
	// Can use the blank identifier if we only want 1 of a functions return values
	area, _ := rectProps2(15.2, 15.6)
	println(area)
}

//func functionname(parametername type) returntype {
func addNum(a int, b int) int {
	var sum = a + b
	return sum
}

// multiple return types
// func function name(paaametername type) (float32, float64)
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// return values can be named
func rectProps2(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

// Blank identifier is _

// Everything in go is passed by value.

// Variadic functions
func variadic(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// Receives the variables in a slice of the type of the parameters. ([]int)

	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Printf("The total: %v\n", sum)
}
