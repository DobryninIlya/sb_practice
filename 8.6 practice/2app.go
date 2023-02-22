package main

import "fmt"

func main() {
	var day string
	fmt.Println("Дни недели")
	fmt.Println("Введите день недели в сокращенной форме")
	fmt.Scan(&day)
	switch day {
	case "пн":
		fmt.Println("Понедельник")
		fallthrough
	case "вт":
		fmt.Println("Вторник")
		fallthrough
	case "ср":
		fmt.Println("Среда")
		fallthrough
	case "чт":
		fmt.Println("Четверг")
		fallthrough
	case "пт":
		fmt.Println("Пятница")
	}
}
