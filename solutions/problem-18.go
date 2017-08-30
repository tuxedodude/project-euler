/*

https://projecteuler.net/problem=18

Find the maximum total from top to bottom of the triangle below:

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

// for each step of the triangle, we find the optimal path
// store the path, and propagate the sum to avoid recalculating
// at each step
type pathnode struct {
	sum  int
	path []int
}

// returns path with the largest sum and the sum.
func Largest_path(triangle [][]int) ([]int, int) {
	top := len(triangle)

	paths := make([]pathnode, top+1)
	//paths[0] = pathnode{triangle[0][0], []int{triangle[0][0]}}

	for R := 0; R < top; R++ {

		row := triangle[R]

		nodes := make([]pathnode, len(row))

		for j, ele := range row {

			var x int
			// choose correct parent to link up
			// first element of row
			if j == 0 {
				x = j
				// last element of row
			} else if j == len(row)-1 || (paths[j].sum < paths[j-1].sum) {
				x = j - 1
				// right parent
			} else {
				x = j
			}

			nodes[j] = _condense(paths[x], ele)
		}

		// move accumulated paths to buffer
		for i, node := range nodes {
			paths[i] = node
		}

	}
	return _largest_path_of_node(paths)
}

func _largest_path_of_node(paths []pathnode) (p []int, sum int) {
	idx := 0
	max_sum := 0

	// pick the index of the path of the greatest sum
	for i, n := range paths {
		if n.sum > max_sum {
			max_sum = n.sum
			idx = i
		}
	}

	// give the path
	return paths[idx].path, max_sum
}

func _condense(parent pathnode, ele int) pathnode {

	child := pathnode{}

	// deep copy to avoid duplicating slice references
	child.path = make([]int, len(parent.path)+1)
	for i, e := range parent.path {
		child.path[i] = e
	}
	child.path[len(child.path)-1] = ele

	// propagate sums
	child.sum = parent.sum + ele

	return child
}

func main() {
	t, _ := parse_triangle(SOLUTION_STRING)
	//printrows(t)

	_, total := Largest_path(t)

	fmt.Println("Result", total)
}
