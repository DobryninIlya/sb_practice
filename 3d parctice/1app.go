package main

import "fmt"

func main() {
	var input int
	fmt.Println("Написание простого цикла")
	fmt.Println("Введите число")
	fmt.Scan(&input)
	for i := 0; i <= input; i++ {
		fmt.Println(i)
	}
}
