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
	//"math"
)

//runs callback as each prime is generated, up to n inclusive
func sieve_apply(n int, callback func(int)) {

	if n < 2 {
		return
	}

	//make sieve
	s := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		if s[i] == false { //this element is prime
			callback(i) // callback on prime number
			//mark all multiples
			for j := 1; j*i <= n; j++ {
				s[i*j] = true
			}
		}
	}
}

// return slice containing all prime numbers up to n inclusive,
// and a map whose keys are the same primes.
func _primes(n int) []int {
	p := make([]int, 0)

	sieve_apply(n,
		func(prime int) {
			p = append(p, prime)
		})

	return p
}

func no_common_factor(p, q int, primes []int) bool {
	for i := 0; i < len(primes) && primes[i] <= q; i++ {
		if (p%primes[i]) == 0 && (q%primes[i]) == 0 {
			//common divisor
			return false
		}
	}
	return true
}

func make_triple(p, q int) (int, int, int) {

	pp, qq := p*p, q*q

	a := p * q
	b := (pp - qq) / 2
	c := (pp + qq) / 2

	if a < b {
		return a, b, c
	} else {
		return b, a, c
	}
}

func Solve(n int) []int {
	//primes := _primes(n)

	for p := 3; p < n/3; p += 2 {
		for q := 1; q < p; q += 2 {

			a, b, c := make_triple(p, q)
			if c*c == a*a+b*b {
				fmt.Println(a, b, c)
				if n%(a+b+c) == 0 {
					return []int{a, b, c}
				}
			}

		}
	}
	return []int{}
}

func main() {
	fmt.Println(Solve(100))
}
