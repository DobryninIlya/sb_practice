package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Println("Проверить, что одно из чисел - положительное")
	fmt.Println("Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	if a > 0 || b > 0 || c > 0 {
		fmt.Println("Хотя бы одно из введенных чисел положительное")
	} else {
		fmt.Println("Все числа отрицательные")
	}

}
