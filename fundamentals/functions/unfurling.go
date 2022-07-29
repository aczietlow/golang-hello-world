package main

import "fmt"

func main() {
	ints := []int{1, 2, 3, 4, 5, 6, 7}
	// the ... can "unfurl" a slice (or any composite variable)
	total := sum(ints...)
	fmt.Println(total)

	// variadic is 0 to inf.
	t2 := sum()
	fmt.Println(t2)
}

func sum(x ...int) int {
	sum := 0
	fmt.Printf("type:%T\n", x)
	for _, v := range x {
		sum += v
	}
	return sum
}
