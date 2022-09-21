package main

import "fmt"

func main() {
	// Strings are an immutable series of bytes.
	s := "this is a string of things"
	fmt.Println(len(s))  // 26 (number of bytes)
	fmt.Println(s[5:16]) // is a string
	fmt.Printf("%T\n", s)

	// the ith value in the string slice is the ith byte, not necessarily the character rune.
	fmt.Println(s[0])

	// Strings are immutable, but the concatenation operator exists.
	// This creates a new string
	t := s
	s += " and stuff."
	fmt.Println(t)
	fmt.Println(s)

	// Because strings are immutable, it's safe for two copies of a string to share the underlying memory.
	// By default, strings are interpreted as unicode-8

	// String literals are encapsulated with double quotes.
	// \ can be used to insert arbitrary byte codes such as \n\r \t - tab \v - vertical tab \a - alert bell
	q := "hello motto"
	fmt.Println(q)
	// with \ insert bytes using hexadecimal or octal escapes.
	// hexadecimal escape would be xhh where hh is 2 hexadecimal digits (upper or lowercase works)
	q += "\x21\x21"
	fmt.Println(q)

	// Octal escapes use the format ooo where o is a an octal code
	q += "\040\046\040\115\157\162\145"
	fmt.Println(q)

	// Raw string literal use ``
	r := `This is interpreted 
as it is entered with out \any \n escapes`

	fmt.Println(r)

	// Unicode code points may be a single or multiple bytes
	// A single unicode code point in go is referred to as a rune
	// Rune is an analogous to int32
	var a rune
	fmt.Printf("%T\n", a)

	// Unicode 16 bit is denoted with uhhhh
	// Unicode 32 bit is denoted with Uhhhhhhhh where h is a hexadecimal digit
	z := "\u4e16\u754c"
	fmt.Println(z)

	// Unicode escapes can be used for rune literals
	a = '\U00004e16'
	fmt.Println(a)
	fmt.Println(string(a))

	// Unicode's characters under 256 can be denoted with xhh
	x := '\x41'
	fmt.Println(x)
}
