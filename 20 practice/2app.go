package main

import (
	"fmt"
)

const (
	rowA = 3
	colA = 5
	rowB = 5
	colB = 4
)

func multiplyAB(a [rowA][colA]int, b [rowB][colB]int) [rowA][colB]int {
	var c [rowA][colB]int
	for i, row := range c {
		for j, _ := range row {
			for t, valA := range a[i] {
				c[i][j] += valA * b[t][j]

			}
		}
	}
	return c
}

func main() {
	matrixA := [rowA][colA]int{
		{1, 2, 3, 6, 7},
		{4, 5, 6, 8, 9},
		{7, 8, 9, 10, 11},
	}
	matrixB := [rowB][colB]int{
		{1, 2, 3, 6},
		{4, 5, 6, 8},
		{7, 8, 9, 10},
		{1, 1, 1, 1},
		{0, 0, 0, 0},
	}
	fmt.Println(multiplyAB(matrixA, matrixB))
}
