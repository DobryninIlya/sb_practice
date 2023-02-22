package main

import "fmt"

func main() {
	var digit1, digit2 int
	fmt.Println("Сумма двух чисел по единице")
	fmt.Println("Введите число1 и число2")
	fmt.Scan(&digit1, &digit2)
	summ := digit1 + digit2
	for ; digit1 < summ; digit1++ {
	}
	fmt.Println(digit1, summ)
}
