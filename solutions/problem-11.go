/*
https://projecteuler.net/problem=11
In the 20×20 grid below, four numbers along a diagonal line have been marked in red.
...
The product of these numbers is 26 × 63 × 78 × 14 = 1788696.
What is the greatest product of four adjacent numbers in the same direction (up, down, left, right, or diagonally) in the 20×20 grid?
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	//"regexp"
)

const GRID string = `08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48`

var calls_to_product int

const SIMPLE_MAX_PRODUCT bool = false

// assumes rows are denoted by newlines
// TO DO: split on white space only, then
// make grid of r, c dimension
func parse_grid(str string) [][]int {

	s := strings.Split(str, "\n")

	grid := make([][]int, len(s))

	for i, ss := range s {

		srow := strings.Split(ss, " ")

		row := make([]int, len(srow))

		// parse into a row of ints
		for j, e := range srow {
			x, _ := strconv.ParseInt(e, 10, 64)
			row[j] = int(x)
		}
		grid[i] = row
	}
	return grid
}

func countcalls() {
	calls_to_product++
}

func product(s []int) (prod, skip int) {
	defer countcalls()

	if len(s) == 0 {
		return 0, 0
	}
	p := 1
	for i := len(s) - 1; i >= 0; i-- {
		// any zero == 0 product
		if s[i] == 0 {
			return 0, i
		}
		p *= s[i]
	}
	return p, 0
}

/* finds largest product of n consecutive elements of slice s */
func max_product_v1(s []int, n int) int {

	if n > len(s) {
		return 0
	}

	max := 0
	for i := 0; i+n <= len(s); i++ {

		consec_n := s[i : i+n]
		//fmt.Println(consec_n)
		p, _ := product(consec_n)

		if p > max {
			max = p
		}
	}
	return max
}

func max_product_v2(s []int, n int) int {

	if n > len(s) {
		return 0
	}
	max := 0

	for i, last := 0, 0; i+n <= len(s); i++ {
		var p, skip int

		ss := s[i : i+n]

		first := s[i+n-1]

		if first > last {
			p, skip = product(ss)
		}

		if skip > 0 {
			i += skip
			last = 0
			continue
		} /*else if p == 0 {
			i += n - 1
			last = 0
			continue
		} */

		if p > max {
			max = p
		}

		last = s[i]
	}
	return max
}

func max_product(s []int, n int) int {
	if SIMPLE_MAX_PRODUCT {
		return max_product_v1(s, n)
	} else {
		return max_product_v2(s, n)
	}
}

func max_horizontals(grid [][]int, n int) int {
	max := 0
	for _, row := range grid {
		p := max_product(row, n)
		if p > max {
			max = p
		}
	}
	return max
}

func transpose(s [][]int) [][]int {

	rows, cols := matrix_dim(s)

	// empty matrix
	if rows == 0 {
		return [][]int{}
	}

	// make rows
	T := make([][]int, cols)

	// populate with empty columns
	for i, _ := range T {
		T[i] = make([]int, rows)
	}

	// transpose
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			T[j][i] = s[i][j]
		}
	}
	return T
}

func max_verticals(grid [][]int, n int) int {
	return max_horizontals(transpose(grid), n)
}

// returns dimensions of this matrix
func matrix_dim(s [][]int) (rows, cols int) {

	r, c := len(s), 0

	// gather smallest column dimension
	for i, row := range s {
		dim := len(row)
		if dim == 0 {
			c = 0
			break
		} else if i == 0 || dim < c {
			c = dim
		}
	}

	return r, c
}

func intmax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func diagonals(s [][]int, collect func([]int)) {
	//diagonals
	//D := make([]int, 0)

	rows, cols := matrix_dim(s)
	maxlen := intmax(rows, cols)

	// increment Diagonals, right descending, by rows
	for i := 0; i < maxlen; i++ {

		// diagonal, right descending
		a, b := make([]int, maxlen), make([]int, maxlen)

		// iterate over right descending diagonals
		for r, c, x := i, 0, 0; r < rows && c < cols; r, c, x = r+1, c+1, x+1 {
			a[x] = s[r][c]
		}
		collect(a)

		for r, c, x := 0, i+1, 0; r < rows && c < cols; r, c, x = r+1, c+1, x+1 {
			b[x] = s[r][c]
		}
		collect(b)

		a1, b1 := make([]int, maxlen), make([]int, maxlen)

		// do reverse diagonals
		for r, c, x := rows-i-1, 0, 0; r >= 0 && c < cols; r, c, x = r-1, c+1, x+1 {
			a1[x] = s[r][c]
		}
		collect(a1)

		for r, c, x := rows-1, i+1, 0; r >= 0 && c < cols; r, c, x = r-1, c+1, x+1 {
			b1[x] = s[r][c]
		}
		collect(b1)

	}
}

func max_diagonals(grid [][]int, n int) int {
	max := 0

	collector := func(s []int) {
		p := max_product(s, n)
		if p > max {
			max = p
		}
	}

	diagonals(grid, collector)

	return max
}

func printgrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println("----------")
}

func printsep() {
	fmt.Println("----------")
}

func showdiagonals(grid [][]int) {

	collector := func(s []int) {
		fmt.Println(s)
	}

	diagonals(grid, collector)
	printsep()
}

func Solve() int {
	s := parse_grid(GRID)
	h, v := max_horizontals(s, 4), max_verticals(s, 4)
	d := max_diagonals(s, 4)

	m := []int{h, v, d}

	//fmt.Println(calls_to_product)
	return max_product(m, 1)
}

func main() {

	fmt.Println(Solve())

}
