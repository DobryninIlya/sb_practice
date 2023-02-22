package main

import "fmt"

func main() {
	flag := false
	var input int
	fmt.Println("Три числа.")
	fmt.Println("Введите первое число")
	fmt.Scan(&input)
	if input > 5 {
		flag = true
	}
	fmt.Println("Введите второе число")
	fmt.Scan(&input)
	if input > 5 {
		flag = true
	}
	fmt.Println("Введите третье число")
	fmt.Scan(&input)
	if input > 5 {
		flag = true
	}
	if flag {
		fmt.Println("Среди введенных чисел есть число больше 5")
	} else {
		fmt.Println("Среди введенных чисел нет чисел больше 5")
	}

}
