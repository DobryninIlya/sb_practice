package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Проверить, что есть совпадающие числа")
	fmt.Println("Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	if a == b || b == a || c == a {
		fmt.Println("Два числа или более совпадают")
	} else {
		fmt.Println("Все числа разные")
	}

}
