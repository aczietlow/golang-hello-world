package main

import "fmt"

func main() {
	// Using the short declaration operator, assign these values to variables with the identifiers x,y,z
	// 42, "james bond", true
	// print the values stored in those variables using a single print statement & multiple print statements
	basicVariables()

	// use var to declare 3 variables
	// the variables should have package level scope
	// do not assign values
	// use the following identifiers
	// x is int
	// y is string
	// z is bool
	instantiateVarialbes()

	// Assign values to variables at the package level.
	// print all of the values into a single string
	// assign the returned values of type string using the short declaration operator to a variable with the identifier "s"
	assignVariables()

	// create your own variable type
	// create a var of your new type with identifier "x"
	// print the value and the type of x
	// assign 42 to "x" using the = operator
	// print out the value of x
	customVarType()

}

func basicVariables() {
	x := 42
	y := "james bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Printf("x is %v\n", x)
	fmt.Printf("y is %v\n", y)
	fmt.Printf("z is %v\n", z)
}

func instantiateVarialbes() {
	var (
		x int
		y string
		z bool
	)

	fmt.Printf("x is %v\n", x)
	fmt.Printf("y is %v\n", y)
	fmt.Printf("z is %v\n", z)
}

func assignVariables() {
	var (
		x int
		y string
		z bool
	)

	x = 42
	y = "james bond"
	z = true

	s := fmt.Sprintf("%v%v%v", x, y, z)
	fmt.Println(s)

}

func customVarType() {
	type linuxPoops int
	var x linuxPoops
	fmt.Printf("value %v\n type %T\n", x, x)
	x = 42
	// conversations
	var y int = 19
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
}
