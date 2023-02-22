package main

import (
	"fmt"
	"math/rand"
)

const n = 10

func insertSort(arr [n]int) [n]int {
	for i := 1; i < n; i++ {
		j := i
		for j > 0 && arr[j] < arr[j-1] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j -= 1
		}
	}
	return arr
}

func main() {
	arr := [n]int{}
	for i, _ := range arr {
		arr[i] = rand.Intn(n)
	}
	fmt.Println(arr)
	fmt.Println("SORTED:\n")
	fmt.Println(insertSort(arr))
}
