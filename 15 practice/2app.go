package main

import "fmt"

func main() {
	var reversed [10]int
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i, val := range arr {
		reversed[9-i] = val
	}
	fmt.Println(reversed)
}
