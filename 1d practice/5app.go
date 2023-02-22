package main

import "fmt"

func main() {
	var dayNum int
	var guestCount int
	var total int
	fmt.Println("Ресторан.")
	fmt.Println("Введите день недели:")
	fmt.Scan(&dayNum)
	fmt.Println("Введите число гостей:")
	fmt.Scan(&guestCount)
	fmt.Println("Введите сумму чека: ")
	fmt.Scan(&total)
	if dayNum == 1 {
		discount := total / 10
		total -= discount
		fmt.Println("Понедельник - день тяжелый! - Держите скидку 10%!:", discount)
	} else if dayNum == 5 && total > 10000 {
		discount := total / 20
		total -= discount
		fmt.Println("Скидка по пятницам:", discount)
	}
	if guestCount > 5 {
		servicePay := total / 10
		total += servicePay
		fmt.Println("Надбавка за обслуживание: ", servicePay)
	}
	fmt.Println("Сумма к оплате:", total)

}
