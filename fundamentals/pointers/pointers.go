package main

import "fmt"

func main() {
	/*
		Pointers are the address on memory of a variable. The `&` operator will return the address value of a variable
		The `*` retrieves the variable from a given memory address.
	*/
	warp := "eleven"
	fmt.Println(warp)
	fmt.Println(&warp)
	fmt.Printf("%T\n", warp)  // string
	fmt.Printf("%T\n", &warp) // *string (a pointer to where a string is stored)
	fmt.Println(*&warp)       // eleven
	fmt.Println(*&*&*&warp)   // eleven (Just because we can, doesn't mean we should)

	warpAddy := &warp
	fmt.Println(warpAddy)
	fmt.Println(*warpAddy)

	parallelWarp := &warp
	*parallelWarp = "ten but more"
	fmt.Println(warp) // ten but more.

}
