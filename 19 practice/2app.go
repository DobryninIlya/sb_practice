package main

import "fmt"

func main() {
	arr := [...]int{10, 30, 15, 2, 9, 1, 43, 28}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

	}
	fmt.Println(arr)
}
