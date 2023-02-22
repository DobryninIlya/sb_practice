package main

import "fmt"

func incrementation(n int) (b int) {
	b = n + 5
	return
}

func multiplication(n int) (a int) {
	a = n * 5
	return
}

func main() {
	var input int
	fmt.Scan(&input)
	fmt.Println(incrementation(multiplication(input)))
}
