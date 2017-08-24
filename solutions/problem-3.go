// https://projecteuler.net/problem=3
// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import (
	"fmt"
	"math"
)

const example_number int64 = 13195

const solution_number int64 = 600851475143

/* returns slice of ints containing all primes below n */
func primes(n int64) []int64 {
	// no primes below 2!
	if n <= 2 {
		return make([]int64, 0)
	}

	/* base case: primes(2) -> [];
	 * primes(3) returns [2] */

	top := int(math.Sqrt(float64(n)) - 2) // top of array

	s := make([]bool, top) // sieve

	//preallocate to avoid growing
	primes := make([]int64, 0)

	for i, _ := range s {
		if s[i] == false {
			primes = append(primes, int64(i+2))
		}

		for k := 1; (k*(i+2) - 2) < top; k++ {
			s[k*(i+2)-2] = true
		}
	}
	return primes
}

/* greatest prime factor */
func Greatest_prime_factor(n int64) int64 {

	p := primes(n + 1)

	if len(p) == 0 {
		return 0
	}

	for i := len(p) - 1; i >= 0; i-- {
		if n%p[i] == 0 {
			return p[i]
		}
	}
	return 0
}

func main() {
	//fmt.Println(Greatest_prime_factor(example_number))
	fmt.Println(Greatest_prime_factor(solution_number))
}
