package main

import "fmt"

func divide(slice []int) (even []int, odd []int) {
	for _, val := range slice {
		if val%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}
	return
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(divide(numbers))
}
