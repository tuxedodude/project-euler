/*
https://projecteuler.net/problem=10
Summation of primes
Problem 10
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
	"fmt"
)

//runs callback as each prime is generated, below n
func sieve_apply(n int, callback func(x int)) {

	if n < 2 {
		return
	}

	//make sieve
	s := make([]bool, n+1)

	for i := 2; i < n; i++ {
		if s[i] == false { //this element is prime

			callback(i)

			//mark all multiples
			for j := 1; j*i <= n; j++ {
				s[i*j] = true
			}
		}
	}
}

func Solve(n int) int64 {
	var sum int64
	sieve_apply(n, func(x int) { sum += int64(x) })
	return sum
}

func main() {
	fmt.Println(Solve(2000000))
}
