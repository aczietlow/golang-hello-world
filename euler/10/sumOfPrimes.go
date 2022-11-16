package main

import "fmt"

// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//
// Find the sum of all the primes below two million.
func main() {
	primes := []int{2}
	sum := 2
	max := 2000000
	for i := 3; i <= max; i++ {
		if isPrime(i) {
			primes = append(primes, i)
			sum += i
		}
	}
	fmt.Printf("Sum %v\n", sum)
}

func isPrime(num int) bool {
	if num%2 == 0 {
		return false
	}

	factor := 3
	for factor*factor <= num {
		if num%factor == 0 {
			return false
		}

		factor += 2
	}

	return true
}
