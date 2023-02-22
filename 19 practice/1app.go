package main

import "fmt"

const (
	val1 = 4
	val2 = 5
)

func main() {
	var sumArr [val1 + val2]int
	arr1 := [val1]int{1, 2, 3, 4}
	arr2 := [val2]int{5, 6, 7, 8, 9}
	for i, val := range arr1 {
		sumArr[i] = val
	}
	for i, val := range arr2 {
		sumArr[i+val1] = val
	}
	fmt.Println(sumArr)
}
