package main

import "fmt"

const (
	x        = 1
	y    int = 2
	year     = 2021
)
const (
	_  = iota
	y1 = year + iota
	y2 = year + iota
	y3 = year + iota
	y4 = year + iota
)

func main() {
	numericalSystem()
	operators()
	createConstants()
	variablesAndBitShifting()
	rawStringLiteral()
	useIota()
}

func numericalSystem() {
	// Write a program that prints a number in decimal, binary, and hex
	z := 42876
	fmt.Printf("decimal %d\n", z)
	fmt.Printf("hexidecimal %X\n", z)
	fmt.Printf("hexidecimal %#X\n", z)
	fmt.Printf("binary %b\n", z)
}

func operators() {
	/*
		Using the following operators, write expressions and assign their values to variables:
		==
		<=
		>=
		!=
		<
		>
		Now print each of the variables.
	*/
	a := 41 == 42
	b := 41 <= 42
	c := 41 >= 42
	d := 41 != 42
	e := 41 < 42
	f := 41 > 42

	fmt.Printf("value %v\ttype %T\t\n", a, a)
	fmt.Printf("value %v\ttype %T\t\n", b, b)
	fmt.Printf("value %v\ttype %T\t\n", c, c)
	fmt.Printf("value %v\ttype %T\t\n", d, d)
	fmt.Printf("value %v\ttype %T\t\n", e, e)
	fmt.Printf("value %v\ttype %T\t\n", f, f)
}

func createConstants() {
	// Create TYPED and UNTYPED constants. Print the values of the constants.
	fmt.Printf("value %v\ttype %T\t\n", x, x)
	fmt.Printf("value %v\ttype %T\t\n", y, y)
}

func variablesAndBitShifting() {
	//Write a program that
	//assigns an int to a variable
	//prints that int in decimal, binary, and hex
	//shifts the bits of that int over 1 position to the left, and assigns that to a variable
	//prints that variable in decimal, binary, and hex
	bar := 2
	fmt.Printf("D-%d\t B-%b\t H-%#x\n", bar, bar, bar)
	baz := bar << 1
	fmt.Printf("D-%d\t B-%b\t H-%#x\n", baz, baz, baz)
}

func rawStringLiteral() {
	// Create a variable of type string using a raw string literal. Print it.
	q := `there 
is 
no 
Zule`
	println(q)
}

func useIota() {
	// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
	println(y1, y2, y3, y4)
}
