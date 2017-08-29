/* Longest Collatz sequence
Problem 14
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even) n → 3n + 1 (n is odd)
Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.*/

package main

import (
	"fmt"
)

// mapping of results to lengths
var cmap map[int64]int64

func unwind(seq []int64, count int64) int64 {

	c := count

	for i := len(seq) - 1; i >= 0; i, c = i-1, c-1 {
		x := seq[i]

		_, member := cmap[x]
		if !member {
			cmap[x] = c
		}
	}

	return count
}

func collatz(x int64) int64 {
	if x <= 0 {
		return 0
	}

	var n, i int64 = x, 0

	// list of numbers generated in this sequence
	seq := []int64{n}

	for ; n != 1; i++ {

		// add to list to be unwound (memoized) later
		seq = append(seq, n)

		_, member := cmap[n]
		if member {
			return unwind(seq, i+cmap[n])
		}

		//fmt.Println(n)
		if n%2 == 0 { // even
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return unwind(seq, i+1)
}

func longest_collatz(upto int64) int64 {
	var i, maxcollatz, maxi int64 = 2, 0, 0

	for ; i < upto; i++ {
		result := collatz(i)
		if result > maxcollatz {
			maxcollatz = result
			maxi = i
		}
	}
	return maxi
}

func main() {
	cmap = make(map[int64]int64)

	fmt.Println("Longest collatz < 1,000,000:", longest_collatz(1000000))
}
