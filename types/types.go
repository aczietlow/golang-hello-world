package main

import (
	"fmt"
	"unsafe"
)

func main() {
	types()
}

func types() {
	var a bool = false
	b := true
	fmt.Println(a, "and", b)

	var c int = 42
	fmt.Println("C is an int:", c)

	// Signed integers
	// int8, int16, int32, int64 are all valid
	// rune is alias for int64
	var d int = -120
	fmt.Printf("D is %T and has a size of %d \n", d, unsafe.Sizeof(d))
	fmt.Println(d)

	// Unsigned integers
	// byte is alias for uint8
	var e uint8 = 50
	fmt.Printf("E is %T and has a size of %d \n", e, unsafe.Sizeof(e))
	fmt.Println(e)

	// Float
	var f float32 = 1.0001
	var g float64 = 1.0001
	fmt.Printf("f is %T \n", f)
	fmt.Printf("g is %T \n", g)

	// Complex
	c1 := complex(6, 7)
	c2 := 5 + 3i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)
}
