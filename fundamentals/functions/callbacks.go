package main

import "fmt"

func main() {
	x := sum1(1, 1, 2, 4, 8, 16)
	fmt.Println(x)

	numbs := []int{3, 6, 9, 12, 15, 18, 21, 24, 27}
	y := sum1(numbs...)
	fmt.Println(y)

	// Passing in a callback as an argument
	yEven := even(sum1, numbs...)
	fmt.Println(yEven)
}

func sum1(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// Sum up the even numbers in a list
func even(f func(xi ...int) int, vi ...int) int {
	var y1 []int
	for _, v := range vi {
		if v%2 == 0 {
			y1 = append(y1, v)
		}
	}
	return f(y1...)
}
