package main

import "fmt"

func main() {
	var month string
	fmt.Println("Времена года")
	fmt.Println("Введите месяц")
	fmt.Scan(&month)
	switch month {
	case "декабрь", "январь", "февраль":
		fmt.Println("зима")
	case "март", "апрель", "май":
		fmt.Println("весна")
	case "июнь", "июль", "август":
		fmt.Println("зима")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("зима")
	default:
		fmt.Println("Вы ввели некорректное значение.")
	}
}
