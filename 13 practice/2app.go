package main

import "fmt"

func swap(i1, i2 *int) {
	c := *i1
	*i1 = *i2
	*i2 = c
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	swap(&a, &b)
	fmt.Println(a, b)
}
