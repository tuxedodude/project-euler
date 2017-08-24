//https://play.golang.org/p/lyTnGqZf_F

package main

import (
	"fmt"
)

var sieve_call_count int = 0

/* https://projecteuler.net/problem=5
 * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 * What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20? */

//runs callback as each prime is generated, up to n inclusive
func sieve_apply(n int, callback func(int)) {

	// this is the most expensive function, so count our calls to it
	defer func() { sieve_call_count++ }()

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
func primes(n int) ([]int, map[int]int) {
	p := make([]int, 0)
	m := make(map[int]int)

	sieve_apply(n,
		func(prime int) {
			p = append(p, prime)
			m[prime] = 1
		})

	return p, m
}

func concat_slice(a, b []int) []int {
	for _, e := range b {
		a = append(a, e)
	}
	return a
}

/* is this number a member of the slice of primes supplied? */
func isprime(x int, pmap map[int]int) bool {
	_, prime := pmap[x]
	return prime
}

func factorize(x int, plist []int, pmap map[int]int) []int {
	s := make([]int, 0)

	if x < 2 {
		// empty set, no prime factors
		return s
	} else if isprime(x, pmap) {
		// base case: if x is prime, it's fully factored
		return append(s, x)
	}

	// recursion case: return the list of factors of x,
	// the first of which is the lowest prime that will divide into x,
	// and the rest which is the quotient factorized recursively.
	for _, p := range plist {
		if x%p == 0 {
			s = append(s, p)
			rest := factorize(x/p, plist, pmap)
			s = concat_slice(s, rest)
			return s
		} else if p > x {
			break
		}
	}
	s[0] = x
	return s
}

// assumes a sorted list of ints!
// calls back on each sub-count
func count_elems(s []int, callback func(element, count int)) {

	// empty slice, do nothing
	if len(s) == 0 {
		return
	}

	var e, e1, c = s[0], s[0], 0

	// iterate over the slice, counting the frequency
	// of the first element
	for ; c < len(s); e1, c = e, c+1 {
		if s[c] == e1 {
			e = s[c]
		} else {
			break
		}
	}
	// update our counts
	callback(e1, c)

	// if there are remaining elements that do not equal
	// the first element, recursively count their frequency
	if c < len(s) {
		defer count_elems(s[c:], callback)
	}
}

func intpow(x, exp int) int {
	if x <= 0 {
		return 0
	}
	prod := x
	for i := 1; i < exp; i++ {
		prod *= x
	}
	return prod
}

func Solve(x int) int {

	// p is a slice of primes;
	// pm is a map of prime to prime frequency
	p, pm := primes(x)

	// create a function to update our prime factor count
	update_counts_cb := func(prime, count int) {
		if count > pm[prime] {
			pm[prime] = count
		}
	}

	for i := 2; i <= x; i++ {
		factors := factorize(i, p, pm)
		count_elems(factors, update_counts_cb)
	}

	//fmt.Println(pm)

	product := 1
	for k, v := range pm {
		product *= intpow(k, v)
	}
	return product

}

func Printsolution(n int) {
	fmt.Println("n =", n, "\t->", Solve(n))
}

func main() {

	Printsolution(10)
	Printsolution(20)
	Printsolution(25)

	fmt.Println("Calls to sieve:", sieve_call_count)
}