package main

import "fmt"

func main() {
	// Arrays
	// ArrayType = [length]ElementType
	var list [5]int
	fmt.Println(list)

	// Slices are a composite of an array
	// Composite literal syntax `t := type{values}`
	listX := []int{1, 2, 5, 8}
	fmt.Println(listX)

	// range allows looping through slices (as well as arrays, maps, strings, and reading from channels)
	for index, value := range listX {
		fmt.Println(index, value)
	}

	// Make function syntax for creating a new slice
	// creates an array of length 5, and returns a slice of length 3, with a capacity of 5.
	var listY = make([]string, 3, 5)
	for i, v := range []string{"foo", "bar", "baz"} {
		listY[i] = v
	}

	fmt.Println(listY)
	fmt.Println("length ", len(listY))
	fmt.Println("capacity", cap(listY))

	// Add items to array with append.
	listOdd := []int{1, 3, 5, 7, 8}
	listOdd = append(listOdd, 9, 11, 13, 15)
	fmt.Println(listOdd)
	// To add a slice to a slice.
	newOddNumbers := []int{17, 19, 21}
	listOdd = append(listOdd, newOddNumbers...)
	// `...` can be appended to an array, slice, map, etc to unfurl values. Useful to pass in values of list as a variadic parameter
	fmt.Println(listOdd)

	// Slicing slices
	fmt.Println(listOdd[0])
	fmt.Println(listOdd[1])
	// slice[n:m] n = index position, m = cap of slice
	fmt.Println(listOdd[2:5])
	// From the beginning to the specified index
	fmt.Println(listOdd[:5])
	// From the specified index to the end
	fmt.Println(listOdd[5:])

	// Deleting element from slice.
	listOdd = append(listOdd, 22, 24, 23, 25)
	fmt.Println(listOdd)
	fmt.Println(listOdd[:12])
	fmt.Println(listOdd[14:])
	// Just append slices of the slice to remove elements
	realOdd := append(listOdd[:12], listOdd[14:]...)
	fmt.Println(realOdd)

}
