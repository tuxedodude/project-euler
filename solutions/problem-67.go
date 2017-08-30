/*
Maximum path sum II
Problem 67 
By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.

3
7 4
2 4 6
8 5 9 3

That is, 3 + 7 + 4 + 9 = 23.

Find the maximum total from top to bottom in triangle.txt (right click and 'Save Link/Target As...'), a 15K text file containing a triangle with one-hundred rows.

NOTE: This is a much more difficult version of Problem 18. It is not possible to try every route to solve this problem, as there are 299 altogether! If you could check one trillion (1012) routes every second it would take over twenty billion years to check them all. There is an efficient algorithm to solve it. ;o)
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

const SOLUTION_STRING string = `75
95 64
17 47 82
18 35 87 10
20 04 82 47 65
19 01 23 75 03 34
88 02 77 73 07 63 67
99 65 04 28 06 16 70 92
41 41 26 56 83 40 80 70 33
41 48 72 33 47 32 37 16 94 29
53 71 44 65 25 43 91 52 97 51 14
70 11 33 28 77 73 17 78 39 68 17 57
91 71 52 38 17 14 91 43 58 50 27 29 48
63 66 04 68 89 53 67 30 73 16 69 87 40 31
04 62 98 27 23 09 70 98 73 93 38 53 60 04 23`

func parse_triangle(s string) (result [][]int, err bool) {

	// assume every newline denotes a new row
	strrows := strings.Split(s, "\n")

	triangle := make([][]int, len(strrows))

	for i := 0; i < len(triangle); i++ {
		triangle[i] = make([]int, i+1)
	}

	for i, row := range strrows {

		r := strings.Split(row, " ")

		// error, row sizes mismatched
		if len(r) != i+1 {
			return nil, true
		}

		// parse each element to place in row
		for j, e := range r {

			ele, converr := strconv.Atoi(e)

			// error: cannot convert to int
			if converr != nil {
				return nil, true
			}

			triangle[i][j] = ele
		}
	}

	return triangle, false
}

func printrows(s [][]int) {
	for _, r := range s {
		fmt.Println(r)
	}
}

// returns sum of the largest path
func Largest_path(triangle [][]int) int {
	top := len(triangle)

	paths := make([]int, top)

	for R := 0; R < top; R++ {

		row := triangle[R]

		sums := make([]int, len(row))

		for j, ele := range row {

			var x int
			// choose correct parent to link up
			// first element of row
			if j == 0 {
				x = j
				// last element of row
			} else if j == len(row)-1 || (paths[j] < paths[j-1]) {
				x = j - 1
				// right parent
			} else {
				x = j
			}

			sums[j] = paths[x] + ele
		}

		// move accumulated paths to buffer
		for i, ele := range sums {
			paths[i] = ele
		}

		fmt.Println(paths)

	}
	return _largest(paths)
}

func _largest(sums []int) int {
	max_sum := 0

	// pick the index of the path of the greatest sum
	for _, n := range sums {
		if n > max_sum {
			max_sum = n
		}
	}

	// give the path
	return max_sum
}

func main() {
	t, _ := parse_triangle(SOLUTION_STRING)

	total := Largest_path(t)

	fmt.Println("Result", total)
}
