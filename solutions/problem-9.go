// unfinished!!! need to account for multiples of found triples

/*https://projecteuler.net/problem=9
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.*/
package main

import (
	"fmt"
)

func gcd(x, y int) int {
	if x == 0 || y == 0 {
		return 0
	}

	if x < y {
		x, y = y, x
	}

	for ; x%y != 0; x, y = y, x%y {
		//fmt.Println(">", x, "%", y, "=", x%y)
	}
	return y
}

//a = mn, b = (m^2 - n^2) / 2, c = (m^2 + n^2) / 2
// Assumes m > n > 0; m and n are odd
func make_triple(m, n int) []int {
	mm, nn := m*m, n*n
	return []int{m * n, (mm - nn) / 2, (mm + nn) / 2}
}

// execute callback function on each generated triple up to upper bound;
// return value of TRUE will stop computation
func pythagorean_triples(upperbound int, cb func([]int) bool) {

	for m := 3; m < upperbound; m += 2 {
		for n := 1; n < m; n += 2 {
			t := make_triple(m, n)

			if cb(t) {
				return
			}
		}
	}
}

// returns slice containing pythagorean triple
// s.t. a+b+c == n; returns empty slice if none found
func Solve(n int) []int {
	var triple []int

	pythagorean_triples(n,
		func(t []int) bool {
			s := sum(t)
			if s == n {
				triple = t
			}
			return s == n
		})

	return triple
}

func coprime(x, y int) bool {
	return gcd(x, y) == 1
}

func sum(s []int) int {
	sum := 0
	for _, e := range s {
		sum += e
	}
	return sum
}

func main() {
	fmt.Println(Solve(1000))
}

