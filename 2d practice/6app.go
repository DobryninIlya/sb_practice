package main

import "fmt"

func main() {
	var ticket int
	fmt.Println("Счастливый билет.")
	fmt.Println("Введите билет!")
	fmt.Scan(&ticket)
	goldTicket := 7234
	if ticket == goldTicket {
		fmt.Println("Билет счастливый!")
		return
	}
	leftPart := ticket / 100
	rightPart := ticket % 100
	rightPart = rightPart%10*10 + rightPart/10 // Реверсируем числа из правой части
	if leftPart == rightPart {
		fmt.Println("Билет зеркальный")
	} else {
		fmt.Println("Обычный билет!")
	}
}
