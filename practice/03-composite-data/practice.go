package main

import "fmt"

func main() {

	// Using a composite literal
	// - create an array which holds 5 values of type int
	// - assign values to each index position
	// Range over the array and print the values out
	// using format printing
	// - print out the type of array
	one()
	// Wouldn't normally create an array

	// Using a composite literal
	// - create a slice of type int
	// - assign 10 values
	// Range over the slice and print the values out
	// using format printing
	// - print out the type of slice
	two()

	// Take code from previous example
	// slice the slice from the previous example
	// display the following
	// [2, 4, 8, 16, 32]
	// [64, 128, 256, 512, 1028]
	// [8, 16, 32, 64, 128, 256]
	// [4, 8, 16, 32, 64, 128]
	three()

	// Provided a given slice append to the slice the value of 52
	// print out the slice
	// in ONE statement append to that slice these values
	// 53, 54, 55
	// print out the slice
	// append to the slice x:=[]int{56,57,58,59,60}
	// print out the slice
	four()

	// To delete from a slice we use append() along with slicing.
	// Start with slice x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	// use append and slicing to get the following values
	// [42,43,44,48,49,50,51]
	five()

	// Create a slice to store the names of all the states in the US.
	// Use make and append to do this.
	// Goal: do not have the array that underlines the slice created more than once.
	// What is the length of your slice?
	// what is the capacity?
	// Print out all the values, along with their index position, without using the range clause.
	six()

	//seven()
	//eight()
	//nine()
	//ten()
}

func one() {
	a := [5]int{2, 4, 8, 16, 32}
	for i, v := range a {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of a is: %T\n", a)
}

func two() {
	s := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1028}
	fmt.Println(len(s))
	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", s)
}

func three() {
	s := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1028}
	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:8])
	fmt.Println(s[1:7])
}

func four() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

	fmt.Println(x)
}

func five() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func six() {
	// This works, but I'm concerned about having the change the len of the slice over and over again.
	x := make([]string, 0, 50)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, `Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
