package main

import (
	"fmt"
	"github.com/aczietlow/stringutil"
	"math/big"
)

// Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.
func main() {
	limit, count, sum := big.NewInt(1000000), big.NewInt(1), big.NewInt(0)
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
