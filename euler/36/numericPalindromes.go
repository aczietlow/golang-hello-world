package main

import (
	"fmt"
	"github.com/aczietlow/stringutil"
	"math/big"
)

// Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
func main() {
	// Use big Int to leverage easy String and Text methods. String comp are faster then numeric computations.
	limit, count, sum := big.NewInt(1000000), big.NewInt(1), big.NewInt(0)
	// Count is 2, b/c since a base 2 number can't have compare leading 0's, implies base 2 numbers will never end in 0.
	for i := big.NewInt(1); i.Cmp(limit) == -1; i.Add(i, count) {
		if isPalindromes(*i) == true {
			sum = sum.Add(sum, i)
		}
	}

	fmt.Println(sum)
}

// Find palindromes that are strings.
func isPalindromes(i big.Int) bool {
	s1, s2 := i.Text(2), i.Text(10)
	return s1 == stringutil.Reverse(s1) && s2 == stringutil.Reverse(s2)
}
