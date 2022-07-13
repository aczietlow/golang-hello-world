package main

import "fmt"

// [] is a composite literal
func main() {
	// make function
	sliceMake()

	// multi-dimension slices
	multi()
}

func sliceMake() {

	// Reminder that capacity is the size of the underlining array. If the len of the slice exceeds the underlying array
	// a new array is created and values are copied to the new array. New capacity is cap * n where n is the number of times cap has been exceeding
	// This is computationally complex and should be avoided.
	s := make([]int, 10, 100)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s[0] = 41
	s[8] = 999
	s[9] = 999999

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	// Add an item to the slice.
	s = append(s, 111)

	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func multi() {
	job1 := []string{"James", "Bond", "Aston Martin", "007"}
	job2 := []string{"Jennie", "James", "Porsche", "004"}

	fmt.Println(job1)
	fmt.Println(job2)

	// a slice of a slice of strings.
	agency := [][]string{job1, job2}

	fmt.Println(agency)
}
