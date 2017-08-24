/* https://projecteuler.net/problem=7
 * find the 10001st prime! */

package main

import (
	"fmt"
	"math"
)

const target int = 10001

//runs callback as each prime is generated, up to n inclusive
func sieve_apply(n int, callback func(ith, x int) bool) {

	if n < 2 {
		return
	}

	//make sieve
	s := make([]bool, n+1)

	count := 0
	for i := 2; i <= n; i++ {
		if s[i] == false { //this element is prime
			count++
			//fmt.Println(i)

			if callback(count, i) {
				return
			}
			//mark all multiples
			for j := 1; j*i <= n; j++ {
				s[i*j] = true
			}
		}
	}
}

// ugly...
func guess_prime_count(n int) int {
	// prime counting function is approximately
	// num_primes(x) = x / ln(x)
	// given that we want n primes, find x

	i := n

	for {
		x := float64(i)
		result := x / math.Log(x)
		if int(result) < n {
			i *= 2
		} else {
			return i
		}
	}
}

func Solve(n int) int {
	x := 0
	count := 0
	sieve_apply(guess_prime_count(n), func(ith, prime int) bool {
		count = ith
		if ith == n {
			x = prime
			return true
		} else {
			return false
		}
	})

	//fmt.Println(count)
	return x
}

func main() {
	fmt.Println(Solve(target))
}
