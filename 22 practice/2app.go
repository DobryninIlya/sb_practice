package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func main() {
	var arr [size]int
	var input int
	rand.Seed(time.Now().UnixNano())
	for i, _ := range arr {
		arr[i] = i*10 + rand.Intn(size)
	}
	fmt.Println(arr, "\nВведите число")
	fmt.Scan(&input)
	result := -1
	min := 0
	max := size - 1
	for max >= min {
		middle := (max + min) / 2
		if input == arr[middle] {
			result = middle
			for arr[result-1] == arr[result] && result > 0 {
				result -= 1
			}
			break
		} else if arr[middle] > input {
			max = middle - 1
		} else {
			min = middle + 1
		}

	}
	fmt.Println(result)

}
