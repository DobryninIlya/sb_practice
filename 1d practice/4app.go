package main

import "fmt"

func main() {
	resultCount := 0
	var input int
	fmt.Println("Три числа.")
	fmt.Println("Введите первое число")
	fmt.Scan(&input)
	if input > +5 {
		resultCount += 1
	}
	fmt.Println("Введите второе число")
	fmt.Scan(&input)
	if input >= 5 {
		resultCount += 1
	}
	fmt.Println("Введите третье число")
	fmt.Scan(&input)
	if input >= 5 {
		resultCount += 1
	}
	if resultCount != 0 {
		fmt.Println("Среди введенных чисел", resultCount, "число(а) больше или равно 5")
	} else {
		fmt.Println("Среди введенных чисел нет чисел больше или равных 5")
	}
}
