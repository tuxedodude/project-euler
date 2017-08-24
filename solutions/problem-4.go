// https://projecteuler.net/problem=4
// Largest Palindrome Product
// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

// To Do: faster implementation of next_pal()

package main

import (
	"fmt"
	"math"
)

// return least significant digit of this decimal number
func lsd(x int) int {
	y := (x / 10) * 10
	return x - y
}

// convert number to a slice of (decimal) digits
func digits_of(n int) []int {

	s := make([]int, 0)

	for x := n; x != 0; x /= 10 {
		s = append(s, lsd(x))
	}
	return s
}

func is_palindrome(n int) bool {
	s := digits_of(n)
	top := len(s)
	for i := 0; i < top/2; i++ {
		if s[i] != s[top-i-1] {
			return false
		}
	}
	return true
}

// next smallest palindrome below n
func next_pal(n int) int {

	// find the largest palindrome <= upper bound
	// naive implementation
	for i := n - 1; i > 1; i-- {
		if is_palindrome(i) {
			return i
		}
	}
	return 0
}

// return lowest and highest factors of the palindrome
func lo_hi_factors(digits int) (lo, hi int) {
	x := int(math.Pow(10, float64(digits))) - 1
	y := int(math.Pow(10, float64(digits-1)))
	return y, x
}

// return the low and high palindrome bounds
func search_bounds(digits int) (lower, upper int) {
	lo, hi := lo_hi_factors(digits)
	return lo * lo, hi * hi
}

//is number n a product of 2 factors of d digits?
func factorable(n, d int) bool {
	//largest number of (d) digits
	lo, hi := lo_hi_factors(d)

	for i := lo; i <= hi; i++ {
		q := n / i
		if n%i == 0 && q >= lo && q <= hi {
			return true
		}
	}

	return false
}

// return the largest palindrome that is the product of two n-digit numbers
func Solve(n int) int {
	lo, hi := search_bounds(n)

	for pal := next_pal(hi + 1); pal >= lo; pal = next_pal(pal) {
		if factorable(pal, n) {
			return pal
		}
	}
	return -1
}

func main() {
	fmt.Println(Solve(3))
}