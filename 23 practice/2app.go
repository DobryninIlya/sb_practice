package main

import (
	"fmt"
	"math/rand"
)

const size = 15

func main() {
	anonymus := func(arr *[size]int) {
		for i := 1; i < size; i++ {
			for j := size - 1; j >= i; j-- {
				if arr[j-1] < arr[j] {
					arr[j], arr[j-1] = arr[j-1], arr[j]
				}
			}
		}
		return
	}
	arr := [size]int{}
	for i, _ := range arr {
		arr[i] = rand.Intn(size)
	}
	fmt.Println(arr)
	fmt.Println("SORTED:\n")
	anonymus(&arr)
	fmt.Println(arr)
}
