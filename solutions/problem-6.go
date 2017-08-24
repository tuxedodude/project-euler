/* https://projecteuler.net/problem=6
 * The sum of the squares of the first ten natural numbers is:
 * 12 + 22 + ... + 102 = 385
 * The square of the sum of the first ten natural numbers is:
 * (1 + 2 + ... + 10)2 = 552 = 3025
 * Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.
 * Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
 */

package main

import (
	"fmt"
)

func sum_to_n(n int64) int64 {
	return n * (n + 1) / 2
}

func square_of_sum(n int64) int64 {
	x := sum_to_n(n)
	return x * x
}

func sum_of_squares(n int64) int64 {
	var sum, i int64 = 0, 1
	for ; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func abs_int64(x int64) int64 {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func Solve(x int) int64 {
	xx := int64(x)
	a, b := sum_of_squares(xx), square_of_sum(xx)
	return abs_int64(a - b)
}

func Sum_square_difference(x int) int64 {
	return Solve(x)
}

func main() {
	fmt.Println(Solve(10))
	fmt.Println(Solve(100))
}
