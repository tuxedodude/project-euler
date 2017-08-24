// https://projecteuler.net/problem=1
// If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
	"fmt"
)

func sum_to_n(n int) int {
	return n * (n + 1) / 2
}

/* returns the sum of the multiples of n below max. */
func sum_of_multiples(n, max int) int {

	var r int

	if max%n > 0 {
		r = max / n
	} else {
		r = max/n - 1
	}

	return n * sum_to_n(r)
}

func Solve() int {
	return sum_of_multiples(3, 1000) + sum_of_multiples(5, 1000) - sum_of_multiples(3*5, 1000)
}

func main() {
	fmt.Println("The sum of all the multiples of 3 or 5 below 1000 is: ", Solve())
}
