package main

import "fmt"

const (
	// throrw away the 0 iota
	_ = iota
	// iota == 1, multiply by 10, right shift bit by 10.
	kb = 1 << (iota * 10)
	// iota == 2, multiply by 10, right shift bit by 20.
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
	tb = 1 << (iota * 10)
)

func main() {
	foo := 4
	// %d formats int as decimal
	// %b formats int as binary
	fmt.Printf("%d\n%b\n", foo, foo)
	memoryExample()
	memoryExampleWithIota()
}

func memoryExample() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t%b\n", kb, kb)
	// increases by 10 0's aka right bit shift.
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

func memoryExampleWithIota() {
	println("with iota\n")
	fmt.Printf("%d\t\t%b\n", kb, kb)
	// increases by 10 0's aka right bit shift.
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t%b\n", tb, tb)
}
