package main

import "fmt"

func summ(i1, i2 int) {
	result := 0
	for i := i1; i <= i2; i++ {
		if i%2 == 0 {
			result += i
		}
	}
	fmt.Println(result)
}

func main() {
	summ(5, 10)
}
